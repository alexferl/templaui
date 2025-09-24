package notification

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestNotification(t *testing.T) {
	tests := []struct {
		name   string
		props  NotificationProps
		expect string
	}{
		{
			name:   "Default",
			props:  NotificationProps{},
			expect: `<div class="notification" role="alert"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  NotificationProps{ID: "notification1", Class: []string{"custom-notification"}},
			expect: `<div id="notification1" class="notification custom-notification" role="alert"></div>`,
		},
		{
			name:   "With color",
			props:  NotificationProps{Color: IsPrimary},
			expect: `<div class="notification is-primary" role="alert"></div>`,
		},
		{
			name:   "With danger color",
			props:  NotificationProps{Color: IsDanger},
			expect: `<div class="notification is-danger" role="alert"></div>`,
		},
		{
			name:   "With success color",
			props:  NotificationProps{Color: IsSuccess},
			expect: `<div class="notification is-success" role="alert"></div>`,
		},
		{
			name:   "With light variant",
			props:  NotificationProps{Color: IsPrimary, Variant: IsLightVariant},
			expect: `<div class="notification is-primary is-light" role="alert"></div>`,
		},
		{
			name:   "With dark variant",
			props:  NotificationProps{Color: IsInfo, Variant: IsDarkVariant},
			expect: `<div class="notification is-info is-dark" role="alert"></div>`,
		},
		{
			name:   "With delete button (default aria-label)",
			props:  NotificationProps{HasDelete: true},
			expect: `<div class="notification" role="alert"><button class="delete" aria-label="Close notification"></button></div>`,
		},
		{
			name:   "With delete button (custom aria-label)",
			props:  NotificationProps{HasDelete: true, DeleteAriaLabel: "Dismiss success message"},
			expect: `<div class="notification" role="alert"><button class="delete" aria-label="Dismiss success message"></button></div>`,
		},
		{
			name:   "With custom role",
			props:  NotificationProps{Role: "status"},
			expect: `<div class="notification" role="status"></div>`,
		},
		{
			name:   "With custom aria-label",
			props:  NotificationProps{AriaLabel: "Important system message"},
			expect: `<div class="notification" role="alert" aria-label="Important system message"></div>`,
		},
		{
			name:   "With color and delete button",
			props:  NotificationProps{Color: IsDanger, HasDelete: true},
			expect: `<div class="notification is-danger" role="alert"><button class="delete" aria-label="Close notification"></button></div>`,
		},
		{
			name:   "With color, variant, and delete",
			props:  NotificationProps{Color: IsWarning, Variant: IsLightVariant, HasDelete: true, DeleteAriaLabel: "Hide warning banner"},
			expect: `<div class="notification is-warning is-light" role="alert"><button class="delete" aria-label="Hide warning banner"></button></div>`,
		},
		{
			name: "All fields combined",
			props: NotificationProps{
				ID:              "main-notification",
				Class:           []string{"foo", "bar"},
				Color:           IsSuccess,
				Variant:         IsLightVariant,
				HasDelete:       true,
				DeleteAriaLabel: "Dismiss success notification",
				Role:            "status",
				AriaLabel:       "Operation completed successfully",
				Attributes:      templ.Attributes{"data-testid": "notification"},
			},
			expect: `<div id="main-notification" class="notification is-success is-light foo bar" role="status" aria-label="Operation completed successfully" data-testid="notification"><button class="delete" aria-label="Dismiss success notification"></button></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Notification(tt.props).Render(context.Background(), &buf)
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

func TestNotificationColors(t *testing.T) {
	tests := []struct {
		name   string
		color  Color
		expect string
	}{
		{"Black", IsBlack, `<div class="notification is-black" role="alert"></div>`},
		{"Danger", IsDanger, `<div class="notification is-danger" role="alert"></div>`},
		{"Dark", IsDark, `<div class="notification is-dark" role="alert"></div>`},
		{"Info", IsInfo, `<div class="notification is-info" role="alert"></div>`},
		{"Light", IsLight, `<div class="notification is-light" role="alert"></div>`},
		{"Link", IsLink, `<div class="notification is-link" role="alert"></div>`},
		{"Primary", IsPrimary, `<div class="notification is-primary" role="alert"></div>`},
		{"Success", IsSuccess, `<div class="notification is-success" role="alert"></div>`},
		{"Text", IsText, `<div class="notification is-text" role="alert"></div>`},
		{"Warning", IsWarning, `<div class="notification is-warning" role="alert"></div>`},
		{"White", IsWhite, `<div class="notification is-white" role="alert"></div>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Notification(NotificationProps{Color: tt.color}).Render(context.Background(), &buf)
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

func TestNotificationVariants(t *testing.T) {
	tests := []struct {
		name    string
		color   Color
		variant Variant
		expect  string
	}{
		{
			name:    "Primary light variant",
			color:   IsPrimary,
			variant: IsLightVariant,
			expect:  `<div class="notification is-primary is-light" role="alert"></div>`,
		},
		{
			name:    "Primary dark variant",
			color:   IsPrimary,
			variant: IsDarkVariant,
			expect:  `<div class="notification is-primary is-dark" role="alert"></div>`,
		},
		{
			name:    "Danger light variant",
			color:   IsDanger,
			variant: IsLightVariant,
			expect:  `<div class="notification is-danger is-light" role="alert"></div>`,
		},
		{
			name:    "Success dark variant",
			color:   IsSuccess,
			variant: IsDarkVariant,
			expect:  `<div class="notification is-success is-dark" role="alert"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Notification(NotificationProps{Color: tt.color, Variant: tt.variant}).Render(context.Background(), &buf)
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

func TestNotificationDeleteButton(t *testing.T) {
	tests := []struct {
		name            string
		hasDelete       bool
		deleteAriaLabel string
		expect          string
	}{
		{
			name:      "No delete button",
			hasDelete: false,
			expect:    `<div class="notification" role="alert"></div>`,
		},
		{
			name:      "Default delete button",
			hasDelete: true,
			expect:    `<div class="notification" role="alert"><button class="delete" aria-label="Close notification"></button></div>`,
		},
		{
			name:            "Custom delete aria-label",
			hasDelete:       true,
			deleteAriaLabel: "Dismiss error message",
			expect:          `<div class="notification" role="alert"><button class="delete" aria-label="Dismiss error message"></button></div>`,
		},
		{
			name:            "Empty delete aria-label (falls back to default)",
			hasDelete:       true,
			deleteAriaLabel: "",
			expect:          `<div class="notification" role="alert"><button class="delete" aria-label="Close notification"></button></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Notification(NotificationProps{
				HasDelete:       tt.hasDelete,
				DeleteAriaLabel: tt.deleteAriaLabel,
			}).Render(context.Background(), &buf)
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

func TestNotificationAccessibility(t *testing.T) {
	tests := []struct {
		name      string
		role      string
		ariaLabel string
		expect    string
	}{
		{
			name:   "Default role (alert)",
			expect: `<div class="notification" role="alert"></div>`,
		},
		{
			name:   "Custom role (status)",
			role:   "status",
			expect: `<div class="notification" role="status"></div>`,
		},
		{
			name:      "With aria-label",
			ariaLabel: "System update notification",
			expect:    `<div class="notification" role="alert" aria-label="System update notification"></div>`,
		},
		{
			name:      "Custom role and aria-label",
			role:      "region",
			ariaLabel: "User preferences saved",
			expect:    `<div class="notification" role="region" aria-label="User preferences saved"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Notification(NotificationProps{
				Role:      tt.role,
				AriaLabel: tt.ariaLabel,
			}).Render(context.Background(), &buf)
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
