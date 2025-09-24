package input

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestInput(t *testing.T) {
	tests := []struct {
		name   string
		props  InputProps
		expect string
	}{
		{
			name:   "Default",
			props:  InputProps{},
			expect: `<input type="text" class="input">`,
		},
		{
			name:   "With ID and classes",
			props:  InputProps{ID: "test-input", Class: []string{"custom", "form-input"}},
			expect: `<input id="test-input" type="text" class="input custom form-input">`,
		},
		{
			name:   "Email type",
			props:  InputProps{Type: TypeEmail},
			expect: `<input type="email" class="input">`,
		},
		{
			name:   "Password type",
			props:  InputProps{Type: TypePassword},
			expect: `<input type="password" class="input">`,
		},
		{
			name:   "With name and value",
			props:  InputProps{Name: "username", Value: "john_doe"},
			expect: `<input type="text" name="username" value="john_doe" class="input">`,
		},
		{
			name:   "With placeholder",
			props:  InputProps{Placeholder: "Enter your name"},
			expect: `<input type="text" placeholder="Enter your name" class="input">`,
		},
		{
			name:   "Small size",
			props:  InputProps{Size: IsSmall},
			expect: `<input type="text" class="input is-small">`,
		},
		{
			name:   "Large size",
			props:  InputProps{Size: IsLarge},
			expect: `<input type="text" class="input is-large">`,
		},
		{
			name:   "Primary color",
			props:  InputProps{Color: IsPrimary},
			expect: `<input type="text" class="input is-primary">`,
		},
		{
			name:   "Danger color",
			props:  InputProps{Color: IsDanger},
			expect: `<input type="text" class="input is-danger">`,
		},
		{
			name:   "Disabled",
			props:  InputProps{Disabled: true},
			expect: `<input type="text" disabled class="input">`,
		},
		{
			name:   "Readonly",
			props:  InputProps{Readonly: true},
			expect: `<input type="text" readonly class="input">`,
		},
		{
			name:   "Required",
			props:  InputProps{Required: true},
			expect: `<input type="text" required class="input">`,
		},
		{
			name:   "Hovered state",
			props:  InputProps{IsHovered: true},
			expect: `<input type="text" class="input is-hovered">`,
		},
		{
			name:   "Focused state",
			props:  InputProps{IsFocused: true},
			expect: `<input type="text" class="input is-focused">`,
		},
		{
			name:   "Static appearance",
			props:  InputProps{IsStatic: true, Readonly: true},
			expect: `<input type="text" readonly class="input is-static">`,
		},
		{
			name:   "Rounded",
			props:  InputProps{IsRounded: true},
			expect: `<input type="text" class="input is-rounded">`,
		},
		{
			name:   "Number input with min/max/step",
			props:  InputProps{Type: TypeNumber, Min: "0", Max: "100", Step: "1"},
			expect: `<input type="number" min="0" max="100" step="1" class="input">`,
		},
		{
			name: "Complete configuration",
			props: InputProps{
				ID:           "user-email",
				Class:        []string{"form-control", "validated"},
				Type:         TypeEmail,
				Name:         "email",
				Placeholder:  "Enter email",
				Required:     true,
				Size:         IsMedium,
				Color:        IsSuccess,
				IsRounded:    true,
				Autocomplete: "email",
			},
			expect: `<input id="user-email" type="email" name="email" placeholder="Enter email" autocomplete="email" required class="input is-medium is-success is-rounded form-control validated">`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Input(tt.props).Render(context.Background(), &buf)
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

func TestInputTypes(t *testing.T) {
	tests := []struct {
		name      string
		inputType InputType
		expect    string
	}{
		{"Text", TypeText, `<input type="text" class="input">`},
		{"Password", TypePassword, `<input type="password" class="input">`},
		{"Email", TypeEmail, `<input type="email" class="input">`},
		{"Tel", TypeTel, `<input type="tel" class="input">`},
		{"Number", TypeNumber, `<input type="number" class="input">`},
		{"Search", TypeSearch, `<input type="search" class="input">`},
		{"URL", TypeURL, `<input type="url" class="input">`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Input(InputProps{Type: tt.inputType}).Render(context.Background(), &buf)
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

func TestInputSizes(t *testing.T) {
	tests := []struct {
		name   string
		size   Size
		expect string
	}{
		{"Small", IsSmall, `<input type="text" class="input is-small">`},
		{"Normal", IsNormal, `<input type="text" class="input is-normal">`},
		{"Medium", IsMedium, `<input type="text" class="input is-medium">`},
		{"Large", IsLarge, `<input type="text" class="input is-large">`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Input(InputProps{Size: tt.size}).Render(context.Background(), &buf)
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

func TestInputColors(t *testing.T) {
	tests := []struct {
		name   string
		color  Color
		expect string
	}{
		{"Primary", IsPrimary, `<input type="text" class="input is-primary">`},
		{"Link", IsLink, `<input type="text" class="input is-link">`},
		{"Info", IsInfo, `<input type="text" class="input is-info">`},
		{"Success", IsSuccess, `<input type="text" class="input is-success">`},
		{"Warning", IsWarning, `<input type="text" class="input is-warning">`},
		{"Danger", IsDanger, `<input type="text" class="input is-danger">`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Input(InputProps{Color: tt.color}).Render(context.Background(), &buf)
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

func TestInputStates(t *testing.T) {
	tests := []struct {
		name   string
		props  InputProps
		expect string
	}{
		{"Hovered", InputProps{IsHovered: true}, `<input type="text" class="input is-hovered">`},
		{"Focused", InputProps{IsFocused: true}, `<input type="text" class="input is-focused">`},
		{"Active", InputProps{IsActive: true}, `<input type="text" class="input is-active">`},
		{"Static", InputProps{IsStatic: true}, `<input type="text" class="input is-static">`},
		{"Rounded", InputProps{IsRounded: true}, `<input type="text" class="input is-rounded">`},
		{"Multiple states", InputProps{IsHovered: true, IsRounded: true}, `<input type="text" class="input is-hovered is-rounded">`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Input(tt.props).Render(context.Background(), &buf)
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

func TestInputAttributes(t *testing.T) {
	props := InputProps{
		Attributes: templ.Attributes{
			"data-testid": "username-input",
			"aria-label":  "Username field",
		},
	}

	var buf strings.Builder
	err := Input(props).Render(context.Background(), &buf)
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
