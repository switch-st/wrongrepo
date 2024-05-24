package t

import (
	"fmt"

	_ "gorm.io/driver/mysql"
)

func T() {
	fmt.Println("t")
}
