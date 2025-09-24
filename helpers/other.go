package helpers

// Other represents Bulma's miscellaneous utility helper system.
//
// Use these constants to apply Bulma's other utility helpers which provide
// various behavioral and visual modifications. These helpers control cursor
// behavior, border radius, shadow effects, and text selection. They are
// commonly used to enhance user experience and modify default element styling.
type Other string

const (
	// Interaction Helpers - Control user interaction behavior
	IsClickable    Other = "is-clickable"    // Apply pointer cursor and indicate clickable element
	IsUnselectable Other = "is-unselectable" // Prevent text selection (user-select: none)

	// Visual Style Helpers - Modify element appearance
	IsRadiusless Other = "is-radiusless" // Remove border-radius from element
	IsShadowless Other = "is-shadowless" // Remove box-shadow from element
)
