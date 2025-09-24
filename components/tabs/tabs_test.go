package tabs

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestTabs(t *testing.T) {
	tests := []struct {
		name   string
		props  TabsProps
		expect string
	}{
		{
			name:   "Default",
			props:  TabsProps{},
			expect: `<div class="tabs"></div>`,
		},
		{
			name:   "With ID",
			props:  TabsProps{ID: "nav-tabs"},
			expect: `<div id="nav-tabs" class="tabs"></div>`,
		},
		{
			name:   "With classes",
			props:  TabsProps{Class: []string{"custom", "nav"}},
			expect: `<div class="tabs custom nav"></div>`,
		},
		{
			name:   "Small size",
			props:  TabsProps{Size: IsSmall},
			expect: `<div class="tabs is-small"></div>`,
		},
		{
			name:   "Medium size",
			props:  TabsProps{Size: IsMedium},
			expect: `<div class="tabs is-medium"></div>`,
		},
		{
			name:   "Large size",
			props:  TabsProps{Size: IsLarge},
			expect: `<div class="tabs is-large"></div>`,
		},
		{
			name:   "Centered",
			props:  TabsProps{Alignment: IsCentered},
			expect: `<div class="tabs is-centered"></div>`,
		},
		{
			name:   "Right aligned",
			props:  TabsProps{Alignment: IsRight},
			expect: `<div class="tabs is-right"></div>`,
		},
		{
			name:   "Fullwidth",
			props:  TabsProps{IsFullwidth: true},
			expect: `<div class="tabs is-fullwidth"></div>`,
		},
		{
			name:   "Boxed",
			props:  TabsProps{IsBoxed: true},
			expect: `<div class="tabs is-boxed"></div>`,
		},
		{
			name:   "Toggle",
			props:  TabsProps{IsToggle: true},
			expect: `<div class="tabs is-toggle"></div>`,
		},
		{
			name:   "Toggle rounded",
			props:  TabsProps{IsToggleRounded: true},
			expect: `<div class="tabs is-toggle-rounded"></div>`,
		},
		{
			name: "Combined modifiers",
			props: TabsProps{
				Size:        IsLarge,
				Alignment:   IsCentered,
				IsFullwidth: true,
				IsBoxed:     true,
			},
			expect: `<div class="tabs is-large is-centered is-boxed is-fullwidth"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Tabs(tt.props).Render(context.Background(), &buf)
			if err != nil {
				t.Fatalf("render failed: %v", err)
			}
			got := strings.TrimSpace(buf.String())
			if got != tt.expect {
				t.Errorf("expected: %s, got: %s", tt.expect, got)
			}
		})
	}
}

func TestTabsWithAttributes(t *testing.T) {
	props := TabsProps{
		Attributes: templ.Attributes{
			"data-role":  "navigation",
			"aria-label": "Main tabs",
		},
	}

	var buf strings.Builder
	err := Tabs(props).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}

	got := buf.String()
	requiredParts := []string{
		`<div class="tabs"`,
		`data-role="navigation"`,
		`aria-label="Main tabs"`,
		`></div>`,
	}

	for _, part := range requiredParts {
		if !strings.Contains(got, part) {
			t.Errorf("missing required part %q in output: %s", part, got)
		}
	}
}

func TestTabsConstants(t *testing.T) {
	sizeTests := []struct {
		size Size
		want string
	}{
		{IsSmall, "is-small"},
		{IsMedium, "is-medium"},
		{IsLarge, "is-large"},
	}
	for _, tt := range sizeTests {
		if string(tt.size) != tt.want {
			t.Errorf("Size %s = %q, want %q", tt.size, string(tt.size), tt.want)
		}
	}

	alignTests := []struct {
		align Alignment
		want  string
	}{
		{IsCentered, "is-centered"},
		{IsRight, "is-right"},
		{IsCenter, "is-center"},
		{IsLeft, "is-left"},
	}
	for _, tt := range alignTests {
		if string(tt.align) != tt.want {
			t.Errorf("Alignment %s = %q, want %q", tt.align, string(tt.align), tt.want)
		}
	}
}
