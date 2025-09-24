package helpers

// Border represents CSS border radius helper modifiers.
//
// Use these constants to apply Bulma's border radius helpers which
// add consistent rounded corners to any element. Essential for
// creating modern, polished interfaces with unified corner styling
// across buttons, cards, images, and containers. Values follow
// Bulma's design system for consistent visual hierarchy.
type Border string

const (
	HasRadiusLarge   Border = "has-radius-large"   // Large border radius (6px) - prominent elements, cards
	HasRadiusNormal  Border = "has-radius-normal"  // Normal border radius (4px) - default buttons, inputs
	HasRadiusRounded Border = "has-radius-rounded" // Fully rounded (50%) - circular avatars, pills, badges
	HasRadiusSmall   Border = "has-radius-small"   // Small border radius (2px) - subtle rounding, tags
)
