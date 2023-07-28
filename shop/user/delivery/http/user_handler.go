package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"net/http"
	domain "shop/domain"
)

// ArticleHandler  represent the httphandler for article
type userHandler struct {
	AUsecase domain.UserUsecase
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewUserHandler(g *gin.Engine, us domain.UserUsecase) {
	handler := &userHandler{
		AUsecase: us,
	}
	g.POST("/user", handler.RegisterUser)
	g.PUT("/increase_debt", handler.IncreaseDebt)
	g.PUT("/decrease_debt", handler.DecreaseDebt)
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var user domain.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": domain.ErrBadParamInput.Error(),
		})
		return
	}
	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{
				"error": domain.ErrBadParamInput.Error(),
			})
			return
		}
	}

	ctx := c.Request.Context()
	err = h.AUsecase.CreateUser(ctx, user)
	if err != nil {
		c.JSON(getStatusCode(err), map[string]string{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func (h *userHandler) IncreaseDebt(c *gin.Context) {
	var debt domain.Debt
	err := c.BindJSON(&debt)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": domain.ErrBadParamInput.Error(),
		})
		return
	}
	validate := validator.New()
	err = validate.Struct(debt)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}
	ctx := c.Request.Context()
	err = h.AUsecase.IncreaseDebt(ctx, debt)
	if err != nil {
		c.JSON(getStatusCode(err), map[string]string{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func (h *userHandler) DecreaseDebt(c *gin.Context) {
	var debt domain.Debt
	err := c.BindJSON(&debt)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": domain.ErrBadParamInput.Error(),
		})
		return
	}
	validate := validator.New()
	err = validate.Struct(debt)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}
	ctx := c.Request.Context()
	err = h.AUsecase.DecreaseDebt(ctx, debt)
	if err != nil {
		c.JSON(getStatusCode(err), map[string]string{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, "ok")
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
