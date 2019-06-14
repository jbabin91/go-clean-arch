package repository

import (
	"context"
	"database/sql"
	"fmt"

	author "github.com/jbabin91/go-clean-arch/models"

	"github.com/sirupsen/logrus"

	article "github.com/jbabin91/go-clean-arch/article"
	models "github.com/jbabin91/go-clean-arch/models"
)

type MsqlArticleRepository struct {
	Conn *sql.DB
}

func NewMysqlArticleRepository(Conn *sql.DB) article.ArticleRepository {
	return &MsqlArticleRepository{Conn}
}

func (m *MsqlArticleRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Article, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer rows.Close()

	result := make([]*models.Article, 0)
	for rows.Next() {
		t := new(model.Article)
		authorID := uint(0)
		err = rows.Scan(
			&t.ID,
			&t.Title,
			&t.Content,
			&authorID,
			&t.UpdatedAt,
			&t.CreatedAt,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		t.Author = author.Author {
			ID: authorID,
		}

		result = append(result, t)
	}

	return result, nil
}

func (m *MsqlArticleRepository) Fetch(ctx context.Context, cursor string, num int64) ([]*models.Article, error) {
	query := `SELECT id, title, content, author_id, updated_at, created_at FROM article WHERE id > ? LIMIT ?`

	return m.fetch(ctx, query, cursor, num)
}

func (m *MysqlArticleRepository) GetByID(ctx context.Context, id int64) (*models.Article, error) {
	query := `SELECT id, title, content, author_id, updated_at, created_at FROM article WHERE id = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	a := &models.Article{}
	if len(list) > 0 {
		a = list[0]
	} else {
		return nil, models.NOT_FOUND_ERROR
	}

	return a, nil
}

func (m *MsqlArticleRepository) GetByTitle(ctx context.Context, title string) (*models.Article, error) {
	query := `SELECT id, title, content, author_id, updated_at, created_at FROM article WHERE title = ?`

	list, err := m.fetch(ctx, query, title)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	a := &models.Article{}
	if len(list) > 0 {
		a = list[0]
	} else {
		return nil, models.NOT_FOUND_ERROR
	}

	return a, nil
}

func (m *MsqlArticleRepository) Store(ctx context.Context, a *models.Article) (int64, error) {
	query := `INSERT article SET title = ?, content = ?, author_id = ?, updated_at = ?, created_at = ?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	logrus.Debug("Created At:", a.CreatedAt)
	
	res, err := stmt.ExecContext(ctx, a.Title, a.Content, a.Author.ID, a.UpdatedAt, a.CreatedAt)
	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	return res.LastInsertId()
}

func (m *MysqlArticleRepository) Delete(ctx context.Context, id int64) (bool, error) {
	query := `DELETE FROM article WHERE id = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		logrus.Error(err)
		return false, err
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		logrus.Error(err)
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logrus.Error(err)
		return false, err
	}
	if rowsAffected != 1 {
		err = fmt.Errorf("Weird Behavior. Total Affected: %d", rowsAffected)
		logrus.Error(err)
		return false, err
	}

	return true, nil
}

func (m *MsqlArticleRepository) Update(ctx context.Context, ar *models.Article) (*models.Article, error) {
	query := `UPDATE article set title = ?, content = ?, author_id = ?, updated_at = ? WHERE ID = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	res, err := stmt.ExecContext(ctx, ar.Title, ar.Content, ar.Author.ID, ar.UpdateAt, ar.ID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	if rowsAffected != 1 {
		err = fmt.ErrorF("Weird Behavior. Total Affected: %d", rowsAffected)
		logrus.Error(err)
		return nil, err
	}

	return ar, nil
}