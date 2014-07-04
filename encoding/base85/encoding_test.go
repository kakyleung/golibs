package base85_test

import (
	"encoding/ascii85"
	"fmt"
	"github.com/wancw/golibs/encoding/base85"
	"testing"
)

func ExampleAllEncodings() {
	data := []byte{0x86, 0x4F, 0xD2, 0x6F, 0xB5, 0x59, 0xF7, 0x5B}

	encodings := []*base85.Encoding{
		base85.Ascii85Encoding,
		base85.Z85Encoding,
		base85.RFC1924Encoding,
	}
	for _, enc := range encodings {
		str, _ := enc.Encode(data)
		fmt.Println(str)
	}

	// Compare with built-in package "encoding/ascii85"
	encoded := make([]byte, ascii85.MaxEncodedLen(len(data)))
	ascii85.Encode(encoded, data)
	fmt.Println(string(encoded))

	// Output:
	// L/669[9<6.
	// HelloWorld
	// hELLOwORLD
	// L/669[9<6.
}

func TestAscii85Encoding(t *testing.T) {
	// data := []byte("Man is distinguished, not only by his reason, but by this singular passion from other animals, which is a lust of the mind, that by a perseverance of delight in the continued and indefatigable generation of knowledge, exceeds the short vehemence of any carnal pleasure.")
	// const expected = "9jqo^BlbD-BleB1DJ+*+F(f,q/0JhKF<GL>Cj@.4Gp$d7F!,L7@<6@)/0JDEF<G%<+EV:2F!,O<DJ+*.@<*K0@<6L(Df-\\0Ec5e;DffZ(EZee.Bl.9pF\"AGXBPCsi+DGm>@3BB/F*&OCAfu2/AKYi(DIb:@FD,*)+C]U=@3BN#EcYf8ATD3s@q?d$AftVqCh[NqF<G:8+EV:.+Cf>-FD5W8ARlolDIal(DId<j@<?3r@:F%a+D58'ATD4$Bl@l3De:,-DJs`8ARoFb/0JMK@qB4^F!,R<AKZ&-DfTqBG%G>uD.RTpAKYo'+CT/5+Cei#DII?(E,9)oF*2M7/c"

	data := []byte("Man is distinguished, not only by his reason, but by this singular passion from other animals, which is a lust of the mind, that by a perseverance of delight in the continued and indefatigable generation of knowledge, exceeds the short vehemence of any carnal pleasure")
	const expected = "9jqo^BlbD-BleB1DJ+*+F(f,q/0JhKF<GL>Cj@.4Gp$d7F!,L7@<6@)/0JDEF<G%<+EV:2F!,O<DJ+*.@<*K0@<6L(Df-\\0Ec5e;DffZ(EZee.Bl.9pF\"AGXBPCsi+DGm>@3BB/F*&OCAfu2/AKYi(DIb:@FD,*)+C]U=@3BN#EcYf8ATD3s@q?d$AftVqCh[NqF<G:8+EV:.+Cf>-FD5W8ARlolDIal(DId<j@<?3r@:F%a+D58'ATD4$Bl@l3De:,-DJs`8ARoFb/0JMK@qB4^F!,R<AKZ&-DfTqBG%G>uD.RTpAKYo'+CT/5+Cei#DII?(E,9)oF*2M7"

	str, err := base85.Ascii85Encoding.Encode(data)
	if err != nil {
		t.Fatal("Encoding Failed", err)
	}

	sample := make([]byte, ascii85.MaxEncodedLen(len(data)))
	sampleLen := ascii85.Encode(sample, data)

	if str != expected {
		t.Fatalf("\nExpected         : %q,\nBut got          : %q.\nBuilt-in package : %q\n", expected, str, sample[:sampleLen])
	}
}

