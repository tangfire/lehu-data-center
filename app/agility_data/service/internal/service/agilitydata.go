package service

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"lehu-data-center/app/agility_data/service/internal/biz"
	"time"

	pb "lehu-data-center/api/agility_data/service/v1"
)

type AgilityDataService struct {
	pb.UnimplementedAgilityDataServer

	uc *biz.AgilityUsecase
}

func NewAgilityDataService(uc *biz.AgilityUsecase) *AgilityDataService {
	return &AgilityDataService{uc: uc}
}

func (s *AgilityDataService) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	// 转换参数
	params := make(map[string]interface{})
	for k, v := range req.Params {
		params[k] = v
	}

	results, total, err := s.uc.List(ctx, req.Sql, req.DataSourceName, params, int(req.Page), int(req.PageSize))
	if err != nil {
		return &pb.ListResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		}, nil
	}

	// 转换结果
	records := make([]*pb.Record, 0, len(results))
	for _, result := range results {
		record := &pb.Record{
			Fields: make(map[string]*pb.Value),
		}
		for k, v := range result {
			record.Fields[k] = convertToProtoValue(v)
		}
		records = append(records, record)
	}

	// 计算总页数
	totalPages := int32(0)
	if req.PageSize > 0 {
		totalPages = int32(total) / req.PageSize
		if int32(total)%req.PageSize != 0 {
			totalPages++
		}
	}

	return &pb.ListResponse{
		Records:    records,
		TotalCount: int32(total),
		Page:       req.Page,
		PageSize:   req.PageSize,
		TotalPages: totalPages,
		Success:    true,
	}, nil
}

func (s *AgilityDataService) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	params := make(map[string]interface{})
	for k, v := range req.Params {
		params[k] = v
	}

	result, err := s.uc.Get(ctx, req.Sql, req.DataSourceName, params)
	if err != nil {
		return &pb.GetResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		}, nil
	}

	if result == nil {
		return &pb.GetResponse{
			Success: true,
		}, nil
	}

	record := &pb.Record{
		Fields: make(map[string]*pb.Value),
	}
	for k, v := range result {
		record.Fields[k] = convertToProtoValue(v)
	}

	return &pb.GetResponse{
		Record:  record,
		Success: true,
	}, nil
}

func (s *AgilityDataService) Execute(ctx context.Context, req *pb.ExecuteRequest) (*pb.ExecuteResponse, error) {
	params := make(map[string]interface{})
	for k, v := range req.Params {
		params[k] = v
	}

	rowsAffected, err := s.uc.Execute(ctx, req.Sql, req.DataSourceName, params)
	if err != nil {
		return &pb.ExecuteResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		}, nil
	}

	return &pb.ExecuteResponse{
		RowsAffected: rowsAffected,
		Success:      true,
	}, nil
}

func (s *AgilityDataService) BatchExecute(ctx context.Context, req *pb.BatchExecuteRequest) (*pb.BatchExecuteResponse, error) {
	paramsList := make([]map[string]interface{}, 0, len(req.ParamsList))
	for _, paramMap := range req.ParamsList {
		params := make(map[string]interface{})
		for k, v := range paramMap.Params {
			params[k] = v
		}
		paramsList = append(paramsList, params)
	}

	rowsAffectedList, err := s.uc.BatchExecute(ctx, req.Sql, req.DataSourceName, paramsList)
	if err != nil {
		return &pb.BatchExecuteResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		}, nil
	}

	return &pb.BatchExecuteResponse{
		RowsAffectedList: rowsAffectedList,
		Success:          true,
	}, nil
}

func (s *AgilityDataService) TestConnection(ctx context.Context, req *pb.TestConnectionRequest) (*pb.TestConnectionResponse, error) {
	startTime := time.Now()

	err := s.uc.TestConnection(ctx, req.DataSourceName)
	responseTime := time.Since(startTime).Milliseconds()

	if err != nil {
		return &pb.TestConnectionResponse{
			Connected:      false,
			Message:        err.Error(),
			ResponseTimeMs: responseTime,
		}, nil
	}

	return &pb.TestConnectionResponse{
		Connected:      true,
		Message:        "Connection successful",
		ResponseTimeMs: responseTime,
	}, nil
}

// 辅助函数：将interface{}转换为proto Value
func convertToProtoValue(v interface{}) *pb.Value {
	if v == nil {
		return &pb.Value{
			Value: &pb.Value_NullValue{NullValue: pb.NullValue_NULL_VALUE},
		}
	}

	switch val := v.(type) {
	case string:
		return &pb.Value{Value: &pb.Value_StringValue{StringValue: val}}
	case int, int8, int16, int32, int64:
		return &pb.Value{Value: &pb.Value_IntValue{IntValue: convertToInt64(val)}}
	case uint, uint8, uint16, uint32, uint64:
		return &pb.Value{Value: &pb.Value_IntValue{IntValue: convertToInt64(val)}}
	case float32, float64:
		return &pb.Value{Value: &pb.Value_DoubleValue{DoubleValue: convertToFloat64(val)}}
	case bool:
		return &pb.Value{Value: &pb.Value_BoolValue{BoolValue: val}}
	case []byte:
		return &pb.Value{Value: &pb.Value_BytesValue{BytesValue: val}}
	case time.Time:
		return &pb.Value{Value: &pb.Value_TimestampValue{TimestampValue: timestamppb.New(val)}}
	default:
		// 尝试转换为字符串
		return &pb.Value{Value: &pb.Value_StringValue{StringValue: fmt.Sprintf("%v", val)}}
	}
}

func convertToInt64(v interface{}) int64 {
	switch val := v.(type) {
	case int:
		return int64(val)
	case int8:
		return int64(val)
	case int16:
		return int64(val)
	case int32:
		return int64(val)
	case int64:
		return val
	case uint:
		return int64(val)
	case uint8:
		return int64(val)
	case uint16:
		return int64(val)
	case uint32:
		return int64(val)
	case uint64:
		return int64(val)
	default:
		return 0
	}
}

func convertToFloat64(v interface{}) float64 {
	switch val := v.(type) {
	case float32:
		return float64(val)
	case float64:
		return val
	default:
		return 0
	}
}
