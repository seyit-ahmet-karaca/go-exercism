package cipher

import (
	"strings"
	"unicode"
)

const firstLetterIndex = 'z' - 'a' + 1

type Vigenere struct {
	key string
}

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance == 0 || distance > 25 || distance < -25 {
		return nil
	}
	return NewVigenere(string(rune('a' + (firstLetterIndex+distance)%firstLetterIndex)))
}

func NewVigenere(key string) Cipher {
	isA := false
	isValid := true
	for i := 0; i < len(key); i++ {
		isLower := unicode.IsLower(rune(key[i]))
		if !isLower && isValid {
			isValid = false
		}
		if !isA && key[i] != 'a' {
			isA = true
		}
	}
	if !isA || !isValid {
		return nil
	}

	return &Vigenere{key}
}

func (cipher *Vigenere) Encode(msg string) string {
	msg = strings.ToLower(msg)
	var encrypted string
	var j int
	for i := 0; i < len(msg); i++ {
		if !unicode.IsLetter(rune(msg[i])) {
			continue
		}
		qq := cipher.key[j] - 'a'
		j++
		if j == len(cipher.key) {
			j = 0
		}
		aa := (msg[i] - 'a' + qq) % 26
		encrypted += string(rune(aa + 'a'))
	}
	return encrypted
}

func (cipher *Vigenere) Decode(msg string) string {
	var decrypted string
	for i := 0; i < len(msg); i++ {
		if !unicode.IsLetter(rune(msg[i])) {
			continue
		}
		cipherChar := cipher.key[i%len(cipher.key)]
		decrypted = decrypted + string(rune('a' + ((msg[i]-'a')-(cipherChar-'a')+firstLetterIndex)%firstLetterIndex))
	}
	return decrypted
}
