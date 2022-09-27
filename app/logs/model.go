package logs

type requestUpdate struct {
	CustomerID *uint64 `json:"customer_id"`
	RoomID     *uint64 `json:"room_id"`
}

type response struct {
	Message string `json:"message"`
}
