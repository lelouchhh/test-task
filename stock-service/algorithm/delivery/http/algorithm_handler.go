package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	domain "stock-service/domain"
)

// ArticleHandler  represent the httphandler for article
type AlgoHandler struct {
	AUsecase domain.AlgorithmUsecase
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewAlgoHandler(g *gin.Engine, us domain.AlgorithmUsecase) {
	handler := &AlgoHandler{
		AUsecase: us,
	}
	g.GET("/algorithm", handler.GetAlgorithm)
	g.POST("/solution", handler.GetSolution)
}

func (a *AlgoHandler) GetAlgorithm(c *gin.Context) {
	ctx := c.Request.Context()
	algorithms, err := a.AUsecase.GetAlgorithms(ctx)
	if err != nil {
		c.JSON(getStatusCode(err), map[string]string{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, algorithms)
}

func (a *AlgoHandler) GetSolution(c *gin.Context) {
	var id domain.Algorithm
	err := c.BindJSON(&id)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}
	ctx := c.Request.Context()
	solution, err := a.AUsecase.GetSolution(ctx, id.Id)
	if err != nil {
		c.JSON(getStatusCode(err), map[string]string{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, solution)
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
	default:
		return http.StatusInternalServerError
	}
}
