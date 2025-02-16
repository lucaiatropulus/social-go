package store

import (
	"net/http"
	"strconv"
	"strings"
	"time"
)

type PaginatedFeedQuery struct {
	Limit  int      `json:"limit"`
	Offset int      `json:"offset"`
	Sort   string   `json:"sort"`
	Tags   []string `json:"tags"`
	Search string   `json:"search"`
	Since  string   `json:"since"`
	Until  string   `json:"until"`
}

func (p *PaginatedFeedQuery) IsValid() bool {
	isLimitValid := p.Limit >= 1 && p.Limit <= 20
	isOffsetValid := p.Offset >= 0
	isSortValid := p.Sort == "ASC" || p.Sort == "DESC"

	return isLimitValid && isOffsetValid && isSortValid
}

func (p PaginatedFeedQuery) Parse(r *http.Request) (PaginatedFeedQuery, error) {
	queryParam := r.URL.Query()

	limit := queryParam.Get("limit")

	if limit != "" {
		l, err := strconv.Atoi(limit)

		if err != nil {
			return p, nil
		}

		p.Limit = l
	}

	offset := queryParam.Get("offset")

	if offset != "" {
		o, err := strconv.Atoi(offset)

		if err != nil {
			return p, nil
		}

		p.Offset = o
	}

	sort := queryParam.Get("sort")

	if sort != "" {
		p.Sort = sort
	}

	tags := queryParam.Get("tags")

	if tags != "" {
		p.Tags = strings.Split(tags, ",")
	}

	search := queryParam.Get("search")

	if search != "" {
		p.Search = search
	}

	since := queryParam.Get("since")

	if since != "" {
		p.Since = parseTime(since)
	}

	until := queryParam.Get("until")

	if until != "" {
		p.Until = parseTime(until)
	}

	return p, nil
}

func parseTime(s string) string {
	t, err := time.Parse(time.DateTime, s)

	if err != nil {
		return ""
	}

	return t.Format(time.DateTime)
}
