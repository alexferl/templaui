package breadcrumb

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestBreadcrumb(t *testing.T) {
	tests := []struct {
		name   string
		props  BreadcrumbProps
		expect string
	}{
		{
			name:   "Default",
			props:  BreadcrumbProps{},
			expect: `<nav class="breadcrumb"></nav>`,
		},
		{
			name:   "With ID and custom classes",
			props:  BreadcrumbProps{ID: "crumb1", Class: []string{"is-small", "custom-foo"}},
			expect: `<nav id="crumb1" class="breadcrumb is-small custom-foo"></nav>`,
		},
		{
			name:   "With alignment and active",
			props:  BreadcrumbProps{IsCentered: true, IsRight: true, IsActive: true},
			expect: `<nav class="breadcrumb is-centered is-right is-active"></nav>`,
		},
		{
			name: "All fields combined",
			props: BreadcrumbProps{
				ID:         "crumb2",
				Class:      []string{"foo", "bar"},
				IsCentered: true,
				IsRight:    true,
				IsActive:   true,
				Size:       IsMedium,
				Style:      HasDotSeparator,
				Attributes: templ.Attributes{"data-role": "navi"},
			},
			expect: `<nav id="crumb2" class="breadcrumb is-centered is-right is-medium has-dot-separator is-active foo bar" data-role="navi"></nav>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Breadcrumb(tt.props).Render(context.Background(), &buf)
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
