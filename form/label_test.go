package form

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestLabel(t *testing.T) {
	tests := []struct {
		name   string
		props  LabelProps
		expect string
	}{
		{
			name:   "Default",
			props:  LabelProps{},
			expect: `<label class="label"></label>`,
		},
		{
			name:   "With ID and classes",
			props:  LabelProps{ID: "test-label", Class: []string{"custom", "required"}},
			expect: `<label id="test-label" class="label custom required"></label>`,
		},
		{
			name:   "With for attribute",
			props:  LabelProps{For: "username"},
			expect: `<label for="username" class="label"></label>`,
		},
		{
			name:   "With for and ID",
			props:  LabelProps{ID: "username-label", For: "username"},
			expect: `<label id="username-label" for="username" class="label"></label>`,
		},
		{
			name: "Complete configuration",
			props: LabelProps{
				ID:    "main-label",
				Class: []string{"form-label", "required"},
				For:   "email-input",
			},
			expect: `<label id="main-label" for="email-input" class="label form-label required"></label>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Label(tt.props).Render(context.Background(), &buf)
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

func TestHelp(t *testing.T) {
	tests := []struct {
		name   string
		props  HelpProps
		expect string
	}{
		{
			name:   "Default",
			props:  HelpProps{},
			expect: `<p class="help"></p>`,
		},
		{
			name:   "With ID and classes",
			props:  HelpProps{ID: "test-help", Class: []string{"custom", "hint"}},
			expect: `<p id="test-help" class="help custom hint"></p>`,
		},
		{
			name:   "With primary color",
			props:  HelpProps{Color: IsPrimary},
			expect: `<p class="help is-primary"></p>`,
		},
		{
			name:   "With danger color",
			props:  HelpProps{Color: IsDanger},
			expect: `<p class="help is-danger"></p>`,
		},
		{
			name:   "With success color",
			props:  HelpProps{Color: IsSuccess},
			expect: `<p class="help is-success"></p>`,
		},
		{
			name:   "With warning color",
			props:  HelpProps{Color: IsWarning},
			expect: `<p class="help is-warning"></p>`,
		},
		{
			name:   "With info color",
			props:  HelpProps{Color: IsInfo},
			expect: `<p class="help is-info"></p>`,
		},
		{
			name: "Complete configuration",
			props: HelpProps{
				ID:    "error-help",
				Class: []string{"error-text", "validation"},
				Color: IsDanger,
			},
			expect: `<p id="error-help" class="help is-danger error-text validation"></p>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Help(tt.props).Render(context.Background(), &buf)
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

func TestHelpColors(t *testing.T) {
	tests := []struct {
		name   string
		color  Color
		expect string
	}{
		{"Primary", IsPrimary, `<p class="help is-primary"></p>`},
		{"Link", IsLink, `<p class="help is-link"></p>`},
		{"Info", IsInfo, `<p class="help is-info"></p>`},
		{"Success", IsSuccess, `<p class="help is-success"></p>`},
		{"Warning", IsWarning, `<p class="help is-warning"></p>`},
		{"Danger", IsDanger, `<p class="help is-danger"></p>`},
		{"Empty color", Color(""), `<p class="help"></p>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Help(HelpProps{Color: tt.color}).Render(context.Background(), &buf)
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

func TestLabelAttributes(t *testing.T) {
	props := LabelProps{
		Attributes: templ.Attributes{
			"data-testid": "label",
			"aria-label":  "Form label",
		},
	}

	var buf strings.Builder
	err := Label(props).Render(context.Background(), &buf)
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

func TestHelpAttributes(t *testing.T) {
	props := HelpProps{
		Attributes: templ.Attributes{
			"data-testid": "help",
			"role":        "alert",
		},
	}

	var buf strings.Builder
	err := Help(props).Render(context.Background(), &buf)
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
