package title

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestTitle(t *testing.T) {
	tests := []struct {
		name   string
		props  TitleProps
		expect string
	}{
		{
			name:   "Default (h1)",
			props:  TitleProps{},
			expect: `<h1 class="title"></h1>`,
		},
		{
			name:   "With ID and custom classes",
			props:  TitleProps{ID: "title1", Class: []string{"custom-title"}},
			expect: `<h1 id="title1" class="title custom-title"></h1>`,
		},
		{
			name:   "With size 1",
			props:  TitleProps{Size: Is1},
			expect: `<h1 class="title is-1"></h1>`,
		},
		{
			name:   "With size 2",
			props:  TitleProps{Size: Is2},
			expect: `<h1 class="title is-2"></h1>`,
		},
		{
			name:   "With size 3",
			props:  TitleProps{Size: Is3},
			expect: `<h1 class="title is-3"></h1>`,
		},
		{
			name:   "With size 4",
			props:  TitleProps{Size: Is4},
			expect: `<h1 class="title is-4"></h1>`,
		},
		{
			name:   "With size 5",
			props:  TitleProps{Size: Is5},
			expect: `<h1 class="title is-5"></h1>`,
		},
		{
			name:   "With size 6",
			props:  TitleProps{Size: Is6},
			expect: `<h1 class="title is-6"></h1>`,
		},
		{
			name:   "With size 7",
			props:  TitleProps{Size: Is7},
			expect: `<h1 class="title is-7"></h1>`,
		},
		{
			name:   "With spaced",
			props:  TitleProps{IsSpaced: true},
			expect: `<h1 class="title is-spaced"></h1>`,
		},
		{
			name:   "With h2 level",
			props:  TitleProps{Level: 2},
			expect: `<h2 class="title"></h2>`,
		},
		{
			name:   "With h3 level",
			props:  TitleProps{Level: 3},
			expect: `<h3 class="title"></h3>`,
		},
		{
			name:   "With h4 level",
			props:  TitleProps{Level: 4},
			expect: `<h4 class="title"></h4>`,
		},
		{
			name:   "With h5 level",
			props:  TitleProps{Level: 5},
			expect: `<h5 class="title"></h5>`,
		},
		{
			name:   "With h6 level",
			props:  TitleProps{Level: 6},
			expect: `<h6 class="title"></h6>`,
		},
		{
			name:   "Level too high (defaults to h6)",
			props:  TitleProps{Level: 10},
			expect: `<h6 class="title"></h6>`,
		},
		{
			name:   "Level too low (defaults to h1)",
			props:  TitleProps{Level: -1},
			expect: `<h1 class="title"></h1>`,
		},
		{
			name:   "With size and level",
			props:  TitleProps{Size: Is2, Level: 3},
			expect: `<h3 class="title is-2"></h3>`,
		},
		{
			name: "All fields combined",
			props: TitleProps{
				ID:         "main-title",
				Class:      []string{"foo", "bar"},
				Size:       Is1,
				Level:      2,
				IsSpaced:   true,
				Attributes: templ.Attributes{"data-testid": "page-title"},
			},
			expect: `<h2 id="main-title" class="title is-1 is-spaced foo bar" data-testid="page-title"></h2>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Title(tt.props).Render(context.Background(), &buf)
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

func TestSubtitle(t *testing.T) {
	tests := []struct {
		name   string
		props  SubtitleProps
		expect string
	}{
		{
			name:   "Default (h2)",
			props:  SubtitleProps{},
			expect: `<h2 class="subtitle"></h2>`,
		},
		{
			name:   "With ID and custom classes",
			props:  SubtitleProps{ID: "subtitle1", Class: []string{"custom-subtitle"}},
			expect: `<h2 id="subtitle1" class="subtitle custom-subtitle"></h2>`,
		},
		{
			name:   "With size 1",
			props:  SubtitleProps{Size: Is1},
			expect: `<h2 class="subtitle is-1"></h2>`,
		},
		{
			name:   "With size 2",
			props:  SubtitleProps{Size: Is2},
			expect: `<h2 class="subtitle is-2"></h2>`,
		},
		{
			name:   "With size 3",
			props:  SubtitleProps{Size: Is3},
			expect: `<h2 class="subtitle is-3"></h2>`,
		},
		{
			name:   "With size 4",
			props:  SubtitleProps{Size: Is4},
			expect: `<h2 class="subtitle is-4"></h2>`,
		},
		{
			name:   "With size 5",
			props:  SubtitleProps{Size: Is5},
			expect: `<h2 class="subtitle is-5"></h2>`,
		},
		{
			name:   "With size 6",
			props:  SubtitleProps{Size: Is6},
			expect: `<h2 class="subtitle is-6"></h2>`,
		},
		{
			name:   "With size 7",
			props:  SubtitleProps{Size: Is7},
			expect: `<h2 class="subtitle is-7"></h2>`,
		},
		{
			name:   "With spaced",
			props:  SubtitleProps{IsSpaced: true},
			expect: `<h2 class="subtitle is-spaced"></h2>`,
		},
		{
			name:   "With h1 level",
			props:  SubtitleProps{Level: 1},
			expect: `<h1 class="subtitle"></h1>`,
		},
		{
			name:   "With h3 level",
			props:  SubtitleProps{Level: 3},
			expect: `<h3 class="subtitle"></h3>`,
		},
		{
			name:   "With h4 level",
			props:  SubtitleProps{Level: 4},
			expect: `<h4 class="subtitle"></h4>`,
		},
		{
			name:   "With h5 level",
			props:  SubtitleProps{Level: 5},
			expect: `<h5 class="subtitle"></h5>`,
		},
		{
			name:   "With h6 level",
			props:  SubtitleProps{Level: 6},
			expect: `<h6 class="subtitle"></h6>`,
		},
		{
			name:   "Level too high (defaults to h6)",
			props:  SubtitleProps{Level: 10},
			expect: `<h6 class="subtitle"></h6>`,
		},
		{
			name:   "Level too low (defaults to h1)",
			props:  SubtitleProps{Level: -1},
			expect: `<h1 class="subtitle"></h1>`,
		},
		{
			name:   "With size and level",
			props:  SubtitleProps{Size: Is5, Level: 3},
			expect: `<h3 class="subtitle is-5"></h3>`,
		},
		{
			name: "All fields combined",
			props: SubtitleProps{
				ID:         "main-subtitle",
				Class:      []string{"foo", "bar"},
				Size:       Is4,
				Level:      3,
				IsSpaced:   true,
				Attributes: templ.Attributes{"data-testid": "page-subtitle"},
			},
			expect: `<h3 id="main-subtitle" class="subtitle is-4 is-spaced foo bar" data-testid="page-subtitle"></h3>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Subtitle(tt.props).Render(context.Background(), &buf)
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

func TestTitleSizes(t *testing.T) {
	tests := []struct {
		name   string
		size   Size
		expect string
	}{
		{"Size 1", Is1, `<h1 class="title is-1"></h1>`},
		{"Size 2", Is2, `<h1 class="title is-2"></h1>`},
		{"Size 3", Is3, `<h1 class="title is-3"></h1>`},
		{"Size 4", Is4, `<h1 class="title is-4"></h1>`},
		{"Size 5", Is5, `<h1 class="title is-5"></h1>`},
		{"Size 6", Is6, `<h1 class="title is-6"></h1>`},
		{"Size 7", Is7, `<h1 class="title is-7"></h1>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Title(TitleProps{Size: tt.size}).Render(context.Background(), &buf)
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

func TestSubtitleSizes(t *testing.T) {
	tests := []struct {
		name   string
		size   Size
		expect string
	}{
		{"Size 1", Is1, `<h2 class="subtitle is-1"></h2>`},
		{"Size 2", Is2, `<h2 class="subtitle is-2"></h2>`},
		{"Size 3", Is3, `<h2 class="subtitle is-3"></h2>`},
		{"Size 4", Is4, `<h2 class="subtitle is-4"></h2>`},
		{"Size 5", Is5, `<h2 class="subtitle is-5"></h2>`},
		{"Size 6", Is6, `<h2 class="subtitle is-6"></h2>`},
		{"Size 7", Is7, `<h2 class="subtitle is-7"></h2>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Subtitle(SubtitleProps{Size: tt.size}).Render(context.Background(), &buf)
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

func TestTitleLevels(t *testing.T) {
	tests := []struct {
		name   string
		level  int
		expect string
	}{
		{"Default level (0 -> h1)", 0, `<h1 class="title"></h1>`},
		{"Level 1", 1, `<h1 class="title"></h1>`},
		{"Level 2", 2, `<h2 class="title"></h2>`},
		{"Level 3", 3, `<h3 class="title"></h3>`},
		{"Level 4", 4, `<h4 class="title"></h4>`},
		{"Level 5", 5, `<h5 class="title"></h5>`},
		{"Level 6", 6, `<h6 class="title"></h6>`},
		{"Too high level (7 -> h6)", 7, `<h6 class="title"></h6>`},
		{"Too low level (-1 -> h1)", -1, `<h1 class="title"></h1>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Title(TitleProps{Level: tt.level}).Render(context.Background(), &buf)
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

func TestSubtitleLevels(t *testing.T) {
	tests := []struct {
		name   string
		level  int
		expect string
	}{
		{"Default level (0 -> h2)", 0, `<h2 class="subtitle"></h2>`},
		{"Level 1", 1, `<h1 class="subtitle"></h1>`},
		{"Level 2", 2, `<h2 class="subtitle"></h2>`},
		{"Level 3", 3, `<h3 class="subtitle"></h3>`},
		{"Level 4", 4, `<h4 class="subtitle"></h4>`},
		{"Level 5", 5, `<h5 class="subtitle"></h5>`},
		{"Level 6", 6, `<h6 class="subtitle"></h6>`},
		{"Too high level (7 -> h6)", 7, `<h6 class="subtitle"></h6>`},
		{"Too low level (-1 -> h1)", -1, `<h1 class="subtitle"></h1>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Subtitle(SubtitleProps{Level: tt.level}).Render(context.Background(), &buf)
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

func TestTitleSubtitleSpacing(t *testing.T) {
	tests := []struct {
		name      string
		component string
		isSpaced  bool
		expect    string
	}{
		{
			name:      "Title without spaced",
			component: "title",
			isSpaced:  false,
			expect:    `<h1 class="title"></h1>`,
		},
		{
			name:      "Title with spaced",
			component: "title",
			isSpaced:  true,
			expect:    `<h1 class="title is-spaced"></h1>`,
		},
		{
			name:      "Subtitle without spaced",
			component: "subtitle",
			isSpaced:  false,
			expect:    `<h2 class="subtitle"></h2>`,
		},
		{
			name:      "Subtitle with spaced",
			component: "subtitle",
			isSpaced:  true,
			expect:    `<h2 class="subtitle is-spaced"></h2>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			var err error

			if tt.component == "title" {
				err = Title(TitleProps{IsSpaced: tt.isSpaced}).Render(context.Background(), &buf)
			} else {
				err = Subtitle(SubtitleProps{IsSpaced: tt.isSpaced}).Render(context.Background(), &buf)
			}

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
