package works_template

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	ApiUrl = "https://jsonmock.hackerrank.com/api/medical_records"
)

type Request struct {
}

type Record struct {
}

type Response struct {
	Page       int      `json:"page"`
	PerPage    int      `json:"per_page"`
	Total      int      `json:"total"`
	TotalPages int      `json:"total_pages"`
	Data       []Record `json:"data"`
}

func GetResponse(pageNo int) *Response {
	searchUrl := "https://jsonmock.hackerrank.com/api/medical_records"
	apiClient := http.Client{Timeout: time.Minute * 2}
	reqBody := &Request{}

	u, _ := url.Parse(searchUrl)
	q := u.Query()
	q.Set("page", strconv.Itoa(pageNo))
	u.RawQuery = q.Encode()

	var buf bytes.Buffer
	_ = json.NewEncoder(&buf).Encode(reqBody)
	req, _ := http.NewRequest(http.MethodGet, u.String(), &buf)

	resp, _ := apiClient.Do(req)
	body, _ := io.ReadAll(resp.Body)
	var res Response
	_ = json.Unmarshal(body, &res)
	return &res
}

func solve() {
	r := GetResponse(1)
	totalPages := r.TotalPages
	for page := 1; page <= totalPages; page++ {
		res := GetResponse(page)
		for i := 0; i < len(res.Data); i++ {
			// do something
		}
	}
}
