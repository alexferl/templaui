package form

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestFieldset(t *testing.T) {
	tests := []struct {
		name   string
		props  FieldsetProps
		expect string
	}{
		{
			name:   "Default",
			props:  FieldsetProps{},
			expect: `<fieldset class=""></fieldset>`,
		},
		{
			name:   "With ID and classes",
			props:  FieldsetProps{ID: "test-fieldset", Class: []string{"custom", "form-group"}},
			expect: `<fieldset id="test-fieldset" class="custom form-group"></fieldset>`,
		},
		{
			name:   "Disabled",
			props:  FieldsetProps{Disabled: true},
			expect: `<fieldset disabled class=""></fieldset>`,
		},
		{
			name:   "Disabled with classes",
			props:  FieldsetProps{Disabled: true, Class: []string{"form-section"}},
			expect: `<fieldset disabled class="form-section"></fieldset>`,
		},
		{
			name: "Complete configuration",
			props: FieldsetProps{
				ID:       "main-fieldset",
				Class:    []string{"form-group", "disabled-section"},
				Disabled: true,
			},
			expect: `<fieldset id="main-fieldset" disabled class="form-group disabled-section"></fieldset>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Fieldset(tt.props).Render(context.Background(), &buf)
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

func TestLegend(t *testing.T) {
	tests := []struct {
		name   string
		props  LegendProps
		expect string
	}{
		{
			name:   "Default",
			props:  LegendProps{},
			expect: `<legend class=""></legend>`,
		},
		{
			name:   "With ID and classes",
			props:  LegendProps{ID: "test-legend", Class: []string{"form-legend", "title"}},
			expect: `<legend id="test-legend" class="form-legend title"></legend>`,
		},
		{
			name: "Complete configuration",
			props: LegendProps{
				ID:    "main-legend",
				Class: []string{"section-title", "required"},
			},
			expect: `<legend id="main-legend" class="section-title required"></legend>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Legend(tt.props).Render(context.Background(), &buf)
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

func TestFieldsetAttributes(t *testing.T) {
	props := FieldsetProps{
		Attributes: templ.Attributes{
			"data-testid": "fieldset",
			"role":        "group",
		},
	}

	var buf strings.Builder
	err := Fieldset(props).Render(context.Background(), &buf)
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

func TestLegendAttributes(t *testing.T) {
	props := LegendProps{
		Attributes: templ.Attributes{
			"data-testid": "legend",
			"aria-label":  "Form section",
		},
	}

	var buf strings.Builder
	err := Legend(props).Render(context.Background(), &buf)
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
