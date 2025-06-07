package utils

import (
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
				ping, err := PingPeer(ExtractPeerFQDN(sipServer.AboutMe.SipUri[0]))
				if err != nil {
					log.Println(err)
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

func ExtractPeerFQDN(sipUrl string) string {
	re := regexp.MustCompile(`sip:([a-zA-Z0-9.-]+):\d+`)
	match := re.FindStringSubmatch(sipUrl)

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

	log.Println("Pinging to", pinger.IPAddr())

	pinger.Count = 5
	err = pinger.Run()
	if err != nil {
		return nil, err
	}

	return pinger.Statistics(), nil
}
