package comment

import "gitlab.com/username/online-service-and-customer-care2.0/entity"

// CommentService specifies customer comment related service
type CommentService interface {
	Comments() ([]entity.Comment, []error)
	Comment(id uint) (*entity.Comment, []error)
	UpdateComment(comment *entity.Comment) (*entity.Comment, []error)
	DeleteComment(id uint) (*entity.Comment, []error)
	StoreComment(comment *entity.Comment) (*entity.Comment, []error)
}
