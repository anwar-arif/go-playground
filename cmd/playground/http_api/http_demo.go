package http_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Diagnosis struct {
	ID int32 `json:"id"`
	//Name     string `json:"name"`
	//Severity int    `json:"severity"`
}

type Vitals struct {
	//BloodPressureDiastole int `json:"bloodPressureDiastole"`
	//BloodPressureSystole  int `json:"bloodPressureSystole"`
	//Pulse                 int `json:"pulse"`
	//BreathingRate         int `json:"breathingRate"`
	BodyTemperature int `json:"bodyTemperature"`
}

type Doctor struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Meta struct {
	Height int `json:"height"`
	Weight int `json:"weight"`
}

type Record struct {
	//ID int `json:"id"`
	//Timestamp int       `json:"timestamp"`
	Diagnosis Diagnosis `json:"diagnosis"`
	Vitals    Vitals    `json:"vitals"`
	Doctor    Doctor    `json:"doctor"`
	//UserId    int       `json:"userId"`
	//UserName  string    `json:"userName"`
	//UserDOB   string    `json:"userDob"`
	//Meta      Meta      `json:"meta"`
}

type Request struct {
	DoctorName  string `json:"doctorName"`
	DiagnosisId int32  `json:"diagnosisId"`
}

type Response struct {
	Page       int      `json:"page"`
	PerPage    int      `json:"per_page"`
	Total      int      `json:"total"`
	TotalPages int      `json:"total_pages"`
	Data       []Record `json:"data"`
}

func GetResponse(doctorName string, diagnosisId int32, pageNo int) Response {
	searchUrl := "https://jsonmock.hackerrank.com/api/medical_records"
	apiClient := http.Client{Timeout: time.Minute * 2}
	reqBody := &Request{
		DoctorName:  doctorName,
		DiagnosisId: diagnosisId,
	}

	u, _ := url.Parse(searchUrl)
	q := u.Query()
	q.Set("page", strconv.Itoa(pageNo))
	u.RawQuery = q.Encode()

	var buf bytes.Buffer
	_ = json.NewEncoder(&buf).Encode(reqBody)
	req, _ := http.NewRequest(http.MethodGet, u.String(), &buf)

	resp, _ := apiClient.Do(req)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var res Response
	_ = json.Unmarshal(body, &res)
	return res
}

func RunHttpDemo() {
	res := GetResponse("Dr Arnold Bullock", 1, 1)
	fmt.Println(res)
}
