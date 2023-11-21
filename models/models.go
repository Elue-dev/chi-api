package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title string `json:"title"`
	Desc string `json:"desc"`
	Category string `json:"category"`
}

type CustomPost struct {
	ID        uint      	  `json:"id"`
	CreatedAt time.Time 	  `json:"created_at"`
	UpdatedAt time.Time 	  `json:"updated_at"`
	DeletedAt gorm.DeletedAt  `json:"deleted_at"`
	Title string 			   `json:"title"`
	Desc string 			   `json:"desc"`
	Category string 		   `json:"category"`
}

