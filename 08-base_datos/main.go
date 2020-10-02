package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type persona struct {
	nombre   string
	apellido string
	edad     int
	id       int
}

var p persona

/*
Las parámetros van según la BD específica
MySQL               PostgreSQL            Oracle
=====               ==========            ======
WHERE col = ?       WHERE col = $1        WHERE col = :col
VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
*/
func main() {

	db, err := sql.Open("mysql", "root:soygroot@tcp(localhost:3306)/go_test?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping() //Veo si estoy conectado
	check(err)

	personas := p.presentaTodo(db)
	fmt.Println("traje Todas las Personas", personas)

	fmt.Println("Voy a traer una persona id 1")

	p1 := p.traePersonaPorID(db, 1)

	fmt.Println("P1:", p1)

	p.nombre = "George"
	p.apellido = "Sansonin"
	p.edad = 39
	p.id = 3
	p.persiste(db)

	p2 := p.traePersonaPorID(db, 3)
	fmt.Println("PT2:", p)
	fmt.Println("P2:", p2)

	operacionesComplejas(db, p)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (p persona) presentaTodo(db *sql.DB) []persona {

	rows, err := db.Query("SELECT id, edad, nombre, apellido FROM persona;")
	check(err)
	defer rows.Close()

	var rP []persona
	for rows.Next() {
		err = rows.Scan(&p.id, &p.edad, &p.nombre, &p.apellido)
		check(err)
		rP = append(rP, p)
	}
	return rP
}

func (p persona) traePersonaPorID(db *sql.DB, id int) persona {

	err := db.QueryRow("SELECT id, edad, nombre, apellido FROM persona WHERE id = ?;", id).
		Scan(&p.id, &p.edad, &p.nombre, &p.apellido)
	check(err)

	return p
}

func (p persona) persiste(db *sql.DB) persona {

	stmt, err := db.Prepare("UPDATE persona set nombre = ?, apellido = ?, edad = ? WHERE id = ?;")
	check(err)

	res, err := stmt.Exec(p.nombre, p.apellido, p.edad, p.id)
	check(err)

	rowCnt, err := res.RowsAffected()
	check(err)

	fmt.Println("Modificaciones: ", rowCnt)

	return p
}

func operacionesComplejas(db *sql.DB, p persona) {

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	check(err)

	_, execErr := tx.ExecContext(ctx, "UPDATE persona set edad = ? WHERE id = ?;", (p.edad + 10), p.id)
	checkRollback(execErr, tx)

	_, execErr = tx.ExecContext(ctx, "INSERT into persona (nombre, apellido, edad) values(?,?,?)", "torcuato", p.apellido, 1)
	checkRollback(execErr, tx)

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

func checkRollback(execErr error, tx *sql.Tx) {
	if execErr != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("Error en update: %v, al hacer Rollback: %v\n", execErr, rollbackErr)
		}
		log.Fatalf("Fallò update %v", execErr)
	}
}
