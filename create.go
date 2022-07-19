package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	layoutISO = "15:04:05"
	layoutUS  = "January 2, 2006"
)

func create(c *gin.Context) {

	reqBody := sls{}

	// s:=(&reqBody)
	err := c.Bind(&reqBody)
	if err != nil {
		res := gin.H{
			"message": "invalid req body",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	time := time.Now().Add(time.Minute * time.Duration(reqBody.ExpiresIn))

	if reqBody.Long_link == "" {
		res := gin.H{
			"status": "link can't be empty",
		}
		c.JSON(http.StatusBadRequest, res)
	} else {

		insert, _ := DB.Query(
			"INSERT INTO test (long_link,timer) VALUES ($1,$2)", reqBody.Long_link, time)

		// if there is an error inserting, handle it
		if insert == nil {
			panic(insert.Err())
		}

		res := gin.H{
			"status": "statusOK",
			"value":  insert,
		}
		c.JSON(http.StatusOK, res)

		// defer insert.Close()

		sqlStatement3 := `SELECT unique_id,timer FROM test where long_link = $1`

		row3 := DB.QueryRow(sqlStatement3, reqBody.Long_link)

		err3 := row3.Scan(&reqBody.Unique_id, &reqBody.Timer)

		fmt.Println(reqBody.Timer)

		// reqBody.Short_link = ("localhost:8080/" + strconv.Itoa(reqBody.Unique_id))

		if err3 != nil {
			res := gin.H{
				"error":   err3.Error(),
				"message": "sqlstatement3",
				"result":  sqlStatement3,
			}
			c.JSON(http.StatusBadRequest, res)
			c.Abort()
			return
		} else {
			res := gin.H{
				"result": reqBody.Unique_id,
				// "value": reqBody.Short_link,
			}
			// c.JSON(http.StatusOK, res)

			c.JSON(http.StatusFound, res)
		}
	}
}

// }
