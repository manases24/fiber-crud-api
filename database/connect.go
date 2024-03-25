package database

import (
	"fmt"
	"strconv"

	"github.com/mnsh5/fiber-crud-api/config"
)

func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		panic("Faild to porse database port")
	}

	dsn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=require",
		config.Config("DB_HOST"),
		port,
		config.Config("DB_NAME"),
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"))
}
