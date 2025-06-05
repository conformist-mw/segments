package models

import "testing"

func TestPaginator_GetOffset(t *testing.T) {
	tests := []struct {
		current int
		want    int
	}{
		{1, 0},
		{2, 10},
		{5, 40},
	}
	for _, tt := range tests {
		p := Paginator{Current: tt.current}
		got := p.GetOffset()
		if got != tt.want {
			t.Errorf("Current %d: expected %d, got %d", tt.current, tt.want, got)
		}
	}
}

func TestPaginator_GetLimit(t *testing.T) {
	p := Paginator{}
	if p.GetLimit() != 10 {
		t.Errorf("expected limit 10, got %d", p.GetLimit())
	}
}

func TestPaginator_GetTotalPages(t *testing.T) {
	tests := []struct {
		items int64
		want  int
	}{
		{0, 0},
		{1, 1},
		{10, 1},
		{11, 2},
		{99, 10},
	}
	for _, tt := range tests {
		p := Paginator{TotalItems: tt.items}
		got := p.GetTotalPages()
		if got != tt.want {
			t.Errorf("items %d: expected %d, got %d", tt.items, tt.want, got)
		}
	}
}
