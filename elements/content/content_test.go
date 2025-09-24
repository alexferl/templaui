package content

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestContent(t *testing.T) {
	tests := []struct {
		name   string
		props  ContentProps
		expect string
	}{
		{
			name:   "Default",
			props:  ContentProps{},
			expect: `<div class="content"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  ContentProps{ID: "content1", Class: []string{"custom-content"}},
			expect: `<div id="content1" class="content custom-content"></div>`,
		},
		{
			name:   "With size",
			props:  ContentProps{Size: IsLarge},
			expect: `<div class="content is-large"></div>`,
		},
		{
			name:   "With small size",
			props:  ContentProps{Size: IsSmall},
			expect: `<div class="content is-small"></div>`,
		},
		{
			name:   "With normal size",
			props:  ContentProps{Size: IsNormal},
			expect: `<div class="content is-normal"></div>`,
		},
		{
			name:   "With medium size",
			props:  ContentProps{Size: IsMedium},
			expect: `<div class="content is-medium"></div>`,
		},
		{
			name:   "With lower-alpha style",
			props:  ContentProps{Style: IsLowerAlpha},
			expect: `<div class="content is-lower-alpha"></div>`,
		},
		{
			name:   "With upper-roman style",
			props:  ContentProps{Style: IsUpperRoman},
			expect: `<div class="content is-upper-roman"></div>`,
		},
		{
			name:   "With lower-roman style",
			props:  ContentProps{Style: IsLowerRoman},
			expect: `<div class="content is-lower-roman"></div>`,
		},
		{
			name:   "With upper-alpha style",
			props:  ContentProps{Style: IsUpperAlpha},
			expect: `<div class="content is-upper-alpha"></div>`,
		},
		{
			name:   "With size and style",
			props:  ContentProps{Size: IsMedium, Style: IsLowerAlpha},
			expect: `<div class="content is-medium is-lower-alpha"></div>`,
		},
		{
			name: "All fields combined",
			props: ContentProps{
				ID:         "main-content",
				Class:      []string{"foo", "bar"},
				Size:       IsLarge,
				Style:      IsUpperRoman,
				Attributes: templ.Attributes{"data-role": "article"},
			},
			expect: `<div id="main-content" class="content is-large is-upper-roman foo bar" data-role="article"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Content(tt.props).Render(context.Background(), &buf)
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
