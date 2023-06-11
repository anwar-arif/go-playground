package works_oa

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"
)

var api_url = "https://jsonmock.hackerrank.com/api/countries/search?region={region}&name={keyword}&page={page_no}"

type Country struct {
	Name       string `json:"name"`
	Population int    `json:"population"`
	//Capital    string   `json:"capital"`
	//Region     string   `json:"region"`
	//Currencies []string `json:"currencies"`
	//Borders    []string `json:"borders"`
}

type Response struct {
	Page       int       `json:"page"`
	PerPage    int       `json:"per_page"`
	Total      int       `json:"total"`
	TotalPages int       `json:"total_pages"`
	Data       []Country `json:"data"`
}

func GetResponse(region string, keyword string, page int) *Response {
	searchUrl := "https://jsonmock.hackerrank.com/api/countries/search"
	apiClient := http.Client{Timeout: time.Minute * 2}

	u, _ := url.Parse(searchUrl)
	q := u.Query()
	q.Set("region", region)
	q.Set("name", keyword)
	q.Set("page", strconv.Itoa(page))
	u.RawQuery = q.Encode()

	req, _ := http.NewRequest(http.MethodGet, u.String(), nil)
	fmt.Println(req.URL.String())

	resp, _ := apiClient.Do(req)
	body, parseErr := io.ReadAll(resp.Body)
	if parseErr != nil {
		log.Fatal("error reading the body")
		return nil
	}

	var res Response
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("error while parsing response body")
		return nil
	}

	return &res
}

func FindCountry(region string, keyword string) {
	res := GetResponse(region, keyword, 1)
	if res == nil {
		return
	}
	countries := make([]Country, 0)
	for _, val := range res.Data {
		countries = append(countries, val)
	}
	totalPage := res.TotalPages
	for page := 2; page <= totalPage; page++ {
		r := GetResponse(region, keyword, page)
		if r == nil {
			continue
		}
		for _, c := range r.Data {
			countries = append(countries, c)
		}
	}
	fmt.Println(len(countries))
	sort.SliceStable(countries, func(i, j int) bool {
		if countries[i].Population == countries[j].Population {
			return countries[i].Name < countries[j].Name
		}
		return countries[i].Population < countries[j].Population
	})

	for _, country := range countries {
		fmt.Println(country.Name, country.Population)
	}
}
