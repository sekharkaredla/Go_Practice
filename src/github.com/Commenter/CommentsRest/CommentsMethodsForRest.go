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

func (controller CommentsController) GetComment(commentId int64) (error, string, string, time.Time) {
	if val, ok := controller.CommentsMap[commentId]; ok {
		content := val.GetComment()
		return nil, content, val.ByPerson, val.LastUpdated

	}
	return errors.New("no comment with that commentId"), "error", "error", time.Now()
}
