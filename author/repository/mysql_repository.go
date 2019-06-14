package repository

import (
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"

	"github.com/jbabin91/go-clean-arch/author"
	"github.com/jbabin91/go-clean-arch/models"
)

// MysqlAuthorRepo defines the db property to be used when getting the author data from the db
type MysqlAuthorRepo struct {
	DB *sql.DB
}

func (m *MysqlAuthorRepo) getOne(ctx context.context, query string, args ...interface{}) (*models.Author, error) {
	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	a := &models.Author{}

	err = row.Scan(
		&a.ID,
		&a.Name,
		&a.CreatedAt,
		&a.UpdatedAt,
	)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return a, nil
}

// GetByID returns the Author's information by using the ID of the author to retrieve the data
func (m *MysqlAuthorRepo) GetByID(ctx context.Context, id uint64) (*models.Author, error) {
	query := `SELECT id, name, created_at, updated_at FROM author WHERE id=?`
	return m.getOne(ctx, query, id)
}