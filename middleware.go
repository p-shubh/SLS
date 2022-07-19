package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Filter(t time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		finish := make(chan struct{})

		go func() {
			c.Next()
			finish <- struct{}{}
		}()

		select {
		case <-time.After(t):
			c.JSON(504, "timeout")
			c.Abort()
		case <-finish:
		}
	}
}

func time_expire(c *gin.Context) {

	reqBody := sls{}
	id, _ := c.Params.Get("id")
	reqBody.Unique_id, _ = strconv.Atoi(id)

	// reqBody.Unique_id = assing

	// present := time.Add(2)

	sqlStatement := (`select timer from test where unique_id = $1`)

	row3 := DB.QueryRow(sqlStatement, reqBody.Unique_id)

	err3 := row3.Scan(&reqBody.Timer)

	if err3 != nil {
		fmt.Println("fail in the middleware")
	} else {
		// res := gin.H{
		// 	"result": reqBody.Unique_id,
		// 	// "value": reqBody.Short_link,
		fmt.Println("got reqBody.Timer")

	}

	fmt.Println("reqBody.Timer := ", reqBody.Timer)

	// fmt.Println("present_time  := ", present_time)

	if reqBody.Timer.After(time.Now()) {
		// link has expired
		res := gin.H{
			"message": "time expires",
		}

		c.JSON(http.StatusOK, res)
	}
	// check_time := reqBody.Timer.Add(time.Minute * 1)

	// if present_time < check_time {

	// }

}
