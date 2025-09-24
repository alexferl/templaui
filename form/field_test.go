package form

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestField(t *testing.T) {
	tests := []struct {
		name   string
		props  FieldProps
		expect string
	}{
		{
			name:   "Default",
			props:  FieldProps{},
			expect: `<div class="field"></div>`,
		},
		{
			name:   "With ID and classes",
			props:  FieldProps{ID: "test-field", Class: []string{"custom"}},
			expect: `<div id="test-field" class="field custom"></div>`,
		},
		{
			name:   "Is grouped",
			props:  FieldProps{IsGrouped: true},
			expect: `<div class="field is-grouped"></div>`,
		},
		{
			name:   "Is grouped centered",
			props:  FieldProps{IsGroupedCentered: true},
			expect: `<div class="field is-grouped is-grouped-centered"></div>`,
		},
		{
			name:   "Has addons",
			props:  FieldProps{HasAddons: true},
			expect: `<div class="field has-addons"></div>`,
		},
		{
			name:   "Has addons centered",
			props:  FieldProps{HasAddonsCentered: true},
			expect: `<div class="field has-addons has-addons-centered"></div>`,
		},
		{
			name:   "Is horizontal",
			props:  FieldProps{IsHorizontal: true},
			expect: `<div class="field is-horizontal"></div>`,
		},
		{
			name: "All modifiers",
			props: FieldProps{
				IsGrouped:          true,
				IsGroupedCentered:  true,
				IsGroupedMultiline: true,
				HasAddons:          true,
				HasAddonsCentered:  true,
				IsHorizontal:       true,
			},
			expect: `<div class="field is-grouped is-grouped-centered is-grouped-multiline has-addons has-addons-centered is-horizontal"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Field(tt.props).Render(context.Background(), &buf)
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

func TestFieldModifiers(t *testing.T) {
	tests := []struct {
		name          string
		props         FieldProps
		expectGrouped bool
		expectAddons  bool
	}{
		{"No modifiers", FieldProps{}, false, false},
		{"Grouped only", FieldProps{IsGrouped: true}, true, false},
		{"Grouped centered implies grouped", FieldProps{IsGroupedCentered: true}, true, false},
		{"Addons only", FieldProps{HasAddons: true}, false, true},
		{"Addons centered implies addons", FieldProps{HasAddonsCentered: true}, false, true},
		{"Both", FieldProps{IsGrouped: true, HasAddons: true}, true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Field(tt.props).Render(context.Background(), &buf)
			if err != nil {
				t.Fatalf("render failed: %v", err)
			}
			got := strings.TrimSpace(buf.String())

			hasGrouped := strings.Contains(got, "is-grouped")
			hasAddons := strings.Contains(got, "has-addons")

			if hasGrouped != tt.expectGrouped {
				t.Errorf("expected is-grouped: %v, got: %v", tt.expectGrouped, hasGrouped)
			}
			if hasAddons != tt.expectAddons {
				t.Errorf("expected has-addons: %v, got: %v", tt.expectAddons, hasAddons)
			}
		})
	}
}

func TestFieldLabel(t *testing.T) {
	tests := []struct {
		name   string
		props  FieldLabelProps
		expect string
	}{
		{
			name:   "Default",
			props:  FieldLabelProps{},
			expect: `<div class="field-label"></div>`,
		},
		{
			name:   "With size",
			props:  FieldLabelProps{Size: IsLarge},
			expect: `<div class="field-label is-large"></div>`,
		},
		{
			name:   "With ID and classes",
			props:  FieldLabelProps{ID: "label", Class: []string{"custom"}},
			expect: `<div id="label" class="field-label custom"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := FieldLabel(tt.props).Render(context.Background(), &buf)
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

func TestFieldBody(t *testing.T) {
	tests := []struct {
		name   string
		props  FieldBodyProps
		expect string
	}{
		{
			name:   "Default",
			props:  FieldBodyProps{},
			expect: `<div class="field-body"></div>`,
		},
		{
			name:   "With ID and classes",
			props:  FieldBodyProps{ID: "body", Class: []string{"custom"}},
			expect: `<div id="body" class="field-body custom"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := FieldBody(tt.props).Render(context.Background(), &buf)
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

func TestFieldAttributes(t *testing.T) {
	props := FieldProps{
		Attributes: templ.Attributes{
			"data-testid": "field",
			"role":        "group",
		},
	}

	var buf strings.Builder
	err := Field(props).Render(context.Background(), &buf)
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
