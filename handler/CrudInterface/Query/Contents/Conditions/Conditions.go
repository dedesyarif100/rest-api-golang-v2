package Conditions

import (
	"gorm.io/gorm"
)

func Conditions(db *gorm.DB, err any) {
	// String Conditions
		StringConditions(db, err)

	// Struct & Map Conditions
		// StructAndMapConditions(db, err)

	// Not Conditions
		// NotConditions(db, err)

	// Or Conditions
		// OrConditions(db, err)

	// Inline Condition
		// InlineConditions(db, err)
}