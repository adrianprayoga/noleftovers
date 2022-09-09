package models

import (
	"database/sql"
	"errors"
	"fmt"
	logger "github.com/adrianprayoga/noleftovers/server/internals/logger"
	"go.uber.org/zap"
)

type FavoritesService struct {
	DB *sql.DB
}

func (service *FavoritesService) GetFavoritesByUserId(userId uint) ([]uint, error) {

	sqlQuery := `SELECT  f.recipe_id FROM favorites f WHERE f.user_id=$1`

	rows, err := service.DB.Query(sqlQuery, userId)
	if err != nil {
		logger.Log.Error("Error getting favorites by userid")
		fmt.Println(err)
		return make([]uint, 0), fmt.Errorf("error getting favorites by userid: %w", err)
	}
	defer rows.Close()

	favorites := make([]uint, 0)
	for rows.Next() {
		var id uint
		err = rows.Scan(&id)
		if err != nil {
			fmt.Println("Error getting favorite recipe")
			return make([]uint, 0), fmt.Errorf("error reading favorite recipe: %w", err)
		}

		favorites = append(favorites, id)
	}

	return favorites, nil
}

func (service *FavoritesService) AddFavoritesForUserId(userId uint, recipeId uint) error {

	exist, err := service.isFavoriteExist(userId, recipeId)
	if err != nil {
		logger.Log.Error("Error checking existing favorites by userid")
		return fmt.Errorf("error checking existing favorites by userid: %w", err)
	}

	if exist {
		return errors.New("favorite already exist")
	}

	sqlQuery := `INSERT INTO favorites (user_id, recipe_id) VALUES ($1, $2)`

	_, err = service.DB.Exec(sqlQuery, userId, recipeId)
	if err != nil {
		logger.Log.Error("Error adding favorites by userid")
		return fmt.Errorf("error adding favorites by userid: %w", err)
	}

	return nil
}

func (service *FavoritesService) RemoveFavoritesForUserId(userId uint, recipeId uint) error {

	exist, err := service.isFavoriteExist(userId, recipeId)
	if err != nil {
		logger.Log.Error("Error checking existing favorites by userid")
		return fmt.Errorf("error checking existing favorites by userid: %w", err)
	}

	if !exist {
		return errors.New("favorites does not exist")
	}

	sqlQuery := `DELETE FROM favorites WHERE user_id=$1 AND recipe_id=$2`

	_, err = service.DB.Exec(sqlQuery, userId, recipeId)
	if err != nil {
		logger.Log.Error("Error removing favorites by userid")
		return fmt.Errorf("error removing favorites by userid: %w", err)
	}

	return nil
}

func (service *FavoritesService) isFavoriteExist(userId uint, recipeId uint) (bool, error) {

	var temp uint
	err := service.DB.QueryRow(`SELECT recipe_id FROM favorites WHERE recipe_id=$1 AND user_id=$2`,
		recipeId, userId).Scan(&temp)

	if err != nil && err != sql.ErrNoRows {
		logger.Log.Error("Error checking whether favorite exists")
		logger.Log.Error("", zap.Error(err))
	} else if err != nil && err == sql.ErrNoRows {
		return false, nil
	}

	return true, nil
}
