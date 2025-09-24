package helpers

// AspectRatio represents CSS aspect ratio helper modifiers.
//
// Use these constants to apply Bulma's aspect ratio helpers which
// maintain specific width-to-height proportions for responsive
// containers. Perfect for videos, images, cards, and any content
// requiring consistent proportional scaling across devices.
// Applied to containers that need to preserve aspect ratios.
type AspectRatio string

const (
	IsAspectRatio16by9 AspectRatio = "is-aspect-ratio-16by9" // 16:9 widescreen ratio (1.78:1) - standard video/TV
	IsAspectRatio1by1  AspectRatio = "is-aspect-ratio-1by1"  // 1:1 square ratio - profile pictures, thumbnails
	IsAspectRatio1by2  AspectRatio = "is-aspect-ratio-1by2"  // 1:2 tall rectangle ratio (0.5:1) - mobile layouts
	IsAspectRatio1by3  AspectRatio = "is-aspect-ratio-1by3"  // 1:3 very tall ratio (0.33:1) - sidebar content
	IsAspectRatio2by1  AspectRatio = "is-aspect-ratio-2by1"  // 2:1 wide rectangle ratio - banner images
	IsAspectRatio2by3  AspectRatio = "is-aspect-ratio-2by3"  // 2:3 portrait ratio (0.67:1) - book covers, posters
	IsAspectRatio3by1  AspectRatio = "is-aspect-ratio-3by1"  // 3:1 very wide ratio - header banners
	IsAspectRatio3by2  AspectRatio = "is-aspect-ratio-3by2"  // 3:2 classic photo ratio (1.5:1) - 35mm film standard
	IsAspectRatio3by4  AspectRatio = "is-aspect-ratio-3by4"  // 3:4 portrait ratio (0.75:1) - traditional portraits
	IsAspectRatio3by5  AspectRatio = "is-aspect-ratio-3by5"  // 3:5 tall portrait ratio (0.6:1) - magazine layouts
	IsAspectRatio4by3  AspectRatio = "is-aspect-ratio-4by3"  // 4:3 classic screen ratio (1.33:1) - old TV/monitor standard
	IsAspectRatio4by5  AspectRatio = "is-aspect-ratio-4by5"  // 4:5 Instagram portrait ratio (0.8:1) - social media
	IsAspectRatio5by3  AspectRatio = "is-aspect-ratio-5by3"  // 5:3 wide ratio (1.67:1) - panoramic views
	IsAspectRatio5by4  AspectRatio = "is-aspect-ratio-5by4"  // 5:4 slightly wide ratio (1.25:1) - large format photos
	IsAspectRatio9by16 AspectRatio = "is-aspect-ratio-9by16" // 9:16 vertical video ratio (0.56:1) - mobile video, stories
)
