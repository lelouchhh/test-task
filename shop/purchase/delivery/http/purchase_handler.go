package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"net/http"
	"shop/domain"
)

type PurchaseHandler struct {
	PUsecase domain.PurchaseUsecase
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewPurchaseHandler(g *gin.Engine, us domain.PurchaseUsecase) {
	handler := &PurchaseHandler{
		PUsecase: us,
	}
	g.POST("/purchase", handler.Purchase)
}

func (a *PurchaseHandler) Purchase(c *gin.Context) {
	var solution domain.Purchase
	err := c.BindJSON(&solution)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": domain.ErrBadParamInput.Error(),
		})
		return
	}
	validate := validator.New()
	err = validate.Struct(solution)
	if err != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{
				"error": domain.ErrBadParamInput.Error(),
			})
			return
		}
	}
	ctx := c.Request.Context()
	algo, err := a.PUsecase.BuyPurchase(ctx, solution)
	if err != nil {
		c.JSON(getStatusCode(err), map[string]string{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, algo)
}
func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	case domain.UserAlreadyExist:
		return http.StatusConflict
	case domain.ErrUnprocessableEntity:
		return http.StatusUnprocessableEntity
	default:
		return http.StatusInternalServerError
	}
}
