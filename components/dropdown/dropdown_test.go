package dropdown

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestDropdown(t *testing.T) {
	tests := []struct {
		name   string
		props  DropdownProps
		expect string
	}{
		{
			name:   "Default",
			props:  DropdownProps{},
			expect: `<div class="dropdown"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  DropdownProps{ID: "dropdown1", Class: []string{"custom-dropdown"}},
			expect: `<div id="dropdown1" class="dropdown custom-dropdown"></div>`,
		},
		{
			name:   "With alignment and states",
			props:  DropdownProps{IsRight: true, IsActive: true, IsHoverable: true},
			expect: `<div class="dropdown is-right is-active is-hoverable"></div>`,
		},
		{
			name: "All fields combined",
			props: DropdownProps{
				ID:          "main-dropdown",
				Class:       []string{"foo", "bar"},
				IsRight:     true,
				IsActive:    true,
				IsHoverable: true,
				IsSelected:  true,
				IsUp:        true,
				Attributes:  templ.Attributes{"data-role": "menu"},
			},
			expect: `<div id="main-dropdown" class="dropdown is-right is-active is-hoverable is-selected is-up foo bar" data-role="menu"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Dropdown(tt.props).Render(context.Background(), &buf)
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

func TestDropdownTrigger(t *testing.T) {
	tests := []struct {
		name   string
		props  DropdownTriggerProps
		expect string
	}{
		{
			name:   "Default",
			props:  DropdownTriggerProps{},
			expect: `<div class="dropdown-trigger"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  DropdownTriggerProps{ID: "trigger1", Class: []string{"custom-trigger"}},
			expect: `<div id="trigger1" class="dropdown-trigger custom-trigger"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := DropdownTrigger(tt.props).Render(context.Background(), &buf)
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

func TestDropdownMenu(t *testing.T) {
	tests := []struct {
		name   string
		props  DropdownMenuProps
		expect string
	}{
		{
			name:   "Default",
			props:  DropdownMenuProps{},
			expect: `<div class="dropdown-menu"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  DropdownMenuProps{ID: "menu1", Class: []string{"custom-menu"}},
			expect: `<div id="menu1" class="dropdown-menu custom-menu"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := DropdownMenu(tt.props).Render(context.Background(), &buf)
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

func TestDropdownContent(t *testing.T) {
	tests := []struct {
		name   string
		props  DropdownContentProps
		expect string
	}{
		{
			name:   "Default",
			props:  DropdownContentProps{},
			expect: `<div class="dropdown-content"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  DropdownContentProps{ID: "content1", Class: []string{"custom-content"}},
			expect: `<div id="content1" class="dropdown-content custom-content"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := DropdownContent(tt.props).Render(context.Background(), &buf)
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

func TestDropdownItem(t *testing.T) {
	tests := []struct {
		name   string
		props  DropdownItemProps
		expect string
	}{
		{
			name:   "Default",
			props:  DropdownItemProps{},
			expect: `<a class="dropdown-item"></a>`,
		},
		{
			name:   "With active state",
			props:  DropdownItemProps{IsActive: true},
			expect: `<a class="dropdown-item is-active"></a>`,
		},
		{
			name: "All fields combined",
			props: DropdownItemProps{
				ID:         "item1",
				Class:      []string{"custom-item"},
				IsActive:   true,
				Attributes: templ.Attributes{"href": "#test"},
			},
			expect: `<a id="item1" class="dropdown-item is-active custom-item" href="#test"></a>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := DropdownItem(tt.props).Render(context.Background(), &buf)
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

func TestDropdownDivider(t *testing.T) {
	tests := []struct {
		name   string
		props  DropdownDividerProps
		expect string
	}{
		{
			name:   "Default",
			props:  DropdownDividerProps{},
			expect: `<hr class="dropdown-divider">`,
		},
		{
			name:   "With ID and custom classes",
			props:  DropdownDividerProps{ID: "divider1", Class: []string{"custom-divider"}},
			expect: `<hr id="divider1" class="dropdown-divider custom-divider">`,
		},
		{
			name: "With attributes",
			props: DropdownDividerProps{
				Attributes: templ.Attributes{"data-separator": "true"},
			},
			expect: `<hr class="dropdown-divider" data-separator="true">`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := DropdownDivider(tt.props).Render(context.Background(), &buf)
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
