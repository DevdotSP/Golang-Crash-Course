package diffuc

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// ReadDifferentFile demonstrates reading and writing various file formats.
func ReadDifferentFile() {
	// Uncomment functions to run specific file operations
	// ReadXML()        // Read data from an XML file
	// ToWriteXML()     // Write data to an XML file
	// ReadJSON()       // Read data from a JSON file
	// ToWriteJSON()    // Write data to a JSON file
	// ReadArigatest()  // Read custom JSON with nested structs
	// ToReadTEXTFILE() // Read a plain text file
	// ToWriteTEXTFILE()// Write data to a plain text file
	// ToReadCSVFILE()  // Read a CSV file
	ToWriteCSVFILE() // Write data to a CSV file
}

/* ================== XML PART ================== */

// Notes represents the XML structure for a note
type Notes struct {
	To      string `xml:"to"`
	From    string `xml:"from"`
	Heading string `xml:"heading"`
	Body    string `xml:"body"`
}

// ReadXML reads an XML file and unmarshals it into a Notes struct
func ReadXML() {
	data, _ := ioutil.ReadFile("Files/notes.xml")
	note := &Notes{}
	_ = xml.Unmarshal([]byte(data), &note)

	fmt.Println(note.To)
	fmt.Println(note.From)
	fmt.Println(note.Heading)
	fmt.Println(note.Body)
}

// ToWriteXML writes a Notes struct to an XML file with indentation
func ToWriteXML() {
	note := Notes{
		To:      "Nicky",
		From:    "Rock",
		Heading: "Meeting",
		Body:    "Meeting at 5pm!",
	}

	file, err := xml.MarshalIndent(note, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling XML: %v", err)
	}

	err = ioutil.WriteFile("Files/notes1.xml", file, 0644)
	if err != nil {
		log.Fatalf("Error writing XML file: %v", err)
	}

	log.Println("XML file 'notes1.xml' created successfully!")
}

/* ================== JSON PART ================== */

// CatlogNodes represents a list of catalog items in JSON
type CatlogNodes struct {
	CatlogNodes []Catlog `json:"catlog_nodes"`
}

// Catlog represents a single catalog item
type Catlog struct {
	ProductId string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

// ReadJSON reads a JSON file and prints the catalog data
func ReadJSON() {
	file, _ := ioutil.ReadFile("Files/test.json")
	data := CatlogNodes{}
	_ = json.Unmarshal([]byte(file), &data)

	for i := range data.CatlogNodes {
		fmt.Println("Product Id:", data.CatlogNodes[i].ProductId)
		fmt.Println("Quantity:", data.CatlogNodes[i].Quantity)
	}
}

// Salary represents a monthly salary structure
type Salary struct {
	Basic, HRA, TA float64
}

// JSONEmployee represents an employee with nested salary info
type JSONEmployee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []Salary
}

// ToWriteJSON writes a JSONEmployee struct to a JSON file
func ToWriteJSON() {
	data := JSONEmployee{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: []Salary{
			{Basic: 15000.00, HRA: 5000.00, TA: 2000.00},
			{Basic: 16000.00, HRA: 5000.00, TA: 2100.00},
			{Basic: 17000.00, HRA: 5000.00, TA: 2200.00},
		},
	}

	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile("Files/Arigatest.json", file, 0644)
}

// ReadArigatest reads a JSON file with nested salary info and prints details
func ReadArigatest() {
	file, _ := ioutil.ReadFile("Files/Arigatest.json")
	data := JSONEmployee{}
	_ = json.Unmarshal([]byte(file), &data)

	fmt.Println(data.FirstName)
	if len(data.FirstName) <= 4 {
		fmt.Printf("Mr.%v %v is %v years old\n", data.FirstName, data.LastName, data.Age)
	} else {
		fmt.Println("Name is longer than 4 characters")
	}
}

/* ================== TEXT FILE PART ================== */

// ToReadTEXTFILE reads a text file line by line into a string slice
func ToReadTEXTFILE() {
	file, err := os.Open("Files/testMoto.txt")
	if err != nil {
		log.Fatalf("Failed opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	// Print each line
	for _, line := range txtlines {
		fmt.Println(strings.Split(line, ","))
	}
}

// ToWriteTEXTFILE writes a slice of strings to a text file, each string in a new line
func ToWriteTEXTFILE() {
	sampledata := []string{
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		"Nunc a mi dapibus, faucibus mauris eu, fermentum ligula.",
		"Donec in mauris ut justo eleifend dapibus.",
		"Donec eu erat sit amet velit auctor tempus id eget mauris.",
	}

	file, err := os.OpenFile("Files/test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed creating file: %s", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range sampledata {
		_, _ = writer.WriteString(line + "\n")
	}
	writer.Flush()
}

/* ================== CSV FILE PART ================== */

// ToReadCSVFILE reads a CSV file line by line and prints each row
func ToReadCSVFILE() {
	file, err := os.Open("Files/csvtest.txt")
	if err != nil {
		log.Fatalf("Failed opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// ToWriteCSVFILE writes a 2D string slice to a CSV file
func ToWriteCSVFILE() {
	rows := [][]string{
		{"Name", "City", "Language"},
		{"Pinky", "London", "Python"},
		{"Nicky", "Paris", "Golang"},
		{"Micky", "Tokyo", "Php"},
	}

	csvfile, err := os.Create("Files/test.csv")
	if err != nil {
		log.Fatalf("Failed creating file: %s", err)
	}
	defer csvfile.Close()

	csvwriter := csv.NewWriter(csvfile)
	for _, row := range rows {
		_ = csvwriter.Write(row)
	}
	csvwriter.Flush()
}
