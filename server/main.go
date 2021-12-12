package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

var todoList = []map[string]interface{}{
	{"id": "0", "todo": "Buy bread"},
}

func main() {
	r := gin.Default()
	r_Cors := r.Group("/", Cors())
	r_Cors.GET("/", showTodoList)
	r_Cors.POST("/", createTodo)
	r_Cors.POST("/deleteAll", deleteAllTodo)
	r_Cors.POST("/delete", deleteTodo)
	r_Cors.POST("/update", updateTodo)
	r.Run()
}

var todoID = 1

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		}
		c.Next()
	}
}

func updateTodo(c *gin.Context) {
	if id := c.Request.FormValue("id"); id == "" {
		c.String(500, fmt.Sprint("Error:", "id is empty"))
	} else if todo := c.Request.FormValue("todo"); todo == "" {
		c.String(500, fmt.Sprint("Error:", "todo is empty"))
	} else {
		for i, v := range todoList {
			if v["id"] == id {
				todoList[i]["todo"] = todo
				c.JSON(200, todoList[i])
				return
			}
		}
		c.String(500, fmt.Sprint("Error:", "id not found"))
	}
}

func deleteTodo(c *gin.Context) {
	if id := c.Request.FormValue("id"); id == "" {
		c.String(500, fmt.Sprint("Error:", "id is empty"))
	} else {
		for i, v := range todoList {
			if v["id"] == id {
				todoList = append(todoList[:i], todoList[i+1:]...)
				fmt.Println(todoList)
				c.String(200, fmt.Sprint("Success"))
				return
			}
		}
		c.String(500, fmt.Sprint("Error:", "id not found"))
	}
}

func deleteAllTodo(c *gin.Context) {
	todoList = []map[string]interface{}{}
	c.String(200, fmt.Sprint("Success"))
}

func showTodoList(c *gin.Context) {
	c.JSON(200, todoList)
}

func createTodo(c *gin.Context) {
	if todo := c.Request.FormValue("todo"); todo == "" {
		c.String(500, fmt.Sprint("Error:", "todo is empty"))
	} else {
		newTodo := map[string]interface{}{"id": strconv.Itoa(todoID), "todo": todo}
		todoList = append(todoList, newTodo)
		todoID++
		fmt.Println(todoList)
		c.JSON(200, newTodo)
	}
}
