package helpers

// Position represents Bulma's CSS positioning helper system.
//
// Use these constants to apply Bulma's position helpers which control
// element positioning behavior and layering. These helpers map directly
// to CSS position properties and provide essential layout control for
// overlays, sticky elements, and precise element positioning. The overlay
// helper provides additional functionality for covering parent elements.
type Position string

const (
	// Special Position Helpers - Enhanced positioning behavior
	IsOverlay Position = "is-overlay" // Cover first positioned parent completely (absolute + full coverage)

	// Standard Position Helpers - Direct CSS position property mapping
	IsPositionAbsolute Position = "is-position-absolute" // Position absolutely relative to nearest positioned ancestor
	IsPositionFixed    Position = "is-position-fixed"    // Position fixed relative to viewport (stays during scroll)
	IsPositionRelative Position = "is-position-relative" // Position relative to normal document flow
	IsPositionStatic   Position = "is-position-static"   // Use normal document flow positioning (default)
	IsPositionSticky   Position = "is-position-sticky"   // Stick to container edge when scrolling (hybrid relative/fixed)
)
