// take the hex encoded string, XOR against with an array of alphabets as key
// return the key that generates the best score

package part_3

import (
	"fmt"
	"encoding/hex"
	"strings"
)


func SingleByteXor() {
	
	
	a := strings.Split(strings.ToUpper("abcdefghijklmnopqrstuvwxyz"), "")
	for _, v := range a {
		data := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
		res, _ := hex.DecodeString(data)

		decrypted := EncryptDecrypt(string(res), v)
		fmt.Println("Decrypted:", string(decrypted))
	}
}

func EncryptDecrypt(input, key string) (output string) {
	for i := range input {
		output += string(input[i] ^ key[i%len(key)])
	}

	return output
}

