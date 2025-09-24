package helpers

// Overflow represents Bulma's overflow control helper system.
//
// Use these constants to apply Bulma's overflow helpers which control
// how content is handled when it exceeds the bounds of its container.
// These helpers map directly to CSS overflow properties and provide
// fine-grained control over both horizontal (x-axis) and vertical
// (y-axis) overflow behavior. Essential for creating scrollable areas,
// clipping content, and managing container boundaries.
type Overflow string

const (
	// Legacy Overflow Helper - Simplified overflow control
	IsClipped Overflow = "is-clipped" // Hide overflow content (overflow: hidden)

	// Universal Overflow Helpers - Control both x and y axes
	IsOverflowAuto    Overflow = "is-overflow-auto"    // Show scrollbars only when needed
	IsOverflowClip    Overflow = "is-overflow-clip"    // Clip content at container boundary
	IsOverflowHidden  Overflow = "is-overflow-hidden"  // Hide overflow content completely
	IsOverflowScroll  Overflow = "is-overflow-scroll"  // Always show scrollbars
	IsOverflowVisible Overflow = "is-overflow-visible" // Allow content to overflow visibly

	// Horizontal Overflow Helpers - Control x-axis only
	IsOverflowXAuto    Overflow = "is-overflow-x-auto"    // Horizontal scrollbar only when needed
	IsOverflowXClip    Overflow = "is-overflow-x-clip"    // Clip horizontal overflow
	IsOverflowXHidden  Overflow = "is-overflow-x-hidden"  // Hide horizontal overflow
	IsOverflowXScroll  Overflow = "is-overflow-x-scroll"  // Always show horizontal scrollbar
	IsOverflowXVisible Overflow = "is-overflow-x-visible" // Allow horizontal overflow

	// Vertical Overflow Helpers - Control y-axis only
	IsOverflowYAuto    Overflow = "is-overflow-y-auto"    // Vertical scrollbar only when needed
	IsOverflowYClip    Overflow = "is-overflow-y-clip"    // Clip vertical overflow
	IsOverflowYHidden  Overflow = "is-overflow-y-hidden"  // Hide vertical overflow
	IsOverflowYScroll  Overflow = "is-overflow-y-scroll"  // Always show vertical scrollbar
	IsOverflowYVisible Overflow = "is-overflow-y-visible" // Allow vertical overflow
)
