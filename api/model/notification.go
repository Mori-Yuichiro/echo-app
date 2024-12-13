package model

type Notification struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	VisitorId uint   `json:"visitor_id" gorm:"not null"`
	VisitedId uint   `json:"visited_id" gorm:"not null"`
	TweetId   *uint  `json:"tweet_id"`
	Action    string `json:"action" gorm:"not null; default:''"`
	Read      bool   `json:"read" gorm:"not null; default:false"`
	Visitor   User   `json:"visitor" gorm:"foreignKey:VisitorId; constraint:OnDelete:CASCADE"`
	Visited   User   `json:"visited" gorm:"foreignKey:VisitedId; constraint:OnDelete:CASCADE"`
	Tweet     Tweet  `json:"tweet" gorm:"foreignKey:TweetId; constraint:OnDelete:CASCADE"`
}

type NotificationResponse struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	VisitorId uint   `json:"visitor_id" gorm:"not null"`
	VisitedId uint   `json:"visited_id" gorm:"not null"`
	TweetId   uint   `json:"tweet_id"`
	Action    string `json:"action" gorm:"not null; default:''"`
	Read      bool   `json:"read" gorm:"not null; default:false"`
}
