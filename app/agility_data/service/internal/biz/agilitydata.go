package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// Agility is a Agility model.
type Agility struct {
	Hello string
}

// AgilityRepo is a Greater repo.
type AgilityRepo interface {
	// List 查询列表数据
	List(ctx context.Context, sql string, dataSourceName string, params map[string]interface{}, page, pageSize int) ([]map[string]interface{}, int64, error)
	// Get 查询单条数据
	Get(ctx context.Context, sql string, dataSourceName string, params map[string]interface{}) (map[string]interface{}, error)
	// Execute 执行SQL（INSERT/UPDATE/DELETE）
	Execute(ctx context.Context, sql string, dataSourceName string, params map[string]interface{}) (int64, error)
	// BatchExecute 批量执行SQL
	BatchExecute(ctx context.Context, sql string, dataSourceName string, paramsList []map[string]interface{}) ([]int64, error)
	// TestConnection 测试数据源连接
	TestConnection(ctx context.Context, dataSourceName string) error
}

// AgilityUsecase is a Agility usecase.
type AgilityUsecase struct {
	repo AgilityRepo
	log  *log.Helper
}

// NewAgilityUsecase new a Agility usecase.
func NewAgilityUsecase(repo AgilityRepo, logger log.Logger) *AgilityUsecase {
	return &AgilityUsecase{repo: repo, log: log.NewHelper(logger)}
}

// List 查询列表数据
func (uc *AgilityUsecase) List(ctx context.Context, sql string, dataSourceName string, params map[string]interface{}, page, pageSize int) ([]map[string]interface{}, int64, error) {
	// 可以在这里添加业务逻辑，比如参数校验、权限检查等
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 1000 {
		pageSize = 1000 // 限制最大页大小
	}

	uc.log.WithContext(ctx).Infof("List query: dataSource=%s, sql=%s", dataSourceName, sql)

	return uc.repo.List(ctx, sql, dataSourceName, params, page, pageSize)
}

// Get 查询单条数据
func (uc *AgilityUsecase) Get(ctx context.Context, sql string, dataSourceName string, params map[string]interface{}) (map[string]interface{}, error) {
	// 业务逻辑：限制查询结果数量
	sql = sql + " LIMIT 1"

	uc.log.WithContext(ctx).Infof("Get query: dataSource=%s, sql=%s", dataSourceName, sql)

	return uc.repo.Get(ctx, sql, dataSourceName, params)
}

// Execute 执行SQL
func (uc *AgilityUsecase) Execute(ctx context.Context, sql string, dataSourceName string, params map[string]interface{}) (int64, error) {
	// 业务逻辑：安全检查，防止恶意SQL
	//if !isSafeSQL(sql) {
	//	return 0, NewError("unsafe SQL detected")
	//}

	uc.log.WithContext(ctx).Infof("Execute SQL: dataSource=%s, sql=%s", dataSourceName, sql)

	return uc.repo.Execute(ctx, sql, dataSourceName, params)
}

// BatchExecute 批量执行SQL
func (uc *AgilityUsecase) BatchExecute(ctx context.Context, sql string, dataSourceName string, paramsList []map[string]interface{}) ([]int64, error) {
	// 业务逻辑：限制批量操作大小
	if len(paramsList) > 1000 {
		return nil, NewError("batch operation too large, max 1000")
	}

	uc.log.WithContext(ctx).Infof("BatchExecute SQL: dataSource=%s, batchSize=%d", dataSourceName, len(paramsList))

	return uc.repo.BatchExecute(ctx, sql, dataSourceName, paramsList)
}

// TestConnection 测试连接
func (uc *AgilityUsecase) TestConnection(ctx context.Context, dataSourceName string) error {
	uc.log.WithContext(ctx).Infof("Testing connection: dataSource=%s", dataSourceName)

	return uc.repo.TestConnection(ctx, dataSourceName)
}

// isSafeSQL 简单的SQL安全检查
//func isSafeSQL(sql string) bool {
//	// 这里可以实现更复杂的安全检查
//	// 例如：禁止DROP、TRUNCATE等高危操作（根据业务需求调整）
//	dangerousKeywords := []string{"DROP", "TRUNCATE", "SHUTDOWN"}
//
//	upperSQL := sql
//	for _, keyword := range dangerousKeywords {
//		// 简单的关键词检查，实际应用中需要更复杂的解析
//		// 这里只是示例
//		// TODO: 实现更安全的SQL检查
//	}
//
//	return true
//}

// Error 业务错误
type Error struct {
	Message string
}

func NewError(message string) error {
	return &Error{Message: message}
}

func (e *Error) Error() string {
	return e.Message
}
