package paragraph

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestParagraph(t *testing.T) {
	tests := []struct {
		name   string
		props  []ParagraphProps
		expect string
	}{
		{
			name:   "Default paragraph",
			props:  []ParagraphProps{},
			expect: `<p></p>`,
		},
		{
			name: "With ID",
			props: []ParagraphProps{{
				ID: "intro-text",
			}},
			expect: `<p id="intro-text"></p>`,
		},
		{
			name: "With size",
			props: []ParagraphProps{{
				Size: IsSize5,
			}},
			expect: `<p class="is-size-5"></p>`,
		},
		{
			name: "With text color",
			props: []ParagraphProps{{
				TextColor: HasTextPrimary,
			}},
			expect: `<p class="has-text-primary"></p>`,
		},
		{
			name: "With custom classes",
			props: []ParagraphProps{{
				Class: []string{"mb-4", "has-text-centered"},
			}},
			expect: `<p class="mb-4 has-text-centered"></p>`,
		},
		{
			name: "With size and color",
			props: []ParagraphProps{{
				Size:      IsSize3,
				TextColor: HasTextInfo,
			}},
			expect: `<p class="is-size-3 has-text-info"></p>`,
		},
		{
			name: "With all properties",
			props: []ParagraphProps{{
				ID:        "main-para",
				Size:      IsSize4,
				TextColor: HasTextSuccess,
				Class:     []string{"mb-3", "is-italic"},
			}},
			expect: `<p id="main-para" class="is-size-4 has-text-success mb-3 is-italic"></p>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Paragraph(tt.props...).Render(context.Background(), &buf)
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

func TestParagraphWithAttributes(t *testing.T) {
	props := ParagraphProps{
		ID: "test-paragraph",
		Attributes: templ.Attributes{
			"data-testid": "intro",
			"role":        "note",
		},
	}

	var buf strings.Builder
	err := Paragraph(props).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}

	got := buf.String()
	requiredParts := []string{
		`<p`,
		`id="test-paragraph"`,
		`data-testid="intro"`,
		`role="note"`,
		`></p>`,
	}

	for _, part := range requiredParts {
		if !strings.Contains(got, part) {
			t.Errorf("missing required part %q in output: %s", part, got)
		}
	}
}

func TestParagraphSizes(t *testing.T) {
	tests := []struct {
		name string
		size Size
	}{
		{"Size 1", IsSize1},
		{"Size 2", IsSize2},
		{"Size 3", IsSize3},
		{"Size 4", IsSize4},
		{"Size 5", IsSize5},
		{"Size 6", IsSize6},
		{"Size 7", IsSize7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Paragraph(ParagraphProps{Size: tt.size}).Render(context.Background(), &buf)
			if err != nil {
				t.Fatalf("render failed: %v", err)
			}
			got := buf.String()
			expected := string(tt.size)
			if !strings.Contains(got, expected) {
				t.Errorf("expected to contain %q, got: %s", expected, got)
			}
		})
	}
}

func TestParagraphTextColors(t *testing.T) {
	tests := []struct {
		name  string
		color TextColor
	}{
		{"Primary", HasTextPrimary},
		{"Info", HasTextInfo},
		{"Success", HasTextSuccess},
		{"Warning", HasTextWarning},
		{"Danger", HasTextDanger},
		{"Dark", HasTextDark},
		{"Light", HasTextLight},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Paragraph(ParagraphProps{TextColor: tt.color}).Render(context.Background(), &buf)
			if err != nil {
				t.Fatalf("render failed: %v", err)
			}
			got := buf.String()
			expected := string(tt.color)
			if !strings.Contains(got, expected) {
				t.Errorf("expected to contain %q, got: %s", expected, got)
			}
		})
	}
}

func TestParagraphClassPrecedence(t *testing.T) {
	props := ParagraphProps{
		Size:      IsSize5,
		TextColor: HasTextPrimary,
		Class:     []string{"custom-class", "another-class"},
	}

	var buf strings.Builder
	err := Paragraph(props).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}

	got := buf.String()

	// Check that all classes are present
	expectedClasses := []string{
		"is-size-5",
		"has-text-primary",
		"custom-class",
		"another-class",
	}

	for _, class := range expectedClasses {
		if !strings.Contains(got, class) {
			t.Errorf("expected to contain class %q, got: %s", class, got)
		}
	}
}

func TestParagraphEmptyValues(t *testing.T) {
	props := ParagraphProps{
		Size:      "",         // Empty size should not add class
		TextColor: "",         // Empty color should not add class
		Class:     []string{}, // Empty class slice should not add class
	}

	var buf strings.Builder
	err := Paragraph(props).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}

	got := strings.TrimSpace(buf.String())
	expected := `<p></p>`

	if got != expected {
		t.Errorf("expected: %s, got: %s", expected, got)
	}
}
