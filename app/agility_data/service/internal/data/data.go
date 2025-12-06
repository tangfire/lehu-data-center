package data

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"lehu-data-center/app/agility_data/service/internal/conf"
	"sync"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewAgilityRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db          *gorm.DB
	dataSources map[string]*DataSource
	mu          sync.RWMutex
}

// DataSource 表示一个数据源连接
type DataSource struct {
	Name   string
	DB     *gorm.DB
	Config *conf.Data_DataSource_Connection
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)
	// 初始化主数据连接
	var db *gorm.DB
	var err error

	if c.Database != nil {
		db, err = initDatabase(c.Database)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to connect database: %v", err)
		}
	}

	// 初始化数据源连接池
	dataSources := make(map[string]*DataSource)

	// 初始化配置中的数据源
	for _, dsConf := range c.DataSources {
		if dsConf.IsActive {
			dsDB, err := initDataSourceDB(dsConf.Connection, logger)
			if err != nil {
				log.Errorf("Failed to init data source %s: %v", dsConf.Name, err)
				continue
			}
			dataSources[dsConf.Name] = &DataSource{
				Name:   dsConf.Name,
				DB:     dsDB,
				Config: dsConf.Connection,
			}
			log.Infof("Data source '%s' initialized", dsConf.Name)
		}
	}

	log.Infof("Total data sources initialized: %d", len(dataSources))

	cleanup := func() {
		log.Info("closing the data resources")

		// 关闭主数据库连接
		if db != nil {
			sqlDB, err := db.DB()
			if err == nil {
				sqlDB.Close()
				log.Info("main database connection closed")
			}
		}

		// 关闭所有数据源连接
		for name, ds := range dataSources {
			if ds.DB != nil {
				sqlDB, err := ds.DB.DB()
				if err == nil {
					sqlDB.Close()
					log.Infof("data source '%s' connection closed", name)
				}
			}
		}
	}

	return &Data{
		db:          db,
		dataSources: dataSources,
	}, cleanup, nil
}

func initDatabase(dbConf *conf.Data_Database) (*gorm.DB, error) {
	var dialector gorm.Dialector

	switch dbConf.Driver {
	case "mysql":
		dialector = mysql.Open(dbConf.Source)
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", dbConf.Driver)
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 配置连接池
	if dbConf.MaxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(int(dbConf.MaxIdleConns))
	}
	if dbConf.MaxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(int(dbConf.MaxOpenConns))
	}
	if dbConf.ConnMaxLifetime != nil {
		sqlDB.SetConnMaxLifetime(dbConf.ConnMaxLifetime.AsDuration())
	}

	return db, nil
}

// 为数据源初始化数据库连接
func initDataSourceDB(conn *conf.Data_DataSource_Connection, l log.Logger) (*gorm.DB, error) {
	var dialector gorm.Dialector
	var dsn string

	switch conn.Driver {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			conn.Username, conn.Password, conn.Host, conn.Port, conn.Database)
		dialector = mysql.Open(dsn)
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", conn.Driver)
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to data source: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

//func initDatabase(dbConf *conf.Data_Database) (*gorm.DB, error) {
//	var dialector gorm.Dialector
//
//	switch dbConf.Driver {
//	case "mysql":
//		dialector = mysql.Open(dbConf.Source)
//	default:
//		return nil, fmt.Errorf("unsupported database driver: %s", dbConf.Driver)
//	}
//
//	db, err := gorm.Open(dialector, &gorm.Config{})
//	if err != nil {
//		return nil, err
//	}
//
//	sqlDB, err := db.DB()
//	if err != nil {
//		return nil, err
//	}
//
//	// 配置连接池
//	if dbConf.MaxIdleConns > 0 {
//		sqlDB.SetMaxIdleConns(int(dbConf.MaxIdleConns))
//	}
//	if dbConf.MaxOpenConns > 0 {
//		sqlDB.SetMaxOpenConns(int(dbConf.MaxOpenConns))
//	}
//	if dbConf.ConnMaxLifetime != nil {
//		sqlDB.SetConnMaxLifetime(dbConf.ConnMaxLifetime.AsDuration())
//	}
//
//	return db, nil
//}
//
//// initDataSourceDB 为数据源初始化数据库连接
//func initDataSourceDB(conn *conf.Data_DataSource_Connection, logger log.Logger) (*gorm.DB, error) {
//	var dialector gorm.Dialector
//	var dsn string
//
//	switch conn.Driver {
//	case "mysql":
//		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
//			conn.Username, conn.Password, conn.Host, conn.Port, conn.Database)
//		dialector = mysql.Open(dsn)
//	default:
//		return nil, fmt.Errorf("unsupported database driver: %s", conn.Driver)
//	}
//
//	db, err := gorm.Open(dialector, &gorm.Config{})
//	if err != nil {
//		return nil, fmt.Errorf("failed to connect to data source: %v", err)
//	}
//
//	sqlDB, err := db.DB()
//	if err != nil {
//		return nil, err
//	}
//
//	// 设置连接池参数
//	sqlDB.SetMaxIdleConns(10)
//	sqlDB.SetMaxOpenConns(100)
//	sqlDB.SetConnMaxLifetime(time.Hour)
//
//	log.NewHelper(logger).Infof("Data source connected: %s@%s:%d/%s",
//		conn.Username, conn.Host, conn.Port, conn.Database)
//
//	return db, nil
//}

// GetDataSource 获取数据源
func (d *Data) GetDataSource(name string) (*DataSource, bool) {
	if name == "" {
		return nil, false
	}

	d.mu.RLock()
	defer d.mu.RUnlock()

	ds, ok := d.dataSources[name]
	return ds, ok
}

// GetDefaultDB 获取默认数据库连接
func (d *Data) GetDefaultDB() *gorm.DB {
	return d.db
}
