package author

import (
	"context"

	"github.com/jbabin91/go-clean-arch/models"
)

// Repository defines the methods used with the db to retrieve data for the author.Repository
type Repository interface {
	GetByID(ctx context.Context, id uint64) (*models.Author, error) 
}