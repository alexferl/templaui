package loader

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestLoader(t *testing.T) {
	tests := []struct {
		name   string
		props  LoaderProps
		expect string
	}{
		{
			name:   "Default",
			props:  LoaderProps{},
			expect: `<div class="loader" role="status" aria-label="Loading" aria-live="polite"></div>`,
		},
		{
			name:   "With ID",
			props:  LoaderProps{ID: "spinner1"},
			expect: `<div id="spinner1" class="loader" role="status" aria-label="Loading" aria-live="polite"></div>`,
		},
		{
			name:   "With custom classes",
			props:  LoaderProps{Class: []string{"custom-loader"}},
			expect: `<div class="loader custom-loader" role="status" aria-label="Loading" aria-live="polite"></div>`,
		},
		{
			name:   "With custom aria-label",
			props:  LoaderProps{AriaLabel: "Saving data"},
			expect: `<div class="loader" role="status" aria-label="Saving data" aria-live="polite"></div>`,
		},
		{
			name:   "With custom aria-label for search",
			props:  LoaderProps{AriaLabel: "Loading search results"},
			expect: `<div class="loader" role="status" aria-label="Loading search results" aria-live="polite"></div>`,
		},
		{
			name:   "With custom aria-label for upload",
			props:  LoaderProps{AriaLabel: "Uploading file"},
			expect: `<div class="loader" role="status" aria-label="Uploading file" aria-live="polite"></div>`,
		},
		{
			name:   "With ID and custom aria-label",
			props:  LoaderProps{ID: "payment-loader", AriaLabel: "Processing payment"},
			expect: `<div id="payment-loader" class="loader" role="status" aria-label="Processing payment" aria-live="polite"></div>`,
		},
		{
			name:   "With ID, classes, and custom aria-label",
			props:  LoaderProps{ID: "loader1", Class: []string{"small", "primary"}, AriaLabel: "Loading content"},
			expect: `<div id="loader1" class="loader small primary" role="status" aria-label="Loading content" aria-live="polite"></div>`,
		},
		{
			name: "With additional attributes",
			props: LoaderProps{
				ID:         "loading-spinner",
				AriaLabel:  "Loading data",
				Attributes: templ.Attributes{"data-testid": "loader"},
			},
			expect: `<div id="loading-spinner" class="loader" role="status" aria-label="Loading data" aria-live="polite" data-testid="loader"></div>`,
		},
		{
			name: "With empty aria-label (falls back to default)",
			props: LoaderProps{
				ID:        "default-loader",
				AriaLabel: "",
			},
			expect: `<div id="default-loader" class="loader" role="status" aria-label="Loading" aria-live="polite"></div>`,
		},
		{
			name: "All fields combined",
			props: LoaderProps{
				ID:         "main-loader",
				Class:      []string{"large", "centered"},
				AriaLabel:  "Processing request",
				Attributes: templ.Attributes{"data-testid": "main-loader", "data-context": "form"},
			},
			expect: `<div id="main-loader" class="loader large centered" role="status" aria-label="Processing request" aria-live="polite" data-context="form" data-testid="main-loader"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Loader(tt.props).Render(context.Background(), &buf)
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

func TestLoaderAriaLabels(t *testing.T) {
	tests := []struct {
		name      string
		ariaLabel string
		expect    string
	}{
		{
			name:      "Default (empty)",
			ariaLabel: "",
			expect:    `<div class="loader" role="status" aria-label="Loading" aria-live="polite"></div>`,
		},
		{
			name:      "Custom saving",
			ariaLabel: "Saving data",
			expect:    `<div class="loader" role="status" aria-label="Saving data" aria-live="polite"></div>`,
		},
		{
			name:      "Custom search",
			ariaLabel: "Loading search results",
			expect:    `<div class="loader" role="status" aria-label="Loading search results" aria-live="polite"></div>`,
		},
		{
			name:      "Custom upload",
			ariaLabel: "Uploading file",
			expect:    `<div class="loader" role="status" aria-label="Uploading file" aria-live="polite"></div>`,
		},
		{
			name:      "Custom payment",
			ariaLabel: "Processing payment",
			expect:    `<div class="loader" role="status" aria-label="Processing payment" aria-live="polite"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Loader(LoaderProps{AriaLabel: tt.ariaLabel}).Render(context.Background(), &buf)
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
