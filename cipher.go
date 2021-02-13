package main

import (
	"fmt"
	"log"
	"strings"
)

// https://ksnctf.sweetduet.info/problem/2

const text = "EBG KVVV vf n fvzcyr yrggre fhofgvghgvba pvcure gung ercynprf n yrggre jvgu gur yrggre KVVV yrggref nsgre vg va gur nycunorg. EBG KVVV vf na rknzcyr bs gur Pnrfne pvcure, qrirybcrq va napvrag Ebzr. Synt vf SYNTFjmtkOWFNZdjkkNH. Vafreg na haqrefpber vzzrqvngryl nsgre SYNT."

var (
	upperByte = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lowerByte = []byte("abcdefghijklmnopqrstuvwxyz")
	max       = len(upperByte)
)

func main() {
	for i := 0; i < max; i++ {
		var s = shift(text, i)

		if strings.Contains(s, "FLAG") {
			fmt.Println(s)
		}
	}
}

func shift(text string, num int) string {
	var retByte = make([]byte, len(text))

	for i, tb := range []byte(text) {
		if strings.Contains(" ,.", string(tb)) {
			retByte[i] = tb
			continue
		}

		b := shiftByte(upperByte, tb, num)
		if b != 0 {
			retByte[i] = b
			continue
		}

		b = shiftByte(lowerByte, tb, num)
		if b != 0 {
			retByte[i] = b
			continue
		}

		log.Fatalf("unexpected character: character: %s, text: %s", string(tb), text)
	}

	return string(retByte)
}

func shiftByte(bytes []byte, b byte, num int) byte {
	for j, byte := range bytes {
		if byte == b {
			if j < num {
				return bytes[max+(j-num)]
			}
			return bytes[j-num]
		}
	}
	return 0
}