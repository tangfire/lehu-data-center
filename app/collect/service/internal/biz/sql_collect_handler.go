package biz

import (
	"context"
	"fmt"
	pb "lehu-data-center/api/agility_data/service/v1"
	"lehu-data-center/app/collect/service/internal/entity"
	"lehu-data-center/app/collect/service/internal/enums"
	"lehu-data-center/app/collect/service/internal/types/transfers"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
)

type SqlCollectHandler struct {
	agilityDataRepo AgilityDataRepo
	log             *log.Helper
}

func NewSqlCollectHandler(agilityDataRepo AgilityDataRepo, logger log.Logger) *SqlCollectHandler {
	return &SqlCollectHandler{
		agilityDataRepo: agilityDataRepo,
		log:             log.NewHelper(logger),
	}
}

func (h *SqlCollectHandler) CollectType() enums.CollectType {
	return enums.CollectTypeSQL
}

func (h *SqlCollectHandler) DoCollect(
	ctx context.Context,
	dimensionGather *entity.DimensionGather,
	requestTime *transfers.RequestTime,
) ([]*transfers.DimensionTransfer, error) {
	h.log.WithContext(ctx).Infof(
		"开始SQL采集: ruleId=%d, collectType=%d, source=%s",
		dimensionGather.RuleId,
		dimensionGather.CollectType,
		dimensionGather.CollectSourceName,
	)

	// 构建SQL参数
	params := make(map[string]string)
	if requestTime.StartTime != nil && requestTime.EndTime != nil {
		params["start_time"] = requestTime.StartTime.Format("2006-01-02 15:04:05")
		params["end_time"] = requestTime.EndTime.Format("2006-01-02 15:04:05")
	}

	// 通过敏捷数据仓库执行SQL查询
	req := &pb.ListRequest{
		Sql:            dimensionGather.CollectDetail,
		DataSourceName: dimensionGather.CollectSourceName,
		Params:         params,
		Page:           1,
		PageSize:       10000,
	}

	resp, err := h.agilityDataRepo.List(ctx, req)
	if err != nil {
		h.log.WithContext(ctx).Errorf("调用敏捷数据服务失败: %v", err)
		return nil, fmt.Errorf("query video dimensions failed: %w", err)
	}

	if !resp.Success {
		h.log.WithContext(ctx).Errorf("SQL执行失败: %s", resp.ErrorMessage)
		return nil, fmt.Errorf("sql execution failed: %s", resp.ErrorMessage)
	}

	h.log.WithContext(ctx).Infof("SQL采集完成: 获取到%d条记录", len(resp.Records))

	if len(resp.Records) == 0 {
		return []*transfers.DimensionTransfer{}, nil
	}

	return h.fillVideoDimension(ctx, resp.Records)
}

func (h *SqlCollectHandler) fillVideoDimension(
	ctx context.Context,
	records []*pb.Record,
) ([]*transfers.DimensionTransfer, error) {
	// 使用map分组，key为"videoTypeId-parentVideoTypeId"
	groupMap := make(map[string]*transfers.DimensionTransfer)

	for _, record := range records {
		// 解析视频维度数据
		videoDimension, err := h.parseVideoDimension(record)
		if err != nil {
			h.log.WithContext(ctx).Warnf("解析视频维度记录失败: %v, 跳过该记录", err)
			continue
		}

		// 构建分组key
		key := fmt.Sprintf("%d-%d", videoDimension.VideoTypeID, videoDimension.ParentVideoTypeID)

		// 如果分组不存在，创建新的维度传输对象
		if _, exists := groupMap[key]; !exists {
			groupMap[key] = &transfers.DimensionTransfer{
				VideoTypeId:       videoDimension.VideoTypeID,
				ParentVideoTypeId: videoDimension.ParentVideoTypeID,
				VideoIdList:       []int64{},
			}
		}

		// 添加视频ID到列表
		groupMap[key].VideoIdList = append(groupMap[key].VideoIdList, videoDimension.VideoID)
	}

	// 将map转换为slice
	dimensionTransferList := make([]*transfers.DimensionTransfer, 0, len(groupMap))
	for _, transfer := range groupMap {
		dimensionTransferList = append(dimensionTransferList, transfer)
	}

	h.log.WithContext(ctx).Infof("填充视频维度完成: 生成%d个维度传输对象", len(dimensionTransferList))
	return dimensionTransferList, nil
}

// parseVideoDimension 从Record中解析视频维度数据
func (h *SqlCollectHandler) parseVideoDimension(record *pb.Record) (*VideoDimension, error) {
	dimension := &VideoDimension{}

	// 解析video_id
	videoID, err := h.extractInt64(record, "video_id")
	if err != nil {
		return nil, fmt.Errorf("parse video_id failed: %w", err)
	}
	dimension.VideoID = videoID

	// 解析video_type_id
	videoTypeID, err := h.extractInt64(record, "video_type_id")
	if err != nil {
		return nil, fmt.Errorf("parse video_type_id failed: %w", err)
	}
	dimension.VideoTypeID = videoTypeID

	// 解析parent_video_type_id
	parentVideoTypeID, err := h.extractInt64(record, "parent_video_type_id")
	if err != nil {
		// 如果parent_video_type_id不存在，尝试使用video_type_id作为默认值
		dimension.ParentVideoTypeID = videoTypeID
	} else {
		dimension.ParentVideoTypeID = parentVideoTypeID
	}

	return dimension, nil
}

// extractInt64 从Record中提取int64字段
func (h *SqlCollectHandler) extractInt64(record *pb.Record, fieldName string) (int64, error) {
	value, ok := record.Fields[fieldName]
	if !ok {
		return 0, fmt.Errorf("field %s not found", fieldName)
	}

	switch v := value.Value.(type) {
	case *pb.Value_IntValue:
		return v.IntValue, nil
	case *pb.Value_StringValue:
		// 尝试将字符串转换为int64
		result, err := strconv.ParseInt(v.StringValue, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("parse string to int64 failed: %s", v.StringValue)
		}
		return result, nil
	case *pb.Value_DoubleValue:
		return int64(v.DoubleValue), nil
	case *pb.Value_BoolValue:
		if v.BoolValue {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("unsupported field type for %s", fieldName)
	}
}
