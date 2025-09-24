package hero

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestHero(t *testing.T) {
	tests := []struct {
		name   string
		props  HeroProps
		expect string
	}{
		{
			name:   "Default",
			props:  HeroProps{},
			expect: `<section class="hero"></section>`,
		},
		{
			name:   "With ID and custom classes",
			props:  HeroProps{ID: "hero1", Class: []string{"custom-hero"}},
			expect: `<section id="hero1" class="hero custom-hero"></section>`,
		},
		{
			name:   "With color and size",
			props:  HeroProps{Color: IsPrimary, Size: IsLarge},
			expect: `<section class="hero is-primary is-large"></section>`,
		},
		{
			name:   "With fullheight size",
			props:  HeroProps{Size: IsFullHeight},
			expect: `<section class="hero is-fullheight"></section>`,
		},
		{
			name:   "With fullheight with navbar size",
			props:  HeroProps{Size: IsFullHeightWithNavbar},
			expect: `<section class="hero is-fullheight-with-navbar"></section>`,
		},
		{
			name: "All fields combined",
			props: HeroProps{
				ID:         "main-hero",
				Class:      []string{"foo", "bar"},
				Color:      IsDanger,
				Size:       IsMedium,
				Attributes: templ.Attributes{"data-role": "banner"},
			},
			expect: `<section id="main-hero" class="hero is-danger is-medium foo bar" data-role="banner"></section>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Hero(tt.props).Render(context.Background(), &buf)
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

func TestHeroHead(t *testing.T) {
	tests := []struct {
		name   string
		props  HeroHeadProps
		expect string
	}{
		{
			name:   "Default",
			props:  HeroHeadProps{},
			expect: `<div class="hero-head"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  HeroHeadProps{ID: "head1", Class: []string{"custom-head"}},
			expect: `<div id="head1" class="hero-head custom-head"></div>`,
		},
		{
			name: "With attributes",
			props: HeroHeadProps{
				ID:         "hero-header",
				Attributes: templ.Attributes{"role": "banner"},
			},
			expect: `<div id="hero-header" class="hero-head" role="banner"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := HeroHead(tt.props).Render(context.Background(), &buf)
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

func TestHeroBody(t *testing.T) {
	tests := []struct {
		name   string
		props  HeroBodyProps
		expect string
	}{
		{
			name:   "Default",
			props:  HeroBodyProps{},
			expect: `<div class="hero-body"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  HeroBodyProps{ID: "body1", Class: []string{"custom-body"}},
			expect: `<div id="body1" class="hero-body custom-body"></div>`,
		},
		{
			name: "With multiple classes and attributes",
			props: HeroBodyProps{
				ID:         "hero-content",
				Class:      []string{"main-content", "centered"},
				Attributes: templ.Attributes{"data-section": "main"},
			},
			expect: `<div id="hero-content" class="hero-body main-content centered" data-section="main"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := HeroBody(tt.props).Render(context.Background(), &buf)
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

func TestHeroFoot(t *testing.T) {
	tests := []struct {
		name   string
		props  HeroFootProps
		expect string
	}{
		{
			name:   "Default",
			props:  HeroFootProps{},
			expect: `<div class="hero-foot"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  HeroFootProps{ID: "foot1", Class: []string{"custom-foot"}},
			expect: `<div id="foot1" class="hero-foot custom-foot"></div>`,
		},
		{
			name: "With attributes",
			props: HeroFootProps{
				ID:         "hero-footer",
				Class:      []string{"navigation"},
				Attributes: templ.Attributes{"data-nav": "tabs"},
			},
			expect: `<div id="hero-footer" class="hero-foot navigation" data-nav="tabs"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := HeroFoot(tt.props).Render(context.Background(), &buf)
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
