package section

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestSection(t *testing.T) {
	tests := []struct {
		name   string
		props  SectionProps
		expect string
	}{
		{
			name:   "Default",
			props:  SectionProps{},
			expect: `<section class="section"></section>`,
		},
		{
			name:   "With ID and custom classes",
			props:  SectionProps{ID: "section1", Class: []string{"custom-section"}},
			expect: `<section id="section1" class="section custom-section"></section>`,
		},
		{
			name:   "With medium size",
			props:  SectionProps{Size: IsMedium},
			expect: `<section class="section is-medium"></section>`,
		},
		{
			name:   "With large size",
			props:  SectionProps{Size: IsLarge},
			expect: `<section class="section is-large"></section>`,
		},
		{
			name:   "With fullheight size",
			props:  SectionProps{Size: IsFullHeight},
			expect: `<section class="section is-fullheight"></section>`,
		},
		{
			name: "All fields combined",
			props: SectionProps{
				ID:         "main-section",
				Class:      []string{"foo", "bar"},
				Size:       IsLarge,
				Attributes: templ.Attributes{"data-role": "main"},
			},
			expect: `<section id="main-section" class="section is-large foo bar" data-role="main"></section>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Section(tt.props).Render(context.Background(), &buf)
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
