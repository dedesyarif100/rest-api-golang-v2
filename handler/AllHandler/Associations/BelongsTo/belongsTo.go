package BelongsTo

import (
	"gorm.io/gorm"
	"rest-api-golang-v2/handler/AllHandler/Associations/BelongsTo/Contents"
)

func BelongsTo(db *gorm.DB, err any) {
	Contents.BelongsTo(db, err)
	Contents.CrudWithBelongsTo(db, err)
	Contents.EagerLoading(db, err)
	Contents.ForeignKeyConstraints(db, err)
	Contents.OverrideForeignKey(db, err)
	Contents.OverrideReferences(db, err)
}