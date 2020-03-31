package main

import (
	"database/sql"
	"fmt"
	"regexp"

	_ "github.com/lib/pq"
)

func normalize(phone string) string {
	re := regexp.MustCompile("\\D")
	return re.ReplaceAllString(phone, "")
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Password@123"
	dbName   = "gophercises_phone"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	pssqlInfo := fmt.Sprintf("host = %s port = %d user = %s password = %s sslmode = disable ", host, port, user, password)
	// db, err := sql.Open("postgres", pssqlInfo)
	// must(err)
	// must(resetDb(db, dbName))
	// defer db.Close()

	pssqlInfo = fmt.Sprintf("%s dbname = %s", pssqlInfo, dbName)
	db, err := sql.Open("postgres", pssqlInfo)
	must(err)
	defer db.Close()
	must(createPhoneNumberTable(db))

	// id, err := insertPhoneNumber(db, "1234567890")
	// must(err)
	// fmt.Println("id = ", id)
	id := 3
	phone, err := getPhone(db, id)
	must(err)
	fmt.Printf("id = %d phone %s\n", id, phone)
	phones, err := allPhones(db)
	must(err)
	for _, p := range phones {
		fmt.Printf("%d = %s\n", p.id, p.number)
		number := normalize(p.number)
		if number != p.number {
			existing, err := findPhone(db, number)
			must(err)
			if existing != nil {
				fmt.Println("Found and Deleting...", number)
				must(deletePhone(db, p))
			} else {
				fmt.Println("Not found and updating...", number)
				p.number = number
				must(updatePhone(db, p))
			}

		}

	}
}

func findPhone(db *sql.DB, number string) (*phone, error) {
	var p phone
	statement := "SELECT * FROM PHONE_NUMBER WHERE VALUE = $1"
	err := db.QueryRow(statement, number).Scan(&p.id, &p.number)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

func getPhone(db *sql.DB, id int) (string, error) {
	var phone string
	statement := "SELECT VALUE FROM PHONE_NUMBER WHERE ID = $1"
	err := db.QueryRow(statement, id).Scan(&phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}
	return phone, nil
}

type phone struct {
	id     int
	number string
}

func allPhones(db *sql.DB) ([]phone, error) {
	var phones []phone
	rows, err := db.Query("SELECT * FROM PHONE_NUMBER")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var p phone
		if err := rows.Scan(&p.id, &p.number); err != nil {
			return nil, err
		}
		phones = append(phones, p)
	}
	return phones, nil
}
func updatePhone(db *sql.DB, p phone) error {
	statement := "UPDATE PHONE_NUMBER SET VALUE = $2 WHERE ID = $1"
	_, err := db.Exec(statement, p.id, p.number)
	return err
}

func deletePhone(db *sql.DB, p phone) error {
	statement := "DELETE FROM PHONE_NUMBER WHERE ID = $1"
	_, err := db.Exec(statement, p.id)
	return err
}

func insertPhoneNumber(db *sql.DB, phone string) (int, error) {
	statement := `
	INSERT INTO PHONE_NUMBER(value) VALUES($1) RETURNING ID
	`
	var id int
	err := db.QueryRow(statement, phone).Scan(&id)
	return id, err
}
func createPhoneNumberTable(db *sql.DB) error {
	statement := `
	CREATE TABLE IF NOT EXISTS PHONE_NUMBER (
	ID SERIAL,
	VALUE VARCHAR(255)
	)
	`
	_, err := db.Exec(statement)
	return err
}

func resetDb(db *sql.DB, name string) error {

	_, err := db.Exec("DROP DATABASE IF EXISTS " + name)
	if err != nil {
		panic(err)
	}
	return createDb(db, name)
}

func createDb(db *sql.DB, name string) error {
	_, err := db.Exec("CREATE DATABASE " + name)
	return err
}
