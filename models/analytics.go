// models/analytics.go

package models

import (
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

