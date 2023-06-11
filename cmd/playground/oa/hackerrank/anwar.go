package hackerrank

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

/*
 * Complete the 'EncodeManager' function below and the struct Manager.
 *
 * The function is expected to return an io.Reader and an error.
 * The function accepts *Manager manager as parameter.
 */

type Manager struct {
	FullName       string
	Position       string
	Age            int32
	YearsInCompany int32
}

func EncodeManager(manager *Manager) (io.Reader, error) {
	m, err := json.Marshal(map[string]interface{}{
		"full_name":        manager.FullName,
		"position":         manager.Position,
		"age":              manager.Age,
		"years_in_company": manager.YearsInCompany,
	})
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(m), nil

}

func Run() {
	m := Manager{
		FullName:       "Jack Oliver",
		Position:       "CEO",
		Age:            44,
		YearsInCompany: 15,
	}

	fmt.Println("manager")

	resultReader, err := EncodeManager(&m)
	if err != nil {
		log.Fatal("error getting the reader", err.Error())
	}

	result, err := ioutil.ReadAll(resultReader)
	if err != nil {
		log.Fatal("error reading from resultReader", err.Error())
	}

	_ = fmt.Sprintf("%s\n", string(result))
}
