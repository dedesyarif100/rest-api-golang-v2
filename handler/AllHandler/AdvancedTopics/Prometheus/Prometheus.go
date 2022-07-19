package Prometheus

import (
	"gorm.io/gorm"
	"rest-api-golang-v2/handler/AllHandler/AdvancedTopics/Prometheus/Contents"
)

func Prometheus(db *gorm.DB, err any) {
	Contents.Usage(db, err)
	Contents.UserDefinedMetrics(db, err)
}