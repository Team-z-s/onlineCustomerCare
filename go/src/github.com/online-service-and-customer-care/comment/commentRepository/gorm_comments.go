package commentRepository

import ("fmt"
		"github.com/jinzhu/gorm"
		"onlineCustomerCare/comment"
		"onlineCustomerCare/entity"
)

// CommentGormRepo implements menu.CommentRepository interface
type CommentGormRepo struct {
	conn *gorm.DB
}

// NewCommentGormRepo returns new object of CommentGormRepo
func NewCommentGormRepo(db *gorm.DB) comment.CommentRepository {
	return &CommentGormRepo{conn: db}
}

// Comments returns all customer comments stored in the database
func (cmntRepo *CommentGormRepo) Comments() ([]entity.Comment, []error) {
	cmnts := []entity.Comment{}
	fmt.Println(cmnts)
	errs := cmntRepo.conn.Find(&cmnts).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnts, errs
}

// StoreComment stores a given customer comment in the database
func (cmntRepo *CommentGormRepo) StoreComment(comment *entity.Comment) (*entity.Comment, []error) {
	cmnt := comment
	errs := cmntRepo.conn.Create(cmnt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}