func TestAscii85Decoding(t *testing.T) {
	data := "9jqo^BlbD-BleB1DJ+*+F(f,q/0JhKF<GL>Cj@.4Gp$d7F!,L7@<6@)/0JDEF<G%<+EV:2F!,O<DJ+*.@<*K0@<6L(Df-\\0Ec5e;DffZ(EZee.Bl.9pF\"AGXBPCsi+DGm>@3BB/F*&OCAfu2/AKYi(DIb:@FD,*)+C]U=@3BN#EcYf8ATD3s@q?d$AftVqCh[NqF<G:8+EV:.+Cf>-FD5W8ARlolDIal(DId<j@<?3r@:F%a+D58'ATD4$Bl@l3De:,-DJs`8ARoFb/0JMK@qB4^F!,R<AKZ&-DfTqBG%G>uD.RTpAKYo'+CT/5+Cei#DII?(E,9)oF*2M7"
	expected := []byte("Man is distinguished, not only by his reason, but by this singular passion from other animals, which is a lust of the mind, that by a perseverance of delight in the continued and indefatigable generation of knowledge, exceeds the short vehemence of any carnal pleasure")

	decoded, err := base85.Ascii85Encoding.Decode(data)
	if err != nil {
		t.Fatal("Encoding Failed", err)
	}

	if str, expectedStr := string(decoded), string(expected); str != expectedStr {
		t.Fatalf("\nExpected : %q,\nBut got  : %q.\n", expectedStr, str)
	}
}

func TestAscii85Encoding_AllZero(t *testing.T) {
	t.Skipf("Not implemented yet.")

	data := []byte{0x00, 0x00, 0x00, 0x00}
	const expected = "z"

	sample := make([]byte, ascii85.MaxEncodedLen(len(data)))
	sampleLen := ascii85.Encode(sample, data)

	encodings := []*base85.Encoding{
		base85.RFC1924Encoding,
		base85.Z85Encoding,
		base85.Ascii85Encoding,
	}

	for i, enc := range encodings {
		str, err := enc.Encode(data)
		if err != nil {
			t.Fatal("Encoding Failed", err)
		}

		if str != expected {
			t.Logf("%d\nExpected         : %q,\nBut got          : %q.\nBuilt-in package : %q\n", i, expected, str, sample[:sampleLen])
		}
	}

}

func TestZ85Encoding(t *testing.T) {
	data := []byte{0x86, 0x4F, 0xD2, 0x6F, 0xB5, 0x59, 0xF7, 0x5B}
	str, err := base85.Z85Encoding.Encode(data)
	if err != nil {
		t.Fatal("Encoding Failed", err)
	}

	const expected = "HelloWorld"
	if str != expected {
		t.Fatalf("Expected %q, but got %q.\n", expected, str)
	}
}

func TestZ85Decoding(t *testing.T) {
	str := "HelloWorld"
	data, err := base85.Z85Encoding.Decode(str)
	if err != nil {
		t.Fatal("Decoding Failed", err)
	}

	expected := []byte{0x86, 0x4F, 0xD2, 0x6F, 0xB5, 0x59, 0xF7, 0x5B}
	if dataStr, expectedStr := string(data), string(expected); dataStr != expectedStr {
		t.Fatalf("Expected %q, but got %q.\n", expectedStr, dataStr)
	}
}

func TestRFC1924Encoding(t *testing.T) {
	data := []byte{0x86, 0x4F, 0xD2, 0x6F, 0xB5, 0x59, 0xF7, 0x5B}
	str, err := base85.RFC1924Encoding.Encode(data)
	if err != nil {
		t.Fatal("Encoding Failed", err)
	}

	const expected = "hELLOwORLD"
	if str != expected {
		t.Fatalf("Expected %q, but got %q.\n", expected, str)
	}
}

func TestRFC1924Decoding(t *testing.T) {
	str := "hELLOwORLD"
	data, err := base85.RFC1924Encoding.Decode(str)
	if err != nil {
		t.Fatal("Decoding Failed", err)
	}

	expected := []byte{0x86, 0x4F, 0xD2, 0x6F, 0xB5, 0x59, 0xF7, 0x5B}
	if dataStr, expectedStr := string(data), string(expected); dataStr != expectedStr {
		t.Fatalf("Expected %q, but got %q.\n", expectedStr, dataStr)
	}
}
