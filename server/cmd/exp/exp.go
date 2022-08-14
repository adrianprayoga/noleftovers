package main

import (
	"database/sql"
	"fmt"
	"github.com/adrianprayoga/noleftovers/server/models"

	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "baloo"
	password = "junglebook"
	dbname   = "lenslocked"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}

func main() {
	cfg := models.DefaultPostgresConfig()
	db, err := sql.Open("pgx", cfg.String())

	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected!")
	//
	//_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
	//  id SERIAL PRIMARY KEY,
	//  name TEXT,
	//  email TEXT NOT NULL
	//);
	//
	//CREATE TABLE IF NOT EXISTS orders (
	//  id SERIAL PRIMARY KEY,
	//  user_id INT NOT NULL,
	//  amount INT,
	//  description TEXT
	//);`)
	//
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("Tables created.")
	//
	//name := "Jon Calhoun"
	//email := "jon@calhoun.io"
	//_, err = db.Exec(`
	//  INSERT INTO users(name, email)
	//  VALUES($1, $2);`, name, email)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("User created.")

	//rs := models.Service{DB: db}
	//recipe, _ := rs.CreateRecipe(models.Recipe{
	//	Name:        "oklahoma burger",
	//	Description: "lorem ipsum ...",
	//})

	//user, err := us.Authenticate("mr@gmail.com", "123")
	//row := us.DB.QueryRow(`SELECT id, password_hash FROM users WHERE email=$1`, "mr@gmail.com")

	//var id uint
	//var pw string
	//row.Scan(&id, &pw)

}
