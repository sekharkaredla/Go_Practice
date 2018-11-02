package Comments

var CommentsMap map[int64]*CommentsStruct

func init() {
	CommentsMap = make(map[int64]*CommentsStruct)
}
