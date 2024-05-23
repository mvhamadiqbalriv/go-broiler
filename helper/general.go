package helper

import (
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

var logger *logrus.Logger

type User struct {
    Password string
}

func HashAndSalt(pwd []byte) string {
    hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
    if err != nil {
        log.Fatalf("Error generating hash: %v", err)
    }
    return string(hash)
}

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
    byteHash := []byte(hashedPwd)
    err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
    if err != nil {
        log.Println("Password comparison failed:", err)
        return false
    }
    return true
}

func ExtractTokenFromHeader(r *http.Request) string {
    // Get the value of the Authorization header
    authHeader := r.Header.Get("Authorization")
    if authHeader == "" {
        return "" // No Authorization header found
    }

    // Split the Authorization header into parts
    parts := strings.Split(authHeader, " ")
    if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
        return "" // Invalid Authorization header format
    }

    // Return the JWT token
    return parts[1]
}

func SaveFile(filePath string, data []byte) error {
    err := os.WriteFile(filePath, data, 0644)
    if err != nil {
        return err
    }
    return nil
}

func DeleteFile(filePath string) error {
    err := os.Remove(filePath)
    if err != nil {
        return err
    }
    return nil
}

func DecodeBase64(base64String string) ([]byte, error) {
    data, err := base64.StdEncoding.DecodeString(base64String)
    if err != nil {
        return nil, err
    }
    return data, nil
}

func GenerateRandomString(length int, prefix string) string {
    if prefix == "" {
        prefix = "filename"
    }
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    b := make([]byte, length)
    for i := range b {
        b[i] = charset[rand.Intn(len(charset))]
    }
    return prefix + "_" + string(b)
}

func OpenLogFile() (*os.File, error) {
	// Initialize a new Logrus logger
	logger := logrus.New()
	// You can customize the logger as needed, for example:
	logger.Formatter = &logrus.JSONFormatter{}

	// Make file name dynamic by date Ymd like application_23052024.log, application_24052024.log, etc
	currentDate := time.Now().Format("20060102")
	// Construct the log filename
	logFilename := "log/application_" + currentDate + ".log"

	// Set output to file
	file, err := os.OpenFile(logFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		// Return error if failed to open or create log file
		return nil, fmt.Errorf("failed to open log file: %v", err)
	}
	logger.Out = file

	// Return the file handle and nil error if successful
	return file, nil
}

//get dot env variable
func GetEnv(key string) string {
    
    err := godotenv.Load(".env")

    if err != nil {
        log.Fatalf("Error loading .env file")
    } 

    return os.Getenv(key)
}
