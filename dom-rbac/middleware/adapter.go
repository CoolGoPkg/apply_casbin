package middleware

import (
	"fmt"

	gormad "github.com/casbin/gorm-adapter/v3"
)

// GetAdapter doc
func GetAdapter() (*gormad.Adapter, error) {
	source := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		"127.0.0.1", 8412, "postgres", "casbin", "1", "disable")
	a, err := gormad.NewAdapter("postgres", source)
	if err != nil {
		fmt.Println("new adapter err :", err)
		return nil, err
	}
	return a, err
}
