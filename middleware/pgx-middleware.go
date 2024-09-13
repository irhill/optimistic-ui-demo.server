package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

// this would all come from config in the real world, but let's hardcode it for this demo
var connectionString = "postgresql://seedlegals:seedlegals@localhost:5432/optimistic?sslmode=disable"

type PgxMiddleware struct {
	Pool *pgxpool.Pool
}

func NewPgxMiddleware() (*PgxMiddleware, error) {
	config, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return &PgxMiddleware{Pool: pool}, nil
}

func (pgxm *PgxMiddleware) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("dbPool", pgxm.Pool)
		c.Next()
	}
}

func (pgmx *PgxMiddleware) Close() {
	if pgmx.Pool != nil {
		log.Println("closing connection to pool...")
		pgmx.Pool.Close()
	}
}
