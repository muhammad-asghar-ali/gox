package handlers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"

	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/common/types"
)

func parseSearchEventParams(c fiber.Ctx) types.SearchEvent {
	req := types.SearchEvent{
		Page:     1,
		PageSize: 10,
	}

	req.Keyword = c.Query("keyword")

	if page, err := strconv.Atoi(c.Query("page", "1")); err == nil {
		req.Page = int32(page)
	}

	if pageSize, err := strconv.Atoi(c.Query("page_size", "10")); err == nil {
		req.PageSize = int32(pageSize)
	}

	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	const dateFormat = "2006-01-02"
	var hasStart, hasEnd bool

	if startDate != "" {
		if parsedStart, err := time.Parse(dateFormat, startDate); err == nil {
			req.Start = parsedStart
			hasStart = true
		}
	}

	if endDate != "" {
		if parsedEnd, err := time.Parse(dateFormat, endDate); err == nil {
			req.End = parsedEnd
			hasEnd = true
		}
	}

	if !hasStart {
		req.Start = time.Now().AddDate(0, -1, 0)
	}

	if !hasEnd {
		req.End = time.Now().AddDate(0, 1, 0)
	}

	return req
}
