package room

type requestUpdate struct {
	Name       string  `json:"name"`
	Active     bool    `json:"active"`
	CustomerID *uint64 `json:"customer_id"`
}

type response struct {
	Message string `json:"message"`
}
