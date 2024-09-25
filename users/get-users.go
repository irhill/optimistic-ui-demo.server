package users

import (
	"context"
	"fmt"
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

	users, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[user])

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{`error`: err.Error()})
		return
	}

	for idx := range users {
		user := &users[idx]
		fullName := fmt.Sprintf("%s %s", user.Forename, user.Surname)
		user.FullName = fullName
	}

	c.JSON(http.StatusOK, users)
}
