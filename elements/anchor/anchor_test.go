package anchor

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestAnchor(t *testing.T) {
	tests := []struct {
		name   string
		props  AnchorProps
		expect string
	}{
		{
			name:   "Basic anchor with href",
			props:  AnchorProps{Href: "/about"},
			expect: `<a href="/about"></a>`,
		},
		{
			name:   "With ID",
			props:  AnchorProps{Href: "/contact", ID: "contact-link"},
			expect: `<a href="/contact" id="contact-link"></a>`,
		},
		{
			name:   "With target blank",
			props:  AnchorProps{Href: "https://example.com", Target: TargetBlank},
			expect: `<a href="https://example.com" target="_blank"></a>`,
		},
		{
			name:   "With single rel",
			props:  AnchorProps{Href: "/help", Rel: []Rel{RelHelp}},
			expect: `<a href="/help" rel="help"></a>`,
		},
		{
			name:   "With multiple rel attributes",
			props:  AnchorProps{Href: "https://external.com", Rel: []Rel{RelNoopener, RelNoreferrer}},
			expect: `<a href="https://external.com" rel="noopener noreferrer"></a>`,
		},
		{
			name:   "With text color",
			props:  AnchorProps{Href: "/products", TextColor: HasTextPrimary},
			expect: `<a href="/products" class="has-text-primary"></a>`,
		},
		{
			name:   "With title",
			props:  AnchorProps{Href: "/info", Title: "More information"},
			expect: `<a href="/info" title="More information"></a>`,
		},
		{
			name:   "With aria-label",
			props:  AnchorProps{Href: "/search", AriaLabel: "Search the website"},
			expect: `<a href="/search" aria-label="Search the website"></a>`,
		},
		{
			name:   "With download",
			props:  AnchorProps{Href: "/file.pdf", Download: "document.pdf"},
			expect: `<a href="/file.pdf" download="document.pdf"></a>`,
		},
		{
			name:   "With custom classes",
			props:  AnchorProps{Href: "/styled", Class: []string{"is-underlined", "has-text-weight-bold"}},
			expect: `<a href="/styled" class="is-underlined has-text-weight-bold"></a>`,
		},
		{
			name: "With all properties",
			props: AnchorProps{
				Href:      "https://example.com/page",
				ID:        "external-link",
				Target:    TargetBlank,
				Rel:       []Rel{RelNoopener, RelExternal},
				TextColor: HasTextInfo,
				Title:     "Visit external site",
				AriaLabel: "Open external website in new tab",
				Class:     []string{"custom-link"},
			},
			expect: `<a href="https://example.com/page" id="external-link" target="_blank" rel="noopener external" title="Visit external site" aria-label="Open external website in new tab" class="has-text-info custom-link"></a>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Anchor(tt.props).Render(context.Background(), &buf)
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

func TestAnchorWithAttributes(t *testing.T) {
	props := AnchorProps{
		Href: "/test",
		Attributes: templ.Attributes{
			"data-testid": "test-link",
			"role":        "button",
		},
	}

	var buf strings.Builder
	err := Anchor(props).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}

	got := buf.String()
	requiredParts := []string{
		`<a`,
		`href="/test"`,
		`data-testid="test-link"`,
		`role="button"`,
		`></a>`,
	}

	for _, part := range requiredParts {
		if !strings.Contains(got, part) {
			t.Errorf("missing required part %q in output: %s", part, got)
		}
	}
}

func TestAnchorAriaLabel(t *testing.T) {
	tests := []struct {
		name      string
		ariaLabel string
	}{
		{"Simple label", "Click here"},
		{"Descriptive label", "Navigate to the about page"},
		{"Icon label", "Search"},
		{"Action label", "Download PDF document"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Anchor(AnchorProps{Href: "/test", AriaLabel: tt.ariaLabel}).Render(context.Background(), &buf)
			if err != nil {
				t.Fatalf("render failed: %v", err)
			}
			got := buf.String()
			expected := `aria-label="` + tt.ariaLabel + `"`
			if !strings.Contains(got, expected) {
				t.Errorf("expected to contain %q, got: %s", expected, got)
			}
		})
	}
}

func TestAnchorTargets(t *testing.T) {
	tests := []struct {
		name   string
		target Target
		expect string
	}{
		{"Target blank", TargetBlank, "_blank"},
		{"Target self", TargetSelf, "_self"},
		{"Target parent", TargetParent, "_parent"},
		{"Target top", TargetTop, "_top"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Anchor(AnchorProps{Href: "/test", Target: tt.target}).Render(context.Background(), &buf)
			if err != nil {
				t.Fatalf("render failed: %v", err)
			}
			got := buf.String()
			if !strings.Contains(got, `target="`+tt.expect+`"`) {
				t.Errorf("expected target=%q, got: %s", tt.expect, got)
			}
		})
	}
}

func TestAnchorRelValues(t *testing.T) {
	tests := []struct {
		name string
		rel  Rel
	}{
		{"Noopener", RelNoopener},
		{"Noreferrer", RelNoreferrer},
		{"Nofollow", RelNofollow},
		{"External", RelExternal},
		{"Bookmark", RelBookmark},
		{"Help", RelHelp},
		{"License", RelLicense},
		{"Next", RelNext},
		{"Prev", RelPrev},
		{"Tag", RelTag},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Anchor(AnchorProps{Href: "/test", Rel: []Rel{tt.rel}}).Render(context.Background(), &buf)
			if err != nil {
				t.Fatalf("render failed: %v", err)
			}
			got := buf.String()
			expected := string(tt.rel)
			if !strings.Contains(got, expected) {
				t.Errorf("expected to contain rel=%q, got: %s", expected, got)
			}
		})
	}
}

func TestAnchorTextColors(t *testing.T) {
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
		{"White", HasTextWhite},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Anchor(AnchorProps{Href: "/test", TextColor: tt.color}).Render(context.Background(), &buf)
			if err != nil {
				t.Fatalf("render failed: %v", err)
			}
			got := buf.String()
			expected := string(tt.color)
			if !strings.Contains(got, expected) {
				t.Errorf("expected to contain color class %q, got: %s", expected, got)
			}
		})
	}
}

func TestExternalLink(t *testing.T) {
	var buf strings.Builder
	err := ExternalLink("https://github.com", "Visit GitHub").Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}

	got := buf.String()
	requiredParts := []string{
		`href="https://github.com"`,
		`target="_blank"`,
		`rel="noopener noreferrer"`,
		`Visit GitHub`,
	}

	for _, part := range requiredParts {
		if !strings.Contains(got, part) {
			t.Errorf("missing required part %q in output: %s", part, got)
		}
	}
}

func TestExternalLinkWithLabel(t *testing.T) {
	var buf strings.Builder
	err := ExternalLinkWithLabel("https://twitter.com/user", "Twitter", "Follow us on Twitter").Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}

	got := buf.String()
	requiredParts := []string{
		`href="https://twitter.com/user"`,
		`target="_blank"`,
		`rel="noopener noreferrer"`,
		`aria-label="Follow us on Twitter"`,
		`Twitter`,
	}

	for _, part := range requiredParts {
		if !strings.Contains(got, part) {
			t.Errorf("missing required part %q in output: %s", part, got)
		}
	}
}

func TestEmailLink(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		subject string
		body    string
		text    string
		expect  string
	}{
		{
			name:   "Basic email",
			email:  "test@example.com",
			text:   "Contact Us",
			expect: `<a href="mailto:test@example.com">Contact Us</a>`,
		},
		{
			name:    "Email with subject",
			email:   "support@example.com",
			subject: "Help Request",
			text:    "Get Help",
			expect:  `<a href="mailto:support@example.com?subject=Help Request">Get Help</a>`,
		},
		{
			name:   "Email with body",
			email:  "info@example.com",
			body:   "Please send more information",
			text:   "More Info",
			expect: `<a href="mailto:info@example.com?body=Please send more information">More Info</a>`,
		},
		{
			name:    "Email with subject and body",
			email:   "contact@example.com",
			subject: "Question",
			body:    "I have a question about...",
			text:    "Ask Question",
			expect:  `<a href="mailto:contact@example.com?subject=Question&amp;body=I have a question about...">Ask Question</a>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := EmailLink(tt.email, tt.subject, tt.body, tt.text).Render(context.Background(), &buf)
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

func TestConvertRelSlice(t *testing.T) {
	rels := []Rel{RelNoopener, RelNoreferrer, RelNofollow}
	expected := []string{"noopener", "noreferrer", "nofollow"}

	result := convertRelSlice(rels)

	if len(result) != len(expected) {
		t.Fatalf("expected length %d, got %d", len(expected), len(result))
	}

	for i, exp := range expected {
		if result[i] != exp {
			t.Errorf("at index %d: expected %q, got %q", i, exp, result[i])
		}
	}
}

func TestAnchorEmptyValues(t *testing.T) {
	props := AnchorProps{
		Href: "/test", // Only href is required
		// All other fields are empty
	}

	var buf strings.Builder
	err := Anchor(props).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}

	got := strings.TrimSpace(buf.String())
	expected := `<a href="/test"></a>`

	if got != expected {
		t.Errorf("expected: %s, got: %s", expected, got)
	}
}

func TestAnchorAccessibilityFeatures(t *testing.T) {
	props := AnchorProps{
		Href:      "https://example.com",
		Target:    TargetBlank,
		Rel:       []Rel{RelNoopener, RelNoreferrer},
		AriaLabel: "Visit example website (opens in new tab)",
		Title:     "Example Website",
	}

	var buf strings.Builder
	err := Anchor(props).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}

	got := buf.String()
	accessibilityFeatures := []string{
		`aria-label="Visit example website (opens in new tab)"`,
		`title="Example Website"`,
		`rel="noopener noreferrer"`,
		`target="_blank"`,
	}

	for _, feature := range accessibilityFeatures {
		if !strings.Contains(got, feature) {
			t.Errorf("missing accessibility feature %q in output: %s", feature, got)
		}
	}
}
