package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"golang.org/x/net/ipv4"
)

// debut, OMIT
const multicastAddr = "224.0.0.1:6666"

func main() {
	go clientReader()
	conn, err := net.Dial("udp", multicastAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(conn, os.Stdin)
}

// milieu, OMIT
func clientReader() {
	// error testing suppressed to compact listing on slides
	conn, _ := net.ListenPacket("udp", multicastAddr) // listen on port
	defer conn.Close()
	p := ipv4.NewPacketConn(conn) // convert to ipv4 packetConn
	addr, _ := net.ResolveUDPAddr("udp", multicastAddr)
	p.JoinGroup(nil, addr) // listen on ip multicast
	buf := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFrom(buf) // n, _, addr, err := p.ReadFrom(buf)
		if err != nil {
			log.Fatal(err)
		}
		// extract text without CR/LF
		var text string
		fmt.Sscanf(string(buf[0:n]), "%s", &text)
		fmt.Printf("%s from %v\n", text, addr)
	}
}

// fin, OMIT
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
