package biz

import (
	"context"
	"fmt"
	"lehu-data-center/app/collect/service/internal/enums"
	"lehu-data-center/app/collect/service/internal/pkg/funcs"
	"lehu-data-center/app/collect/service/internal/types/transfers"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
)

// ChainHandler 责任链处理器接口
type ChainHandler interface {
	Handler(ctx context.Context, totalParam *transfers.TotalParamTransfers) error
	Order() int
}

// UpgradeChainHandler 维度升级责任链处理器
type UpgradeChainHandler struct {
	BaseChainHandler
	upgradeHandler UpgradeHandler
	messageService MessageService
	log            *log.Helper
}

func NewUpgradeChainHandler(
	upgradeHandler UpgradeHandler,
	messageService MessageService,
	logger log.Logger,
) *UpgradeChainHandler {
	return &UpgradeChainHandler{
		upgradeHandler: upgradeHandler,
		messageService: messageService,
		log:            log.NewHelper(logger),
	}
}

func (h *UpgradeChainHandler) Handler(ctx context.Context, totalParam *transfers.TotalParamTransfers) error {
	// 执行维度升级
	newParam, err := h.doUpgradeDimension(ctx, totalParam)
	if err != nil {
		return fmt.Errorf("upgrade dimension failed: %w", err)
	}

	// 如果新参数为空，说明维度已经升级到最高
	if newParam == nil {
		h.log.WithContext(ctx).Infof("All dimensions collected, ruleId: %d, gatherDate: %s",
			totalParam.ParamTransfers.RuleId,
			totalParam.ParamTransfers.RequestTime.GatherDate)
		return nil
	}

	totalParam.ParamTransfers = newParam

	// 创建消息记录
	messageRecord, err := h.messageService.InsertMessageRecord(ctx, totalParam)
	if err != nil {
		return fmt.Errorf("insert message record failed: %w", err)
	}

	// 执行维度升级
	upgradeDomain, err := h.upgradeHandler.UpgradeDimension(
		ctx,
		messageRecord.ID,
		strconv.FormatInt(messageRecord.MessageID, 10),
		totalParam,
	)
	if err != nil {
		return fmt.Errorf("upgrade dimension domain failed: %w", err)
	}

	if upgradeDomain.Error != nil {
		return upgradeDomain.Error
	}

	return nil
}

func (h *UpgradeChainHandler) doUpgradeDimension(
	ctx context.Context,
	totalParam *transfers.TotalParamTransfers,
) (*transfers.ParamTransfers, error) {

	param := totalParam.ParamTransfers
	if param == nil {
		return nil, ErrInvalidParam
	}

	// 如果视频维度是父级分类且时间维度是年，则不需要升级
	if param.VideoDimensionType == enums.VideoDimensionTypeParentVideoType &&
		param.RequestTime.DateType == enums.DateTypeYear {
		return nil, nil
	}

	// 升级维度参数
	newParam, err := h.upgradeDimensionParam(ctx, param)
	if err != nil {
		return nil, fmt.Errorf("upgrade dimension param failed: %w", err)
	}

	if newParam == nil {
		return nil, nil
	}

	return newParam, nil
}

func (h *UpgradeChainHandler) upgradeDimensionParam(
	ctx context.Context,
	param *transfers.ParamTransfers,
) (*transfers.ParamTransfers, error) {

	dateType := param.RequestTime.DateType
	videoDimensionType := param.VideoDimensionType

	switch dateType {
	case enums.DateTypeDay, enums.DateTypeWeek, enums.DateTypeMonth:
		// 升级时间维度
		upgradeDateType := funcs.UpgradeDateType(dateType)
		if upgradeDateType == 0 {
			return nil, fmt.Errorf("upgrade DateType failed")
		}

		// 创建升级后的时间
		upgradeRequestTime, err := funcs.CreateRequestTime(
			upgradeDateType,
			param.RequestTime.GatherDate,
		)
		if err != nil {
			return nil, fmt.Errorf("create request time failed: %w", err)
		}

		newParam := *param
		newParam.RequestTime = upgradeRequestTime
		return &newParam, nil

	case enums.DateTypeYear:
		// 升级视频维度
		upgradeVideoDimensionType := funcs.UpGradeVideoDimension(videoDimensionType)
		if upgradeVideoDimensionType == 0 {
			return nil, fmt.Errorf("upgrade VideoDimensionType failed")
		}

		newParam := *param
		newParam.VideoDimensionType = enums.VideoDimensionType(upgradeVideoDimensionType)

		// 如果升级到父级视频分类，时间维度重置为日维度
		if upgradeVideoDimensionType == enums.VideoDimensionTypeParentVideoType {
			dayRequestTime, err := funcs.CreateRequestTime(
				enums.DateTypeDay,
				param.RequestTime.GatherDate,
			)
			if err != nil {
				return nil, fmt.Errorf("create day request time failed: %w", err)
			}
			newParam.RequestTime = dayRequestTime
		}

		return &newParam, nil

	default:
		return nil, fmt.Errorf("unsupported date type: %d", dateType)
	}
}

func (h *UpgradeChainHandler) Order() int {
	return 4
}
