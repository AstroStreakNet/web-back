package middlewares

import (
	"fmt"
	"log/slog"
	"os"
)

// Retrieves the URL for serving static images
func GetStaticURL(imageID string) string {
	var imagePath = fmt.Sprintf("/static/%s.webp", imageID)
	_, err := os.Stat(imagePath)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to retrieve static path for image: %s", err))
		// Return placeholder image if file can't be accessed
		return "/static/billy.jpeg"
	}

	return imagePath
}
