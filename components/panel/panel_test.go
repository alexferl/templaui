package panel

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestPanel(t *testing.T) {
	tests := []struct {
		name   string
		props  PanelProps
		expect string
	}{
		{
			name:   "Default",
			props:  PanelProps{},
			expect: `<nav class="panel"></nav>`,
		},
		{
			name:   "With ID and custom classes",
			props:  PanelProps{ID: "panel1", Class: []string{"custom-panel"}},
			expect: `<nav id="panel1" class="panel custom-panel"></nav>`,
		},
		{
			name:   "With color",
			props:  PanelProps{Color: IsPrimary},
			expect: `<nav class="panel is-primary"></nav>`,
		},
		{
			name:   "With danger color",
			props:  PanelProps{Color: IsDanger},
			expect: `<nav class="panel is-danger"></nav>`,
		},
		{
			name:   "With active style",
			props:  PanelProps{IsActive: true},
			expect: `<nav class="panel is-active"></nav>`,
		},
		{
			name:   "With wrapped layout",
			props:  PanelProps{IsWrapped: true},
			expect: `<nav class="panel is-wrapped"></nav>`,
		},
		{
			name: "All fields combined",
			props: PanelProps{
				ID:         "main-panel",
				Class:      []string{"foo", "bar"},
				Color:      IsInfo,
				IsActive:   true,
				IsWrapped:  true,
				Attributes: templ.Attributes{"data-role": "sidebar"},
			},
			expect: `<nav id="main-panel" class="panel is-info is-active is-wrapped foo bar" data-role="sidebar"></nav>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Panel(tt.props).Render(context.Background(), &buf)
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

func TestPanelHeading(t *testing.T) {
	tests := []struct {
		name   string
		props  PanelHeadingProps
		expect string
	}{
		{
			name:   "Default",
			props:  PanelHeadingProps{},
			expect: `<p class="panel-heading"></p>`,
		},
		{
			name:   "With ID and custom classes",
			props:  PanelHeadingProps{ID: "heading1", Class: []string{"custom-heading"}},
			expect: `<p id="heading1" class="panel-heading custom-heading"></p>`,
		},
		{
			name: "With attributes",
			props: PanelHeadingProps{
				ID:         "panel-title",
				Attributes: templ.Attributes{"data-title": "main"},
			},
			expect: `<p id="panel-title" class="panel-heading" data-title="main"></p>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := PanelHeading(tt.props).Render(context.Background(), &buf)
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

func TestPanelTabs(t *testing.T) {
	tests := []struct {
		name   string
		props  PanelTabsProps
		expect string
	}{
		{
			name:   "Default",
			props:  PanelTabsProps{},
			expect: `<p class="panel-tabs"></p>`,
		},
		{
			name:   "With ID and custom classes",
			props:  PanelTabsProps{ID: "tabs1", Class: []string{"custom-tabs"}},
			expect: `<p id="tabs1" class="panel-tabs custom-tabs"></p>`,
		},
		{
			name: "With attributes",
			props: PanelTabsProps{
				ID:         "navigation-tabs",
				Attributes: templ.Attributes{"data-nav": "tabs"},
			},
			expect: `<p id="navigation-tabs" class="panel-tabs" data-nav="tabs"></p>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := PanelTabs(tt.props).Render(context.Background(), &buf)
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

func TestPanelBlock(t *testing.T) {
	tests := []struct {
		name   string
		props  PanelBlockProps
		expect string
	}{
		{
			name:   "Default (div)",
			props:  PanelBlockProps{},
			expect: `<div class="panel-block"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  PanelBlockProps{ID: "block1", Class: []string{"custom-block"}},
			expect: `<div id="block1" class="panel-block custom-block"></div>`,
		},
		{
			name:   "With active state",
			props:  PanelBlockProps{IsActive: true},
			expect: `<div class="panel-block is-active"></div>`,
		},
		{
			name:   "As anchor",
			props:  PanelBlockProps{IsAnchor: true},
			expect: `<a class="panel-block"></a>`,
		},
		{
			name:   "As label",
			props:  PanelBlockProps{IsLabel: true},
			expect: `<label class="panel-block"></label>`,
		},
		{
			name:   "Anchor with active state",
			props:  PanelBlockProps{IsAnchor: true, IsActive: true},
			expect: `<a class="panel-block is-active"></a>`,
		},
		{
			name:   "Label with ID and classes",
			props:  PanelBlockProps{ID: "checkbox-block", IsLabel: true, Class: []string{"with-checkbox"}},
			expect: `<label id="checkbox-block" class="panel-block with-checkbox"></label>`,
		},
		{
			name: "All fields combined (div)",
			props: PanelBlockProps{
				ID:         "main-block",
				Class:      []string{"foo", "bar"},
				IsActive:   true,
				Attributes: templ.Attributes{"data-block": "content"},
			},
			expect: `<div id="main-block" class="panel-block is-active foo bar" data-block="content"></div>`,
		},
		{
			name: "All fields combined (anchor)",
			props: PanelBlockProps{
				ID:         "link-block",
				Class:      []string{"clickable"},
				IsActive:   true,
				IsAnchor:   true,
				Attributes: templ.Attributes{"href": "/path"},
			},
			expect: `<a id="link-block" class="panel-block is-active clickable" href="/path"></a>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := PanelBlock(tt.props).Render(context.Background(), &buf)
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

func TestPanelIcon(t *testing.T) {
	tests := []struct {
		name   string
		props  PanelIconProps
		expect string
	}{
		{
			name:   "Default",
			props:  PanelIconProps{},
			expect: `<span class="panel-icon"></span>`,
		},
		{
			name:   "With ID and custom classes",
			props:  PanelIconProps{ID: "icon1", Class: []string{"custom-icon"}},
			expect: `<span id="icon1" class="panel-icon custom-icon"></span>`,
		},
		{
			name: "With attributes",
			props: PanelIconProps{
				ID:         "search-icon",
				Class:      []string{"fa", "fa-search"},
				Attributes: templ.Attributes{"aria-hidden": "true"},
			},
			expect: `<span id="search-icon" class="panel-icon fa fa-search" aria-hidden="true"></span>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := PanelIcon(tt.props).Render(context.Background(), &buf)
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
