package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
} // struc opening

type OpeningResponse struct {
	ID        uint      `json: "id"`
	CreatedAt time.Time `json: "createdat"`
	UpdatedAt time.Time `json: "updatedat"`
	DeletedAt time.Time `json: "deletedat"`
	Role      string    `json: "role"`
	Company   string    `json: "company"`
	Location  string    `json: "location"`
	Remote    bool      `json: "remote"`
	Link      string    `json: "link"`
	Salary    int64     `json: "salary"`
}
