package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/koines/code-cadets-2021/lecture_4/03_bet_acceptance_api/internal/api/controllers/models"
	"net/http"
)

// Controller implements handlers for web server requests.
type Controller struct {
	betValidator BetValidator
	betService   BetService
}

// NewController creates a new instance of Controller
func NewController(betValidator BetValidator, betService BetService) *Controller {
	return &Controller{
		betValidator: betValidator,
		betService:   betService,
	}
}

// CreateBet handlers bet request.
func (e *Controller) CreateBet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var betRequestDto models.BetRequest
		err := ctx.ShouldBindWith(&betRequestDto, binding.JSON)
		if err != nil {
			ctx.String(http.StatusBadRequest, "bind: bet request is not valid.")
			return
		}

		if !e.betValidator.BetIsValid(betRequestDto) {
			ctx.String(http.StatusBadRequest, "validator: bet request is not valid.")
			return
		}

		err = e.betService.CreateBet(betRequestDto)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "request could not be processed.")
			return
		}

		ctx.Status(http.StatusOK)
	}
}
