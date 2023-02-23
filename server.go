package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/kleklai/todoAppv1/graph"
	"github.com/kleklai/todoAppv1/repository"
	"github.com/kleklai/todoAppv1/service"
)

const defaultPort = ":8080"

func graphHandler() gin.HandlerFunc {

	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		// Inject automatically New Repository and New Service
		Service: service.NewService(*repository.NewRepository()),
	}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {

	h := playground.Handler("GraphQL playground", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = defaultPort
	// }

	// srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{

	// 	// Inject automatically New Repository and New Service
	// 	Service: service.NewService(*repository.NewRepository()),
	// }}))

	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", srv)

	// log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// log.Fatal(http.ListenAndServe(":"+port, nil))

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/query", graphHandler())
	r.GET("/", playgroundHandler())
	err := r.Run(defaultPort)

	if err != nil {
		panic(err)
	}
}
