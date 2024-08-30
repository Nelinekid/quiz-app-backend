package models

import (
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"` // or `default:gen_random_uuid()`
	UserName string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

type Score struct {
	gorm.Model
	UserName	string 	`json:"user_name"`
	Score1 	    int    	`json:"score1" gorm:"default:0"`
	Score2 	    int    	`json:"score2" gorm:"default:0"`
	Score3 	    int    	`json:"score3" gorm:"default:0"`
	Score4 	    int    	`json:"score4" gorm:"default:0"`
	ScoreT 		int  	`json:"scoret" gorm:"default:0"`
}