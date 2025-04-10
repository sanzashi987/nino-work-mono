package group

import (
	"errors"

	"github.com/sanzashi987/nino-work/apps/canvix/consts"
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
)

var errTagNotSupported = errors.New("Not find a corresponding interface related to the give type tag")

var typeTagToGroupCountHandler = map[string]any{
	consts.PROJECT:    model.ProjectModel{},
	consts.BLOCK:      model.BaseModel{},
	consts.DESIGN:     model.AssetModel{},
	consts.FONT:       model.AssetModel{},
	consts.COMPONENT:  model.AssetModel{},
	consts.DATASOURCE: model.DataSourceModel{},
}
