package customdriver

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

type CustomConn struct {
	name string
}

type CustomDriver struct{}

// implements database/sql/driver.Driver interface
// Definition is here https://github.com/golang/go/blob/e5da18df52/src/database/sql/driver/driver.go#L94
func (d *CustomDriver) Open(name string) (driver.Conn, error) {
	fmt.Printf("[custom]: open(%s)", name)
	return nil, nil
}

// import _ "customdriver"
func init() {
	// mysql one https://github.com/go-sql-driver/mysql/blob/6be42e0ff99645d7d9626d779001a46e39c5f280/driver.go#L168
	sql.Register("custom", &CustomDriver{})
	// List registered drivers
	fmt.Printf("Drivers=%+v\n", sql.Drivers())
}
