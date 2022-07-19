package AdvanceQuery

import (
	"rest-api-golang-v2/handler/AllHandler/CrudInterface/AdvanceQuery/Contents"
	"gorm.io/gorm"
)

func AdvanceQuery(db *gorm.DB, err any) {
	// ADVANCE QUERY
	// Smart Select Fields
	// Contents.SmartSelectFields(db, err)

	// SubQuery
	// Contents.SubQuery(db, err)

	// Group Conditions
	// Contents.GroupConditions(db, err)

	// IN with multiple columns
	// Contents.InWithMultipleColumns(db, err)

	// Named Argument
	// Contents.NamedArgument(db, err)

	// Find To Map
	// Contents.FindToMap(db, err)

	// FirstOrInit
	// Contents.FirstOrInit(db, err)

	// FirstOrCreate
	// Contents.FirstOrCreate(db, err)

	// Iteration
	// Contents.Iteration(db, err)

	// FindInBatches
	// Contents.FindInBatches(db, err)

	// Count
	// Contents.Count(db, err)

	// Pluck
	Contents.Pluck(db, err)
}
