package Comments

import (
	"time"
)

type Comment interface {
	getComment() string
	setComment(comment string)
	getLastUpdatedTime() time.Time
	getPersonName() string
}
