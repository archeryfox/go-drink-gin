package repository

import "time"

// UserModel represents a user/customer
// swagger:model UserModel
type UserModel struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"size:100;unique;not null" json:"username"`
	Email    string `gorm:"size:200;unique;not null" json:"email"`
}

// FetishModel represents a fetish/tag
// swagger:model FetishModel
type FetishModel struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size:100;unique;not null" json:"name"`
}

// ProductFetishModel many-to-many products <-> fetishes
// swagger:model ProductFetishModel
type ProductFetishModel struct {
	ProductID uint `gorm:"primaryKey" json:"product_id"`
	FetishID  uint `gorm:"primaryKey" json:"fetish_id"`
}

// UserFetishModel user preferences
// swagger:model UserFetishModel
type UserFetishModel struct {
	UserID   uint `gorm:"primaryKey" json:"user_id"`
	FetishID uint `gorm:"primaryKey" json:"fetish_id"`
}

// LikeModel represents a like by a user on a product
// swagger:model LikeModel
type LikeModel struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"index;not null" json:"user_id"`
	ProductID uint      `gorm:"index;not null" json:"product_id"`
	CreatedAt time.Time `json:"created_at"`
}

// RecommendationModel represents a recommended product for a user
// swagger:model RecommendationModel
type RecommendationModel struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"index;not null" json:"user_id"`
	ProductID uint      `gorm:"index;not null" json:"product_id"`
	Score     float64   `json:"score"`
	CreatedAt time.Time `json:"created_at"`
}

// NotificationModel represents notifications for users
// swagger:model NotificationModel
type NotificationModel struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"index;not null" json:"user_id"`
	Title     string    `gorm:"size:255" json:"title"`
	Body      string    `gorm:"type:text" json:"body"`
	Read      bool      `gorm:"default:false" json:"read"`
	CreatedAt time.Time `json:"created_at"`
}

// ReviewModel simple product review
// swagger:model ReviewModel
type ReviewModel struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"index;not null" json:"user_id"`
	ProductID uint      `gorm:"index;not null" json:"product_id"`
	Rating    int       `gorm:"not null" json:"rating"`
	Text      string    `gorm:"type:text" json:"text"`
	CreatedAt time.Time `json:"created_at"`
}
