package menu

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestMenu(t *testing.T) {
	tests := []struct {
		name   string
		props  MenuProps
		expect string
	}{
		{
			name:   "Default",
			props:  MenuProps{},
			expect: `<aside class="menu"></aside>`,
		},
		{
			name:   "With ID and custom classes",
			props:  MenuProps{ID: "menu1", Class: []string{"custom-menu"}},
			expect: `<aside id="menu1" class="menu custom-menu"></aside>`,
		},
		{
			name:   "With size and states",
			props:  MenuProps{Size: IsLarge, IsActive: true, IsSelected: true},
			expect: `<aside class="menu is-large is-active is-selected"></aside>`,
		},
		{
			name: "All fields combined",
			props: MenuProps{
				ID:         "main-menu",
				Class:      []string{"foo", "bar"},
				Size:       IsMedium,
				IsActive:   true,
				IsSelected: true,
				Attributes: templ.Attributes{"data-role": "navigation"},
			},
			expect: `<aside id="main-menu" class="menu is-medium is-active is-selected foo bar" data-role="navigation"></aside>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Menu(tt.props).Render(context.Background(), &buf)
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

func TestMenuLabel(t *testing.T) {
	tests := []struct {
		name   string
		props  MenuLabelProps
		expect string
	}{
		{
			name:   "Default",
			props:  MenuLabelProps{},
			expect: `<p class="menu-label"></p>`,
		},
		{
			name:   "With ID and custom classes",
			props:  MenuLabelProps{ID: "label1", Class: []string{"custom-label"}},
			expect: `<p id="label1" class="menu-label custom-label"></p>`,
		},
		{
			name: "With attributes",
			props: MenuLabelProps{
				ID:         "section-label",
				Attributes: templ.Attributes{"data-section": "main"},
			},
			expect: `<p id="section-label" class="menu-label" data-section="main"></p>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := MenuLabel(tt.props).Render(context.Background(), &buf)
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

func TestMenuList(t *testing.T) {
	tests := []struct {
		name   string
		props  MenuListProps
		expect string
	}{
		{
			name:   "Default",
			props:  MenuListProps{},
			expect: `<ul class="menu-list"></ul>`,
		},
		{
			name:   "With ID and custom classes",
			props:  MenuListProps{ID: "list1", Class: []string{"custom-list"}},
			expect: `<ul id="list1" class="menu-list custom-list"></ul>`,
		},
		{
			name: "With multiple classes and attributes",
			props: MenuListProps{
				ID:         "nav-list",
				Class:      []string{"primary", "vertical"},
				Attributes: templ.Attributes{"role": "menu"},
			},
			expect: `<ul id="nav-list" class="menu-list primary vertical" role="menu"></ul>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := MenuList(tt.props).Render(context.Background(), &buf)
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
