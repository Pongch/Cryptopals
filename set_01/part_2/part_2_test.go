package part_2

import "testing"

func TestXor(t *testing.T) {
	s := "1c0111001f010100061a024b53535009181c"
	x := "686974207468652062756c6c277320657965"
	expect := "746865206b696420646f6e277420706c6179"
	res := FixedXor(s, x)
	if expect != res {
		t.Errorf("not correct Xor, expected %v, got %v", expect, res)
	}
}
