package CommentsRest

import (
	"errors"
	"time"

	CommentsPackage "github.com/Commenter/Comments"
)

type CommentsController struct {
	CommentsMap map[int64]*CommentsPackage.CommentsStruct
}

var Controller CommentsController

func init() {
	Controller = CommentsController{CommentsPackage.CommentsMap}
}

func (controller CommentsController) AddComment(commentId int64, content, user string) error {
	if _, ok := controller.CommentsMap[commentId]; ok {
		return errors.New("commentId Already Present")
	}
	newComment := CommentsPackage.CommentsStruct{content, time.Now(), user}
	controller.CommentsMap[commentId] = &newComment
	return nil
}

func (controller CommentsController) ModifyComment(commentId int64, content string) error {
	if _, ok := controller.CommentsMap[commentId]; ok {
		controller.CommentsMap[commentId].SetComment(content)
		return nil

	}
	return errors.New("no comment with that commentId")
}

func (controller CommentsController) GetComment(commentId int64) (error, CommentsPackage.CommentsStruct) {
	if val, ok := controller.CommentsMap[commentId]; ok {
		content := val.GetComment()
		return nil, CommentsPackage.CommentsStruct{content, val.LastUpdated, val.ByPerson}

	}
	return errors.New("no comment with that commentId"), CommentsPackage.CommentsStruct{"", time.Now(), ""}
}

func (controller CommentsController) GetAllComments() []CommentsPackage.CommentsStruct {
	ReturnSlice := make([]CommentsPackage.CommentsStruct, 0, 1)
	for _, v := range CommentsPackage.CommentsMap {
		ReturnSlice = append(ReturnSlice, CommentsPackage.CommentsStruct{v.GetComment(), v.GetLastUpdateTime(), v.GetPersonName()})
	}
	return ReturnSlice
}

func (controller CommentsController) DeleteComments(commentId int64) error {
	if _, ok := controller.CommentsMap[commentId]; ok {
		delete(controller.CommentsMap, commentId)
		return nil
	}
	return errors.New("No comment with that comment id")
}
