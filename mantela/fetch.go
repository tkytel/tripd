package mantela

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func FetchMantela(mantelaUrl string) (Mantela, error) {
	resp, err := http.Get(mantelaUrl)
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
