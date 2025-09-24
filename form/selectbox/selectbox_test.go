package selectbox

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestSelect(t *testing.T) {
	tests := []struct {
		name   string
		props  SelectProps
		expect string
	}{
		{
			name:   "Default",
			props:  SelectProps{},
			expect: `<div class="select"></div>`,
		},
		{
			name:   "With ID and classes",
			props:  SelectProps{ID: "test-select", Class: []string{"custom", "form-select"}},
			expect: `<div id="test-select" class="select custom form-select"></div>`,
		},
		{
			name:   "Is multiple",
			props:  SelectProps{IsMultiple: true},
			expect: `<div class="select is-multiple"></div>`,
		},
		{
			name:   "Is rounded",
			props:  SelectProps{IsRounded: true},
			expect: `<div class="select is-rounded"></div>`,
		},
		{
			name:   "Is fullwidth",
			props:  SelectProps{IsFullwidth: true},
			expect: `<div class="select is-fullwidth"></div>`,
		},
		{
			name:   "Small size",
			props:  SelectProps{Size: IsSmall},
			expect: `<div class="select is-small"></div>`,
		},
		{
			name:   "Large size",
			props:  SelectProps{Size: IsLarge},
			expect: `<div class="select is-large"></div>`,
		},
		{
			name:   "Primary color",
			props:  SelectProps{Color: IsPrimary},
			expect: `<div class="select is-primary"></div>`,
		},
		{
			name:   "Danger color",
			props:  SelectProps{Color: IsDanger},
			expect: `<div class="select is-danger"></div>`,
		},
		{
			name: "Complete configuration",
			props: SelectProps{
				ID:          "country-select",
				Class:       []string{"form-control", "validated"},
				IsMultiple:  true,
				IsRounded:   true,
				IsFullwidth: true,
				Size:        IsMedium,
				Color:       IsSuccess,
			},
			expect: `<div id="country-select" class="select is-multiple is-rounded is-fullwidth is-medium is-success form-control validated"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Select(tt.props).Render(context.Background(), &buf)
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

func TestSelectElement(t *testing.T) {
	tests := []struct {
		name   string
		props  SelectElementProps
		expect string
	}{
		{
			name:   "Default",
			props:  SelectElementProps{},
			expect: `<select class=""></select>`,
		},
		{
			name:   "With ID and name",
			props:  SelectElementProps{ID: "country", Name: "country"},
			expect: `<select id="country" name="country" class=""></select>`,
		},
		{
			name:   "Multiple",
			props:  SelectElementProps{Multiple: true},
			expect: `<select multiple class=""></select>`,
		},
		{
			name:   "With size",
			props:  SelectElementProps{Size: 5},
			expect: `<select size="5" class=""></select>`,
		},
		{
			name:   "Disabled",
			props:  SelectElementProps{Disabled: true},
			expect: `<select disabled class=""></select>`,
		},
		{
			name:   "Required",
			props:  SelectElementProps{Required: true},
			expect: `<select required class=""></select>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := SelectElement(tt.props).Render(context.Background(), &buf)
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

func TestOption(t *testing.T) {
	tests := []struct {
		name   string
		props  OptionProps
		expect string
	}{
		{
			name:   "Default",
			props:  OptionProps{},
			expect: `<option class=""></option>`,
		},
		{
			name:   "With value",
			props:  OptionProps{Value: "us"},
			expect: `<option value="us" class=""></option>`,
		},
		{
			name:   "Selected",
			props:  OptionProps{Value: "us", Selected: true},
			expect: `<option value="us" selected class=""></option>`,
		},
		{
			name:   "Disabled",
			props:  OptionProps{Value: "us", Disabled: true},
			expect: `<option value="us" disabled class=""></option>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Option(tt.props).Render(context.Background(), &buf)
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

func TestOptgroup(t *testing.T) {
	tests := []struct {
		name   string
		props  OptgroupProps
		expect string
	}{
		{
			name:   "Default",
			props:  OptgroupProps{},
			expect: `<optgroup class=""></optgroup>`,
		},
		{
			name:   "With label",
			props:  OptgroupProps{Label: "North America"},
			expect: `<optgroup label="North America" class=""></optgroup>`,
		},
		{
			name:   "Disabled",
			props:  OptgroupProps{Label: "Europe", Disabled: true},
			expect: `<optgroup label="Europe" disabled class=""></optgroup>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Optgroup(tt.props).Render(context.Background(), &buf)
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

func TestSelectSizes(t *testing.T) {
	tests := []struct {
		name   string
		size   Size
		expect string
	}{
		{"Small", IsSmall, `<div class="select is-small"></div>`},
		{"Normal", IsNormal, `<div class="select is-normal"></div>`},
		{"Medium", IsMedium, `<div class="select is-medium"></div>`},
		{"Large", IsLarge, `<div class="select is-large"></div>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Select(SelectProps{Size: tt.size}).Render(context.Background(), &buf)
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

func TestSelectColors(t *testing.T) {
	tests := []struct {
		name   string
		color  Color
		expect string
	}{
		{"Primary", IsPrimary, `<div class="select is-primary"></div>`},
		{"Link", IsLink, `<div class="select is-link"></div>`},
		{"Info", IsInfo, `<div class="select is-info"></div>`},
		{"Success", IsSuccess, `<div class="select is-success"></div>`},
		{"Warning", IsWarning, `<div class="select is-warning"></div>`},
		{"Danger", IsDanger, `<div class="select is-danger"></div>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Select(SelectProps{Color: tt.color}).Render(context.Background(), &buf)
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

func TestSelectAttributes(t *testing.T) {
	props := SelectProps{
		Attributes: templ.Attributes{
			"data-testid": "country-select",
			"aria-label":  "Select country",
		},
	}

	var buf strings.Builder
	err := Select(props).Render(context.Background(), &buf)
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
