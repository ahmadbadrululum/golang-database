package repository

import (
	"context"
	"golang-database/entity"
)

// contract interface seperti model, mengumpulkan fungsi2 yg masuk pada database, seperti insert, update, delete, dll
type CommentRepository interface {
	// (parsing context, entyty) dan (return value nya)
	Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	FindById(ctx context.Context, id int32) (entity.Comment, error)
	FindAll(ctx context.Context) ([]entity.Comment, error)
}
