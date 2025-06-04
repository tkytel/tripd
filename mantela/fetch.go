package mantela

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/tkytel/tripd/config"
)

func FetchMantela() (Mantela, error) {
	c := config.Get()
	resp, err := http.Get(c.Mantela.Url)
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
