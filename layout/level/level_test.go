package level

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestLevel(t *testing.T) {
	tests := []struct {
		name   string
		props  LevelProps
		expect string
	}{
		{
			name:   "Default",
			props:  LevelProps{},
			expect: `<nav class="level"></nav>`,
		},
		{
			name:   "With ID and custom classes",
			props:  LevelProps{ID: "level1", Class: []string{"custom-level"}},
			expect: `<nav id="level1" class="level custom-level"></nav>`,
		},
		{
			name:   "With mobile",
			props:  LevelProps{IsMobile: true},
			expect: `<nav class="level is-mobile"></nav>`,
		},
		{
			name: "All fields combined",
			props: LevelProps{
				ID:         "main-level",
				Class:      []string{"foo", "bar"},
				IsMobile:   true,
				Attributes: templ.Attributes{"data-role": "navigation"},
			},
			expect: `<nav id="main-level" class="level is-mobile foo bar" data-role="navigation"></nav>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Level(tt.props).Render(context.Background(), &buf)
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

func TestLevelLeft(t *testing.T) {
	tests := []struct {
		name   string
		props  LevelLeftProps
		expect string
	}{
		{
			name:   "Default",
			props:  LevelLeftProps{},
			expect: `<div class="level-left"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  LevelLeftProps{ID: "left1", Class: []string{"custom-left"}},
			expect: `<div id="left1" class="level-left custom-left"></div>`,
		},
		{
			name: "With attributes",
			props: LevelLeftProps{
				ID:         "level-start",
				Attributes: templ.Attributes{"data-align": "left"},
			},
			expect: `<div id="level-start" class="level-left" data-align="left"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := LevelLeft(tt.props).Render(context.Background(), &buf)
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

func TestLevelRight(t *testing.T) {
	tests := []struct {
		name   string
		props  LevelRightProps
		expect string
	}{
		{
			name:   "Default",
			props:  LevelRightProps{},
			expect: `<div class="level-right"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  LevelRightProps{ID: "right1", Class: []string{"custom-right"}},
			expect: `<div id="right1" class="level-right custom-right"></div>`,
		},
		{
			name: "With attributes",
			props: LevelRightProps{
				ID:         "level-end",
				Attributes: templ.Attributes{"data-align": "right"},
			},
			expect: `<div id="level-end" class="level-right" data-align="right"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := LevelRight(tt.props).Render(context.Background(), &buf)
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

func TestLevelItem(t *testing.T) {
	tests := []struct {
		name   string
		props  LevelItemProps
		expect string
	}{
		{
			name:   "Default",
			props:  LevelItemProps{},
			expect: `<div class="level-item"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  LevelItemProps{ID: "item1", Class: []string{"custom-item"}},
			expect: `<div id="item1" class="level-item custom-item"></div>`,
		},
		{
			name:   "With flexible",
			props:  LevelItemProps{IsFlexible: true},
			expect: `<div class="level-item is-flexible"></div>`,
		},
		{
			name:   "With narrow",
			props:  LevelItemProps{IsNarrow: true},
			expect: `<div class="level-item is-narrow"></div>`,
		},
		{
			name: "All fields combined",
			props: LevelItemProps{
				ID:         "flex-item",
				Class:      []string{"content", "centered"},
				IsFlexible: true,
				IsNarrow:   true,
				Attributes: templ.Attributes{"data-grow": "true"},
			},
			expect: `<div id="flex-item" class="level-item is-flexible is-narrow content centered" data-grow="true"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := LevelItem(tt.props).Render(context.Background(), &buf)
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
