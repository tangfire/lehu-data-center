package data

import (
	"lehu-data-center/app/collect/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type collectRepo struct {
	data *Data
	log  *log.Helper
}

// NewCollectRepo .
func NewCollectRepo(data *Data, logger log.Logger) biz.CollectRepo {
	return &collectRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
