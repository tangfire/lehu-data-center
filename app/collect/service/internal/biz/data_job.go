package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"lehu-data-center/app/collect/service/internal/enums"
	"lehu-data-center/app/collect/service/internal/pkg/datetime"
	"lehu-data-center/app/collect/service/internal/pkg/timeutil"
	"time"
)

// RequestTime 请求时间
type RequestTime struct {
	GatherDate time.Time      `json:"gather_date"`
	DateType   enums.DateType `json:"date_type"`
	StartTime  *time.Time     `json:"start_time"`
	EndTime    *time.Time     `json:"end_time"`
}

type DataJobRepo interface {
	GetUid(context.Context, string) (int64, error)
}

type DataJobUsecase struct {
	messageConsumerRecordRepo          MessageConsumerRecordRepo
	messageProducerRecordRepo          MessageProducerRecordRepo
	videoBusinessMessageConsumerRecord VideoBusinessMessageConsumerRecordRepo
	videoBusinessMessageProducerRecord VideoBusinessMessageProducerRecordRepo
	dataJobRepo                        DataJobRepo
	log                                *log.Helper
}

func NewDataJobUsecase(messageConsumerRecordRepo MessageConsumerRecordRepo,
	messageProducerRecordRepo MessageProducerRecordRepo,
	videoBusinessMessageConsumerRecord VideoBusinessMessageConsumerRecordRepo,
	videoBusinessMessageProducerRecord VideoBusinessMessageProducerRecordRepo,
	dataJobRepo DataJobRepo,
	logger log.Logger) *DataJobUsecase {
	return &DataJobUsecase{messageConsumerRecordRepo: messageConsumerRecordRepo,
		messageProducerRecordRepo:          messageProducerRecordRepo,
		videoBusinessMessageConsumerRecord: videoBusinessMessageConsumerRecord,
		videoBusinessMessageProducerRecord: videoBusinessMessageProducerRecord,
		dataJobRepo:                        dataJobRepo,
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
	requestTime, err := datetime.CreateRequestTime(enums.DateTypeDay, yesterdayDate)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("Job|CreateRequestTime fail, ruleId:%d, err:%v", ruleId, err)
		return err
	}

	fmt.Println(messageParentTraceId)

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
