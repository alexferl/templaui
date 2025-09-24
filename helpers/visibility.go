package helpers

// Visibility represents Bulma's comprehensive responsive visibility helper system.
//
// Use these constants to control element display and visibility across Bulma's
// breakpoint system. Includes display property helpers (block, flex, grid, inline),
// visibility helpers (hidden, invisible), and accessibility helpers (sr-only).
// All helpers support responsive breakpoints: mobile (≤768px), tablet (769-1023px),
// desktop (1024-1215px), widescreen (1216-1407px), fullhd (≥1408px).
type Visibility string

const (
	// Legacy Display Helpers - Simplified display control (use is-display-* for consistency)
	IsBlock               Visibility = "is-block"                 // Display as block element
	IsBlockDesktop        Visibility = "is-block-desktop"         // Block on desktop and larger
	IsBlockDesktopOnly    Visibility = "is-block-desktop-only"    // Block on desktop only
	IsBlockFullHD         Visibility = "is-block-fullhd"          // Block on fullhd screens
	IsBlockMobile         Visibility = "is-block-mobile"          // Block on mobile only
	IsBlockTablet         Visibility = "is-block-tablet"          // Block on tablet and larger
	IsBlockTabletOnly     Visibility = "is-block-tablet-only"     // Block on tablet only
	IsBlockTouch          Visibility = "is-block-touch"           // Block on mobile and tablet
	IsBlockWidescreen     Visibility = "is-block-widescreen"      // Block on widescreen and larger
	IsBlockWidescreenOnly Visibility = "is-block-widescreen-only" // Block on widescreen only

	IsFlex               Visibility = "is-flex"                 // Display as flexbox container
	IsFlexDesktop        Visibility = "is-flex-desktop"         // Flex on desktop and larger
	IsFlexDesktopOnly    Visibility = "is-flex-desktop-only"    // Flex on desktop only
	IsFlexFullHD         Visibility = "is-flex-fullhd"          // Flex on fullhd screens
	IsFlexMobile         Visibility = "is-flex-mobile"          // Flex on mobile only
	IsFlexTablet         Visibility = "is-flex-tablet"          // Flex on tablet and larger
	IsFlexTabletOnly     Visibility = "is-flex-tablet-only"     // Flex on tablet only
	IsFlexTouch          Visibility = "is-flex-touch"           // Flex on mobile and tablet
	IsFlexWidescreen     Visibility = "is-flex-widescreen"      // Flex on widescreen and larger
	IsFlexWidescreenOnly Visibility = "is-flex-widescreen-only" // Flex on widescreen only

	IsInline               Visibility = "is-inline"                 // Display as inline element
	IsInlineDesktop        Visibility = "is-inline-desktop"         // Inline on desktop and larger
	IsInlineDesktopOnly    Visibility = "is-inline-desktop-only"    // Inline on desktop only
	IsInlineFullHD         Visibility = "is-inline-fullhd"          // Inline on fullhd screens
	IsInlineMobile         Visibility = "is-inline-mobile"          // Inline on mobile only
	IsInlineTablet         Visibility = "is-inline-tablet"          // Inline on tablet and larger
	IsInlineTabletOnly     Visibility = "is-inline-tablet-only"     // Inline on tablet only
	IsInlineTouch          Visibility = "is-inline-touch"           // Inline on mobile and tablet
	IsInlineWidescreen     Visibility = "is-inline-widescreen"      // Inline on widescreen and larger
	IsInlineWidescreenOnly Visibility = "is-inline-widescreen-only" // Inline on widescreen only

	IsInlineBlock               Visibility = "is-inline-block"                 // Display as inline-block element
	IsInlineBlockDesktop        Visibility = "is-inline-block-desktop"         // Inline-block on desktop and larger
	IsInlineBlockDesktopOnly    Visibility = "is-inline-block-desktop-only"    // Inline-block on desktop only
	IsInlineBlockFullHD         Visibility = "is-inline-block-fullhd"          // Inline-block on fullhd screens
	IsInlineBlockMobile         Visibility = "is-inline-block-mobile"          // Inline-block on mobile only
	IsInlineBlockTablet         Visibility = "is-inline-block-tablet"          // Inline-block on tablet and larger
	IsInlineBlockTabletOnly     Visibility = "is-inline-block-tablet-only"     // Inline-block on tablet only
	IsInlineBlockTouch          Visibility = "is-inline-block-touch"           // Inline-block on mobile and tablet
	IsInlineBlockWidescreen     Visibility = "is-inline-block-widescreen"      // Inline-block on widescreen and larger
	IsInlineBlockWidescreenOnly Visibility = "is-inline-block-widescreen-only" // Inline-block on widescreen only

	IsInlineFlex               Visibility = "is-inline-flex"                 // Display as inline-flex container
	IsInlineFlexDesktop        Visibility = "is-inline-flex-desktop"         // Inline-flex on desktop and larger
	IsInlineFlexDesktopOnly    Visibility = "is-inline-flex-desktop-only"    // Inline-flex on desktop only
	IsInlineFlexFullHD         Visibility = "is-inline-flex-fullhd"          // Inline-flex on fullhd screens
	IsInlineFlexMobile         Visibility = "is-inline-flex-mobile"          // Inline-flex on mobile only
	IsInlineFlexTablet         Visibility = "is-inline-flex-tablet"          // Inline-flex on tablet and larger
	IsInlineFlexTabletOnly     Visibility = "is-inline-flex-tablet-only"     // Inline-flex on tablet only
	IsInlineFlexTouch          Visibility = "is-inline-flex-touch"           // Inline-flex on mobile and tablet
	IsInlineFlexWidescreen     Visibility = "is-inline-flex-widescreen"      // Inline-flex on widescreen and larger
	IsInlineFlexWidescreenOnly Visibility = "is-inline-flex-widescreen-only" // Inline-flex on widescreen only

	IsGrid               Visibility = "is-grid"                 // Display as CSS Grid container
	IsGridDesktop        Visibility = "is-grid-desktop"         // Grid on desktop and larger
	IsGridDesktopOnly    Visibility = "is-grid-desktop-only"    // Grid on desktop only
	IsGridFullHD         Visibility = "is-grid-fullhd"          // Grid on fullhd screens
	IsGridMobile         Visibility = "is-grid-mobile"          // Grid on mobile only
	IsGridTablet         Visibility = "is-grid-tablet"          // Grid on tablet and larger
	IsGridTabletOnly     Visibility = "is-grid-tablet-only"     // Grid on tablet only
	IsGridTouch          Visibility = "is-grid-touch"           // Grid on mobile and tablet
	IsGridWidescreen     Visibility = "is-grid-widescreen"      // Grid on widescreen and larger
	IsGridWidescreenOnly Visibility = "is-grid-widescreen-only" // Grid on widescreen only

	// Explicit Display Helpers - Preferred for consistency
	IsDisplayBlock               Visibility = "is-display-block"                 // Explicit display: block
	IsDisplayBlockDesktop        Visibility = "is-display-block-desktop"         // Block on desktop and larger
	IsDisplayBlockDesktopOnly    Visibility = "is-display-block-desktop-only"    // Block on desktop only
	IsDisplayBlockFullHD         Visibility = "is-display-block-fullhd"          // Block on fullhd screens
	IsDisplayBlockMobile         Visibility = "is-display-block-mobile"          // Block on mobile only
	IsDisplayBlockTablet         Visibility = "is-display-block-tablet"          // Block on tablet and larger
	IsDisplayBlockTabletOnly     Visibility = "is-display-block-tablet-only"     // Block on tablet only
	IsDisplayBlockTouch          Visibility = "is-display-block-touch"           // Block on mobile and tablet
	IsDisplayBlockWidescreen     Visibility = "is-display-block-widescreen"      // Block on widescreen and larger
	IsDisplayBlockWidescreenOnly Visibility = "is-display-block-widescreen-only" // Block on widescreen only

	IsDisplayFlex               Visibility = "is-display-flex"                 // Explicit display: flex
	IsDisplayFlexDesktop        Visibility = "is-display-flex-desktop"         // Flex on desktop and larger
	IsDisplayFlexDesktopOnly    Visibility = "is-display-flex-desktop-only"    // Flex on desktop only
	IsDisplayFlexFullHD         Visibility = "is-display-flex-fullhd"          // Flex on fullhd screens
	IsDisplayFlexMobile         Visibility = "is-display-flex-mobile"          // Flex on mobile only
	IsDisplayFlexTablet         Visibility = "is-display-flex-tablet"          // Flex on tablet and larger
	IsDisplayFlexTabletOnly     Visibility = "is-display-flex-tablet-only"     // Flex on tablet only
	IsDisplayFlexTouch          Visibility = "is-display-flex-touch"           // Flex on mobile and tablet
	IsDisplayFlexWidescreen     Visibility = "is-display-flex-widescreen"      // Flex on widescreen and larger
	IsDisplayFlexWidescreenOnly Visibility = "is-display-flex-widescreen-only" // Flex on widescreen only

	IsDisplayGrid               Visibility = "is-display-grid"                 // Explicit display: grid
	IsDisplayGridDesktop        Visibility = "is-display-grid-desktop"         // Grid on desktop and larger
	IsDisplayGridDesktopOnly    Visibility = "is-display-grid-desktop-only"    // Grid on desktop only
	IsDisplayGridFullHD         Visibility = "is-display-grid-fullhd"          // Grid on fullhd screens
	IsDisplayGridMobile         Visibility = "is-display-grid-mobile"          // Grid on mobile only
	IsDisplayGridTablet         Visibility = "is-display-grid-tablet"          // Grid on tablet and larger
	IsDisplayGridTabletOnly     Visibility = "is-display-grid-tablet-only"     // Grid on tablet only
	IsDisplayGridTouch          Visibility = "is-display-grid-touch"           // Grid on mobile and tablet
	IsDisplayGridWidescreen     Visibility = "is-display-grid-widescreen"      // Grid on widescreen and larger
	IsDisplayGridWidescreenOnly Visibility = "is-display-grid-widescreen-only" // Grid on widescreen only

	IsDisplayInline               Visibility = "is-display-inline"                 // Explicit display: inline
	IsDisplayInlineDesktop        Visibility = "is-display-inline-desktop"         // Inline on desktop and larger
	IsDisplayInlineDesktopOnly    Visibility = "is-display-inline-desktop-only"    // Inline on desktop only
	IsDisplayInlineFullHD         Visibility = "is-display-inline-fullhd"          // Inline on fullhd screens
	IsDisplayInlineMobile         Visibility = "is-display-inline-mobile"          // Inline on mobile only
	IsDisplayInlineTablet         Visibility = "is-display-inline-tablet"          // Inline on tablet and larger
	IsDisplayInlineTabletOnly     Visibility = "is-display-inline-tablet-only"     // Inline on tablet only
	IsDisplayInlineTouch          Visibility = "is-display-inline-touch"           // Inline on mobile and tablet
	IsDisplayInlineWidescreen     Visibility = "is-display-inline-widescreen"      // Inline on widescreen and larger
	IsDisplayInlineWidescreenOnly Visibility = "is-display-inline-widescreen-only" // Inline on widescreen only

	IsDisplayInlineBlock               Visibility = "is-display-inline-block"                 // Explicit display: inline-block
	IsDisplayInlineBlockDesktop        Visibility = "is-display-inline-block-desktop"         // Inline-block on desktop and larger
	IsDisplayInlineBlockDesktopOnly    Visibility = "is-display-inline-block-desktop-only"    // Inline-block on desktop only
	IsDisplayInlineBlockFullHD         Visibility = "is-display-inline-block-fullhd"          // Inline-block on fullhd screens
	IsDisplayInlineBlockMobile         Visibility = "is-display-inline-block-mobile"          // Inline-block on mobile only
	IsDisplayInlineBlockTablet         Visibility = "is-display-inline-block-tablet"          // Inline-block on tablet and larger
	IsDisplayInlineBlockTabletOnly     Visibility = "is-display-inline-block-tablet-only"     // Inline-block on tablet only
	IsDisplayInlineBlockTouch          Visibility = "is-display-inline-block-touch"           // Inline-block on mobile and tablet
	IsDisplayInlineBlockWidescreen     Visibility = "is-display-inline-block-widescreen"      // Inline-block on widescreen and larger
	IsDisplayInlineBlockWidescreenOnly Visibility = "is-display-inline-block-widescreen-only" // Inline-block on widescreen only

	IsDisplayInlineFlex               Visibility = "is-display-inline-flex"                 // Explicit display: inline-flex
	IsDisplayInlineFlexDesktop        Visibility = "is-display-inline-flex-desktop"         // Inline-flex on desktop and larger
	IsDisplayInlineFlexDesktopOnly    Visibility = "is-display-inline-flex-desktop-only"    // Inline-flex on desktop only
	IsDisplayInlineFlexFullHD         Visibility = "is-display-inline-flex-fullhd"          // Inline-flex on fullhd screens
	IsDisplayInlineFlexMobile         Visibility = "is-display-inline-flex-mobile"          // Inline-flex on mobile only
	IsDisplayInlineFlexTablet         Visibility = "is-display-inline-flex-tablet"          // Inline-flex on tablet and larger
	IsDisplayInlineFlexTabletOnly     Visibility = "is-display-inline-flex-tablet-only"     // Inline-flex on tablet only
	IsDisplayInlineFlexTouch          Visibility = "is-display-inline-flex-touch"           // Inline-flex on mobile and tablet
	IsDisplayInlineFlexWidescreen     Visibility = "is-display-inline-flex-widescreen"      // Inline-flex on widescreen and larger
	IsDisplayInlineFlexWidescreenOnly Visibility = "is-display-inline-flex-widescreen-only" // Inline-flex on widescreen only

	IsDisplayNone               Visibility = "is-display-none"                 // Explicit display: none (hide element)
	IsDisplayNoneDesktop        Visibility = "is-display-none-desktop"         // Hide on desktop and larger
	IsDisplayNoneDesktopOnly    Visibility = "is-display-none-desktop-only"    // Hide on desktop only
	IsDisplayNoneFullHD         Visibility = "is-display-none-fullhd"          // Hide on fullhd screens
	IsDisplayNoneMobile         Visibility = "is-display-none-mobile"          // Hide on mobile only
	IsDisplayNoneTablet         Visibility = "is-display-none-tablet"          // Hide on tablet and larger
	IsDisplayNoneTabletOnly     Visibility = "is-display-none-tablet-only"     // Hide on tablet only
	IsDisplayNoneTouch          Visibility = "is-display-none-touch"           // Hide on mobile and tablet
	IsDisplayNoneWidescreen     Visibility = "is-display-none-widescreen"      // Hide on widescreen and larger
	IsDisplayNoneWidescreenOnly Visibility = "is-display-none-widescreen-only" // Hide on widescreen only

	// Visibility Helpers - Control element visibility without affecting layout
	IsHidden               Visibility = "is-hidden"                 // Hide element (display: none)
	IsHiddenDesktop        Visibility = "is-hidden-desktop"         // Hide on desktop and larger
	IsHiddenDesktopOnly    Visibility = "is-hidden-desktop-only"    // Hide on desktop only
	IsHiddenFullHD         Visibility = "is-hidden-fullhd"          // Hide on fullhd screens
	IsHiddenMobile         Visibility = "is-hidden-mobile"          // Hide on mobile only
	IsHiddenTablet         Visibility = "is-hidden-tablet"          // Hide on tablet and larger
	IsHiddenTabletOnly     Visibility = "is-hidden-tablet-only"     // Hide on tablet only
	IsHiddenTouch          Visibility = "is-hidden-touch"           // Hide on mobile and tablet
	IsHiddenWidescreen     Visibility = "is-hidden-widescreen"      // Hide on widescreen and larger
	IsHiddenWidescreenOnly Visibility = "is-hidden-widescreen-only" // Hide on widescreen only

	IsInvisible               Visibility = "is-invisible"                 // Make invisible but preserve layout space
	IsInvisibleDesktop        Visibility = "is-invisible-desktop"         // Invisible on desktop and larger
	IsInvisibleDesktopOnly    Visibility = "is-invisible-desktop-only"    // Invisible on desktop only
	IsInvisibleFullHD         Visibility = "is-invisible-fullhd"          // Invisible on fullhd screens
	IsInvisibleMobile         Visibility = "is-invisible-mobile"          // Invisible on mobile only
	IsInvisibleTablet         Visibility = "is-invisible-tablet"          // Invisible on tablet and larger
	IsInvisibleTabletOnly     Visibility = "is-invisible-tablet-only"     // Invisible on tablet only
	IsInvisibleTouch          Visibility = "is-invisible-touch"           // Invisible on mobile and tablet
	IsInvisibleWidescreen     Visibility = "is-invisible-widescreen"      // Invisible on widescreen and larger
	IsInvisibleWidescreenOnly Visibility = "is-invisible-widescreen-only" // Invisible on widescreen only

	IsVisibilityHidden               Visibility = "is-visibility-hidden"                 // Explicit visibility: hidden
	IsVisibilityHiddenDesktop        Visibility = "is-visibility-hidden-desktop"         // Visibility hidden on desktop and larger
	IsVisibilityHiddenDesktopOnly    Visibility = "is-visibility-hidden-desktop-only"    // Visibility hidden on desktop only
	IsVisibilityHiddenFullHD         Visibility = "is-visibility-hidden-fullhd"          // Visibility hidden on fullhd screens
	IsVisibilityHiddenMobile         Visibility = "is-visibility-hidden-mobile"          // Visibility hidden on mobile only
	IsVisibilityHiddenTablet         Visibility = "is-visibility-hidden-tablet"          // Visibility hidden on tablet and larger
	IsVisibilityHiddenTabletOnly     Visibility = "is-visibility-hidden-tablet-only"     // Visibility hidden on tablet only
	IsVisibilityHiddenTouch          Visibility = "is-visibility-hidden-touch"           // Visibility hidden on mobile and tablet
	IsVisibilityHiddenWidescreen     Visibility = "is-visibility-hidden-widescreen"      // Visibility hidden on widescreen and larger
	IsVisibilityHiddenWidescreenOnly Visibility = "is-visibility-hidden-widescreen-only" // Visibility hidden on widescreen only

	// Accessibility Helpers
	IsSrOnly Visibility = "is-sr-only" // Screen reader only (visually hidden but accessible)
)
