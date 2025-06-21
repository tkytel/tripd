package utils

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func SendSipOptions(sipServer string, port string) (res bool, error error) {
	addr := net.JoinHostPort(sipServer, port)
	raddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return false, err
	}

	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		return false, err
	}
	defer conn.Close()

	msg := fmt.Sprintf(
		"OPTIONS sip:%s SIP/2.0\r\n"+
			"Via: SIP/2.0/UDP client.local;branch=z9hG4bK123456\r\n"+
			"Max-Forwards: 70\r\n"+
			"From: <sip:monitor@%s>;tag=12345\r\n"+
			"To: <sip:monitor@%s>\r\n"+
			"Call-ID: %s\r\n"+
			"CSeq: 0 OPTIONS\r\n"+
			"Contact: <sip:monitor@%s>\r\n"+
			"User-Agent: tripd\r\n"+
			"Content-Length: 0\r\n\r\n",
		sipServer, sipServer, sipServer, RandStringRunes(20), sipServer,
	)

	_, err = conn.Write([]byte(msg))
	if err != nil {
		return false, err
	}

	conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	buf := make([]byte, 2048)
	n, _, err := conn.ReadFromUDP(buf)
	if err != nil {
		return false, err
	}

	response := string(buf[:n])

	if strings.Contains(response, "SIP/2.0 200 OK") {
		log.Println("Server", sipServer, "is responding to SIP OPTIONS")
		return true, nil
	}
	if strings.Contains(response, "SIP/2.0 401 Unauthorized") {
		log.Println("Server", sipServer, "is responding to SIP OPTIONS")
		return true, nil
	}

	fmt.Println("⚠️ Unexpected response:", response)
	return false, nil
}

func ParseSipURI(uri string) (host string, port string, err error) {
	if !strings.HasPrefix(uri, "sip:") {
		return "", "", fmt.Errorf("invalid sip uri; must start with sip: ")
	}
	uri = strings.TrimPrefix(uri, "sip:")

	semicolonIndex := strings.Index(uri, ";")
	if semicolonIndex != -1 {
		uri = uri[:semicolonIndex]
	}

	atIndex := strings.LastIndex(uri, "@")
	var hostPort string
	if atIndex != -1 {
		hostPort = uri[atIndex+1:]
	} else {
		hostPort = uri
	}

	if strings.Contains(hostPort, ":") {
		host, port, err = net.SplitHostPort(hostPort)
		if err != nil {
			return "", "", fmt.Errorf("invalid sip uri: %v", err)
		}
	} else {
		host = hostPort
		port = "5060"
	}

	return host, port, nil
}
