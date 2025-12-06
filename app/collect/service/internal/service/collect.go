package service

import (
	"context"
	"lehu-data-center/app/collect/service/internal/biz"

	pb "lehu-data-center/api/collect/service/v1"
)

type CollectService struct {
	pb.UnimplementedCollectServer

	uc *biz.CollectUsecase
}

func NewCollectService(uc *biz.CollectUsecase) *CollectService {
	return &CollectService{uc: uc}
}

func (s *CollectService) CreateCollect(ctx context.Context, req *pb.CreateCollectRequest) (*pb.CreateCollectReply, error) {
	return &pb.CreateCollectReply{}, nil
}
func (s *CollectService) UpdateCollect(ctx context.Context, req *pb.UpdateCollectRequest) (*pb.UpdateCollectReply, error) {
	return &pb.UpdateCollectReply{}, nil
}
func (s *CollectService) DeleteCollect(ctx context.Context, req *pb.DeleteCollectRequest) (*pb.DeleteCollectReply, error) {
	return &pb.DeleteCollectReply{}, nil
}
func (s *CollectService) GetCollect(ctx context.Context, req *pb.GetCollectRequest) (*pb.GetCollectReply, error) {
	return &pb.GetCollectReply{}, nil
}
func (s *CollectService) ListCollect(ctx context.Context, req *pb.ListCollectRequest) (*pb.ListCollectReply, error) {
	return &pb.ListCollectReply{}, nil
}
