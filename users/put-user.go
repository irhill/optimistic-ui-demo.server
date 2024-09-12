package users

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

type PutUserRequest struct {
	Forename string      `json:"forename" binding:"required"`
	Surname  string      `json:"surname" binding:"required"`
	Dob      pgtype.Date `json:"dob" binding:"required"`
}

func PutUser(c *gin.Context) {
	var request PutUserRequest

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
	query := "UPDATE users set forename = $2, surname = $3, dob = $4 where id = $1"
	id := c.Param("id")

	tag, err := dbPool.Exec(ctx, query, id, request.Forename, request.Surname, request.Dob)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if tag.RowsAffected() == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no rows were updated, check request id"})
		return
	}

	c.Status(http.StatusAccepted)
}
