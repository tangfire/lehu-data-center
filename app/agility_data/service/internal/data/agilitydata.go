package data

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"lehu-data-center/app/agility_data/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type agilityRepo struct {
	data *Data
	log  *log.Helper
}

// NewAgilityRepo 创建敏捷数据仓库
func NewAgilityRepo(data *Data, logger log.Logger) biz.AgilityRepo {
	return &agilityRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *agilityRepo) List(ctx context.Context, sql string, dataSourceName string, params map[string]interface{}, page, pageSize int) ([]map[string]interface{}, int64, error) {
	db, err := r.data.GetDB(dataSourceName)
	if err != nil {
		return nil, 0, fmt.Errorf("get database connection failed: %v", err)
	}

	// 构建计数SQL（移除子查询的别名，使用更兼容的方式）
	countSQL := "SELECT COUNT(*) FROM (" + sql + ") as count_table"

	var total int64
	if err := db.Raw(countSQL, params).Scan(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count query failed: %v", err)
	}

	// 添加分页
	var results []map[string]interface{}
	if page > 0 && pageSize > 0 {
		offset := (page - 1) * pageSize
		sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, pageSize, offset)
	}

	// 执行查询
	if err := db.Raw(sql, params).Scan(&results).Error; err != nil {
		return nil, 0, fmt.Errorf("query failed: %v", err)
	}

	return results, total, nil
}

func (r *agilityRepo) Get(ctx context.Context, sql string, dataSourceName string, params map[string]interface{}) (map[string]interface{}, error) {
	db, err := r.data.GetDB(dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("get database connection failed: %v", err)
	}

	var result map[string]interface{}
	if err := db.Raw(sql, params).Take(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("query failed: %v", err)
	}

	return result, nil
}

func (r *agilityRepo) Execute(ctx context.Context, sql string, dataSourceName string, params map[string]interface{}) (int64, error) {
	db, err := r.data.GetDB(dataSourceName)
	if err != nil {
		return 0, fmt.Errorf("get database connection failed: %v", err)
	}

	// 判断SQL类型
	upperSQL := strings.ToUpper(strings.TrimSpace(sql))

	// 对于INSERT/UPDATE/DELETE操作
	if strings.HasPrefix(upperSQL, "INSERT") ||
		strings.HasPrefix(upperSQL, "UPDATE") ||
		strings.HasPrefix(upperSQL, "DELETE") {

		result := db.Exec(sql, params)
		if result.Error != nil {
			return 0, fmt.Errorf("execute failed: %v", result.Error)
		}
		return result.RowsAffected, nil
	}

	// 对于SELECT查询，返回0行影响
	if strings.HasPrefix(upperSQL, "SELECT") {
		return 0, nil
	}

	return 0, fmt.Errorf("unsupported SQL statement type")
}

func (r *agilityRepo) BatchExecute(ctx context.Context, sql string, dataSourceName string, paramsList []map[string]interface{}) ([]int64, error) {
	if len(paramsList) == 0 {
		return []int64{}, nil
	}

	db, err := r.data.GetDB(dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("get database connection failed: %v", err)
	}

	var rowsAffectedList []int64

	// 开启事务
	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	// 确保事务回滚
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 批量执行
	for _, params := range paramsList {
		result := tx.Exec(sql, params)
		if result.Error != nil {
			tx.Rollback()
			return nil, fmt.Errorf("batch execute failed: %v", result.Error)
		}
		rowsAffectedList = append(rowsAffectedList, result.RowsAffected)
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("commit failed: %v", err)
	}

	return rowsAffectedList, nil
}

func (r *agilityRepo) TestConnection(ctx context.Context, dataSourceName string) error {
	db, err := r.data.GetDB(dataSourceName)
	if err != nil {
		return fmt.Errorf("get database connection failed: %v", err)
	}

	// 执行简单查询测试连接
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("get underlying sql.DB failed: %v", err)
	}

	// Ping测试连接
	if err := sqlDB.PingContext(ctx); err != nil {
		return fmt.Errorf("ping failed: %v", err)
	}

	return nil
}
