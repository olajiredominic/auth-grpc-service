package helpers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base32"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/lerryjay/auth-grpc-service/pkg/config"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ValidatePasswordHash(hashpassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(password))
	return err == nil
}

func GetOTP(length int) string {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Println("Invalid OTP", err)
		return ""
	}
	return base32.StdEncoding.EncodeToString(randomBytes)[:length]
}

// Function for generating the tokens.
func GenerateToken(payload map[string]string) (string, error) {

	header := "HS256"
	config, err := config.LoadConfig()
	if err != nil {
		return "", err
	}

	jsonStr, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	claims := map[string]string{
		"aud":  config.APP_NAME,
		"iss":  config.APP_URL,
		"exp":  fmt.Sprint(time.Now().Add(time.Minute * 1).Unix()),
		"user": string(jsonStr),
	}

	secret := config.JWT_SECRET

	// create a new hash of type sha256. We pass the secret key to it
	// sha256 is a symmetric cryptographic algorithm
	h := hmac.New(sha256.New, []byte(secret))

	// We base encode the header which is a normal string
	header64 := base64.StdEncoding.EncodeToString([]byte(header))
	// We then Marshal the payload which is a map. This converts it to a string of JSON.
	// Now we base encode this string
	payloadstr, err := json.Marshal(claims)
	if err != nil {
		return string(payloadstr), err
	}
	payload64 := base64.StdEncoding.EncodeToString(payloadstr)

	// Now add the encoded string.
	message := header64 + "." + payload64

	// We have the unsigned message ready. This is simply concat of header and payload
	unsignedStr := header + string(payloadstr)

	// we write this to the SHA256 to hash it. We can use this to generate the signature now
	h.Write([]byte(unsignedStr))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	//Finally we have the token
	tokenStr := message + "." + signature
	return tokenStr, nil
}

// This helps in validating the token
func ValidateJWTToken(token string) (string, error) {

	config, err := config.LoadConfig()
	if err != nil {
		log.Println("Error generating Token")
		return "", err
	}

	secret := config.JWT_SECRET
	// JWT has 3 parts separated by '.'
	splitToken := strings.Split(token, ".")
	// if length is not 3, we know that the token is corrupt
	if len(splitToken) != 3 {
		return "", nil
	}

	// decode the header and payload back to strings
	header, err := base64.StdEncoding.DecodeString(splitToken[0])
	if err != nil {
		return "", err
	}
	payload, err := base64.StdEncoding.DecodeString(splitToken[1])
	if err != nil {
		return "", err
	}

	//again create the signature
	unsignedStr := string(header) + string(payload)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(unsignedStr))

	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	// if both the signature dont match, this means token is wrong
	if signature != splitToken[2] {
		return "", nil
	}

	data := map[string]string{}
	json.Unmarshal(payload, &data)
	// This means the token matches
	return data["user"], nil
}
