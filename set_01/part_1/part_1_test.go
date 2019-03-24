package part_1

import "testing"

//should output correct result
func TestEncoding(t *testing.T) {
	inputHex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expectOutput := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	res := Encode(inputHex)
	if expectOutput != res {
		t.Errorf("not correct encoding, expected %s, got %v ", expectOutput, res)
	}

}
