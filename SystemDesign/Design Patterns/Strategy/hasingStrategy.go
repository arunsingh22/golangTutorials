package main

// import (
// 	"crypto/sha256"
// 	"fmt"

// 	"golang.org/x/crypto/bcrypt"
// )

// // Strategy function type for encryption/decryption (though bcrypt is one-way)
// type CryptoStrategy func(string) (string, error)

// // Context: Holds the currently selected crypto strategy
// type Cryptographer struct {
// 	strategy CryptoStrategy
// }

// // SetStrategy: Allows changing the active strategy at runtime
// func (c *Cryptographer) SetStrategy(s CryptoStrategy) {
// 	c.strategy = s
// }

// // Execute: Uses the current strategy to perform the operation
// func (c *Cryptographer) Execute(data string) (string, error) {
// 	if c.strategy == nil {
// 		return "", fmt.Errorf("no crypto strategy set")
// 	}
// 	return c.strategy(data)
// }

// // SHA256 Hashing Strategy
// func sha256Hash(data string) (string, error) {
// 	fmt.Println("Performing SHA256 hashing...")
// 	hasher := sha256.New()
// 	_, err := hasher.Write([]byte(data))
// 	if err != nil {
// 		return "", err
// 	}
// 	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
// }

// // Bcrypt Hashing Strategy (for password hashing, not reversible decryption)
// func bcryptHash(data string) (string, error) {
// 	fmt.Println("Performing Bcrypt hashing...")
// 	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(hashedBytes), nil
// }

// // Bcrypt Verification Strategy (separate function for verification)
// func bcryptVerify(plainText string, hashedPassword string) error {
// 	fmt.Println("Performing Bcrypt verification...")
// 	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainText))
// }

// func main() {
// 	message := "This is a secret message."
// 	password := "mysecretpassword"

// 	cryptographer := Cryptographer{}

// 	// SHA256 Encryption (Hashing)
// 	cryptographer.SetStrategy(sha256Hash)
// 	sha256HashResult, err := cryptographer.Execute(message)
// 	if err != nil {
// 		fmt.Println("SHA256 Error:", err)
// 	} else {
// 		fmt.Println("SHA256 Hash:", sha256HashResult)
// 	}

// 	fmt.Println()

// 	// Bcrypt Encryption (Hashing)
// 	cryptographer.SetStrategy(bcryptHash)
// 	bcryptHashResult, err := cryptographer.Execute(password)
// 	if err != nil {
// 		fmt.Println("Bcrypt Hash Error:", err)
// 	} else {
// 		fmt.Println("Bcrypt Hash:", bcryptHashResult)
// 	}

// 	fmt.Println()

// 	// Bcrypt Verification
// 	err = bcryptVerify(password, bcryptHashResult)
// 	if err == nil {
// 		fmt.Println("Bcrypt Verification: Password matches!")
// 	} else {
// 		fmt.Println("Bcrypt Verification: Password does NOT match:", err)
// 	}

// 	wrongPassword := "incorrectpassword"
// 	err = bcryptVerify(wrongPassword, bcryptHashResult)
// 	if err == nil {
// 		fmt.Println("Bcrypt Verification (Incorrect): Password matches!")
// 	} else {
// 		fmt.Println("Bcrypt Verification (Incorrect): Password does NOT match:", err)
// 	}
// }
