package db

import "time"

type SoftwareItem struct {
	RefPointer int       `sql:"-"`
	tableName  struct{}  `sql:"software_collection"`
	ID         int       `sql:"id,pk" json:"id"`
	Name       string    `sql:"name,unique" json:"name"`
	Desc       string    `sql:"desc" json:"desc"`
	Image      string    `sql:"image" json:"image"`
	Licence    string    `sql:"licence" json:"licence"`
	Git        string    `sql:"git" json:"git"`
	Website    string    `sql:"website" json:"website"`
	OS         string    `sql:"os" json:"os"`
	CreatedAt  time.Time `sql:"created_at" json:"created_at"`
	UpdatedAt  time.Time `sql:"updated_at" json:"updated_at"`
	IsActive   bool      `sql:"is_active" json:"is_active"`
}
