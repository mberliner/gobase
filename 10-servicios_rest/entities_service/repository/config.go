package repository

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/mberliner/gobase/10-servicios_rest/entities_service/logger"
)

var (
	db *sql.DB
	//UserRepo Repositorio para manejo de acceso a datos de usuario
	UserRepo UserRepository
	//PersonaRepo Repositorio para manejo de acceso a datos de persona
	PersonaRepo PersonaRepository
)

func init() {

	initRepo()

	UserRepo = NewUserRepository(db)
	PersonaRepo = NewPersonaRepository(db)

}

func initRepo() {
	var err error
	db, err = sql.Open("mysql", "test:test987@tcp(localhost:3306)/go_test?charset=utf8&parseTime=true")
	if err != nil {
		logger.Error("Error en Open:", err)
		panic(err)
	}

	if err = db.Ping(); err != nil {
		logger.Error("Error Ping:", err)
		panic(err)
	}
	err = mysql.SetLogger(logger.GetLogger())
	if err != nil {
		logger.Error("Error seteando logger a Mysql", err)
	}

	logger.Info("Conectado a Mysql")
}
