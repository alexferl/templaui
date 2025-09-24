package pagination

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestPagination(t *testing.T) {
	tests := []struct {
		name   string
		props  PaginationProps
		expect string
	}{
		{
			name:   "Default",
			props:  PaginationProps{},
			expect: `<nav class="pagination" role="navigation" aria-label="pagination"></nav>`,
		},
		{
			name:   "With ID and custom classes",
			props:  PaginationProps{ID: "pagination1", Class: []string{"custom-pagination"}},
			expect: `<nav id="pagination1" class="pagination custom-pagination" role="navigation" aria-label="pagination"></nav>`,
		},
		{
			name:   "With size",
			props:  PaginationProps{Size: IsLarge},
			expect: `<nav class="pagination is-large" role="navigation" aria-label="pagination"></nav>`,
		},
		{
			name:   "With alignment",
			props:  PaginationProps{Alignment: IsCentered},
			expect: `<nav class="pagination is-centered" role="navigation" aria-label="pagination"></nav>`,
		},
		{
			name:   "With style modifiers",
			props:  PaginationProps{IsDisabled: true, IsRounded: true},
			expect: `<nav class="pagination is-disabled is-rounded" role="navigation" aria-label="pagination"></nav>`,
		},
		{
			name:   "With states",
			props:  PaginationProps{IsActive: true, IsCurrent: true},
			expect: `<nav class="pagination is-active is-current" role="navigation" aria-label="pagination"></nav>`,
		},
		{
			name: "All fields combined",
			props: PaginationProps{
				ID:         "main-pagination",
				Class:      []string{"foo", "bar"},
				Size:       IsMedium,
				Alignment:  IsRight,
				IsRounded:  true,
				IsActive:   true,
				Attributes: templ.Attributes{"data-role": "pagination"},
			},
			expect: `<nav id="main-pagination" class="pagination is-medium is-right is-rounded is-active foo bar" role="navigation" aria-label="pagination" data-role="pagination"></nav>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Pagination(tt.props).Render(context.Background(), &buf)
			if err != nil {
				t.Fatalf("render failed: %v", err)
			}
			got := strings.TrimSpace(buf.String())
			if got != tt.expect {
				t.Errorf("expected:\n%s\ngot:\n%s", tt.expect, got)
			}
		})
	}
}

func TestPaginationPrevious(t *testing.T) {
	tests := []struct {
		name   string
		props  PaginationPreviousProps
		expect string
	}{
		{
			name:   "Default",
			props:  PaginationPreviousProps{},
			expect: `<a class="pagination-previous"></a>`,
		},
		{
			name:   "With ID and custom classes",
			props:  PaginationPreviousProps{ID: "prev1", Class: []string{"custom-prev"}},
			expect: `<a id="prev1" class="pagination-previous custom-prev"></a>`,
		},
		{
			name:   "With disabled state",
			props:  PaginationPreviousProps{IsDisabled: true},
			expect: `<a class="pagination-previous" disabled></a>`,
		},
		{
			name: "All fields combined",
			props: PaginationPreviousProps{
				ID:         "main-prev",
				Class:      []string{"custom"},
				IsDisabled: true,
				Attributes: templ.Attributes{"data-action": "prev"},
			},
			expect: `<a id="main-prev" class="pagination-previous custom" disabled data-action="prev"></a>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := PaginationPrevious(tt.props).Render(context.Background(), &buf)
			if err != nil {
				t.Fatalf("render failed: %v", err)
			}
			got := strings.TrimSpace(buf.String())
			if got != tt.expect {
				t.Errorf("expected:\n%s\ngot:\n%s", tt.expect, got)
			}
		})
	}
}

