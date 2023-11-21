package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/elue-dev/chi-api/helpers"
	"github.com/elue-dev/chi-api/initializers"
	"github.com/elue-dev/chi-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddPost (c *gin.Context) {
	var body models.Post
	c.Bind(&body)

	post := models.Post{
		Title:body.Title,
		Desc: body.Desc, 
		Category: body.Category,
	}

	validated := helpers.ValidatePostFields(body.Title, body.Desc, body.Category)

	if !validated  {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please fill all fields (Title, Description, Category)",
		})
		return
	}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"post": helpers.DatabasePostToPostModel(post),
	})
}

func GetPosts (c *gin.Context) {
	var posts []models.Post

    result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error fetching posts",
		})
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"posts": helpers.DatabasePostsArrToPostModel(posts),
	})
}

func GetPost (c *gin.Context) {
	var post models.Post
	id := c.Param("id")

	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Post with id " +  id + " could not be found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Something wennt wrong",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": helpers.DatabasePostToPostModel(post),
	})
}

func UpdatePost (c *gin.Context) {
	var body models.Post
	c.Bind(&body)

	var post models.Post
	id := c.Param("id")

	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf("Post with id %s could not be found", id),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Something wennt wrong",
			})
		}
		return
	}


	validated := helpers.ValidatePostFieldForUpdates(body.Title, body.Desc, body.Category)

	if !validated  {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please enter at least on field to update (Title, Description, Category)",
		})
		return
	}

	initializers.DB.Model(&post).Updates(models.Post{
		Title: helpers.UpdateFieldBasedOfValuePresence(body.Title, post.Title),
		Desc:  helpers.UpdateFieldBasedOfValuePresence(body.Desc, post.Desc),
		Category: helpers.UpdateFieldBasedOfValuePresence(body.Category, post.Category),
	})

	c.JSON(http.StatusOK, gin.H{
		"post": helpers.DatabasePostToPostModel(post),
	})
}

func DeletePost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")

	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf("Post with id %s could not be found", id),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Something wennt wrong",
			})
		}
		return
	}

	initializers.DB.Delete(&post, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Post has been deleted.",
	})
}