package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func main() {
	appConfig, err := envConfigVariable("config.yaml")
	if err != nil {
		log.Println("config file not found")
		return
	}

	path := "db"

	db, err := sql.Open("postgres", "postgresql://"+
		appConfig.GetString(fmt.Sprintf("%s.username", path))+":"+
		appConfig.GetString(fmt.Sprintf("%s.password", path))+"@"+
		appConfig.GetString(fmt.Sprintf("%s.host", path))+":"+
		appConfig.GetString(fmt.Sprintf("%s.port", path))+"/"+
		appConfig.GetString(fmt.Sprintf("%s.schema", path))+"?sslmode=disable")
	if err != nil {
		log.Println("error opening database", err)
		return
	}

	schemaDatabase := flag.String("schema", "default value", "a string for description")
	migrationFolder := flag.String("folder", "default value", "a string for description")
	flag.Parse()

	driver, err := postgres.WithInstance(db, &postgres.Config{
		SchemaName: *schemaDatabase,
	})
	if err != nil {
		log.Println("error opening instance", err)
		return
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./scripts/pgbo/migrations/"+*migrationFolder,
		appConfig.GetString(fmt.Sprintf("%s.schema", path)), driver)

	if err != nil {
		log.Println("error migrate database", err)
		return
	}

	err = m.Up()
	if err != nil {
		log.Println("error migrate database", err)
		return
	}
}

func envConfigVariable(filePath string) (cfg *viper.Viper, err error) {
	cfg = viper.New()
	cfg.SetConfigFile(filePath)

	if err = cfg.ReadInConfig(); err != nil {
		err = errors.Wrap(err, "Error while reading config file")

		return
	}

	return
}
