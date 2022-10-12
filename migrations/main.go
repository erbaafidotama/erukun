package migrations

import (
	"erukunrukun/config"
	"erukunrukun/modules/lookup"
	"erukunrukun/modules/lookupItem"
	"erukunrukun/modules/user"
)

func Migration() {
	db := config.InitDB()

	// if err = db.AutoMigrate(&Person{}, &Address{}); err != nil {
	//     log.Fatal(err)
	// }

	// db.AutoMigrate(&user.UserModel{})

	db.AutoMigrate(&user.UserModel{}, &lookup.LookupModel{}, &lookupItem.LookupItemModel{})
}
