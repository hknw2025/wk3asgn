# Week 3 Asignment - CSV to JSON Command Line Tool

This tool takes the given data in housesInput.csv and output properly ordered JSON data. 
It finds the input data at a user given path and outputs the data to a file at the user given path.

## How to use this tool:
* After cloning this git repository you can 
`go run main.go {path to input housing CSV} {path to output jsonl file}`
* you will have a new file in the output location with JSONL extension.

## Example:

  {"value", "income", "age", "rooms", "bedrooms", "pop", "hh"},
        {"100", "10.5", "5", "2", "5", "10", "1"},

### Basic Transformation
If I have a csv file contains a test data `test.csv`:
```csv
Value,Income,Age,Rooms,Bedrooms,Pop,HH
100,10.5,5,2,5,10,1
```
and we need to convert it to json file:

`go run main.go ./data/test.csv ./data/test.jsonl`

after writing this command you will get another file in the `data` directory called `test.jsonl` with the previously CSV data in JSONL format:
```json
[
  {
    "Value":100,
    "Income":"10.5",
    "Age":5,
    "Rooms":2,
    "Bedrooms":5,
    "Pop":10,
    "HH":1
  }
]
```







