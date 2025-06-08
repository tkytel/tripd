package mantela

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

func FetchMantela(mantelaUri string) (Mantela, error) {
	// if mantelaUrl is data uri, parse it directly
	if strings.Contains(mantelaUri, "data") {
		b := strings.TrimPrefix(mantelaUri, `data:application/json;base64,`)
		j, err := base64.StdEncoding.DecodeString(b)
		if err != nil {
			return Mantela{}, err
		}

		var m Mantela
		if err := json.Unmarshal(j, &m); err != nil {
			return Mantela{}, err
		}

		return m, nil
	}

	// if mantelaUri is truly uri, retrieve / parse it
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get(mantelaUri)
	if err != nil {
		return Mantela{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return Mantela{}, errors.New("mantela.json not found")
	}
	if resp.StatusCode != http.StatusOK {
		return Mantela{}, errors.New("failed to fetch mantela.json")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Mantela{}, err
	}

	var m Mantela
	if err := json.Unmarshal(body, &m); err != nil {
		return Mantela{}, err
	}

	return m, nil
}
