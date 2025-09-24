package templaui

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestDocument(t *testing.T) {
	tests := []struct {
		name         string
		props        DocumentProps
		headContents []templ.Component
		contains     []string
	}{
		{
			name:  "Default props",
			props: DocumentProps{},
			contains: []string{
				`<!doctype html>`,
				`<html lang="en">`,
				`<title>My App</title>`,
				`<meta charset="UTF-8">`,
				`<meta name="viewport" content="width=device-width, initial-scale=1.0">`,
				`<meta name="color-scheme" content="light dark">`,
				`<meta name="theme-color" content="#00d1b2">`,
				`<link rel="icon" href="data:,">`,
				`<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bulma/1.0.4/css/bulma.min.css"`,
			},
		},
		{
			name: "With all fields",
			props: DocumentProps{
				Title:        "Test App",
				Description:  "Test description",
				Lang:         "es",
				Favicon:      "/favicon.ico",
				UseLocalCSS:  true,
				LocalCSSPath: "/css/custom.css",
				Charset:      "UTF-8",
				Viewport:     "width=device-width, initial-scale=1.0",
				Theme:        "#ff3860",
				ColorScheme:  "dark",
				BodyClass:    "has-navbar-fixed-top",
			},
			contains: []string{
				`<html lang="es">`,
				`<title>Test App</title>`,
				`<meta name="description" content="Test description">`,
				`<meta name="color-scheme" content="dark">`,
				`<meta name="theme-color" content="#ff3860">`,
				`<link rel="icon" href="/favicon.ico">`,
				`<link rel="stylesheet" href="/css/custom.css">`,
				`<body class="has-navbar-fixed-top">`,
			},
		},
		{
			name: "With local CSS default path",
			props: DocumentProps{
				Title:       "Test",
				UseLocalCSS: true,
			},
			contains: []string{
				`<title>Test</title>`,
				`<meta name="color-scheme" content="light dark">`,
				`<link rel="icon" href="data:,">`,
				`<link rel="stylesheet" href="/static/css/bulma.min.css">`,
			},
		},
		{
			name: "With NoFavicon true",
			props: DocumentProps{
				Title:     "No Favicon App",
				NoFavicon: true,
			},
			contains: []string{
				`<title>No Favicon App</title>`,
				`<meta name="color-scheme" content="light dark">`,
				`<meta name="theme-color" content="#00d1b2">`,
			},
		},
		{
			name: "With custom favicon and NoFavicon false",
			props: DocumentProps{
				Title:     "Custom Favicon App",
				Favicon:   "/my-favicon.ico",
				NoFavicon: false,
			},
			contains: []string{
				`<title>Custom Favicon App</title>`,
				`<link rel="icon" href="/my-favicon.ico">`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Document(tt.props, tt.headContents...).Render(context.Background(), &buf)
			if err != nil {
				t.Fatalf("render failed: %v", err)
			}
			got := buf.String()

			for _, expected := range tt.contains {
				if !strings.Contains(got, expected) {
					t.Errorf("expected to contain %q, got: %s", expected, got)
				}
			}
		})
	}
}

func TestDocumentWithAttributes(t *testing.T) {
	props := DocumentProps{
		Title:     "Test App",
		BodyClass: "dark-theme",
		BodyAttrs: templ.Attributes{"data-theme": "dark", "data-version": "1.0"},
	}

	var buf strings.Builder
	err := Document(props).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}

	got := buf.String()
	requiredParts := []string{
		`<!doctype html>`,
		`<html lang="en">`,
		`<title>Test App</title>`,
		`<meta name="color-scheme" content="light dark">`,
		`<body class="dark-theme"`,
		`data-theme="dark"`,
		`data-version="1.0"`,
		`></body></html>`,
	}

	for _, part := range requiredParts {
		if !strings.Contains(got, part) {
			t.Errorf("missing required part %q in output: %s", part, got)
		}
	}
}

