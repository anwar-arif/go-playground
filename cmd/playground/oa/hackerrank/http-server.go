package hackerrank

type Lake struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Area int64  `json:"area"`
}

type CreateRequest struct {
	Type    string `json:"type"`
	Payload Lake   `json:"payload"`
}

type GenericRequest struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
}
