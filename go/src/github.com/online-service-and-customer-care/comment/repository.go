package comment

import (
	"onlineCustomerCare/entity"
)

// CommentRepository specifies customer comment related database operations
type CommentRepository interface {
	Comments() ([]entity.Comment, []error)
	StoreComment(comment *entity.Comment) (*entity.Comment, []error)
}
