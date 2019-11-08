package main

import (
	"fmt"

	"github.com/nsysu/teacher-education/src/persistence/gorm"
)

func main() {
	defer gorm.Close()

	fmt.Println("Seeding Admin")
	adminSeeder()

	fmt.Println("Seeding Student")
	studentSeeder()

	fmt.Println("Seeding Leture")
	letureSeeder()
}
