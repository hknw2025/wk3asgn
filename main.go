package main

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"
	"log"
)

type house struct {
	Value int `json:"Value"`
	Income string `json:"Income"`
	Age int `json:"Age"`
	Rooms int `json:"Rooms"`
	Bedrooms int `json:"Bedrooms"`
	Pop int `json:"Pop"`
	HH int `json:"HH"`

}

func CreateHousingData(data [][]string) []house {
    var house_data []house
    for i, line := range data {
        if i > 0 { // omit header line
            var rec house
            for j, field := range line {
                if j == 0 {
                    rec.Value, _ = strconv.Atoi(field)
                } else if j == 1 {
                    rec.Income = field
                } else if j == 2 {
                    rec.Age, _ = strconv.Atoi(field)
                } else if j == 3 {
                    rec.Rooms, _ = strconv.Atoi(field)
                } else if j == 4 {
                    rec.Bedrooms, _ = strconv.Atoi(field)
                } else if j == 5 {
                    rec.Pop, _ = strconv.Atoi(field)
                } else if j == 6 {
                    rec.HH, _ = strconv.Atoi(field)
                }
            }
            house_data = append(house_data, rec)
        }
    }
    return house_data
}

func main() {


    // open file
    f, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    // remember to close the file at the end of the program
    defer f.Close()

    // 2. Read CSV file using csv.Reader
    csvReader := csv.NewReader(f)
    data, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal(err)
    }

    // 3. Assign successive lines of raw CSV data to fields of the created structs
    house_data := CreateHousingData(data)

    // 4. Convert an array of structs to JSON using marshaling functions from the encoding/json package
    jsonData, err := json.MarshalIndent(house_data, "", "  ")
    if err != nil {
        log.Fatal(err)
    }

	os.WriteFile(os.Args[2], jsonData, os.ModePerm)


}








