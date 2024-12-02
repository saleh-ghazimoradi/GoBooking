package service_models

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type PaginationFeedQuery struct {
	Limit  int       `json:"limit" validate:"gte=1,lte=20"`
	Offset int       `json:"offset" validate:"gte=0"`
	Sort   string    `json:"sort" validate:"oneof=asc desc"`
	Search string    `json:"search" validate:"max=100"`
	Since  time.Time `json:"since"`
	Until  time.Time `json:"until"`
}

func (fq PaginationFeedQuery) Parse(r *http.Request) (PaginationFeedQuery, error) {
	qs := r.URL.Query()

	limit := qs.Get("limit")
	if limit != "" {
		l, err := strconv.Atoi(limit)
		if err != nil {
			return fq, nil
		}
		fq.Limit = l
	}

	offset := qs.Get("offset")
	if offset != "" {
		l, err := strconv.Atoi(offset)
		if err != nil {
			return fq, nil
		}

		fq.Offset = l
	}

	sort := qs.Get("sort")
	if sort != "" {
		fq.Sort = sort
	}

	search := qs.Get("search")
	if search != "" {
		fq.Search = search
	}

	since := qs.Get("since")
	t, err := parseTime(since)
	if err != nil {
		return fq, err
	}
	fq.Since = t

	until := qs.Get("until")
	t, err = parseTime(until)
	if err != nil {
		return fq, err
	}
	fq.Until = t

	return fq, nil
}

func parseTime(value string) (time.Time, error) {
	if value == "" {
		return time.Time{}, nil
	}
	t, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid time format: %v", err)
	}
	return t, nil
}
