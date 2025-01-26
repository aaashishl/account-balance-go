package accountbalance

import (
	"net/http"
	"errors"
	"io/ioutil"
	"encoding/json"
)
type AccountBalance struct {
	Balance string `json:"balance"`
	LockedBalance string `json:"lockedBalance"`
}
func GetAccountBalance(baseURL, address string) (AccountBalance, error){
	if address == "" {
		return AccountBalance{}, errors.New("Address is required")
	}
	var ab AccountBalance
	
	response, err := http.Get(baseURL + "/addresses/" + address + "/balance")

	if err != nil {
		return AccountBalance{}, errors.New("Error while fetching account balance")
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err := json.Unmarshal(body, &ab); err != nil {
		return AccountBalance{}, errors.New("Error while unmarshalling response")
	}
	return ab, nil
}