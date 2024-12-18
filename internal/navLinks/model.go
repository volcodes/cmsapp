package navLinks

import "time"

// Menu represents a menu item
type NavLink struct {
	ID          int       `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Label       string    `db:"label" json:"label"`
	Destination string    `db:"destination" json:"destination"`
	Href        string    `db:"href" json:"href"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}
