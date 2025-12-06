package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/sony/sonyflake"
	v1 "lehu-data-center/api/id_generator/service/v1"
	"lehu-data-center/app/id_generator/service/internal/conf"
	"time"
)

type IdGeneratorRepo interface {
}

type IdGeneratorUsecase struct {
	repo      IdGeneratorRepo
	log       *log.Helper
	snoyflake *sonyflake.Sonyflake
}

// GeneratedId Id生成结果
type GeneratedId struct {
	Id        int64
	Timestamp time.Time
	WorkerId  int64
	Sequence  int64
}

func NewIdGeneratorUsecase(repo IdGeneratorRepo, snoyflake *sonyflake.Sonyflake, logger log.Logger) *IdGeneratorUsecase {
	return &IdGeneratorUsecase{repo: repo, snoyflake: snoyflake, log: log.NewHelper(logger)}
}

func NewSonyflake(cfg *conf.Snoyflake) *sonyflake.Sonyflake {
	sf := sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		MachineID: func() (uint16, error) {
			return uint16(cfg.WorkerId), nil
		},
	})
	return sf
}

// GenerateSingle 生成单个Id
func (uc *IdGeneratorUsecase) GenerateSingle(ctx context.Context, req *v1.GenerateIdReq) (*v1.GenerateIdResp, error) {
	// todo 加入缓存

	// 缓存为空，实时生成
	id, err := uc.generateSnowflakeId(ctx)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("GenerateSingle|generateSnowflakeId fail,data:%+v,err:%v", id, err)
		return nil, err
	}
	generatedId := uc.parseSnowflakeID(id)
	return &v1.GenerateIdResp{
		Id:        generatedId.Id,
		Timestamp: generatedId.Timestamp.Unix(),
		WorkerId:  int32(generatedId.WorkerId),
		Sequence:  int32(generatedId.Sequence),
	}, nil
}

// GenerateBatch 批量生成Id
func (uc *IdGeneratorUsecase) GenerateBatch(ctx context.Context, req *v1.GenerateBatchReq) (*v1.GenerateBatchResp, error) {
	startTime := time.Now()

	// 生成Id
	idList := make([]int64, 0, req.Count)
	detailList := make([]*GeneratedId, 0, req.Count)

	for i := 0; i < int(req.Count); i++ {
		id, err := uc.generateSnowflakeId(ctx)
		if err != nil {
			return nil, err
		}
		idList = append(idList, id)
		generateId := uc.parseSnowflakeID(id)
		detailList = append(detailList, &GeneratedId{
			Id:        generateId.Id,
			Timestamp: generateId.Timestamp,
			WorkerId:  generateId.WorkerId,
			Sequence:  generateId.Sequence,
		})
	}
	var retList []*v1.IdDetail
	for _, v := range detailList {
		retList = append(retList, &v1.IdDetail{
			Id:        v.Id,
			Timestamp: v.Timestamp.Unix(),
			WorkerId:  int32(v.WorkerId),
			Sequence:  int32(v.Sequence),
		})
	}
	return &v1.GenerateBatchResp{
		Ids:            idList,
		Details:        retList,
		GenerateTimeMs: int64(time.Since(startTime)),
	}, nil
}

func (uc *IdGeneratorUsecase) generateSnowflakeId(ctx context.Context) (int64, error) {
	nextId, err := uc.snoyflake.NextID()
	if err != nil {
		uc.log.WithContext(ctx).Errorf("generateSnowflakeId|NextID fail,data:%+v,err:%v", nextId, err)
		return 0, err
	}
	return int64(nextId), nil
}

// parseSnowflakeID 解析雪花ID
func (uc *IdGeneratorUsecase) parseSnowflakeID(id int64) *GeneratedId {
	parseMap := sonyflake.Decompose(uint64(id))

	// Sonyflake 的起始时间（epoch）通常是 2014-09-01 00:00:00 +0000 UTC
	// 时间单位是 10 毫秒
	sonyflakeEpoch := int64(1409529600000) // 2014-09-01 在毫秒时间戳

	// 获取解析出的时间单位数
	timeUnits := parseMap["time"]

	// 转换为毫秒：每个时间单位是 10 毫秒
	timestampMs := int64(timeUnits) * 10

	// 转换为绝对时间戳（毫秒）
	absoluteTimestampMs := sonyflakeEpoch + timestampMs

	// 转换为 time.Time
	// 注意：time.Unix 接受秒和纳秒，所以需要转换
	timestamp := time.Unix(absoluteTimestampMs/1000, (absoluteTimestampMs%1000)*1e6)

	return &GeneratedId{
		Id:        int64(parseMap["id"]),
		Timestamp: timestamp,
		WorkerId:  int64(parseMap["machine-id"]),
		Sequence:  int64(parseMap["sequence"]),
	}
}
