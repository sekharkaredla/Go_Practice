package Comments

import (
	"time"
)

type CommentsStruct struct {
	Content     string
	LastUpdated time.Time
	ByPerson    string
}

func (cmt CommentsStruct) GetComment() string {
	return cmt.Content
}

func (cmt *CommentsStruct) SetComment(newCmt string) {
	cmt.Content = newCmt
}

func (cmt CommentsStruct) GetLastUpdateTime() time.Time {
	return cmt.LastUpdated
}

func (cmt CommentsStruct) GetPersonName() string {
	return cmt.ByPerson
}
