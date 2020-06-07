package main

import (
	"log"
	"os"
	"tdd/Database/mysql"
	"tdd/Tools/jwt"
	"tdd/router"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	db := getDatabase()
	jwtTool := getJwtTool()

	router.Start(db, jwtTool)
}

func getDatabase() *mysql.Database {
	db, err := mysql.ConnectDB()

	if err != nil {
		log.Panic("can't connect database mysql", err.Error())
	}

	return db
}

func getJwtTool() jwt.JWT {
	secret := os.Getenv("JWT_SECRET")

	if secret == "" {
		log.Panic("can't got jwt secret in env")
	}

	return jwt.JWT{Secret: secret}
}
