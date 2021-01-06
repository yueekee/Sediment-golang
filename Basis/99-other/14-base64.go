package main

import "fmt"

func main() {
	toString := URLEncoding.EncodeToString([]byte("31"))
	fmt.Println(toString)
}

type Encoding struct {
	encode    [64]byte
	decodeMap [256]byte
	padChar   rune
	strict    bool
}

var URLEncoding = NewEncoding(encodeURL)

const encodeURL = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
const (
	StdPadding rune = '=' // Standard padding character
	NoPadding  rune = -1  // No padding
)

func NewEncoding(encoder string) *Encoding {
	if len(encoder) != 64 {
		panic("encoding alphabet is not 64-bytes long")
	}
	for i := 0; i < len(encoder); i++ {
		if encoder[i] == '\n' || encoder[i] == '\r' {
			panic("encoding alphabet contains newline character")
		}
	}

	e := new(Encoding)
	e.padChar = StdPadding
	copy(e.encode[:], encoder)

	for i := 0; i < len(e.decodeMap); i++ {
		e.decodeMap[i] = 0xFF
	}

	for i := 0; i < len(encoder); i++ {
		e.decodeMap[encoder[i]] = byte(i)
	}

	return e
}

func (enc *Encoding) EncodeToString(src []byte) string {
	fmt.Println("len(src):", len(src))
	fmt.Println("enc.EncodedLen(len(src)):", enc.EncodedLen(len(src)))

	buf := make([]byte, enc.EncodedLen(len(src)))

	enc.Encode(buf, src)
	return string(buf)
}

func (enc *Encoding) EncodedLen(n int) int {
	if enc.padChar == NoPadding {
		return (n*8 + 5) / 6 // minimum # chars at 6 bits  perchar
	}
	return (n + 2) / 3 * 4 // minimum # 4-char quanta, 3 bytes each
}

func (enc *Encoding) Encode(dst, src []byte) {
	if len(src) == 0 {
		return
	}
	// enc is a pointer receiver, so the use of enc.encode within the hot
	// loop below means a nil check at every operation. Lift that nil check
	// outside of the loop to speed up the encoder.
	_ = enc.encode

	di, si := 0, 0
	n := (len(src) / 3) * 3
	fmt.Println("n:", n)
	for si < n {
		// Convert 3x 8bit source bytes into 4 bytes
		val := uint(src[si+0])<<16 | uint(src[si+1])<<8 | uint(src[si+2])

		dst[di+0] = enc.encode[val>>18&0x3F]
		dst[di+1] = enc.encode[val>>12&0x3F]
		dst[di+2] = enc.encode[val>>6&0x3F]
		dst[di+3] = enc.encode[val&0x3F]

		si += 3
		di += 4
	}

	remain := len(src) - si
	if remain == 0 {
		return
	}
	// Add the remaining small block
	fmt.Println("si:", si)
	fmt.Println("si+0:", si+0)
	fmt.Println("src[si+0]:", src[si+0])

	val := uint(src[si+0]) << 16
	fmt.Println("val:", val)
	if remain == 2 {
		val |= uint(src[si+1]) << 8
	}

	dst[di+0] = enc.encode[val>>18&0x3F]
	dst[di+1] = enc.encode[val>>12&0x3F]

	switch remain {
	case 2:
		dst[di+2] = enc.encode[val>>6&0x3F]
		if enc.padChar != NoPadding {
			dst[di+3] = byte(enc.padChar)
		}
	case 1:
		if enc.padChar != NoPadding {
			dst[di+2] = byte(enc.padChar)
			dst[di+3] = byte(enc.padChar)
		}
	}
}