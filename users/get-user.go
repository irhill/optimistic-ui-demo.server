package users

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

func GetUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	dbPool, exists := c.MustGet("dbPool").(*pgxpool.Pool)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db connection failed, no dbPool configuration"})
		return
	}

	ctx := context.Background()
	query := "SELECT id, forename, surname, dob FROM users WHERE id=$1"
	id := c.Param("id")

	var user user
	err := dbPool.QueryRow(ctx, query, id).Scan(&user.ID, &user.Forename, &user.Surname, &user.Dob)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying data"})
		return
	}

	c.JSON(http.StatusOK, user)
}
