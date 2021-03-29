package main

import (
	"log"
	"smart_pharmacy_api/app"

	cr "github.com/brkelkar/common_utils/configreader"
	db "github.com/brkelkar/common_utils/databases"
	"gorm.io/gorm"
)

var (
	envVerables = []string{"SERVER_PORT", "SERVER_HOST", "AWACS_DB", "AWACS_SMART_DB", "AWACS_SMART_STOCKIST_DB",
		"DB_PORT", "DB_HOST", "DB_USERNAME", "DB_PASSWORD", "GRPC_SERVER", "GRPC_CLIENT"}
	err error
)

func main() {
	log.Println("Starting Sync upload service")

	var cfg cr.Config
	//Read configuration from config.yml file
	cfg.ReadFile("config.yml")

	//Read enviroment variables
	m := cfg.ReadEnv(envVerables)

	//Over write config file variables if enviroment variable is set
	cfg.MapEnv(m)

	// Create connention map incase of multiple db
	// map will hold value for *db with respective string identifier
	var dbObj db.DbObj

	db.DB = make(map[string]*gorm.DB, 1)
	db.DB["smartdb"], err = dbObj.GetConnection("smartdb", cfg)
	if err != nil {
		log.Printf("Error while connecting to smartdb : %v\n", err)
		log.Fatal(err)
	}

	dbObj.CreateConnectionPool(db.DB["smartdb"], 10, 100, 60)
	log.Println("Connection successful smartdb")

	// db.DB["awacs"], err = dbObj.GetConnection("awacs", cfg)
	// if err != nil {
	// 	log.Printf("Error while connecting to awacs : %v\n", err)
	// 	log.Fatal(err)
	// }

	// dbObj.CreateConnectionPool(db.DB["awacs"], 10, 100, 60)
	// log.Printf("Connection successful awacs")

	app.StartApplication(cfg)
}
