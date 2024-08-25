package status_provider

import (
	"encoding/json"
	"io"
	"kia-logix/internal/orders"
	"log"
	"net/http"
)

type Podro struct {
	statusURL string
}

var PodroStatusesMap = map[string]orders.OrderStatus{
	"در انبار":      orders.PROVIDER_SEEN,
	"در حال پردازش": orders.IN_PROGRESS,
	"تحویل شده":     orders.DELIVERED,
}

func (p *Podro) getOrderStatus(podroStatus string) orders.OrderStatus {
	return PodroStatusesMap[podroStatus]
}

var podroClient IStatusProvider

func NewPodro(statusURL string) IStatusProvider {
	if podroClient == nil {
		return &Podro{statusURL: statusURL}
	}
	return podroClient
}

func (p *Podro) GetOrderStatus(orderID uint) (orders.OrderStatus, error) {
	resp, err := http.Get(p.statusURL)
	if err != nil {
		log.Println(err)
		return "", err
	}
	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}
	// Unmarshal the JSON data into the Response struct
	var response PodroResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("Error unmarshaling JSON: %v", err)
	}

	// Variables to track the highest status and its corresponding fa_status
	var maxStatus string
	var maxFaStatus string

	// Iterate over the data array to find the highest status
	for _, entry := range response.Data {
		if entry.Status > maxStatus {
			maxStatus = entry.Status
			maxFaStatus = entry.FaStatus
		}
	}
	return p.getOrderStatus(maxFaStatus), nil
}

type PodroResponse struct {
	Message string `json:"message"`
	Data    []struct {
		Status    string `json:"status"`
		FaStatus  string `json:"fa_status"`
		CreatedAt string `json:"created_at"`
	} `json:"data"`
}
