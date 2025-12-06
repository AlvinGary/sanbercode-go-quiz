package structs

import (
	"time"
)

type Kategori struct {
	Id         int
	Name       string
	CreatedAt  time.Time
	CreatedBy  string
	ModifiedAt time.Time
	ModifiedBy string
}

type Buku struct {
	Id          int
	Title       string
	Description string
	ImageUrl    string
	ReleaseYear int
	Price       int
	TotalPage   int
	Thickness   string
	CategoryId  int
	CreatedAt   time.Time
	CreatedBy   string
	ModifiedAt  time.Time
	ModifiedBy  string
}