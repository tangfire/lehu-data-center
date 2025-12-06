package data

import (
	"lehu-data-center/app/id_generator/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type idGeneratorRepo struct {
	data *Data
	log  *log.Helper
}

// NewIdGeneratorRepo .
func NewIdGeneratorRepo(data *Data, logger log.Logger) biz.IdGeneratorRepo {
	return &idGeneratorRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
