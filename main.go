package main

// import needed packages
import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"
	"log"
	//"fmt"
	"bufio"
)

// define struct to hold data from the csv that will become JSON
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
		// we dont start at i = 0 because that is the column name row
        if i > 0 { 
			//define empty house struct variable to hold data
            var rec house
            for j, field := range line {
                if j == 0 {
                    rec.Value, _ = strconv.Atoi(field)
                } else if j == 1 {
					// income is parsed as a string because of the period in this field
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
    // open file from the first user argument
    f, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    // remember to close the file at the end of the program
    defer f.Close()

    // Read CSV file using csv.Reader
    csvReader := csv.NewReader(f)
    data, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal(err)
    }

	// use the previously defined function to turn the csv data into a list of house structs
    house_data := CreateHousingData(data)

    // convert an array of structs to JSON using marshaling functions from the encoding/json package
    jsonData, err := json.MarshalIndent(house_data, "", "  ")
    if err != nil {
        log.Fatal(err)
    }

	// Open the JSONL file that will be written to
	file, err := os.Create(os.Args[2])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a buffered writer for efficient writing
	writer := bufio.NewWriter(file)


	var jsonArray []map[string]interface{}
	err = json.Unmarshal([]byte(jsonData), &jsonArray)
	if err != nil {
		panic(err)
	}

	for _, element := range jsonArray {
		data, err := json.Marshal(element)
		if err != nil {
			panic(err)
		}

		// Append a newline after each JSON object
		data = append(data, '\n')

		// write each element in json array to jsonl file
		if _, err := writer.Write(data); err != nil {
			panic(err)
		}
	}
	writer.Flush()

}








