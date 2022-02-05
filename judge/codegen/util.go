package codegen

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

const (
	Python     = "Python"
	Cpp        = "Cpp"
	Go         = "Go"
	Javascript = "Javascript"
)

var (
	LanguageToExt map[string]string = map[string]string{
		Python:     ".py",
		Cpp:        ".cpp",
		Go:         ".go",
		Javascript: ".js",
	}

	Languages          = []string{Python, Cpp, Go, Javascript}
	ErrInvalidLanguage = errors.New("invalid language")
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func randomSecret(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

func RandomSecret() string {
	return randomSecret(32)
}

func RandomName() string {
	return randomSecret(8)
}
