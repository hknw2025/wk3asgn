package main

import (
	"testing"
)





func TestCreateHousingData(t *testing.T) {

	testdata := [][]string{
        {"value", "income", "age", "rooms", "bedrooms", "pop", "hh"},
        {"100", "10.5", "5", "2", "5", "10", "1"},
    }

	data := CreateHousingData(testdata)

	var ans_list []house

	ans := house{
		Value: 100,
		Income: "10.5",
		Age:  5,
		Rooms: 2,
		Bedrooms: 5,
		Pop: 10, 
		HH: 1,
	}

	ans_list = append(ans_list, ans)

	if data[0] != ans_list[0] {
        t.Errorf("parsed incorrectly")
    }


}