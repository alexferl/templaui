package helpers

// Float represents Bulma's float and clear helper system.
//
// Use these constants to apply Bulma's float helpers which control
// element positioning and clearing behavior. Float helpers are legacy
// CSS properties primarily used for text wrapping around elements.
// For modern layouts, consider using flexbox or CSS Grid instead.
// Includes both standard float/clear properties and Bulma's pulled helpers.
type Float string

const (
	// Clear Helpers - Control clearing of floated elements
	IsClearBoth  Float = "is-clear-both"  // Clear floated elements on both sides
	IsClearLeft  Float = "is-clear-left"  // Clear floated elements on the left side
	IsClearNone  Float = "is-clear-none"  // Do not clear any floated elements
	IsClearRight Float = "is-clear-right" // Clear floated elements on the right side

	// Clearfix Helper - Fix container collapse with floated children
	IsClearfix Float = "is-clearfix" // Apply clearfix to contain floated children

	// Float Helpers - Control element floating behavior
	IsFloatLeft  Float = "is-float-left"  // Float element to the left side
	IsFloatNone  Float = "is-float-none"  // Remove floating (default behavior)
	IsFloatRight Float = "is-float-right" // Float element to the right side

	// Pulled Helpers - Bulma's semantic float alternatives
	IsPulledLeft  Float = "is-pulled-left"  // Pull element to the left (same as float: left)
	IsPulledRight Float = "is-pulled-right" // Pull element to the right (same as float: right)
)
