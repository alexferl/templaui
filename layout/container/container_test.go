package container

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestContainer(t *testing.T) {
	tests := []struct {
		name   string
		props  ContainerProps
		expect string
	}{
		{
			name:   "Default",
			props:  ContainerProps{},
			expect: `<div class="container"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  ContainerProps{ID: "container1", Class: []string{"custom-container"}},
			expect: `<div id="container1" class="container custom-container"></div>`,
		},
		{
			name:   "With fluid responsive",
			props:  ContainerProps{Responsive: IsFluid},
			expect: `<div class="container is-fluid"></div>`,
		},
		{
			name:   "With fullhd responsive",
			props:  ContainerProps{Responsive: IsFullHD},
			expect: `<div class="container is-fullhd"></div>`,
		},
		{
			name:   "With max-desktop responsive",
			props:  ContainerProps{Responsive: IsMaxDesktop},
			expect: `<div class="container is-max-desktop"></div>`,
		},
		{
			name:   "With max-tablet responsive",
			props:  ContainerProps{Responsive: IsMaxTablet},
			expect: `<div class="container is-max-tablet"></div>`,
		},
		{
			name:   "With max-widescreen responsive",
			props:  ContainerProps{Responsive: IsMaxWidescreen},
			expect: `<div class="container is-max-widescreen"></div>`,
		},
		{
			name:   "With widescreen responsive",
			props:  ContainerProps{Responsive: IsWidescreen},
			expect: `<div class="container is-widescreen"></div>`,
		},
		{
			name: "All fields combined",
			props: ContainerProps{
				ID:         "main-container",
				Class:      []string{"foo", "bar"},
				Responsive: IsFluid,
				Attributes: templ.Attributes{"data-role": "main"},
			},
			expect: `<div id="main-container" class="container is-fluid foo bar" data-role="main"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Container(tt.props).Render(context.Background(), &buf)
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
