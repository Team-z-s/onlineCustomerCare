package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"html/template"

	"github.com/julienschmidt/httprouter"
	"gitlab.com/username/online-service-and-customer-care2.0/comment"
	"gitlab.com/username/online-service-and-customer-care2.0/entity"
)

// AdminCommentHandler handles comment related http requests
type AdminCommentHandler struct {
	commentService comment.CommentService
	tmpl        *template.Template
}

// NewAdminCommentHandler returns new AdminCommentHandler object
func NewAdminCommentHandler(cmntService comment.CommentService, T *template.Template) *AdminCommentHandler {
	return &AdminCommentHandler{commentService: cmntService, tmpl:T}
}

// GetComments handles GET  request company_dashboard/getcomment
func (ach *AdminCommentHandler) GetComments(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	comments, errs := ach.commentService.Comments()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	ach.tmpl.ExecuteTemplate(w,"showcomment.layout",comments)
	return

}

// PostComment handles POST  request
func (ach *AdminCommentHandler) PostComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	comment := &entity.Comment{}

	err := json.Unmarshal(body, comment)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	comment, errs := ach.commentService.StoreComment(comment)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	p := fmt.Sprintf("/v1/admin/comments/%d", comment.ID)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}

// DeleteComment handles DELETE request
func (ach *AdminCommentHandler) DeleteComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := ach.commentService.DeleteComment(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
