package form

// Color represents form control color modifiers for validation states and styling.
type Color string

const (
	IsBlack   Color = "is-black"   // Black form control styling
	IsDark    Color = "is-dark"    // Dark form control styling
	IsDanger  Color = "is-danger"  // Danger/error state (red) - validation failures
	IsInfo    Color = "is-info"    // Info state (blue) - informational feedback
	IsLight   Color = "is-light"   // Light form control styling
	IsLink    Color = "is-link"    // Link-style form control coloring
	IsPrimary Color = "is-primary" // Primary brand color styling
	IsSuccess Color = "is-success" // Success state (green) - validation success
	IsText    Color = "is-text"    // Text-only form control styling
	IsWarning Color = "is-warning" // Warning state (yellow) - validation warnings
	IsWhite   Color = "is-white"   // White form control styling
)

// Size represents form control size modifiers for consistent scaling.
type Size string

const (
	IsSmall  Size = "is-small"  // Small form control size
	IsNormal Size = "is-normal" // Normal form control size (default)
	IsMedium Size = "is-medium" // Medium form control size
	IsLarge  Size = "is-large"  // Large form control size
)

// IconPosition represents icon positioning within form controls.
type IconPosition string

const (
	IsLeft  IconPosition = "is-left"  // Position icon on the left side of form control
	IsRight IconPosition = "is-right" // Position icon on the right side of form control
)
