package progress

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestProgress(t *testing.T) {
	tests := []struct {
		name   string
		props  ProgressProps
		expect string
	}{
		{
			name:   "Default (indeterminate)",
			props:  ProgressProps{},
			expect: `<progress class="progress" max="100"></progress>`,
		},
		{
			name:   "With ID and custom classes",
			props:  ProgressProps{ID: "progress1", Class: []string{"custom-progress"}},
			expect: `<progress id="progress1" class="progress custom-progress" max="100"></progress>`,
		},
		{
			name:   "With small size",
			props:  ProgressProps{Size: IsSmall},
			expect: `<progress class="progress is-small" max="100"></progress>`,
		},
		{
			name:   "With medium size",
			props:  ProgressProps{Size: IsMedium},
			expect: `<progress class="progress is-medium" max="100"></progress>`,
		},
		{
			name:   "With large size",
			props:  ProgressProps{Size: IsLarge},
			expect: `<progress class="progress is-large" max="100"></progress>`,
		},
		{
			name:   "With color",
			props:  ProgressProps{Color: IsPrimary},
			expect: `<progress class="progress is-primary" max="100"></progress>`,
		},
		{
			name:   "With danger color",
			props:  ProgressProps{Color: IsDanger},
			expect: `<progress class="progress is-danger" max="100"></progress>`,
		},
		{
			name:   "Determinate progress (30%)",
			props:  ProgressProps{Value: &[]int{30}[0]},
			expect: `<progress class="progress" max="100" value="30"></progress>`,
		},
		{
			name:   "Determinate progress with custom max",
			props:  ProgressProps{Value: &[]int{7}[0], Max: 10},
			expect: `<progress class="progress" max="10" value="7"></progress>`,
		},
		{
			name:   "Zero progress",
			props:  ProgressProps{Value: &[]int{0}[0]},
			expect: `<progress class="progress" max="100" value="0"></progress>`,
		},
		{
			name:   "Full progress",
			props:  ProgressProps{Value: &[]int{100}[0]},
			expect: `<progress class="progress" max="100" value="100"></progress>`,
		},
		{
			name:   "With aria-label",
			props:  ProgressProps{AriaLabel: "File upload progress"},
			expect: `<progress class="progress" max="100" aria-label="File upload progress"></progress>`,
		},
		{
			name:   "With aria-valuetext",
			props:  ProgressProps{Value: &[]int{3}[0], Max: 10, AriaValueText: "3 of 10 files"},
			expect: `<progress class="progress" max="10" value="3" aria-valuetext="3 of 10 files"></progress>`,
		},
		{
			name:   "With size and color",
			props:  ProgressProps{Size: IsLarge, Color: IsSuccess},
			expect: `<progress class="progress is-large is-success" max="100"></progress>`,
		},
		{
			name: "All fields combined",
			props: ProgressProps{
				ID:            "main-progress",
				Class:         []string{"foo", "bar"},
				Size:          IsMedium,
				Color:         IsInfo,
				Value:         &[]int{75}[0],
				Max:           100,
				AriaLabel:     "Download progress",
				AriaValueText: "75 percent complete",
				Attributes:    templ.Attributes{"data-testid": "progress"},
			},
			expect: `<progress id="main-progress" class="progress is-medium is-info foo bar" max="100" value="75" aria-label="Download progress" aria-valuetext="75 percent complete" data-testid="progress"></progress>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Progress(tt.props).Render(context.Background(), &buf)
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

func TestProgressSizes(t *testing.T) {
	tests := []struct {
		name   string
		size   Size
		expect string
	}{
		{"Small", IsSmall, `<progress class="progress is-small" max="100"></progress>`},
		{"Medium", IsMedium, `<progress class="progress is-medium" max="100"></progress>`},
		{"Large", IsLarge, `<progress class="progress is-large" max="100"></progress>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Progress(ProgressProps{Size: tt.size}).Render(context.Background(), &buf)
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

func TestProgressColors(t *testing.T) {
	tests := []struct {
		name   string
		color  Color
		expect string
	}{
		{"Black", IsBlack, `<progress class="progress is-black" max="100"></progress>`},
		{"Danger", IsDanger, `<progress class="progress is-danger" max="100"></progress>`},
		{"Dark", IsDark, `<progress class="progress is-dark" max="100"></progress>`},
		{"Info", IsInfo, `<progress class="progress is-info" max="100"></progress>`},
		{"Light", IsLight, `<progress class="progress is-light" max="100"></progress>`},
		{"Link", IsLink, `<progress class="progress is-link" max="100"></progress>`},
		{"Primary", IsPrimary, `<progress class="progress is-primary" max="100"></progress>`},
		{"Success", IsSuccess, `<progress class="progress is-success" max="100"></progress>`},
		{"Text", IsText, `<progress class="progress is-text" max="100"></progress>`},
		{"Warning", IsWarning, `<progress class="progress is-warning" max="100"></progress>`},
		{"White", IsWhite, `<progress class="progress is-white" max="100"></progress>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Progress(ProgressProps{Color: tt.color}).Render(context.Background(), &buf)
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

func TestProgressValues(t *testing.T) {
	tests := []struct {
		name   string
		props  ProgressProps
		expect string
	}{
		{
			name:   "Indeterminate (no value)",
			props:  ProgressProps{},
			expect: `<progress class="progress" max="100"></progress>`,
		},
		{
			name:   "Zero progress",
			props:  ProgressProps{Value: &[]int{0}[0]},
			expect: `<progress class="progress" max="100" value="0"></progress>`,
		},
		{
			name:   "25% progress",
			props:  ProgressProps{Value: &[]int{25}[0]},
			expect: `<progress class="progress" max="100" value="25"></progress>`,
		},
		{
			name:   "50% progress",
			props:  ProgressProps{Value: &[]int{50}[0]},
			expect: `<progress class="progress" max="100" value="50"></progress>`,
		},
		{
			name:   "100% progress",
			props:  ProgressProps{Value: &[]int{100}[0]},
			expect: `<progress class="progress" max="100" value="100"></progress>`,
		},
		{
			name:   "Custom max (3 of 5)",
			props:  ProgressProps{Value: &[]int{3}[0], Max: 5},
			expect: `<progress class="progress" max="5" value="3"></progress>`,
		},
		{
			name:   "Custom max (7 of 10)",
			props:  ProgressProps{Value: &[]int{7}[0], Max: 10},
			expect: `<progress class="progress" max="10" value="7"></progress>`,
		},
		{
			name:   "Zero max defaults to 100",
			props:  ProgressProps{Max: 0},
			expect: `<progress class="progress" max="100"></progress>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Progress(tt.props).Render(context.Background(), &buf)
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

func TestProgressAccessibility(t *testing.T) {
	tests := []struct {
		name          string
		ariaLabel     string
		ariaValueText string
		expect        string
	}{
		{
			name:   "No accessibility attributes",
			expect: `<progress class="progress" max="100"></progress>`,
		},
		{
			name:      "With aria-label",
			ariaLabel: "Loading progress",
			expect:    `<progress class="progress" max="100" aria-label="Loading progress"></progress>`,
		},
		{
			name:          "With aria-valuetext",
			ariaValueText: "Step 2 of 5",
			expect:        `<progress class="progress" max="100" aria-valuetext="Step 2 of 5"></progress>`,
		},
		{
			name:          "With both aria attributes",
			ariaLabel:     "File upload",
			ariaValueText: "3 of 10 files uploaded",
			expect:        `<progress class="progress" max="100" aria-label="File upload" aria-valuetext="3 of 10 files uploaded"></progress>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Progress(ProgressProps{
				AriaLabel:     tt.ariaLabel,
				AriaValueText: tt.ariaValueText,
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

// Helper function for creating pointer values in tests
func intPtr(i int) *int {
	return &i
}

func TestProgressWithHelperFunction(t *testing.T) {
	tests := []struct {
		name   string
		value  *int
		expect string
	}{
		{
			name:   "Using helper - nil value",
			value:  nil,
			expect: `<progress class="progress" max="100"></progress>`,
		},
		{
			name:   "Using helper - 45% progress",
			value:  intPtr(45),
			expect: `<progress class="progress" max="100" value="45"></progress>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Progress(ProgressProps{Value: tt.value}).Render(context.Background(), &buf)
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
