package tag

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestTag(t *testing.T) {
	tests := []struct {
		name   string
		props  TagProps
		expect string
	}{
		{
			name:   "Default (span)",
			props:  TagProps{},
			expect: `<span class="tag"></span>`,
		},
		{
			name:   "With ID and custom classes",
			props:  TagProps{ID: "tag1", Class: []string{"custom-tag"}},
			expect: `<span id="tag1" class="tag custom-tag"></span>`,
		},
		{
			name:   "As button",
			props:  TagProps{IsButton: true},
			expect: `<button class="tag"></button>`,
		},
		{
			name:   "As anchor",
			props:  TagProps{IsAnchor: true},
			expect: `<a class="tag"></a>`,
		},
		{
			name:   "With color",
			props:  TagProps{Color: IsPrimary},
			expect: `<span class="tag is-primary"></span>`,
		},
		{
			name:   "With light variant",
			props:  TagProps{Color: IsDanger, IsLightVariant: true},
			expect: `<span class="tag is-danger is-light"></span>`,
		},
		{
			name:   "With size",
			props:  TagProps{Size: IsLarge},
			expect: `<span class="tag is-large"></span>`,
		},
		{
			name:   "With hoverable state",
			props:  TagProps{IsHoverable: true},
			expect: `<span class="tag is-hoverable"></span>`,
		},
		{
			name:   "With rounded style",
			props:  TagProps{IsRounded: true},
			expect: `<span class="tag is-rounded"></span>`,
		},
		{
			name:   "Delete tag",
			props:  TagProps{IsDelete: true},
			expect: `<span class="tag is-delete"></span>`,
		},
		{
			name:   "With delete button (default aria-label)",
			props:  TagProps{HasDelete: true},
			expect: `<span class="tag"><button class="delete is-small" aria-label="Remove tag"></button></span>`,
		},
		{
			name:   "With delete button (custom aria-label)",
			props:  TagProps{HasDelete: true, DeleteAriaLabel: "Remove filter"},
			expect: `<span class="tag"><button class="delete is-small" aria-label="Remove filter"></button></span>`,
		},
		{
			name:   "Delete tag with content suppressed",
			props:  TagProps{IsDelete: true},
			expect: `<span class="tag is-delete"></span>`,
		},
		{
			name:   "Button with color and size",
			props:  TagProps{IsButton: true, Color: IsSuccess, Size: IsMedium},
			expect: `<button class="tag is-medium is-success"></button>`,
		},
		{
			name:   "Anchor with hover and rounded",
			props:  TagProps{IsAnchor: true, IsHoverable: true, IsRounded: true},
			expect: `<a class="tag is-rounded is-hoverable"></a>`,
		},
		{
			name: "All fields combined",
			props: TagProps{
				ID:              "main-tag",
				Class:           []string{"foo", "bar"},
				Color:           IsInfo,
				IsLightVariant:  true,
				Size:            IsLarge,
				IsRounded:       true,
				IsHoverable:     true,
				HasDelete:       true,
				DeleteAriaLabel: "Remove item tag",
				Attributes:      templ.Attributes{"data-value": "react"},
			},
			expect: `<span id="main-tag" class="tag is-large is-info is-light is-rounded is-hoverable foo bar" data-value="react"><button class="delete is-small" aria-label="Remove item tag"></button></span>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Tag(tt.props).Render(context.Background(), &buf)
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

func TestTagSizes(t *testing.T) {
	tests := []struct {
		name   string
		size   Size
		expect string
	}{
		{"Normal", IsNormal, `<span class="tag is-normal"></span>`},
		{"Medium", IsMedium, `<span class="tag is-medium"></span>`},
		{"Large", IsLarge, `<span class="tag is-large"></span>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Tag(TagProps{Size: tt.size}).Render(context.Background(), &buf)
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

func TestTagColors(t *testing.T) {
	tests := []struct {
		name   string
		color  Color
		expect string
	}{
		{"Black", IsBlack, `<span class="tag is-black"></span>`},
		{"Danger", IsDanger, `<span class="tag is-danger"></span>`},
		{"Dark", IsDark, `<span class="tag is-dark"></span>`},
		{"Info", IsInfo, `<span class="tag is-info"></span>`},
		{"Light", IsLight, `<span class="tag is-light"></span>`},
		{"Link", IsLink, `<span class="tag is-link"></span>`},
		{"Primary", IsPrimary, `<span class="tag is-primary"></span>`},
		{"Success", IsSuccess, `<span class="tag is-success"></span>`},
		{"Text", IsText, `<span class="tag is-text"></span>`},
		{"Warning", IsWarning, `<span class="tag is-warning"></span>`},
		{"White", IsWhite, `<span class="tag is-white"></span>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Tag(TagProps{Color: tt.color}).Render(context.Background(), &buf)
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

func TestTagElements(t *testing.T) {
	tests := []struct {
		name   string
		props  TagProps
		expect string
	}{
		{
			name:   "Default span element",
			props:  TagProps{},
			expect: `<span class="tag"></span>`,
		},
		{
			name:   "Button element",
			props:  TagProps{IsButton: true, Color: IsPrimary},
			expect: `<button class="tag is-primary"></button>`,
		},
		{
			name:   "Anchor element",
			props:  TagProps{IsAnchor: true, Color: IsSuccess},
			expect: `<a class="tag is-success"></a>`,
		},
		{
			name:   "Button with delete",
			props:  TagProps{IsButton: true, HasDelete: true},
			expect: `<button class="tag"><button class="delete is-small" aria-label="Remove tag"></button></button>`,
		},
		{
			name:   "Anchor with delete",
			props:  TagProps{IsAnchor: true, HasDelete: true, DeleteAriaLabel: "Remove link"},
			expect: `<a class="tag"><button class="delete is-small" aria-label="Remove link"></button></a>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Tag(tt.props).Render(context.Background(), &buf)
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

func TestTags(t *testing.T) {
	tests := []struct {
		name   string
		props  TagsProps
		expect string
	}{
		{
			name:   "Default",
			props:  TagsProps{},
			expect: `<div class="tags"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  TagsProps{ID: "tags1", Class: []string{"custom-tags"}},
			expect: `<div id="tags1" class="tags custom-tags"></div>`,
		},
		{
			name:   "With medium size",
			props:  TagsProps{Size: AreMedium},
			expect: `<div class="tags are-medium"></div>`,
		},
		{
			name:   "With large size",
			props:  TagsProps{Size: AreLarge},
			expect: `<div class="tags are-large"></div>`,
		},
		{
			name:   "With centered alignment",
			props:  TagsProps{Alignment: IsCentered},
			expect: `<div class="tags is-centered"></div>`,
		},
		{
			name:   "With right alignment",
			props:  TagsProps{Alignment: IsRight},
			expect: `<div class="tags is-right"></div>`,
		},
		{
			name:   "With addons",
			props:  TagsProps{HasAddons: true},
			expect: `<div class="tags has-addons"></div>`,
		},
		{
			name:   "With size and alignment",
			props:  TagsProps{Size: AreLarge, Alignment: IsCentered},
			expect: `<div class="tags are-large is-centered"></div>`,
		},
		{
			name: "All fields combined",
			props: TagsProps{
				ID:         "main-tags",
				Class:      []string{"foo", "bar"},
				Size:       AreMedium,
				Alignment:  IsRight,
				HasAddons:  true,
				Attributes: templ.Attributes{"data-group": "filters"},
			},
			expect: `<div id="main-tags" class="tags are-medium is-right has-addons foo bar" data-group="filters"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Tags(tt.props).Render(context.Background(), &buf)
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

func TestTagLightVariants(t *testing.T) {
	tests := []struct {
		name   string
		color  Color
		expect string
	}{
		{
			name:   "Primary light",
			color:  IsPrimary,
			expect: `<span class="tag is-primary is-light"></span>`,
		},
		{
			name:   "Danger light",
			color:  IsDanger,
			expect: `<span class="tag is-danger is-light"></span>`,
		},
		{
			name:   "Success light",
			color:  IsSuccess,
			expect: `<span class="tag is-success is-light"></span>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Tag(TagProps{Color: tt.color, IsLightVariant: true}).Render(context.Background(), &buf)
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

func TestTagDeleteFunctionality(t *testing.T) {
	tests := []struct {
		name   string
		props  TagProps
		expect string
	}{
		{
			name:   "Delete tag (no content)",
			props:  TagProps{IsDelete: true},
			expect: `<span class="tag is-delete"></span>`,
		},
		{
			name:   "Tag with delete button (default aria)",
			props:  TagProps{HasDelete: true},
			expect: `<span class="tag"><button class="delete is-small" aria-label="Remove tag"></button></span>`,
		},
		{
			name:   "Tag with delete button (custom aria)",
			props:  TagProps{HasDelete: true, DeleteAriaLabel: "Remove JavaScript tag"},
			expect: `<span class="tag"><button class="delete is-small" aria-label="Remove JavaScript tag"></button></span>`,
		},
		{
			name:   "Delete tag cannot have delete button",
			props:  TagProps{IsDelete: true, HasDelete: true},
			expect: `<span class="tag is-delete"><button class="delete is-small" aria-label="Remove tag"></button></span>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Tag(tt.props).Render(context.Background(), &buf)
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
