package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type sls struct {
	Unique_id  int    `unique_id`
	Short_link string `short_link`
	Long_link  string `ling_link "binding:"required"`
}

var DB *sql.DB

func main() {
	connection_with_db()
	defer DB.Close()
	router := gin.Default()
	setupRoutes(router)
	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

func setupRoutes(r *gin.Engine) {
	// panic("unimplemented")

	r.POST("/short_link/create", create)
	r.GET("/:id", redirect)
}
