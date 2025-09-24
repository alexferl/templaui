package footer

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestFooter(t *testing.T) {
	tests := []struct {
		name   string
		props  FooterProps
		expect string
	}{
		{
			name:   "Default",
			props:  FooterProps{},
			expect: `<footer class="footer"></footer>`,
		},
		{
			name:   "With ID and custom classes",
			props:  FooterProps{ID: "footer1", Class: []string{"custom-footer"}},
			expect: `<footer id="footer1" class="footer custom-footer"></footer>`,
		},
		{
			name: "With multiple classes",
			props: FooterProps{
				Class: []string{"site-footer", "dark-theme"},
			},
			expect: `<footer class="footer site-footer dark-theme"></footer>`,
		},
		{
			name: "With attributes",
			props: FooterProps{
				ID:         "main-footer",
				Attributes: templ.Attributes{"role": "contentinfo"},
			},
			expect: `<footer id="main-footer" class="footer" role="contentinfo"></footer>`,
		},
		{
			name: "All fields combined",
			props: FooterProps{
				ID:         "site-footer",
				Class:      []string{"foo", "bar"},
				Attributes: templ.Attributes{"data-section": "footer", "role": "contentinfo"},
			},
			expect: `<footer id="site-footer" class="footer foo bar" data-section="footer" role="contentinfo"></footer>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Footer(tt.props).Render(context.Background(), &buf)
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
