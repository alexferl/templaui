package form

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestControl(t *testing.T) {
	tests := []struct {
		name   string
		props  ControlProps
		expect string
	}{
		{
			name:   "Default",
			props:  ControlProps{},
			expect: `<div class="control"></div>`,
		},
		{
			name:   "With ID and classes",
			props:  ControlProps{ID: "test-control", Class: []string{"custom", "wrapper"}},
			expect: `<div id="test-control" class="control custom wrapper"></div>`,
		},
		{
			name:   "Has icons left",
			props:  ControlProps{HasIconsLeft: true},
			expect: `<div class="control has-icons-left"></div>`,
		},
		{
			name:   "Has icons right",
			props:  ControlProps{HasIconsRight: true},
			expect: `<div class="control has-icons-right"></div>`,
		},
		{
			name:   "Has icons both sides",
			props:  ControlProps{HasIconsLeft: true, HasIconsRight: true},
			expect: `<div class="control has-icons-left has-icons-right"></div>`,
		},
		{
			name:   "Is expanded",
			props:  ControlProps{IsExpanded: true},
			expect: `<div class="control is-expanded"></div>`,
		},
		{
			name:   "Is loading",
			props:  ControlProps{IsLoading: true},
			expect: `<div class="control is-loading"></div>`,
		},
		{
			name: "All modifiers",
			props: ControlProps{
				HasIconsLeft:  true,
				HasIconsRight: true,
				IsExpanded:    true,
				IsLoading:     true,
			},
			expect: `<div class="control has-icons-left has-icons-right is-expanded is-loading"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Control(tt.props).Render(context.Background(), &buf)
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

func TestControlModifiers(t *testing.T) {
	tests := []struct {
		name             string
		props            ControlProps
		expectIconsLeft  bool
		expectIconsRight bool
		expectExpanded   bool
		expectLoading    bool
	}{
		{"No modifiers", ControlProps{}, false, false, false, false},
		{"Icons left", ControlProps{HasIconsLeft: true}, true, false, false, false},
		{"Icons right", ControlProps{HasIconsRight: true}, false, true, false, false},
		{"Both icons", ControlProps{HasIconsLeft: true, HasIconsRight: true}, true, true, false, false},
		{"Expanded", ControlProps{IsExpanded: true}, false, false, true, false},
		{"Loading", ControlProps{IsLoading: true}, false, false, false, true},
		{"Mixed", ControlProps{HasIconsLeft: true, IsLoading: true}, true, false, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Control(tt.props).Render(context.Background(), &buf)
			if err != nil {
				t.Fatalf("render failed: %v", err)
			}
			got := strings.TrimSpace(buf.String())

			hasIconsLeft := strings.Contains(got, "has-icons-left")
			hasIconsRight := strings.Contains(got, "has-icons-right")
			hasExpanded := strings.Contains(got, "is-expanded")
			hasLoading := strings.Contains(got, "is-loading")

			if hasIconsLeft != tt.expectIconsLeft {
				t.Errorf("expected has-icons-left: %v, got: %v", tt.expectIconsLeft, hasIconsLeft)
			}
			if hasIconsRight != tt.expectIconsRight {
				t.Errorf("expected has-icons-right: %v, got: %v", tt.expectIconsRight, hasIconsRight)
			}
			if hasExpanded != tt.expectExpanded {
				t.Errorf("expected is-expanded: %v, got: %v", tt.expectExpanded, hasExpanded)
			}
			if hasLoading != tt.expectLoading {
				t.Errorf("expected is-loading: %v, got: %v", tt.expectLoading, hasLoading)
			}
		})
	}
}

func TestControlAttributes(t *testing.T) {
	props := ControlProps{
		Attributes: templ.Attributes{
			"data-testid": "control",
			"role":        "group",
		},
	}

	var buf strings.Builder
	err := Control(props).Render(context.Background(), &buf)
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
