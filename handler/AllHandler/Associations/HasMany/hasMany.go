package HasMany

import (
	"gorm.io/gorm"
	"rest-api-golang-v2/handler/AllHandler/Associations/HasMany/Contents"
)

func HasMany(db *gorm.DB, err any) {
	Contents.CrudWithHasMany(db, err)
	Contents.EagerLoading(db, err)
	Contents.ForeignKeyConstraints(db, err)
	Contents.HasMany(db, err)
	Contents.OverrideForeignKey(db, err)
	Contents.OverrideReferences(db, err)
	Contents.PolymorphismAssociation(db, err)
	Contents.SelfReferentialHasMany(db, err)
}