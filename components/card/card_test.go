package card

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestCard(t *testing.T) {
	tests := []struct {
		name   string
		props  CardProps
		expect string
	}{
		{
			name:   "Default",
			props:  CardProps{},
			expect: `<div class="card"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  CardProps{ID: "card1", Class: []string{"custom-card", "special"}},
			expect: `<div id="card1" class="card custom-card special"></div>`,
		},
		{
			name: "All fields combined",
			props: CardProps{
				ID:         "main-card",
				Class:      []string{"foo", "bar"},
				Attributes: templ.Attributes{"data-role": "container"},
			},
			expect: `<div id="main-card" class="card foo bar" data-role="container"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Card(tt.props).Render(context.Background(), &buf)
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

func TestCardHeader(t *testing.T) {
	tests := []struct {
		name   string
		props  CardHeaderProps
		expect string
	}{
		{
			name:   "Default",
			props:  CardHeaderProps{},
			expect: `<header class="card-header"></header>`,
		},
		{
			name:   "With ID and custom classes",
			props:  CardHeaderProps{ID: "header1", Class: []string{"custom-header"}},
			expect: `<header id="header1" class="card-header custom-header"></header>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := CardHeader(tt.props).Render(context.Background(), &buf)
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

func TestCardHeaderTitle(t *testing.T) {
	tests := []struct {
		name   string
		props  CardHeaderTitleProps
		expect string
	}{
		{
			name:   "Default",
			props:  CardHeaderTitleProps{},
			expect: `<p class="card-header-title"></p>`,
		},
		{
			name:   "With centered",
			props:  CardHeaderTitleProps{IsCentered: true},
			expect: `<p class="card-header-title is-centered"></p>`,
		},
		{
			name: "All fields combined",
			props: CardHeaderTitleProps{
				ID:         "title1",
				IsCentered: true,
				Class:      []string{"custom-title"},
				Attributes: templ.Attributes{"data-test": "title"},
			},
			expect: `<p id="title1" class="card-header-title is-centered custom-title" data-test="title"></p>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := CardHeaderTitle(tt.props).Render(context.Background(), &buf)
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

func TestCardHeaderIcon(t *testing.T) {
	tests := []struct {
		name   string
		props  CardHeaderIconProps
		expect string
	}{
		{
			name:   "Default",
			props:  CardHeaderIconProps{},
			expect: `<span class="card-header-icon"></span>`,
		},
		{
			name:   "With ID and custom classes",
			props:  CardHeaderIconProps{ID: "icon1", Class: []string{"fa", "fa-angle-down"}},
			expect: `<span id="icon1" class="card-header-icon fa fa-angle-down"></span>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := CardHeaderIcon(tt.props).Render(context.Background(), &buf)
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

func TestCardContent(t *testing.T) {
	tests := []struct {
		name   string
		props  CardContentProps
		expect string
	}{
		{
			name:   "Default",
			props:  CardContentProps{},
			expect: `<div class="card-content"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  CardContentProps{ID: "content1", Class: []string{"custom-content"}},
			expect: `<div id="content1" class="card-content custom-content"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := CardContent(tt.props).Render(context.Background(), &buf)
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

func TestCardImage(t *testing.T) {
	tests := []struct {
		name   string
		props  CardImageProps
		expect string
	}{
		{
			name:   "Default",
			props:  CardImageProps{},
			expect: `<div class="card-image"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  CardImageProps{ID: "image1", Class: []string{"custom-image"}},
			expect: `<div id="image1" class="card-image custom-image"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := CardImage(tt.props).Render(context.Background(), &buf)
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

func TestCardFooter(t *testing.T) {
	tests := []struct {
		name   string
		props  CardFooterProps
		expect string
	}{
		{
			name:   "Default",
			props:  CardFooterProps{},
			expect: `<footer class="card-footer"></footer>`,
		},
		{
			name:   "With ID and custom classes",
			props:  CardFooterProps{ID: "footer1", Class: []string{"custom-footer"}},
			expect: `<footer id="footer1" class="card-footer custom-footer"></footer>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := CardFooter(tt.props).Render(context.Background(), &buf)
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

func TestCardFooterItem(t *testing.T) {
	tests := []struct {
		name   string
		props  CardFooterItemProps
		expect string
	}{
		{
			name:   "Default",
			props:  CardFooterItemProps{},
			expect: `<a class="card-footer-item"></a>`,
		},
		{
			name:   "With ID and custom classes",
			props:  CardFooterItemProps{ID: "item1", Class: []string{"custom-item"}},
			expect: `<a id="item1" class="card-footer-item custom-item"></a>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := CardFooterItem(tt.props).Render(context.Background(), &buf)
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
