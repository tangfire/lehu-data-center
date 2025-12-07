// file: internal/biz/agility_data_repo.go
package biz

import (
	"context"
	pb "lehu-data-center/api/agility_data/service/v1"
)

// VideoDimension 视频维度业务对象
type VideoDimension struct {
	VideoID           int64
	VideoTypeID       int64
	ParentVideoTypeID int64
	UserID            int64 // 如果需要
}

// AgilityDataRepo 敏捷数据仓库接口
type AgilityDataRepo interface {
	List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error)
	Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error)
	Execute(ctx context.Context, req *pb.ExecuteRequest) (*pb.ExecuteResponse, error)
	BatchExecute(ctx context.Context, req *pb.BatchExecuteRequest) (*pb.BatchExecuteResponse, error)
	TestConnection(ctx context.Context, req *pb.TestConnectionRequest) (*pb.TestConnectionResponse, error)
}
