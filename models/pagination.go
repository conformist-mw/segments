package models

import (
	"math"
	"net/url"
	"strconv"
)

type Paginator struct {
	TotalItems int64
	Current    int
	Next       int
	Previous   int
	SearchForm SearchForm
}

func (p Paginator) GetOffset() int {
	return 10 * (p.Current - 1)
}

func (p Paginator) GetLimit() int {
	return 10
}

func (p Paginator) GetTotalPages() int {
	return int(math.Ceil(float64(p.TotalItems) / float64(p.GetLimit())))
}

func (p Paginator) HasPrevious() bool {
	return p.Current > 1
}

func (p Paginator) GetPrevious() int {
	if p.HasPrevious() {
		return p.Current - 1
	}
	return p.Current
}

func (p Paginator) HasNext() bool {
	return p.Current < p.GetTotalPages()
}

func (p Paginator) GetNext() int {
	if p.HasNext() {
		return p.Current + 1
	}
	return p.Current
}

func (p Paginator) Start() int {
	start := p.Current - 6
	if start < 1 {
		start = 1
	}
	return start
}

func (p Paginator) End() int {
	end := p.Current + 6
	if end > p.GetTotalPages() {
		end = p.GetTotalPages()
	}
	return end
}

func (p Paginator) PageUrl() string {
	params := make(url.Values)
	params.Set("color", p.SearchForm.Color)
	params.Set("color_type", p.SearchForm.ColorType)
	params.Set("deleted", p.SearchForm.Deleted)
	params.Set("height", strconv.Itoa(p.SearchForm.Height))
	params.Set("width", strconv.Itoa(p.SearchForm.Width))
	params.Set("page", "__page_number__")
	return "?" + params.Encode()
}
