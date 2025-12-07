package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"lehu-data-center/app/collect/service/internal/pkg/funcs"
	"lehu-data-center/app/collect/service/internal/types/transfers"

	"lehu-data-center/app/collect/service/internal/enums"
	"lehu-data-center/app/collect/service/internal/pkg/timeutil"
	"time"
)

type DataJobRepo interface {
	GetUid(context.Context, string) (int64, error)
}

type DataJobUsecase struct {
	messageConsumerRecordRepo          MessageConsumerRecordRepo
	messageProducerRecordRepo          MessageProducerRecordRepo
	videoBusinessMessageConsumerRecord VideoBusinessMessageConsumerRecordRepo
	videoBusinessMessageProducerRecord VideoBusinessMessageProducerRecordRepo
	dataJobRepo                        DataJobRepo
	dimensionGatherUsecase             DimensionGatherUsecase
	ruleUsecase                        RuleUsecase
	log                                *log.Helper
}

func NewDataJobUsecase(messageConsumerRecordRepo MessageConsumerRecordRepo,
	messageProducerRecordRepo MessageProducerRecordRepo,
	videoBusinessMessageConsumerRecord VideoBusinessMessageConsumerRecordRepo,
	videoBusinessMessageProducerRecord VideoBusinessMessageProducerRecordRepo,
	dataJobRepo DataJobRepo,
	dimensionGatherUsecase DimensionGatherUsecase,
	logger log.Logger) *DataJobUsecase {
	return &DataJobUsecase{messageConsumerRecordRepo: messageConsumerRecordRepo,
		messageProducerRecordRepo:          messageProducerRecordRepo,
		videoBusinessMessageConsumerRecord: videoBusinessMessageConsumerRecord,
		videoBusinessMessageProducerRecord: videoBusinessMessageProducerRecord,
		dataJobRepo:                        dataJobRepo,
		dimensionGatherUsecase:             dimensionGatherUsecase,
		log:                                log.NewHelper(logger)}
}

func (uc *DataJobUsecase) GetIdGenerateType() string {
	return "sonyflake"
}

func (uc *DataJobUsecase) Job(ctx context.Context, ruleId int64) error {
	// 1. 删除今天的消息
	err := uc.deleteTodyMessageRecord(ctx)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("Job|deleteTodyMessageRecord fail,data:%+v,err:%v", ruleId, err)
		return err
	}

	// 2. 生成父级链路Id
	messageParentTraceId, err := uc.dataJobRepo.GetUid(ctx, uc.GetIdGenerateType())
	if err != nil {
		uc.log.WithContext(ctx).Errorf("Job|GetUid fail,data:%+v,err:%v", ruleId, err)
		return err
	}

	// 3. 获取昨天的日期
	yesterdayDate := timeutil.AddDay(timeutil.Now(), -1)

	// 4. 根据时间维度创建请求时间(按天维度)
	requestTime, err := funcs.CreateRequestTime(enums.DateTypeDay, yesterdayDate)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("Job|CreateRequestTime fail, ruleId:%d, err:%v", ruleId, err)
		return err
	}
	// 5. 获取维度数据
	dimensionTransferList, err := uc.dimensionGatherUsecase.HandleDimensionData(ctx, ruleId, requestTime)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("Job|HandleDimensionData fail,ruleId:%d,err:%v", ruleId, err)
		return err
	}

	// 6. 处理每个维度
	for _, dimensionTransfer := range dimensionTransferList {
		// 生成消息链路ID
		messageTraceId, err := uc.dataJobRepo.GetUid(ctx, uc.GetIdGenerateType())
		if err != nil {
			uc.log.WithContext(ctx).Errorf("Job|GetUid for traceId fail,ruleId:%d,err:%v", ruleId, err)
			continue
		}

		// 7. 构建参数并处理规则
		paramTransfers := &transfers.ParamTransfers{
			RuleId:               ruleId,
			RuleType:             enums.RuleTypeGather,
			VideoDimensionType:   enums.VideoDimensionTypeVideo,
			DimensionTransfers:   dimensionTransfer,
			RequestTime:          requestTime,
			CollectType:          enums.CollectTypeSQL,
			MessageParentTraceId: messageParentTraceId,
			MessageTraceId:       messageTraceId,
		}

		// 8. 调用规则处理器
		_, err = uc.ruleUsecase.Handle(ctx, paramTransfers)
		if err != nil {
			uc.log.WithContext(ctx).Errorf("Job|RuleHandler.Handle fail,ruleId:%d,err:%v", ruleId, err)
			continue
		}
	}

	uc.log.WithContext(ctx).Infof("Job completed successfully, ruleId:%d, processed %d dimensions",
		ruleId, len(dimensionTransferList))
	return nil

}

func (uc *DataJobUsecase) deleteTodyMessageRecord(ctx context.Context) error {
	currentTime := time.Now().Format(time.DateOnly)
	_, _ = uc.messageConsumerRecordRepo.DeleteByConsumerTime(currentTime)
	_, _ = uc.messageProducerRecordRepo.DeleteBySendTime(currentTime)
	_, _ = uc.videoBusinessMessageConsumerRecord.DeleteByCreateTime(currentTime)
	_, _ = uc.videoBusinessMessageProducerRecord.DeleteByCreateTime(currentTime)
	return nil
}
