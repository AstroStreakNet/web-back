// models/analytics.go

// model copied over from database API

package models

import (
	"webback/config"

	"gorm.io/datatypes"
)

type Analytics struct {
	ID                uint           `gorm:"primaryKey; column:analytics_id; type:INT"`
	PageViews         uint           `gorm:"column:page_views; type:INT"`
	ImageViews        uint           `gorm:"column:image_views; type:INT"`
	Clicks            uint           `gorm:"column:clicks; type:INT"`
	SessionData       string         `gorm:"column:session_data; type:TEXT"`
	ReferralSources   string         `gorm:"column:referral_sources; type:TEXT"`
	NavigationPaths   string         `gorm:"column:navigation_paths; type:TEXT"`
	AbandonmentRate   float32        `gorm:"column:abandonment_rate; type:FLOAT"`
	RepeatVisits      uint           `gorm:"column:repeat_visits; type:TINYINT"`
	TimeBetweenVisits datatypes.Time `gorm:"column:time_between_visits; type:TIME"`
	ImageID           string         `gorm:"foreignKey:image_id; column:fk_image_id; type:VARCHAR(45)"`
}

//var db *gorm.DB

func init() {
	config.Connect()
	db.AutoMigrate(&Analytics{})
}

func CreateAnalytics(analytics *Analytics) *Analytics {
	db.Create(analytics)
	return analytics
}

func GetAllAnalytics() []Analytics {
	var analytics []Analytics
	db.Find(&analytics)
	return analytics
}

func GetAnalyticsByID(ID int64) (*Analytics, error) {
	var analytics Analytics
	if err := db.First(&analytics, ID).Error; err != nil {
		return nil, err
	}
	return &analytics, nil
}

func DeleteAnalytics(ID int64) error {
	if err := db.Delete(&Analytics{}, ID).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAnalytics(analytics *Analytics) (*Analytics, error) {
	if err := db.Save(analytics).Error; err != nil {
		return nil, err
	}
	return analytics, nil
}
