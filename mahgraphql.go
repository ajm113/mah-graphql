package main

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ajm113/mah-graphql/graph"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	_ "github.com/uptrace/bun/driver/pgdriver"
)

func graphqlHandler(db *bun.DB) gin.HandlerFunc {
	h := handler.NewDefaultServer(
		graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DB: db}}),
	)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func main() {
	c, err := loadConfig("config.yml")

	if err != nil {
		log.Fatal().Err(err).Msg("failed loading config")
	}

	dsn := "postgres://" + c.PG.User + ":" + c.PG.Password + "@" + c.PG.Host + ":" + strconv.FormatInt(int64(c.PG.Port), 10) + "/" + c.PG.Database + "?sslmode=" + c.PG.SSLMode

	sqldb, err := sql.Open("pg", dsn)

	if err != nil {
		log.Fatal().Err(err).Msg("failed connecting to PG")
	}

	db := bun.NewDB(sqldb, sqlitedialect.New())

	// Ping the database to make sure it works.
	log.Debug().Msg("pinging PG database")
	err = db.Ping()
	if err != nil {
		log.Fatal().Err(err).Msg("failed pinging database")
	}

	r := gin.New()
	r.Use(GinContextToContextMiddleware())

	r.POST("/query", graphqlHandler(db))
	r.GET("/", playgroundHandler())
	r.Run(c.Server.Addr)
}
