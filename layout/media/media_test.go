package media

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestMedia(t *testing.T) {
	tests := []struct {
		name   string
		props  MediaProps
		expect string
	}{
		{
			name:   "Default",
			props:  MediaProps{},
			expect: `<article class="media"></article>`,
		},
		{
			name:   "With ID and custom classes",
			props:  MediaProps{ID: "media1", Class: []string{"custom-media"}},
			expect: `<article id="media1" class="media custom-media"></article>`,
		},
		{
			name:   "With large spacing",
			props:  MediaProps{IsLarge: true},
			expect: `<article class="media is-large"></article>`,
		},
		{
			name: "All fields combined",
			props: MediaProps{
				ID:         "main-media",
				Class:      []string{"foo", "bar"},
				IsLarge:    true,
				Attributes: templ.Attributes{"data-role": "comment"},
			},
			expect: `<article id="main-media" class="media is-large foo bar" data-role="comment"></article>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Media(tt.props).Render(context.Background(), &buf)
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

func TestMediaLeft(t *testing.T) {
	tests := []struct {
		name   string
		props  MediaLeftProps
		expect string
	}{
		{
			name:   "Default",
			props:  MediaLeftProps{},
			expect: `<div class="media-left"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  MediaLeftProps{ID: "left1", Class: []string{"custom-left"}},
			expect: `<div id="left1" class="media-left custom-left"></div>`,
		},
		{
			name: "With attributes",
			props: MediaLeftProps{
				ID:         "avatar-section",
				Attributes: templ.Attributes{"data-align": "left"},
			},
			expect: `<div id="avatar-section" class="media-left" data-align="left"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := MediaLeft(tt.props).Render(context.Background(), &buf)
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

func TestMediaContent(t *testing.T) {
	tests := []struct {
		name   string
		props  MediaContentProps
		expect string
	}{
		{
			name:   "Default",
			props:  MediaContentProps{},
			expect: `<div class="media-content"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  MediaContentProps{ID: "content1", Class: []string{"custom-content"}},
			expect: `<div id="content1" class="media-content custom-content"></div>`,
		},
		{
			name: "With multiple classes and attributes",
			props: MediaContentProps{
				ID:         "main-content",
				Class:      []string{"text-content", "expandable"},
				Attributes: templ.Attributes{"data-expandable": "true"},
			},
			expect: `<div id="main-content" class="media-content text-content expandable" data-expandable="true"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := MediaContent(tt.props).Render(context.Background(), &buf)
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

func TestMediaRight(t *testing.T) {
	tests := []struct {
		name   string
		props  MediaRightProps
		expect string
	}{
		{
			name:   "Default",
			props:  MediaRightProps{},
			expect: `<div class="media-right"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  MediaRightProps{ID: "right1", Class: []string{"custom-right"}},
			expect: `<div id="right1" class="media-right custom-right"></div>`,
		},
		{
			name: "With attributes",
			props: MediaRightProps{
				ID:         "actions-section",
				Class:      []string{"actions"},
				Attributes: templ.Attributes{"data-actions": "reply,delete"},
			},
			expect: `<div id="actions-section" class="media-right actions" data-actions="reply,delete"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := MediaRight(tt.props).Render(context.Background(), &buf)
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
