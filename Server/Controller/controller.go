package controller

import (
	"fmt"
	"log"
	database "server/Database"
	models "server/Models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetBlogList(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "ok",
		"message":    "Get Blog List",
	}

	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		context["statusText"] = "error"
		context["message"] = "unauthenticated"
		return c.Status(fiber.StatusUnauthorized).JSON(context)
	}

	claims := token.Claims.(*jwt.RegisteredClaims)
	var user models.User
	if err := database.DB.Where("id = ?", claims.Issuer).First(&user).Error; err != nil {
		context["statusText"] = "error"
		context["message"] = "User not found"
		return c.Status(fiber.StatusNotFound).JSON(context)
	}

	var bloglist []models.Blog
	result := database.DB.Where("user_id = ?", user.Id).Find(&bloglist)
	if result.Error != nil {
		log.Println("Error occurred while fetching blogs:", result.Error)
		context["statusText"] = "error"
		context["message"] = "Could not find blog data in the database"
		return c.Status(fiber.StatusInternalServerError).JSON(context)
	}

	context["data"] = bloglist
	context["message"] = "Request Successful"
	return c.Status(fiber.StatusOK).JSON(context)
}

func GetBlogDetail(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "ok",
		"message":    "Get Blog Details",
	}
	id := c.Params("id")
	var record models.Blog
	database.DB.First(&record, id)
	if record.ID == 0 {
		c.Status(404)
		context["statusText"] = ""
		context["message"] = "Blog not found"
	}
	context["data"] = record
	context["message"] = "Blog loaded successfully"
	c.Status(200)
	return c.JSON(context)
}
func CreateBlog(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "ok",
		"message":    "Create a Blog",
	}
	blog := new(models.Blog)
	if err := c.BodyParser(blog); err != nil {
		log.Println("Error occurred while parsing body:", err)
		context["statusText"] = "error"
		context["message"] = "Failed to parse request body"
		return c.Status(fiber.StatusBadRequest).JSON(context)
	}
	cookie := c.Cookies("jwt")
	fmt.Println(cookie)
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		context["statusText"] = "error"
		context["message"] = "unauthenticated"
		return c.Status(fiber.StatusUnauthorized).JSON(context)
	}
	claims := token.Claims.(*jwt.RegisteredClaims)
	blog.UserID, _ = strconv.Atoi(claims.Issuer)

	form, err := c.MultipartForm()
	if err != nil {
		log.Println("Error occurred while retrieving form data:", err)
		context["statusText"] = "error"
		context["message"] = "Failed to retrieve form data"
		return c.Status(fiber.StatusBadRequest).JSON(context)
	}
	files := form.File["image"]
	if len(files) > 0 {
		file := files[0]
		if file.Size > 0 {
			fileName := "./static/uploads/" + file.Filename
			if err := c.SaveFile(file, fileName); err != nil {
				log.Println("Error occurred while saving file:", err)
				context["statusText"] = "error"
				context["message"] = "Failed to save file"
				return c.Status(fiber.StatusInternalServerError).JSON(context)
			}
			blog.Image = "/static/uploads/" + file.Filename
		}
	} else {
		blog.Image = ""
	}
	result := database.DB.Create(blog)
	if result.Error != nil {
		log.Println("Error occurred while creating blog:", result.Error)
		context["statusText"] = "error"
		context["message"] = "Could not create blog"
		return c.Status(fiber.StatusInternalServerError).JSON(context)
	}
	context["data"] = blog
	context["message"] = "Blog Created Successfully"
	return c.Status(fiber.StatusCreated).JSON(context)
}

func UpdateBlog(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "ok",
		"message":    "Update a Blog",
	}
	id := c.Params("id")
	var record models.Blog
	database.DB.First(&record, id)
	if record.ID == 0 {
		c.Status(404)
		context["statusText"] = ""
		context["message"] = "Blog not found"
	}
	if err := c.BodyParser(&record); err != nil {
		context["statusText"] = ""
		context["message"] = "Something went wrong"
	}
	form, err := c.MultipartForm()
	if err != nil {
		log.Println("Error occurred while retrieving form data:", err)
		context["statusText"] = "error"
		context["message"] = "Failed to retrieve form data"
		return c.Status(fiber.StatusBadRequest).JSON(context)
	}
	files := form.File["image"]
	if len(files) > 0 {
		file := files[0]
		if file.Size > 0 {
			fileName := "./static/uploads/" + file.Filename
			if err := c.SaveFile(file, fileName); err != nil {
				log.Println("Error occurred while saving file:", err)
				context["statusText"] = "error"
				context["message"] = "Failed to save file"
				return c.Status(fiber.StatusInternalServerError).JSON(context)
			}
			record.Image = "/static/uploads/" + file.Filename
		}
	} else {
		record.Image = ""
	}
	context["data"] = record
	context["message"] = "Blog Updated Successfully"
	database.DB.Save(record)
	c.Status(200)
	return c.JSON(context)
}
func DeleteBlog(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "ok",
		"message":    "Delete a Blog",
	}
	id := c.Params("id")
	var record models.Blog
	result := database.DB.First(&record, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.Status(404)
			context["statusText"] = ""
			context["message"] = "Blog not found"
			return c.JSON(context)
		}
		log.Println("Error occurred while retrieving blog:", result.Error)
		c.Status(500)
		context["statusText"] = ""
		context["message"] = "Internal Server Error"
		return c.JSON(context)
	}
	result = database.DB.Delete(&record)
	if result.Error != nil {
		log.Println("Error occurred while deleting blog:", result.Error)
		c.Status(500)
		context["statusText"] = ""
		context["message"] = "Could not delete blog"
		return c.JSON(context)
	}
	context["message"] = "Blog deleted successfully"
	c.Status(200)
	return c.JSON(context)
}

const SecretKey = "secret"

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	pass, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: pass,
	}

	database.DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email=?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(user.Id),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Request Successful",
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	if token == nil || token.Claims == nil {
		return c.JSON(fiber.Map{
			"message": "token or claims is nil",
		})
	}
	claims := token.Claims.(*jwt.RegisteredClaims)
	var user models.User
	database.DB.Where("id=?", claims.Issuer).First(&user)
	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		HTTPOnly: true,
		Expires:  time.Now().Add(-time.Hour),
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
