package form

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestIcon(t *testing.T) {
	tests := []struct {
		name   string
		props  IconProps
		expect string
	}{
		{
			name:   "Default",
			props:  IconProps{},
			expect: `<span class="icon"></span>`,
		},
		{
			name:   "With ID and classes",
			props:  IconProps{ID: "test-icon", Class: []string{"custom", "icon-wrapper"}},
			expect: `<span id="test-icon" class="icon custom icon-wrapper"></span>`,
		},
		{
			name:   "With left position",
			props:  IconProps{Position: IsLeft},
			expect: `<span class="icon is-left"></span>`,
		},
		{
			name:   "With right position",
			props:  IconProps{Position: IsRight},
			expect: `<span class="icon is-right"></span>`,
		},
		{
			name:   "With small size",
			props:  IconProps{Size: IsSmall},
			expect: `<span class="icon is-small"></span>`,
		},
		{
			name:   "With medium size",
			props:  IconProps{Size: IsMedium},
			expect: `<span class="icon is-medium"></span>`,
		},
		{
			name:   "With large size",
			props:  IconProps{Size: IsLarge},
			expect: `<span class="icon is-large"></span>`,
		},
		{
			name:   "Left position with size",
			props:  IconProps{Position: IsLeft, Size: IsLarge},
			expect: `<span class="icon is-left is-large"></span>`,
		},
		{
			name:   "Right position with size",
			props:  IconProps{Position: IsRight, Size: IsSmall},
			expect: `<span class="icon is-right is-small"></span>`,
		},
		{
			name: "Complete configuration",
			props: IconProps{
				ID:       "main-icon",
				Class:    []string{"fa", "fa-user"},
				Position: IsLeft,
				Size:     IsMedium,
			},
			expect: `<span id="main-icon" class="icon is-left is-medium fa fa-user"></span>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Icon(tt.props).Render(context.Background(), &buf)
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

func TestIconModifiers(t *testing.T) {
	tests := []struct {
		name             string
		props            IconProps
		expectPosition   bool
		expectSize       bool
		expectedPosition string
		expectedSize     string
	}{
		{"No modifiers", IconProps{}, false, false, "", ""},
		{"Left position", IconProps{Position: IsLeft}, true, false, "is-left", ""},
		{"Right position", IconProps{Position: IsRight}, true, false, "is-right", ""},
		{"Small size", IconProps{Size: IsSmall}, false, true, "", "is-small"},
		{"Large size", IconProps{Size: IsLarge}, false, true, "", "is-large"},
		{"Left and medium", IconProps{Position: IsLeft, Size: IsMedium}, true, true, "is-left", "is-medium"},
		{"Right and small", IconProps{Position: IsRight, Size: IsSmall}, true, true, "is-right", "is-small"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Icon(tt.props).Render(context.Background(), &buf)
			if err != nil {
				t.Fatalf("render failed: %v", err)
			}
			got := strings.TrimSpace(buf.String())

			hasPosition := strings.Contains(got, "is-left") || strings.Contains(got, "is-right")
			hasSize := strings.Contains(got, "is-small") || strings.Contains(got, "is-medium") || strings.Contains(got, "is-large")

			if hasPosition != tt.expectPosition {
				t.Errorf("expected position modifier: %v, got: %v", tt.expectPosition, hasPosition)
			}
			if hasSize != tt.expectSize {
				t.Errorf("expected size modifier: %v, got: %v", tt.expectSize, hasSize)
			}

			if tt.expectedPosition != "" && !strings.Contains(got, tt.expectedPosition) {
				t.Errorf("expected %s to be present in: %s", tt.expectedPosition, got)
			}
			if tt.expectedSize != "" && !strings.Contains(got, tt.expectedSize) {
				t.Errorf("expected %s to be present in: %s", tt.expectedSize, got)
			}
		})
	}
}

func TestIconAttributes(t *testing.T) {
	props := IconProps{
		Attributes: templ.Attributes{
			"data-testid": "icon",
			"aria-hidden": "true",
		},
	}

	var buf strings.Builder
	err := Icon(props).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}
	got := strings.TrimSpace(buf.String())

	for key, value := range props.Attributes {
		expectedAttr := key + `="` + fmt.Sprintf("%v", value) + `"`
		if !strings.Contains(got, expectedAttr) {
			t.Errorf("expected attribute %s to be present in: %s", expectedAttr, got)
		}
	}
}
