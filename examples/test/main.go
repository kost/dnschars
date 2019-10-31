package main

import (
	"github.com/kost/dnschars"
	"log"
)

func main() {
	str:="abcde#gzxxxxxnnfjdjkjksjdkjkjksjk2837297391793897937897498zzzzzzzzzzzzzz00010187823782"
	log.Printf("String %s, %d\n", str, len(str))
	y:=dnschars.DnsEncodeSize([]byte(str), 63)
	log.Printf("Encoded %s, %d\n", y, len(y))
	z:=dnschars.DnsDecode([]byte(y))
	log.Printf("Decoded %s, %d\n", z, len(z))
}
