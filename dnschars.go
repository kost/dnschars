package dnschars

import (
	"fmt"
	"strconv"
)

var allowedchrs = []byte("0123456789abcdefghijklmnopqrstuvwxy")
var specialchr = "z"

func DnsEncode (buffer []byte) (r string) {
	str:=""

	for _, element := range buffer {
		found := false
		for _, e := range allowedchrs {
			if e == element {
				found=true
			}
		}
		if found {
			str = str + string(element)
		} else {
			hex := fmt.Sprintf("%x", int(element))
			str = str + specialchr + hex
		}
	}

	return str
}

func DnsEncodeSize (buffer []byte, limit int) (r string) {
	str:=""

	for _, element := range buffer {
		found := false
		for _, e := range allowedchrs {
			if e == element {
				found=true
			}
		}
		if found {
			if (len(str)+1>=limit) {
				return str
			}
			str = str + string(element)
		} else {
			if (len(str)+3>=limit) {
				return str
			}
			hex := fmt.Sprintf("%x", int(element))
			str = str + specialchr + hex
		}
	}
	return str
}

func DnsDecode (buffer []byte) (r string) {
	str:=""

	index :=0
	for index < len(buffer) {
		element:=buffer[index]
		if string(element) == specialchr {
			hexstr:=string(buffer[index+1])+string(buffer[index+2])
			result, _:=strconv.ParseInt(hexstr, 16, 0)
			strch := rune(result)
			index=index+2
			str = str + string(strch)
		} else {
			str = str + string(element)
		}
		index=index+1
	}
	return str
}

