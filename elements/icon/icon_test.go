package icon

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
	"github.com/alexferl/templaui/helpers"
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
			name:   "With ID and custom classes",
			props:  IconProps{ID: "icon1", Class: []string{"custom-icon"}},
			expect: `<span id="icon1" class="icon custom-icon"></span>`,
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
			name:   "With color",
			props:  IconProps{Color: helpers.IsPrimary},
			expect: `<span class="icon is-primary"></span>`,
		},
		{
			name:   "With danger color",
			props:  IconProps{Color: helpers.IsDanger},
			expect: `<span class="icon is-danger"></span>`,
		},
		{
			name:   "With color and size",
			props:  IconProps{Color: helpers.IsSuccess, Size: IsLarge},
			expect: `<span class="icon is-success is-large"></span>`,
		},
		{
			name: "All fields combined",
			props: IconProps{
				ID:         "main-icon",
				Class:      []string{"foo", "bar"},
				Color:      helpers.IsWarning,
				Size:       IsMedium,
				Attributes: templ.Attributes{"aria-hidden": "true"},
			},
			expect: `<span id="main-icon" class="icon is-warning is-medium foo bar" aria-hidden="true"></span>`,
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

func TestIconText(t *testing.T) {
	tests := []struct {
		name   string
		props  IconTextProps
		expect string
	}{
		{
			name:   "Default",
			props:  IconTextProps{},
			expect: `<span class="icon-text"></span>`,
		},
		{
			name:   "With ID and custom classes",
			props:  IconTextProps{ID: "icontext1", Class: []string{"custom-icontext"}},
			expect: `<span id="icontext1" class="icon-text custom-icontext"></span>`,
		},
		{
			name:   "With size",
			props:  IconTextProps{Size: IsLarge},
			expect: `<span class="icon-text is-large"></span>`,
		},
		{
			name:   "With color",
			props:  IconTextProps{Color: helpers.IsInfo},
			expect: `<span class="icon-text is-info"></span>`,
		},
		{
			name:   "With color and size",
			props:  IconTextProps{Color: helpers.IsDark, Size: IsSmall},
			expect: `<span class="icon-text is-dark is-small"></span>`,
		},
		{
			name: "All fields combined",
			props: IconTextProps{
				ID:         "text-icon",
				Class:      []string{"inline", "aligned"},
				Color:      helpers.IsLink,
				Size:       IsMedium,
				Attributes: templ.Attributes{"data-inline": "true"},
			},
			expect: `<span id="text-icon" class="icon-text is-link is-medium inline aligned" data-inline="true"></span>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := IconText(tt.props).Render(context.Background(), &buf)
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
