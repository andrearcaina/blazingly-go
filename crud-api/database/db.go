package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	DbName string
	DbHost string
	DbUser string
	DbPass string
)

func init() {
	loadEnv()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	DbName = os.Getenv("DB_NAME")
	DbHost = os.Getenv("DB_HOST")
	DbUser = os.Getenv("DB_USER")
	DbPass = os.Getenv("DB_PASS")

	log.Println("Loaded environment variables")
}

func dbURI() string {
	return fmt.Sprintf("%s:%s@tcp(%s:8080)/%s", DbUser, DbPass, DbHost, DbName)
}

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", dbURI())

	if err != nil {
		log.Fatalf("Error connecting to database")
	}

	err = db.Ping()

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	MigrateDB(db)

	log.Println("Connected to database")

	return db
}

func MigrateDB(DB *sql.DB) {
	_, useErr := DB.Exec("USE " + DbName)

	if useErr != nil {
		log.Fatalf("Error using database: %v", useErr)
	}

	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS students (
		id INT AUTO_INCREMENT PRIMARY KEY,
		first_name VARCHAR(100),
		last_name VARCHAR(100),
		age INT,
		major VARCHAR(100)
	)`)
	if err != nil {
		log.Fatalf("Error creating students table: %v", err)
	}

	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS professors (
		id INT AUTO_INCREMENT PRIMARY KEY,
		first_name VARCHAR(100),
		last_name VARCHAR(100),
		age INT,
		department VARCHAR(100)
	)`)
	if err != nil {
		log.Fatalf("Error creating professors table: %v", err)
	}

	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS courses (
		id INT AUTO_INCREMENT PRIMARY KEY,
		students_id INT,
		professors_id INT,
		course_name VARCHAR(100),
		course_code VARCHAR(100)
	)`)
	if err != nil {
		log.Fatalf("Error creating courses table: %v", err)
	}
}
