package db

import "time"

type AppItem struct {
	RefPointer int       `sql:"-"`
	tableName  struct{}  `sql:"apps_collection"`
	ID         int       `sql:"id,pk"`
	Name       string    `sql:"name,unique"`
	Desc       string    `sql:"desc"`
	Image      string    `sql:"image"`
	Licence    string    `sql:"licence"`
	Git        string    `sql:"git"`
	Website    string    `sql:"website"`
	Itunes     string    `sql:"itunes"`
	PlayStore  string    `sql:"play_store"`
	OS         string    `sql:"os"`
	CreatedAt  time.Time `sql:"created_at"`
	UpdatedAt  time.Time `sql:"updated_at"`
	IsActive   bool      `sql:"is_active"`
}
