package HasOne

import (
	"gorm.io/gorm"
	"rest-api-golang-v2/handler/AllHandler/Associations/HasOne/Contents"
)

func HasOne(db *gorm.DB, err any) {
	Contents.CrudWithHasOne(db, err)
	Contents.EagerLoading(db, err)
	Contents.ForeignKeyConstraints(db, err)
	Contents.HasOne(db, err)
	Contents.OverrideForeignKey(db, err)
	Contents.OverrideReferences(db, err)
	Contents.PolymorphismAssociation(db, err)
	Contents.SelfReferentialHasOne(db, err)
}