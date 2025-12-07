package transfers

// DimensionTransfer 维度传输对象
type DimensionTransfer struct {
	// 视频ID列表
	VideoIdList []int64 `json:"video_id_list"`
	// 视频类型ID
	VideoTypeId int64 `json:"video_type_id"`
	// 父级视频类型ID
	ParentVideoTypeId int64 `json:"parent_video_type_id"`
}
