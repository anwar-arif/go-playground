package dns_msg

import (
	"encoding/hex"
	"fmt"
	"github.com/miekg/dns"
	"log"
)

func Run() {
	// Hex-encoded DNS message
	//hexMessage := "a01d81800001000100000000076578616d706c6503636f6d0000010001c00c0001000100001bbc00045db8d822"

	var hexMessage string
	_, err := fmt.Scanf("%s", &hexMessage)
	if err != nil {
		fmt.Println("error reading input")
		return
	}

	// Decode the hex-encoded string into bytes
	bytesMessage, err := hex.DecodeString(hexMessage)
	if err != nil {
		log.Fatal("Error decoding hex string:", err)
	}

	// Parse the DNS message
	msg := new(dns.Msg)
	if err := msg.Unpack(bytesMessage); err != nil {
		log.Fatal("Error parsing DNS message:", err)
	}
	dnsMsg := string(msg.String())
	prefix := ";; ->>HEADER<<-"
	format := prefix + dnsMsg[2:]
	fmt.Println(format)

	//// Print the human-readable representation of the DNS message
	//fmt.Println("DNS Header:")
	//fmt.Printf("ID: %d\n", msg.Id)
	//fmt.Printf("QR: %s\n", map[bool]string{true: "Response", false: "Query"}[msg.Response])
	//fmt.Printf("Opcode: %d\n", msg.Opcode)
	//fmt.Printf("Msg: %s\n", msg.String())
	//fmt.Printf("Header: %s\n", msg.MsgHdr.String())
	//fmt.Printf("Opcode str: %s\n", msg.MsgHdr.String())
	//fmt.Printf("AA: %t\n", msg.Authoritative)
	//fmt.Printf("TC: %t\n", msg.Truncated)
	//fmt.Printf("RD: %t\n", msg.RecursionDesired)
	//fmt.Printf("RA: %t\n", msg.RecursionAvailable)
	////fmt.Printf("Z: %d\n", msg.Z)
	//fmt.Printf("RCODE: %d\n", msg.Rcode)
	//fmt.Printf("QDCOUNT: %d\n", len(msg.Question))
	//fmt.Printf("ANCOUNT: %d\n", len(msg.Answer))
	//fmt.Printf("NSCOUNT: %d\n", len(msg.Ns))
	//fmt.Printf("ARCOUNT: %d\n", len(msg.Extra))
	//
	//for _, q := range msg.Question {
	//	fmt.Println("\nQuestion Section:")
	//	fmt.Printf("QNAME: %s\n", q.Name)
	//	fmt.Printf("QTYPE: %d\n", q.Qtype)
	//	fmt.Printf("QCLASS: %d\n", q.Qclass)
	//}
	//
	//for _, rr := range msg.Answer {
	//	fmt.Println("\nAnswer Section:")
	//	fmt.Printf("NAME: %s\n", rr.Header().Name)
	//	fmt.Printf("TYPE: %d\n", rr.Header().Rrtype)
	//	fmt.Printf("RCLASS: %d\n", rr.Header().Class)
	//	fmt.Printf("TTL: %d\n", rr.Header().Ttl)
	//	fmt.Printf("RDLENGTH: %d\n", rr.Header().Rdlength)
	//
	//	switch rr.Header().Rrtype {
	//	case dns.TypeA:
	//		rrA, ok := rr.(*dns.A)
	//		if ok {
	//			fmt.Printf("Address: %s\n", rrA.A)
	//		}
	//	case dns.TypeAAAA:
	//		rrAAAA, ok := rr.(*dns.AAAA)
	//		if ok {
	//			fmt.Printf("IPv6 Address: %s\n", rrAAAA.AAAA)
	//		}
	//	case dns.TypeCNAME:
	//		rrCNAME, ok := rr.(*dns.CNAME)
	//		if ok {
	//			fmt.Printf("CNAME: %s\n", rrCNAME.Target)
	//		}
	//	}
	//}
	//
	//for _, auth := range msg.Ns {
	//	fmt.Println("\nAuthority Section:")
	//	fmt.Printf("NAME: %s\n", auth.Header().Name)
	//	fmt.Printf("TYPE: %d\n", auth.Header().Rrtype)
	//	fmt.Printf("RCLASS: %d\n", auth.Header().Class)
	//	fmt.Printf("TTL: %d\n", auth.Header().Ttl)
	//	fmt.Printf("RDLENGTH: %d\n", auth.Header().Rdlength)
	//}
	//
	//for _, extra := range msg.Extra {
	//	fmt.Println("\nAdditional Section:")
	//	fmt.Printf("NAME: %s\n", extra.Header().Name)
	//	fmt.Printf("TYPE: %d\n", extra.Header().Rrtype)
	//	fmt.Printf("RCLASS: %d\n", extra.Header().Class)
	//	fmt.Printf("TTL: %d\n", extra.Header().Ttl)
	//	fmt.Printf("RDLENGTH: %d\n", extra.Header().Rdlength)
	//}
}
