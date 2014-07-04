package base85

import (
	"fmt"
)

type Encoding struct {
	representation []byte
}

func (enc *Encoding) Encode(data []byte) (string, error) {
	dataLen := len(data)
	if dataLen%4 != 0 {
		return "", fmt.Errorf("Base85 encode: Length of data not a multiple of 4")
	}

	strLen := 5 * dataLen / 4
	strBytes := make([]byte, strLen+1)

	for dataCount, strCount, value := 0, 0, uint(0); dataCount < dataLen; {
		value = value*256 + uint(data[dataCount])
		dataCount++
		if dataCount%4 == 0 {
			for divisor := uint(85 * 85 * 85 * 85); divisor != 0; divisor /= 85 {
				strBytes[strCount] = enc.representation[value/divisor%85]
				strCount++
			}
			value = 0
		}
	}

	return string(strBytes[:strLen]), nil
}

func (enc *Encoding) Decode(str string) ([]byte, error) {
	strBytes := []byte(str)
	strLen := len(strBytes)
	if strLen%5 != 0 {
		return nil, fmt.Errorf("Base85 decode: Length of Base85 string not a multiple of 5")
	}

	dataLen := 4 * strLen / 5
	data := make([]byte, dataLen)

	for dataCount, strCount, value := 0, 0, uint(0); strCount < strLen; {
		rep := strBytes[strCount]
		repVal, err := enc.repToVal(rep)
		if err != nil {
			return nil, err
		}

		value = value*85 + repVal
		strCount++
		if strCount%5 == 0 {
			for divisor := uint(256 * 256 * 256); divisor != 0; divisor /= 256 {
				data[dataCount] = byte(value / divisor % 256)
				dataCount++
			}
			value = 0
		}
	}
	return data[:dataLen], nil
}

func (enc *Encoding) repToVal(b byte) (uint, error) {
	for i, r := range enc.representation {
		if b == r {
			return uint(i), nil
		}
	}

	return 0, fmt.Errorf("Base85 decode: invalid representation %q", b)
}

// Ascii85Encoding use original representation of Base85.
var Ascii85Encoding = &Encoding{
	representation: []byte(
		"!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstu",
	),
}

// Z85Encoding use representation from ZeroMQ version of Base85.
var Z85Encoding = &Encoding{
	representation: []byte(
		"0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ.-:+=^!/*?&<>()[]{}@%$#",
	),
}

// RFC1924Encoding use representation from RFC1924 version of Base85.
var RFC1924Encoding = &Encoding{
	representation: []byte(
		"0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!#$%&()*+-;<=>?@^_`{|}~",
	),
}
