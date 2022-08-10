package helper

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	lowerCharSet = "abcdefghij"
	upperCharSet = "KLMNOPQRS"
	numberSet    = "0123456789"
	allCharSet   = lowerCharSet + upperCharSet + numberSet
)

func HashSha1Pass() string {

	rand.Seed(time.Now().Unix())
	minSpecialChar := 1
	minNum := 1
	minUpperCase := 1
	passwordLength := 8

	pwd := generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase)
	var sha = sha1.New()
	sha.Write([]byte(pwd))
	var encrypted = sha.Sum(nil)

	var encryptedString = fmt.Sprintf("%x", encrypted)

	return encryptedString
}

func generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase int) string {

	var password strings.Builder

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}

	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})

	return string(inRune)
}
