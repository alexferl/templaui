package radio

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestRadio(t *testing.T) {
	tests := []struct {
		name   string
		props  RadioProps
		expect string
	}{
		{
			name:   "Default",
			props:  RadioProps{},
			expect: `<label class="radio"><input type="radio"></label>`,
		},
		{
			name:   "With ID and classes",
			props:  RadioProps{ID: "test-radio", Class: []string{"custom", "form-radio"}},
			expect: `<label id="test-radio" class="radio custom form-radio"><input type="radio"></label>`,
		},
		{
			name:   "With name and value",
			props:  RadioProps{Name: "size", Value: "large"},
			expect: `<label class="radio"><input type="radio" name="size" value="large"></label>`,
		},
		{
			name:   "Checked",
			props:  RadioProps{Checked: true},
			expect: `<label class="radio"><input type="radio" checked></label>`,
		},
		{
			name:   "Disabled",
			props:  RadioProps{Disabled: true},
			expect: `<label class="radio"><input type="radio" disabled></label>`,
		},
		{
			name:   "Required",
			props:  RadioProps{Required: true},
			expect: `<label class="radio"><input type="radio" required></label>`,
		},
		{
			name:   "Checked and disabled",
			props:  RadioProps{Checked: true, Disabled: true},
			expect: `<label class="radio"><input type="radio" checked disabled></label>`,
		},
		{
			name: "Complete configuration",
			props: RadioProps{
				ID:       "size-large",
				Class:    []string{"form-control", "size-option"},
				Name:     "size",
				Value:    "large",
				Checked:  true,
				Required: true,
			},
			expect: `<label id="size-large" class="radio form-control size-option"><input type="radio" name="size" value="large" checked required></label>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Radio(tt.props).Render(context.Background(), &buf)
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

func TestRadios(t *testing.T) {
	tests := []struct {
		name   string
		props  RadiosProps
		expect string
	}{
		{
			name:   "Default",
			props:  RadiosProps{},
			expect: `<div class="radios"></div>`,
		},
		{
			name:   "With ID and classes",
			props:  RadiosProps{ID: "test-radios", Class: []string{"custom", "form-group"}},
			expect: `<div id="test-radios" class="radios custom form-group"></div>`,
		},
		{
			name: "Complete configuration",
			props: RadiosProps{
				ID:    "size-radios",
				Class: []string{"form-group", "size-selection"},
			},
			expect: `<div id="size-radios" class="radios form-group size-selection"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Radios(tt.props).Render(context.Background(), &buf)
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

func TestRadioInput(t *testing.T) {
	tests := []struct {
		name   string
		props  RadioInputProps
		expect string
	}{
		{
			name:   "Default",
			props:  RadioInputProps{},
			expect: `<input type="radio" class="">`,
		},
		{
			name:   "With ID and classes",
			props:  RadioInputProps{ID: "test-input", Class: []string{"custom"}},
			expect: `<input id="test-input" type="radio" class="custom">`,
		},
		{
			name:   "With name and value",
			props:  RadioInputProps{Name: "color", Value: "red"},
			expect: `<input type="radio" name="color" value="red" class="">`,
		},
		{
			name:   "Checked",
			props:  RadioInputProps{Checked: true},
			expect: `<input type="radio" checked class="">`,
		},
		{
			name:   "Disabled",
			props:  RadioInputProps{Disabled: true},
			expect: `<input type="radio" disabled class="">`,
		},
		{
			name:   "Required",
			props:  RadioInputProps{Required: true},
			expect: `<input type="radio" required class="">`,
		},
		{
			name: "Complete configuration",
			props: RadioInputProps{
				ID:       "standalone-radio",
				Class:    []string{"custom-radio"},
				Name:     "subscription",
				Value:    "premium",
				Checked:  false,
				Required: true,
			},
			expect: `<input id="standalone-radio" type="radio" name="subscription" value="premium" required class="custom-radio">`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := RadioInput(tt.props).Render(context.Background(), &buf)
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

func TestRadioStates(t *testing.T) {
	tests := []struct {
		name     string
		props    RadioProps
		expected map[string]bool
	}{
		{
			name:  "Default state",
			props: RadioProps{},
			expected: map[string]bool{
				"checked":  false,
				"disabled": false,
				"required": false,
			},
		},
		{
			name:  "Checked state",
			props: RadioProps{Checked: true},
			expected: map[string]bool{
				"checked":  true,
				"disabled": false,
				"required": false,
			},
		},
		{
			name:  "Disabled state",
			props: RadioProps{Disabled: true},
			expected: map[string]bool{
				"checked":  false,
				"disabled": true,
				"required": false,
			},
		},
		{
			name:  "Required state",
			props: RadioProps{Required: true},
			expected: map[string]bool{
				"checked":  false,
				"disabled": false,
				"required": true,
			},
		},
		{
			name:  "All states",
			props: RadioProps{Checked: true, Disabled: true, Required: true},
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
			err := Radio(tt.props).Render(context.Background(), &buf)
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

func TestRadioGrouping(t *testing.T) {
	tests := []struct {
		name       string
		radios     []RadioProps
		expectSame bool // should have same name attribute
		groupName  string
	}{
		{
			name: "Linked radio buttons",
			radios: []RadioProps{
				{Name: "size", Value: "small"},
				{Name: "size", Value: "medium"},
				{Name: "size", Value: "large"},
			},
			expectSame: true,
			groupName:  "size",
		},
		{
			name: "Different radio groups",
			radios: []RadioProps{
				{Name: "size", Value: "small"},
				{Name: "color", Value: "red"},
			},
			expectSame: false,
			groupName:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var names []string

			for _, radioProps := range tt.radios {
				var buf strings.Builder
				err := Radio(radioProps).Render(context.Background(), &buf)
				if err != nil {
					t.Fatalf("render failed: %v", err)
				}
				got := strings.TrimSpace(buf.String())

				if radioProps.Name != "" {
					expectedName := `name="` + radioProps.Name + `"`
					if !strings.Contains(got, expectedName) {
						t.Errorf("expected name attribute %s to be present in: %s", expectedName, got)
					}
					names = append(names, radioProps.Name)
				}
			}

			if tt.expectSame && len(names) > 1 {
				firstName := names[0]
				for _, name := range names[1:] {
					if name != firstName {
						t.Errorf("expected all radio buttons to have same name, got different names: %v", names)
					}
				}
			}
		})
	}
}

func TestRadioAttributes(t *testing.T) {
	props := RadioProps{
		Attributes: templ.Attributes{
			"data-testid": "size-radio",
			"aria-label":  "Size selection",
		},
	}

	var buf strings.Builder
	err := Radio(props).Render(context.Background(), &buf)
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
