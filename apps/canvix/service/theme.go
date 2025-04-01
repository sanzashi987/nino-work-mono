package service

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/canvix/consts"
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
	"github.com/sanzashi987/nino-work/pkg/shared"
)

type UpdateThemeReq struct {
	Id     uint64  `json:"id" binding:"required"`
	Name   *string `json:"name"`
	Config *string `json:"config"`
}

func UpdateTheme(ctx context.Context, workspaceId uint64, req *UpdateThemeReq) error {
	tx := db.NewTx(ctx)
	toUpdate := map[string]string{}

	if req.Name != nil {
		toUpdate["name"] = *req.Name
	}

	if req.Config != nil {
		toUpdate["config"] = *req.Config
	}

	return tx.Model(&model.ThemeModel{}).Where("id = ? and workspace = ?", req.Id, workspaceId).Updates(toUpdate).Error
}

type CreateThemeReq struct {
	Name   string `json:"name" binding:"required"`
	Config string `json:"theme" binding:"required"`
	/**
	 * 0: default theme
	 * 1: custom theme
	 */
	Flag int `json:"flag" binding:"required"`
}

func CreateTheme(ctx context.Context, workspaceId uint64, req *CreateThemeReq) error {
	tx := db.NewTx(ctx)

	toCreate := model.ThemeModel{
		Type:   int8(req.Flag),
		Config: req.Config,
	}

	toCreate.Workspace = workspaceId
	toCreate.Name, toCreate.TypeTag = req.Name, consts.THEME

	tx = tx.Begin()

	if err := tx.Model(&model.ThemeModel{}).Create(&toCreate).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

type ThemeItem struct {
	Name  string `json:"name"`
	Theme string `json:"theme"`
	Id    uint64 `json:"id"`
	shared.DBTime
}

func GetThemes(ctx context.Context, workspaceId uint64) ([]*ThemeItem, error) {
	tx := db.NewTx(ctx)
	models := []model.ThemeModel{}
	if err := tx.Where("workspace = ? ", workspaceId).Find(&models).Error; err != nil {
		return nil, err
	}

	res := []*ThemeItem{}

	for _, model := range models {
		res = append(res, &ThemeItem{
			DBTime: shared.DBTime{
				CreateTime: model.GetCreatedDate(),
				UpdateTime: model.GetUpdatedDate(),
			},
			Name:  model.Name,
			Id:    model.Id,
			Theme: model.Config,
		})
	}

	return res, nil
}

type DeleteThemeReq struct {
	Data []uint64 `json:"data" binding:"required"`
}

func DeleteThemes(ctx context.Context, workspaceId uint64, ids []uint64) error {
	tx := db.NewTx(ctx)
	tx = tx.Begin()
	if err := tx.Model(&model.ThemeModel{}).Where("id in ? AND workspace = ?", ids, workspaceId).Delete(&model.ThemeModel{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
