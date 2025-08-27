package diffuc

import (
	"fmt"
	"regexp"
	"strings"
)

/*

Regular Expression (Regex) - Examples and Solutions
*/

func ExampleRegex() {
	// Uncomment individual functions to run examples
	// RegularRegexInBracket()   // Extract text inside square brackets
	// RegexNonAlphaChar()       // Extract non-alphanumeric characters
	// RegexExtractDate()        // Extract dates in YYYY-MM-DD format
	// RegexDNS()                // Extract IP addresses from text
	// RegexURL()                // Extract domain from URL
	// RegexEmailAdd()           // Validate email addresses
	// RegexPhonenumber()        // Validate phone numbers
	RegexDate()                 // Validate dates in DD/MM/YYYY format
}

// -------------------------------
// Extract text between square brackets
func RegularRegexInBracket() {
	str := "this is a [sample] [[string]] with [SOME] special words"

	// Pattern explanation:
	// \[      => matches literal '['
	// ([^\[\]]*) => captures any character except '[' or ']' zero or more times
	// \]      => matches literal ']'
	re := regexp.MustCompile(`\[([^\[\]]*)\]`)
	fmt.Println("Pattern:", re.String())
	fmt.Println("Matched:", re.MatchString(str))

	// Extract and trim brackets
	submatches := re.FindAllString(str, -1)
	for _, element := range submatches {
		fmt.Println("With brackets:", element)
		element = strings.Trim(element, "[]")
		fmt.Println("Without brackets:", element)
	}
}

// -------------------------------
// Extract all Non-Alphanumeric Characters from a String
func RegexNonAlphaChar() {
	str := "We @@@Love@@@@ #Go!$! ****Programming****Language^^^"

	// Pattern: [^a-zA-Z0-9]+ => match one or more characters that are NOT letters or digits
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	fmt.Println("Pattern:", re.String())
	fmt.Println("Matched:", re.MatchString(str))

	submatches := re.FindAllString(str, -1)
	for _, element := range submatches {
		fmt.Println(element)
	}
}

// -------------------------------
// Extract dates in YYYY-MM-DD format from string
func RegexExtractDate() {
	str := "DOB: 1995-10-03, some text here."

	// Pattern: \d{4}-\d{2}-\d{2} => matches 4-digit year, 2-digit month, 2-digit day
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	fmt.Println("Pattern:", re.String())
	fmt.Println("Matched:", re.MatchString(str))

	submatches := re.FindAllString(str, -1)
	for _, element := range submatches {
		fmt.Println(element)
	}
}

// -------------------------------
// Extract IP addresses from text
func RegexDNS() {
	text := `Some IPs: 118.99.81.204, 202.118.236.130, 62.201.207.9`

	// Pattern to match IPv4 addresses
	re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)
	fmt.Println("Pattern:", re.String())
	fmt.Println("Matched:", re.MatchString(text))

	submatches := re.FindAllString(text, -1)
	for _, element := range submatches {
		fmt.Println(element)
	}
}

// -------------------------------
// Extract domain from URL
func RegexURL() {
	url := "http://www.example.co.uk/product/1"

	// Pattern explanation:
	// ^(?:https?:\/\/)? => optional http/https at start
	// (?:www\.)? => optional www.
	// ([^:\/\n]+) => capture domain
	re := regexp.MustCompile(`^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?([^:\/\n]+)`)
	fmt.Println("Pattern:", re.String())
	fmt.Println("Matched:", re.MatchString(url))

	submatches := re.FindAllString(url, -1)
	for _, element := range submatches {
		fmt.Println(element)
	}
}

// -------------------------------
// Validate email addresses
func RegexEmailAdd() {
	emails := []string{
		"ç$€§/az@gmail.com",
		"abcd@gmail_yahoo.com",
		"abcd@gmail-yahoo.com",
		"abcd@gmailyahoo",
		"abcd@gmail.yahoo",
	}

	re := regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`)
	fmt.Println("Pattern:", re.String())

	for _, email := range emails {
		fmt.Printf("Email: %v, Valid? %v\n", email, re.MatchString(email))
	}
}

// -------------------------------
// Validate phone numbers (various formats)
func RegexPhonenumber() {
	phones := []string{
		"1(234)5678901x1234",
		"(+351) 282 43 50 50",
		"90191919908",
		"555-8909",
		"001 6867684",
		"001 6867684x1",
		"1 (234) 567-8901",
		"1-234-567-8901 ext1234",
	}

	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	fmt.Println("Pattern:", re.String())

	for _, phone := range phones {
		fmt.Printf("Phone: %v, Valid? %v\n", phone, re.MatchString(phone))
	}
}

// -------------------------------
// Validate dates in "DD/MM/YYYY" format
func RegexDate() {
	dates := []string{
		"31/07/2010",
		"1/13/2010",
		"29/2/2007",
		"31/08/2010",
		"29/02/200a",
		"55/02/200a",
		"2_/02/2009",
	}

	// Pattern explanation:
	// (0?[1-9]|[12][0-9]|3[01]) => day 01-31
	// (0?[1-9]|1[012]) => month 01-12
	// ((19|20)\d\d) => year 1900-2099
	re := regexp.MustCompile(`(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\d\d)`)
	fmt.Println("Pattern:", re.String())

	for _, date := range dates {
		fmt.Printf("Date: %v, Valid? %v\n", date, re.MatchString(date))
	}
}
