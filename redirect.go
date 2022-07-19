package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func redirect(c *gin.Context) {

	reqBody := sls{}

	id, _ := c.Params.Get("id")

	// (&id)
	fmt.Println("myid", id)

	// strId._ := strconv.Atoi(id)

	sqlStatement := `select long_link from test where unique_id = $1`

	// c.Query.row3(sqlStatement,id)

	r := DB.QueryRow(sqlStatement, id)

	r.Scan(&reqBody.Long_link)

	fmt.Println(reqBody.Long_link)

	c.Redirect(http.StatusMovedPermanently, reqBody.Long_link)

}
