package conv

import (
	"encoding/json"

	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/common/types"
)

func ByteToPerformers(row []byte) ([]types.Performer, error) {
	performers := make([]types.Performer, 0)

	err := json.Unmarshal(row, &performers)
	if err != nil {
		return nil, err
	}

	return performers, nil
}
