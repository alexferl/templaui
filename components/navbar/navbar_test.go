package navbar

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestNavbar(t *testing.T) {
	tests := []struct {
		name   string
		props  NavbarProps
		expect string
	}{
		{
			name:   "Default",
			props:  NavbarProps{},
			expect: `<nav class="navbar" role="navigation"></nav>`,
		},
		{
			name:   "With ID and custom classes",
			props:  NavbarProps{ID: "navbar1", Class: []string{"custom-navbar"}},
			expect: `<nav id="navbar1" class="navbar custom-navbar" role="navigation"></nav>`,
		},
		{
			name:   "With color",
			props:  NavbarProps{Color: IsPrimary},
			expect: `<nav class="navbar is-primary" role="navigation"></nav>`,
		},
		{
			name:   "With fixed positioning",
			props:  NavbarProps{Fixed: FixedTop},
			expect: `<nav class="navbar is-fixed-top" role="navigation"></nav>`,
		},
		{
			name:   "With responsive fixed",
			props:  NavbarProps{FixedResponsive: FixedTopDesktop},
			expect: `<nav class="navbar is-fixed-top-desktop" role="navigation"></nav>`,
		},
		{
			name:   "With dropdown direction",
			props:  NavbarProps{DropdownDirection: HasDropdown},
			expect: `<nav class="navbar has-dropdown" role="navigation"></nav>`,
		},
		{
			name:   "With style and layout flags",
			props:  NavbarProps{IsSpaced: true, HasShadow: true, IsTransparent: true},
			expect: `<nav class="navbar is-spaced has-shadow is-transparent" role="navigation"></nav>`,
		},
		{
			name:   "With states",
			props:  NavbarProps{IsActive: true, IsExpanded: true},
			expect: `<nav class="navbar is-active is-expanded" role="navigation"></nav>`,
		},
		{
			name:   "With aria-label",
			props:  NavbarProps{AriaLabel: "Main Navigation"},
			expect: `<nav aria-label="Main Navigation" class="navbar" role="navigation"></nav>`,
		},
		{
			name: "All fields combined",
			props: NavbarProps{
				ID:                "main-navbar",
				Class:             []string{"foo", "bar"},
				Color:             IsDanger,
				Fixed:             FixedTop,
				DropdownDirection: HasDropdownUp,
				IsSpaced:          true,
				HasShadow:         true,
				IsActive:          true,
				AriaLabel:         "Primary Navigation",
				Attributes:        templ.Attributes{"data-role": "navigation"},
			},
			expect: `<nav id="main-navbar" aria-label="Primary Navigation" class="navbar is-danger is-fixed-top has-dropdown-up is-spaced has-shadow is-active foo bar" role="navigation" data-role="navigation"></nav>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Navbar(tt.props).Render(context.Background(), &buf)
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

func TestNavbarBrand(t *testing.T) {
	tests := []struct {
		name   string
		props  NavbarBrandProps
		expect string
	}{
		{
			name:   "Default",
			props:  NavbarBrandProps{},
			expect: `<div class="navbar-brand"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  NavbarBrandProps{ID: "brand1", Class: []string{"custom-brand"}},
			expect: `<div id="brand1" class="navbar-brand custom-brand"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := NavbarBrand(tt.props).Render(context.Background(), &buf)
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

func TestNavbarBurger(t *testing.T) {
	tests := []struct {
		name   string
		props  NavbarBurgerProps
		expect string
	}{
		{
			name:   "Default",
			props:  NavbarBurgerProps{},
			expect: `<a class="navbar-burger" role="button" aria-label="menu" aria-expanded="false" aria-controls="navbar-menu"></a>`,
		},
		{
			name:   "With active state",
			props:  NavbarBurgerProps{IsActive: true},
			expect: `<a class="navbar-burger is-active" role="button" aria-label="menu" aria-expanded="true" aria-controls="navbar-menu"></a>`,
		},
		{
			name:   "With ID and custom classes",
			props:  NavbarBurgerProps{ID: "burger1", Class: []string{"custom-burger"}},
			expect: `<a id="burger1" class="navbar-burger custom-burger" role="button" aria-label="menu" aria-expanded="false" aria-controls="navbar-menu"></a>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := NavbarBurger(tt.props).Render(context.Background(), &buf)
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

func TestNavbarMenu(t *testing.T) {
	tests := []struct {
		name   string
		props  NavbarMenuProps
		expect string
	}{
		{
			name:   "Default",
			props:  NavbarMenuProps{},
			expect: `<div id="navbar-menu" class="navbar-menu" aria-labelledby="navbar-burger"></div>`,
		},
		{
			name:   "With custom ID",
			props:  NavbarMenuProps{ID: "custom-menu"},
			expect: `<div id="custom-menu" class="navbar-menu" aria-labelledby="navbar-burger"></div>`,
		},
		{
			name:   "With active state",
			props:  NavbarMenuProps{IsActive: true},
			expect: `<div id="navbar-menu" class="navbar-menu is-active" aria-labelledby="navbar-burger"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := NavbarMenu(tt.props).Render(context.Background(), &buf)
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

func TestNavbarStart(t *testing.T) {
	tests := []struct {
		name   string
		props  NavbarStartProps
		expect string
	}{
		{
			name:   "Default",
			props:  NavbarStartProps{},
			expect: `<div class="navbar-start"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  NavbarStartProps{ID: "start1", Class: []string{"custom-start"}},
			expect: `<div id="start1" class="navbar-start custom-start"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := NavbarStart(tt.props).Render(context.Background(), &buf)
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

func TestNavbarEnd(t *testing.T) {
	tests := []struct {
		name   string
		props  NavbarEndProps
		expect string
	}{
		{
			name:   "Default",
			props:  NavbarEndProps{},
			expect: `<div class="navbar-end"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  NavbarEndProps{ID: "end1", Class: []string{"custom-end"}},
			expect: `<div id="end1" class="navbar-end custom-end"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := NavbarEnd(tt.props).Render(context.Background(), &buf)
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

func TestNavbarItem(t *testing.T) {
	tests := []struct {
		name   string
		props  NavbarItemProps
		expect string
	}{
		{
			name:   "Default",
			props:  NavbarItemProps{},
			expect: `<div class="navbar-item"></div>`,
		},
		{
			name:   "With states",
			props:  NavbarItemProps{IsActive: true, IsHoverable: true},
			expect: `<div class="navbar-item is-active is-hoverable"></div>`,
		},
		{
			name:   "With dropdown direction",
			props:  NavbarItemProps{DropdownDirection: HasDropdown},
			expect: `<div class="navbar-item has-dropdown"></div>`,
		},
		{
			name:   "With tab style",
			props:  NavbarItemProps{IsTab: true},
			expect: `<div class="navbar-item is-tab"></div>`,
		},
		{
			name: "All fields combined",
			props: NavbarItemProps{
				ID:                "item1",
				Class:             []string{"custom-item"},
				IsActive:          true,
				IsExpanded:        true,
				DropdownDirection: HasDropdownUp,
				IsTab:             true,
			},
			expect: `<div id="item1" class="navbar-item is-active is-expanded has-dropdown-up is-tab custom-item"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := NavbarItem(tt.props).Render(context.Background(), &buf)
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

func TestNavbarLink(t *testing.T) {
	tests := []struct {
		name   string
		props  NavbarLinkProps
		expect string
	}{
		{
			name:   "Default",
			props:  NavbarLinkProps{},
			expect: `<a class="navbar-link"></a>`,
		},
		{
			name:   "With active state",
			props:  NavbarLinkProps{IsActive: true},
			expect: `<a class="navbar-link is-active"></a>`,
		},
		{
			name:   "With arrowless style",
			props:  NavbarLinkProps{IsArrowless: true},
			expect: `<a class="navbar-link is-arrowless"></a>`,
		},
		{
			name: "All fields combined",
			props: NavbarLinkProps{
				ID:          "link1",
				Class:       []string{"custom-link"},
				IsActive:    true,
				IsArrowless: true,
			},
			expect: `<a id="link1" class="navbar-link is-active is-arrowless custom-link"></a>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := NavbarLink(tt.props).Render(context.Background(), &buf)
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

func TestNavbarDropdown(t *testing.T) {
	tests := []struct {
		name   string
		props  NavbarDropdownProps
		expect string
	}{
		{
			name:   "Default",
			props:  NavbarDropdownProps{},
			expect: `<div class="navbar-dropdown"></div>`,
		},
		{
			name:   "With right alignment",
			props:  NavbarDropdownProps{IsRight: true},
			expect: `<div class="navbar-dropdown is-right"></div>`,
		},
		{
			name:   "With boxed style",
			props:  NavbarDropdownProps{IsBoxed: true},
			expect: `<div class="navbar-dropdown is-boxed"></div>`,
		},
		{
			name: "All fields combined",
			props: NavbarDropdownProps{
				ID:      "dropdown1",
				Class:   []string{"custom-dropdown"},
				IsRight: true,
				IsBoxed: true,
			},
			expect: `<div id="dropdown1" class="navbar-dropdown is-right is-boxed custom-dropdown"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := NavbarDropdown(tt.props).Render(context.Background(), &buf)
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

func TestNavbarDivider(t *testing.T) {
	tests := []struct {
		name   string
		props  NavbarDividerProps
		expect string
	}{
		{
			name:   "Default",
			props:  NavbarDividerProps{},
			expect: `<hr class="navbar-divider">`,
		},
		{
			name:   "With ID and custom classes",
			props:  NavbarDividerProps{ID: "divider1", Class: []string{"custom-divider"}},
			expect: `<hr id="divider1" class="navbar-divider custom-divider">`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := NavbarDivider(tt.props).Render(context.Background(), &buf)
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
