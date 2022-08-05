package models

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/copier"
)

type Recipe struct {
	Id 	uint
	Name string
	Description string
	Author uint
}

type RecipeService struct {
	DB *sql.DB
}

type RecipeValidator struct {
}


func (us *RecipeService) CreateRecipe(recipe Recipe) (*Recipe, error) {
	newRecipe := Recipe{}
	copier.Copy(&newRecipe, &recipe)

	row := us.DB.QueryRow(`INSERT INTO recipes (name, description) 
								 VALUES ($1, $2) RETURNING id`, recipe.Name, recipe.Description)
	err := row.Scan(&newRecipe.Id)
	if err != nil {
		return nil, fmt.Errorf("error creating recipe: %w", err)
	}
	return &newRecipe, nil
}

func (us *RecipeService) SearchRecipeById(id uint) (*Recipe, error) {
	recipe := Recipe{
		Id: id,
	}

	row := us.DB.QueryRow(`SELECT name, description FROM recipes WHERE id=$1`, id)
	err := row.Scan(&recipe.Name, &recipe.Description)
	if err != nil {
		fmt.Println("Error authenticating")
		return nil, fmt.Errorf("user login: %w", err)
	}

	fmt.Printf("Recipe %+v\n", recipe)

	return &recipe, nil
}