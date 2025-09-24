package helpers

// Spacing represents Bulma's comprehensive spacing helper system.
//
// Use these constants to apply Bulma's spacing helpers which control
// margin and padding on elements. The system uses a scale from 0-6
// corresponding to rem values: 0=0rem, 1=0.25rem, 2=0.5rem, 3=0.75rem,
// 4=1rem, 5=1.25rem, 6=1.5rem. Auto values use CSS auto for centering.
// Supports directional control with prefixes: t(op), r(ight), b(ottom),
// l(eft), x(horizontal), y(vertical).
type Spacing string

const (
	// Universal Margin Helpers - Apply margin to all sides
	M0    Spacing = "m-0"    // Remove all margins (0rem)
	M1    Spacing = "m-1"    // Small margin all sides (0.25rem)
	M2    Spacing = "m-2"    // Medium margin all sides (0.5rem)
	M3    Spacing = "m-3"    // Standard margin all sides (0.75rem)
	M4    Spacing = "m-4"    // Large margin all sides (1rem)
	M5    Spacing = "m-5"    // Extra large margin all sides (1.25rem)
	M6    Spacing = "m-6"    // Maximum margin all sides (1.5rem)
	MAuto Spacing = "m-auto" // Auto margin all sides (centering)

	// Top Margin Helpers - Apply margin to top edge
	MT0    Spacing = "mt-0"    // Remove top margin (0rem)
	MT1    Spacing = "mt-1"    // Small top margin (0.25rem)
	MT2    Spacing = "mt-2"    // Medium top margin (0.5rem)
	MT3    Spacing = "mt-3"    // Standard top margin (0.75rem)
	MT4    Spacing = "mt-4"    // Large top margin (1rem)
	MT5    Spacing = "mt-5"    // Extra large top margin (1.25rem)
	MT6    Spacing = "mt-6"    // Maximum top margin (1.5rem)
	MTAuto Spacing = "mt-auto" // Auto top margin

	// Right Margin Helpers - Apply margin to right edge
	MR0    Spacing = "mr-0"    // Remove right margin (0rem)
	MR1    Spacing = "mr-1"    // Small right margin (0.25rem)
	MR2    Spacing = "mr-2"    // Medium right margin (0.5rem)
	MR3    Spacing = "mr-3"    // Standard right margin (0.75rem)
	MR4    Spacing = "mr-4"    // Large right margin (1rem)
	MR5    Spacing = "mr-5"    // Extra large right margin (1.25rem)
	MR6    Spacing = "mr-6"    // Maximum right margin (1.5rem)
	MRAuto Spacing = "mr-auto" // Auto right margin

	// Bottom Margin Helpers - Apply margin to bottom edge
	MB0    Spacing = "mb-0"    // Remove bottom margin (0rem)
	MB1    Spacing = "mb-1"    // Small bottom margin (0.25rem)
	MB2    Spacing = "mb-2"    // Medium bottom margin (0.5rem)
	MB3    Spacing = "mb-3"    // Standard bottom margin (0.75rem)
	MB4    Spacing = "mb-4"    // Large bottom margin (1rem)
	MB5    Spacing = "mb-5"    // Extra large bottom margin (1.25rem)
	MB6    Spacing = "mb-6"    // Maximum bottom margin (1.5rem)
	MBAuto Spacing = "mb-auto" // Auto bottom margin

	// Left Margin Helpers - Apply margin to left edge
	ML0    Spacing = "ml-0"    // Remove left margin (0rem)
	ML1    Spacing = "ml-1"    // Small left margin (0.25rem)
	ML2    Spacing = "ml-2"    // Medium left margin (0.5rem)
	ML3    Spacing = "ml-3"    // Standard left margin (0.75rem)
	ML4    Spacing = "ml-4"    // Large left margin (1rem)
	ML5    Spacing = "ml-5"    // Extra large left margin (1.25rem)
	ML6    Spacing = "ml-6"    // Maximum left margin (1.5rem)
	MLAuto Spacing = "ml-auto" // Auto left margin

	// Horizontal Margin Helpers - Apply margin to left and right
	MX0    Spacing = "mx-0"    // Remove horizontal margins (0rem)
	MX1    Spacing = "mx-1"    // Small horizontal margins (0.25rem)
	MX2    Spacing = "mx-2"    // Medium horizontal margins (0.5rem)
	MX3    Spacing = "mx-3"    // Standard horizontal margins (0.75rem)
	MX4    Spacing = "mx-4"    // Large horizontal margins (1rem)
	MX5    Spacing = "mx-5"    // Extra large horizontal margins (1.25rem)
	MX6    Spacing = "mx-6"    // Maximum horizontal margins (1.5rem)
	MXAuto Spacing = "mx-auto" // Auto horizontal margins (horizontal centering)

	// Vertical Margin Helpers - Apply margin to top and bottom
	MY0    Spacing = "my-0"    // Remove vertical margins (0rem)
	MY1    Spacing = "my-1"    // Small vertical margins (0.25rem)
	MY2    Spacing = "my-2"    // Medium vertical margins (0.5rem)
	MY3    Spacing = "my-3"    // Standard vertical margins (0.75rem)
	MY4    Spacing = "my-4"    // Large vertical margins (1rem)
	MY5    Spacing = "my-5"    // Extra large vertical margins (1.25rem)
	MY6    Spacing = "my-6"    // Maximum vertical margins (1.5rem)
	MYAuto Spacing = "my-auto" // Auto vertical margins

	// Universal Padding Helpers - Apply padding to all sides
	P0    Spacing = "p-0"    // Remove all padding (0rem)
	P1    Spacing = "p-1"    // Small padding all sides (0.25rem)
	P2    Spacing = "p-2"    // Medium padding all sides (0.5rem)
	P3    Spacing = "p-3"    // Standard padding all sides (0.75rem)
	P4    Spacing = "p-4"    // Large padding all sides (1rem)
	P5    Spacing = "p-5"    // Extra large padding all sides (1.25rem)
	P6    Spacing = "p-6"    // Maximum padding all sides (1.5rem)
	PAuto Spacing = "p-auto" // Auto padding all sides

	// Top Padding Helpers - Apply padding to top edge
	PT0    Spacing = "pt-0"    // Remove top padding (0rem)
	PT1    Spacing = "pt-1"    // Small top padding (0.25rem)
	PT2    Spacing = "pt-2"    // Medium top padding (0.5rem)
	PT3    Spacing = "pt-3"    // Standard top padding (0.75rem)
	PT4    Spacing = "pt-4"    // Large top padding (1rem)
	PT5    Spacing = "pt-5"    // Extra large top padding (1.25rem)
	PT6    Spacing = "pt-6"    // Maximum top padding (1.5rem)
	PTAuto Spacing = "pt-auto" // Auto top padding

	// Right Padding Helpers - Apply padding to right edge
	PR0    Spacing = "pr-0"    // Remove right padding (0rem)
	PR1    Spacing = "pr-1"    // Small right padding (0.25rem)
	PR2    Spacing = "pr-2"    // Medium right padding (0.5rem)
	PR3    Spacing = "pr-3"    // Standard right padding (0.75rem)
	PR4    Spacing = "pr-4"    // Large right padding (1rem)
	PR5    Spacing = "pr-5"    // Extra large right padding (1.25rem)
	PR6    Spacing = "pr-6"    // Maximum right padding (1.5rem)
	PRAuto Spacing = "pr-auto" // Auto right padding

	// Bottom Padding Helpers - Apply padding to bottom edge
	PB0    Spacing = "pb-0"    // Remove bottom padding (0rem)
	PB1    Spacing = "pb-1"    // Small bottom padding (0.25rem)
	PB2    Spacing = "pb-2"    // Medium bottom padding (0.5rem)
	PB3    Spacing = "pb-3"    // Standard bottom padding (0.75rem)
	PB4    Spacing = "pb-4"    // Large bottom padding (1rem)
	PB5    Spacing = "pb-5"    // Extra large bottom padding (1.25rem)
	PB6    Spacing = "pb-6"    // Maximum bottom padding (1.5rem)
	PBAuto Spacing = "pb-auto" // Auto bottom padding

	// Left Padding Helpers - Apply padding to left edge
	PL0    Spacing = "pl-0"    // Remove left padding (0rem)
	PL1    Spacing = "pl-1"    // Small left padding (0.25rem)
	PL2    Spacing = "pl-2"    // Medium left padding (0.5rem)
	PL3    Spacing = "pl-3"    // Standard left padding (0.75rem)
	PL4    Spacing = "pl-4"    // Large left padding (1rem)
	PL5    Spacing = "pl-5"    // Extra large left padding (1.25rem)
	PL6    Spacing = "pl-6"    // Maximum left padding (1.5rem)
	PLAuto Spacing = "pl-auto" // Auto left padding

	// Horizontal Padding Helpers - Apply padding to left and right
	PX0    Spacing = "px-0"    // Remove horizontal padding (0rem)
	PX1    Spacing = "px-1"    // Small horizontal padding (0.25rem)
	PX2    Spacing = "px-2"    // Medium horizontal padding (0.5rem)
	PX3    Spacing = "px-3"    // Standard horizontal padding (0.75rem)
	PX4    Spacing = "px-4"    // Large horizontal padding (1rem)
	PX5    Spacing = "px-5"    // Extra large horizontal padding (1.25rem)
	PX6    Spacing = "px-6"    // Maximum horizontal padding (1.5rem)
	PXAuto Spacing = "px-auto" // Auto horizontal padding

	// Vertical Padding Helpers - Apply padding to top and bottom
	PY0    Spacing = "py-0"    // Remove vertical padding (0rem)
	PY1    Spacing = "py-1"    // Small vertical padding (0.25rem)
	PY2    Spacing = "py-2"    // Medium vertical padding (0.5rem)
	PY3    Spacing = "py-3"    // Standard vertical padding (0.75rem)
	PY4    Spacing = "py-4"    // Large vertical padding (1rem)
	PY5    Spacing = "py-5"    // Extra large vertical padding (1.25rem)
	PY6    Spacing = "py-6"    // Maximum vertical padding (1.5rem)
	PYAuto Spacing = "py-auto" // Auto vertical padding
)
