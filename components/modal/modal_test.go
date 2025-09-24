package modal

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestModal(t *testing.T) {
	tests := []struct {
		name   string
		props  ModalProps
		expect string
	}{
		{
			name:   "Default",
			props:  ModalProps{},
			expect: `<div class="modal"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  ModalProps{ID: "modal1", Class: []string{"custom-modal"}},
			expect: `<div id="modal1" class="modal custom-modal"></div>`,
		},
		{
			name:   "With size and active state",
			props:  ModalProps{Size: IsLarge, IsActive: true},
			expect: `<div class="modal is-large is-active"></div>`,
		},
		{
			name: "All fields combined",
			props: ModalProps{
				ID:         "main-modal",
				Class:      []string{"foo", "bar"},
				Size:       IsMedium,
				IsActive:   true,
				Attributes: templ.Attributes{"data-role": "dialog"},
			},
			expect: `<div id="main-modal" class="modal is-medium is-active foo bar" data-role="dialog"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Modal(tt.props).Render(context.Background(), &buf)
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

func TestModalBackground(t *testing.T) {
	tests := []struct {
		name   string
		props  ModalBackgroundProps
		expect string
	}{
		{
			name:   "Default",
			props:  ModalBackgroundProps{},
			expect: `<div class="modal-background"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  ModalBackgroundProps{ID: "bg1", Class: []string{"custom-bg"}},
			expect: `<div id="bg1" class="modal-background custom-bg"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := ModalBackground(tt.props).Render(context.Background(), &buf)
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

func TestModalContent(t *testing.T) {
	tests := []struct {
		name   string
		props  ModalContentProps
		expect string
	}{
		{
			name:   "Default",
			props:  ModalContentProps{},
			expect: `<div class="modal-content"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  ModalContentProps{ID: "content1", Class: []string{"custom-content"}},
			expect: `<div id="content1" class="modal-content custom-content"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := ModalContent(tt.props).Render(context.Background(), &buf)
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

func TestModalClose(t *testing.T) {
	tests := []struct {
		name   string
		props  ModalCloseProps
		expect string
	}{
		{
			name:   "Default",
			props:  ModalCloseProps{},
			expect: `<button class="modal-close"></button>`,
		},
		{
			name:   "With ID and custom classes",
			props:  ModalCloseProps{ID: "close1", Class: []string{"custom-close"}},
			expect: `<button id="close1" class="modal-close custom-close"></button>`,
		},
		{
			name: "With attributes",
			props: ModalCloseProps{
				Attributes: templ.Attributes{"aria-label": "close"},
			},
			expect: `<button class="modal-close" aria-label="close"></button>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := ModalClose(tt.props).Render(context.Background(), &buf)
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

func TestModalCard(t *testing.T) {
	tests := []struct {
		name   string
		props  ModalCardProps
		expect string
	}{
		{
			name:   "Default",
			props:  ModalCardProps{},
			expect: `<div class="modal-card"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  ModalCardProps{ID: "card1", Class: []string{"custom-card"}},
			expect: `<div id="card1" class="modal-card custom-card"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := ModalCard(tt.props).Render(context.Background(), &buf)
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

func TestModalCardHead(t *testing.T) {
	tests := []struct {
		name   string
		props  ModalCardHeadProps
		expect string
	}{
		{
			name:   "Default",
			props:  ModalCardHeadProps{},
			expect: `<header class="modal-card-head"></header>`,
		},
		{
			name:   "With ID and custom classes",
			props:  ModalCardHeadProps{ID: "head1", Class: []string{"custom-head"}},
			expect: `<header id="head1" class="modal-card-head custom-head"></header>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := ModalCardHead(tt.props).Render(context.Background(), &buf)
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

func TestModalCardTitle(t *testing.T) {
	tests := []struct {
		name   string
		props  ModalCardTitleProps
		expect string
	}{
		{
			name:   "Default",
			props:  ModalCardTitleProps{},
			expect: `<p class="modal-card-title"></p>`,
		},
		{
			name:   "With ID and custom classes",
			props:  ModalCardTitleProps{ID: "title1", Class: []string{"custom-title"}},
			expect: `<p id="title1" class="modal-card-title custom-title"></p>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := ModalCardTitle(tt.props).Render(context.Background(), &buf)
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

func TestModalCardBody(t *testing.T) {
	tests := []struct {
		name   string
		props  ModalCardBodyProps
		expect string
	}{
		{
			name:   "Default",
			props:  ModalCardBodyProps{},
			expect: `<section class="modal-card-body"></section>`,
		},
		{
			name:   "With ID and custom classes",
			props:  ModalCardBodyProps{ID: "body1", Class: []string{"custom-body"}},
			expect: `<section id="body1" class="modal-card-body custom-body"></section>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := ModalCardBody(tt.props).Render(context.Background(), &buf)
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

func TestModalCardFoot(t *testing.T) {
	tests := []struct {
		name   string
		props  ModalCardFootProps
		expect string
	}{
		{
			name:   "Default",
			props:  ModalCardFootProps{},
			expect: `<footer class="modal-card-foot"></footer>`,
		},
		{
			name:   "With ID and custom classes",
			props:  ModalCardFootProps{ID: "foot1", Class: []string{"custom-foot"}},
			expect: `<footer id="foot1" class="modal-card-foot custom-foot"></footer>`,
		},
		{
			name: "With attributes",
			props: ModalCardFootProps{
				ID:         "footer",
				Attributes: templ.Attributes{"data-actions": "confirm"},
			},
			expect: `<footer id="footer" class="modal-card-foot" data-actions="confirm"></footer>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := ModalCardFoot(tt.props).Render(context.Background(), &buf)
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
