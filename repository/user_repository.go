package repository

import (
	"context"
	"hyphen-backend-hellog/ent"
)

type UserRepository interface {
	Create(ctx context.Context, idFromUserMS int, name string, profileImage string) (*ent.User, error)
	QueryByID(ctx context.Context, id int) (*ent.User, error)
	UpdateByID(ctx context.Context, id int, idFromUserMS int, name string, profileImage string) (*ent.User, error)
	DeleteByID(ctx context.Context, id int) error
}

func NewUserRepository(database *ent.Client) UserRepository {
	return &IUserRepository{client: database}
}

type IUserRepository struct {
	client *ent.Client
}

func (r *IUserRepository) Create(ctx context.Context, idFromUserMS int, name string, profileImage string) (*ent.User, error) {
	return r.client.User.Create().
		SetIDFromUserMs(idFromUserMS).
		SetName(name).
		SetProfileImage(profileImage).
		Save(ctx)

}

func (r *IUserRepository) QueryByID(ctx context.Context, id int) (*ent.User, error) {
	return r.client.User.Get(ctx, id)
}

func (r *IUserRepository) UpdateByID(ctx context.Context, id int, idFromUserMS int, name string, profileImage string) (*ent.User, error) {
	return r.client.User.UpdateOneID(id).
		SetIDFromUserMs(idFromUserMS).
		SetName(name).
		SetProfileImage(profileImage).
		Save(ctx)
}

func (r *IUserRepository) DeleteByID(ctx context.Context, id int) error {
	return r.client.User.DeleteOneID(id).Exec(ctx)
}