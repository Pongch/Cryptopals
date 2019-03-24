//encoded
package part_2

import (
	"encoding/hex"
	"log"

	"github.com/hashicorp/vault/helper/xor"
)

func fixedXor(s string, x string) string {
	encoded_s, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}

	encoded_x, err := hex.DecodeString(x)
	if err != nil {
		log.Fatal(err)
	}

	xor_encoded, err := xor.XORBytes(encoded_s, encoded_x)
	if err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(xor_encoded)
}
