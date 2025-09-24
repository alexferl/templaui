package delete

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestDelete(t *testing.T) {
	tests := []struct {
		name   string
		props  DeleteProps
		expect string
	}{
		{
			name:   "Default",
			props:  DeleteProps{},
			expect: `<button class="delete"></button>`,
		},
		{
			name:   "With ID and custom classes",
			props:  DeleteProps{ID: "delete1", Class: []string{"custom-delete"}},
			expect: `<button id="delete1" class="delete custom-delete"></button>`,
		},
		{
			name:   "With small size",
			props:  DeleteProps{Size: IsSmall},
			expect: `<button class="delete is-small"></button>`,
		},
		{
			name:   "With medium size",
			props:  DeleteProps{Size: IsMedium},
			expect: `<button class="delete is-medium"></button>`,
		},
		{
			name:   "With large size",
			props:  DeleteProps{Size: IsLarge},
			expect: `<button class="delete is-large"></button>`,
		},
		{
			name: "With attributes",
			props: DeleteProps{
				ID:         "close-btn",
				Attributes: templ.Attributes{"aria-label": "close"},
			},
			expect: `<button id="close-btn" class="delete" aria-label="close"></button>`,
		},
		{
			name: "All fields combined",
			props: DeleteProps{
				ID:         "main-delete",
				Class:      []string{"foo", "bar"},
				Size:       IsLarge,
				Attributes: templ.Attributes{"aria-label": "delete", "data-action": "remove"},
			},
			expect: `<button id="main-delete" class="delete is-large foo bar" aria-label="delete" data-action="remove"></button>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Delete(tt.props).Render(context.Background(), &buf)
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
