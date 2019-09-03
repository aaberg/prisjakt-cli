package main

type ProductItem struct {
	Name  string `json:"name"`
	Id    int64  `json:"id"`
	Url   string `json:"url"`
	Price struct {
		Regular     float32
		Alternative float32
		Display     string
		InStock     string
	}
}

type Product struct {
	Items             []ProductItem `json:"items"`
	MoreHitsAvailable bool          `json:"more_hits_available"`
}

type Message struct {
	Product Product `json:"product"`
}

type Response struct {
	Error     bool
	Message   Message `json:"message"`
	RequestId int     `json:"request_id"`
	Timestamp float64
}
