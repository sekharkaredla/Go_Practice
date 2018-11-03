package main

import (
	"fmt"
	//	"time"

	//	CommentPackage "github.com/Commenter/Comments"
	CommentMethods "github.com/Commenter/CommentsRest"
)

//func main() {
//	CommentPackage.CommentsMap[1] = &CommentPackage.CommentsStruct{Content: "comment1", LastUpdated: time.Now(), ByPerson: "sekhar"}
//	fmt.Println("Map:", CommentPackage.CommentsMap)
//	fmt.Println("old content : ", CommentPackage.CommentsMap[1].GetComment())
//	CommentPackage.CommentsMap[1].SetComment("comment1_new")
//	fmt.Println("new content : ", CommentPackage.CommentsMap[1].GetComment())
//	fmt.Println(CommentPackage.CommentsMap[1].GetPersonName())
//	fmt.Println(CommentPackage.CommentsMap[1].GetLastUpdateTime())
//}

func main() {
	CommentMethods.Controller.AddComment(1, "comment1", "user")
	fmt.Println(CommentMethods.Controller.GetComment(1))
	fmt.Println(CommentMethods.Controller.CommentsMap)
	fmt.Println(CommentMethods.Controller.ModifyComment(2, "wanted Error"))
	fmt.Println(CommentMethods.Controller.ModifyComment(1, "comment1_new"))
	fmt.Println(CommentMethods.Controller.GetComment(1))
}
