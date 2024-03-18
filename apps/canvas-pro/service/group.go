package service

import (
	"context"
	"errors"
	"regexp"
)

type GroupService struct{}

var groupService *GroupService

func init() {
	groupService = &GroupService{}
}

func GetGroupService() *GroupService {
	return groupService
}

const legalNameRegex = `^[\u4E00-\u9FA5\uF900-\uFA2D\w][\u4E00-\u9FA5\uF900-\uFA2D\w-_]*[\u4E00-\u9FA5\uF900-\uFA2D\w]*$`

var reg, _ = regexp.Compile(legalNameRegex)

var ErrorNameContainIllegalChar = errors.New("error name contains illegal character")

func (serv *GroupService) Create(ctx context.Context, name, workspace string) (err error) {

	if reg.FindStringIndex(name) == nil {
		err = ErrorNameContainIllegalChar
		return
	}

	return
}
