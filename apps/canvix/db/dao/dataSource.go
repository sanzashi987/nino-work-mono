package dao

import (
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
	"github.com/sanzashi987/nino-work/pkg/shared"
	"gorm.io/gorm"
)

func FindByNameOrType(tx *gorm.DB, workspace uint64, name string, sourceType []string, pagination shared.PaginationRequest) ([]*model.DataSourceModel, error) {
	result := []*model.DataSourceModel{}
	query := tx.Scopes(db.Paginate(pagination.Page, pagination.Size)).Model(&model.DataSourceModel{}).Where("workspace = ?", workspace)

	if name != "" {
		query = query.Where("name LIKE ?", name)
	}

	if len(sourceType) != 0 {
		sourceTypeEnums := []uint8{}

		for _, v := range sourceType {
			enum, exist := model.SourceTypeStringToEnum[v]
			if exist {
				sourceTypeEnums = append(sourceTypeEnums, enum)
			}
		}

		query = query.Where("source_type IN ?", sourceTypeEnums)
	}

	err := query.Find(&result).Error
	return result, err

}

func UpdateProject(tx *gorm.DB, workspace, id uint64, sourceName, sourceInfo *string) (*model.DataSourceModel, error) {
	toUpdate := map[string]string{}
	var result = model.DataSourceModel{}
	if sourceName != nil {
		toUpdate["name"] = *sourceName
	}

	if sourceInfo != nil {
		toUpdate["source_info"] = *sourceInfo
	}
	err := tx.Model(&result).Where("id = ? and workspace = ?", id, workspace).Updates(toUpdate).Error

	return &result, err
}
