package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

type persona struct {
	nombre   string
	apellido string
	edad     int
	id       int
}

type personaRepository struct {
	db *sql.DB
}

func newPersonaRepository(db *sql.DB) *personaRepository {
	return &personaRepository{db}
}

func (pR personaRepository) traeTodas() []persona {

	var p persona
	rows, err := pR.db.Query("SELECT id, edad, nombre, apellido FROM persona;")
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

func (pR personaRepository) findByID(id int) persona {
	var p persona
	//Uso exclusivamente QueryRow que es para retornar sólo 1 registro
	err := pR.db.QueryRow("SELECT id, edad, nombre, apellido FROM persona WHERE id = ?;", id).
		Scan(&p.id, &p.edad, &p.nombre, &p.apellido)
	if err != sql.ErrNoRows {
		check(err)
	}

	return p
}

func (pR personaRepository) persiste(p persona) persona {
	stmt, err := pR.db.Prepare("INSERT into persona(nombre,apellido,edad) VALUES(?,?,?);")
	check(err)

	res, err := stmt.Exec(p.nombre, p.apellido, p.edad)
	check(err)

	rowCnt, err := res.RowsAffected()
	check(err)

	fmt.Println("Modificaciones: ", rowCnt)

	return p
}

func (pR personaRepository) actualiza(p persona) persona {
	/*
	   Las parámetros van según la BD específica
	   MySQL               PostgreSQL            Oracle
	   =====               ==========            ======
	   WHERE col = ?       WHERE col = $1        WHERE col = :col
	   VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
	*/

	stmt, err := pR.db.Prepare("UPDATE persona set nombre = ?, apellido = ?, edad = ? WHERE id = ?;")
	check(err)

	res, err := stmt.Exec(p.nombre, p.apellido, p.edad, p.id)
	check(err)

	rowCnt, err := res.RowsAffected()
	check(err)

	fmt.Println("Actualizaciones: ", rowCnt)

	return p
}

// Sólo para mostrar una transacción con varias operaciones
func (pR personaRepository) operacionesComplejas(p persona) {
	ctx := context.Background()
	tx, err := pR.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	check(err)

	_, execErr := tx.ExecContext(ctx, "UPDATE persona set edad = ? WHERE id = ?;", (p.edad + 10), p.id)
	checkRollback(execErr, tx)

	_, execErr = tx.ExecContext(ctx, "INSERT into persona (nombre, apellido, edad) values(?,?,?)", "Torcuato", p.apellido, 1)
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
		log.Fatalf("Falló update %v", execErr)
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (pR personaRepository) crea() {
	stmt, err := pR.db.Prepare(`DROP TABLE IF EXISTS persona; `)
	check(err)

	_, err = stmt.Exec()
	check(err)

	stmt, err = pR.db.Prepare(`CREATE TABLE IF NOT EXISTS persona (
		id bigint(20) NOT NULL AUTO_INCREMENT,
		nombre varchar(100) NOT NULL,
		apellido varchar(100) DEFAULT NULL,
		edad int(10) DEFAULT NULL,
		PRIMARY KEY (id));`)
	check(err)

	_, err = stmt.Exec()
	check(err)
}
