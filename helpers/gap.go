package helpers

// Gap represents Bulma's gap helper system for CSS Grid and Flexbox layouts.
//
// Use these constants to apply Bulma's gap helpers which control spacing
// between grid items, flex items, and columns. Gap values are based on
// rem units with a scale from 0 to 8 (0rem to 2rem). The 0.5 increments
// provide fine-grained spacing control. These helpers work with CSS Grid
// containers, Flexbox containers, and Bulma's column system.
type Gap string

const (
	// Column Gap Helpers - Control horizontal spacing between columns/grid items
	IsColumnGap0  Gap = "is-column-gap-0"   // Remove horizontal gap (0rem)
	IsColumnGap05 Gap = "is-column-gap-0.5" // Minimal horizontal gap (0.125rem)
	IsColumnGap1  Gap = "is-column-gap-1"   // Small horizontal gap (0.25rem)
	IsColumnGap15 Gap = "is-column-gap-1.5" // Small-medium horizontal gap (0.375rem)
	IsColumnGap2  Gap = "is-column-gap-2"   // Medium horizontal gap (0.5rem)
	IsColumnGap25 Gap = "is-column-gap-2.5" // Medium-large horizontal gap (0.625rem)
	IsColumnGap3  Gap = "is-column-gap-3"   // Standard horizontal gap (0.75rem)
	IsColumnGap35 Gap = "is-column-gap-3.5" // Standard-large horizontal gap (0.875rem)
	IsColumnGap4  Gap = "is-column-gap-4"   // Large horizontal gap (1rem)
	IsColumnGap45 Gap = "is-column-gap-4.5" // Large-extra horizontal gap (1.125rem)
	IsColumnGap5  Gap = "is-column-gap-5"   // Extra large horizontal gap (1.25rem)
	IsColumnGap55 Gap = "is-column-gap-5.5" // Extra-maximum horizontal gap (1.375rem)
	IsColumnGap6  Gap = "is-column-gap-6"   // Maximum horizontal gap (1.5rem)
	IsColumnGap65 Gap = "is-column-gap-6.5" // Maximum-plus horizontal gap (1.625rem)
	IsColumnGap7  Gap = "is-column-gap-7"   // Super large horizontal gap (1.75rem)
	IsColumnGap75 Gap = "is-column-gap-7.5" // Super-maximum horizontal gap (1.875rem)
	IsColumnGap8  Gap = "is-column-gap-8"   // Largest horizontal gap (2rem)

	// Universal Gap Helpers - Control spacing in both directions
	IsGap0  Gap = "is-gap-0"   // Remove all gaps (0rem)
	IsGap05 Gap = "is-gap-0.5" // Minimal gap in both directions (0.125rem)
	IsGap1  Gap = "is-gap-1"   // Small gap in both directions (0.25rem)
	IsGap15 Gap = "is-gap-1.5" // Small-medium gap in both directions (0.375rem)
	IsGap2  Gap = "is-gap-2"   // Medium gap in both directions (0.5rem)
	IsGap25 Gap = "is-gap-2.5" // Medium-large gap in both directions (0.625rem)
	IsGap3  Gap = "is-gap-3"   // Standard gap in both directions (0.75rem)
	IsGap35 Gap = "is-gap-3.5" // Standard-large gap in both directions (0.875rem)
	IsGap4  Gap = "is-gap-4"   // Large gap in both directions (1rem)
	IsGap45 Gap = "is-gap-4.5" // Large-extra gap in both directions (1.125rem)
	IsGap5  Gap = "is-gap-5"   // Extra large gap in both directions (1.25rem)
	IsGap55 Gap = "is-gap-5.5" // Extra-maximum gap in both directions (1.375rem)
	IsGap6  Gap = "is-gap-6"   // Maximum gap in both directions (1.5rem)
	IsGap65 Gap = "is-gap-6.5" // Maximum-plus gap in both directions (1.625rem)
	IsGap7  Gap = "is-gap-7"   // Super large gap in both directions (1.75rem)
	IsGap75 Gap = "is-gap-7.5" // Super-maximum gap in both directions (1.875rem)
	IsGap8  Gap = "is-gap-8"   // Largest gap in both directions (2rem)

	// Row Gap Helpers - Control vertical spacing between rows/grid items
	IsRowGap0  Gap = "is-row-gap-0"   // Remove vertical gap (0rem)
	IsRowGap05 Gap = "is-row-gap-0.5" // Minimal vertical gap (0.125rem)
	IsRowGap1  Gap = "is-row-gap-1"   // Small vertical gap (0.25rem)
	IsRowGap15 Gap = "is-row-gap-1.5" // Small-medium vertical gap (0.375rem)
	IsRowGap2  Gap = "is-row-gap-2"   // Medium vertical gap (0.5rem)
	IsRowGap25 Gap = "is-row-gap-2.5" // Medium-large vertical gap (0.625rem)
	IsRowGap3  Gap = "is-row-gap-3"   // Standard vertical gap (0.75rem)
	IsRowGap35 Gap = "is-row-gap-3.5" // Standard-large vertical gap (0.875rem)
	IsRowGap4  Gap = "is-row-gap-4"   // Large vertical gap (1rem)
	IsRowGap45 Gap = "is-row-gap-4.5" // Large-extra vertical gap (1.125rem)
	IsRowGap5  Gap = "is-row-gap-5"   // Extra large vertical gap (1.25rem)
	IsRowGap55 Gap = "is-row-gap-5.5" // Extra-maximum vertical gap (1.375rem)
	IsRowGap6  Gap = "is-row-gap-6"   // Maximum vertical gap (1.5rem)
	IsRowGap65 Gap = "is-row-gap-6.5" // Maximum-plus vertical gap (1.625rem)
	IsRowGap7  Gap = "is-row-gap-7"   // Super large vertical gap (1.75rem)
	IsRowGap75 Gap = "is-row-gap-7.5" // Super-maximum vertical gap (1.875rem)
	IsRowGap8  Gap = "is-row-gap-8"   // Largest vertical gap (2rem)
)
