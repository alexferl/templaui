package block

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestBlock(t *testing.T) {
	tests := []struct {
		name   string
		props  BlockProps
		expect string
	}{
		{
			name:   "Default",
			props:  BlockProps{},
			expect: `<div class="block"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  BlockProps{ID: "block1", Class: []string{"custom-block"}},
			expect: `<div id="block1" class="block custom-block"></div>`,
		},
		{
			name: "With multiple classes",
			props: BlockProps{
				Class: []string{"content-block", "spaced"},
			},
			expect: `<div class="block content-block spaced"></div>`,
		},
		{
			name: "With attributes",
			props: BlockProps{
				ID:         "main-block",
				Attributes: templ.Attributes{"data-spacing": "bottom"},
			},
			expect: `<div id="main-block" class="block" data-spacing="bottom"></div>`,
		},
		{
			name: "All fields combined",
			props: BlockProps{
				ID:         "content-block",
				Class:      []string{"foo", "bar"},
				Attributes: templ.Attributes{"data-role": "content", "data-spacing": "margin"},
			},
			expect: `<div id="content-block" class="block foo bar" data-role="content" data-spacing="margin"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Block(tt.props).Render(context.Background(), &buf)
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
