package utils

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strings"
	"sync"
	"time"

	probing "github.com/prometheus-community/pro-bing"
	"github.com/tkytel/tripd/config"
	"github.com/tkytel/tripd/mantela"
)

var Peers []Peer
var LastMeasured time.Time

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

	sem := make(chan struct{}, 3)

	for _, v := range res.Providers {
		v := v
		wg.Add(1)

		go func() {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			// ignore the provider which is not mantela available
			if strings.Contains(v.Identifier, "XXX") || v.Mantela == "" {
				return
			}

			isMeasurable := true
			sipServer, err := mantela.FetchMantela(v.Mantela)
			if err != nil {
				log.Printf("Failed to fetch mantela (%v): %v", v.Identifier, err)
				isMeasurable = false
			}

			// if this provider is me, skipping
			// this happens when I am hoppable
			if sipServer.AboutMe.Identifier == res.AboutMe.Identifier {
				return
			}

			var rtt *float64
			var loss *float64
			var min *float64
			var max *float64
			var mdev *float64

			if len(sipServer.AboutMe.SipUri) == 0 {
				if sipServer.AboutMe.SipServer != "" {
					sipServer.AboutMe.SipUri = append(
						sipServer.AboutMe.SipUri,
						sipServer.AboutMe.SipServer,
					)
				} else {
					isMeasurable = false
				}
			}

			if isMeasurable {
				dest := ExtractPeerAddress(sipServer.AboutMe.SipUri[0])
				ping, err := PingPeer(dest)
				if err != nil {
					log.Printf("Ignored host %v: %v", dest, err)
					isMeasurable = false
					goto End
				}

				rttVal := float64(ping.AvgRtt.Microseconds()) * 0.001
				lossVal := ping.PacketLoss
				maxVal := float64(ping.MaxRtt.Microseconds()) * 0.001
				minVal := float64(ping.MinRtt.Microseconds()) * 0.001
				mdevVal := float64(ping.StdDevRtt.Microseconds()) * 0.001
				roundedRtt := math.Round(rttVal*1000) / 1000
				roundedLoss := math.Round(lossVal*1000) / 1000
				roundedMax := math.Round(maxVal*1000) / 1000
				roundedMin := math.Round(minVal*1000) / 1000
				roundedMdev := math.Round(mdevVal*1000) / 1000

				rtt = &roundedRtt
				loss = &roundedLoss
				max = &roundedMax
				min = &roundedMin
				mdev = &roundedMdev
			}

		End:
			peer := Peer{
				Measurable: isMeasurable,
				Identifier: v.Identifier,
				Rtt:        rtt,
				Loss:       loss,
				Min:        min,
				Max:        max,
				Mdev:       mdev,
			}

			mutex.Lock()
			p = append(p, peer)
			mutex.Unlock()
		}()
	}

	wg.Wait()
	Peers = p
	LastMeasured = time.Now()
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
	cfg := config.Get()
	pinger, err := probing.NewPinger(fqdn)
	if err != nil {
		return nil, err
	}
	pinger.SetPrivileged(true)

	msg := fmt.Sprintf("Pinging to %v", pinger.IPAddr())
	if fqdn != pinger.IPAddr().IP.String() {
		msg += fmt.Sprintf(" (%v)", fqdn)
	}
	log.Println(msg)

	pinger.Count = cfg.Ping.Count
	err = pinger.Run()
	if err != nil {
		return nil, err
	}

	return pinger.Statistics(), nil
}
