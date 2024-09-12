package users

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

func DeleteUser(c *gin.Context) {

	dbPool, exists := c.MustGet("dbPool").(*pgxpool.Pool)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db connection failed, no dbPool configuration"})
		return
	}

	ctx := context.Background()
	query := "DELETE FROM users WHERE id = $1"
	id := c.Param("id")

	tag, err := dbPool.Exec(ctx, query, id)
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
