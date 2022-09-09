package models

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/copier"
	"strconv"
	"strings"
)

type Recipe struct {
	Id          uint         `json:"id,omitempty"`
	Name        string       `json:"name,omitempty"`
	Description string       `json:"description,omitempty"`
	Author      uint         `json:"author,omitempty"`
	ImageLink   string       `json:"imageLink,omitempty"`
	ModifiedAt  string       `json:"modifiedAt,omitempty"`
	CreatedAt   string       `json:"createdAt,omitempty"`
	Ingredients []Ingredient `json:"ingredients"`
	Steps       []RecipeStep `json:"steps"`
}

type RecipeStep struct {
	Id         uint   `json:"id,omitempty"`
	RecipeId   uint   `json:"recipeId,omitempty"`
	Position   uint   `json:"position"`
	Text       string `json:"text,omitempty"`
	ModifiedAt string `json:"modifiedAt,omitempty"`
	CreatedAt  string `json:"createdAt,omitempty"`
}

type Ingredient struct {
	Id              uint           `json:"id,omitempty"`
	RecipeId        uint           `json:"recipeId,omitempty"`
	Position        uint           `json:"position"`
	Name            string         `json:"name"`
	Amount          float32        `json:"amount,omitempty"`
	Measure         uint           `json:"measure,omitempty"`
	MeasureResolved sql.NullString `json:"measureValue,omitempty"`
	ModifiedAt      string         `json:"modifiedAt,omitempty"`
	CreatedAt       string         `json:"createdAt,omitempty"`
}

type Favorites struct {
	RecipeId uint `json:"recipe_id"`
	UserId   uint `json:"user_id"`
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

	sqlQuery := `SELECT 
		r.name, r.description, r.author, r.image_link, i.position, i.name, i.amount, i.measure, m.name
	FROM recipes r
			LEFT JOIN ingredients i ON r.id = i.recipe_id
			LEFT JOIN measure m 	ON i.measure = m.id
	WHERE r.id=$1
	ORDER BY i.position`

	rows, err := service.DB.Query(sqlQuery, id)
	if err != nil {
		fmt.Println("Error getting recipe by id")
		return nil, fmt.Errorf("error getting recipe by id: %w", err)
	}
	defer rows.Close()

	ingredients := make([]Ingredient, 0)
	for rows.Next() {
		i := Ingredient{}
		err = rows.Scan(&recipe.Name, &recipe.Description, &recipe.Author, &recipe.ImageLink, &i.Position, &i.Name, &i.Amount, &i.Measure, &i.MeasureResolved)
		if err != nil {
			fmt.Println("Error reading recipe")
			return nil, fmt.Errorf("error reading recipe: %w", err)
		}

		ingredients = append(ingredients, i)
	}
	recipe.Ingredients = ingredients

	sqlQuery = `SELECT s.text, s.position FROM steps s
	WHERE s.recipe_id=$1
	ORDER BY s.position`

	rows, err = service.DB.Query(sqlQuery, id)
	if err != nil {
		fmt.Println("Error getting recipe by id")
		return nil, fmt.Errorf("error getting recipe by id: %w", err)
	}
	defer rows.Close()

	steps := make([]RecipeStep, 0)
	for rows.Next() {
		step := RecipeStep{}
		err = rows.Scan(&step.Text, &step.Position)
		if err != nil {
			fmt.Println("Error reading recipe")
			return nil, fmt.Errorf("error reading recipe: %w", err)
		}

		steps = append(steps, step)
	}
	recipe.Steps = steps

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
