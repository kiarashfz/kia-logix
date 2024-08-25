package app

import (
	"gorm.io/gorm"
	"kia-logix/config"
	"kia-logix/internal/addresses"
	"kia-logix/internal/orders"
	"kia-logix/internal/providers"
	"kia-logix/internal/user"
	"kia-logix/pkg/adapters/storage"
	"kia-logix/service"
	"log"
)

type Container struct {
	cfg             config.Config
	dbConn          *gorm.DB
	authService     service.IAuthService
	providerService service.IProviderService
	addressService  service.IAddressService
	orderService    service.IOrderService
}

func NewAppContainer(cfg config.Config) (*Container, error) {
	app := &Container{
		cfg: cfg,
	}

	app.mustInitDB()

	app.setAuthService()
	app.setProviderService()
	app.setAddressService()
	app.setOrderService()

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

func (a *Container) setAuthService() {
	// singleton
	if a.authService != nil {
		return
	}
	a.authService = service.NewAuthService(user.NewOps(storage.NewUserRepo(a.RawDBConnection())), a.cfg.Auth)
}

func (a *Container) GetAuthService() service.IAuthService {
	return a.authService
}

func (a *Container) setProviderService() {
	// singleton
	if a.providerService != nil {
		return
	}
	a.providerService = service.NewProviderService(providers.NewOps(storage.NewProviderRepo(a.RawDBConnection())))
}

func (a *Container) GetProviderService() service.IProviderService {
	return a.providerService
}

func (a *Container) setOrderService() {
	// singleton
	if a.orderService != nil {
		return
	}
	a.orderService = service.NewOrderService(orders.NewOps(storage.NewOrderRepo(a.RawDBConnection())))
}

func (a *Container) GetOrderService() service.IOrderService {
	return a.orderService
}

func (a *Container) setAddressService() {
	// singleton
	if a.addressService != nil {
		return
	}
	a.addressService = service.NewAddressService(addresses.NewOps(storage.NewAddressRepo(a.RawDBConnection())))
}

func (a *Container) GetAddressService() service.IAddressService {
	return a.addressService
}
