package data

import (
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"lehu-data-center/app/agility_data/service/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewAgilityRepo,
	NewDefaultDB,
	NewDataSourceMap,
)

// Data 数据访问层
type Data struct {
	defaultDB  *gorm.DB       // 默认数据库连接
	dataSource *DataSourceMap // 数据源映射
	log        *log.Helper
}

// DataSourceMap 数据源映射，线程安全
type DataSourceMap struct {
	sources map[string]*gorm.DB
	log     *log.Helper
}

// NewDataSourceMap 创建数据源映射
func NewDataSourceMap(c *conf.Data, logger log.Logger) (*DataSourceMap, func(), error) {
	log := log.NewHelper(logger)
	sources := make(map[string]*gorm.DB)

	var cleanupFuncs []func()

	// 初始化配置中的数据源
	for _, dsConf := range c.DataSources {
		if !dsConf.IsActive {
			continue
		}

		db, cleanup, err := initDataSource(dsConf.Connection, logger)
		if err != nil {
			log.Errorf("Failed to init data source %s: %v", dsConf.Name, err)
			continue
		}

		sources[dsConf.Name] = db
		cleanupFuncs = append(cleanupFuncs, cleanup)
		log.Infof("Data source '%s' initialized", dsConf.Name)
	}

	log.Infof("Total data sources initialized: %d", len(sources))

	cleanup := func() {
		for _, f := range cleanupFuncs {
			f()
		}
		log.Info("All data sources closed")
	}

	return &DataSourceMap{
		sources: sources,
		log:     log,
	}, cleanup, nil
}

// NewDefaultDB 创建默认数据库连接
func NewDefaultDB(c *conf.Data, logger log.Logger) (*gorm.DB, func(), error) {
	log := log.NewHelper(logger)

	if c.Database == nil {
		return nil, nil, fmt.Errorf("database config is required")
	}

	db, cleanup, err := initDatabase(c.Database, logger)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect default database: %v", err)
	}

	log.Info("Default database connected")
	return db, cleanup, nil
}

// NewData 创建数据访问层
func NewData(
	defaultDB *gorm.DB,
	dataSource *DataSourceMap,
	logger log.Logger,
) (*Data, func(), error) {
	log := log.NewHelper(logger)

	cleanup := func() {
		log.Info("closing data resources")
	}

	return &Data{
		defaultDB:  defaultDB,
		dataSource: dataSource,
		log:        log,
	}, cleanup, nil
}

// initDatabase 初始化数据库连接
func initDatabase(dbConf *conf.Data_Database, logger log.Logger) (*gorm.DB, func(), error) {
	log := log.NewHelper(logger)

	var dialector gorm.Dialector
	switch dbConf.Driver {
	case "mysql":
		dialector = mysql.Open(dbConf.Source)
	default:
		return nil, nil, fmt.Errorf("unsupported database driver: %s", dbConf.Driver)
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, err
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

	cleanup := func() {
		if err := sqlDB.Close(); err != nil {
			log.Errorf("Failed to close database connection: %v", err)
		}
	}

	return db, cleanup, nil
}

// initDataSource 初始化数据源连接
func initDataSource(conn *conf.Data_DataSource_Connection, logger log.Logger) (*gorm.DB, func(), error) {
	log := log.NewHelper(logger)

	var dialector gorm.Dialector
	switch conn.Driver {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			conn.Username,
			conn.Password,
			conn.Host,
			conn.Port,
			conn.Database,
		)

		// 添加参数
		params := "?charset=utf8mb4&parseTime=True&loc=Local"
		for k, v := range conn.Parameters {
			if len(params) > 1 {
				params += "&"
			}
			params += k + "=" + v
		}

		dialector = mysql.Open(dsn + params)
	default:
		return nil, nil, fmt.Errorf("unsupported database driver: %s", conn.Driver)
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to data source: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, err
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Infof("Data source connected: %s@%s:%d/%s",
		conn.Username, conn.Host, conn.Port, conn.Database)

	cleanup := func() {
		if err := sqlDB.Close(); err != nil {
			log.Errorf("Failed to close data source connection: %v", err)
		}
	}

	return db, cleanup, nil
}

// GetDB 获取数据库连接
func (d *Data) GetDB(dataSourceName string) (*gorm.DB, error) {
	if dataSourceName == "" || dataSourceName == "default" {
		if d.defaultDB == nil {
			return nil, fmt.Errorf("default database not configured")
		}
		return d.defaultDB, nil
	}

	return d.dataSource.Get(dataSourceName)
}

// Get 从数据源映射中获取数据库连接
func (m *DataSourceMap) Get(name string) (*gorm.DB, error) {
	if m == nil {
		return nil, fmt.Errorf("data source map is nil")
	}

	db, exists := m.sources[name]
	if !exists {
		return nil, fmt.Errorf("data source '%s' not found", name)
	}

	return db, nil
}

// GetAllNames 获取所有数据源名称
func (m *DataSourceMap) GetAllNames() []string {
	if m == nil {
		return []string{}
	}

	names := make([]string, 0, len(m.sources))
	for name := range m.sources {
		names = append(names, name)
	}

	return names
}
