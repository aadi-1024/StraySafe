package handlers

import (
	"github.com/aadi-1024/StraySafe/backend/pkg/database"
	"github.com/aadi-1024/StraySafe/backend/pkg/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func IncidentPostHandler(d *database.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Get("typ").(string) != "usr" {
			return c.JSON(http.StatusUnauthorized, models.JsonResponse{
				Message: "user expected",
				Content: nil,
			})
		}

		latitude := c.FormValue("latitude")
		lat, _ := strconv.ParseFloat(latitude, 32)

		longitude := c.FormValue("longitude")
		lon, _ := strconv.ParseFloat(longitude, 32)

		image, _ := c.FormFile("image")
		buf := make([]byte, image.Size)

		file, err := image.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.JsonResponse{
				Message: "error parsing image",
				Content: err.Error(),
			})
		}

		_, err = file.Read(buf)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.JsonResponse{
				Message: "error parsing image",
				Content: err.Error(),
			})
		}

		inc := models.Incident{
			Uid:         c.Get("id").(int),
			Latitude:    float32(lat),
			Longitude:   float32(lon),
			Title:       c.FormValue("title"),
			Description: c.FormValue("description"),
			Image:       buf,
		}

		err = d.NewIncident(inc)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.JsonResponse{
				Message: "database error",
				Content: err.Error(),
			})
		}

		return c.JSON(http.StatusOK, models.JsonResponse{
			Message: "successful",
			Content: nil,
		})
	}
}