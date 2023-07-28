package postgres

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"shop/domain"
)

type restPurchaseRepository struct {
	baseURL string
}

func (r restPurchaseRepository) BuyPurchase(ctx context.Context, p domain.Purchase) (domain.Algorithm, error) {
	url := r.baseURL + "/solution"

	// Create a map representing the JSON payload with the 'id' parameter
	data := map[string]interface{}{
		"id": p.UserId, // Replace 123 with the desired 'id' value
	}

	// Marshal the data map into a JSON byte slice
	body, err := json.Marshal(data)
	if err != nil {
		return domain.Algorithm{}, domain.ErrInternalServerError
	}

	// Create the POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return domain.Algorithm{}, domain.ErrInternalServerError
	}

	// Set the Content-Type header to indicate JSON data
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return domain.Algorithm{}, domain.ErrInternalServerError
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return domain.Algorithm{}, domain.ErrInternalServerError
	}

	// Read the response body
	var algorithm domain.Algorithm
	err = json.NewDecoder(resp.Body).Decode(&algorithm)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return domain.Algorithm{}, nil
	}

	// Process the response data as needed
	return algorithm, nil
}

func NewPurchaseRepository(baseURL string) domain.PurchaseRepository {
	return &restPurchaseRepository{baseURL}
}
