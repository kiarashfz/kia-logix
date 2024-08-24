package app

import (
	"gorm.io/gorm"
	"kia-logix/config"
	"kia-logix/internal/providers"
	"kia-logix/pkg/adapters/storage"
	"kia-logix/service"
	"log"
)

type Container struct {
	cfg             config.Config
	dbConn          *gorm.DB
	providerService service.IProviderService
}

func NewAppContainer(cfg config.Config) (*Container, error) {
	app := &Container{
		cfg: cfg,
	}

	app.mustInitDB()

	app.setProviderService()

	return app, nil
}

func (a *Container) RawDBConnection() *gorm.DB {
	return a.dbConn
}

func (a *Container) mustInitDB() {
	// singleton
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

func (a *Container) setProviderService() {
	// singleton
	if a.providerService != nil { // Fixed condition
		return
	}
	a.providerService = service.NewProviderService(providers.NewOps(storage.NewProviderRepo(a.RawDBConnection())))
}

func (a *Container) ProviderService() service.IProviderService {
	return a.providerService
}
