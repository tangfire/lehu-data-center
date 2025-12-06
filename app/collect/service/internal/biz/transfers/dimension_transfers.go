package transfers

type DimensionTransfer struct {
	VideoIdList       []int64 // 视频id集合
	VideoTypeId       int64   // 视频分类id
	ParentVideoTypeId int64   // 父级视频分类id
}
