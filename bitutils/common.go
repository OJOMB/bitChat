package bitutils

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

const pathSeperator = string(os.PathSeparator)

// ConstructPath constructs a filepath from a slice of strings
func ConstructPath(fields []string, fileExtension string) string {
	return strings.Join(fields, pathSeperator) + fileExtension
}

// Convenience function to redirect to the error message page
func errorMessage(w http.ResponseWriter, r *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(w, r, strings.Join(url, ""), 302)
}

// parse HTML templates
// pass in a list of file names, and get a template
func parseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

// version
func version() string {
	return "0.1"
}

func createUUID() (string, error) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		return "", err
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)

	return fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:]), nil
}

// CreateUUID create a random UUID with from RFC 4122
// adapted from http://github.com/nu7hatch/gouuid
func CreateUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

// Encrypt hash plaintext with SHA-1
func Encrypt(plaintext string) (crypt string) {
	crypt = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}
