package http

import (
	"net/http"
	"strconv"

	"rasyidcode/simpleapi-clean-arch/models"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type ResponseError struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type TodoHandler struct {
	TUsecase models.TodoUsecase
}

func NewTodoHandler(e *echo.Echo, us models.TodoUsecase) {
	handler := &TodoHandler{
		TUsecase: us,
	}

	e.GET("/todo/:id", handler.GetByID)
}

func (a *TodoHandler) GetByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, models.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()

	todo, err := a.TUsecase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error(), Success: false})
	}

	return c.JSON(http.StatusOK, todo)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case models.ErrInternalServerError:
		return http.StatusInternalServerError
	case models.ErrNotFound:
		return http.StatusNotFound
	case models.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
