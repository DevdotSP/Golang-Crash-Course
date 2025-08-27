package diffuc

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
	"unicode"
)

// ExampleCryptography demonstrates hashing, HMAC, and Caesar cipher encryption/decryption
func ExampleCryptography() {
	// -------------------------------
	// Caesar Cipher Examples
	// -------------------------------
	// c := NewCaesar(1)
	// fmt.Println("Encrypt Key(01) abcd =>", c.Encryption("abcd"))
	// fmt.Println("Decrypt Key(01) bcde =>", c.Decryption("bcde"))
	// fmt.Println()

	// c = NewCaesar(10)
	// fmt.Println("Encrypt Key(10) abcd =>", c.Encryption("abcd"))
	// fmt.Println("Decrypt Key(10) klmn =>", c.Decryption("klmn"))
	// fmt.Println()

	// c = NewCaesar(15)
	// fmt.Println("Encrypt Key(15) abcd =>", c.Encryption("abcd"))
	// fmt.Println("Decrypt Key(15) pqrs =>", c.Decryption("pqrs"))

	// -------------------------------
	// HMAC with Salt Example
	// -------------------------------
	// SampleToGenerateSalt()

	// -------------------------------
	// Hashing Examples
	// -------------------------------
	// SampleHashing()
}

var secretKey = "4234kxzjcjj3nxnxbcvsjfj"

// generateSalt generates a cryptographically secure random salt
func generateSalt() string {
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		fmt.Println("Failed generating salt:", err)
		return ""
	}
	// Encode bytes to URL-safe Base64 string
	return base64.URLEncoding.EncodeToString(randomBytes)
}

// SampleToGenerateSalt shows HMAC signing with SHA256 and SHA512 using a random salt
func SampleToGenerateSalt() {
	message := "Today web engineering has modern apps adhere to SPA model."
	salt := generateSalt()
	fmt.Println("Message: " + message)
	fmt.Println("Salt: " + salt)

	// HMAC with SHA256
	hash := hmac.New(sha256.New, []byte(secretKey))
	io.WriteString(hash, message+salt)
	fmt.Printf("HMAC-SHA256: %x\n", hash.Sum(nil))

	// HMAC with SHA512
	hash = hmac.New(sha512.New, []byte(secretKey))
	io.WriteString(hash, message+salt)
	fmt.Printf("HMAC-SHA512: %x\n", hash.Sum(nil))
}

// SampleHashing demonstrates various hash algorithms for small and large messages
func SampleHashing() {
	// Small message
	message := []byte("Small message example.")
	fmt.Printf("MD5: %x\n", md5.Sum(message))
	fmt.Printf("SHA1: %x\n", sha1.Sum(message))
	fmt.Printf("SHA256: %x\n", sha256.Sum256(message))
	fmt.Printf("SHA512: %x\n", sha512.Sum512(message))

	// Large message example
	message = []byte("This is a larger message for demonstrating MD5, SHA1, SHA256, SHA512 hashing in Go.")
	fmt.Printf("MD5: %x\n", md5.Sum(message))
	fmt.Printf("SHA1: %x\n", sha1.Sum(message))
	fmt.Printf("SHA256: %x\n", sha256.Sum256(message))
	fmt.Printf("SHA512: %x\n", sha512.Sum512(message))
}

// -------------------------------
// Caesar Cipher Implementation
// -------------------------------

// Cipher interface defines encryption and decryption methods
type Cipher interface {
	Encryption(string) string
	Decryption(string) string
}

// cipher stores the shift value(s) for the Caesar cipher
type cipher []int

// cipherAlgorithm performs character shifting
func (c cipher) cipherAlgorithm(letters string, shift func(int, int) int) string {
	shiftedText := ""
	for _, letter := range letters {
		if !unicode.IsLetter(letter) {
			continue // skip non-letter characters
		}
		shiftDist := c[len(shiftedText)%len(c)]
		s := shift(int(unicode.ToLower(letter)), shiftDist)
		// Wrap around alphabet
		switch {
		case s < 'a':
			s += 'z' - 'a' + 1
		case s > 'z':
			s -= 'z' - 'a' + 1
		}
		shiftedText += string(s)
	}
	return shiftedText
}

// Encryption applies Caesar cipher to plaintext
func (c *cipher) Encryption(plainText string) string {
	return c.cipherAlgorithm(plainText, func(a, b int) int { return a + b })
}

// Decryption reverses Caesar cipher
func (c *cipher) Decryption(cipherText string) string {
	return c.cipherAlgorithm(cipherText, func(a, b int) int { return a - b })
}

// NewCaesar creates a new Caesar cipher with a shift key
func NewCaesar(key int) Cipher {
	return NewShift(key)
}

// NewShift validates shift key and returns a Cipher
func NewShift(shift int) Cipher {
	if shift < -25 || shift > 25 || shift == 0 {
		return nil // invalid shift
	}
	c := cipher([]int{shift})
	return &c
}
