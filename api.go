package main

import (
	"example/go-orm-api/model"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "warawit:admin@tcp(127.0.0.1:3306)/app_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	r := gin.Default()
	r.GET("/users", func(ctx *gin.Context) {
		var users []model.User
		db.Find(&users)
		ctx.JSON(200, users)
	})
	r.GET("/user/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		var user model.User
		db.First(&user, id)
		ctx.JSON(200, user)
	})
	r.POST("/users", func(ctx *gin.Context) {
		var user model.User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result := db.Create(&user)
		ctx.JSON(200, gin.H{"RowAffected": result.RowsAffected})
	})
	r.DELETE("/user/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		var user model.User
		db.First(&user, id)
		db.Delete(&user)
		ctx.JSON(200, user)
	})
	r.PUT("/users", func(ctx *gin.Context) {
		var user model.User
		var updatedUser model.User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.First(&updatedUser, user.ID)
		updatedUser.Fname = user.Fname
		updatedUser.Lname = user.Lname
		updatedUser.Username = user.Username
		updatedUser.Avatar = user.Avatar
		db.Save(updatedUser)
		ctx.JSON(200, user)
	})
	r.Use(cors.Default())
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
