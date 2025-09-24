package button

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestButton(t *testing.T) {
	tests := []struct {
		name   string
		props  ButtonProps
		expect string
	}{
		{
			name:   "Default button element",
			props:  ButtonProps{},
			expect: `<button type="button" class="button"></button>`,
		},
		{
			name:   "With ID and custom classes",
			props:  ButtonProps{ID: "btn1", Class: []string{"custom-btn"}},
			expect: `<button id="btn1" type="button" class="button custom-btn"></button>`,
		},
		{
			name:   "With custom type",
			props:  ButtonProps{Type: "submit"},
			expect: `<button type="submit" class="button"></button>`,
		},
		{
			name:   "As anchor",
			props:  ButtonProps{IsAnchor: true},
			expect: `<a class="button"></a>`,
		},
		{
			name:   "As input",
			props:  ButtonProps{IsInput: true},
			expect: `<input type="button" class="button">`,
		},
		{
			name:   "As input with submit type",
			props:  ButtonProps{IsInput: true, Type: "submit"},
			expect: `<input type="submit" class="button">`,
		},
		{
			name:   "With size",
			props:  ButtonProps{Size: IsLarge},
			expect: `<button type="button" class="button is-large"></button>`,
		},
		{
			name:   "With color",
			props:  ButtonProps{Color: IsPrimary},
			expect: `<button type="button" class="button is-primary"></button>`,
		},
		{
			name:   "With responsive",
			props:  ButtonProps{IsResponsive: true},
			expect: `<button type="button" class="button is-responsive"></button>`,
		},
		{
			name:   "With style modifiers",
			props:  ButtonProps{IsOutlined: true, IsRounded: true, IsLoading: true},
			expect: `<button type="button" class="button is-loading is-outlined is-rounded"></button>`,
		},
		{
			name:   "With layout modifier",
			props:  ButtonProps{IsFullwidth: true},
			expect: `<button type="button" class="button is-fullwidth"></button>`,
		},
		{
			name:   "With states",
			props:  ButtonProps{IsActive: true, IsHovered: true, IsSelected: true},
			expect: `<button type="button" class="button is-active is-hovered is-selected"></button>`,
		},
		{
			name:   "Anchor with all modifiers",
			props:  ButtonProps{IsAnchor: true, Size: IsMedium, Color: IsDanger, IsOutlined: true, IsActive: true},
			expect: `<a class="button is-medium is-danger is-outlined is-active"></a>`,
		},
		{
			name: "All fields combined (button)",
			props: ButtonProps{
				ID:           "main-btn",
				Class:        []string{"foo", "bar"},
				Type:         "submit",
				Size:         IsLarge,
				Color:        IsSuccess,
				IsResponsive: true,
				IsBold:       true,
				IsGhost:      true,
				IsFullwidth:  true,
				IsActive:     true,
				Attributes:   templ.Attributes{"data-action": "save"},
			},
			expect: `<button id="main-btn" type="submit" class="button is-large is-success is-responsive is-bold is-ghost is-fullwidth is-active foo bar" data-action="save"></button>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Button(tt.props).Render(context.Background(), &buf)
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

func TestButtonSizes(t *testing.T) {
	tests := []struct {
		name   string
		props  ButtonProps
		expect string
	}{
		{
			name:   "Small size",
			props:  ButtonProps{Size: IsSmall},
			expect: `<button type="button" class="button is-small"></button>`,
		},
		{
			name:   "Normal size",
			props:  ButtonProps{Size: IsNormal},
			expect: `<button type="button" class="button is-normal"></button>`,
		},
		{
			name:   "Medium size",
			props:  ButtonProps{Size: IsMedium},
			expect: `<button type="button" class="button is-medium"></button>`,
		},
		{
			name:   "Large size",
			props:  ButtonProps{Size: IsLarge},
			expect: `<button type="button" class="button is-large"></button>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Button(tt.props).Render(context.Background(), &buf)
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

func TestButtonColors(t *testing.T) {
	tests := []struct {
		name   string
		props  ButtonProps
		expect string
	}{
		{
			name:   "Black color",
			props:  ButtonProps{Color: IsBlack},
			expect: `<button type="button" class="button is-black"></button>`,
		},
		{
			name:   "Primary color",
			props:  ButtonProps{Color: IsPrimary},
			expect: `<button type="button" class="button is-primary"></button>`,
		},
		{
			name:   "Danger color",
			props:  ButtonProps{Color: IsDanger},
			expect: `<button type="button" class="button is-danger"></button>`,
		},
		{
			name:   "Success color",
			props:  ButtonProps{Color: IsSuccess},
			expect: `<button type="button" class="button is-success"></button>`,
		},
		{
			name:   "Warning color",
			props:  ButtonProps{Color: IsWarning},
			expect: `<button type="button" class="button is-warning"></button>`,
		},
		{
			name:   "Info color",
			props:  ButtonProps{Color: IsInfo},
			expect: `<button type="button" class="button is-info"></button>`,
		},
		{
			name:   "Light color",
			props:  ButtonProps{Color: IsLight},
			expect: `<button type="button" class="button is-light"></button>`,
		},
		{
			name:   "Dark color",
			props:  ButtonProps{Color: IsDark},
			expect: `<button type="button" class="button is-dark"></button>`,
		},
		{
			name:   "Text color",
			props:  ButtonProps{Color: IsText},
			expect: `<button type="button" class="button is-text"></button>`,
		},
		{
			name:   "White color",
			props:  ButtonProps{Color: IsWhite},
			expect: `<button type="button" class="button is-white"></button>`,
		},
		{
			name:   "Link color",
			props:  ButtonProps{Color: IsLink},
			expect: `<button type="button" class="button is-link"></button>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Button(tt.props).Render(context.Background(), &buf)
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

func TestButtons(t *testing.T) {
	tests := []struct {
		name   string
		props  ButtonsProps
		expect string
	}{
		{
			name:   "Default",
			props:  ButtonsProps{},
			expect: `<div class="buttons"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  ButtonsProps{ID: "buttons1", Class: []string{"custom-buttons"}},
			expect: `<div id="buttons1" class="buttons custom-buttons"></div>`,
		},
		{
			name:   "With small size",
			props:  ButtonsProps{Size: AreSmall},
			expect: `<div class="buttons are-small"></div>`,
		},
		{
			name:   "With medium size",
			props:  ButtonsProps{Size: AreMedium},
			expect: `<div class="buttons are-medium"></div>`,
		},
		{
			name:   "With large size",
			props:  ButtonsProps{Size: AreLarge},
			expect: `<div class="buttons are-large"></div>`,
		},
		{
			name:   "With centered alignment",
			props:  ButtonsProps{Alignment: IsCentered},
			expect: `<div class="buttons is-centered"></div>`,
		},
		{
			name:   "With right alignment",
			props:  ButtonsProps{Alignment: IsRight},
			expect: `<div class="buttons is-right"></div>`,
		},
		{
			name:   "With addons",
			props:  ButtonsProps{HasAddons: true},
			expect: `<div class="buttons has-addons"></div>`,
		},
		{
			name:   "With size and alignment",
			props:  ButtonsProps{Size: AreLarge, Alignment: IsCentered},
			expect: `<div class="buttons are-large is-centered"></div>`,
		},
		{
			name: "All fields combined",
			props: ButtonsProps{
				ID:         "main-buttons",
				Class:      []string{"foo", "bar"},
				Size:       AreMedium,
				Alignment:  IsRight,
				HasAddons:  true,
				Attributes: templ.Attributes{"data-group": "actions"},
			},
			expect: `<div id="main-buttons" class="buttons are-medium is-right has-addons foo bar" data-group="actions"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Buttons(tt.props).Render(context.Background(), &buf)
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
