package users

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

func GetUsers(c *gin.Context) {
	c.Header(`Content-Type`, `application/json`)

	dbPool, exists := c.MustGet(`dbPool`).(*pgxpool.Pool)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{`error`: `db connection failed, no dbPool configuration`})
		return
	}

	ctx := context.Background()
	query := `SELECT id, forename, surname, dob FROM users`

	rows, err := dbPool.Query(ctx, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{`error`: err.Error()})
		return
	}

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[user])

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{`error`: err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
