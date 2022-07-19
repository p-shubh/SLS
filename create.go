package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func create(c *gin.Context) {

	reqBody := sls{}

	// s:=(&reqBody)
	c.Bind(&reqBody)

	if reqBody.Long_link == "" {
		res := gin.H{
			"status": "link can't be empty",
		}
		c.JSON(http.StatusBadRequest, res)
	} else {
		// res := gin.H{
		// 	"status": "statusOK",
		// 	"value":  reqBody.Long_link,
		// }
		// c.JSON(http.StatusOK, res)

		// sqlStatement := (`insert into test (long_link) values ('$1')`)

		insert, _ := DB.Query(
			"INSERT INTO test (long_link) VALUES ($1)", reqBody.Long_link)

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

		sqlStatement3 := `SELECT unique_id FROM test where long_link = $1`

		row3 := DB.QueryRow(sqlStatement3, reqBody.Long_link)

		err3 := row3.Scan(&reqBody.Unique_id)

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
