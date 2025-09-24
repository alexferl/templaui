package image

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestImage(t *testing.T) {
	tests := []struct {
		name   string
		props  ImageProps
		expect string
	}{
		{
			name:   "Default",
			props:  ImageProps{},
			expect: `<figure class="image"></figure>`,
		},
		{
			name:   "With ID and custom classes",
			props:  ImageProps{ID: "image1", Class: []string{"custom-image"}},
			expect: `<figure id="image1" class="image custom-image"></figure>`,
		},
		{
			name:   "With 16x16 size",
			props:  ImageProps{Size: Is16x16},
			expect: `<figure class="image is-16x16"></figure>`,
		},
		{
			name:   "With 24x24 size",
			props:  ImageProps{Size: Is24x24},
			expect: `<figure class="image is-24x24"></figure>`,
		},
		{
			name:   "With 32x32 size",
			props:  ImageProps{Size: Is32x32},
			expect: `<figure class="image is-32x32"></figure>`,
		},
		{
			name:   "With 48x48 size",
			props:  ImageProps{Size: Is48x48},
			expect: `<figure class="image is-48x48"></figure>`,
		},
		{
			name:   "With 64x64 size",
			props:  ImageProps{Size: Is64x64},
			expect: `<figure class="image is-64x64"></figure>`,
		},
		{
			name:   "With 96x96 size",
			props:  ImageProps{Size: Is96x96},
			expect: `<figure class="image is-96x96"></figure>`,
		},
		{
			name:   "With 128x128 size",
			props:  ImageProps{Size: Is128x128},
			expect: `<figure class="image is-128x128"></figure>`,
		},
		{
			name:   "With square ratio",
			props:  ImageProps{Ratio: IsSquare},
			expect: `<figure class="image is-square"></figure>`,
		},
		{
			name:   "With 1by1 ratio",
			props:  ImageProps{Ratio: Is1by1},
			expect: `<figure class="image is-1by1"></figure>`,
		},
		{
			name:   "With 16by9 ratio",
			props:  ImageProps{Ratio: Is16by9},
			expect: `<figure class="image is-16by9"></figure>`,
		},
		{
			name:   "With 4by3 ratio",
			props:  ImageProps{Ratio: Is4by3},
			expect: `<figure class="image is-4by3"></figure>`,
		},
		{
			name:   "With 3by2 ratio",
			props:  ImageProps{Ratio: Is3by2},
			expect: `<figure class="image is-3by2"></figure>`,
		},
		{
			name:   "With 9by16 ratio",
			props:  ImageProps{Ratio: Is9by16},
			expect: `<figure class="image is-9by16"></figure>`,
		},
		{
			name:   "With rounded style",
			props:  ImageProps{IsRounded: true},
			expect: `<figure class="image is-rounded"></figure>`,
		},
		{
			name:   "With fullwidth layout",
			props:  ImageProps{IsFullwidth: true},
			expect: `<figure class="image is-fullwidth"></figure>`,
		},
		{
			name:   "With has-ratio utility",
			props:  ImageProps{HasRatio: true},
			expect: `<figure class="image has-ratio"></figure>`,
		},
		{
			name:   "With size and rounded",
			props:  ImageProps{Size: Is64x64, IsRounded: true},
			expect: `<figure class="image is-64x64 is-rounded"></figure>`,
		},
		{
			name:   "With ratio and fullwidth",
			props:  ImageProps{Ratio: Is16by9, IsFullwidth: true},
			expect: `<figure class="image is-16by9 is-fullwidth"></figure>`,
		},
		{
			name:   "With ratio and has-ratio utility",
			props:  ImageProps{Ratio: Is4by3, HasRatio: true},
			expect: `<figure class="image is-4by3 has-ratio"></figure>`,
		},
		{
			name: "All fields combined",
			props: ImageProps{
				ID:          "main-image",
				Class:       []string{"foo", "bar"},
				Size:        Is128x128,
				Ratio:       Is1by1,
				IsRounded:   true,
				IsFullwidth: true,
				HasRatio:    true,
				Attributes:  templ.Attributes{"data-src": "avatar.jpg"},
			},
			expect: `<figure id="main-image" class="image is-128x128 is-1by1 is-rounded is-fullwidth has-ratio foo bar" data-src="avatar.jpg"></figure>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Image(tt.props).Render(context.Background(), &buf)
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

func TestImageSizes(t *testing.T) {
	tests := []struct {
		name   string
		size   Size
		expect string
	}{
		{"16x16", Is16x16, `<figure class="image is-16x16"></figure>`},
		{"24x24", Is24x24, `<figure class="image is-24x24"></figure>`},
		{"32x32", Is32x32, `<figure class="image is-32x32"></figure>`},
		{"48x48", Is48x48, `<figure class="image is-48x48"></figure>`},
		{"64x64", Is64x64, `<figure class="image is-64x64"></figure>`},
		{"96x96", Is96x96, `<figure class="image is-96x96"></figure>`},
		{"128x128", Is128x128, `<figure class="image is-128x128"></figure>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Image(ImageProps{Size: tt.size}).Render(context.Background(), &buf)
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

func TestImageRatios(t *testing.T) {
	tests := []struct {
		name   string
		ratio  Ratio
		expect string
	}{
		{"Square", IsSquare, `<figure class="image is-square"></figure>`},
		{"1by1", Is1by1, `<figure class="image is-1by1"></figure>`},
		{"1by2", Is1by2, `<figure class="image is-1by2"></figure>`},
		{"1by3", Is1by3, `<figure class="image is-1by3"></figure>`},
		{"2by1", Is2by1, `<figure class="image is-2by1"></figure>`},
		{"2by3", Is2by3, `<figure class="image is-2by3"></figure>`},
		{"3by1", Is3by1, `<figure class="image is-3by1"></figure>`},
		{"3by2", Is3by2, `<figure class="image is-3by2"></figure>`},
		{"3by4", Is3by4, `<figure class="image is-3by4"></figure>`},
		{"3by5", Is3by5, `<figure class="image is-3by5"></figure>`},
		{"4by3", Is4by3, `<figure class="image is-4by3"></figure>`},
		{"4by5", Is4by5, `<figure class="image is-4by5"></figure>`},
		{"5by3", Is5by3, `<figure class="image is-5by3"></figure>`},
		{"5by4", Is5by4, `<figure class="image is-5by4"></figure>`},
		{"9by16", Is9by16, `<figure class="image is-9by16"></figure>`},
		{"16by9", Is16by9, `<figure class="image is-16by9"></figure>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Image(ImageProps{Ratio: tt.ratio}).Render(context.Background(), &buf)
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

func TestImageModifiers(t *testing.T) {
	tests := []struct {
		name   string
		props  ImageProps
		expect string
	}{
		{
			name:   "Rounded style",
			props:  ImageProps{IsRounded: true},
			expect: `<figure class="image is-rounded"></figure>`,
		},
		{
			name:   "Fullwidth layout",
			props:  ImageProps{IsFullwidth: true},
			expect: `<figure class="image is-fullwidth"></figure>`,
		},
		{
			name:   "Has ratio utility",
			props:  ImageProps{HasRatio: true},
			expect: `<figure class="image has-ratio"></figure>`,
		},
		{
			name:   "All modifiers combined",
			props:  ImageProps{IsRounded: true, IsFullwidth: true, HasRatio: true},
			expect: `<figure class="image is-rounded is-fullwidth has-ratio"></figure>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Image(tt.props).Render(context.Background(), &buf)
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
