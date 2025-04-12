package group

import (
	"errors"

	"github.com/sanzashi987/nino-work/apps/canvix/consts"
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
)

var errTagNotSupported = errors.New("not find a corresponding interface related to the give type tag")

var typeTagToGroupCountHandler = map[string]any{
	consts.PROJECT: model.ProjectModel{},
	consts.BLOCK:   model.BlockModel{},
	consts.DESIGN: model.AssetModel{
		BaseModel: model.BaseModel{
			TypeTag: consts.DESIGN,
		},
	},
	consts.FONT: model.AssetModel{
		BaseModel: model.BaseModel{
			TypeTag: consts.FONT,
		},
	},
	consts.COMPONENT: model.AssetModel{
		BaseModel: model.BaseModel{
			TypeTag: consts.COMPONENT,
		},
	},
	consts.DATASOURCE: model.DataSourceModel{},
}
