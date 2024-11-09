package works_oa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"
)

const (
	ApiUrl = "https://jsonmock.hackerrank.com/api/food_outlets"
)

type Rating struct {
	AverageRating float32 `json:"average_rating"`
	Votes         int32   `json:"votes"`
}

type Request struct {
}

type Record struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	City          string `json:"city"`
	EstimatedCost int    `json:"estimated_cost"`
	UserRating    Rating `json:"user_rating"`
}

type Response struct {
	Page       int      `json:"page"`
	PerPage    int      `json:"per_page"`
	Total      int      `json:"total"`
	TotalPages int      `json:"total_pages"`
	Data       []Record `json:"data"`
}

func GetResponse(city string, pageNo int) *Response {
	apiClient := http.Client{Timeout: time.Minute * 2}
	reqBody := &Request{}
	searchUrl := ApiUrl

	u, _ := url.Parse(searchUrl)
	q := u.Query()
	q.Set("city", city)
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

func solve(city string, minVotes int32) string {
	r := GetResponse(city, 1)
	totalPages := r.TotalPages

	outlets := make([]Record, 0)
	for page := 1; page <= totalPages; page++ {
		res := GetResponse(city, page)
		for i := 0; i < len(res.Data); i++ {
			// do something
			if res.Data[i].UserRating.Votes >= minVotes {
				outlets = append(outlets, res.Data[i])
			}
		}
	}

	sort.SliceStable(outlets, func(i, j int) bool {
		if outlets[i].UserRating.AverageRating == outlets[j].UserRating.AverageRating {
			return outlets[i].UserRating.Votes > outlets[j].UserRating.Votes
		} else {
			return outlets[i].UserRating.AverageRating > outlets[j].UserRating.AverageRating
		}
	})

	return outlets[0].Name
}

func finestFoodOutlet(city string, votes int32) string {
	return solve(city, votes)
}

func Run() {
	city := "Seattle"
	votes := 500
	fmt.Println(finestFoodOutlet(city, int32(votes)))
}
