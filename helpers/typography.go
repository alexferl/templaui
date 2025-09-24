package helpers

// Typography represents Bulma's comprehensive typography helper system.
//
// Use these constants to apply Bulma's typography helpers which control
// text size, weight, alignment, transformation, and font families. The size
// system uses 7 levels (1-7) with rem values: 1=3rem, 2=2.5rem, 3=2rem,
// 4=1.5rem, 5=1.25rem, 6=1rem, 7=0.75rem. All helpers support Bulma's
// responsive breakpoints: mobile, tablet, desktop, widescreen, fullhd.
type Typography string

const (
	// Text Transformation Helpers - Control text case and styling
	IsCapitalized Typography = "is-capitalized" // Capitalize first letter of each word
	IsItalic      Typography = "is-italic"      // Italic text style
	IsLowercase   Typography = "is-lowercase"   // Transform text to lowercase
	IsUnderlined  Typography = "is-underlined"  // Add underline decoration
	IsUppercase   Typography = "is-uppercase"   // Transform text to uppercase

	// Font Family Helpers - Control font family
	IsFamilyCode      Typography = "is-family-code"       // Monospace font for code (alias)
	IsFamilyMonospace Typography = "is-family-monospace"  // Monospace font family
	IsFamilyPrimary   Typography = "is-family-primary"    // Primary font family
	IsFamilySansSerif Typography = "is-family-sans-serif" // Sans-serif font family
	IsFamilySecondary Typography = "is-family-secondary"  // Secondary font family

	// Text Alignment Helpers - Control horizontal text alignment
	HasTextCentered               Typography = "has-text-centered"                 // Center align text
	HasTextCenteredDesktop        Typography = "has-text-centered-desktop"         // Center align on desktop and larger
	HasTextCenteredDesktopOnly    Typography = "has-text-centered-desktop-only"    // Center align on desktop only
	HasTextCenteredFullHD         Typography = "has-text-centered-fullhd"          // Center align on fullhd screens
	HasTextCenteredMobile         Typography = "has-text-centered-mobile"          // Center align on mobile only
	HasTextCenteredTablet         Typography = "has-text-centered-tablet"          // Center align on tablet and larger
	HasTextCenteredTabletOnly     Typography = "has-text-centered-tablet-only"     // Center align on tablet only
	HasTextCenteredTouch          Typography = "has-text-centered-touch"           // Center align on mobile and tablet
	HasTextCenteredWidescreen     Typography = "has-text-centered-widescreen"      // Center align on widescreen and larger
	HasTextCenteredWidescreenOnly Typography = "has-text-centered-widescreen-only" // Center align on widescreen only

	HasTextJustified               Typography = "has-text-justified"                 // Justify text alignment
	HasTextJustifiedDesktop        Typography = "has-text-justified-desktop"         // Justify on desktop and larger
	HasTextJustifiedDesktopOnly    Typography = "has-text-justified-desktop-only"    // Justify on desktop only
	HasTextJustifiedFullHD         Typography = "has-text-justified-fullhd"          // Justify on fullhd screens
	HasTextJustifiedMobile         Typography = "has-text-justified-mobile"          // Justify on mobile only
	HasTextJustifiedTablet         Typography = "has-text-justified-tablet"          // Justify on tablet and larger
	HasTextJustifiedTabletOnly     Typography = "has-text-justified-tablet-only"     // Justify on tablet only
	HasTextJustifiedTouch          Typography = "has-text-justified-touch"           // Justify on mobile and tablet
	HasTextJustifiedWidescreen     Typography = "has-text-justified-widescreen"      // Justify on widescreen and larger
	HasTextJustifiedWidescreenOnly Typography = "has-text-justified-widescreen-only" // Justify on widescreen only

	HasTextLeft               Typography = "has-text-left"                 // Left align text
	HasTextLeftDesktop        Typography = "has-text-left-desktop"         // Left align on desktop and larger
	HasTextLeftDesktopOnly    Typography = "has-text-left-desktop-only"    // Left align on desktop only
	HasTextLeftFullHD         Typography = "has-text-left-fullhd"          // Left align on fullhd screens
	HasTextLeftMobile         Typography = "has-text-left-mobile"          // Left align on mobile only
	HasTextLeftTablet         Typography = "has-text-left-tablet"          // Left align on tablet and larger
	HasTextLeftTabletOnly     Typography = "has-text-left-tablet-only"     // Left align on tablet only
	HasTextLeftTouch          Typography = "has-text-left-touch"           // Left align on mobile and tablet
	HasTextLeftWidescreen     Typography = "has-text-left-widescreen"      // Left align on widescreen and larger
	HasTextLeftWidescreenOnly Typography = "has-text-left-widescreen-only" // Left align on widescreen only

	HasTextRight               Typography = "has-text-right"                 // Right align text
	HasTextRightDesktop        Typography = "has-text-right-desktop"         // Right align on desktop and larger
	HasTextRightDesktopOnly    Typography = "has-text-right-desktop-only"    // Right align on desktop only
	HasTextRightFullHD         Typography = "has-text-right-fullhd"          // Right align on fullhd screens
	HasTextRightMobile         Typography = "has-text-right-mobile"          // Right align on mobile only
	HasTextRightTablet         Typography = "has-text-right-tablet"          // Right align on tablet and larger
	HasTextRightTabletOnly     Typography = "has-text-right-tablet-only"     // Right align on tablet only
	HasTextRightTouch          Typography = "has-text-right-touch"           // Right align on mobile and tablet
	HasTextRightWidescreen     Typography = "has-text-right-widescreen"      // Right align on widescreen and larger
	HasTextRightWidescreenOnly Typography = "has-text-right-widescreen-only" // Right align on widescreen only

	// Font Weight Helpers - Control text weight
	HasTextWeightBold      Typography = "has-text-weight-bold"      // Bold text weight (700)
	HasTextWeightExtrabold Typography = "has-text-weight-extrabold" // Extra bold text weight (800)
	HasTextWeightLight     Typography = "has-text-weight-light"     // Light text weight (300)
	HasTextWeightMedium    Typography = "has-text-weight-medium"    // Medium text weight (500)
	HasTextWeightNormal    Typography = "has-text-weight-normal"    // Normal text weight (400)
	HasTextWeightSemibold  Typography = "has-text-weight-semibold"  // Semi-bold text weight (600)

	// Font Size Helpers - Size 1 (3rem) - Largest heading
	IsSize1           Typography = "is-size-1"            // Font size 1 (3rem)
	IsSize1Desktop    Typography = "is-size-1-desktop"    // Size 1 on desktop and larger
	IsSize1FullHD     Typography = "is-size-1-fullhd"     // Size 1 on fullhd screens
	IsSize1Mobile     Typography = "is-size-1-mobile"     // Size 1 on mobile only
	IsSize1Tablet     Typography = "is-size-1-tablet"     // Size 1 on tablet and larger
	IsSize1Touch      Typography = "is-size-1-touch"      // Size 1 on mobile and tablet
	IsSize1Widescreen Typography = "is-size-1-widescreen" // Size 1 on widescreen and larger

	// Font Size Helpers - Size 2 (2.5rem) - Large heading
	IsSize2           Typography = "is-size-2"            // Font size 2 (2.5rem)
	IsSize2Desktop    Typography = "is-size-2-desktop"    // Size 2 on desktop and larger
	IsSize2FullHD     Typography = "is-size-2-fullhd"     // Size 2 on fullhd screens
	IsSize2Mobile     Typography = "is-size-2-mobile"     // Size 2 on mobile only
	IsSize2Tablet     Typography = "is-size-2-tablet"     // Size 2 on tablet and larger
	IsSize2Touch      Typography = "is-size-2-touch"      // Size 2 on mobile and tablet
	IsSize2Widescreen Typography = "is-size-2-widescreen" // Size 2 on widescreen and larger

	// Font Size Helpers - Size 3 (2rem) - Medium heading
	IsSize3           Typography = "is-size-3"            // Font size 3 (2rem)
	IsSize3Desktop    Typography = "is-size-3-desktop"    // Size 3 on desktop and larger
	IsSize3FullHD     Typography = "is-size-3-fullhd"     // Size 3 on fullhd screens
	IsSize3Mobile     Typography = "is-size-3-mobile"     // Size 3 on mobile only
	IsSize3Tablet     Typography = "is-size-3-tablet"     // Size 3 on tablet and larger
	IsSize3Touch      Typography = "is-size-3-touch"      // Size 3 on mobile and tablet
	IsSize3Widescreen Typography = "is-size-3-widescreen" // Size 3 on widescreen and larger

	// Font Size Helpers - Size 4 (1.5rem) - Small heading
	IsSize4           Typography = "is-size-4"            // Font size 4 (1.5rem)
	IsSize4Desktop    Typography = "is-size-4-desktop"    // Size 4 on desktop and larger
	IsSize4FullHD     Typography = "is-size-4-fullhd"     // Size 4 on fullhd screens
	IsSize4Mobile     Typography = "is-size-4-mobile"     // Size 4 on mobile only
	IsSize4Tablet     Typography = "is-size-4-tablet"     // Size 4 on tablet and larger
	IsSize4Touch      Typography = "is-size-4-touch"      // Size 4 on mobile and tablet
	IsSize4Widescreen Typography = "is-size-4-widescreen" // Size 4 on widescreen and larger

	// Font Size Helpers - Size 5 (1.25rem) - Large body text
	IsSize5           Typography = "is-size-5"            // Font size 5 (1.25rem)
	IsSize5Desktop    Typography = "is-size-5-desktop"    // Size 5 on desktop and larger
	IsSize5FullHD     Typography = "is-size-5-fullhd"     // Size 5 on fullhd screens
	IsSize5Mobile     Typography = "is-size-5-mobile"     // Size 5 on mobile only
	IsSize5Tablet     Typography = "is-size-5-tablet"     // Size 5 on tablet and larger
	IsSize5Touch      Typography = "is-size-5-touch"      // Size 5 on mobile and tablet
	IsSize5Widescreen Typography = "is-size-5-widescreen" // Size 5 on widescreen and larger

	// Font Size Helpers - Size 6 (1rem) - Default body text
	IsSize6           Typography = "is-size-6"            // Font size 6 (1rem) - default
	IsSize6Desktop    Typography = "is-size-6-desktop"    // Size 6 on desktop and larger
	IsSize6FullHD     Typography = "is-size-6-fullhd"     // Size 6 on fullhd screens
	IsSize6Mobile     Typography = "is-size-6-mobile"     // Size 6 on mobile only
	IsSize6Tablet     Typography = "is-size-6-tablet"     // Size 6 on tablet and larger
	IsSize6Touch      Typography = "is-size-6-touch"      // Size 6 on mobile and tablet
	IsSize6Widescreen Typography = "is-size-6-widescreen" // Size 6 on widescreen and larger

	// Font Size Helpers - Size 7 (0.75rem) - Small text
	IsSize7           Typography = "is-size-7"            // Font size 7 (0.75rem) - smallest
	IsSize7Desktop    Typography = "is-size-7-desktop"    // Size 7 on desktop and larger
	IsSize7FullHD     Typography = "is-size-7-fullhd"     // Size 7 on fullhd screens
	IsSize7Mobile     Typography = "is-size-7-mobile"     // Size 7 on mobile only
	IsSize7Tablet     Typography = "is-size-7-tablet"     // Size 7 on tablet and larger
	IsSize7Touch      Typography = "is-size-7-touch"      // Size 7 on mobile and tablet
	IsSize7Widescreen Typography = "is-size-7-widescreen" // Size 7 on widescreen and larger
)
