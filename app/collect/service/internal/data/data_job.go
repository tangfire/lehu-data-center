package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "lehu-data-center/api/id_generator/service/v1"
	"lehu-data-center/app/collect/service/internal/biz"
)

type dataJobRepo struct {
	data *Data
	log  *log.Helper
}

func NewJobRepo(data *Data, logger log.Logger) biz.DataJobRepo {
	return &dataJobRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *dataJobRepo) GetUid(ctx context.Context, idGenerateType string) (int64, error) {
	resp, err := r.data.idGeneratorClient.GenerateId(ctx, &v1.GenerateIdReq{Type: idGenerateType})
	if err != nil {
		r.log.WithContext(ctx).Errorf("GetUid|GenerateId fail,data:%+v,err:%v", idGenerateType, err)
		return 0, err
	}
	return resp.Id, nil
}
