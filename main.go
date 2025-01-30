package main

import (
	"fmt"
	"log"
	"os"

	server "github.com/ZolaraProject/pki-vault-service/pkivault"
)

const (
	defaultDbPort     string = "5432"
	defaultDbHostname string = "postgres.database.svc.cluster.local"
	defaultDbUser     string = "postgres"
)

func main() {
	var ok bool
	server.DbPort, ok = os.LookupEnv("DB_PORT")
	if !ok {
		log.Printf("Warning: could not read $DB_PORT, starting server with default DB port (%s)", defaultDbPort)
		server.DbPort = defaultDbPort
	}
	server.DbPort, ok = os.LookupEnv("DB_PORT_SINGLE")
	if !ok {
		log.Printf("Warning: could not read $DB_PORT_SINGLE, starting server with default DB port (%s)", defaultDbPort)
		server.DbPort = defaultDbPort
	}
	server.DbHostname, ok = os.LookupEnv("DB_HOSTNAME")
	if !ok {
		log.Printf("Warning: could not read $DB_HOSTNAME, starting server with default DB hostname (%s)", defaultDbHostname)
		server.DbHostname = defaultDbHostname
	}
	dbUser, err := os.ReadFile("/root/trinityDataMaster/userName")
	if err != nil {
		log.Printf("Warning: could not read DB user, starting server with default DB user (%s)", defaultDbUser)
		server.DbUser = defaultDbUser
	} else {
		server.DbUser = fmt.Sprintf("%s", dbUser)
	}
	dbPassword, err := os.ReadFile("/")
	if err != nil {
		log.Fatal("Fatal: could not read DB password")
	}
	server.DbPassword = fmt.Sprintf("%s", dbPassword)

	server.Run()
}
