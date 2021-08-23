package todo

import (
	"context"
	"github.com/acoshift/pgsql"
	"github.com/acoshift/pgsql/pgctx"
)

// type BaseRepository interface {
// 	FindAll() ([]Todo, error)
// 	FindByID(id string) (Todo, error)
// }

type Repository struct {
	// ctx context.Context
}

func (r Repository) FindAll(ctx context.Context) ([]*Todo, error) {
	result := make([]*Todo, 0)

	// language=SQL
	err := pgctx.Iter(ctx, func(scan pgsql.Scanner) error {
		var t Todo
		if err := scan(&t.ID, &t.OwnerID, &t.Content, &t.IsActive, &t.DueDate, &t.CreatedAt); err != nil {
			return err
		}
		result = append(result, &t)
		return nil
	}, `select id, owner_id, content, is_active, due_date, created_at from todos `)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r Repository) FindByID(ctx context.Context, id string) (Todo, error) {
	panic("implement me")
}