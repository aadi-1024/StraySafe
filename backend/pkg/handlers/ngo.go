package handlers

import (
	"github.com/aadi-1024/StraySafe/backend/pkg/database"
	"github.com/aadi-1024/StraySafe/backend/pkg/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func MarkResolvedHandler(db *database.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Get("typ") != "ngo" {
			return c.JSON(http.StatusBadRequest, models.JsonResponse{
				Message: "expected ngo",
				Content: nil,
			})
		}
		incId, _ := strconv.Atoi(c.FormValue("id"))
		inc := models.Incident{Id: incId}
		err := db.MarkResolved(inc, c.Get("id").(int))

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
