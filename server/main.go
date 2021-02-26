package main

import (
	"crowdfunding/user"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	UserInput := user.RegisterUserInput{}
	UserInput.Name = "tes simpan dari service"
	UserInput.Email = "tes@gmail.com"
	UserInput.Occupation = "google"
	UserInput.Password = "ini password"

	userService.RegisterUser(UserInput)

	// fmt.Println("Connected to Database")

	// var users []user.User
	// db.Find(&users)

	// fmt.Println(len(users))

	// for _, user := range users {
	// 	fmt.Println(user.Name)
	// 	fmt.Println(user.Email)
	// 	fmt.Println(user.Occupation)
	// }

	// router := gin.Default()
	// router.GET("/handler", handler)
	// router.Run()

}

// func handler(c *gin.Context) {
// 	dsn := "root:@tcp(127.0.0.1:3306)/crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	var users []user.User
// 	db.Find(&users)

// 	c.JSON(http.StatusOK, users)

//handler
//service
//repository
//db
// }
