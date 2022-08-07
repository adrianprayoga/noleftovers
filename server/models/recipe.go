package models

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/copier"
)

type Recipe struct {
	Id          uint          `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Author      sql.NullInt32 `json:"author"`
	ModifiedAt  string        `json:"modified_at"`
	CreatedAt   string        `json:"created_at"`
}

type RecipeService struct {
	DB *sql.DB
}

type RecipeValidator struct{}

func (service *RecipeService) CreateRecipe(recipe Recipe) (*Recipe, error) {
	newRecipe := Recipe{}
	copier.Copy(&newRecipe, &recipe)

	row := service.DB.QueryRow(`INSERT INTO recipes (name, description) 
								 VALUES ($1, $2) RETURNING id`, recipe.Name, recipe.Description)
	err := row.Scan(&newRecipe.Id)
	if err != nil {
		return nil, fmt.Errorf("error creating recipe: %w", err)
	}
	return &newRecipe, nil
}

func (service *RecipeService) GetRecipes() ([]Recipe, error) {

	rows, err := service.DB.Query(`SELECT name, description, author FROM recipes`)
	if err != nil {
		fmt.Println("Error getting recipe list")
		return nil, fmt.Errorf("error getting recipe list: %w", err)
	}
	defer rows.Close()

	var recipes []Recipe
	for rows.Next() {
		recipe := Recipe{}
		err = rows.Scan(&recipe.Name, &recipe.Description, &recipe.Author)
		if err != nil {
			fmt.Println("Error reading recipe")
			return nil, fmt.Errorf("error reading recipe: %w", err)
		}

		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

func (service *RecipeService) GetRecipeById(id uint) (*Recipe, error) {
	recipe := Recipe{
		Id: id,
	}

	row := service.DB.QueryRow(`SELECT name, description FROM recipes WHERE id=$1`, id)
	err := row.Scan(&recipe.Name, &recipe.Description)
	if err != nil {
		fmt.Println("Error authenticating")
		return nil, fmt.Errorf("user login: %w", err)
	}

	fmt.Printf("Recipe %+v\n", recipe)

	return &recipe, nil
}
