package utils

import (
	"log"
	"regexp"
	"strings"
	"sync"

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

			isMeasurable := false
			sipServer, _ := mantela.FetchMantela(v.Mantela)

			if sipServer.AboutMe.SipServer != "" {
				isMeasurable = true
			}

			peer := Peer{
				Measurable: isMeasurable,
				Identifier: v.Identifier,
				Rtt:        nil,
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
