package part_3

import (
	"fmt"
	"testing"
)

func TestSingleByteXor(t *testing.T) {
	encryptedDataString := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	ak := "X"
	ap := "Cooking MC's like a pound of bacon"
	s := BruteForceDecrypt(encryptedDataString)
	n := s.ScoreTexts()
	a := n.GetHighScore()
	if a.key != ak || a.plainText != ap {
		t.Error("Not correct answer")
	} else {
		fmt.Printf("the key is %v \n decrypted plaintext is %s \n", a.key, a.plainText)
	}
}
