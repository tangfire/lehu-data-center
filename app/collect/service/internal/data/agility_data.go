// file: internal/data/agility_data_repo.go
package data

import (
	"context"
	"fmt"
	pb "lehu-data-center/api/agility_data/service/v1"
	"lehu-data-center/app/collect/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type agilityDataRepo struct {
	data *Data
	log  *log.Helper
}

// NewAgilityDataRepo 创建敏捷数据仓库
func NewAgilityDataRepo(data *Data, logger log.Logger) biz.AgilityDataRepo {
	return &agilityDataRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *agilityDataRepo) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	return r.data.agilityDataClient.List(ctx, req)
}

func (r *agilityDataRepo) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	return r.data.agilityDataClient.Get(ctx, req)
}

func (r *agilityDataRepo) Execute(ctx context.Context, req *pb.ExecuteRequest) (*pb.ExecuteResponse, error) {
	return r.data.agilityDataClient.Execute(ctx, req)
}

func (r *agilityDataRepo) BatchExecute(ctx context.Context, req *pb.BatchExecuteRequest) (*pb.BatchExecuteResponse, error) {
	return r.data.agilityDataClient.BatchExecute(ctx, req)
}

func (r *agilityDataRepo) TestConnection(ctx context.Context, req *pb.TestConnectionRequest) (*pb.TestConnectionResponse, error) {
	return r.data.agilityDataClient.TestConnection(ctx, req)
}

// 自定义方法：执行SQL查询并返回业务对象
func (r *agilityDataRepo) QueryVideoDimensions(ctx context.Context, sql string, dataSourceName string, params map[string]string) ([]*biz.VideoDimension, error) {
	resp, err := r.data.agilityDataClient.List(ctx, &pb.ListRequest{
		Sql:            sql,
		DataSourceName: dataSourceName,
		Params:         params,
		Page:           1,
		PageSize:       10000,
	})

	if err != nil {
		return nil, fmt.Errorf("query video dimensions failed: %v", err)
	}

	if !resp.Success {
		return nil, fmt.Errorf("query failed: %s", resp.ErrorMessage)
	}

	// 转换proto记录为业务对象
	var dimensions []*biz.VideoDimension
	for _, record := range resp.Records {
		dimension, err := r.convertToVideoDimension(record)
		if err != nil {
			r.log.Errorf("convert record failed: %v", err)
			continue
		}
		dimensions = append(dimensions, dimension)
	}

	return dimensions, nil
}

func (r *agilityDataRepo) convertToVideoDimension(record *pb.Record) (*biz.VideoDimension, error) {
	dimension := &biz.VideoDimension{}

	// 解析video_id
	if videoID, ok := record.Fields["video_id"]; ok {
		if intValue, ok := videoID.Value.(*pb.Value_IntValue); ok {
			dimension.VideoID = intValue.IntValue
		}
	}

	// 解析video_type_id
	if videoTypeID, ok := record.Fields["video_type_id"]; ok {
		if intValue, ok := videoTypeID.Value.(*pb.Value_IntValue); ok {
			dimension.VideoTypeID = intValue.IntValue
		}
	}

	// 解析parent_video_type_id
	if parentVideoTypeID, ok := record.Fields["parent_video_type_id"]; ok {
		if intValue, ok := parentVideoTypeID.Value.(*pb.Value_IntValue); ok {
			dimension.ParentVideoTypeID = intValue.IntValue
		}
	}

	return dimension, nil
}
