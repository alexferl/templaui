package div

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestDiv(t *testing.T) {
	tests := []struct {
		name   string
		props  []DivProps
		expect string
	}{
		{
			name:   "Default div",
			props:  []DivProps{},
			expect: `<div></div>`,
		},
		{
			name:   "With ID",
			props:  []DivProps{{ID: "main-content"}},
			expect: `<div id="main-content"></div>`,
		},
		{
			name:   "With single class",
			props:  []DivProps{{Class: []string{"container"}}},
			expect: `<div class="container"></div>`,
		},
		{
			name:   "With multiple classes",
			props:  []DivProps{{Class: []string{"box", "content", "is-large"}}},
			expect: `<div class="box content is-large"></div>`,
		},
		{
			name: "With ID and classes",
			props: []DivProps{{
				ID:    "hero-section",
				Class: []string{"hero", "is-primary"},
			}},
			expect: `<div id="hero-section" class="hero is-primary"></div>`,
		},
		{
			name: "With all properties",
			props: []DivProps{{
				ID:    "wrapper",
				Class: []string{"section", "has-background-light"},
			}},
			expect: `<div id="wrapper" class="section has-background-light"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Div(tt.props...).Render(context.Background(), &buf)
			if err != nil {
				t.Fatalf("render failed: %v", err)
			}
			got := strings.TrimSpace(buf.String())
			if got != tt.expect {
				t.Errorf("expected: %s, got: %s", tt.expect, got)
			}
		})
	}
}

func TestDivWithAttributes(t *testing.T) {
	props := DivProps{
		ID: "custom-div",
		Attributes: templ.Attributes{
			"data-testid": "main-wrapper",
			"role":        "banner",
			"aria-label":  "Main content area",
		},
	}

	var buf strings.Builder
	err := Div(props).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}

	got := buf.String()
	requiredParts := []string{
		`<div`,
		`id="custom-div"`,
		`data-testid="main-wrapper"`,
		`role="banner"`,
		`aria-label="Main content area"`,
		`></div>`,
	}

	for _, part := range requiredParts {
		if !strings.Contains(got, part) {
			t.Errorf("missing required part %q in output: %s", part, got)
		}
	}
}

func TestDivBulmaClasses(t *testing.T) {
	tests := []struct {
		name    string
		classes []string
		expect  []string
	}{
		{
			name:    "Container classes",
			classes: []string{"container", "is-fluid"},
			expect:  []string{"container", "is-fluid"},
		},
		{
			name:    "Layout classes",
			classes: []string{"columns", "is-multiline", "is-centered"},
			expect:  []string{"columns", "is-multiline", "is-centered"},
		},
		{
			name:    "Hero classes",
			classes: []string{"hero", "is-primary", "is-large"},
			expect:  []string{"hero", "is-primary", "is-large"},
		},
		{
			name:    "Box and content classes",
			classes: []string{"box", "content", "notification", "is-info"},
			expect:  []string{"box", "content", "notification", "is-info"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Div(DivProps{Class: tt.classes}).Render(context.Background(), &buf)
			if err != nil {
				t.Fatalf("render failed: %v", err)
			}
			got := buf.String()

			for _, expectedClass := range tt.expect {
				if !strings.Contains(got, expectedClass) {
					t.Errorf("expected to contain class %q, got: %s", expectedClass, got)
				}
			}
		})
	}
}

func TestDivEmptyValues(t *testing.T) {
	props := DivProps{
		// All fields are empty
		ID:    "",
		Class: []string{},
	}

	var buf strings.Builder
	err := Div(props).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}

	got := strings.TrimSpace(buf.String())
	expected := `<div></div>`

	if got != expected {
		t.Errorf("expected: %s, got: %s", expected, got)
	}
}

func TestDivNoProps(t *testing.T) {
	var buf strings.Builder
	err := Div().Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}

	got := strings.TrimSpace(buf.String())
	expected := `<div></div>`

	if got != expected {
		t.Errorf("expected: %s, got: %s", expected, got)
	}
}

func TestDivClassComposition(t *testing.T) {
	props := DivProps{
		Class: []string{"first-class", "second-class", "third-class"},
	}

	var buf strings.Builder
	err := Div(props).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}

	got := buf.String()
	expected := `class="first-class second-class third-class"`

	if !strings.Contains(got, expected) {
		t.Errorf("expected to contain %q, got: %s", expected, got)
	}
}

func TestDivAccessibilityAttributes(t *testing.T) {
	props := DivProps{
		ID:    "accessible-div",
		Class: []string{"section"},
		Attributes: templ.Attributes{
			"role":          "main",
			"aria-label":    "Main content section",
			"aria-expanded": "true",
			"tabindex":      "0",
		},
	}

	var buf strings.Builder
	err := Div(props).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}

	got := buf.String()
	accessibilityFeatures := []string{
		`role="main"`,
		`aria-label="Main content section"`,
		`aria-expanded="true"`,
		`tabindex="0"`,
	}

	for _, feature := range accessibilityFeatures {
		if !strings.Contains(got, feature) {
			t.Errorf("missing accessibility feature %q in output: %s", feature, got)
		}
	}
}

func TestDivDataAttributes(t *testing.T) {
	props := DivProps{
		Class: []string{"component"},
		Attributes: templ.Attributes{
			"data-component": "hero-banner",
			"data-version":   "1.0",
			"data-testid":    "hero-test",
			"data-analytics": "track-banner",
		},
	}

	var buf strings.Builder
	err := Div(props).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}

	got := buf.String()
	dataAttributes := []string{
		`data-component="hero-banner"`,
		`data-version="1.0"`,
		`data-testid="hero-test"`,
		`data-analytics="track-banner"`,
	}

	for _, attr := range dataAttributes {
		if !strings.Contains(got, attr) {
			t.Errorf("missing data attribute %q in output: %s", attr, got)
		}
	}
}
