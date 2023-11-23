package helpers

import (
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hash))
	return err == nil
}

func GenerateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(900000) + 100000)
}

func GetTokenClaims(c *context.Context) map[string]interface{} {
	token_claims := c.Input.GetData("LoginUserData")
	user_id := token_claims.(jwt.MapClaims)["user_id"]
	user_email := token_claims.(jwt.MapClaims)["user_email"]
	response := map[string]interface{}{"User_id": user_id, "User_Email": user_email}
	return response
}

func UploadFile(fileToUpload multipart.File, fileHeader *multipart.FileHeader, uploadDir string) (string, error) {
	defer fileToUpload.Close()
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename))
	if err := os.MkdirAll(uploadDir, 0777); err != nil {
		return "", fmt.Errorf("failed to create upload directory: %v", err)
	}
	filePath := filepath.Join(uploadDir, filename)
	out, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create the destination file: %v", err)
	}
	defer out.Close()

	_, err = io.Copy(out, fileToUpload)
	if err != nil {
		return "", fmt.Errorf("failed to copy the file: %v", err)
	}

	return filePath, nil
}

func RemoveFile(fileName, directory string) error {
	err := os.Remove(filepath.Join(directory, fileName))
	if err != nil {
		return err
	}
	return nil
}

func SplitFilePath(car string) (string, string) {
	lastIndex := strings.LastIndex(car, "/")

	var fileDirectory string
	var fileName string

	if lastIndex != -1 {
		fileDirectory = car[:lastIndex]
		fileName = car[lastIndex+1:]
	} else {
		fileDirectory = "No '/' found in the string."
		fileName = fileDirectory
	}

	return fileDirectory, fileName
}
