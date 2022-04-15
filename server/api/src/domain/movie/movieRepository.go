package movie

import "web-service/api/database"

type MovieRepository struct {
	database.IDBHandler
}

func (entity *MovieRepository) SaveMovie() {

}
