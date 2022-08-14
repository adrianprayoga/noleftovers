package models

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/copier"
	"strconv"
	"strings"
)

type Recipe struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Author      uint   `json:"author"`
	ImageLink   string `json:"imageLink"`
	ModifiedAt  string `json:"modifiedAt"`
	CreatedAt   string `json:"createdAt"`
}

type RecipeStep struct {
	Id         uint   `json:"id"`
	RecipeId   uint   `json:"recipeId"`
	Position   uint   `json:"position"`
	Text       string `json:"text"`
	ModifiedAt string `json:"modifiedAt"`
	CreatedAt  string `json:"createdAt"`
}

type Ingredient struct {
	Id         uint    `json:"id"`
	RecipeId   uint    `json:"recipeId"`
	Position   uint    `json:"position"`
	Name       string  `json:"name"`
	Amount     float32 `json:"amount"`
	Measure    uint    `json:"measure"`
	ModifiedAt string  `json:"modifiedAt"`
	CreatedAt  string  `json:"createdAt"`
}

type RecipeService struct {
	DB *sql.DB
}

type RecipeValidator struct{}

func (service *RecipeService) CreateRecipe(recipe Recipe, steps []RecipeStep, ingredients []Ingredient) (*Recipe, error) {
	newRecipe := Recipe{}
	copier.Copy(&newRecipe, &recipe)

	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Entry to recipes table
	stmt, err := tx.Prepare(`INSERT INTO recipes (name, description, author, image_link) 
								 VALUES ($1, $2, $3, $4) RETURNING id`)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(recipe.Name, recipe.Description, recipe.Author, recipe.ImageLink)
	err = row.Scan(&newRecipe.Id)
	if err != nil {
		tx.Rollback() // return an error too, we may want to wrap them
		return nil, fmt.Errorf("error creating recipe: %w", err)
	}

	// Insert steps
	sqlString := `INSERT INTO steps (recipe_id, position, text) VALUES %s`
	numArguments := 3
	valueArgs := make([]interface{}, 0, numArguments*len(steps))
	for i, step := range steps {
		valueArgs = append(
			valueArgs,
			newRecipe.Id,
			i,
			step.Text,
		)
	}
	err = batchInsert(tx, sqlString, numArguments, len(steps), valueArgs...)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error inserting steps: %w", err)
	}

	// Insert ingredients
	sqlString = `INSERT INTO ingredients (recipe_id, position, name, amount, measure) VALUES %s`
	numArguments = 5
	valueArgs = make([]interface{}, 0, numArguments*len(ingredients))
	for i, ingredient := range ingredients {
		valueArgs = append(
			valueArgs,
			newRecipe.Id,
			i,
			ingredient.Name,
			ingredient.Amount,
			ingredient.Measure,
		)
	}
	err = batchInsert(tx, sqlString, numArguments, len(ingredients), valueArgs...)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error inserting steps: %w", err)
	}

	tx.Commit()

	return &newRecipe, nil
}

func batchInsert(tx *sql.Tx, sqlString string, numArguments int, numEntry int, arguments ...any) error {
	stmt, err := tx.Prepare(getBulkInsertSQLSimple(sqlString, numArguments, numEntry))
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(arguments...)
	if err != nil {
		return err
	}

	return nil
}

func (service *RecipeService) GetRecipes() ([]Recipe, error) {

	rows, err := service.DB.Query(`SELECT id, name, description, author, image_link FROM recipes`)
	if err != nil {
		fmt.Println("Error getting recipe list")
		return nil, fmt.Errorf("error getting recipe list: %w", err)
	}
	defer rows.Close()

	var recipes []Recipe
	for rows.Next() {
		recipe := Recipe{}
		err = rows.Scan(&recipe.Id, &recipe.Name, &recipe.Description, &recipe.Author, &recipe.ImageLink)
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

	row := service.DB.QueryRow(`SELECT name, description, author, image_link FROM recipes WHERE id=$1`, id)
	err := row.Scan(&recipe.Name, &recipe.Description, &recipe.Author, &recipe.ImageLink)
	if err != nil {
		fmt.Println("Error authenticating")
		return nil, fmt.Errorf("user login: %w", err)
	}

	fmt.Printf("Recipe %+v\n", recipe)

	return &recipe, nil
}

func getBulkInsertSQL(SQLString string, rowValueSQL string, numRows int) string {
	// Combine the base SQL string and N value strings
	valueStrings := make([]string, 0, numRows)
	for i := 0; i < numRows; i++ {
		valueStrings = append(valueStrings, "("+rowValueSQL+")")
	}
	allValuesString := strings.Join(valueStrings, ",")
	SQLString = fmt.Sprintf(SQLString, allValuesString)

	// Convert all of the "?" to "$1", "$2", "$3", etc.
	// (which is the way that pgx expects query variables to be)
	numArgs := strings.Count(SQLString, "?")
	SQLString = strings.ReplaceAll(SQLString, "?", "$%v")
	numbers := make([]interface{}, 0, numRows)
	for i := 1; i <= numArgs; i++ {
		numbers = append(numbers, strconv.Itoa(i))
	}
	return fmt.Sprintf(SQLString, numbers...)
}

func getBulkInsertSQLSimple(SQLString string, numArgsPerRow int, numRows int) string {
	questionMarks := make([]string, 0, numArgsPerRow)
	for i := 0; i < numArgsPerRow; i++ {
		questionMarks = append(questionMarks, "?")
	}
	rowValueSQL := strings.Join(questionMarks, ", ")
	return getBulkInsertSQL(SQLString, rowValueSQL, numRows)
}
