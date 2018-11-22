package models

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	car := NewCar("21121", "111", "BMW110", 8000)
	rent := NewRent("21121", "111", 7, 64)
	rent.Car = car

	rules := rent.Rules()
	data, err := Transform(rent, rules)

	if data == nil || err != nil {
		t.Errorf("Failed")
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		t.Errorf("Failed")
	}
	fmt.Println(string(bytes))
}
