package data

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	consulAPI "github.com/hashicorp/consul/api"
	agilitydatav1 "lehu-data-center/api/agility_data/service/v1"
	idGeneratorv1 "lehu-data-center/api/id_generator/service/v1"
	"lehu-data-center/app/collect/service/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData,
	NewDiscovery,
	NewIdGeneratorClient,
	NewDB,
	NewVideoBusinessMessageConsumerRecordRepo,
	NewVideoBusinessMessageProducerRecordRepo,
	NewMessageProducerRecordRepo,
	NewMessageConsumerRecordRepo,
	NewJobRepo,
	NewCollectRepo,
	NewAgilityDataRepo,
	NewDimensionGatherRepo,
	NewRuleRepo,
	NewAgilityDataClient,
	NewMetricRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db                *gorm.DB
	log               *log.Helper
	idGeneratorClient idGeneratorv1.IdGeneratorClient
	agilityDataClient agilitydatav1.AgilityDataClient
}

// NewData .
func NewData(c *conf.Data, db *gorm.DB, idGeneratorClient idGeneratorv1.IdGeneratorClient, agilityDataClient agilitydatav1.AgilityDataClient, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)
	cleanup := func() {
		log.Info("closing the data resources")
	}
	return &Data{db: db, log: log, idGeneratorClient: idGeneratorClient, agilityDataClient: agilityDataClient}, cleanup, nil
}

func NewDB(conf *conf.Data, logger log.Logger) *gorm.DB {
	log := log.NewHelper(logger)
	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	return db
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewIdGeneratorClient(r registry.Discovery) idGeneratorv1.IdGeneratorClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///lehu-data-center-id_generator"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := idGeneratorv1.NewIdGeneratorClient(conn)
	return c
}

func NewAgilityDataClient(r registry.Discovery) agilitydatav1.AgilityDataClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///lehu-data-center-agility_data"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := agilitydatav1.NewAgilityDataClient(conn)
	return c
}
