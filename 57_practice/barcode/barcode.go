package barcode

import (
	"fmt"
	"image/png"
	"net/http"
	"os"
	"text/template"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

// Page struct used to pass data to HTML template
type Page struct {
	Title string
}

// ExampleBarcode starts a simple HTTP server to serve the QR code generator
func ExampleBarcode() {
	// Route for home page with the form
	http.HandleFunc("/", homeHandler)

	// Route for generating and displaying QR codes
	http.HandleFunc("/generator/", viewCodeHandler)

	fmt.Println("Server running at http://localhost:8080/")

	// Start the HTTP server on port 8080
	http.ListenAndServe(":8080", nil)
}

// homeHandler serves the HTML page containing the QR code form
func homeHandler(w http.ResponseWriter, r *http.Request) {
	p := Page{Title: "QR Code Generator"}

	// Debug log: show current working directory
	cwd, _ := os.Getwd()
	fmt.Println("Current Working Directory:", cwd)

	// Parse the HTML template file
	t, err := template.ParseFiles("barcode/templates/generator.html")
	if err != nil {
		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template, injecting the Page struct
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

// viewCodeHandler generates a QR code from the input data string
func viewCodeHandler(w http.ResponseWriter, r *http.Request) {
	// Get the input string from the form
	dataString := r.FormValue("dataString")

	// Encode the string as a QR code
	qrCode, err := qr.Encode(dataString, qr.L, qr.Auto)
	if err != nil {
		http.Error(w, "Error generating QR code: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Scale the QR code to 512x512 pixels
	qrCode, err = barcode.Scale(qrCode, 512, 512)
	if err != nil {
		http.Error(w, "Error scaling QR code: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Set response content type as PNG image
	w.Header().Set("Content-Type", "image/png")

	// Write the QR code image to the HTTP response
	png.Encode(w, qrCode)
}
