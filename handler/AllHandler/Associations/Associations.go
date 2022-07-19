package Associations

import (
	"rest-api-golang-v2/handler/AllHandler/Associations/AssociationMode"
	"rest-api-golang-v2/handler/AllHandler/Associations/BelongsTo"
	"rest-api-golang-v2/handler/AllHandler/Associations/HasMany"
	"rest-api-golang-v2/handler/AllHandler/Associations/HasOne"
	"rest-api-golang-v2/handler/AllHandler/Associations/ManyToMany"
	"rest-api-golang-v2/handler/AllHandler/Associations/Preloading"
	"gorm.io/gorm"
)

func Associations(db *gorm.DB, err any) {
	
	AssociationMode.AssociationMode(db, err)

	BelongsTo.BelongsTo(db, err)
	
	HasMany.HasMany(db, err)
	
	HasOne.HasOne(db, err)
	
	ManyToMany.ManyToMany(db, err)
	
	Preloading.Preloading(db, err)

}
