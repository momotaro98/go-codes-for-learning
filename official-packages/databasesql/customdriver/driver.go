package customdriver

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"os"
)

type CustomConn struct {
	name string
}

// implements sql.Conn interface

func (c *CustomConn) Begin() (driver.Tx, error) {
	return nil, fmt.Errorf("Begin method not implemented")
}

// Prepare -
func (c *CustomConn) Prepare(query string) (driver.Stmt, error) {
	return nil, fmt.Errorf("Prepare method not implemented")
}

// Close -
func (c *CustomConn) Close() error {
	return nil
}

type CustomDriver struct{}

// implements database/sql/driver.Driver interface
// Definition is here https://github.com/golang/go/blob/e5da18df52/src/database/sql/driver/driver.go#L94
func (d *CustomDriver) Open(name string) (driver.Conn, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	fmt.Printf("[custom]: open(%s)", name)
	return &CustomConn{name}, nil
}

// import _ "customdriver"
func init() {
	// mysql one https://github.com/go-sql-driver/mysql/blob/6be42e0ff99645d7d9626d779001a46e39c5f280/driver.go#L168
	sql.Register("custom", &CustomDriver{})
	// List registered drivers
	fmt.Printf("Drivers=%+v\n", sql.Drivers())
}
