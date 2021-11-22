package utils

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"strings"
	"unicode"
)

//removeDublicatedSpace ...
func removeDublicatedSpace(str *string) string {
	var buf bytes.Buffer
	var pc rune
	for _, c := range *str {
		if pc == c && pc == ' ' {
			continue
		}

		pc = c
		buf.WriteRune(c)
	}
	fmt.Println(buf.String())
	return buf.String()
}

// PrepareString ...
func PrepareString(str *string, lower bool) string {
	if str == nil {
		return ""
	}
	*str = strings.Trim(*str, " ")
	*str = removeDublicatedSpace(str)
	if lower {
		return strings.ToLower(*str)
	}
	return *str
}

// ConvertToSha ...
func ConvertToSha(password string) string {
	hasher := sha1.New()
	hasher.Write([]byte(password))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

// ValidString ...
func ValidString(str string) error {
	if strings.Trim(str, " ") == "" {
		return fmt.Errorf("Invalid string")
	}
	if strings.Contains(str, "'") {
		return fmt.Errorf("Invalid char")
	}

	return nil
}

// ValidPassword ...
func ValidPassword(pass string) error {
	var (
		upp, low, num, sym bool
		tot                uint8
	)

	for _, char := range pass {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		default:
			return fmt.Errorf("Invalid char")
		}
	}
	if !upp {
		return fmt.Errorf("password should contain upper case char")
	}
	if !low {
		return fmt.Errorf("password should contain lawer case char")
	}
	if !num {
		return fmt.Errorf("password should contain number")
	}
	if !sym {
		return fmt.Errorf("password should contain symbol char")
	}
	if tot < 8 {
		return fmt.Errorf("password length should be more than 8 char")
	}

	return nil
}
