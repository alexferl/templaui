package helpers

// Flexbox represents Bulma's comprehensive flexbox helper system.
//
// Use these constants to apply Bulma's flexbox helpers which provide
// complete control over CSS flexbox properties. Must be combined with
// the "is-flex" class to enable flexbox behavior. These helpers cover
// all major flexbox properties: flex-direction, flex-wrap, justify-content,
// align-content, align-items, align-self, flex-grow, and flex-shrink.
type Flexbox string

const (
	// Align Content Helpers - Control line spacing in multi-line flex containers
	IsAlignContentBaseline     Flexbox = "is-align-content-baseline"      // Align lines to text baseline
	IsAlignContentCenter       Flexbox = "is-align-content-center"        // Center lines vertically
	IsAlignContentEnd          Flexbox = "is-align-content-end"           // Pack lines to end of container
	IsAlignContentFlexEnd      Flexbox = "is-align-content-flex-end"      // Pack lines to flex end (legacy)
	IsAlignContentFlexStart    Flexbox = "is-align-content-flex-start"    // Pack lines to flex start (legacy)
	IsAlignContentSpaceAround  Flexbox = "is-align-content-space-around"  // Distribute lines with space around each
	IsAlignContentSpaceBetween Flexbox = "is-align-content-space-between" // Distribute lines with space between
	IsAlignContentSpaceEvenly  Flexbox = "is-align-content-space-evenly"  // Distribute lines with equal space
	IsAlignContentStart        Flexbox = "is-align-content-start"         // Pack lines to start of container
	IsAlignContentStretch      Flexbox = "is-align-content-stretch"       // Stretch lines to fill container height

	// Align Items Helpers - Control cross-axis alignment of flex items
	IsAlignItemsBaseline  Flexbox = "is-align-items-baseline"   // Align items to text baseline
	IsAlignItemsCenter    Flexbox = "is-align-items-center"     // Center items on cross-axis
	IsAlignItemsEnd       Flexbox = "is-align-items-end"        // Align items to cross-axis end
	IsAlignItemsFlexEnd   Flexbox = "is-align-items-flex-end"   // Align items to flex end (legacy)
	IsAlignItemsFlexStart Flexbox = "is-align-items-flex-start" // Align items to flex start (legacy)
	IsAlignItemsSelfEnd   Flexbox = "is-align-items-self-end"   // Align items to self end
	IsAlignItemsSelfStart Flexbox = "is-align-items-self-start" // Align items to self start
	IsAlignItemsStart     Flexbox = "is-align-items-start"      // Align items to cross-axis start
	IsAlignItemsStretch   Flexbox = "is-align-items-stretch"    // Stretch items to fill cross-axis

	// Align Self Helpers - Override container alignment for individual items
	IsAlignSelfAuto      Flexbox = "is-align-self-auto"       // Use container's align-items value
	IsAlignSelfBaseline  Flexbox = "is-align-self-baseline"   // Align item to text baseline
	IsAlignSelfCenter    Flexbox = "is-align-self-center"     // Center item on cross-axis
	IsAlignSelfFlexEnd   Flexbox = "is-align-self-flex-end"   // Align item to flex end
	IsAlignSelfFlexStart Flexbox = "is-align-self-flex-start" // Align item to flex start
	IsAlignSelfStretch   Flexbox = "is-align-self-stretch"    // Stretch item to fill cross-axis

	// Flex Direction Helpers - Control main axis direction and order
	IsFlexDirectionColumn        Flexbox = "is-flex-direction-column"         // Stack items vertically (top to bottom)
	IsFlexDirectionColumnReverse Flexbox = "is-flex-direction-column-reverse" // Stack items vertically (bottom to top)
	IsFlexDirectionRow           Flexbox = "is-flex-direction-row"            // Arrange items horizontally (left to right)
	IsFlexDirectionRowReverse    Flexbox = "is-flex-direction-row-reverse"    // Arrange items horizontally (right to left)

	// Flex Grow Helpers - Control how items expand to fill space
	IsFlexGrow0 Flexbox = "is-flex-grow-0" // Item will not grow (flex-grow: 0)
	IsFlexGrow1 Flexbox = "is-flex-grow-1" // Item will grow 1x rate (flex-grow: 1)
	IsFlexGrow2 Flexbox = "is-flex-grow-2" // Item will grow 2x rate (flex-grow: 2)
	IsFlexGrow3 Flexbox = "is-flex-grow-3" // Item will grow 3x rate (flex-grow: 3)
	IsFlexGrow4 Flexbox = "is-flex-grow-4" // Item will grow 4x rate (flex-grow: 4)
	IsFlexGrow5 Flexbox = "is-flex-grow-5" // Item will grow 5x rate (flex-grow: 5)

	// Flex Shrink Helpers - Control how items contract when space is limited
	IsFlexShrink0 Flexbox = "is-flex-shrink-0" // Item will not shrink (flex-shrink: 0)
	IsFlexShrink1 Flexbox = "is-flex-shrink-1" // Item will shrink 1x rate (flex-shrink: 1)
	IsFlexShrink2 Flexbox = "is-flex-shrink-2" // Item will shrink 2x rate (flex-shrink: 2)
	IsFlexShrink3 Flexbox = "is-flex-shrink-3" // Item will shrink 3x rate (flex-shrink: 3)
	IsFlexShrink4 Flexbox = "is-flex-shrink-4" // Item will shrink 4x rate (flex-shrink: 4)
	IsFlexShrink5 Flexbox = "is-flex-shrink-5" // Item will shrink 5x rate (flex-shrink: 5)

	// Flex Wrap Helpers - Control wrapping behavior of flex items
	IsFlexWrapNowrap      Flexbox = "is-flex-wrap-nowrap"       // Items stay on single line (may overflow)
	IsFlexWrapWrap        Flexbox = "is-flex-wrap-wrap"         // Items wrap to new lines as needed
	IsFlexWrapWrapReverse Flexbox = "is-flex-wrap-wrap-reverse" // Items wrap to new lines in reverse order

	// Justify Content Helpers - Control main-axis alignment and distribution
	IsJustifyContentCenter       Flexbox = "is-justify-content-center"        // Center items on main axis
	IsJustifyContentEnd          Flexbox = "is-justify-content-end"           // Pack items to main axis end
	IsJustifyContentFlexEnd      Flexbox = "is-justify-content-flex-end"      // Pack items to flex end (legacy)
	IsJustifyContentFlexStart    Flexbox = "is-justify-content-flex-start"    // Pack items to flex start (legacy)
	IsJustifyContentLeft         Flexbox = "is-justify-content-left"          // Pack items to left edge
	IsJustifyContentRight        Flexbox = "is-justify-content-right"         // Pack items to right edge
	IsJustifyContentSpaceAround  Flexbox = "is-justify-content-space-around"  // Distribute items with space around each
	IsJustifyContentSpaceBetween Flexbox = "is-justify-content-space-between" // Distribute items with space between
	IsJustifyContentSpaceEvenly  Flexbox = "is-justify-content-space-evenly"  // Distribute items with equal space
	IsJustifyContentStart        Flexbox = "is-justify-content-start"         // Pack items to main axis start
)
