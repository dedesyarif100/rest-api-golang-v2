package AssociationMode

import (
	"gorm.io/gorm"
	"rest-api-golang-v2/handler/AllHandler/Associations/AssociationMode/Contents"
)

func AssociationMode(db *gorm.DB, err any) {
	Contents.AssociationMode(db, err)
	Contents.AssociationTags(db, err)
	Contents.AutoCreateUpdate(db, err)
	Contents.DeleteWithSelect(db, err)
	Contents.SelectOmitAssociationfields(db, err)
	Contents.SkipAutoCreateUpdate(db, err)
}
