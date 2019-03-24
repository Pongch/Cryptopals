// part_1 of set 1, convert hex to base64
package part_1

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

func Encode(s string) string {
	decoded, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(decoded)
}
