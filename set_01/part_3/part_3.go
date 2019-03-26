// take the hex encoded string, XOR against with an array of alphabets as key
// return the key that generates the best score

package part_3

import (
	"encoding/hex"
	"strings"
)

type KeyToScore struct {
	key       string
	plainText string
	score     float64
}

type KeysToScore []KeyToScore

var s string = "!#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[]^_`abcdefghijklmnopqrstuvwxyz{|}~"

func BruteForceDecrypt(encryptedData string) KeysToScore {
	a := strings.Split(s, "")
	ks := []KeyToScore{}
	for _, v := range a {
		res, _ := hex.DecodeString(encryptedData)
		decrypted := EncryptDecrypt(string(res), v)
		k := KeyToScore{
			key:       v,
			plainText: string(decrypted),
			score:     0,
		}
		ks = append(ks, k)
	}
	return ks
}

func (ks KeysToScore) ScoreTexts() KeysToScore {
	var nks KeysToScore

	for i := 0; i < len(ks); i++ {
		val := ks[i]
		txt := val.plainText
		for _, ch := range txt {
			val.score = ScoreText(val.score, ch)
		}
		nks = append(nks, val)
	}
	return nks
}

func ScoreText(score float64, char rune) float64 {
	var s float64
	cc := map[string]float64{
		"e": 13,
		"t": 12,
		"a": 11,
		"o": 10,
		"i": 9,
		"n": 8,
		"s": 7,
		"h": 6,
		"r": 5,
		"d": 4,
		"l": 3,
		"u": 2,
	}
	if cc[string(char)] == 0 {
		s = score - 5
	} else {
		s = score + cc[string(char)]
	}
	return s
}

func (ks KeysToScore) GetHighScore() KeyToScore {
	var as float64
	var ap string
	var ak string
	for i, _ := range ks {
		k := ks[i]
		if k.score > as {
			as = k.score
			ak = k.key
			ap = k.plainText
		}
	}

	return KeyToScore{key: ak, plainText: ap, score: as}
}

func EncryptDecrypt(input, key string) (output string) {
	for i := range input {
		output += string(input[i] ^ key[i%len(key)])
	}

	return output
}
