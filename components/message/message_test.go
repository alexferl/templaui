package message

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestMessage(t *testing.T) {
	tests := []struct {
		name   string
		props  MessageProps
		expect string
	}{
		{
			name:   "Default",
			props:  MessageProps{},
			expect: `<article class="message"></article>`,
		},
		{
			name:   "With ID and custom classes",
			props:  MessageProps{ID: "msg1", Class: []string{"custom-message"}},
			expect: `<article id="msg1" class="message custom-message"></article>`,
		},
		{
			name:   "With color and size",
			props:  MessageProps{Color: IsDanger, Size: IsLarge},
			expect: `<article class="message is-danger is-large"></article>`,
		},
		{
			name: "All fields combined",
			props: MessageProps{
				ID:         "main-message",
				Class:      []string{"foo", "bar"},
				Color:      IsSuccess,
				Size:       IsMedium,
				Attributes: templ.Attributes{"data-role": "alert"},
			},
			expect: `<article id="main-message" class="message is-success is-medium foo bar" data-role="alert"></article>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Message(tt.props).Render(context.Background(), &buf)
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

func TestMessageHeader(t *testing.T) {
	tests := []struct {
		name   string
		props  MessageHeaderProps
		expect string
	}{
		{
			name:   "Default",
			props:  MessageHeaderProps{},
			expect: `<div class="message-header"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  MessageHeaderProps{ID: "header1", Class: []string{"custom-header"}},
			expect: `<div id="header1" class="message-header custom-header"></div>`,
		},
		{
			name: "With attributes",
			props: MessageHeaderProps{
				ID:         "msg-header",
				Attributes: templ.Attributes{"data-dismiss": "message"},
			},
			expect: `<div id="msg-header" class="message-header" data-dismiss="message"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := MessageHeader(tt.props).Render(context.Background(), &buf)
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

func TestMessageBody(t *testing.T) {
	tests := []struct {
		name   string
		props  MessageBodyProps
		expect string
	}{
		{
			name:   "Default",
			props:  MessageBodyProps{},
			expect: `<div class="message-body"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  MessageBodyProps{ID: "body1", Class: []string{"custom-body"}},
			expect: `<div id="body1" class="message-body custom-body"></div>`,
		},
		{
			name: "With multiple classes and attributes",
			props: MessageBodyProps{
				ID:         "msg-body",
				Class:      []string{"content", "formatted"},
				Attributes: templ.Attributes{"data-content": "main"},
			},
			expect: `<div id="msg-body" class="message-body content formatted" data-content="main"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := MessageBody(tt.props).Render(context.Background(), &buf)
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
