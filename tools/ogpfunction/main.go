package main

import (
	"net/http"
	"os"

	"github.com/dyatlov/go-opengraph/opengraph"
	"github.com/dyatlov/go-opengraph/opengraph/types/image"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ResponsePayload struct {
	URL         string  `json:"url"`
	SiteName    string  `json:"site_name"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ImageURL    *string `json:"image_url,omitempty"`
}

func getenv(key, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/api/ogp", getOGPHandler)

	port := getenv("PORT", getenv("FUNCTIONS_CUSTOMHANDLER_PORT", "9000"))
	e.Logger.Fatal(e.Start(":" + port))
}

func getOGPHandler(c echo.Context) error {
	url := c.QueryParam("url")
	empty := ResponsePayload{
		URL:         url,
		SiteName:    "",
		Title:       url,
		Description: "",
	}
	if url == "" {
		return c.JSON(http.StatusOK, empty)
	}

	resp, err := http.Get(url)
	if err != nil {
		c.Logger().Errorf("Could not get url[%s]. %w", url, err)
		return c.JSON(http.StatusOK, empty)
	}
	defer resp.Body.Close()

	og := opengraph.NewOpenGraph()
	err = og.ProcessHTML(resp.Body)
	if err != nil {
		c.Logger().Errorf("Could not process ogp for url[%s]. %w", url, err)
		return c.JSON(http.StatusOK, empty)
	}

	var imageURL *string
	image := getTopImage(og)
	if image != nil {
		imageURL = &image.URL
	}
	title := og.Title
	if title == "" {
		title = url
	}

	payload := ResponsePayload{
		URL:         url,
		SiteName:    og.SiteName,
		Title:       og.Title,
		Description: og.Description,
		ImageURL:    imageURL,
	}

	c.Response().Header().Set(echo.HeaderCacheControl, "public, max-age=86400")
	return c.JSON(http.StatusOK, payload)
}

func getTopImage(og *opengraph.OpenGraph) *image.Image {
	if len(og.Images) > 0 {
		return og.Images[0]
	}
	return nil
}
