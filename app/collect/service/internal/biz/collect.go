package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)


type Collect struct {
	Hello string
}


type CollectRepo interface {
}


type CollectUsecase struct {
	repo CollectRepo
	log  *log.Helper
}


func NewCollectUsecase(repo CollectRepo, logger log.Logger) *CollectUsecase {
	return &CollectUsecase{repo: repo, log: log.NewHelper(logger)}
}
