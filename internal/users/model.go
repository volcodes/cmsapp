package users

import "time"

type User struct {
	ID          int       `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Email       string    `db:"email" json:"email"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt   time.Time `db:"deleted_at" json:"deleted_at"`
	Status      string    `db:"status" json:"status"`
	Role        string    `db:"role" json:"role"`
	Password    string    `db:"password" json:"password"`
	Phone       string    `db:"phone" json:"phone"`
	Address     string    `db:"address" json:"address"`
	Avatar      string    `db:"avatar" json:"avatar"`
	IsActive    bool      `db:"is_active" json:"is_active"`
	IsDeleted   bool      `db:"is_deleted" json:"is_deleted"`
	IsVerified  bool      `db:"is_verified" json:"is_verified"`
	IsBlocked   bool      `db:"is_blocked" json:"is_blocked"`
	IsSuspended bool      `db:"is_suspended" json:"is_suspended"`
}
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
