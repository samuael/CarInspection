package helper

import (
	"encoding/json"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
	"github.com/samuael/Project/CarInspection/pkg/constants/state"
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))


// NUMBERS const numbers
const NUMBERS = "1234567890"

// CHARACTERS const field
const CHARACTERS = "abcdefghijelmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_1234567890"

// GenerateRandomString  function
func GenerateRandomString(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// MarshalThis function
func MarshalThis(inter interface{}) []byte {
	val, era := json.Marshal(inter)
	if era != nil {
		return nil
	}
	return val
}


// IsImage function checking whether the file is an image or not
func IsImage(filepath string ) bool {
	extension := GetExtension(filepath)
	extension = strings.ToLower(extension)
	for _ , e := range state.ImageExtensions  {
		if e== extension {
			return true 
		}
	}
	return  false 
}

// GetExtension function to return the extension of the File Input FileName
func GetExtension(Filename string) string {
	fileSlice := strings.Split(Filename, ".")
	if len(fileSlice) >= 1 {
		return fileSlice[len(fileSlice)-1]
	}
	return ""
}
// ValidateUsername  function to validate whether the string is a valid Username or not
func ValidateUsername(  username string , minLength uint  ) bool {
	trim := func()bool {
		name := strings.Trim(username  , " ") 
		return (len(name) < int(minLength ))
	}

	numbercheck := func() bool{
		_ , err := strconv.Atoi(username)
		return (err != nil)
		}

	if (len( username) < int(minLength))  || 
		trim() ||
		numbercheck(){
		return false 
	}
	return true
}
// ValidatePassword  function to validate whether the string is a valid Username or not
func ValidatePassword(  password string , minLength uint  ) bool {
	if (len( password) < int(minLength))  || 
		(
	func()bool {
		name := strings.Trim(password  , " ") 
		return (len(name) < int(minLength ))
	}()){
		return false 
	}
	return true
}

// EmailRX represents email address maching pattern
var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// MatchesPattern checks if a given input form field matchs a given pattern
func MatchesPattern(email string, pattern *regexp.Regexp) bool  {
	if email == "" {
		return false 
	}
	if !pattern.MatchString(email) {
		return false
	}
	return true
}