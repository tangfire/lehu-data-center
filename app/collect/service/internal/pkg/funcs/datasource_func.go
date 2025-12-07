package funcs

import (
	"fmt"
	"lehu-data-center/app/collect/service/internal/enums"
	"strings"
)

const (
	Stats      = "stats"
	StatsRange = "stats_range"
	UnderLine  = "_"
)

// CreateTableName 创建统计表名
func CreateTableName(tablePrefixName, ruleName string, videoDimensionType enums.VideoDimensionType, dateType enums.DateType) string {
	datePart := Stats
	if dateType != enums.DateTypeDay {
		datePart = StatsRange
	}

	videoDimValue := strings.ToLower(videoDimensionType.Value())
	return fmt.Sprintf("%s%s%s%s%s%s",
		tablePrefixName, ruleName, UnderLine, videoDimValue, UnderLine, datePart)
}

// CreateDeleteTableDataSql 创建删除表数据的SQL
func CreateDeleteTableDataSql(tableName string, videoDimensionType enums.VideoDimensionType) string {
	sql := fmt.Sprintf("DELETE FROM %s WHERE stats_time = :stats_time AND parent_video_type_id = :parent_video_type_id", tableName)

	// 注意：这里使用枚举值进行比较，不是使用字符串
	if videoDimensionType == enums.VideoDimensionTypeVideoType {
		sql += " AND video_type_id = :video_type_id"
	}

	if videoDimensionType == enums.VideoDimensionTypeVideo {
		sql += " AND video_type_id = :video_type_id AND video_id IN (:video_ids)"
	}

	return sql
}

// CreateInsertTableDataSql 创建插入表数据的SQL
func CreateInsertTableDataSql(tableName string, resultDataList []map[string]interface{}) string {
	if len(resultDataList) == 0 {
		return ""
	}

	// 获取列名（从第一个元素中获取）
	var columns []string
	for key := range resultDataList[0] {
		columns = append(columns, key)
	}

	// 创建占位符
	var placeholders []string
	for _, col := range columns {
		placeholders = append(placeholders, ":"+col)
	}

	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		tableName,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "))
}
