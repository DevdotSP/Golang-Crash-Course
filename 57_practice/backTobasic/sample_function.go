package backtobasic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// sample is a custom int type used for demonstrating methods on basic types
type sample int

// Custom type definitions for clarity
type (
	Name     string
	Age      int
	NickName string
)

// SampleModel demonstrates a struct with JSON tags
type SampleModel struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	NickName string `json:"nick_name"`
}

// ComputationSample demonstrates a method on a custom type
func (s sample) ComputationSample(x int, y int) int {
	return x + y
}

// SampleModelFunc returns any type, demonstrating generics / interface usage
func (sm SampleModel) SampleModelFunc(T any) any {
	return T
}

// SampleMultiType returns interface{}, useful for multiple types
func (sm SampleModel) SampleMultiType(T any) interface{} {
	return T
}

// PaymentRequest represents the payload to send to a payment gateway
type PaymentRequest struct {
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	CardNumber  string  `json:"card_number"`
	ExpiryMonth string  `json:"expiry_month"`
	ExpiryYear  string  `json:"expiry_year"`
	CVV         string  `json:"cvv"`
}

// PaymentResponse represents the response from a payment gateway
type PaymentResponse struct {
	TransactionID string `json:"transaction_id"`
	Status        string `json:"status"`
	Message       string `json:"message"`
}

// SamplePayment demonstrates how to send a POST request to a payment API
func SamplePayment() {
	// Example payment request
	paymentRequest := PaymentRequest{
		Amount:      100.50,
		Currency:    "PHP",
		CardNumber:  "4111111111111111",
		ExpiryMonth: "12",
		ExpiryYear:  "2025",
		CVV:         "123",
	}

	// Convert struct to JSON
	requestBody, err := json.Marshal(paymentRequest)
	if err != nil {
		fmt.Printf("Error marshaling payment request: %v\n", err)
		return
	}

	// Send POST request to payment gateway
	url := "https://api.paymentgateway.com/v1/charge" // Example URL
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Decode the response JSON
	var paymentResponse PaymentResponse
	err = json.NewDecoder(resp.Body).Decode(&paymentResponse)
	if err != nil {
		fmt.Printf("Error decoding response: %v\n", err)
		return
	}

	// Handle the response
	if paymentResponse.Status == "success" {
		fmt.Printf("Payment successful! Transaction ID: %s\n", paymentResponse.TransactionID)
	} else {
		fmt.Printf("Payment failed: %s\n", paymentResponse.Message)
	}
}

// PaymongoSample demonstrates how to make an authorized GET request
func PaymongoSample() {
	url := "https://api.paymongo.com/v1/merchants/capabilities/payment_methods"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", "Basic c2tfdGVzdF9Gb21yeEQ0QmdKQzRCR0Y5QnY1YTN4QmQ6") // Example API key

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))
}

// User represents a user from the external API
type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	PhoneNo   string `json:"phone_no"`
	ID        string `json:"id"`
}

// Fetchusers retrieves users from an external API
func Fetchusers() ([]User, error) {
	url := "https://6756619c11ce847c992c9c36.mockapi.io/user/verify-member/users"

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response: %v", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var users []User
	if err := json.Unmarshal(body, &users); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %v", err)
	}

	return users, nil
}

// SampleToExtractAPI demonstrates iterating over fetched API users
func SampleToExtractAPI() {
	users, err := Fetchusers()
	if err != nil {
		fmt.Printf("Error fetching users: %v\n", err)
		return
	}

	for _, user := range users {
		fmt.Printf("User ID: %s\nFirstName: %s\nLastName: %s\nPhoneNo: %s\n\n",
			user.ID, user.FirstName, user.LastName, user.PhoneNo)
	}
}

// SamplePrint demonstrates iterating over SampleModel slice
func SamplePrint() {
	sampleModels := []SampleModel{
		{Name: "John Doe", Age: 30, NickName: "Johnny"},
		{Name: "Jane Smith", Age: 25, NickName: "Janey"},
	}

	for _, m := range sampleModels {
		fmt.Printf("%v is %v years old with the nickname %v\n",
			m.Name, m.Age, m.NickName)
	}

	// Access first element and demonstrate type conversion
	sm := sampleModels[0]
	fmt.Printf("Converted types: %v, %v, %v\n",
		Name(sm.Name), Age(sm.Age), NickName(sm.NickName))
}

// SampleFunction demonstrates multiple examples
func SampleFunction() {
	// Uncomment to run examples
	// FirstExample()
	// SecondExample()
	// ThirdExample()
	// SamplePayment()
	PaymongoSample()
}

func ThirdExample() {
	SampleToExtractAPI()
}

func SecondExample() {
	SamplePrint()
}

func FirstExample() {
	// Demonstrate ComputationSample method on sample type
	var numSample int
	sampleInstance := sample(numSample)
	result := sampleInstance.ComputationSample(5, 6)
	fmt.Printf("The result is %v\n", result)

	// Demonstrate SampleModel usage
	sampleModel := SampleModel{Name: "John Doe", Age: 30, NickName: "Johnny"}
	funcResult := sampleModel.SampleMultiType(sampleModel)

	if model, ok := funcResult.(SampleModel); ok {
		fmt.Printf("%v is %v years old with nickname %v\n", model.Name, model.Age, model.NickName)
	} else {
		fmt.Println("Type assertion failed")
	}
}
