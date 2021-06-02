package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Controller implements handlers for web server requests.
type Controller struct {
	idValidator     BetDtoIdValidator
	statusValidator BetDtoStatusValidator
	betDtoService   BetDtoService
}

// NewController creates a new instance of Controller
func NewController(idValidator BetDtoIdValidator, statusValidator BetDtoStatusValidator, betDtoService BetDtoService) *Controller {
	return &Controller{
		idValidator:     idValidator,
		statusValidator: statusValidator,
		betDtoService:   betDtoService,
	}
}

func (e *Controller) FetchAll() gin.HandlerFunc {
	return func(context *gin.Context) {
		userId := context.Param("id")

		bets, err := e.betDtoService.GetAll(context, userId)
		if err != nil {
			context.String(http.StatusInternalServerError, "request could not be processed.")
			return
		}

		context.JSON(http.StatusOK, bets)
	}
}

func (e *Controller) FetchSpecificId() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")

		bet, err := e.betDtoService.GetByID(context, id)
		if err != nil {
			//context.String(http.StatusInternalServerError, "request could not be processed.")
			context.String(http.StatusInternalServerError, err.Error())
			return
		}

		context.JSON(http.StatusOK, bet)
	}
}

func (e *Controller) FetchSpecificStatus() gin.HandlerFunc {
	return func(context *gin.Context) {
		status := context.Query("status")

		bets, err := e.betDtoService.GetByStatus(context, status)
		if err != nil {
			context.String(http.StatusInternalServerError, err.Error())
			//context.String(http.StatusInternalServerError, "request could not be processed.")
			return
		}

		context.JSON(http.StatusOK, bets)
	}

}
