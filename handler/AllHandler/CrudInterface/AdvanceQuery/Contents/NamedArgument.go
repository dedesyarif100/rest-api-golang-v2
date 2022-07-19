package Contents

import (
	"database/sql"
	"fmt"

	// "log"
	"rest-api-golang-v2/entity"

	"gorm.io/gorm"
)

func NamedArgument(db *gorm.DB, err any) {
	var users []entity.User

	db.Debug().Where("name = @nameuser OR name = @nameuser", sql.Named("nameuser", "USER B")).Find(&users)

	db.Debug().Where("name = @nameuser AND age = @ageuser AND role = @roleuser",
		sql.Named("nameuser", "USER A"),
		sql.Named("ageuser", 22),
		sql.Named("roleuser", "superadmin"),
	).Find(&users)

	db.Debug().Where("name = @nameuser AND age = @ageuser AND role = @roleuser",
		map[string]interface{}{
			"nameuser": "USER B",
			"ageuser":  22,
			"roleuser": "superadmin",
		},
	).First(&users)
	println()
	for _, data := range users {
		fmt.Println("----------------------------")
		fmt.Println("ID 		:", data.ID)
		fmt.Println("NAME 		:", data.Name)
		fmt.Println("AGE 		:", data.Age)
		fmt.Println("GENDER 		:", data.Gender)
		fmt.Println("ROLE 		:", data.Role)
		fmt.Println("----------------------------")
	}
	println()
}
