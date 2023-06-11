package body_temperature

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	ApiUrl = "https://jsonmock.hackerrank.com/api/medical_records"
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

/*
 * Complete the 'bodyTemperature' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. STRING doctorName
 *  2. INTEGER diagnosisId
 * API URL: https://jsonmock.hackerrank.com/api/medical_records?page={page_no}
 */

func GetResponse(doctorName string, diagnosisId int32, pageNo int) *Response {
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
	body, _ := io.ReadAll(resp.Body)
	var res Response
	_ = json.Unmarshal(body, &res)
	return &res
}

func bodyTemperature(doctorName string, diagnosisId int32) []int {
	r := GetResponse(doctorName, diagnosisId, 1)
	totalPages := r.TotalPages
	minTemp := math.MaxInt
	maxTemp := math.MinInt
	for page := 1; page <= totalPages; page++ {
		res := GetResponse(doctorName, diagnosisId, page)
		fmt.Println(page, ", res: ", res)
		for i := 0; i < len(res.Data); i++ {
			if res.Data[i].Diagnosis.ID != diagnosisId {
				continue
			}
			if res.Data[i].Doctor.Name != doctorName {
				continue
			}
			if minTemp > res.Data[i].Vitals.BodyTemperature {
				minTemp = res.Data[i].Vitals.BodyTemperature
			}
			if maxTemp < res.Data[i].Vitals.BodyTemperature {
				maxTemp = res.Data[i].Vitals.BodyTemperature
			}
		}
	}

	return []int{minTemp, maxTemp}
}

func RunBodyTemperature() {
	doctorName := "Dr Allysa Ellis"
	diagnosisId := 2
	res := bodyTemperature(doctorName, int32(diagnosisId))
	fmt.Println(res)
}
