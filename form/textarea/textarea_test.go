package textarea

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestTextarea(t *testing.T) {
	tests := []struct {
		name   string
		props  TextareaProps
		expect string
	}{
		{
			name:   "Default",
			props:  TextareaProps{},
			expect: `<textarea class="textarea"></textarea>`,
		},
		{
			name:   "With ID and classes",
			props:  TextareaProps{ID: "test-textarea", Class: []string{"custom", "form-textarea"}},
			expect: `<textarea id="test-textarea" class="textarea custom form-textarea"></textarea>`,
		},
		{
			name:   "With name and value",
			props:  TextareaProps{Name: "message", Value: "Hello world"},
			expect: `<textarea name="message" class="textarea">Hello world</textarea>`,
		},
		{
			name:   "With placeholder",
			props:  TextareaProps{Placeholder: "Enter your message"},
			expect: `<textarea placeholder="Enter your message" class="textarea"></textarea>`,
		},
		{
			name:   "With rows",
			props:  TextareaProps{Rows: 5},
			expect: `<textarea rows="5" class="textarea"></textarea>`,
		},
		{
			name:   "With cols",
			props:  TextareaProps{Cols: 50},
			expect: `<textarea cols="50" class="textarea"></textarea>`,
		},
		{
			name:   "With rows and cols",
			props:  TextareaProps{Rows: 4, Cols: 60},
			expect: `<textarea rows="4" cols="60" class="textarea"></textarea>`,
		},
		{
			name:   "Small size",
			props:  TextareaProps{Size: IsSmall},
			expect: `<textarea class="textarea is-small"></textarea>`,
		},
		{
			name:   "Large size",
			props:  TextareaProps{Size: IsLarge},
			expect: `<textarea class="textarea is-large"></textarea>`,
		},
		{
			name:   "Primary color",
			props:  TextareaProps{Color: IsPrimary},
			expect: `<textarea class="textarea is-primary"></textarea>`,
		},
		{
			name:   "Danger color",
			props:  TextareaProps{Color: IsDanger},
			expect: `<textarea class="textarea is-danger"></textarea>`,
		},
		{
			name:   "Disabled",
			props:  TextareaProps{Disabled: true},
			expect: `<textarea disabled class="textarea"></textarea>`,
		},
		{
			name:   "Readonly",
			props:  TextareaProps{Readonly: true},
			expect: `<textarea readonly class="textarea"></textarea>`,
		},
		{
			name:   "Required",
			props:  TextareaProps{Required: true},
			expect: `<textarea required class="textarea"></textarea>`,
		},
		{
			name:   "Hovered state",
			props:  TextareaProps{IsHovered: true},
			expect: `<textarea class="textarea is-hovered"></textarea>`,
		},
		{
			name:   "Focused state",
			props:  TextareaProps{IsFocused: true},
			expect: `<textarea class="textarea is-focused"></textarea>`,
		},
		{
			name:   "Fixed size",
			props:  TextareaProps{HasFixedSize: true},
			expect: `<textarea class="textarea has-fixed-size"></textarea>`,
		},
		{
			name: "Complete configuration",
			props: TextareaProps{
				ID:           "user-message",
				Class:        []string{"form-control", "validated"},
				Name:         "message",
				Value:        "Sample text",
				Placeholder:  "Enter message",
				Rows:         6,
				Required:     true,
				Size:         IsMedium,
				Color:        IsSuccess,
				HasFixedSize: true,
			},
			expect: `<textarea id="user-message" name="message" placeholder="Enter message" rows="6" required class="textarea is-medium is-success has-fixed-size form-control validated">Sample text</textarea>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Textarea(tt.props).Render(context.Background(), &buf)
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

func TestTextareaSizes(t *testing.T) {
	tests := []struct {
		name   string
		size   Size
		expect string
	}{
		{"Small", IsSmall, `<textarea class="textarea is-small"></textarea>`},
		{"Normal", IsNormal, `<textarea class="textarea is-normal"></textarea>`},
		{"Medium", IsMedium, `<textarea class="textarea is-medium"></textarea>`},
		{"Large", IsLarge, `<textarea class="textarea is-large"></textarea>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Textarea(TextareaProps{Size: tt.size}).Render(context.Background(), &buf)
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

func TestTextareaColors(t *testing.T) {
	tests := []struct {
		name   string
		color  Color
		expect string
	}{
		{"Primary", IsPrimary, `<textarea class="textarea is-primary"></textarea>`},
		{"Link", IsLink, `<textarea class="textarea is-link"></textarea>`},
		{"Info", IsInfo, `<textarea class="textarea is-info"></textarea>`},
		{"Success", IsSuccess, `<textarea class="textarea is-success"></textarea>`},
		{"Warning", IsWarning, `<textarea class="textarea is-warning"></textarea>`},
		{"Danger", IsDanger, `<textarea class="textarea is-danger"></textarea>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Textarea(TextareaProps{Color: tt.color}).Render(context.Background(), &buf)
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

func TestTextareaStates(t *testing.T) {
	tests := []struct {
		name   string
		props  TextareaProps
		expect string
	}{
		{"Hovered", TextareaProps{IsHovered: true}, `<textarea class="textarea is-hovered"></textarea>`},
		{"Focused", TextareaProps{IsFocused: true}, `<textarea class="textarea is-focused"></textarea>`},
		{"Active", TextareaProps{IsActive: true}, `<textarea class="textarea is-active"></textarea>`},
		{"Fixed size", TextareaProps{HasFixedSize: true}, `<textarea class="textarea has-fixed-size"></textarea>`},
		{"Multiple states", TextareaProps{IsHovered: true, HasFixedSize: true}, `<textarea class="textarea is-hovered has-fixed-size"></textarea>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Textarea(tt.props).Render(context.Background(), &buf)
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

func TestTextareaAttributes(t *testing.T) {
	props := TextareaProps{
		Attributes: templ.Attributes{
			"data-testid": "message-textarea",
			"aria-label":  "Message field",
		},
	}

	var buf strings.Builder
	err := Textarea(props).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}
	got := strings.TrimSpace(buf.String())

	for key, value := range props.Attributes {
		expectedAttr := key + `="` + fmt.Sprintf("%v", value) + `"`
		if !strings.Contains(got, expectedAttr) {
			t.Errorf("expected attribute %s to be present in: %s", expectedAttr, got)
		}
	}
}
