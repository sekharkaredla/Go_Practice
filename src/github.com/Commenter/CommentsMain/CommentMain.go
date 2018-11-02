package main

import (
	"fmt"
	"time"

	CommentPackage "github.com/Commenter/Comments"
)

func main() {
	CommentPackage.CommentsMap[1] = &CommentPackage.CommentsStruct{Content: "comment1", LastUpdated: time.Now(), ByPerson: "sekhar"}
	fmt.Println("Map:", CommentPackage.CommentsMap)
	fmt.Println("old content : ", CommentPackage.CommentsMap[1].GetComment())
	CommentPackage.CommentsMap[1].SetComment("comment1_new")
	fmt.Println("new content : ", CommentPackage.CommentsMap[1].GetComment())
	fmt.Println(CommentPackage.CommentsMap[1].GetPersonName())
	fmt.Println(CommentPackage.CommentsMap[1].GetLastUpdateTime())
}
