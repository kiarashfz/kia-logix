package app

import (
	"gorm.io/gorm"
	"kia-logix/config"
	"kia-logix/pkg/adapters/storage"
	"log"
)

type Container struct {
	cfg    config.Config
	dbConn *gorm.DB
}

func NewAppContainer(cfg config.Config) (*Container, error) {
	app := &Container{
		cfg: cfg,
	}

	app.mustInitDB()

	return app, nil
}

func (a *Container) RawDBConnection() *gorm.DB {
	return a.dbConn
}

func (a *Container) mustInitDB() {
	if a.dbConn != nil {
		return
	}

	db, err := storage.NewPostgresGormConnection(a.cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	a.dbConn = db

	err = storage.Migrate(a.dbConn)
	if err != nil {
		log.Fatal("Migration failed: ", err)
	}
}
