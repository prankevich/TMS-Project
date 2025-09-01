package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type Num struct {
	A int `json:"a"`
	B int `json:"b"`
}

func main() {
	r := gin.Default()
	r.GET("/hello", helloHandler)
	r.POST("/divide", divide)
	err := r.Run(":8181")
	if err != nil {
		panic(err)
	}

}
func helloHandler(c *gin.Context) {
	name := c.Query("name")
	c.JSON(200, gin.H{"msg": "Привет, " + name})
}
func divide(c *gin.Context) {
	var num Num
	err := c.BindJSON(&num)
	if err != nil {
		c.JSON(400, gin.H{"err": "Ошибка чтения"})
		return
	}
	if num.B == 0 {
		c.JSON(400, gin.H{"err": "На ноль делить нельзя"})
		return
	}
	result := num.A / num.B
	c.JSON(200, gin.H{"msg": "result: " + strconv.Itoa(result)})
}
