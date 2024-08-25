package addresses

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"kia-logix/internal/user"
)

type Repo interface {
	Create(ctx context.Context, address *Address) (*Address, error)
}

type Address struct {
	ID          uint
	State       string
	City        string
	AddressLine string
	PostalCode  string
	Coordinates *Coordinates
	OwnerID     uint
	Owner       *user.User
}

type Coordinates struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func (c Coordinates) GormDataType() string {
	return "geography(POINT, 4326)"
}

func (c Coordinates) Value() (driver.Value, error) {
	return fmt.Sprintf("SRID=4326;POINT(%f %f)", c.Lng, c.Lat), nil
}

func (c *Coordinates) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("type assertion to string failed")
	}
	_, err := fmt.Sscanf(str, "POINT(%f %f)", &c.Lng, &c.Lat)
	if err != nil {
		return err
	}
	return nil
}
