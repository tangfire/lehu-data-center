package service

import (
	"context"
	"lehu-data-center/app/id_generator/service/internal/biz"

	pb "lehu-data-center/api/id_generator/service/v1"
)

type IdGeneratorService struct {
	pb.UnimplementedIdGeneratorServer

	uc *biz.IdGeneratorUsecase
}

func NewIdGeneratorService(uc *biz.IdGeneratorUsecase) *IdGeneratorService {
	return &IdGeneratorService{uc: uc}
}

func (s *IdGeneratorService) GenerateId(ctx context.Context, req *pb.GenerateIdReq) (*pb.GenerateIdResp, error) {
	return &pb.GenerateIdResp{}, nil
}
func (s *IdGeneratorService) GenerateBatch(ctx context.Context, req *pb.GenerateBatchReq) (*pb.GenerateBatchResp, error) {
	return &pb.GenerateBatchResp{}, nil
}
