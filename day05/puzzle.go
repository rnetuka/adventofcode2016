package day05

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
)

func Hash(doorId string, index int) string {
	text := doorId + strconv.Itoa(index)
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func FindHash(doorId string, index int) (string, int) {
	for true {
		hash := Hash(doorId, index)
		if strings.HasPrefix(hash, "00000") {
			return hash, index
		} else {
			index++
		}
	}
	return "", -1
}

func Solution1() string {
	password := make([]rune, 8)
	var hash string
	doorId := "ugkcyxxp"
	index := 0
	for i := 0; i < len(password); i++ {
		hash, index = FindHash(doorId, index)
		password[i] = rune(hash[5])
		index++
	}
	return string(password)
}

func passwordReady(password []rune) bool {
	for i := 0; i < len(password); i++ {
		if password[i] == 0 {
			return false
		}
	}
	return true
}

func Solution2() string {
	password := make([]rune, 8)
	var hash string
	doorId := "ugkcyxxp"
	index := 0
	for true {
		hash, index = FindHash(doorId, index)
		j, _ := strconv.ParseInt(hash[5:6], 16, 8)
		if j < 8 && password[j] == 0 {
			password[j] = rune(hash[6])
		}
		if passwordReady(password) {
			break
		}
		index++
	}
	return string(password)
}