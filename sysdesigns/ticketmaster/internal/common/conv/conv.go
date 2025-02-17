package conv

import (
	"encoding/json"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/common/types"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/entities"
)

func ByteToPerformers(row []byte) ([]types.Performer, error) {
	performers := make([]types.Performer, 0)

	err := json.Unmarshal(row, &performers)
	if err != nil {
		return nil, err
	}

	return performers, nil
}

func ByteToTickets(row []byte) ([]types.Ticket, error) {
	tickets := make([]types.Ticket, 0)

	err := json.Unmarshal(row, &tickets)
	if err != nil {
		return nil, err
	}

	return tickets, nil
}

func ConvertEventSearchParams(req types.SearchEvent) entities.SearchEventsParams {
	if req.Page < 1 {
		req.Page = 1
	}
	offset := (req.Page - 1) * req.PageSize

	return entities.SearchEventsParams{
		Column1: pgtype.Text{
			String: req.Keyword,
			Valid:  req.Keyword != "",
		},
		EventDate: pgtype.Timestamp{
			Time:  req.Start,
			Valid: !req.Start.IsZero(),
		},
		EventDate_2: pgtype.Timestamp{
			Time:  req.End,
			Valid: !req.End.IsZero(),
		},
		Limit:  req.PageSize,
		Offset: offset,
	}
}
