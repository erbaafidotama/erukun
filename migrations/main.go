package migrations

import (
	"erukunrukun/config"
	"erukunrukun/modules/user"
)

func Migration() {
	db := config.InitDB()

	// if err = db.AutoMigrate(&Person{}, &Address{}); err != nil {
	//     log.Fatal(err)
	// }
	db.AutoMigrate(&user.UserModel{})
}