func TestHead(t *testing.T) {
	tests := []struct {
		name         string
		props        HeadProps
		headContents []templ.Component
		expect       string
	}{
		{
			name:   "Basic head with title",
			props:  HeadProps{Title: "Test Title"},
			expect: `<head><title>Test Title</title></head>`,
		},
		{
			name:   "Empty title",
			props:  HeadProps{},
			expect: `<head><title></title></head>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Head(tt.props, tt.headContents...).Render(context.Background(), &buf)
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

func TestMetaHead(t *testing.T) {
	tests := []struct {
		name     string
		props    MetaHeadProps
		contains []string
	}{
		{
			name: "With all fields",
			props: MetaHeadProps{
				Description: "Test description",
				Favicon:     "/favicon.ico",
				Charset:     "UTF-8",
				Viewport:    "width=device-width, initial-scale=1.0",
				Theme:       "#00d1b2",
				ColorScheme: "dark",
				Keywords:    "test,app,bulma",
				Author:      "Test Author",
			},
			contains: []string{
				`<meta charset="UTF-8">`,
				`<meta name="viewport" content="width=device-width, initial-scale=1.0">`,
				`<meta name="color-scheme" content="dark">`,
				`<meta name="description" content="Test description">`,
				`<meta name="keywords" content="test,app,bulma">`,
				`<meta name="author" content="Test Author">`,
				`<meta name="theme-color" content="#00d1b2">`,
				`<link rel="icon" href="/favicon.ico">`,
			},
		},
		{
			name: "With minimal fields (defaults)",
			props: MetaHeadProps{
				Charset:  "UTF-8",
				Viewport: "width=device-width, initial-scale=1.0",
			},
			contains: []string{
				`<meta charset="UTF-8">`,
				`<meta name="viewport" content="width=device-width, initial-scale=1.0">`,
				`<meta name="color-scheme" content="light dark">`,
				`<meta name="theme-color" content="#00d1b2">`,
				`<link rel="icon" href="data:,">`,
			},
		},
		{
			name: "With custom theme and color scheme",
			props: MetaHeadProps{
				Charset:     "UTF-8",
				Viewport:    "width=device-width, initial-scale=1.0",
				Theme:       "#ff3860",
				ColorScheme: "light",
			},
			contains: []string{
				`<meta name="color-scheme" content="light">`,
				`<meta name="theme-color" content="#ff3860">`,
			},
		},
		{
			name: "Only color scheme specified",
			props: MetaHeadProps{
				Charset:     "UTF-8",
				Viewport:    "width=device-width, initial-scale=1.0",
				ColorScheme: "dark light",
			},
			contains: []string{
				`<meta name="color-scheme" content="dark light">`,
				`<meta name="theme-color" content="#00d1b2">`,
			},
		},
		{
			name: "With NoFavicon true",
			props: MetaHeadProps{
				Charset:     "UTF-8",
				Viewport:    "width=device-width, initial-scale=1.0",
				Description: "Test description",
				NoFavicon:   true,
			},
			contains: []string{
				`<meta charset="UTF-8">`,
				`<meta name="viewport" content="width=device-width, initial-scale=1.0">`,
				`<meta name="description" content="Test description">`,
				`<meta name="color-scheme" content="light dark">`,
				`<meta name="theme-color" content="#00d1b2">`,
			},
		},
		{
			name: "With custom favicon and NoFavicon false",
			props: MetaHeadProps{
				Charset:   "UTF-8",
				Viewport:  "width=device-width, initial-scale=1.0",
				Favicon:   "/test-favicon.png",
				NoFavicon: false,
			},
			contains: []string{
				`<meta charset="UTF-8">`,
				`<meta name="viewport" content="width=device-width, initial-scale=1.0">`,
				`<link rel="icon" href="/test-favicon.png">`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := MetaHead(tt.props).Render(context.Background(), &buf)
			if err != nil {
				t.Fatalf("render failed: %v", err)
			}
			got := buf.String()

			for _, expected := range tt.contains {
				if !strings.Contains(got, expected) {
					t.Errorf("expected to contain %q, got: %s", expected, got)
				}
			}
		})
	}
}

func TestMetaHeadDefaults(t *testing.T) {
	props := MetaHeadProps{
		Charset:  "UTF-8",
		Viewport: "width=device-width, initial-scale=1.0",
	}

	var buf strings.Builder
	err := MetaHead(props).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}

	got := buf.String()
	requiredParts := []string{
		`<meta charset="UTF-8">`,
		`<meta name="viewport" content="width=device-width, initial-scale=1.0">`,
		`<meta name="color-scheme" content="light dark">`,
		`<meta name="theme-color" content="#00d1b2">`,
		`<link rel="icon" href="data:,">`,
	}

	for _, part := range requiredParts {
		if !strings.Contains(got, part) {
			t.Errorf("missing required part %q in output: %s", part, got)
		}
	}
}

func TestMetaHeadColorSchemeDefaults(t *testing.T) {
	tests := []struct {
		name        string
		colorScheme string
		expect      string
	}{
		{
			name:        "Default light dark",
			colorScheme: "",
			expect:      `<meta name="color-scheme" content="light dark">`,
		},
		{
			name:        "Custom light only",
			colorScheme: "light",
			expect:      `<meta name="color-scheme" content="light">`,
		},
		{
			name:        "Custom dark only",
			colorScheme: "dark",
			expect:      `<meta name="color-scheme" content="dark">`,
		},
		{
			name:        "Custom dark light order",
			colorScheme: "dark light",
			expect:      `<meta name="color-scheme" content="dark light">`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			props := MetaHeadProps{
				Charset:     "UTF-8",
				Viewport:    "width=device-width, initial-scale=1.0",
				ColorScheme: tt.colorScheme,
			}

			var buf strings.Builder
			err := MetaHead(props).Render(context.Background(), &buf)
			if err != nil {
				t.Fatalf("render failed: %v", err)
			}
			got := buf.String()

			if !strings.Contains(got, tt.expect) {
				t.Errorf("expected to contain %q, got: %s", tt.expect, got)
			}
		})
	}
}

func TestMeta(t *testing.T) {
	tests := []struct {
		name   string
		props  MetaProps
		expect string
	}{
		{
			name:   "Charset meta",
			props:  MetaProps{Charset: "UTF-8"},
			expect: `<meta charset="UTF-8">`,
		},
		{
			name:   "Name meta",
			props:  MetaProps{Name: "description", Content: "Test description"},
			expect: `<meta name="description" content="Test description">`,
		},
		{
			name:   "Property meta (OpenGraph)",
			props:  MetaProps{Property: "og:title", Content: "Page Title"},
			expect: `<meta property="og:title" content="Page Title">`,
		},
		{
			name:   "HttpEquiv meta",
			props:  MetaProps{HttpEquiv: "refresh", Content: "30"},
			expect: `<meta http-equiv="refresh" content="30">`,
		},
		{
			name:   "Empty props",
			props:  MetaProps{},
			expect: ``,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Meta(tt.props).Render(context.Background(), &buf)
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

func TestMetaColorScheme(t *testing.T) {
	tests := []struct {
		name   string
		props  MetaProps
		expect string
	}{
		{
			name:   "Color scheme light dark",
			props:  MetaProps{Name: "color-scheme", Content: "light dark"},
			expect: `<meta name="color-scheme" content="light dark">`,
		},
		{
			name:   "Color scheme dark only",
			props:  MetaProps{Name: "color-scheme", Content: "dark"},
			expect: `<meta name="color-scheme" content="dark">`,
		},
		{
			name:   "Color scheme light only",
			props:  MetaProps{Name: "color-scheme", Content: "light"},
			expect: `<meta name="color-scheme" content="light">`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Meta(tt.props).Render(context.Background(), &buf)
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

func TestLink(t *testing.T) {
	tests := []struct {
		name   string
		props  LinkProps
		expect string
	}{
		{
			name:   "Stylesheet link",
			props:  LinkProps{Rel: "stylesheet", Href: "/css/style.css"},
			expect: `<link rel="stylesheet" href="/css/style.css">`,
		},
		{
			name:   "Icon link with type and sizes",
			props:  LinkProps{Rel: "icon", Type: "image/png", Sizes: "32x32", Href: "/favicon.png"},
			expect: `<link rel="icon" href="/favicon.png" type="image/png" sizes="32x32">`,
		},
		{
			name: "CDN stylesheet with integrity",
			props: LinkProps{
				Rel:            "stylesheet",
				Href:           "https://cdn.example.com/style.css",
				Integrity:      "sha256-abcd1234",
				CrossOrigin:    "anonymous",
				ReferrerPolicy: "no-referrer",
			},
			expect: `<link rel="stylesheet" href="https://cdn.example.com/style.css" integrity="sha256-abcd1234" crossorigin="anonymous" referrerpolicy="no-referrer">`,
		},
		{
			name:   "Preload link",
			props:  LinkProps{Rel: "preload", As: "font", Href: "/fonts/font.woff2"},
			expect: `<link rel="preload" href="/fonts/font.woff2" as="font">`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Link(tt.props).Render(context.Background(), &buf)
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

func TestLinkDataURI(t *testing.T) {
	props := LinkProps{
		Rel:  "icon",
		Href: "data:,",
	}

	var buf strings.Builder
	err := Link(props).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}

	got := strings.TrimSpace(buf.String())
	expected := `<link rel="icon" href="data:,">`

	if got != expected {
		t.Errorf("expected: %s, got: %s", expected, got)
	}
}

func TestLinkWithAttributes(t *testing.T) {
	props := LinkProps{
		Rel:  "stylesheet",
		Href: "/css/style.css",
		Attributes: templ.Attributes{
			"data-version": "1.0",
			"data-theme":   "dark",
		},
	}

	var buf strings.Builder
	err := Link(props).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}

	got := buf.String()
	requiredParts := []string{
		`<link rel="stylesheet" href="/css/style.css"`,
		`data-version="1.0"`,
		`data-theme="dark"`,
		`>`,
	}

	for _, part := range requiredParts {
		if !strings.Contains(got, part) {
			t.Errorf("missing required part %q in output: %s", part, got)
		}
	}
}

func TestBody(t *testing.T) {
	tests := []struct {
		name   string
		props  []BodyProps
		expect string
	}{
		{
			name:   "Default body",
			props:  []BodyProps{},
			expect: `<body></body>`,
		},
		{
			name:   "With ID and class",
			props:  []BodyProps{{ID: "main-body", Class: "custom-body"}},
			expect: `<body id="main-body" class="custom-body"></body>`,
		},
		{
			name:   "With attributes",
			props:  []BodyProps{{Attributes: templ.Attributes{"data-role": "main"}}},
			expect: `<body data-role="main"></body>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Body(tt.props...).Render(context.Background(), &buf)
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

func TestBodyWithAttributes(t *testing.T) {
	props := BodyProps{
		ID:         "main-body",
		Class:      "custom-body dark-theme",
		Attributes: templ.Attributes{"data-role": "main", "data-theme": "dark"},
	}

	var buf strings.Builder
	err := Body(props).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}

	got := buf.String()
	requiredParts := []string{
		`<body`,
		`id="main-body"`,
		`class="custom-body dark-theme"`,
		`data-role="main"`,
		`data-theme="dark"`,
		`></body>`,
	}

	for _, part := range requiredParts {
		if !strings.Contains(got, part) {
			t.Errorf("missing required part %q in output: %s", part, got)
		}
	}
}

func TestPreloadCSS(t *testing.T) {
	var buf strings.Builder
	err := PreloadCSS().Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}

	got := buf.String()
	expectedParts := []string{
		`<link rel="preconnect" href="https://cdnjs.cloudflare.com" crossorigin="anonymous">`,
		`<link rel="dns-prefetch" href="https://cdnjs.cloudflare.com">`,
	}

	for _, part := range expectedParts {
		if !strings.Contains(got, part) {
			t.Errorf("missing required preload link: %s", part)
		}
	}
}

func TestBulmaCSS(t *testing.T) {
	tests := []struct {
		name     string
		config   BulmaConfig
		contains []string
	}{
		{
			name:   "Default CDN configuration",
			config: BulmaConfig{},
			contains: []string{
				`<link rel="stylesheet"`,
				`href="https://cdnjs.cloudflare.com/ajax/libs/bulma/1.0.4/css/bulma.min.css"`,
				`integrity="sha512-yh2RE0wZCVZeysGiqTwDTO/dKelCbS9bP2L94UvOFtl/FKXcNAje3Y2oBg/ZMZ3LS1sicYk4dYVGtDex75fvvA=="`,
				`crossorigin="anonymous"`,
				`referrerpolicy="no-referrer"`,
			},
		},
		{
			name: "Local CSS with custom path",
			config: BulmaConfig{
				UseLocal:  true,
				LocalPath: "/css/custom-bulma.css",
			},
			contains: []string{
				`<link rel="stylesheet" href="/css/custom-bulma.css">`,
			},
		},
		{
			name: "Local CSS with default path",
			config: BulmaConfig{
				UseLocal: true,
			},
			contains: []string{
				`<link rel="stylesheet" href="/static/css/bulma.min.css">`,
			},
		},
		{
			name: "CDN with preload",
			config: BulmaConfig{
				Preload: true,
			},
			contains: []string{
				`<link rel="preconnect" href="https://cdnjs.cloudflare.com" crossorigin="anonymous">`,
				`<link rel="dns-prefetch" href="https://cdnjs.cloudflare.com">`,
				`<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bulma/1.0.4/css/bulma.min.css"`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := BulmaCSS(tt.config).Render(context.Background(), &buf)
			if err != nil {
				t.Fatalf("render failed: %v", err)
			}
			got := buf.String()

			for _, expected := range tt.contains {
				if !strings.Contains(got, expected) {
					t.Errorf("expected to contain %q, got: %s", expected, got)
				}
			}
		})
	}
}

func TestHTML(t *testing.T) {
	tests := []struct {
		name         string
		props        HTMLProps
		headContents []templ.Component
		expect       string
	}{
		{
			name:   "Default props",
			props:  HTMLProps{},
			expect: `<!doctype html><html lang="en"><head><title>My App</title></head><body></body></html>`,
		},
		{
			name: "With custom lang and title",
			props: HTMLProps{
				Title: "Custom App",
				Lang:  "de",
			},
			expect: `<!doctype html><html lang="de"><head><title>Custom App</title></head><body></body></html>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := HTML(tt.props, tt.headContents...).Render(context.Background(), &buf)
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

func TestHTMLWithAttributes(t *testing.T) {
	props := HTMLProps{
		Title:      "Test App",
		Lang:       "en",
		Attributes: templ.Attributes{"data-theme": "dark", "class": "no-js"},
	}

	var buf strings.Builder
	err := HTML(props).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}

	got := buf.String()
	requiredParts := []string{
		`<!doctype html>`,
		`<html lang="en"`,
		`data-theme="dark"`,
		`class="no-js"`,
		`<head><title>Test App</title></head>`,
		`<body></body></html>`,
	}

	for _, part := range requiredParts {
		if !strings.Contains(got, part) {
			t.Errorf("missing required part %q in output: %s", part, got)
		}
	}
}
