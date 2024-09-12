package users

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

type PostUserRequest struct {
	Forename string      `json:"forename" binding:"required"`
	Surname  string      `json:"surname" binding:"required"`
	Dob      pgtype.Date `json:"dob" binding:"required"`
}

type PostUserResponse struct {
	Id   int    `json:"id"`
	Link string `json:"link"`
}

func PostUser(c *gin.Context) {
	var request PostUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbPool, exists := c.MustGet("dbPool").(*pgxpool.Pool)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db connection failed, no dbPool configuration"})
		return
	}

	ctx := context.Background()
	query := "INSERT INTO users (forename, surname, dob) VALUES ($1, $2, $3) RETURNING id"

	var id int
	err := dbPool.QueryRow(ctx, query, request.Forename, request.Surname, request.Dob).Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := PostUserResponse{Id: id, Link: fmt.Sprintf("/users/%d", id)}
	c.JSON(http.StatusCreated, response)
}
