package repository

import (
	"context"
	"hyphen-backend-hellog/ent"
	"hyphen-backend-hellog/ent/comment"
	"time"
)

type CommentRepository interface {
	Create(ctx context.Context, content string, post *ent.Post, author *ent.User, parent *ent.Comment) (*ent.Comment, error)
	QueryByID(ctx context.Context, id int) (*ent.Comment, error)
	QueryChilds(ctx context.Context, parent *ent.Comment) ([]*ent.Comment, error)
	QueryParentsByPost(ctx context.Context, post *ent.Post) ([]*ent.Comment, error)
	UpdateByID(ctx context.Context, id int, content string) (*ent.Comment, error)
	DeleteByID(ctx context.Context, id int) error
}

func NewCommentRepository(database *ent.Client) CommentRepository {
	return &ICommentRepository{client: database}
}

type ICommentRepository struct {
	client *ent.Client
}

func (r *ICommentRepository) Create(ctx context.Context, content string, post *ent.Post, author *ent.User, parent *ent.Comment) (*ent.Comment, error) {
	comment := r.client.Comment.Create().
		SetContent(content).
		SetPost(post).
		SetAuthor(author)

	if parent != nil {
		comment.SetParent(parent)
	}

	return comment.
		Save(ctx)
}

func (r *ICommentRepository) QueryByID(ctx context.Context, id int) (*ent.Comment, error) {
	return r.client.Comment.Get(ctx, id)
}

func (r *ICommentRepository) QueryChilds(ctx context.Context, parent *ent.Comment) ([]*ent.Comment, error) {
	return parent.QueryChildrens().All(ctx)

}

func (r *ICommentRepository) QueryParentsByPost(ctx context.Context, post *ent.Post) ([]*ent.Comment, error) {
	return post.QueryComments().Where(comment.Not(comment.HasParent())).All(ctx)
}

func (r *ICommentRepository) UpdateByID(ctx context.Context, id int, content string) (*ent.Comment, error) {
	return r.client.Comment.UpdateOneID(id).
		SetContent(content).
		SetUpdatedAt(time.Now()).
		Save(ctx)

}

func (r *ICommentRepository) DeleteByID(ctx context.Context, id int) error {
	return r.client.Comment.DeleteOneID(id).
		Exec(ctx)
}
