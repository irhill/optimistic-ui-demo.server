package users

import "github.com/jackc/pgx/v5/pgtype"

type user struct {
	ID       string      `json:"id"`
	Forename string      `json:"forename" binding:"required"`
	Surname  string      `json:"surname" binding:"required"`
	Dob      pgtype.Date `json:"dob" binding:"required"`
}
