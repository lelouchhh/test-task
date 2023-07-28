package postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"shop/domain"
)

type restAlgorithmRepository struct {
	baseURL string
}

func (r restAlgorithmRepository) GetAlgorithms(ctx context.Context) ([]domain.Algorithm, error) {
	client := &http.Client{}

	// Create a GET request
	url := r.baseURL + "/algorithm"
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logrus.Error(err)
		return []domain.Algorithm{}, domain.ErrInternalServerError
	}

	// Set any necessary headers (optional)

	// Send the request and get the response
	resp, err := client.Do(req)
	if err != nil {
		logrus.Error(err)

		return []domain.Algorithm{}, domain.ErrInternalServerError
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return []domain.Algorithm{}, domain.ErrInternalServerError
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error(err)

		return []domain.Algorithm{}, domain.ErrInternalServerError
	}

	// Unmarshal the JSON response into the Algorithm struct
	var algorithm []domain.Algorithm
	err = json.Unmarshal(body, &algorithm)
	if err != nil {
		logrus.Error(err)
		return []domain.Algorithm{}, domain.ErrInternalServerError

	}
	return algorithm, nil
}

func NewAlgorithmRepository(baseURL string) domain.AlgorithmRepository {
	return &restAlgorithmRepository{baseURL}
}
