package form

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestForm(t *testing.T) {
	tests := []struct {
		name   string
		props  FormProps
		expect string
	}{
		{
			name:   "Default",
			props:  FormProps{},
			expect: `<form class=""></form>`,
		},
		{
			name:   "With ID and classes",
			props:  FormProps{ID: "test-form", Class: []string{"custom", "login-form"}},
			expect: `<form id="test-form" class="custom login-form"></form>`,
		},
		{
			name:   "With action and method",
			props:  FormProps{Action: "/submit", Method: "POST"},
			expect: `<form action="/submit" method="POST" class=""></form>`,
		},
		{
			name:   "With enctype",
			props:  FormProps{Enctype: "multipart/form-data"},
			expect: `<form enctype="multipart/form-data" class=""></form>`,
		},
		{
			name:   "With target",
			props:  FormProps{Target: "_blank"},
			expect: `<form target="_blank" class=""></form>`,
		},
		{
			name:   "With novalidate",
			props:  FormProps{NoValidate: true},
			expect: `<form novalidate class=""></form>`,
		},
		{
			name: "Complete configuration",
			props: FormProps{
				ID:         "contact-form",
				Class:      []string{"form", "contact"},
				Action:     "/contact",
				Method:     "POST",
				Enctype:    "multipart/form-data",
				Target:     "_self",
				NoValidate: true,
			},
			expect: `<form id="contact-form" action="/contact" method="POST" enctype="multipart/form-data" target="_self" novalidate class="form contact"></form>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Form(tt.props).Render(context.Background(), &buf)
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

func TestFormAttributes(t *testing.T) {
	tests := []struct {
		name   string
		props  FormProps
		checks map[string]bool
	}{
		{
			name: "All attributes",
			props: FormProps{
				Action:     "/submit",
				Method:     "POST",
				Enctype:    "application/x-www-form-urlencoded",
				Target:     "_parent",
				NoValidate: true,
			},
			checks: map[string]bool{
				"action":     true,
				"method":     true,
				"enctype":    true,
				"target":     true,
				"novalidate": true,
			},
		},
		{
			name: "Partial attributes",
			props: FormProps{
				Action: "/login",
				Method: "GET",
			},
			checks: map[string]bool{
				"action":     true,
				"method":     true,
				"enctype":    false,
				"target":     false,
				"novalidate": false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Form(tt.props).Render(context.Background(), &buf)
			if err != nil {
				t.Fatalf("render failed: %v", err)
			}
			got := strings.TrimSpace(buf.String())

			for attr, shouldHave := range tt.checks {
				hasAttr := strings.Contains(got, attr+"=")
				if attr == "novalidate" {
					hasAttr = strings.Contains(got, "novalidate")
				}

				if hasAttr != shouldHave {
					t.Errorf("expected %s attribute: %v, got: %v in: %s", attr, shouldHave, hasAttr, got)
				}
			}
		})
	}
}

func TestFormCustomAttributes(t *testing.T) {
	props := FormProps{
		Attributes: templ.Attributes{
			"data-testid":    "contact-form",
			"autocomplete":   "off",
			"accept-charset": "UTF-8",
		},
	}

	var buf strings.Builder
	err := Form(props).Render(context.Background(), &buf)
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
