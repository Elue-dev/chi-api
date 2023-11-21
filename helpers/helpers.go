package helpers

import "github.com/elue-dev/gin-api/models"

func ValidatePostFields(title, desc, category string) bool  {
	if title == "" || desc == "" || category == "" {
		return false
	} else {
		return true
	}
}

func ValidatePostFieldForUpdates(title, desc, category string) bool  {
	if title == "" && desc == "" && category == "" {
		return false
	} else {
		return true
	}
}

func DatabasePostToPostModel (dbPost models.Post) models.CustomPost {
	return models.CustomPost{
		ID:           dbPost.ID,
		CreatedAt:    dbPost.CreatedAt,
		UpdatedAt:    dbPost.UpdatedAt,
		DeletedAt:    dbPost.DeletedAt,
		Title:        dbPost.Title,
		Desc:  	      dbPost.Desc,
		Category:     dbPost.Category,
	}
}


func DatabasePostsArrToPostModel (dbPosts []models.Post) []models.CustomPost {
	posts := []models.CustomPost{}

	for _, dbPost := range dbPosts {
		posts = append(posts, DatabasePostToPostModel(dbPost))
	}
	return posts
}

func UpdateFieldBasedOfValuePresence(newVal, oldVal string) string {
	if newVal != "" {
		return newVal
	}
	return oldVal
}

