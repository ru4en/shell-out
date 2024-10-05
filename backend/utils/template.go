package shellout

import "github.com/gin-gonic/gin"

// Template struct
type Template struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DockerFile  string `json:"dockerfile"`
}

func GetTemplateList(c *gin.Context) {
	// Get test data
	templates := []Template{
		{
			ID:          "1",
			Title:       "ubuntu",
			Description: "Ubuntu template",
			DockerFile:  "\nFROM ubuntu:latest\n\nRUN apt-get update\nRUN apt-get install -y nginx\n\nCMD [\"nginx\", \"-g\", \"daemon off;\"]\n",
		},
		{
			ID:          "2",
			Title:       "debian",
			Description: "Debian template",
			DockerFile:  "\nFROM debian:latest\n\nRUN apt-get update\nRUN apt-get install -y apache2\n\nCMD [\"apache2ctl\", \"-D\", \"FOREGROUND\"]\n",
		},
	}
	c.JSON(200, templates)
}

func GetTemplate(id string) {
	// Get a template
}

func CreateTemplate() {
	// Create a template
}

func UpdateTemplate(id string) {
	// Update a template
}

func DeleteTemplate(id string) {
	// Delete a template
}
