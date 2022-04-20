package movie

import (
	"web-service/api/database"
)

type MovieRepository struct {
	database.IDBHandler
}