func TestPaginationNext(t *testing.T) {
	tests := []struct {
		name   string
		props  PaginationNextProps
		expect string
	}{
		{
			name:   "Default",
			props:  PaginationNextProps{},
			expect: `<a class="pagination-next"></a>`,
		},
		{
			name:   "With ID and custom classes",
			props:  PaginationNextProps{ID: "next1", Class: []string{"custom-next"}},
			expect: `<a id="next1" class="pagination-next custom-next"></a>`,
		},
		{
			name:   "With disabled state",
			props:  PaginationNextProps{IsDisabled: true},
			expect: `<a class="pagination-next" disabled></a>`,
		},
		{
			name: "All fields combined",
			props: PaginationNextProps{
				ID:         "main-next",
				Class:      []string{"custom"},
				IsDisabled: true,
				Attributes: templ.Attributes{"data-action": "next"},
			},
			expect: `<a id="main-next" class="pagination-next custom" disabled data-action="next"></a>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := PaginationNext(tt.props).Render(context.Background(), &buf)
			if err != nil {
				t.Fatalf("render failed: %v", err)
			}
			got := strings.TrimSpace(buf.String())
			if got != tt.expect {
				t.Errorf("expected:\n%s\ngot:\n%s", tt.expect, got)
			}
		})
	}
}

func TestPaginationList(t *testing.T) {
	tests := []struct {
		name   string
		props  PaginationListProps
		expect string
	}{
		{
			name:   "Default",
			props:  PaginationListProps{},
			expect: `<ul class="pagination-list"></ul>`,
		},
		{
			name:   "With ID and custom classes",
			props:  PaginationListProps{ID: "list1", Class: []string{"custom-list"}},
			expect: `<ul id="list1" class="pagination-list custom-list"></ul>`,
		},
		{
			name: "With attributes",
			props: PaginationListProps{
				ID:         "page-list",
				Attributes: templ.Attributes{"data-pages": "10"},
			},
			expect: `<ul id="page-list" class="pagination-list" data-pages="10"></ul>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := PaginationList(tt.props).Render(context.Background(), &buf)
			if err != nil {
				t.Fatalf("render failed: %v", err)
			}
			got := strings.TrimSpace(buf.String())
			if got != tt.expect {
				t.Errorf("expected:\n%s\ngot:\n%s", tt.expect, got)
			}
		})
	}
}

func TestPaginationLink(t *testing.T) {
	tests := []struct {
		name   string
		props  PaginationLinkProps
		expect string
	}{
		{
			name:   "Default",
			props:  PaginationLinkProps{},
			expect: `<a class="pagination-link"></a>`,
		},
		{
			name:   "With ID and custom classes",
			props:  PaginationLinkProps{ID: "link1", Class: []string{"custom-link"}},
			expect: `<a id="link1" class="pagination-link custom-link"></a>`,
		},
		{
			name:   "With current state",
			props:  PaginationLinkProps{IsCurrent: true},
			expect: `<a class="pagination-link is-current" aria-current="page"></a>`,
		},
		{
			name:   "With active state",
			props:  PaginationLinkProps{IsActive: true},
			expect: `<a class="pagination-link is-active"></a>`,
		},
		{
			name:   "With multiple states",
			props:  PaginationLinkProps{IsActive: true, IsFocused: true, IsSelected: true},
			expect: `<a class="pagination-link is-active is-focused is-selected"></a>`,
		},
		{
			name: "All fields combined",
			props: PaginationLinkProps{
				ID:         "page-link",
				Class:      []string{"custom"},
				IsActive:   true,
				IsCurrent:  true,
				IsSelected: true,
				Attributes: templ.Attributes{"data-page": "5"},
			},
			expect: `<a id="page-link" class="pagination-link is-active is-current is-selected custom" aria-current="page" data-page="5"></a>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := PaginationLink(tt.props).Render(context.Background(), &buf)
			if err != nil {
				t.Fatalf("render failed: %v", err)
			}
			got := strings.TrimSpace(buf.String())
			if got != tt.expect {
				t.Errorf("expected:\n%s\ngot:\n%s", tt.expect, got)
			}
		})
	}
}

func TestPaginationEllipsis(t *testing.T) {
	tests := []struct {
		name   string
		props  PaginationEllipsisProps
		expect string
	}{
		{
			name:   "Default",
			props:  PaginationEllipsisProps{},
			expect: `<span class="pagination-ellipsis"></span>`,
		},
		{
			name:   "With ID and custom classes",
			props:  PaginationEllipsisProps{ID: "ellipsis1", Class: []string{"custom-ellipsis"}},
			expect: `<span id="ellipsis1" class="pagination-ellipsis custom-ellipsis"></span>`,
		},
		{
			name: "With attributes",
			props: PaginationEllipsisProps{
				ID:         "separator",
				Attributes: templ.Attributes{"data-separator": "true"},
			},
			expect: `<span id="separator" class="pagination-ellipsis" data-separator="true"></span>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := PaginationEllipsis(tt.props).Render(context.Background(), &buf)
			if err != nil {
				t.Fatalf("render failed: %v", err)
			}
			got := strings.TrimSpace(buf.String())
			if got != tt.expect {
				t.Errorf("expected:\n%s\ngot:\n%s", tt.expect, got)
			}
		})
	}
}
