package ManyToMany

import (
	"gorm.io/gorm"
	"rest-api-golang-v2/handler/AllHandler/Associations/ManyToMany/Contents"
)

func ManyToMany(db *gorm.DB, err any) {
	Contents.BackReference(db, err)
	Contents.CompositeForeignKeys(db, err)
	Contents.CrudWithManyToMany(db, err)
	Contents.CustomizeJoinTable(db, err)
	Contents.EagerLoading(db, err)
	Contents.ForeignKeyConstraints(db, err)
	Contents.ManyToMany(db, err)
	Contents.OverrideForeignKey(db, err)
	Contents.SelfReferentialManyToMany(db, err)
}