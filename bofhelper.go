package bofhelper

import (
	"bytes"
	"strconv"
	"strings"
)

//defaultBlackList := []string{`\x00`, `\x0a`, `\x0b`, `\x0c`, `\x0d`, `\x1b`, `\x3f`, `\x5c`, `\x07`, `\x08`, `\x09`, `\x22`, `\x27`}

// enumRSP builds a string to aid in detirmining rsp location from buffer
// m for the amount of prefixed padding to the stub and n for the amount of suffexing padding
func enumRSP(m, n int) (result []byte) {
	var aStr string = "A"
	var bStr string = "BBBB"
	var cStr string = "C"

	gen := bytes.Buffer{}
	for i := 0; i <= m-n-4; i++ {
		gen.WriteString(aStr)
	}
	gen.WriteString(bStr)
	for j := 0; j <= n; j++ {
		gen.WriteString(cStr)
	}
	result = ([]byte(gen.String()))
	return result
}

func createOffsetFuzzer(n int) string {
	builder := strings.Builder{}
	counter := 0
	letters := []string{"0", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for j := 1; j <= 52; j++ {
		for k := 1; k <= 52; k++ {
			for l := 1; l <= 52; l++ {
				builder.WriteString(letters[j] + letters[k] + letters[l])
				counter++
				if (counter) == n {
					return builder.String()
				}
			}
		}
	}
	return "0"
}

//
func userInputBadChars() {}

//converts user input in to hex
func convBadChars() []string {
}

//checks uniqueness of string slice elems
//removing any duplicates
func uniqBlackListsChar() []string {}

//Combine defaultBlack, addtions and new addition into one set
func joinBlacklists() []string {
	uniqBlackListsChar()
	return combinedSet
}

// outputBadChars is either provided with user input string slice of bad characters
// from blacklist struct to remove known string 0-255 in hex to enumerate ascii characters
func outputBadChars(blacklistSet []string) string {
	allChars := []string{`\x00`, `\x01`, `\x02`, `\x03`, `\x04`, `\x05`, `\x06`, `\x07`, `\x08`, `\x09`, `\x0a`, `\x0b`, `\x0c`, `\x0d`, `\x0e`, `\x0f`, `\x10`, `\x11`, `\x12`, `\x13`, `\x14`, `\x15`, `\x16`, `\x17`, `\x18`, `\x19`, `\x1a`, `\x1b`, `\x1c`, `\x1d`, `\x1e`, `\x1f`, `\x20`, `\x21`, `\x22`, `\x23`, `\x24`, `\x25`, `\x26`, `\x27`, `\x28`, `\x29`, `\x2a`, `\x2b`, `\x2c`, `\x2d`, `\x2e`, `\x2f`, `\x30`, `\x31`, `\x32`, `\x33`, `\x34`, `\x35`, `\x36`, `\x37`, `\x38`, `\x39`, `\x3a`, `\x3b`, `\x3c`, `\x3d`, `\x3e`, `\x3f`, `\x40`, `\x41`, `\x42`, `\x43`, `\x44`, `\x45`, `\x46`, `\x47`, `\x48`, `\x49`, `\x4a`, `\x4b`, `\x4c`, `\x4d`, `\x4e`, `\x4f`, `\x50`, `\x51`, `\x52`, `\x53`, `\x54`, `\x55`, `\x56`, `\x57`, `\x58`, `\x59`, `\x5a`, `\x5b`, `\x5c`, `\x5d`, `\x5e`, `\x5f`, `\x60`, `\x61`, `\x62`, `\x63`, `\x64`, `\x65`, `\x66`, `\x67`, `\x68`, `\x69`, `\x6a`, `\x6b`, `\x6c`, `\x6d`, `\x6e`, `\x6f`, `\x70`, `\x71`, `\x72`, `\x73`, `\x74`, `\x75`, `\x76`, `\x77`, `\x78`, `\x79`, `\x7a`, `\x7b`, `\x7c`, `\x7d`, `\x7e`, `\x7f`, `\x80`, `\x81`, `\x82`, `\x83`, `\x84`, `\x85`, `\x86`, `\x87`, `\x88`, `\x89`, `\x8a`, `\x8b`, `\x8c`, `\x8d`, `\x8e`, `\x8f`, `\x90`, `\x91`, `\x92`, `\x93`, `\x94`, `\x95`, `\x96`, `\x97`, `\x98`, `\x99`, `\x9a`, `\x9b`, `\x9c`, `\x9d`, `\x9e`, `\x9f`, `\xa0`, `\xa1`, `\xa2`, `\xa3`, `\xa4`, `\xa5`, `\xa6`, `\xa7`, `\xa8`, `\xa9`, `\xaa`, `\xab`, `\xac`, `\xad`, `\xae`, `\xaf`, `\xb0`, `\xb1`, `\xb2`, `\xb3`, `\xb4`, `\xb5`, `\xb6`, `\xb7`, `\xb8`, `\xb9`, `\xba`, `\xbb`, `\xbc`, `\xbd`, `\xbe`, `\xbf`, `\xc0`, `\xc1`, `\xc2`, `\xc3`, `\xc4`, `\xc5`, `\xc6`, `\xc7`, `\xc8`, `\xc9`, `\xca`, `\xcb`, `\xcc`, `\xcd`, `\xce`, `\xcf`, `\xd0`, `\xd1`, `\xd2`, `\xd3`, `\xd4`, `\xd5`, `\xd6`, `\xd7`, `\xd8`, `\xd9`, `\xda`, `\xdb`, `\xdc`, `\xdd`, `\xde`, `\xdf`, `\xe0`, `\xe1`, `\xe2`, `\xe3`, `\xe4`, `\xe5`, `\xe6`, `\xe7`, `\xe8`, `\xe9`, `\xea`, `\xeb`, `\xec`, `\xed`, `\xee`, `\xef`, `\xf0`, `\xf1`, `\xf2`, `\xf3`, `\xf4`, `\xf5`, `\xf6`, `\xf7`, `\xf8`, `\xf9`, `\xfa`, `\xfb`, `\xfc`, `\xfd`, `\xfe`, `\xff`, `\x100`}
	switch blacklistSet {
	case nil:
		return strings.Join(allChars, "")
	default:
		goodChars := strings.Builder{}
		for i, hex := range allChars {
			for j := 0; j <= len(blacklistSet)-1; j++ {
				if allChars[i] == blacklistSet[j] {
					hex, blacklistSet[j] = "", ""
				}
			}
			goodChars.WriteString(hex)
		}
		return goodChars.String()
	}
}

//Reverse a string slice into string
func reverseSStoS(s []string) string {
	builder := strings.Builder{}
	for i := len(s) - 1; i >= 0; i-- {
		builder.WriteString(s[i])
	}
	return builder.String()
}

//Generate a padding string with choice of a single character
func paddingGen(c string, n int) string {
	builder := strings.Builder{}
	for i := 0; i <= n; i++ {
		builder.WriteString(c)
	}
	return builder.String()
}

//Generate a load of NOPs
func genNopsled(n int) string {
	var nopStr string = `\x90`
	builder := strings.Builder{}
	for i := 0; i <= n; i++ {
		builder.WriteString(nopStr)
	}
	return builder.String()
}

//
func reverseCharEndianness(s string) string {
	builder := strings.Builder{}
	var hexform string = `\x`
	if strings.Contains(s, hexform) {
		noSlashX := strings.ReplaceAll(s, hexform, "")
		reversed := reverseStrEndianness(noSlashX)
		reSlashX := insertSlashX(reversed)
		builder.WriteString(reSlashX)
	} else {
		err := fmt.Errorf("Error reverseCharEndianness %s does not contain %s!\n", s, hexform)
		fmt.Println(err.Error())
		return "why is this getting to this point"
	}
	return builder.String()
}

//
func insertSlashX(s string) string {
	builder := strings.Builder{}
	hexform := `\x`
	edit := strings.SplitAfter(s, "")
	editSize := len(edit)
	for i := 0; i < editSize-1; i += 2 {
		builder.WriteString(hexform + edit[i] + edit[i+1])
	}
	return builder.String()
}

func add0xStub(s string) string {
	addzeroX := "0x" + s
	return addzeroX
}

func remove0xStub(s string) string {
	stubless := strings.TrimPrefix(s, "0x")
	return stubless
}

func evenOrOdd(m int) bool {
	if m&1 == 0 {
		return true
	} else {
		return false
	}
}

//seperate into pairs of chars
func explodeToPairs(s string, n int) ([]string, int) {
	var pair string
	var resize int
	if evenOrOdd(n) {
		resize = n / 2
	} else {
		resize = (n + 1) / 2
	}
	pairSlice := make([]string, resize)
	for i, j := 0, 0; i <= n; i, j = i+2, j+1 {
		pair = string(s[i]) + string(s[i+1])
		pairSlice[j] = pair
	}
	return pairSlice, resize
}

//Used to reverse the endianess strings both number and char
func reverseStrEndianness(s string) string {
	builder := strings.Builder{}
	var editStr string
	if strings.Contains(s, "0x") {
		editStr = strings.TrimPrefix(s, "0x")
	} else {
		editStr = s
	}
	sizeS := len(editStr) - 1
	pieces, piecesSize := explodeToPairs(editStr, sizeS)
	for i := piecesSize - 1; i >= 0; i-- {
		builder.WriteString(pieces[i])
	}
	return builder.String()
}

func covertStrToByte(s string) []byte {
  result := byte[](s)
	return result
}
