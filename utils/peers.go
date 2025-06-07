package utils

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"sync"

	probing "github.com/prometheus-community/pro-bing"
	"github.com/tkytel/tripd/config"
	"github.com/tkytel/tripd/mantela"
)

var Peers []Peer

func RetrievePeers() {
	cfg := config.Get()

	res, err := mantela.FetchMantela(cfg.Mantela.Url)
	if err != nil {
		log.Println("Failed to fetch mantela:", err)
	}
	var (
		p     = make([]Peer, 0, len(res.Providers))
		mutex sync.Mutex
		wg    sync.WaitGroup
	)

	for _, v := range res.Providers {
		v := v
		wg.Add(1)

		go func() {
			defer wg.Done()

			if strings.Contains(v.Identifier, "XXX") {
				return
			}

			isMeasurable := true
			sipServer, _ := mantela.FetchMantela(v.Mantela)

			if sipServer.AboutMe.Identifier == res.AboutMe.Identifier {
				return
			}

			var rtt *int64
			var loss *float64

			if len(sipServer.AboutMe.SipUri) == 0 {
				isMeasurable = false
			}

			if isMeasurable {
				dest := ExtractPeerAddress(sipServer.AboutMe.SipUri[0])
				ping, err := PingPeer(dest)
				if err != nil {
					log.Printf("Ignored host %v: %v", dest, err)
					isMeasurable = false
					goto End
				}

				rttVal := ping.AvgRtt.Milliseconds()
				lossVal := ping.PacketLoss
				rtt = &rttVal
				loss = &lossVal
			}

		End:
			peer := Peer{
				Measurable: isMeasurable,
				Identifier: v.Identifier,
				Rtt:        rtt,
				Loss:       loss,
			}

			mutex.Lock()
			p = append(p, peer)
			mutex.Unlock()
		}()
	}

	wg.Wait()
	Peers = p

	log.Println("Updated peers with", len(Peers), "entries")
}

func ExtractPeerAddress(sipUri string) string {
	re := regexp.MustCompile(`sip:(?:\/\/)?(?:[^@]+@)?([^:;]+)`)
	match := re.FindStringSubmatch(sipUri)
	if len(match) >= 2 {
		return match[1]
	}

	re = regexp.MustCompile(`(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})`)
	match = re.FindStringSubmatch(sipUri)
	if len(match) > 1 {
		return match[1]
	}

	return ""
}

func PingPeer(fqdn string) (*probing.Statistics, error) {
	pinger, err := probing.NewPinger(fqdn)
	if err != nil {
		return nil, err
	}

	msg := fmt.Sprintf("Pinging to %v", pinger.IPAddr())
	if fqdn != pinger.IPAddr().IP.String() {
		msg += fmt.Sprintf(" (%v)", fqdn)
	}
	log.Println(msg)

	pinger.Count = 20
	err = pinger.Run()
	if err != nil {
		return nil, err
	}

	return pinger.Statistics(), nil
}
