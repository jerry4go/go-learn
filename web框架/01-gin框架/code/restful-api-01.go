package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 配置mysql连接池
	db, err := sql.Open("mysql", "root:a123456@tcp(192.168.xx.xx:3306)/gintestdb?parseTime=True")

	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
    // 配置mysql最大空闲连接数
	db.SetMaxIdleConns(20)
	// 配置mysql最大并发连接数
	db.SetMaxOpenConns(20)
    // 验证mysql连接串是否配置正确
	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}

	// gin 框架使用
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "It works\n")
	})

	// 插入数据  POST
	router.POST("/person", func(c *gin.Context) {
		firstName := c.Request.FormValue("first_name")
		lastName := c.Request.FormValue("last_name")
		rs, err := db.Exec("INSERT INTO person(first_name,last_name) VALUES(?,?)", firstName, lastName)
		if err != nil {
			log.Fatalln(err)
		}
		id, err := rs.LastInsertId()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("insert person ID {}", id)
		msg := fmt.Sprintf("insert successful %d", id)
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
	})

	// 查询表的所有数据   GET
	type Person struct {
		Id        int    `json:"id" form:"id"`
		FirstName string `json:"first_name" form:"first_name"`
		LastName  string `json:"last_name" form:"last_name"`
	}

	router.GET("/persons", func(c *gin.Context) {
	    // 返回一个rows对象
		rows, err := db.Query("SELECT id,first_name,last_name FROM person")

		if err != nil {
			log.Fatalln(err)
		}
		defer rows.Close()
		// make 声明对象，当切片没有元素时，JSON会返回[]
		persons := make([]Person, 0)
		// 直接声明切片，当切片没有元素时，JSON会返回null
		//var persons []Person

		// 通过Next()方法遍历每行数据，并且绑定到person对象
		// 最后append到persons切片
		for rows.Next() {
			var person Person
			rows.Scan(&person.Id, &person.FirstName, &person.LastName)
			persons = append(persons, person)
		}
		if err = rows.Err(); err != nil {
			log.Fatalln(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"persons": persons,
		})
	})

	// 查询单条数据   GET
	router.GET("/person/:id", func(c *gin.Context) {
		id := c.Param("id")
		var person Person
		err := db.QueryRow("SELECT id,first_name,last_name FROM person WHERE id=?", id).Scan(
			&person.Id, &person.FirstName, &person.LastName,
		)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"person": nil,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"person": person,
		})
	})

	// 更新数据   PUT
	router.PUT("/person/:id", func(c *gin.Context) {
		cid := c.Param("id")
		id, err := strconv.Atoi(cid)
		person := Person{Id: id}
		err = c.Bind(&person)
		if err != nil {
			log.Fatalln(err)
		}
		stmt, err := db.Prepare("UPDATE person SET first_name=?,last_name=? WHERE id=?")

		if err != nil {
			log.Fatalln(err)
		}
		defer stmt.Close()

		rs, err := stmt.Exec(person.FirstName, person.LastName, person.Id)

		if err != nil {
			log.Fatalln(err)
		}

		ra, err := rs.RowsAffected()
		if err != nil {
			log.Fatalln(err)
		}
		msg := fmt.Sprintf("Update person %d successful %d", person.Id, ra)

		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})

	})

	// 删除数据   DELETE
	router.DELETE("/person/:id", func(c *gin.Context) {
		cid := c.Param("id")
		id, err := strconv.Atoi(cid)
		if err != nil {
			log.Fatalln(err)
		}
		rs, err := db.Exec("DELETE FROM person WHERE id=?", id)
		if err != nil {
			log.Fatalln(err)
		}
		ra, err := rs.RowsAffected()
		if err != nil {
			log.Fatalln(err)
		}
		msg := fmt.Sprintf("Delete person %d successful %d", id, ra)
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
	})

	router.Run(":8000")
}
