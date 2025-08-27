package photogallery

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// tpl stores parsed HTML templates for rendering the photo gallery
var tpl *template.Template

func init() {
	// Parse all template files in the templates folder
	// Must panics if there is an error in parsing
	tpl = template.Must(template.ParseGlob("photo_gallery/templates/*"))
}

// PhotoGallery starts the HTTP server and sets up routes for the gallery
func PhotoGallery() {
	// Route for homepage
	http.HandleFunc("/", index)

	// Serve static files (images, CSS, JS) from ./static directory
	// http.StripPrefix removes /static from URL path
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	// Handle favicon requests (return 404)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	// Start HTTP server on port 8080
	sd := http.ListenAndServe(":8080", nil)
	fmt.Println(sd)
}

// index handles requests to the homepage
// - GET: renders the gallery
// - POST: handles file upload
func index(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		// Retrieve uploaded file from form field "nf"
		mf, fh, err := req.FormFile("nf")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer mf.Close()

		// Generate SHA1 hash of file content for unique filename
		ext := strings.Split(fh.Filename, ".")[1] // extract file extension
		h := sha1.New()
		io.Copy(h, mf) // copy file data into hash
		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext // hashed filename

		// Determine path to save the file
		wd, err := os.Getwd() // get current working directory
		if err != nil {
			fmt.Println(err)
			return
		}
		path := filepath.Join(wd, "static/images", fname)

		// Create new file on disk
		nf, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer nf.Close()

		// Reset file pointer and copy uploaded content to the new file
		mf.Seek(0, 0)
		io.Copy(nf, mf)
	}

	// Open image directory
	file, err := os.Open("static/images")
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer file.Close()

	// Read all file names in the directory
	list, _ := file.Readdirnames(0) // 0 reads all files/folders

	// Render template with list of image filenames
	tpl.ExecuteTemplate(w, "index.gohtml", list)
}
