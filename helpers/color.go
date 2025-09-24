package helpers

// Color represents Bulma's comprehensive color helper system.
//
// Use these constants to apply Bulma's color helpers which provide
// background colors, text colors, and interactive states for any element.
// Includes the full range of semantic colors (primary, danger, success, etc.),
// monochrome variations (black, white, grey), and fine-grained opacity
// levels (00-100) with corresponding inverted colors for accessibility.
type Color string

const (
	// Background Color Helpers - Base Colors
	HasBackground        Color = "has-background"         // Inherit background color from parent
	HasBackgroundBlack   Color = "has-background-black"   // Pure black background
	HasBackgroundDanger  Color = "has-background-danger"  // Danger/error background (red)
	HasBackgroundDark    Color = "has-background-dark"    // Dark background color
	HasBackgroundInfo    Color = "has-background-info"    // Info background (blue)
	HasBackgroundLight   Color = "has-background-light"   // Light background color
	HasBackgroundLink    Color = "has-background-link"    // Link background color
	HasBackgroundPrimary Color = "has-background-primary" // Primary brand background
	HasBackgroundSuccess Color = "has-background-success" // Success background (green)
	HasBackgroundText    Color = "has-background-text"    // Text color as background
	HasBackgroundWarning Color = "has-background-warning" // Warning background (yellow)
	HasBackgroundWhite   Color = "has-background-white"   // Pure white background

	// Background Color Helpers - Black Opacity Variations (00-100)
	HasBackgroundBlack00  Color = "has-background-black-00"  // Black 0% opacity (transparent)
	HasBackgroundBlack05  Color = "has-background-black-05"  // Black 5% opacity
	HasBackgroundBlack10  Color = "has-background-black-10"  // Black 10% opacity
	HasBackgroundBlack15  Color = "has-background-black-15"  // Black 15% opacity
	HasBackgroundBlack20  Color = "has-background-black-20"  // Black 20% opacity
	HasBackgroundBlack25  Color = "has-background-black-25"  // Black 25% opacity
	HasBackgroundBlack30  Color = "has-background-black-30"  // Black 30% opacity
	HasBackgroundBlack35  Color = "has-background-black-35"  // Black 35% opacity
	HasBackgroundBlack40  Color = "has-background-black-40"  // Black 40% opacity
	HasBackgroundBlack45  Color = "has-background-black-45"  // Black 45% opacity
	HasBackgroundBlack50  Color = "has-background-black-50"  // Black 50% opacity
	HasBackgroundBlack55  Color = "has-background-black-55"  // Black 55% opacity
	HasBackgroundBlack60  Color = "has-background-black-60"  // Black 60% opacity
	HasBackgroundBlack65  Color = "has-background-black-65"  // Black 65% opacity
	HasBackgroundBlack70  Color = "has-background-black-70"  // Black 70% opacity
	HasBackgroundBlack75  Color = "has-background-black-75"  // Black 75% opacity
	HasBackgroundBlack80  Color = "has-background-black-80"  // Black 80% opacity
	HasBackgroundBlack85  Color = "has-background-black-85"  // Black 85% opacity
	HasBackgroundBlack90  Color = "has-background-black-90"  // Black 90% opacity
	HasBackgroundBlack95  Color = "has-background-black-95"  // Black 95% opacity
	HasBackgroundBlack100 Color = "has-background-black-100" // Black 100% opacity (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid)

	// Background Color Helpers - Black Opacity Inverts (for contrast)
	HasBackgroundBlack00Invert  Color = "has-background-black-00-invert"  // Inverted contrast for black-00
	HasBackgroundBlack05Invert  Color = "has-background-black-05-invert"  // Inverted contrast for black-05
	HasBackgroundBlack10Invert  Color = "has-background-black-10-invert"  // Inverted contrast for black-10
	HasBackgroundBlack15Invert  Color = "has-background-black-15-invert"  // Inverted contrast for black-15
	HasBackgroundBlack20Invert  Color = "has-background-black-20-invert"  // Inverted contrast for black-20
	HasBackgroundBlack25Invert  Color = "has-background-black-25-invert"  // Inverted contrast for black-25
	HasBackgroundBlack30Invert  Color = "has-background-black-30-invert"  // Inverted contrast for black-30
	HasBackgroundBlack35Invert  Color = "has-background-black-35-invert"  // Inverted contrast for black-35
	HasBackgroundBlack40Invert  Color = "has-background-black-40-invert"  // Inverted contrast for black-40
	HasBackgroundBlack45Invert  Color = "has-background-black-45-invert"  // Inverted contrast for black-45
	HasBackgroundBlack50Invert  Color = "has-background-black-50-invert"  // Inverted contrast for black-50
	HasBackgroundBlack55Invert  Color = "has-background-black-55-invert"  // Inverted contrast for black-55
	HasBackgroundBlack60Invert  Color = "has-background-black-60-invert"  // Inverted contrast for black-60
	HasBackgroundBlack65Invert  Color = "has-background-black-65-invert"  // Inverted contrast for black-65
	HasBackgroundBlack70Invert  Color = "has-background-black-70-invert"  // Inverted contrast for black-70
	HasBackgroundBlack75Invert  Color = "has-background-black-75-invert"  // Inverted contrast for black-75
	HasBackgroundBlack80Invert  Color = "has-background-black-80-invert"  // Inverted contrast for black-80
	HasBackgroundBlack85Invert  Color = "has-background-black-85-invert"  // Inverted contrast for black-85
	HasBackgroundBlack90Invert  Color = "has-background-black-90-invert"  // Inverted contrast for black-90
	HasBackgroundBlack95Invert  Color = "has-background-black-95-invert"  // Inverted contrast for black-95
	HasBackgroundBlack100Invert Color = "has-background-black-100-invert" // Inverted contrast for black-100

	// Background Color Helpers - Black Semantic Variations
	HasBackgroundBlackBis         Color = "has-background-black-bis"          // Black-bis monochrome variant
	HasBackgroundBlackTer         Color = "has-background-black-ter"          // Black-ter monochrome variant
	HasBackgroundBlackInvert      Color = "has-background-black-invert"       // Black inverted color
	HasBackgroundBlackLight       Color = "has-background-black-light"        // Light variant of black
	HasBackgroundBlackDark        Color = "has-background-black-dark"         // Dark variant of black
	HasBackgroundBlackBold        Color = "has-background-black-bold"         // Bold variant of black
	HasBackgroundBlackSoft        Color = "has-background-black-soft"         // Soft variant of black
	HasBackgroundBlackOnScheme    Color = "has-background-black-on-scheme"    // Black adapted for current scheme
	HasBackgroundBlackLightInvert Color = "has-background-black-light-invert" // Inverted light black
	HasBackgroundBlackDarkInvert  Color = "has-background-black-dark-invert"  // Inverted dark black
	HasBackgroundBlackBoldInvert  Color = "has-background-black-bold-invert"  // Inverted bold black
	HasBackgroundBlackSoftInvert  Color = "has-background-black-soft-invert"  // Inverted soft black

	// Background Color Helpers - Special Values
	HasBackgroundCurrent Color = "has-background-current" // Use current color as background
	HasBackgroundInherit Color = "has-background-inherit" // Inherit background from parent

	// Background Color Helpers - Danger Opacity Variations (00-100)
	HasBackgroundDanger00  Color = "has-background-danger-00"  // Danger 0% opacity (transparent)
	HasBackgroundDanger05  Color = "has-background-danger-05"  // Danger 5% opacity
	HasBackgroundDanger10  Color = "has-background-danger-10"  // Danger 10% opacity
	HasBackgroundDanger15  Color = "has-background-danger-15"  // Danger 15% opacity
	HasBackgroundDanger20  Color = "has-background-danger-20"  // Danger 20% opacity
	HasBackgroundDanger25  Color = "has-background-danger-25"  // Danger 25% opacity
	HasBackgroundDanger30  Color = "has-background-danger-30"  // Danger 30% opacity
	HasBackgroundDanger35  Color = "has-background-danger-35"  // Danger 35% opacity
	HasBackgroundDanger40  Color = "has-background-danger-40"  // Danger 40% opacity
	HasBackgroundDanger45  Color = "has-background-danger-45"  // Danger 45% opacity
	HasBackgroundDanger50  Color = "has-background-danger-50"  // Danger 50% opacity
	HasBackgroundDanger55  Color = "has-background-danger-55"  // Danger 55% opacity
	HasBackgroundDanger60  Color = "has-background-danger-60"  // Danger 60% opacity
	HasBackgroundDanger65  Color = "has-background-danger-65"  // Danger 65% opacity
	HasBackgroundDanger70  Color = "has-background-danger-70"  // Danger 70% opacity
	HasBackgroundDanger75  Color = "has-background-danger-75"  // Danger 75% opacity
	HasBackgroundDanger80  Color = "has-background-danger-80"  // Danger 80% opacity
	HasBackgroundDanger85  Color = "has-background-danger-85"  // Danger 85% opacity
	HasBackgroundDanger90  Color = "has-background-danger-90"  // Danger 90% opacity
	HasBackgroundDanger95  Color = "has-background-danger-95"  // Danger 95% opacity
	HasBackgroundDanger100 Color = "has-background-danger-100" // Danger 100% opacity (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid)

	// Background Color Helpers - Danger Opacity Inverts
	HasBackgroundDanger00Invert  Color = "has-background-danger-00-invert"  // Inverted contrast for danger-00
	HasBackgroundDanger05Invert  Color = "has-background-danger-05-invert"  // Inverted contrast for danger-05
	HasBackgroundDanger10Invert  Color = "has-background-danger-10-invert"  // Inverted contrast for danger-10
	HasBackgroundDanger15Invert  Color = "has-background-danger-15-invert"  // Inverted contrast for danger-15
	HasBackgroundDanger20Invert  Color = "has-background-danger-20-invert"  // Inverted contrast for danger-20
	HasBackgroundDanger25Invert  Color = "has-background-danger-25-invert"  // Inverted contrast for danger-25
	HasBackgroundDanger30Invert  Color = "has-background-danger-30-invert"  // Inverted contrast for danger-30
	HasBackgroundDanger35Invert  Color = "has-background-danger-35-invert"  // Inverted contrast for danger-35
	HasBackgroundDanger40Invert  Color = "has-background-danger-40-invert"  // Inverted contrast for danger-40
	HasBackgroundDanger45Invert  Color = "has-background-danger-45-invert"  // Inverted contrast for danger-45
	HasBackgroundDanger50Invert  Color = "has-background-danger-50-invert"  // Inverted contrast for danger-50
	HasBackgroundDanger55Invert  Color = "has-background-danger-55-invert"  // Inverted contrast for danger-55
	HasBackgroundDanger60Invert  Color = "has-background-danger-60-invert"  // Inverted contrast for danger-60
	HasBackgroundDanger65Invert  Color = "has-background-danger-65-invert"  // Inverted contrast for danger-65
	HasBackgroundDanger70Invert  Color = "has-background-danger-70-invert"  // Inverted contrast for danger-70
	HasBackgroundDanger75Invert  Color = "has-background-danger-75-invert"  // Inverted contrast for danger-75
	HasBackgroundDanger80Invert  Color = "has-background-danger-80-invert"  // Inverted contrast for danger-80
	HasBackgroundDanger85Invert  Color = "has-background-danger-85-invert"  // Inverted contrast for danger-85
	HasBackgroundDanger90Invert  Color = "has-background-danger-90-invert"  // Inverted contrast for danger-90
	HasBackgroundDanger95Invert  Color = "has-background-danger-95-invert"  // Inverted contrast for danger-95
	HasBackgroundDanger100Invert Color = "has-background-danger-100-invert" // Inverted contrast for danger-100

	// Background Color Helpers - Danger Semantic Variations
	HasBackgroundDangerInvert      Color = "has-background-danger-invert"       // Danger inverted color
	HasBackgroundDangerLight       Color = "has-background-danger-light"        // Light variant of danger
	HasBackgroundDangerDark        Color = "has-background-danger-dark"         // Dark variant of danger
	HasBackgroundDangerBold        Color = "has-background-danger-bold"         // Bold variant of danger
	HasBackgroundDangerSoft        Color = "has-background-danger-soft"         // Soft variant of danger
	HasBackgroundDangerOnScheme    Color = "has-background-danger-on-scheme"    // Danger adapted for current scheme
	HasBackgroundDangerLightInvert Color = "has-background-danger-light-invert" // Inverted light danger
	HasBackgroundDangerDarkInvert  Color = "has-background-danger-dark-invert"  // Inverted dark danger
	HasBackgroundDangerBoldInvert  Color = "has-background-danger-bold-invert"  // Inverted bold danger
	HasBackgroundDangerSoftInvert  Color = "has-background-danger-soft-invert"  // Inverted soft danger

	// Background Color Helpers - Dark Opacity Variations (00-100)
	HasBackgroundDark00  Color = "has-background-dark-00"  // Dark 0% opacity (transparent)
	HasBackgroundDark05  Color = "has-background-dark-05"  // Dark 5% opacity
	HasBackgroundDark10  Color = "has-background-dark-10"  // Dark 10% opacity
	HasBackgroundDark15  Color = "has-background-dark-15"  // Dark 15% opacity
	HasBackgroundDark20  Color = "has-background-dark-20"  // Dark 20% opacity
	HasBackgroundDark25  Color = "has-background-dark-25"  // Dark 25% opacity
	HasBackgroundDark30  Color = "has-background-dark-30"  // Dark 30% opacity
	HasBackgroundDark35  Color = "has-background-dark-35"  // Dark 35% opacity
	HasBackgroundDark40  Color = "has-background-dark-40"  // Dark 40% opacity
	HasBackgroundDark45  Color = "has-background-dark-45"  // Dark 45% opacity
	HasBackgroundDark50  Color = "has-background-dark-50"  // Dark 50% opacity
	HasBackgroundDark55  Color = "has-background-dark-55"  // Dark 55% opacity
	HasBackgroundDark60  Color = "has-background-dark-60"  // Dark 60% opacity
	HasBackgroundDark65  Color = "has-background-dark-65"  // Dark 65% opacity
	HasBackgroundDark70  Color = "has-background-dark-70"  // Dark 70% opacity
	HasBackgroundDark75  Color = "has-background-dark-75"  // Dark 75% opacity
	HasBackgroundDark80  Color = "has-background-dark-80"  // Dark 80% opacity
	HasBackgroundDark85  Color = "has-background-dark-85"  // Dark 85% opacity
	HasBackgroundDark90  Color = "has-background-dark-90"  // Dark 90% opacity
	HasBackgroundDark95  Color = "has-background-dark-95"  // Dark 95% opacity
	HasBackgroundDark100 Color = "has-background-dark-100" // Dark 100% opacity (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid)

	// Background Color Helpers - Dark Opacity Inverts
	HasBackgroundDark00Invert  Color = "has-background-dark-00-invert"  // Inverted contrast for dark-00
	HasBackgroundDark05Invert  Color = "has-background-dark-05-invert"  // Inverted contrast for dark-05
	HasBackgroundDark10Invert  Color = "has-background-dark-10-invert"  // Inverted contrast for dark-10
	HasBackgroundDark15Invert  Color = "has-background-dark-15-invert"  // Inverted contrast for dark-15
	HasBackgroundDark20Invert  Color = "has-background-dark-20-invert"  // Inverted contrast for dark-20
	HasBackgroundDark25Invert  Color = "has-background-dark-25-invert"  // Inverted contrast for dark-25
	HasBackgroundDark30Invert  Color = "has-background-dark-30-invert"  // Inverted contrast for dark-30
	HasBackgroundDark35Invert  Color = "has-background-dark-35-invert"  // Inverted contrast for dark-35
	HasBackgroundDark40Invert  Color = "has-background-dark-40-invert"  // Inverted contrast for dark-40
	HasBackgroundDark45Invert  Color = "has-background-dark-45-invert"  // Inverted contrast for dark-45
	HasBackgroundDark50Invert  Color = "has-background-dark-50-invert"  // Inverted contrast for dark-50
	HasBackgroundDark55Invert  Color = "has-background-dark-55-invert"  // Inverted contrast for dark-55
	HasBackgroundDark60Invert  Color = "has-background-dark-60-invert"  // Inverted contrast for dark-60
	HasBackgroundDark65Invert  Color = "has-background-dark-65-invert"  // Inverted contrast for dark-65
	HasBackgroundDark70Invert  Color = "has-background-dark-70-invert"  // Inverted contrast for dark-70
	HasBackgroundDark75Invert  Color = "has-background-dark-75-invert"  // Inverted contrast for dark-75
	HasBackgroundDark80Invert  Color = "has-background-dark-80-invert"  // Inverted contrast for dark-80
	HasBackgroundDark85Invert  Color = "has-background-dark-85-invert"  // Inverted contrast for dark-85
	HasBackgroundDark90Invert  Color = "has-background-dark-90-invert"  // Inverted contrast for dark-90
	HasBackgroundDark95Invert  Color = "has-background-dark-95-invert"  // Inverted contrast for dark-95
	HasBackgroundDark100Invert Color = "has-background-dark-100-invert" // Inverted contrast for dark-100

	// Background Color Helpers - Dark Semantic Variations
	HasBackgroundDarkInvert      Color = "has-background-dark-invert"       // Dark inverted color
	HasBackgroundDarkLight       Color = "has-background-dark-light"        // Light variant of dark
	HasBackgroundDarkDark        Color = "has-background-dark-dark"         // Dark variant of dark
	HasBackgroundDarkBold        Color = "has-background-dark-bold"         // Bold variant of dark
	HasBackgroundDarkSoft        Color = "has-background-dark-soft"         // Soft variant of dark
	HasBackgroundDarkOnScheme    Color = "has-background-dark-on-scheme"    // Dark adapted for current scheme
	HasBackgroundDarkLightInvert Color = "has-background-dark-light-invert" // Inverted light dark
	HasBackgroundDarkDarkInvert  Color = "has-background-dark-dark-invert"  // Inverted dark dark
	HasBackgroundDarkBoldInvert  Color = "has-background-dark-bold-invert"  // Inverted bold dark
	HasBackgroundDarkSoftInvert  Color = "has-background-dark-soft-invert"  // Inverted soft dark

	// Background Color Helpers - Monochrome Grey Variations
	HasBackgroundGrey        Color = "has-background-grey"         // Standard grey background
	HasBackgroundGreyDark    Color = "has-background-grey-dark"    // Dark grey background
	HasBackgroundGreyDarker  Color = "has-background-grey-darker"  // Darker grey background
	HasBackgroundGreyLight   Color = "has-background-grey-light"   // Light grey background
	HasBackgroundGreyLighter Color = "has-background-grey-lighter" // Lighter grey background

	// Background Color Helpers - Info Opacity Variations (00-100)
	HasBackgroundInfo00  Color = "has-background-info-00"  // Info 00% opacity (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent)
	HasBackgroundInfo05  Color = "has-background-info-05"  // Info 05% opacity
	HasBackgroundInfo10  Color = "has-background-info-10"  // Info 10% opacity
	HasBackgroundInfo15  Color = "has-background-info-15"  // Info 15% opacity
	HasBackgroundInfo20  Color = "has-background-info-20"  // Info 20% opacity
	HasBackgroundInfo25  Color = "has-background-info-25"  // Info 25% opacity
	HasBackgroundInfo30  Color = "has-background-info-30"  // Info 30% opacity
	HasBackgroundInfo35  Color = "has-background-info-35"  // Info 35% opacity
	HasBackgroundInfo40  Color = "has-background-info-40"  // Info 40% opacity
	HasBackgroundInfo45  Color = "has-background-info-45"  // Info 45% opacity
	HasBackgroundInfo50  Color = "has-background-info-50"  // Info 50% opacity
	HasBackgroundInfo55  Color = "has-background-info-55"  // Info 55% opacity
	HasBackgroundInfo60  Color = "has-background-info-60"  // Info 60% opacity
	HasBackgroundInfo65  Color = "has-background-info-65"  // Info 65% opacity
	HasBackgroundInfo70  Color = "has-background-info-70"  // Info 70% opacity
	HasBackgroundInfo75  Color = "has-background-info-75"  // Info 75% opacity
	HasBackgroundInfo80  Color = "has-background-info-80"  // Info 80% opacity
	HasBackgroundInfo85  Color = "has-background-info-85"  // Info 85% opacity
	HasBackgroundInfo90  Color = "has-background-info-90"  // Info 90% opacity
	HasBackgroundInfo95  Color = "has-background-info-95"  // Info 95% opacity
	HasBackgroundInfo100 Color = "has-background-info-100" // Info 100% opacity (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid)

	// Background Color Helpers - Info Opacity Inverts
	HasBackgroundInfo00Invert  Color = "has-background-info-00-invert"  // Inverted contrast for info-00
	HasBackgroundInfo05Invert  Color = "has-background-info-05-invert"  // Inverted contrast for info-05
	HasBackgroundInfo10Invert  Color = "has-background-info-10-invert"  // Inverted contrast for info-10
	HasBackgroundInfo15Invert  Color = "has-background-info-15-invert"  // Inverted contrast for info-15
	HasBackgroundInfo20Invert  Color = "has-background-info-20-invert"  // Inverted contrast for info-20
	HasBackgroundInfo25Invert  Color = "has-background-info-25-invert"  // Inverted contrast for info-25
	HasBackgroundInfo30Invert  Color = "has-background-info-30-invert"  // Inverted contrast for info-30
	HasBackgroundInfo35Invert  Color = "has-background-info-35-invert"  // Inverted contrast for info-35
	HasBackgroundInfo40Invert  Color = "has-background-info-40-invert"  // Inverted contrast for info-40
	HasBackgroundInfo45Invert  Color = "has-background-info-45-invert"  // Inverted contrast for info-45
	HasBackgroundInfo50Invert  Color = "has-background-info-50-invert"  // Inverted contrast for info-50
	HasBackgroundInfo55Invert  Color = "has-background-info-55-invert"  // Inverted contrast for info-55
	HasBackgroundInfo60Invert  Color = "has-background-info-60-invert"  // Inverted contrast for info-60
	HasBackgroundInfo65Invert  Color = "has-background-info-65-invert"  // Inverted contrast for info-65
	HasBackgroundInfo70Invert  Color = "has-background-info-70-invert"  // Inverted contrast for info-70
	HasBackgroundInfo75Invert  Color = "has-background-info-75-invert"  // Inverted contrast for info-75
	HasBackgroundInfo80Invert  Color = "has-background-info-80-invert"  // Inverted contrast for info-80
	HasBackgroundInfo85Invert  Color = "has-background-info-85-invert"  // Inverted contrast for info-85
	HasBackgroundInfo90Invert  Color = "has-background-info-90-invert"  // Inverted contrast for info-90
	HasBackgroundInfo95Invert  Color = "has-background-info-95-invert"  // Inverted contrast for info-95
	HasBackgroundInfo100Invert Color = "has-background-info-100-invert" // Inverted contrast for info-100

	// Background Color Helpers - Info Semantic Variations
	HasBackgroundInfoInvert      Color = "has-background-info-invert"       // Info inverted color
	HasBackgroundInfoLight       Color = "has-background-info-light"        // Light variant of info
	HasBackgroundInfoDark        Color = "has-background-info-dark"         // Dark variant of info
	HasBackgroundInfoBold        Color = "has-background-info-bold"         // Bold variant of info
	HasBackgroundInfoSoft        Color = "has-background-info-soft"         // Soft variant of info
	HasBackgroundInfoOnScheme    Color = "has-background-info-on-scheme"    // Info adapted for current scheme
	HasBackgroundInfoLightInvert Color = "has-background-info-light-invert" // Inverted light info
	HasBackgroundInfoDarkInvert  Color = "has-background-info-dark-invert"  // Inverted dark info
	HasBackgroundInfoBoldInvert  Color = "has-background-info-bold-invert"  // Inverted bold info
	HasBackgroundInfoSoftInvert  Color = "has-background-info-soft-invert"  // Inverted soft info

	// Background Color Helpers - Light Opacity Variations (00-100)
	HasBackgroundLight00  Color = "has-background-light-00"  // Light 00% opacity (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent)
	HasBackgroundLight05  Color = "has-background-light-05"  // Light 05% opacity
	HasBackgroundLight10  Color = "has-background-light-10"  // Light 10% opacity
	HasBackgroundLight15  Color = "has-background-light-15"  // Light 15% opacity
	HasBackgroundLight20  Color = "has-background-light-20"  // Light 20% opacity
	HasBackgroundLight25  Color = "has-background-light-25"  // Light 25% opacity
	HasBackgroundLight30  Color = "has-background-light-30"  // Light 30% opacity
	HasBackgroundLight35  Color = "has-background-light-35"  // Light 35% opacity
	HasBackgroundLight40  Color = "has-background-light-40"  // Light 40% opacity
	HasBackgroundLight45  Color = "has-background-light-45"  // Light 45% opacity
	HasBackgroundLight50  Color = "has-background-light-50"  // Light 50% opacity
	HasBackgroundLight55  Color = "has-background-light-55"  // Light 55% opacity
	HasBackgroundLight60  Color = "has-background-light-60"  // Light 60% opacity
	HasBackgroundLight65  Color = "has-background-light-65"  // Light 65% opacity
	HasBackgroundLight70  Color = "has-background-light-70"  // Light 70% opacity
	HasBackgroundLight75  Color = "has-background-light-75"  // Light 75% opacity
	HasBackgroundLight80  Color = "has-background-light-80"  // Light 80% opacity
	HasBackgroundLight85  Color = "has-background-light-85"  // Light 85% opacity
	HasBackgroundLight90  Color = "has-background-light-90"  // Light 90% opacity
	HasBackgroundLight95  Color = "has-background-light-95"  // Light 95% opacity
	HasBackgroundLight100 Color = "has-background-light-100" // Light 100% opacity (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid)

	// Background Color Helpers - Light Opacity Inverts
	HasBackgroundLight00Invert  Color = "has-background-light-00-invert"  // Inverted contrast for light-00
	HasBackgroundLight05Invert  Color = "has-background-light-05-invert"  // Inverted contrast for light-05
	HasBackgroundLight10Invert  Color = "has-background-light-10-invert"  // Inverted contrast for light-10
	HasBackgroundLight15Invert  Color = "has-background-light-15-invert"  // Inverted contrast for light-15
	HasBackgroundLight20Invert  Color = "has-background-light-20-invert"  // Inverted contrast for light-20
	HasBackgroundLight25Invert  Color = "has-background-light-25-invert"  // Inverted contrast for light-25
	HasBackgroundLight30Invert  Color = "has-background-light-30-invert"  // Inverted contrast for light-30
	HasBackgroundLight35Invert  Color = "has-background-light-35-invert"  // Inverted contrast for light-35
	HasBackgroundLight40Invert  Color = "has-background-light-40-invert"  // Inverted contrast for light-40
	HasBackgroundLight45Invert  Color = "has-background-light-45-invert"  // Inverted contrast for light-45
	HasBackgroundLight50Invert  Color = "has-background-light-50-invert"  // Inverted contrast for light-50
	HasBackgroundLight55Invert  Color = "has-background-light-55-invert"  // Inverted contrast for light-55
	HasBackgroundLight60Invert  Color = "has-background-light-60-invert"  // Inverted contrast for light-60
	HasBackgroundLight65Invert  Color = "has-background-light-65-invert"  // Inverted contrast for light-65
	HasBackgroundLight70Invert  Color = "has-background-light-70-invert"  // Inverted contrast for light-70
	HasBackgroundLight75Invert  Color = "has-background-light-75-invert"  // Inverted contrast for light-75
	HasBackgroundLight80Invert  Color = "has-background-light-80-invert"  // Inverted contrast for light-80
	HasBackgroundLight85Invert  Color = "has-background-light-85-invert"  // Inverted contrast for light-85
	HasBackgroundLight90Invert  Color = "has-background-light-90-invert"  // Inverted contrast for light-90
	HasBackgroundLight95Invert  Color = "has-background-light-95-invert"  // Inverted contrast for light-95
	HasBackgroundLight100Invert Color = "has-background-light-100-invert" // Inverted contrast for light-100

	// Background Color Helpers - Light Semantic Variations
	HasBackgroundLightInvert      Color = "has-background-light-invert"       // Light inverted color
	HasBackgroundLightLight       Color = "has-background-light-light"        // Light variant of light
	HasBackgroundLightDark        Color = "has-background-light-dark"         // Dark variant of light
	HasBackgroundLightBold        Color = "has-background-light-bold"         // Bold variant of light
	HasBackgroundLightSoft        Color = "has-background-light-soft"         // Soft variant of light
	HasBackgroundLightOnScheme    Color = "has-background-light-on-scheme"    // Light adapted for current scheme
	HasBackgroundLightLightInvert Color = "has-background-light-light-invert" // Inverted light light
	HasBackgroundLightDarkInvert  Color = "has-background-light-dark-invert"  // Inverted dark light
	HasBackgroundLightBoldInvert  Color = "has-background-light-bold-invert"  // Inverted bold light
	HasBackgroundLightSoftInvert  Color = "has-background-light-soft-invert"  // Inverted soft light

	// Background Color Helpers - Link Opacity Variations (00-100)
	HasBackgroundLink00  Color = "has-background-link-00"  // Link 00% opacity (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent)
	HasBackgroundLink05  Color = "has-background-link-05"  // Link 05% opacity
	HasBackgroundLink10  Color = "has-background-link-10"  // Link 10% opacity
	HasBackgroundLink15  Color = "has-background-link-15"  // Link 15% opacity
	HasBackgroundLink20  Color = "has-background-link-20"  // Link 20% opacity
	HasBackgroundLink25  Color = "has-background-link-25"  // Link 25% opacity
	HasBackgroundLink30  Color = "has-background-link-30"  // Link 30% opacity
	HasBackgroundLink35  Color = "has-background-link-35"  // Link 35% opacity
	HasBackgroundLink40  Color = "has-background-link-40"  // Link 40% opacity
	HasBackgroundLink45  Color = "has-background-link-45"  // Link 45% opacity
	HasBackgroundLink50  Color = "has-background-link-50"  // Link 50% opacity
	HasBackgroundLink55  Color = "has-background-link-55"  // Link 55% opacity
	HasBackgroundLink60  Color = "has-background-link-60"  // Link 60% opacity
	HasBackgroundLink65  Color = "has-background-link-65"  // Link 65% opacity
	HasBackgroundLink70  Color = "has-background-link-70"  // Link 70% opacity
	HasBackgroundLink75  Color = "has-background-link-75"  // Link 75% opacity
	HasBackgroundLink80  Color = "has-background-link-80"  // Link 80% opacity
	HasBackgroundLink85  Color = "has-background-link-85"  // Link 85% opacity
	HasBackgroundLink90  Color = "has-background-link-90"  // Link 90% opacity
	HasBackgroundLink95  Color = "has-background-link-95"  // Link 95% opacity
	HasBackgroundLink100 Color = "has-background-link-100" // Link 100% opacity (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid)

	// Background Color Helpers - Link Opacity Inverts
	HasBackgroundLink00Invert  Color = "has-background-link-00-invert"  // Inverted contrast for link-00
	HasBackgroundLink05Invert  Color = "has-background-link-05-invert"  // Inverted contrast for link-05
	HasBackgroundLink10Invert  Color = "has-background-link-10-invert"  // Inverted contrast for link-10
	HasBackgroundLink15Invert  Color = "has-background-link-15-invert"  // Inverted contrast for link-15
	HasBackgroundLink20Invert  Color = "has-background-link-20-invert"  // Inverted contrast for link-20
	HasBackgroundLink25Invert  Color = "has-background-link-25-invert"  // Inverted contrast for link-25
	HasBackgroundLink30Invert  Color = "has-background-link-30-invert"  // Inverted contrast for link-30
	HasBackgroundLink35Invert  Color = "has-background-link-35-invert"  // Inverted contrast for link-35
	HasBackgroundLink40Invert  Color = "has-background-link-40-invert"  // Inverted contrast for link-40
	HasBackgroundLink45Invert  Color = "has-background-link-45-invert"  // Inverted contrast for link-45
	HasBackgroundLink50Invert  Color = "has-background-link-50-invert"  // Inverted contrast for link-50
	HasBackgroundLink55Invert  Color = "has-background-link-55-invert"  // Inverted contrast for link-55
	HasBackgroundLink60Invert  Color = "has-background-link-60-invert"  // Inverted contrast for link-60
	HasBackgroundLink65Invert  Color = "has-background-link-65-invert"  // Inverted contrast for link-65
	HasBackgroundLink70Invert  Color = "has-background-link-70-invert"  // Inverted contrast for link-70
	HasBackgroundLink75Invert  Color = "has-background-link-75-invert"  // Inverted contrast for link-75
	HasBackgroundLink80Invert  Color = "has-background-link-80-invert"  // Inverted contrast for link-80
	HasBackgroundLink85Invert  Color = "has-background-link-85-invert"  // Inverted contrast for link-85
	HasBackgroundLink90Invert  Color = "has-background-link-90-invert"  // Inverted contrast for link-90
	HasBackgroundLink95Invert  Color = "has-background-link-95-invert"  // Inverted contrast for link-95
	HasBackgroundLink100Invert Color = "has-background-link-100-invert" // Inverted contrast for link-100

	// Background Color Helpers - Link Semantic Variations
	HasBackgroundLinkInvert      Color = "has-background-link-invert"       // Link inverted color
	HasBackgroundLinkLight       Color = "has-background-link-light"        // Light variant of link
	HasBackgroundLinkDark        Color = "has-background-link-dark"         // Dark variant of link
	HasBackgroundLinkBold        Color = "has-background-link-bold"         // Bold variant of link
	HasBackgroundLinkSoft        Color = "has-background-link-soft"         // Soft variant of link
	HasBackgroundLinkOnScheme    Color = "has-background-link-on-scheme"    // Link adapted for current scheme
	HasBackgroundLinkLightInvert Color = "has-background-link-light-invert" // Inverted light link
	HasBackgroundLinkDarkInvert  Color = "has-background-link-dark-invert"  // Inverted dark link
	HasBackgroundLinkBoldInvert  Color = "has-background-link-bold-invert"  // Inverted bold link
	HasBackgroundLinkSoftInvert  Color = "has-background-link-soft-invert"  // Inverted soft link

	// Background Color Helpers - Primary Opacity Variations (00-100)
	HasBackgroundPrimary00  Color = "has-background-primary-00"  // Primary 00% opacity (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent)
	HasBackgroundPrimary05  Color = "has-background-primary-05"  // Primary 05% opacity
	HasBackgroundPrimary10  Color = "has-background-primary-10"  // Primary 10% opacity
	HasBackgroundPrimary15  Color = "has-background-primary-15"  // Primary 15% opacity
	HasBackgroundPrimary20  Color = "has-background-primary-20"  // Primary 20% opacity
	HasBackgroundPrimary25  Color = "has-background-primary-25"  // Primary 25% opacity
	HasBackgroundPrimary30  Color = "has-background-primary-30"  // Primary 30% opacity
	HasBackgroundPrimary35  Color = "has-background-primary-35"  // Primary 35% opacity
	HasBackgroundPrimary40  Color = "has-background-primary-40"  // Primary 40% opacity
	HasBackgroundPrimary45  Color = "has-background-primary-45"  // Primary 45% opacity
	HasBackgroundPrimary50  Color = "has-background-primary-50"  // Primary 50% opacity
	HasBackgroundPrimary55  Color = "has-background-primary-55"  // Primary 55% opacity
	HasBackgroundPrimary60  Color = "has-background-primary-60"  // Primary 60% opacity
	HasBackgroundPrimary65  Color = "has-background-primary-65"  // Primary 65% opacity
	HasBackgroundPrimary70  Color = "has-background-primary-70"  // Primary 70% opacity
	HasBackgroundPrimary75  Color = "has-background-primary-75"  // Primary 75% opacity
	HasBackgroundPrimary80  Color = "has-background-primary-80"  // Primary 80% opacity
	HasBackgroundPrimary85  Color = "has-background-primary-85"  // Primary 85% opacity
	HasBackgroundPrimary90  Color = "has-background-primary-90"  // Primary 90% opacity
	HasBackgroundPrimary95  Color = "has-background-primary-95"  // Primary 95% opacity
	HasBackgroundPrimary100 Color = "has-background-primary-100" // Primary 100% opacity (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid)

	// Background Color Helpers - Primary Opacity Inverts
	HasBackgroundPrimary00Invert  Color = "has-background-primary-00-invert"  // Inverted contrast for primary-00
	HasBackgroundPrimary05Invert  Color = "has-background-primary-05-invert"  // Inverted contrast for primary-05
	HasBackgroundPrimary10Invert  Color = "has-background-primary-10-invert"  // Inverted contrast for primary-10
	HasBackgroundPrimary15Invert  Color = "has-background-primary-15-invert"  // Inverted contrast for primary-15
	HasBackgroundPrimary20Invert  Color = "has-background-primary-20-invert"  // Inverted contrast for primary-20
	HasBackgroundPrimary25Invert  Color = "has-background-primary-25-invert"  // Inverted contrast for primary-25
	HasBackgroundPrimary30Invert  Color = "has-background-primary-30-invert"  // Inverted contrast for primary-30
	HasBackgroundPrimary35Invert  Color = "has-background-primary-35-invert"  // Inverted contrast for primary-35
	HasBackgroundPrimary40Invert  Color = "has-background-primary-40-invert"  // Inverted contrast for primary-40
	HasBackgroundPrimary45Invert  Color = "has-background-primary-45-invert"  // Inverted contrast for primary-45
	HasBackgroundPrimary50Invert  Color = "has-background-primary-50-invert"  // Inverted contrast for primary-50
	HasBackgroundPrimary55Invert  Color = "has-background-primary-55-invert"  // Inverted contrast for primary-55
	HasBackgroundPrimary60Invert  Color = "has-background-primary-60-invert"  // Inverted contrast for primary-60
	HasBackgroundPrimary65Invert  Color = "has-background-primary-65-invert"  // Inverted contrast for primary-65
	HasBackgroundPrimary70Invert  Color = "has-background-primary-70-invert"  // Inverted contrast for primary-70
	HasBackgroundPrimary75Invert  Color = "has-background-primary-75-invert"  // Inverted contrast for primary-75
	HasBackgroundPrimary80Invert  Color = "has-background-primary-80-invert"  // Inverted contrast for primary-80
	HasBackgroundPrimary85Invert  Color = "has-background-primary-85-invert"  // Inverted contrast for primary-85
	HasBackgroundPrimary90Invert  Color = "has-background-primary-90-invert"  // Inverted contrast for primary-90
	HasBackgroundPrimary95Invert  Color = "has-background-primary-95-invert"  // Inverted contrast for primary-95
	HasBackgroundPrimary100Invert Color = "has-background-primary-100-invert" // Inverted contrast for primary-100

	// Background Color Helpers - Primary Semantic Variations
	HasBackgroundPrimaryInvert      Color = "has-background-primary-invert"       // Primary inverted color
	HasBackgroundPrimaryLight       Color = "has-background-primary-light"        // Light variant of primary
	HasBackgroundPrimaryDark        Color = "has-background-primary-dark"         // Dark variant of primary
	HasBackgroundPrimaryBold        Color = "has-background-primary-bold"         // Bold variant of primary
	HasBackgroundPrimarySoft        Color = "has-background-primary-soft"         // Soft variant of primary
	HasBackgroundPrimaryOnScheme    Color = "has-background-primary-on-scheme"    // Primary adapted for current scheme
	HasBackgroundPrimaryLightInvert Color = "has-background-primary-light-invert" // Inverted light primary
	HasBackgroundPrimaryDarkInvert  Color = "has-background-primary-dark-invert"  // Inverted dark primary
	HasBackgroundPrimaryBoldInvert  Color = "has-background-primary-bold-invert"  // Inverted bold primary
	HasBackgroundPrimarySoftInvert  Color = "has-background-primary-soft-invert"  // Inverted soft primary

	// Background Color Helpers - Success Opacity Variations (00-100)
	HasBackgroundSuccess00  Color = "has-background-success-00"  // Success 00% opacity (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent)
	HasBackgroundSuccess05  Color = "has-background-success-05"  // Success 05% opacity
	HasBackgroundSuccess10  Color = "has-background-success-10"  // Success 10% opacity
	HasBackgroundSuccess15  Color = "has-background-success-15"  // Success 15% opacity
	HasBackgroundSuccess20  Color = "has-background-success-20"  // Success 20% opacity
	HasBackgroundSuccess25  Color = "has-background-success-25"  // Success 25% opacity
	HasBackgroundSuccess30  Color = "has-background-success-30"  // Success 30% opacity
	HasBackgroundSuccess35  Color = "has-background-success-35"  // Success 35% opacity
	HasBackgroundSuccess40  Color = "has-background-success-40"  // Success 40% opacity
	HasBackgroundSuccess45  Color = "has-background-success-45"  // Success 45% opacity
	HasBackgroundSuccess50  Color = "has-background-success-50"  // Success 50% opacity
	HasBackgroundSuccess55  Color = "has-background-success-55"  // Success 55% opacity
	HasBackgroundSuccess60  Color = "has-background-success-60"  // Success 60% opacity
	HasBackgroundSuccess65  Color = "has-background-success-65"  // Success 65% opacity
	HasBackgroundSuccess70  Color = "has-background-success-70"  // Success 70% opacity
	HasBackgroundSuccess75  Color = "has-background-success-75"  // Success 75% opacity
	HasBackgroundSuccess80  Color = "has-background-success-80"  // Success 80% opacity
	HasBackgroundSuccess85  Color = "has-background-success-85"  // Success 85% opacity
	HasBackgroundSuccess90  Color = "has-background-success-90"  // Success 90% opacity
	HasBackgroundSuccess95  Color = "has-background-success-95"  // Success 95% opacity
	HasBackgroundSuccess100 Color = "has-background-success-100" // Success 100% opacity (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid)

	// Background Color Helpers - Success Opacity Inverts
	HasBackgroundSuccess00Invert  Color = "has-background-success-00-invert"  // Inverted contrast for success-00
	HasBackgroundSuccess05Invert  Color = "has-background-success-05-invert"  // Inverted contrast for success-05
	HasBackgroundSuccess10Invert  Color = "has-background-success-10-invert"  // Inverted contrast for success-10
	HasBackgroundSuccess15Invert  Color = "has-background-success-15-invert"  // Inverted contrast for success-15
	HasBackgroundSuccess20Invert  Color = "has-background-success-20-invert"  // Inverted contrast for success-20
	HasBackgroundSuccess25Invert  Color = "has-background-success-25-invert"  // Inverted contrast for success-25
	HasBackgroundSuccess30Invert  Color = "has-background-success-30-invert"  // Inverted contrast for success-30
	HasBackgroundSuccess35Invert  Color = "has-background-success-35-invert"  // Inverted contrast for success-35
	HasBackgroundSuccess40Invert  Color = "has-background-success-40-invert"  // Inverted contrast for success-40
	HasBackgroundSuccess45Invert  Color = "has-background-success-45-invert"  // Inverted contrast for success-45
	HasBackgroundSuccess50Invert  Color = "has-background-success-50-invert"  // Inverted contrast for success-50
	HasBackgroundSuccess55Invert  Color = "has-background-success-55-invert"  // Inverted contrast for success-55
	HasBackgroundSuccess60Invert  Color = "has-background-success-60-invert"  // Inverted contrast for success-60
	HasBackgroundSuccess65Invert  Color = "has-background-success-65-invert"  // Inverted contrast for success-65
	HasBackgroundSuccess70Invert  Color = "has-background-success-70-invert"  // Inverted contrast for success-70
	HasBackgroundSuccess75Invert  Color = "has-background-success-75-invert"  // Inverted contrast for success-75
	HasBackgroundSuccess80Invert  Color = "has-background-success-80-invert"  // Inverted contrast for success-80
	HasBackgroundSuccess85Invert  Color = "has-background-success-85-invert"  // Inverted contrast for success-85
	HasBackgroundSuccess90Invert  Color = "has-background-success-90-invert"  // Inverted contrast for success-90
	HasBackgroundSuccess95Invert  Color = "has-background-success-95-invert"  // Inverted contrast for success-95
	HasBackgroundSuccess100Invert Color = "has-background-success-100-invert" // Inverted contrast for success-100

	// Background Color Helpers - Success Semantic Variations
	HasBackgroundSuccessInvert      Color = "has-background-success-invert"       // Success inverted color
	HasBackgroundSuccessLight       Color = "has-background-success-light"        // Light variant of success
	HasBackgroundSuccessDark        Color = "has-background-success-dark"         // Dark variant of success
	HasBackgroundSuccessBold        Color = "has-background-success-bold"         // Bold variant of success
	HasBackgroundSuccessSoft        Color = "has-background-success-soft"         // Soft variant of success
	HasBackgroundSuccessOnScheme    Color = "has-background-success-on-scheme"    // Success adapted for current scheme
	HasBackgroundSuccessLightInvert Color = "has-background-success-light-invert" // Inverted light success
	HasBackgroundSuccessDarkInvert  Color = "has-background-success-dark-invert"  // Inverted dark success
	HasBackgroundSuccessBoldInvert  Color = "has-background-success-bold-invert"  // Inverted bold success
	HasBackgroundSuccessSoftInvert  Color = "has-background-success-soft-invert"  // Inverted soft success

	// Background Color Helpers - Text Opacity Variations (00-100)
	HasBackgroundText00  Color = "has-background-text-00"  // Text 00% opacity (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent)
	HasBackgroundText05  Color = "has-background-text-05"  // Text 05% opacity
	HasBackgroundText10  Color = "has-background-text-10"  // Text 10% opacity
	HasBackgroundText15  Color = "has-background-text-15"  // Text 15% opacity
	HasBackgroundText20  Color = "has-background-text-20"  // Text 20% opacity
	HasBackgroundText25  Color = "has-background-text-25"  // Text 25% opacity
	HasBackgroundText30  Color = "has-background-text-30"  // Text 30% opacity
	HasBackgroundText35  Color = "has-background-text-35"  // Text 35% opacity
	HasBackgroundText40  Color = "has-background-text-40"  // Text 40% opacity
	HasBackgroundText45  Color = "has-background-text-45"  // Text 45% opacity
	HasBackgroundText50  Color = "has-background-text-50"  // Text 50% opacity
	HasBackgroundText55  Color = "has-background-text-55"  // Text 55% opacity
	HasBackgroundText60  Color = "has-background-text-60"  // Text 60% opacity
	HasBackgroundText65  Color = "has-background-text-65"  // Text 65% opacity
	HasBackgroundText70  Color = "has-background-text-70"  // Text 70% opacity
	HasBackgroundText75  Color = "has-background-text-75"  // Text 75% opacity
	HasBackgroundText80  Color = "has-background-text-80"  // Text 80% opacity
	HasBackgroundText85  Color = "has-background-text-85"  // Text 85% opacity
	HasBackgroundText90  Color = "has-background-text-90"  // Text 90% opacity
	HasBackgroundText95  Color = "has-background-text-95"  // Text 95% opacity
	HasBackgroundText100 Color = "has-background-text-100" // Text 100% opacity (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid)

	// Background Color Helpers - Text Opacity Inverts
	HasBackgroundText00Invert  Color = "has-background-text-00-invert"  // Inverted contrast for text-00
	HasBackgroundText05Invert  Color = "has-background-text-05-invert"  // Inverted contrast for text-05
	HasBackgroundText10Invert  Color = "has-background-text-10-invert"  // Inverted contrast for text-10
	HasBackgroundText15Invert  Color = "has-background-text-15-invert"  // Inverted contrast for text-15
	HasBackgroundText20Invert  Color = "has-background-text-20-invert"  // Inverted contrast for text-20
	HasBackgroundText25Invert  Color = "has-background-text-25-invert"  // Inverted contrast for text-25
	HasBackgroundText30Invert  Color = "has-background-text-30-invert"  // Inverted contrast for text-30
	HasBackgroundText35Invert  Color = "has-background-text-35-invert"  // Inverted contrast for text-35
	HasBackgroundText40Invert  Color = "has-background-text-40-invert"  // Inverted contrast for text-40
	HasBackgroundText45Invert  Color = "has-background-text-45-invert"  // Inverted contrast for text-45
	HasBackgroundText50Invert  Color = "has-background-text-50-invert"  // Inverted contrast for text-50
	HasBackgroundText55Invert  Color = "has-background-text-55-invert"  // Inverted contrast for text-55
	HasBackgroundText60Invert  Color = "has-background-text-60-invert"  // Inverted contrast for text-60
	HasBackgroundText65Invert  Color = "has-background-text-65-invert"  // Inverted contrast for text-65
	HasBackgroundText70Invert  Color = "has-background-text-70-invert"  // Inverted contrast for text-70
	HasBackgroundText75Invert  Color = "has-background-text-75-invert"  // Inverted contrast for text-75
	HasBackgroundText80Invert  Color = "has-background-text-80-invert"  // Inverted contrast for text-80
	HasBackgroundText85Invert  Color = "has-background-text-85-invert"  // Inverted contrast for text-85
	HasBackgroundText90Invert  Color = "has-background-text-90-invert"  // Inverted contrast for text-90
	HasBackgroundText95Invert  Color = "has-background-text-95-invert"  // Inverted contrast for text-95
	HasBackgroundText100Invert Color = "has-background-text-100-invert" // Inverted contrast for text-100

	// Background Color Helpers - Text Semantic Variations
	HasBackgroundTextInvert      Color = "has-background-text-invert"       // Text inverted color
	HasBackgroundTextLight       Color = "has-background-text-light"        // Light variant of text
	HasBackgroundTextDark        Color = "has-background-text-dark"         // Dark variant of text
	HasBackgroundTextBold        Color = "has-background-text-bold"         // Bold variant of text
	HasBackgroundTextSoft        Color = "has-background-text-soft"         // Soft variant of text
	HasBackgroundTextOnScheme    Color = "has-background-text-on-scheme"    // Text adapted for current scheme
	HasBackgroundTextLightInvert Color = "has-background-text-light-invert" // Inverted light text
	HasBackgroundTextDarkInvert  Color = "has-background-text-dark-invert"  // Inverted dark text
	HasBackgroundTextBoldInvert  Color = "has-background-text-bold-invert"  // Inverted bold text
	HasBackgroundTextSoftInvert  Color = "has-background-text-soft-invert"  // Inverted soft text

	// Background Color Helpers - Warning Opacity Variations (00-100)
	HasBackgroundWarning00  Color = "has-background-warning-00"  // Warning 00% opacity (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent)
	HasBackgroundWarning05  Color = "has-background-warning-05"  // Warning 05% opacity
	HasBackgroundWarning10  Color = "has-background-warning-10"  // Warning 10% opacity
	HasBackgroundWarning15  Color = "has-background-warning-15"  // Warning 15% opacity
	HasBackgroundWarning20  Color = "has-background-warning-20"  // Warning 20% opacity
	HasBackgroundWarning25  Color = "has-background-warning-25"  // Warning 25% opacity
	HasBackgroundWarning30  Color = "has-background-warning-30"  // Warning 30% opacity
	HasBackgroundWarning35  Color = "has-background-warning-35"  // Warning 35% opacity
	HasBackgroundWarning40  Color = "has-background-warning-40"  // Warning 40% opacity
	HasBackgroundWarning45  Color = "has-background-warning-45"  // Warning 45% opacity
	HasBackgroundWarning50  Color = "has-background-warning-50"  // Warning 50% opacity
	HasBackgroundWarning55  Color = "has-background-warning-55"  // Warning 55% opacity
	HasBackgroundWarning60  Color = "has-background-warning-60"  // Warning 60% opacity
	HasBackgroundWarning65  Color = "has-background-warning-65"  // Warning 65% opacity
	HasBackgroundWarning70  Color = "has-background-warning-70"  // Warning 70% opacity
	HasBackgroundWarning75  Color = "has-background-warning-75"  // Warning 75% opacity
	HasBackgroundWarning80  Color = "has-background-warning-80"  // Warning 80% opacity
	HasBackgroundWarning85  Color = "has-background-warning-85"  // Warning 85% opacity
	HasBackgroundWarning90  Color = "has-background-warning-90"  // Warning 90% opacity
	HasBackgroundWarning95  Color = "has-background-warning-95"  // Warning 95% opacity
	HasBackgroundWarning100 Color = "has-background-warning-100" // Warning 100% opacity (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid)

	// Background Color Helpers - Warning Opacity Inverts
	HasBackgroundWarning00Invert  Color = "has-background-warning-00-invert"  // Inverted contrast for warning-00
	HasBackgroundWarning05Invert  Color = "has-background-warning-05-invert"  // Inverted contrast for warning-05
	HasBackgroundWarning10Invert  Color = "has-background-warning-10-invert"  // Inverted contrast for warning-10
	HasBackgroundWarning15Invert  Color = "has-background-warning-15-invert"  // Inverted contrast for warning-15
	HasBackgroundWarning20Invert  Color = "has-background-warning-20-invert"  // Inverted contrast for warning-20
	HasBackgroundWarning25Invert  Color = "has-background-warning-25-invert"  // Inverted contrast for warning-25
	HasBackgroundWarning30Invert  Color = "has-background-warning-30-invert"  // Inverted contrast for warning-30
	HasBackgroundWarning35Invert  Color = "has-background-warning-35-invert"  // Inverted contrast for warning-35
	HasBackgroundWarning40Invert  Color = "has-background-warning-40-invert"  // Inverted contrast for warning-40
	HasBackgroundWarning45Invert  Color = "has-background-warning-45-invert"  // Inverted contrast for warning-45
	HasBackgroundWarning50Invert  Color = "has-background-warning-50-invert"  // Inverted contrast for warning-50
	HasBackgroundWarning55Invert  Color = "has-background-warning-55-invert"  // Inverted contrast for warning-55
	HasBackgroundWarning60Invert  Color = "has-background-warning-60-invert"  // Inverted contrast for warning-60
	HasBackgroundWarning65Invert  Color = "has-background-warning-65-invert"  // Inverted contrast for warning-65
	HasBackgroundWarning70Invert  Color = "has-background-warning-70-invert"  // Inverted contrast for warning-70
	HasBackgroundWarning75Invert  Color = "has-background-warning-75-invert"  // Inverted contrast for warning-75
	HasBackgroundWarning80Invert  Color = "has-background-warning-80-invert"  // Inverted contrast for warning-80
	HasBackgroundWarning85Invert  Color = "has-background-warning-85-invert"  // Inverted contrast for warning-85
	HasBackgroundWarning90Invert  Color = "has-background-warning-90-invert"  // Inverted contrast for warning-90
	HasBackgroundWarning95Invert  Color = "has-background-warning-95-invert"  // Inverted contrast for warning-95
	HasBackgroundWarning100Invert Color = "has-background-warning-100-invert" // Inverted contrast for warning-100

	// Background Color Helpers - Warning Semantic Variations
	HasBackgroundWarningInvert      Color = "has-background-warning-invert"       // Warning inverted color
	HasBackgroundWarningLight       Color = "has-background-warning-light"        // Light variant of warning
	HasBackgroundWarningDark        Color = "has-background-warning-dark"         // Dark variant of warning
	HasBackgroundWarningBold        Color = "has-background-warning-bold"         // Bold variant of warning
	HasBackgroundWarningSoft        Color = "has-background-warning-soft"         // Soft variant of warning
	HasBackgroundWarningOnScheme    Color = "has-background-warning-on-scheme"    // Warning adapted for current scheme
	HasBackgroundWarningLightInvert Color = "has-background-warning-light-invert" // Inverted light warning
	HasBackgroundWarningDarkInvert  Color = "has-background-warning-dark-invert"  // Inverted dark warning
	HasBackgroundWarningBoldInvert  Color = "has-background-warning-bold-invert"  // Inverted bold warning
	HasBackgroundWarningSoftInvert  Color = "has-background-warning-soft-invert"  // Inverted soft warning

	// Background Color Helpers - White Opacity Variations (00-100)
	HasBackgroundWhite00  Color = "has-background-white-00"  // White 00% opacity (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent)
	HasBackgroundWhite05  Color = "has-background-white-05"  // White 05% opacity
	HasBackgroundWhite10  Color = "has-background-white-10"  // White 10% opacity
	HasBackgroundWhite15  Color = "has-background-white-15"  // White 15% opacity
	HasBackgroundWhite20  Color = "has-background-white-20"  // White 20% opacity
	HasBackgroundWhite25  Color = "has-background-white-25"  // White 25% opacity
	HasBackgroundWhite30  Color = "has-background-white-30"  // White 30% opacity
	HasBackgroundWhite35  Color = "has-background-white-35"  // White 35% opacity
	HasBackgroundWhite40  Color = "has-background-white-40"  // White 40% opacity
	HasBackgroundWhite45  Color = "has-background-white-45"  // White 45% opacity
	HasBackgroundWhite50  Color = "has-background-white-50"  // White 50% opacity
	HasBackgroundWhite55  Color = "has-background-white-55"  // White 55% opacity
	HasBackgroundWhite60  Color = "has-background-white-60"  // White 60% opacity
	HasBackgroundWhite65  Color = "has-background-white-65"  // White 65% opacity
	HasBackgroundWhite70  Color = "has-background-white-70"  // White 70% opacity
	HasBackgroundWhite75  Color = "has-background-white-75"  // White 75% opacity
	HasBackgroundWhite80  Color = "has-background-white-80"  // White 80% opacity
	HasBackgroundWhite85  Color = "has-background-white-85"  // White 85% opacity
	HasBackgroundWhite90  Color = "has-background-white-90"  // White 90% opacity
	HasBackgroundWhite95  Color = "has-background-white-95"  // White 95% opacity
	HasBackgroundWhite100 Color = "has-background-white-100" // White 100% opacity (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid)

	// Background Color Helpers - White Opacity Inverts
	HasBackgroundWhite00Invert  Color = "has-background-white-00-invert"  // Inverted contrast for white-00
	HasBackgroundWhite05Invert  Color = "has-background-white-05-invert"  // Inverted contrast for white-05
	HasBackgroundWhite10Invert  Color = "has-background-white-10-invert"  // Inverted contrast for white-10
	HasBackgroundWhite15Invert  Color = "has-background-white-15-invert"  // Inverted contrast for white-15
	HasBackgroundWhite20Invert  Color = "has-background-white-20-invert"  // Inverted contrast for white-20
	HasBackgroundWhite25Invert  Color = "has-background-white-25-invert"  // Inverted contrast for white-25
	HasBackgroundWhite30Invert  Color = "has-background-white-30-invert"  // Inverted contrast for white-30
	HasBackgroundWhite35Invert  Color = "has-background-white-35-invert"  // Inverted contrast for white-35
	HasBackgroundWhite40Invert  Color = "has-background-white-40-invert"  // Inverted contrast for white-40
	HasBackgroundWhite45Invert  Color = "has-background-white-45-invert"  // Inverted contrast for white-45
	HasBackgroundWhite50Invert  Color = "has-background-white-50-invert"  // Inverted contrast for white-50
	HasBackgroundWhite55Invert  Color = "has-background-white-55-invert"  // Inverted contrast for white-55
	HasBackgroundWhite60Invert  Color = "has-background-white-60-invert"  // Inverted contrast for white-60
	HasBackgroundWhite65Invert  Color = "has-background-white-65-invert"  // Inverted contrast for white-65
	HasBackgroundWhite70Invert  Color = "has-background-white-70-invert"  // Inverted contrast for white-70
	HasBackgroundWhite75Invert  Color = "has-background-white-75-invert"  // Inverted contrast for white-75
	HasBackgroundWhite80Invert  Color = "has-background-white-80-invert"  // Inverted contrast for white-80
	HasBackgroundWhite85Invert  Color = "has-background-white-85-invert"  // Inverted contrast for white-85
	HasBackgroundWhite90Invert  Color = "has-background-white-90-invert"  // Inverted contrast for white-90
	HasBackgroundWhite95Invert  Color = "has-background-white-95-invert"  // Inverted contrast for white-95
	HasBackgroundWhite100Invert Color = "has-background-white-100-invert" // Inverted contrast for white-100

	// Background Color Helpers - White Semantic Variations
	HasBackgroundWhiteInvert      Color = "has-background-white-invert"       // White inverted color
	HasBackgroundWhiteLight       Color = "has-background-white-light"        // Light variant of white
	HasBackgroundWhiteDark        Color = "has-background-white-dark"         // Dark variant of white
	HasBackgroundWhiteBold        Color = "has-background-white-bold"         // Bold variant of white
	HasBackgroundWhiteSoft        Color = "has-background-white-soft"         // Soft variant of white
	HasBackgroundWhiteOnScheme    Color = "has-background-white-on-scheme"    // White adapted for current scheme
	HasBackgroundWhiteLightInvert Color = "has-background-white-light-invert" // Inverted light white
	HasBackgroundWhiteDarkInvert  Color = "has-background-white-dark-invert"  // Inverted dark white
	HasBackgroundWhiteBoldInvert  Color = "has-background-white-bold-invert"  // Inverted bold white
	HasBackgroundWhiteSoftInvert  Color = "has-background-white-soft-invert"  // Inverted soft white
	HasBackgroundWhiteBis         Color = "has-background-white-bis"          // White-bis monochrome variant
	HasBackgroundWhiteTer         Color = "has-background-white-ter"          // White-ter monochrome variant

	// Text Color Helpers - Base Colors
	HasTextBlack   Color = "has-text-black"   // Black text color
	HasTextDanger  Color = "has-text-danger"  // Danger/error text (red)
	HasTextDark    Color = "has-text-dark"    // Dark text color
	HasTextInfo    Color = "has-text-info"    // Info text (blue)
	HasTextLight   Color = "has-text-light"   // Light text color
	HasTextLink    Color = "has-text-link"    // Link text color
	HasTextPrimary Color = "has-text-primary" // Primary brand text
	HasTextSuccess Color = "has-text-success" // Success text (green)
	HasTextText    Color = "has-text-text"    // Default text color
	HasTextWarning Color = "has-text-warning" // Warning text (yellow)
	HasTextWhite   Color = "has-text-white"   // White text color

	// Text Color Helpers - Monochrome Grey Variations
	HasTextGrey        Color = "has-text-grey"         // Standard grey text
	HasTextGreyDark    Color = "has-text-grey-dark"    // Dark grey text
	HasTextGreyDarker  Color = "has-text-grey-darker"  // Darker grey text
	HasTextGreyLight   Color = "has-text-grey-light"   // Light grey text
	HasTextGreyLighter Color = "has-text-grey-lighter" // Lighter grey text

	// Text Color Helpers - Special Values
	HasTextCurrent Color = "has-text-current" // Use current text color
	HasTextInherit Color = "has-text-inherit" // Inherit text color from parent

	// Text Color Helpers - Black Opacity Variations (00-100)
	HasTextBlack00  Color = "has-text-black-00"  // Black text 00% opacity (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent)
	HasTextBlack05  Color = "has-text-black-05"  // Black text 05% opacity
	HasTextBlack10  Color = "has-text-black-10"  // Black text 10% opacity
	HasTextBlack15  Color = "has-text-black-15"  // Black text 15% opacity
	HasTextBlack20  Color = "has-text-black-20"  // Black text 20% opacity
	HasTextBlack25  Color = "has-text-black-25"  // Black text 25% opacity
	HasTextBlack30  Color = "has-text-black-30"  // Black text 30% opacity
	HasTextBlack35  Color = "has-text-black-35"  // Black text 35% opacity
	HasTextBlack40  Color = "has-text-black-40"  // Black text 40% opacity
	HasTextBlack45  Color = "has-text-black-45"  // Black text 45% opacity
	HasTextBlack50  Color = "has-text-black-50"  // Black text 50% opacity
	HasTextBlack55  Color = "has-text-black-55"  // Black text 55% opacity
	HasTextBlack60  Color = "has-text-black-60"  // Black text 60% opacity
	HasTextBlack65  Color = "has-text-black-65"  // Black text 65% opacity
	HasTextBlack70  Color = "has-text-black-70"  // Black text 70% opacity
	HasTextBlack75  Color = "has-text-black-75"  // Black text 75% opacity
	HasTextBlack80  Color = "has-text-black-80"  // Black text 80% opacity
	HasTextBlack85  Color = "has-text-black-85"  // Black text 85% opacity
	HasTextBlack90  Color = "has-text-black-90"  // Black text 90% opacity
	HasTextBlack95  Color = "has-text-black-95"  // Black text 95% opacity
	HasTextBlack100 Color = "has-text-black-100" // Black text 100% opacity (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid)

	// Text Color Helpers - Black Opacity Inverts
	HasTextBlack00Invert  Color = "has-text-black-00-invert"  // Inverted contrast for black text-00
	HasTextBlack05Invert  Color = "has-text-black-05-invert"  // Inverted contrast for black text-05
	HasTextBlack10Invert  Color = "has-text-black-10-invert"  // Inverted contrast for black text-10
	HasTextBlack15Invert  Color = "has-text-black-15-invert"  // Inverted contrast for black text-15
	HasTextBlack20Invert  Color = "has-text-black-20-invert"  // Inverted contrast for black text-20
	HasTextBlack25Invert  Color = "has-text-black-25-invert"  // Inverted contrast for black text-25
	HasTextBlack30Invert  Color = "has-text-black-30-invert"  // Inverted contrast for black text-30
	HasTextBlack35Invert  Color = "has-text-black-35-invert"  // Inverted contrast for black text-35
	HasTextBlack40Invert  Color = "has-text-black-40-invert"  // Inverted contrast for black text-40
	HasTextBlack45Invert  Color = "has-text-black-45-invert"  // Inverted contrast for black text-45
	HasTextBlack50Invert  Color = "has-text-black-50-invert"  // Inverted contrast for black text-50
	HasTextBlack55Invert  Color = "has-text-black-55-invert"  // Inverted contrast for black text-55
	HasTextBlack60Invert  Color = "has-text-black-60-invert"  // Inverted contrast for black text-60
	HasTextBlack65Invert  Color = "has-text-black-65-invert"  // Inverted contrast for black text-65
	HasTextBlack70Invert  Color = "has-text-black-70-invert"  // Inverted contrast for black text-70
	HasTextBlack75Invert  Color = "has-text-black-75-invert"  // Inverted contrast for black text-75
	HasTextBlack80Invert  Color = "has-text-black-80-invert"  // Inverted contrast for black text-80
	HasTextBlack85Invert  Color = "has-text-black-85-invert"  // Inverted contrast for black text-85
	HasTextBlack90Invert  Color = "has-text-black-90-invert"  // Inverted contrast for black text-90
	HasTextBlack95Invert  Color = "has-text-black-95-invert"  // Inverted contrast for black text-95
	HasTextBlack100Invert Color = "has-text-black-100-invert" // Inverted contrast for black text-100

	// Text Color Helpers - Black Semantic Variations
	HasTextBlackInvert      Color = "has-text-black-invert"       // Black inverted text color
	HasTextBlackLight       Color = "has-text-black-light"        // Light variant of black text
	HasTextBlackDark        Color = "has-text-black-dark"         // Dark variant of black text
	HasTextBlackBold        Color = "has-text-black-bold"         // Bold variant of black text
	HasTextBlackSoft        Color = "has-text-black-soft"         // Soft variant of black text
	HasTextBlackOnScheme    Color = "has-text-black-on-scheme"    // Black text adapted for current scheme
	HasTextBlackLightInvert Color = "has-text-black-light-invert" // Inverted light black text
	HasTextBlackDarkInvert  Color = "has-text-black-dark-invert"  // Inverted dark black text
	HasTextBlackBoldInvert  Color = "has-text-black-bold-invert"  // Inverted bold black text
	HasTextBlackSoftInvert  Color = "has-text-black-soft-invert"  // Inverted soft black text
	HasTextBlackBis         Color = "has-text-black-bis"          // Black-bis monochrome text variant
	HasTextBlackTer         Color = "has-text-black-ter"          // Black-ter monochrome text variant

	// Text Color Helpers - Danger Opacity Variations (00-100)
	HasTextDanger00  Color = "has-text-danger-00"  // Danger text 00% opacity (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent)
	HasTextDanger05  Color = "has-text-danger-05"  // Danger text 05% opacity
	HasTextDanger10  Color = "has-text-danger-10"  // Danger text 10% opacity
	HasTextDanger15  Color = "has-text-danger-15"  // Danger text 15% opacity
	HasTextDanger20  Color = "has-text-danger-20"  // Danger text 20% opacity
	HasTextDanger25  Color = "has-text-danger-25"  // Danger text 25% opacity
	HasTextDanger30  Color = "has-text-danger-30"  // Danger text 30% opacity
	HasTextDanger35  Color = "has-text-danger-35"  // Danger text 35% opacity
	HasTextDanger40  Color = "has-text-danger-40"  // Danger text 40% opacity
	HasTextDanger45  Color = "has-text-danger-45"  // Danger text 45% opacity
	HasTextDanger50  Color = "has-text-danger-50"  // Danger text 50% opacity
	HasTextDanger55  Color = "has-text-danger-55"  // Danger text 55% opacity
	HasTextDanger60  Color = "has-text-danger-60"  // Danger text 60% opacity
	HasTextDanger65  Color = "has-text-danger-65"  // Danger text 65% opacity
	HasTextDanger70  Color = "has-text-danger-70"  // Danger text 70% opacity
	HasTextDanger75  Color = "has-text-danger-75"  // Danger text 75% opacity
	HasTextDanger80  Color = "has-text-danger-80"  // Danger text 80% opacity
	HasTextDanger85  Color = "has-text-danger-85"  // Danger text 85% opacity
	HasTextDanger90  Color = "has-text-danger-90"  // Danger text 90% opacity
	HasTextDanger95  Color = "has-text-danger-95"  // Danger text 95% opacity
	HasTextDanger100 Color = "has-text-danger-100" // Danger text 100% opacity (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid)

	// Text Color Helpers - Danger Opacity Inverts
	HasTextDanger00Invert  Color = "has-text-danger-00-invert"  // Inverted contrast for danger text-00
	HasTextDanger05Invert  Color = "has-text-danger-05-invert"  // Inverted contrast for danger text-05
	HasTextDanger10Invert  Color = "has-text-danger-10-invert"  // Inverted contrast for danger text-10
	HasTextDanger15Invert  Color = "has-text-danger-15-invert"  // Inverted contrast for danger text-15
	HasTextDanger20Invert  Color = "has-text-danger-20-invert"  // Inverted contrast for danger text-20
	HasTextDanger25Invert  Color = "has-text-danger-25-invert"  // Inverted contrast for danger text-25
	HasTextDanger30Invert  Color = "has-text-danger-30-invert"  // Inverted contrast for danger text-30
	HasTextDanger35Invert  Color = "has-text-danger-35-invert"  // Inverted contrast for danger text-35
	HasTextDanger40Invert  Color = "has-text-danger-40-invert"  // Inverted contrast for danger text-40
	HasTextDanger45Invert  Color = "has-text-danger-45-invert"  // Inverted contrast for danger text-45
	HasTextDanger50Invert  Color = "has-text-danger-50-invert"  // Inverted contrast for danger text-50
	HasTextDanger55Invert  Color = "has-text-danger-55-invert"  // Inverted contrast for danger text-55
	HasTextDanger60Invert  Color = "has-text-danger-60-invert"  // Inverted contrast for danger text-60
	HasTextDanger65Invert  Color = "has-text-danger-65-invert"  // Inverted contrast for danger text-65
	HasTextDanger70Invert  Color = "has-text-danger-70-invert"  // Inverted contrast for danger text-70
	HasTextDanger75Invert  Color = "has-text-danger-75-invert"  // Inverted contrast for danger text-75
	HasTextDanger80Invert  Color = "has-text-danger-80-invert"  // Inverted contrast for danger text-80
	HasTextDanger85Invert  Color = "has-text-danger-85-invert"  // Inverted contrast for danger text-85
	HasTextDanger90Invert  Color = "has-text-danger-90-invert"  // Inverted contrast for danger text-90
	HasTextDanger95Invert  Color = "has-text-danger-95-invert"  // Inverted contrast for danger text-95
	HasTextDanger100Invert Color = "has-text-danger-100-invert" // Inverted contrast for danger text-100

	// Text Color Helpers - Danger Semantic Variations
	HasTextDangerInvert      Color = "has-text-danger-invert"       // Danger inverted text color
	HasTextDangerLight       Color = "has-text-danger-light"        // Light variant of danger text
	HasTextDangerDark        Color = "has-text-danger-dark"         // Dark variant of danger text
	HasTextDangerBold        Color = "has-text-danger-bold"         // Bold variant of danger text
	HasTextDangerSoft        Color = "has-text-danger-soft"         // Soft variant of danger text
	HasTextDangerOnScheme    Color = "has-text-danger-on-scheme"    // Danger text adapted for current scheme
	HasTextDangerLightInvert Color = "has-text-danger-light-invert" // Inverted light danger text
	HasTextDangerDarkInvert  Color = "has-text-danger-dark-invert"  // Inverted dark danger text
	HasTextDangerBoldInvert  Color = "has-text-danger-bold-invert"  // Inverted bold danger text
	HasTextDangerSoftInvert  Color = "has-text-danger-soft-invert"  // Inverted soft danger text

	// Text Color Helpers - Dark Opacity Variations (00-100)
	HasTextDark00  Color = "has-text-dark-00"  // Dark text 00% opacity (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent)
	HasTextDark05  Color = "has-text-dark-05"  // Dark text 05% opacity
	HasTextDark10  Color = "has-text-dark-10"  // Dark text 10% opacity
	HasTextDark15  Color = "has-text-dark-15"  // Dark text 15% opacity
	HasTextDark20  Color = "has-text-dark-20"  // Dark text 20% opacity
	HasTextDark25  Color = "has-text-dark-25"  // Dark text 25% opacity
	HasTextDark30  Color = "has-text-dark-30"  // Dark text 30% opacity
	HasTextDark35  Color = "has-text-dark-35"  // Dark text 35% opacity
	HasTextDark40  Color = "has-text-dark-40"  // Dark text 40% opacity
	HasTextDark45  Color = "has-text-dark-45"  // Dark text 45% opacity
	HasTextDark50  Color = "has-text-dark-50"  // Dark text 50% opacity
	HasTextDark55  Color = "has-text-dark-55"  // Dark text 55% opacity
	HasTextDark60  Color = "has-text-dark-60"  // Dark text 60% opacity
	HasTextDark65  Color = "has-text-dark-65"  // Dark text 65% opacity
	HasTextDark70  Color = "has-text-dark-70"  // Dark text 70% opacity
	HasTextDark75  Color = "has-text-dark-75"  // Dark text 75% opacity
	HasTextDark80  Color = "has-text-dark-80"  // Dark text 80% opacity
	HasTextDark85  Color = "has-text-dark-85"  // Dark text 85% opacity
	HasTextDark90  Color = "has-text-dark-90"  // Dark text 90% opacity
	HasTextDark95  Color = "has-text-dark-95"  // Dark text 95% opacity
	HasTextDark100 Color = "has-text-dark-100" // Dark text 100% opacity (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid)

	// Text Color Helpers - Dark Opacity Inverts
	HasTextDark00Invert  Color = "has-text-dark-00-invert"  // Inverted contrast for dark text-00
	HasTextDark05Invert  Color = "has-text-dark-05-invert"  // Inverted contrast for dark text-05
	HasTextDark10Invert  Color = "has-text-dark-10-invert"  // Inverted contrast for dark text-10
	HasTextDark15Invert  Color = "has-text-dark-15-invert"  // Inverted contrast for dark text-15
	HasTextDark20Invert  Color = "has-text-dark-20-invert"  // Inverted contrast for dark text-20
	HasTextDark25Invert  Color = "has-text-dark-25-invert"  // Inverted contrast for dark text-25
	HasTextDark30Invert  Color = "has-text-dark-30-invert"  // Inverted contrast for dark text-30
	HasTextDark35Invert  Color = "has-text-dark-35-invert"  // Inverted contrast for dark text-35
	HasTextDark40Invert  Color = "has-text-dark-40-invert"  // Inverted contrast for dark text-40
	HasTextDark45Invert  Color = "has-text-dark-45-invert"  // Inverted contrast for dark text-45
	HasTextDark50Invert  Color = "has-text-dark-50-invert"  // Inverted contrast for dark text-50
	HasTextDark55Invert  Color = "has-text-dark-55-invert"  // Inverted contrast for dark text-55
	HasTextDark60Invert  Color = "has-text-dark-60-invert"  // Inverted contrast for dark text-60
	HasTextDark65Invert  Color = "has-text-dark-65-invert"  // Inverted contrast for dark text-65
	HasTextDark70Invert  Color = "has-text-dark-70-invert"  // Inverted contrast for dark text-70
	HasTextDark75Invert  Color = "has-text-dark-75-invert"  // Inverted contrast for dark text-75
	HasTextDark80Invert  Color = "has-text-dark-80-invert"  // Inverted contrast for dark text-80
	HasTextDark85Invert  Color = "has-text-dark-85-invert"  // Inverted contrast for dark text-85
	HasTextDark90Invert  Color = "has-text-dark-90-invert"  // Inverted contrast for dark text-90
	HasTextDark95Invert  Color = "has-text-dark-95-invert"  // Inverted contrast for dark text-95
	HasTextDark100Invert Color = "has-text-dark-100-invert" // Inverted contrast for dark text-100

	// Text Color Helpers - Dark Semantic Variations
	HasTextDarkInvert      Color = "has-text-dark-invert"       // Dark inverted text color
	HasTextDarkLight       Color = "has-text-dark-light"        // Light variant of dark text
	HasTextDarkDark        Color = "has-text-dark-dark"         // Dark variant of dark text
	HasTextDarkBold        Color = "has-text-dark-bold"         // Bold variant of dark text
	HasTextDarkSoft        Color = "has-text-dark-soft"         // Soft variant of dark text
	HasTextDarkOnScheme    Color = "has-text-dark-on-scheme"    // Dark text adapted for current scheme
	HasTextDarkLightInvert Color = "has-text-dark-light-invert" // Inverted light dark text
	HasTextDarkDarkInvert  Color = "has-text-dark-dark-invert"  // Inverted dark dark text
	HasTextDarkBoldInvert  Color = "has-text-dark-bold-invert"  // Inverted bold dark text
	HasTextDarkSoftInvert  Color = "has-text-dark-soft-invert"  // Inverted soft dark text

	// Text Color Helpers - Info Opacity Variations (00-100)
	HasTextInfo00  Color = "has-text-info-00"  // Info text 00% opacity (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent)
	HasTextInfo05  Color = "has-text-info-05"  // Info text 05% opacity
	HasTextInfo10  Color = "has-text-info-10"  // Info text 10% opacity
	HasTextInfo15  Color = "has-text-info-15"  // Info text 15% opacity
	HasTextInfo20  Color = "has-text-info-20"  // Info text 20% opacity
	HasTextInfo25  Color = "has-text-info-25"  // Info text 25% opacity
	HasTextInfo30  Color = "has-text-info-30"  // Info text 30% opacity
	HasTextInfo35  Color = "has-text-info-35"  // Info text 35% opacity
	HasTextInfo40  Color = "has-text-info-40"  // Info text 40% opacity
	HasTextInfo45  Color = "has-text-info-45"  // Info text 45% opacity
	HasTextInfo50  Color = "has-text-info-50"  // Info text 50% opacity
	HasTextInfo55  Color = "has-text-info-55"  // Info text 55% opacity
	HasTextInfo60  Color = "has-text-info-60"  // Info text 60% opacity
	HasTextInfo65  Color = "has-text-info-65"  // Info text 65% opacity
	HasTextInfo70  Color = "has-text-info-70"  // Info text 70% opacity
	HasTextInfo75  Color = "has-text-info-75"  // Info text 75% opacity
	HasTextInfo80  Color = "has-text-info-80"  // Info text 80% opacity
	HasTextInfo85  Color = "has-text-info-85"  // Info text 85% opacity
	HasTextInfo90  Color = "has-text-info-90"  // Info text 90% opacity
	HasTextInfo95  Color = "has-text-info-95"  // Info text 95% opacity
	HasTextInfo100 Color = "has-text-info-100" // Info text 100% opacity (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid)

	// Text Color Helpers - Info Opacity Inverts
	HasTextInfo00Invert  Color = "has-text-info-00-invert"  // Inverted contrast for info text-00
	HasTextInfo05Invert  Color = "has-text-info-05-invert"  // Inverted contrast for info text-05
	HasTextInfo10Invert  Color = "has-text-info-10-invert"  // Inverted contrast for info text-10
	HasTextInfo15Invert  Color = "has-text-info-15-invert"  // Inverted contrast for info text-15
	HasTextInfo20Invert  Color = "has-text-info-20-invert"  // Inverted contrast for info text-20
	HasTextInfo25Invert  Color = "has-text-info-25-invert"  // Inverted contrast for info text-25
	HasTextInfo30Invert  Color = "has-text-info-30-invert"  // Inverted contrast for info text-30
	HasTextInfo35Invert  Color = "has-text-info-35-invert"  // Inverted contrast for info text-35
	HasTextInfo40Invert  Color = "has-text-info-40-invert"  // Inverted contrast for info text-40
	HasTextInfo45Invert  Color = "has-text-info-45-invert"  // Inverted contrast for info text-45
	HasTextInfo50Invert  Color = "has-text-info-50-invert"  // Inverted contrast for info text-50
	HasTextInfo55Invert  Color = "has-text-info-55-invert"  // Inverted contrast for info text-55
	HasTextInfo60Invert  Color = "has-text-info-60-invert"  // Inverted contrast for info text-60
	HasTextInfo65Invert  Color = "has-text-info-65-invert"  // Inverted contrast for info text-65
	HasTextInfo70Invert  Color = "has-text-info-70-invert"  // Inverted contrast for info text-70
	HasTextInfo75Invert  Color = "has-text-info-75-invert"  // Inverted contrast for info text-75
	HasTextInfo80Invert  Color = "has-text-info-80-invert"  // Inverted contrast for info text-80
	HasTextInfo85Invert  Color = "has-text-info-85-invert"  // Inverted contrast for info text-85
	HasTextInfo90Invert  Color = "has-text-info-90-invert"  // Inverted contrast for info text-90
	HasTextInfo95Invert  Color = "has-text-info-95-invert"  // Inverted contrast for info text-95
	HasTextInfo100Invert Color = "has-text-info-100-invert" // Inverted contrast for info text-100

	// Text Color Helpers - Info Semantic Variations
	HasTextInfoInvert      Color = "has-text-info-invert"       // Info inverted text color
	HasTextInfoLight       Color = "has-text-info-light"        // Light variant of info text
	HasTextInfoDark        Color = "has-text-info-dark"         // Dark variant of info text
	HasTextInfoBold        Color = "has-text-info-bold"         // Bold variant of info text
	HasTextInfoSoft        Color = "has-text-info-soft"         // Soft variant of info text
	HasTextInfoOnScheme    Color = "has-text-info-on-scheme"    // Info text adapted for current scheme
	HasTextInfoLightInvert Color = "has-text-info-light-invert" // Inverted light info text
	HasTextInfoDarkInvert  Color = "has-text-info-dark-invert"  // Inverted dark info text
	HasTextInfoBoldInvert  Color = "has-text-info-bold-invert"  // Inverted bold info text
	HasTextInfoSoftInvert  Color = "has-text-info-soft-invert"  // Inverted soft info text

	// Text Color Helpers - Light Opacity Variations (00-100)
	HasTextLight00  Color = "has-text-light-00"  // Light text 00% opacity (transparent) (transparent) (transparent) (transparent) (transparent) (transparent) (transparent)
	HasTextLight05  Color = "has-text-light-05"  // Light text 05% opacity
	HasTextLight10  Color = "has-text-light-10"  // Light text 10% opacity
	HasTextLight15  Color = "has-text-light-15"  // Light text 15% opacity
	HasTextLight20  Color = "has-text-light-20"  // Light text 20% opacity
	HasTextLight25  Color = "has-text-light-25"  // Light text 25% opacity
	HasTextLight30  Color = "has-text-light-30"  // Light text 30% opacity
	HasTextLight35  Color = "has-text-light-35"  // Light text 35% opacity
	HasTextLight40  Color = "has-text-light-40"  // Light text 40% opacity
	HasTextLight45  Color = "has-text-light-45"  // Light text 45% opacity
	HasTextLight50  Color = "has-text-light-50"  // Light text 50% opacity
	HasTextLight55  Color = "has-text-light-55"  // Light text 55% opacity
	HasTextLight60  Color = "has-text-light-60"  // Light text 60% opacity
	HasTextLight65  Color = "has-text-light-65"  // Light text 65% opacity
	HasTextLight70  Color = "has-text-light-70"  // Light text 70% opacity
	HasTextLight75  Color = "has-text-light-75"  // Light text 75% opacity
	HasTextLight80  Color = "has-text-light-80"  // Light text 80% opacity
	HasTextLight85  Color = "has-text-light-85"  // Light text 85% opacity
	HasTextLight90  Color = "has-text-light-90"  // Light text 90% opacity
	HasTextLight95  Color = "has-text-light-95"  // Light text 95% opacity
	HasTextLight100 Color = "has-text-light-100" // Light text 100% opacity (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid)

	// Text Color Helpers - Light Opacity Inverts
	HasTextLight00Invert  Color = "has-text-light-00-invert"  // Inverted contrast for light text-00
	HasTextLight05Invert  Color = "has-text-light-05-invert"  // Inverted contrast for light text-05
	HasTextLight10Invert  Color = "has-text-light-10-invert"  // Inverted contrast for light text-10
	HasTextLight15Invert  Color = "has-text-light-15-invert"  // Inverted contrast for light text-15
	HasTextLight20Invert  Color = "has-text-light-20-invert"  // Inverted contrast for light text-20
	HasTextLight25Invert  Color = "has-text-light-25-invert"  // Inverted contrast for light text-25
	HasTextLight30Invert  Color = "has-text-light-30-invert"  // Inverted contrast for light text-30
	HasTextLight35Invert  Color = "has-text-light-35-invert"  // Inverted contrast for light text-35
	HasTextLight40Invert  Color = "has-text-light-40-invert"  // Inverted contrast for light text-40
	HasTextLight45Invert  Color = "has-text-light-45-invert"  // Inverted contrast for light text-45
	HasTextLight50Invert  Color = "has-text-light-50-invert"  // Inverted contrast for light text-50
	HasTextLight55Invert  Color = "has-text-light-55-invert"  // Inverted contrast for light text-55
	HasTextLight60Invert  Color = "has-text-light-60-invert"  // Inverted contrast for light text-60
	HasTextLight65Invert  Color = "has-text-light-65-invert"  // Inverted contrast for light text-65
	HasTextLight70Invert  Color = "has-text-light-70-invert"  // Inverted contrast for light text-70
	HasTextLight75Invert  Color = "has-text-light-75-invert"  // Inverted contrast for light text-75
	HasTextLight80Invert  Color = "has-text-light-80-invert"  // Inverted contrast for light text-80
	HasTextLight85Invert  Color = "has-text-light-85-invert"  // Inverted contrast for light text-85
	HasTextLight90Invert  Color = "has-text-light-90-invert"  // Inverted contrast for light text-90
	HasTextLight95Invert  Color = "has-text-light-95-invert"  // Inverted contrast for light text-95
	HasTextLight100Invert Color = "has-text-light-100-invert" // Inverted contrast for light text-100

	// Text Color Helpers - Light Semantic Variations
	HasTextLightInvert      Color = "has-text-light-invert"       // Light inverted text color
	HasTextLightLight       Color = "has-text-light-light"        // Light variant of light text
	HasTextLightDark        Color = "has-text-light-dark"         // Dark variant of light text
	HasTextLightBold        Color = "has-text-light-bold"         // Bold variant of light text
	HasTextLightSoft        Color = "has-text-light-soft"         // Soft variant of light text
	HasTextLightOnScheme    Color = "has-text-light-on-scheme"    // Light text adapted for current scheme
	HasTextLightLightInvert Color = "has-text-light-light-invert" // Inverted light light text
	HasTextLightDarkInvert  Color = "has-text-light-dark-invert"  // Inverted dark light text
	HasTextLightBoldInvert  Color = "has-text-light-bold-invert"  // Inverted bold light text
	HasTextLightSoftInvert  Color = "has-text-light-soft-invert"  // Inverted soft light text

	// Text Color Helpers - Link Opacity Variations (00-100)
	HasTextLink00  Color = "has-text-link-00"  // Link text 00% opacity (transparent) (transparent) (transparent) (transparent) (transparent) (transparent)
	HasTextLink05  Color = "has-text-link-05"  // Link text 05% opacity
	HasTextLink10  Color = "has-text-link-10"  // Link text 10% opacity
	HasTextLink15  Color = "has-text-link-15"  // Link text 15% opacity
	HasTextLink20  Color = "has-text-link-20"  // Link text 20% opacity
	HasTextLink25  Color = "has-text-link-25"  // Link text 25% opacity
	HasTextLink30  Color = "has-text-link-30"  // Link text 30% opacity
	HasTextLink35  Color = "has-text-link-35"  // Link text 35% opacity
	HasTextLink40  Color = "has-text-link-40"  // Link text 40% opacity
	HasTextLink45  Color = "has-text-link-45"  // Link text 45% opacity
	HasTextLink50  Color = "has-text-link-50"  // Link text 50% opacity
	HasTextLink55  Color = "has-text-link-55"  // Link text 55% opacity
	HasTextLink60  Color = "has-text-link-60"  // Link text 60% opacity
	HasTextLink65  Color = "has-text-link-65"  // Link text 65% opacity
	HasTextLink70  Color = "has-text-link-70"  // Link text 70% opacity
	HasTextLink75  Color = "has-text-link-75"  // Link text 75% opacity
	HasTextLink80  Color = "has-text-link-80"  // Link text 80% opacity
	HasTextLink85  Color = "has-text-link-85"  // Link text 85% opacity
	HasTextLink90  Color = "has-text-link-90"  // Link text 90% opacity
	HasTextLink95  Color = "has-text-link-95"  // Link text 95% opacity
	HasTextLink100 Color = "has-text-link-100" // Link text 100% opacity (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid)

	// Text Color Helpers - Link Opacity Inverts
	HasTextLink00Invert  Color = "has-text-link-00-invert"  // Inverted contrast for link text-00
	HasTextLink05Invert  Color = "has-text-link-05-invert"  // Inverted contrast for link text-05
	HasTextLink10Invert  Color = "has-text-link-10-invert"  // Inverted contrast for link text-10
	HasTextLink15Invert  Color = "has-text-link-15-invert"  // Inverted contrast for link text-15
	HasTextLink20Invert  Color = "has-text-link-20-invert"  // Inverted contrast for link text-20
	HasTextLink25Invert  Color = "has-text-link-25-invert"  // Inverted contrast for link text-25
	HasTextLink30Invert  Color = "has-text-link-30-invert"  // Inverted contrast for link text-30
	HasTextLink35Invert  Color = "has-text-link-35-invert"  // Inverted contrast for link text-35
	HasTextLink40Invert  Color = "has-text-link-40-invert"  // Inverted contrast for link text-40
	HasTextLink45Invert  Color = "has-text-link-45-invert"  // Inverted contrast for link text-45
	HasTextLink50Invert  Color = "has-text-link-50-invert"  // Inverted contrast for link text-50
	HasTextLink55Invert  Color = "has-text-link-55-invert"  // Inverted contrast for link text-55
	HasTextLink60Invert  Color = "has-text-link-60-invert"  // Inverted contrast for link text-60
	HasTextLink65Invert  Color = "has-text-link-65-invert"  // Inverted contrast for link text-65
	HasTextLink70Invert  Color = "has-text-link-70-invert"  // Inverted contrast for link text-70
	HasTextLink75Invert  Color = "has-text-link-75-invert"  // Inverted contrast for link text-75
	HasTextLink80Invert  Color = "has-text-link-80-invert"  // Inverted contrast for link text-80
	HasTextLink85Invert  Color = "has-text-link-85-invert"  // Inverted contrast for link text-85
	HasTextLink90Invert  Color = "has-text-link-90-invert"  // Inverted contrast for link text-90
	HasTextLink95Invert  Color = "has-text-link-95-invert"  // Inverted contrast for link text-95
	HasTextLink100Invert Color = "has-text-link-100-invert" // Inverted contrast for link text-100

	// Text Color Helpers - Link Semantic Variations
	HasTextLinkInvert      Color = "has-text-link-invert"       // Link inverted text color
	HasTextLinkLight       Color = "has-text-link-light"        // Light variant of link text
	HasTextLinkDark        Color = "has-text-link-dark"         // Dark variant of link text
	HasTextLinkBold        Color = "has-text-link-bold"         // Bold variant of link text
	HasTextLinkSoft        Color = "has-text-link-soft"         // Soft variant of link text
	HasTextLinkOnScheme    Color = "has-text-link-on-scheme"    // Link text adapted for current scheme
	HasTextLinkLightInvert Color = "has-text-link-light-invert" // Inverted light link text
	HasTextLinkDarkInvert  Color = "has-text-link-dark-invert"  // Inverted dark link text
	HasTextLinkBoldInvert  Color = "has-text-link-bold-invert"  // Inverted bold link text
	HasTextLinkSoftInvert  Color = "has-text-link-soft-invert"  // Inverted soft link text

	// Text Color Helpers - Primary Opacity Variations (00-100)
	HasTextPrimary00  Color = "has-text-primary-00"  // Primary text 00% opacity (transparent) (transparent) (transparent) (transparent) (transparent)
	HasTextPrimary05  Color = "has-text-primary-05"  // Primary text 05% opacity
	HasTextPrimary10  Color = "has-text-primary-10"  // Primary text 10% opacity
	HasTextPrimary15  Color = "has-text-primary-15"  // Primary text 15% opacity
	HasTextPrimary20  Color = "has-text-primary-20"  // Primary text 20% opacity
	HasTextPrimary25  Color = "has-text-primary-25"  // Primary text 25% opacity
	HasTextPrimary30  Color = "has-text-primary-30"  // Primary text 30% opacity
	HasTextPrimary35  Color = "has-text-primary-35"  // Primary text 35% opacity
	HasTextPrimary40  Color = "has-text-primary-40"  // Primary text 40% opacity
	HasTextPrimary45  Color = "has-text-primary-45"  // Primary text 45% opacity
	HasTextPrimary50  Color = "has-text-primary-50"  // Primary text 50% opacity
	HasTextPrimary55  Color = "has-text-primary-55"  // Primary text 55% opacity
	HasTextPrimary60  Color = "has-text-primary-60"  // Primary text 60% opacity
	HasTextPrimary65  Color = "has-text-primary-65"  // Primary text 65% opacity
	HasTextPrimary70  Color = "has-text-primary-70"  // Primary text 70% opacity
	HasTextPrimary75  Color = "has-text-primary-75"  // Primary text 75% opacity
	HasTextPrimary80  Color = "has-text-primary-80"  // Primary text 80% opacity
	HasTextPrimary85  Color = "has-text-primary-85"  // Primary text 85% opacity
	HasTextPrimary90  Color = "has-text-primary-90"  // Primary text 90% opacity
	HasTextPrimary95  Color = "has-text-primary-95"  // Primary text 95% opacity
	HasTextPrimary100 Color = "has-text-primary-100" // Primary text 100% opacity (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid)

	// Text Color Helpers - Primary Opacity Inverts
	HasTextPrimary00Invert  Color = "has-text-primary-00-invert"  // Inverted contrast for primary text-00
	HasTextPrimary05Invert  Color = "has-text-primary-05-invert"  // Inverted contrast for primary text-05
	HasTextPrimary10Invert  Color = "has-text-primary-10-invert"  // Inverted contrast for primary text-10
	HasTextPrimary15Invert  Color = "has-text-primary-15-invert"  // Inverted contrast for primary text-15
	HasTextPrimary20Invert  Color = "has-text-primary-20-invert"  // Inverted contrast for primary text-20
	HasTextPrimary25Invert  Color = "has-text-primary-25-invert"  // Inverted contrast for primary text-25
	HasTextPrimary30Invert  Color = "has-text-primary-30-invert"  // Inverted contrast for primary text-30
	HasTextPrimary35Invert  Color = "has-text-primary-35-invert"  // Inverted contrast for primary text-35
	HasTextPrimary40Invert  Color = "has-text-primary-40-invert"  // Inverted contrast for primary text-40
	HasTextPrimary45Invert  Color = "has-text-primary-45-invert"  // Inverted contrast for primary text-45
	HasTextPrimary50Invert  Color = "has-text-primary-50-invert"  // Inverted contrast for primary text-50
	HasTextPrimary55Invert  Color = "has-text-primary-55-invert"  // Inverted contrast for primary text-55
	HasTextPrimary60Invert  Color = "has-text-primary-60-invert"  // Inverted contrast for primary text-60
	HasTextPrimary65Invert  Color = "has-text-primary-65-invert"  // Inverted contrast for primary text-65
	HasTextPrimary70Invert  Color = "has-text-primary-70-invert"  // Inverted contrast for primary text-70
	HasTextPrimary75Invert  Color = "has-text-primary-75-invert"  // Inverted contrast for primary text-75
	HasTextPrimary80Invert  Color = "has-text-primary-80-invert"  // Inverted contrast for primary text-80
	HasTextPrimary85Invert  Color = "has-text-primary-85-invert"  // Inverted contrast for primary text-85
	HasTextPrimary90Invert  Color = "has-text-primary-90-invert"  // Inverted contrast for primary text-90
	HasTextPrimary95Invert  Color = "has-text-primary-95-invert"  // Inverted contrast for primary text-95
	HasTextPrimary100Invert Color = "has-text-primary-100-invert" // Inverted contrast for primary text-100

	// Text Color Helpers - Primary Semantic Variations
	HasTextPrimaryInvert      Color = "has-text-primary-invert"       // Primary inverted text color
	HasTextPrimaryLight       Color = "has-text-primary-light"        // Light variant of primary text
	HasTextPrimaryDark        Color = "has-text-primary-dark"         // Dark variant of primary text
	HasTextPrimaryBold        Color = "has-text-primary-bold"         // Bold variant of primary text
	HasTextPrimarySoft        Color = "has-text-primary-soft"         // Soft variant of primary text
	HasTextPrimaryOnScheme    Color = "has-text-primary-on-scheme"    // Primary text adapted for current scheme
	HasTextPrimaryLightInvert Color = "has-text-primary-light-invert" // Inverted light primary text
	HasTextPrimaryDarkInvert  Color = "has-text-primary-dark-invert"  // Inverted dark primary text
	HasTextPrimaryBoldInvert  Color = "has-text-primary-bold-invert"  // Inverted bold primary text
	HasTextPrimarySoftInvert  Color = "has-text-primary-soft-invert"  // Inverted soft primary text

	// Text Color Helpers - Success Opacity Variations (00-100)
	HasTextSuccess00  Color = "has-text-success-00"  // Success text 00% opacity (transparent) (transparent) (transparent) (transparent)
	HasTextSuccess05  Color = "has-text-success-05"  // Success text 05% opacity
	HasTextSuccess10  Color = "has-text-success-10"  // Success text 10% opacity
	HasTextSuccess15  Color = "has-text-success-15"  // Success text 15% opacity
	HasTextSuccess20  Color = "has-text-success-20"  // Success text 20% opacity
	HasTextSuccess25  Color = "has-text-success-25"  // Success text 25% opacity
	HasTextSuccess30  Color = "has-text-success-30"  // Success text 30% opacity
	HasTextSuccess35  Color = "has-text-success-35"  // Success text 35% opacity
	HasTextSuccess40  Color = "has-text-success-40"  // Success text 40% opacity
	HasTextSuccess45  Color = "has-text-success-45"  // Success text 45% opacity
	HasTextSuccess50  Color = "has-text-success-50"  // Success text 50% opacity
	HasTextSuccess55  Color = "has-text-success-55"  // Success text 55% opacity
	HasTextSuccess60  Color = "has-text-success-60"  // Success text 60% opacity
	HasTextSuccess65  Color = "has-text-success-65"  // Success text 65% opacity
	HasTextSuccess70  Color = "has-text-success-70"  // Success text 70% opacity
	HasTextSuccess75  Color = "has-text-success-75"  // Success text 75% opacity
	HasTextSuccess80  Color = "has-text-success-80"  // Success text 80% opacity
	HasTextSuccess85  Color = "has-text-success-85"  // Success text 85% opacity
	HasTextSuccess90  Color = "has-text-success-90"  // Success text 90% opacity
	HasTextSuccess95  Color = "has-text-success-95"  // Success text 95% opacity
	HasTextSuccess100 Color = "has-text-success-100" // Success text 100% opacity (solid) (transparent) (solid) (transparent) (solid) (transparent) (solid)

	// Text Color Helpers - Success Opacity Inverts
	HasTextSuccess00Invert  Color = "has-text-success-00-invert"  // Inverted contrast for success text-00
	HasTextSuccess05Invert  Color = "has-text-success-05-invert"  // Inverted contrast for success text-05
	HasTextSuccess10Invert  Color = "has-text-success-10-invert"  // Inverted contrast for success text-10
	HasTextSuccess15Invert  Color = "has-text-success-15-invert"  // Inverted contrast for success text-15
	HasTextSuccess20Invert  Color = "has-text-success-20-invert"  // Inverted contrast for success text-20
	HasTextSuccess25Invert  Color = "has-text-success-25-invert"  // Inverted contrast for success text-25
	HasTextSuccess30Invert  Color = "has-text-success-30-invert"  // Inverted contrast for success text-30
	HasTextSuccess35Invert  Color = "has-text-success-35-invert"  // Inverted contrast for success text-35
	HasTextSuccess40Invert  Color = "has-text-success-40-invert"  // Inverted contrast for success text-40
	HasTextSuccess45Invert  Color = "has-text-success-45-invert"  // Inverted contrast for success text-45
	HasTextSuccess50Invert  Color = "has-text-success-50-invert"  // Inverted contrast for success text-50
	HasTextSuccess55Invert  Color = "has-text-success-55-invert"  // Inverted contrast for success text-55
	HasTextSuccess60Invert  Color = "has-text-success-60-invert"  // Inverted contrast for success text-60
	HasTextSuccess65Invert  Color = "has-text-success-65-invert"  // Inverted contrast for success text-65
	HasTextSuccess70Invert  Color = "has-text-success-70-invert"  // Inverted contrast for success text-70
	HasTextSuccess75Invert  Color = "has-text-success-75-invert"  // Inverted contrast for success text-75
	HasTextSuccess80Invert  Color = "has-text-success-80-invert"  // Inverted contrast for success text-80
	HasTextSuccess85Invert  Color = "has-text-success-85-invert"  // Inverted contrast for success text-85
	HasTextSuccess90Invert  Color = "has-text-success-90-invert"  // Inverted contrast for success text-90
	HasTextSuccess95Invert  Color = "has-text-success-95-invert"  // Inverted contrast for success text-95
	HasTextSuccess100Invert Color = "has-text-success-100-invert" // Inverted contrast for success text-100

	// Text Color Helpers - Success Semantic Variations
	HasTextSuccessInvert      Color = "has-text-success-invert"       // Success inverted text color
	HasTextSuccessLight       Color = "has-text-success-light"        // Light variant of success text
	HasTextSuccessDark        Color = "has-text-success-dark"         // Dark variant of success text
	HasTextSuccessBold        Color = "has-text-success-bold"         // Bold variant of success text
	HasTextSuccessSoft        Color = "has-text-success-soft"         // Soft variant of success text
	HasTextSuccessOnScheme    Color = "has-text-success-on-scheme"    // Success text adapted for current scheme
	HasTextSuccessLightInvert Color = "has-text-success-light-invert" // Inverted light success text
	HasTextSuccessDarkInvert  Color = "has-text-success-dark-invert"  // Inverted dark success text
	HasTextSuccessBoldInvert  Color = "has-text-success-bold-invert"  // Inverted bold success text
	HasTextSuccessSoftInvert  Color = "has-text-success-soft-invert"  // Inverted soft success text

	// Text Color Helpers - Text Opacity Variations (00-100)
	HasTextText00  Color = "has-text-text-00"  // Text text 00% opacity (transparent) (transparent) (transparent)
	HasTextText05  Color = "has-text-text-05"  // Text text 05% opacity
	HasTextText10  Color = "has-text-text-10"  // Text text 10% opacity
	HasTextText15  Color = "has-text-text-15"  // Text text 15% opacity
	HasTextText20  Color = "has-text-text-20"  // Text text 20% opacity
	HasTextText25  Color = "has-text-text-25"  // Text text 25% opacity
	HasTextText30  Color = "has-text-text-30"  // Text text 30% opacity
	HasTextText35  Color = "has-text-text-35"  // Text text 35% opacity
	HasTextText40  Color = "has-text-text-40"  // Text text 40% opacity
	HasTextText45  Color = "has-text-text-45"  // Text text 45% opacity
	HasTextText50  Color = "has-text-text-50"  // Text text 50% opacity
	HasTextText55  Color = "has-text-text-55"  // Text text 55% opacity
	HasTextText60  Color = "has-text-text-60"  // Text text 60% opacity
	HasTextText65  Color = "has-text-text-65"  // Text text 65% opacity
	HasTextText70  Color = "has-text-text-70"  // Text text 70% opacity
	HasTextText75  Color = "has-text-text-75"  // Text text 75% opacity
	HasTextText80  Color = "has-text-text-80"  // Text text 80% opacity
	HasTextText85  Color = "has-text-text-85"  // Text text 85% opacity
	HasTextText90  Color = "has-text-text-90"  // Text text 90% opacity
	HasTextText95  Color = "has-text-text-95"  // Text text 95% opacity
	HasTextText100 Color = "has-text-text-100" // Text text 100% opacity (solid) (transparent) (solid) (transparent) (solid)

	// Text Color Helpers - Text Opacity Inverts
	HasTextText00Invert  Color = "has-text-text-00-invert"  // Inverted contrast for text text-00
	HasTextText05Invert  Color = "has-text-text-05-invert"  // Inverted contrast for text text-05
	HasTextText10Invert  Color = "has-text-text-10-invert"  // Inverted contrast for text text-10
	HasTextText15Invert  Color = "has-text-text-15-invert"  // Inverted contrast for text text-15
	HasTextText20Invert  Color = "has-text-text-20-invert"  // Inverted contrast for text text-20
	HasTextText25Invert  Color = "has-text-text-25-invert"  // Inverted contrast for text text-25
	HasTextText30Invert  Color = "has-text-text-30-invert"  // Inverted contrast for text text-30
	HasTextText35Invert  Color = "has-text-text-35-invert"  // Inverted contrast for text text-35
	HasTextText40Invert  Color = "has-text-text-40-invert"  // Inverted contrast for text text-40
	HasTextText45Invert  Color = "has-text-text-45-invert"  // Inverted contrast for text text-45
	HasTextText50Invert  Color = "has-text-text-50-invert"  // Inverted contrast for text text-50
	HasTextText55Invert  Color = "has-text-text-55-invert"  // Inverted contrast for text text-55
	HasTextText60Invert  Color = "has-text-text-60-invert"  // Inverted contrast for text text-60
	HasTextText65Invert  Color = "has-text-text-65-invert"  // Inverted contrast for text text-65
	HasTextText70Invert  Color = "has-text-text-70-invert"  // Inverted contrast for text text-70
	HasTextText75Invert  Color = "has-text-text-75-invert"  // Inverted contrast for text text-75
	HasTextText80Invert  Color = "has-text-text-80-invert"  // Inverted contrast for text text-80
	HasTextText85Invert  Color = "has-text-text-85-invert"  // Inverted contrast for text text-85
	HasTextText90Invert  Color = "has-text-text-90-invert"  // Inverted contrast for text text-90
	HasTextText95Invert  Color = "has-text-text-95-invert"  // Inverted contrast for text text-95
	HasTextText100Invert Color = "has-text-text-100-invert" // Inverted contrast for text text-100

	// Text Color Helpers - Text Semantic Variations
	HasTextTextInvert      Color = "has-text-text-invert"       // Text inverted text color
	HasTextTextLight       Color = "has-text-text-light"        // Light variant of text text
	HasTextTextDark        Color = "has-text-text-dark"         // Dark variant of text text
	HasTextTextBold        Color = "has-text-text-bold"         // Bold variant of text text
	HasTextTextSoft        Color = "has-text-text-soft"         // Soft variant of text text
	HasTextTextOnScheme    Color = "has-text-text-on-scheme"    // Text text adapted for current scheme
	HasTextTextLightInvert Color = "has-text-text-light-invert" // Inverted light text text
	HasTextTextDarkInvert  Color = "has-text-text-dark-invert"  // Inverted dark text text
	HasTextTextBoldInvert  Color = "has-text-text-bold-invert"  // Inverted bold text text
	HasTextTextSoftInvert  Color = "has-text-text-soft-invert"  // Inverted soft text text

	// Text Color Helpers - Warning Opacity Variations (00-100)
	HasTextWarning00  Color = "has-text-warning-00"  // Warning text 00% opacity (transparent) (transparent)
	HasTextWarning05  Color = "has-text-warning-05"  // Warning text 05% opacity
	HasTextWarning10  Color = "has-text-warning-10"  // Warning text 10% opacity
	HasTextWarning15  Color = "has-text-warning-15"  // Warning text 15% opacity
	HasTextWarning20  Color = "has-text-warning-20"  // Warning text 20% opacity
	HasTextWarning25  Color = "has-text-warning-25"  // Warning text 25% opacity
	HasTextWarning30  Color = "has-text-warning-30"  // Warning text 30% opacity
	HasTextWarning35  Color = "has-text-warning-35"  // Warning text 35% opacity
	HasTextWarning40  Color = "has-text-warning-40"  // Warning text 40% opacity
	HasTextWarning45  Color = "has-text-warning-45"  // Warning text 45% opacity
	HasTextWarning50  Color = "has-text-warning-50"  // Warning text 50% opacity
	HasTextWarning55  Color = "has-text-warning-55"  // Warning text 55% opacity
	HasTextWarning60  Color = "has-text-warning-60"  // Warning text 60% opacity
	HasTextWarning65  Color = "has-text-warning-65"  // Warning text 65% opacity
	HasTextWarning70  Color = "has-text-warning-70"  // Warning text 70% opacity
	HasTextWarning75  Color = "has-text-warning-75"  // Warning text 75% opacity
	HasTextWarning80  Color = "has-text-warning-80"  // Warning text 80% opacity
	HasTextWarning85  Color = "has-text-warning-85"  // Warning text 85% opacity
	HasTextWarning90  Color = "has-text-warning-90"  // Warning text 90% opacity
	HasTextWarning95  Color = "has-text-warning-95"  // Warning text 95% opacity
	HasTextWarning100 Color = "has-text-warning-100" // Warning text 100% opacity (solid) (transparent) (solid)

	// Text Color Helpers - Warning Opacity Inverts
	HasTextWarning00Invert  Color = "has-text-warning-00-invert"  // Inverted contrast for warning text-00
	HasTextWarning05Invert  Color = "has-text-warning-05-invert"  // Inverted contrast for warning text-05
	HasTextWarning10Invert  Color = "has-text-warning-10-invert"  // Inverted contrast for warning text-10
	HasTextWarning15Invert  Color = "has-text-warning-15-invert"  // Inverted contrast for warning text-15
	HasTextWarning20Invert  Color = "has-text-warning-20-invert"  // Inverted contrast for warning text-20
	HasTextWarning25Invert  Color = "has-text-warning-25-invert"  // Inverted contrast for warning text-25
	HasTextWarning30Invert  Color = "has-text-warning-30-invert"  // Inverted contrast for warning text-30
	HasTextWarning35Invert  Color = "has-text-warning-35-invert"  // Inverted contrast for warning text-35
	HasTextWarning40Invert  Color = "has-text-warning-40-invert"  // Inverted contrast for warning text-40
	HasTextWarning45Invert  Color = "has-text-warning-45-invert"  // Inverted contrast for warning text-45
	HasTextWarning50Invert  Color = "has-text-warning-50-invert"  // Inverted contrast for warning text-50
	HasTextWarning55Invert  Color = "has-text-warning-55-invert"  // Inverted contrast for warning text-55
	HasTextWarning60Invert  Color = "has-text-warning-60-invert"  // Inverted contrast for warning text-60
	HasTextWarning65Invert  Color = "has-text-warning-65-invert"  // Inverted contrast for warning text-65
	HasTextWarning70Invert  Color = "has-text-warning-70-invert"  // Inverted contrast for warning text-70
	HasTextWarning75Invert  Color = "has-text-warning-75-invert"  // Inverted contrast for warning text-75
	HasTextWarning80Invert  Color = "has-text-warning-80-invert"  // Inverted contrast for warning text-80
	HasTextWarning85Invert  Color = "has-text-warning-85-invert"  // Inverted contrast for warning text-85
	HasTextWarning90Invert  Color = "has-text-warning-90-invert"  // Inverted contrast for warning text-90
	HasTextWarning95Invert  Color = "has-text-warning-95-invert"  // Inverted contrast for warning text-95
	HasTextWarning100Invert Color = "has-text-warning-100-invert" // Inverted contrast for warning text-100

	// Text Color Helpers - Warning Semantic Variations
	HasTextWarningInvert      Color = "has-text-warning-invert"       // Warning inverted text color
	HasTextWarningLight       Color = "has-text-warning-light"        // Light variant of warning text
	HasTextWarningDark        Color = "has-text-warning-dark"         // Dark variant of warning text
	HasTextWarningBold        Color = "has-text-warning-bold"         // Bold variant of warning text
	HasTextWarningSoft        Color = "has-text-warning-soft"         // Soft variant of warning text
	HasTextWarningOnScheme    Color = "has-text-warning-on-scheme"    // Warning text adapted for current scheme
	HasTextWarningLightInvert Color = "has-text-warning-light-invert" // Inverted light warning text
	HasTextWarningDarkInvert  Color = "has-text-warning-dark-invert"  // Inverted dark warning text
	HasTextWarningBoldInvert  Color = "has-text-warning-bold-invert"  // Inverted bold warning text
	HasTextWarningSoftInvert  Color = "has-text-warning-soft-invert"  // Inverted soft warning text

	// Text Color Helpers - White Opacity Variations (00-100)
	HasTextWhite00  Color = "has-text-white-00"  // White text 00% opacity (transparent)
	HasTextWhite05  Color = "has-text-white-05"  // White text 05% opacity
	HasTextWhite10  Color = "has-text-white-10"  // White text 10% opacity
	HasTextWhite15  Color = "has-text-white-15"  // White text 15% opacity
	HasTextWhite20  Color = "has-text-white-20"  // White text 20% opacity
	HasTextWhite25  Color = "has-text-white-25"  // White text 25% opacity
	HasTextWhite30  Color = "has-text-white-30"  // White text 30% opacity
	HasTextWhite35  Color = "has-text-white-35"  // White text 35% opacity
	HasTextWhite40  Color = "has-text-white-40"  // White text 40% opacity
	HasTextWhite45  Color = "has-text-white-45"  // White text 45% opacity
	HasTextWhite50  Color = "has-text-white-50"  // White text 50% opacity
	HasTextWhite55  Color = "has-text-white-55"  // White text 55% opacity
	HasTextWhite60  Color = "has-text-white-60"  // White text 60% opacity
	HasTextWhite65  Color = "has-text-white-65"  // White text 65% opacity
	HasTextWhite70  Color = "has-text-white-70"  // White text 70% opacity
	HasTextWhite75  Color = "has-text-white-75"  // White text 75% opacity
	HasTextWhite80  Color = "has-text-white-80"  // White text 80% opacity
	HasTextWhite85  Color = "has-text-white-85"  // White text 85% opacity
	HasTextWhite90  Color = "has-text-white-90"  // White text 90% opacity
	HasTextWhite95  Color = "has-text-white-95"  // White text 95% opacity
	HasTextWhite100 Color = "has-text-white-100" // White text 100% opacity (solid)

	// Text Color Helpers - White Opacity Inverts
	HasTextWhite00Invert  Color = "has-text-white-00-invert"  // Inverted contrast for white text-00
	HasTextWhite05Invert  Color = "has-text-white-05-invert"  // Inverted contrast for white text-05
	HasTextWhite10Invert  Color = "has-text-white-10-invert"  // Inverted contrast for white text-10
	HasTextWhite15Invert  Color = "has-text-white-15-invert"  // Inverted contrast for white text-15
	HasTextWhite20Invert  Color = "has-text-white-20-invert"  // Inverted contrast for white text-20
	HasTextWhite25Invert  Color = "has-text-white-25-invert"  // Inverted contrast for white text-25
	HasTextWhite30Invert  Color = "has-text-white-30-invert"  // Inverted contrast for white text-30
	HasTextWhite35Invert  Color = "has-text-white-35-invert"  // Inverted contrast for white text-35
	HasTextWhite40Invert  Color = "has-text-white-40-invert"  // Inverted contrast for white text-40
	HasTextWhite45Invert  Color = "has-text-white-45-invert"  // Inverted contrast for white text-45
	HasTextWhite50Invert  Color = "has-text-white-50-invert"  // Inverted contrast for white text-50
	HasTextWhite55Invert  Color = "has-text-white-55-invert"  // Inverted contrast for white text-55
	HasTextWhite60Invert  Color = "has-text-white-60-invert"  // Inverted contrast for white text-60
	HasTextWhite65Invert  Color = "has-text-white-65-invert"  // Inverted contrast for white text-65
	HasTextWhite70Invert  Color = "has-text-white-70-invert"  // Inverted contrast for white text-70
	HasTextWhite75Invert  Color = "has-text-white-75-invert"  // Inverted contrast for white text-75
	HasTextWhite80Invert  Color = "has-text-white-80-invert"  // Inverted contrast for white text-80
	HasTextWhite85Invert  Color = "has-text-white-85-invert"  // Inverted contrast for white text-85
	HasTextWhite90Invert  Color = "has-text-white-90-invert"  // Inverted contrast for white text-90
	HasTextWhite95Invert  Color = "has-text-white-95-invert"  // Inverted contrast for white text-95
	HasTextWhite100Invert Color = "has-text-white-100-invert" // Inverted contrast for white text-100

	// Text Color Helpers - White Semantic Variations
	HasTextWhiteInvert      Color = "has-text-white-invert"       // White inverted text color
	HasTextWhiteLight       Color = "has-text-white-light"        // Light variant of white text
	HasTextWhiteDark        Color = "has-text-white-dark"         // Dark variant of white text
	HasTextWhiteBold        Color = "has-text-white-bold"         // Bold variant of white text
	HasTextWhiteSoft        Color = "has-text-white-soft"         // Soft variant of white text
	HasTextWhiteOnScheme    Color = "has-text-white-on-scheme"    // White text adapted for current scheme
	HasTextWhiteLightInvert Color = "has-text-white-light-invert" // Inverted light white text
	HasTextWhiteDarkInvert  Color = "has-text-white-dark-invert"  // Inverted dark white text
	HasTextWhiteBoldInvert  Color = "has-text-white-bold-invert"  // Inverted bold white text
	HasTextWhiteSoftInvert  Color = "has-text-white-soft-invert"  // Inverted soft white text
	HasTextWhiteBis         Color = "has-text-white-bis"          // White-bis monochrome text variant
	HasTextWhiteTer         Color = "has-text-white-ter"          // White-ter monochrome text variant

	// State Modifiers
	IsBlack     Color = "is-black"     // Black color state
	IsDanger    Color = "is-danger"    // Danger color state
	IsDark      Color = "is-dark"      // Dark color state
	IsInfo      Color = "is-info"      // Info color state
	IsLight     Color = "is-light"     // Light color state
	IsLink      Color = "is-link"      // Link color state
	IsPrimary   Color = "is-primary"   // Primary color state
	IsSuccess   Color = "is-success"   // Success color state
	IsText      Color = "is-text"      // Text color state
	IsWarning   Color = "is-warning"   // Warning color state
	IsWhite     Color = "is-white"     // White color state
	IsHoverable Color = "is-hoverable" // Enable hover effects
)
