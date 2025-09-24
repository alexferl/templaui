package box

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestBox(t *testing.T) {
	tests := []struct {
		name   string
		props  BoxProps
		expect string
	}{
		{
			name:   "Default",
			props:  BoxProps{},
			expect: `<div class="box"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  BoxProps{ID: "box1", Class: []string{"custom-box"}},
			expect: `<div id="box1" class="box custom-box"></div>`,
		},
		{
			name: "With multiple classes",
			props: BoxProps{
				Class: []string{"content-box", "elevated"},
			},
			expect: `<div class="box content-box elevated"></div>`,
		},
		{
			name: "With attributes",
			props: BoxProps{
				ID:         "main-box",
				Attributes: templ.Attributes{"data-content": "widget"},
			},
			expect: `<div id="main-box" class="box" data-content="widget"></div>`,
		},
		{
			name: "All fields combined",
			props: BoxProps{
				ID:         "content-box",
				Class:      []string{"foo", "bar"},
				Attributes: templ.Attributes{"data-role": "container", "data-type": "card"},
			},
			expect: `<div id="content-box" class="box foo bar" data-role="container" data-type="card"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Box(tt.props).Render(context.Background(), &buf)
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
