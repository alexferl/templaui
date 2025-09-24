package checkbox

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestCheckbox(t *testing.T) {
	tests := []struct {
		name   string
		props  CheckboxProps
		expect string
	}{
		{
			name:   "Default",
			props:  CheckboxProps{},
			expect: `<label class="checkbox"><input type="checkbox"></label>`,
		},
		{
			name:   "With ID and classes",
			props:  CheckboxProps{ID: "test-checkbox", Class: []string{"custom", "form-checkbox"}},
			expect: `<label id="test-checkbox" class="checkbox custom form-checkbox"><input type="checkbox"></label>`,
		},
		{
			name:   "With name and value",
			props:  CheckboxProps{Name: "agree", Value: "yes"},
			expect: `<label class="checkbox"><input type="checkbox" name="agree" value="yes"></label>`,
		},
		{
			name:   "Checked",
			props:  CheckboxProps{Checked: true},
			expect: `<label class="checkbox"><input type="checkbox" checked></label>`,
		},
		{
			name:   "Disabled",
			props:  CheckboxProps{Disabled: true},
			expect: `<label class="checkbox"><input type="checkbox" disabled></label>`,
		},
		{
			name:   "Required",
			props:  CheckboxProps{Required: true},
			expect: `<label class="checkbox"><input type="checkbox" required></label>`,
		},
		{
			name:   "Checked and disabled",
			props:  CheckboxProps{Checked: true, Disabled: true},
			expect: `<label class="checkbox"><input type="checkbox" checked disabled></label>`,
		},
		{
			name: "Complete configuration",
			props: CheckboxProps{
				ID:       "terms-checkbox",
				Class:    []string{"form-control", "required"},
				Name:     "terms",
				Value:    "accepted",
				Checked:  true,
				Required: true,
			},
			expect: `<label id="terms-checkbox" class="checkbox form-control required"><input type="checkbox" name="terms" value="accepted" checked required></label>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Checkbox(tt.props).Render(context.Background(), &buf)
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

func TestCheckboxes(t *testing.T) {
	tests := []struct {
		name   string
		props  CheckboxesProps
		expect string
	}{
		{
			name:   "Default",
			props:  CheckboxesProps{},
			expect: `<div class="checkboxes"></div>`,
		},
		{
			name:   "With ID and classes",
			props:  CheckboxesProps{ID: "test-checkboxes", Class: []string{"custom", "form-group"}},
			expect: `<div id="test-checkboxes" class="checkboxes custom form-group"></div>`,
		},
		{
			name: "Complete configuration",
			props: CheckboxesProps{
				ID:    "interests-checkboxes",
				Class: []string{"form-group", "multi-select"},
			},
			expect: `<div id="interests-checkboxes" class="checkboxes form-group multi-select"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Checkboxes(tt.props).Render(context.Background(), &buf)
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

func TestCheckboxInput(t *testing.T) {
	tests := []struct {
		name   string
		props  CheckboxInputProps
		expect string
	}{
		{
			name:   "Default",
			props:  CheckboxInputProps{},
			expect: `<input type="checkbox" class="">`,
		},
		{
			name:   "With ID and classes",
			props:  CheckboxInputProps{ID: "test-input", Class: []string{"custom"}},
			expect: `<input id="test-input" type="checkbox" class="custom">`,
		},
		{
			name:   "With name and value",
			props:  CheckboxInputProps{Name: "option", Value: "1"},
			expect: `<input type="checkbox" name="option" value="1" class="">`,
		},
		{
			name:   "Checked",
			props:  CheckboxInputProps{Checked: true},
			expect: `<input type="checkbox" checked class="">`,
		},
		{
			name:   "Disabled",
			props:  CheckboxInputProps{Disabled: true},
			expect: `<input type="checkbox" disabled class="">`,
		},
		{
			name:   "Required",
			props:  CheckboxInputProps{Required: true},
			expect: `<input type="checkbox" required class="">`,
		},
		{
			name: "Complete configuration",
			props: CheckboxInputProps{
				ID:       "standalone-checkbox",
				Class:    []string{"custom-checkbox"},
				Name:     "newsletter",
				Value:    "subscribe",
				Checked:  false,
				Required: true,
			},
			expect: `<input id="standalone-checkbox" type="checkbox" name="newsletter" value="subscribe" required class="custom-checkbox">`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := CheckboxInput(tt.props).Render(context.Background(), &buf)
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

func TestCheckboxStates(t *testing.T) {
	tests := []struct {
		name     string
		props    CheckboxProps
		expected map[string]bool
	}{
		{
			name:  "Default state",
			props: CheckboxProps{},
			expected: map[string]bool{
				"checked":  false,
				"disabled": false,
				"required": false,
			},
		},
		{
			name:  "Checked state",
			props: CheckboxProps{Checked: true},
			expected: map[string]bool{
				"checked":  true,
				"disabled": false,
				"required": false,
			},
		},
		{
			name:  "Disabled state",
			props: CheckboxProps{Disabled: true},
			expected: map[string]bool{
				"checked":  false,
				"disabled": true,
				"required": false,
			},
		},
		{
			name:  "Required state",
			props: CheckboxProps{Required: true},
			expected: map[string]bool{
				"checked":  false,
				"disabled": false,
				"required": true,
			},
		},
		{
			name:  "All states",
			props: CheckboxProps{Checked: true, Disabled: true, Required: true},
			expected: map[string]bool{
				"checked":  true,
				"disabled": true,
				"required": true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Checkbox(tt.props).Render(context.Background(), &buf)
			if err != nil {
				t.Fatalf("render failed: %v", err)
			}
			got := strings.TrimSpace(buf.String())

			for state, shouldHave := range tt.expected {
				hasState := strings.Contains(got, state)
				if hasState != shouldHave {
					t.Errorf("expected %s attribute: %v, got: %v in: %s", state, shouldHave, hasState, got)
				}
			}
		})
	}
}

func TestCheckboxAttributes(t *testing.T) {
	props := CheckboxProps{
		Attributes: templ.Attributes{
			"data-testid": "agree-checkbox",
			"aria-label":  "Agreement checkbox",
		},
	}

	var buf strings.Builder
	err := Checkbox(props).Render(context.Background(), &buf)
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
