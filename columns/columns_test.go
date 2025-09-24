package columns

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestColumns(t *testing.T) {
	tests := []struct {
		name   string
		props  ColumnsProps
		expect string
	}{
		{
			name:   "Default",
			props:  ColumnsProps{},
			expect: `<div class="columns"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  ColumnsProps{ID: "main-columns", Class: []string{"custom-grid"}},
			expect: `<div id="main-columns" class="columns custom-grid"></div>`,
		},
		{
			name:   "Mobile enabled",
			props:  ColumnsProps{IsMobile: true},
			expect: `<div class="columns is-mobile"></div>`,
		},
		{
			name:   "Multiline",
			props:  ColumnsProps{IsMultiline: true},
			expect: `<div class="columns is-multiline"></div>`,
		},
		{
			name:   "Desktop only",
			props:  ColumnsProps{IsDesktop: true},
			expect: `<div class="columns is-desktop"></div>`,
		},
		{
			name:   "Gapless",
			props:  ColumnsProps{IsGapless: true},
			expect: `<div class="columns is-gapless"></div>`,
		},
		{
			name:   "Centered",
			props:  ColumnsProps{IsCentered: true},
			expect: `<div class="columns is-centered"></div>`,
		},
		{
			name:   "Vertically centered",
			props:  ColumnsProps{IsVCentered: true},
			expect: `<div class="columns is-vcentered"></div>`,
		},
		{
			name:   "With gap 0",
			props:  ColumnsProps{Gap: Gap0},
			expect: `<div class="columns is-0"></div>`,
		},
		{
			name:   "With gap 3",
			props:  ColumnsProps{Gap: Gap3},
			expect: `<div class="columns is-3"></div>`,
		},
		{
			name:   "With gap 8",
			props:  ColumnsProps{Gap: Gap8},
			expect: `<div class="columns is-8"></div>`,
		},
		{
			name:   "Mobile gap",
			props:  ColumnsProps{GapMobile: Gap2},
			expect: `<div class="columns is-2-mobile"></div>`,
		},
		{
			name:   "Tablet gap",
			props:  ColumnsProps{GapTablet: Gap4},
			expect: `<div class="columns is-4-tablet"></div>`,
		},
		{
			name:   "Desktop gap",
			props:  ColumnsProps{GapDesktop: Gap6},
			expect: `<div class="columns is-6-desktop"></div>`,
		},
		{
			name:   "Widescreen gap",
			props:  ColumnsProps{GapWidescreen: Gap5},
			expect: `<div class="columns is-5-widescreen"></div>`,
		},
		{
			name:   "FullHD gap",
			props:  ColumnsProps{GapFullHD: Gap7},
			expect: `<div class="columns is-7-fullhd"></div>`,
		},
		{
			name:   "Tablet only gap",
			props:  ColumnsProps{GapTabletOnly: Gap1},
			expect: `<div class="columns is-1-tablet-only"></div>`,
		},
		{
			name:   "Touch gap",
			props:  ColumnsProps{GapTouch: Gap3},
			expect: `<div class="columns is-3-touch"></div>`,
		},
		{
			name:   "Desktop only gap",
			props:  ColumnsProps{GapDesktopOnly: Gap4},
			expect: `<div class="columns is-4-desktop-only"></div>`,
		},
		{
			name:   "Widescreen only gap",
			props:  ColumnsProps{GapWidescreenOnly: Gap2},
			expect: `<div class="columns is-2-widescreen-only"></div>`,
		},
		{
			name: "All layout options",
			props: ColumnsProps{
				IsMobile:    true,
				IsMultiline: true,
				IsCentered:  true,
				IsVCentered: true,
			},
			expect: `<div class="columns is-mobile is-multiline is-centered is-vcentered"></div>`,
		},
		{
			name: "Multiple gaps",
			props: ColumnsProps{
				Gap:        Gap3,
				GapMobile:  Gap1,
				GapDesktop: Gap5,
			},
			expect: `<div class="columns is-3 is-1-mobile is-5-desktop"></div>`,
		},
		{
			name: "Complete configuration",
			props: ColumnsProps{
				ID:          "complex-grid",
				Class:       []string{"custom", "responsive"},
				IsMobile:    true,
				IsMultiline: true,
				Gap:         Gap4,
				GapMobile:   Gap2,
				GapTablet:   Gap6,
				Attributes:  templ.Attributes{"data-testid": "columns-container"},
			},
			expect: `<div id="complex-grid" class="columns is-mobile is-multiline is-4 is-2-mobile is-6-tablet custom responsive" data-testid="columns-container"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Columns(tt.props).Render(context.Background(), &buf)
			if err != nil {
				t.Fatalf("render failed: %v", err)
			}
			got := strings.TrimSpace(buf.String())
			if got != tt.expect {
				t.Errorf("expected:\n%s\ngot:\n%s", tt.expect, got)
			}
		})
	}
}

func TestColumnsGaps(t *testing.T) {
	tests := []struct {
		name   string
		props  ColumnsProps
		expect string
	}{
		{"Gap 0", ColumnsProps{Gap: Gap0}, `<div class="columns is-0"></div>`},
		{"Gap 1", ColumnsProps{Gap: Gap1}, `<div class="columns is-1"></div>`},
		{"Gap 2", ColumnsProps{Gap: Gap2}, `<div class="columns is-2"></div>`},
		{"Gap 3", ColumnsProps{Gap: Gap3}, `<div class="columns is-3"></div>`},
		{"Gap 4", ColumnsProps{Gap: Gap4}, `<div class="columns is-4"></div>`},
		{"Gap 5", ColumnsProps{Gap: Gap5}, `<div class="columns is-5"></div>`},
		{"Gap 6", ColumnsProps{Gap: Gap6}, `<div class="columns is-6"></div>`},
		{"Gap 7", ColumnsProps{Gap: Gap7}, `<div class="columns is-7"></div>`},
		{"Gap 8", ColumnsProps{Gap: Gap8}, `<div class="columns is-8"></div>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Columns(tt.props).Render(context.Background(), &buf)
			if err != nil {
				t.Fatalf("render failed: %v", err)
			}
			got := strings.TrimSpace(buf.String())
			if got != tt.expect {
				t.Errorf("expected:\n%s\ngot:\n%s", tt.expect, got)
			}
		})
	}
}

func TestColumn(t *testing.T) {
	tests := []struct {
		name   string
		props  ColumnProps
		expect string
	}{
		{
			name:   "Default",
			props:  ColumnProps{},
			expect: `<div class="column"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  ColumnProps{ID: "main-column", Class: []string{"custom-column"}},
			expect: `<div id="main-column" class="column custom-column"></div>`,
		},
		{
			name:   "Half size",
			props:  ColumnProps{Size: SizeHalf},
			expect: `<div class="column is-half"></div>`,
		},
		{
			name:   "One third size",
			props:  ColumnProps{Size: SizeOneThird},
			expect: `<div class="column is-one-third"></div>`,
		},
		{
			name:   "Two thirds size",
			props:  ColumnProps{Size: SizeTwoThirds},
			expect: `<div class="column is-two-thirds"></div>`,
		},
		{
			name:   "Three quarters size",
			props:  ColumnProps{Size: SizeThreeQuarters},
			expect: `<div class="column is-three-quarters"></div>`,
		},
		{
			name:   "One quarter size",
			props:  ColumnProps{Size: SizeOneQuarter},
			expect: `<div class="column is-one-quarter"></div>`,
		},
		{
			name:   "One fifth size",
			props:  ColumnProps{Size: SizeOneFifth},
			expect: `<div class="column is-one-fifth"></div>`,
		},
		{
			name:   "Two fifths size",
			props:  ColumnProps{Size: SizeTwoFifths},
			expect: `<div class="column is-two-fifths"></div>`,
		},
		{
			name:   "Three fifths size",
			props:  ColumnProps{Size: SizeThreeFifths},
			expect: `<div class="column is-three-fifths"></div>`,
		},
		{
			name:   "Four fifths size",
			props:  ColumnProps{Size: SizeFourFifths},
			expect: `<div class="column is-four-fifths"></div>`,
		},
		{
			name:   "Full size",
			props:  ColumnProps{Size: SizeFull},
			expect: `<div class="column is-full"></div>`,
		},
		{
			name:   "Narrow size",
			props:  ColumnProps{Size: SizeNarrow},
			expect: `<div class="column is-narrow"></div>`,
		},
		{
			name:   "Size 1",
			props:  ColumnProps{Size: Size1},
			expect: `<div class="column is-1"></div>`,
		},
		{
			name:   "Size 6",
			props:  ColumnProps{Size: Size6},
			expect: `<div class="column is-6"></div>`,
		},
		{
			name:   "Size 12",
			props:  ColumnProps{Size: Size12},
			expect: `<div class="column is-12"></div>`,
		},
		{
			name:   "With offset half",
			props:  ColumnProps{Offset: SizeHalf},
			expect: `<div class="column is-offset-half"></div>`,
		},
		{
			name:   "With offset 3",
			props:  ColumnProps{Offset: Size3},
			expect: `<div class="column is-offset-3"></div>`,
		},
		{
			name:   "Size and offset",
			props:  ColumnProps{Size: SizeHalf, Offset: SizeOneQuarter},
			expect: `<div class="column is-half is-offset-one-quarter"></div>`,
		},
		{
			name:   "Mobile size",
			props:  ColumnProps{SizeMobile: SizeFull},
			expect: `<div class="column is-full-mobile"></div>`,
		},
		{
			name:   "Tablet size",
			props:  ColumnProps{SizeTablet: SizeHalf},
			expect: `<div class="column is-half-tablet"></div>`,
		},
		{
			name:   "Touch size",
			props:  ColumnProps{SizeTouch: SizeTwoThirds},
			expect: `<div class="column is-two-thirds-touch"></div>`,
		},
		{
			name:   "Desktop size",
			props:  ColumnProps{SizeDesktop: SizeOneThird},
			expect: `<div class="column is-one-third-desktop"></div>`,
		},
		{
			name:   "Widescreen size",
			props:  ColumnProps{SizeWidescreen: SizeOneQuarter},
			expect: `<div class="column is-one-quarter-widescreen"></div>`,
		},
		{
			name:   "FullHD size",
			props:  ColumnProps{SizeFullHD: Size4},
			expect: `<div class="column is-4-fullhd"></div>`,
		},
		{
			name:   "Mobile offset",
			props:  ColumnProps{OffsetMobile: SizeOneQuarter},
			expect: `<div class="column is-offset-one-quarter-mobile"></div>`,
		},
		{
			name:   "Desktop offset",
			props:  ColumnProps{OffsetDesktop: Size2},
			expect: `<div class="column is-offset-2-desktop"></div>`,
		},
		{
			name: "Responsive sizes",
			props: ColumnProps{
				Size:        Size12,
				SizeMobile:  SizeFull,
				SizeTablet:  SizeHalf,
				SizeDesktop: SizeOneThird,
			},
			expect: `<div class="column is-12 is-full-mobile is-half-tablet is-one-third-desktop"></div>`,
		},
		{
			name: "Complex responsive layout",
			props: ColumnProps{
				Size:          SizeHalf,
				Offset:        SizeOneQuarter,
				SizeMobile:    SizeFull,
				OffsetMobile:  Size0,
				SizeTablet:    SizeTwoThirds,
				SizeDesktop:   SizeOneThird,
				OffsetDesktop: Size1,
			},
			expect: `<div class="column is-half is-offset-one-quarter is-full-mobile is-offset-0-mobile is-two-thirds-tablet is-one-third-desktop is-offset-1-desktop"></div>`,
		},
		{
			name: "All breakpoints",
			props: ColumnProps{
				Size:           Size6,
				SizeMobile:     Size12,
				SizeTablet:     Size8,
				SizeTouch:      Size10,
				SizeDesktop:    Size4,
				SizeWidescreen: Size3,
				SizeFullHD:     Size2,
			},
			expect: `<div class="column is-6 is-12-mobile is-8-tablet is-10-touch is-4-desktop is-3-widescreen is-2-fullhd"></div>`,
		},
		{
			name: "Complete configuration",
			props: ColumnProps{
				ID:           "main-col",
				Class:        []string{"featured", "primary"},
				Size:         SizeHalf,
				Offset:       SizeOneQuarter,
				SizeMobile:   SizeFull,
				OffsetMobile: Size0,
				SizeDesktop:  SizeOneThird,
				Attributes:   templ.Attributes{"data-role": "main-content"},
			},
			expect: `<div id="main-col" class="column is-half is-offset-one-quarter is-full-mobile is-offset-0-mobile is-one-third-desktop featured primary" data-role="main-content"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Column(tt.props).Render(context.Background(), &buf)
			if err != nil {
				t.Fatalf("render failed: %v", err)
			}
			got := strings.TrimSpace(buf.String())
			if got != tt.expect {
				t.Errorf("expected:\n%s\ngot:\n%s", tt.expect, got)
			}
		})
	}
}

func TestColumnSizes(t *testing.T) {
	tests := []struct {
		name   string
		size   Size
		expect string
	}{
		{"Size 0", Size0, `<div class="column is-0"></div>`},
		{"Size 1", Size1, `<div class="column is-1"></div>`},
		{"Size 2", Size2, `<div class="column is-2"></div>`},
		{"Size 3", Size3, `<div class="column is-3"></div>`},
		{"Size 4", Size4, `<div class="column is-4"></div>`},
		{"Size 5", Size5, `<div class="column is-5"></div>`},
		{"Size 6", Size6, `<div class="column is-6"></div>`},
		{"Size 7", Size7, `<div class="column is-7"></div>`},
		{"Size 8", Size8, `<div class="column is-8"></div>`},
		{"Size 9", Size9, `<div class="column is-9"></div>`},
		{"Size 10", Size10, `<div class="column is-10"></div>`},
		{"Size 11", Size11, `<div class="column is-11"></div>`},
		{"Size 12", Size12, `<div class="column is-12"></div>`},
		{"Narrow", SizeNarrow, `<div class="column is-narrow"></div>`},
		{"Full", SizeFull, `<div class="column is-full"></div>`},
		{"Half", SizeHalf, `<div class="column is-half"></div>`},
		{"One third", SizeOneThird, `<div class="column is-one-third"></div>`},
		{"Two thirds", SizeTwoThirds, `<div class="column is-two-thirds"></div>`},
		{"One quarter", SizeOneQuarter, `<div class="column is-one-quarter"></div>`},
		{"Three quarters", SizeThreeQuarters, `<div class="column is-three-quarters"></div>`},
		{"One fifth", SizeOneFifth, `<div class="column is-one-fifth"></div>`},
		{"Two fifths", SizeTwoFifths, `<div class="column is-two-fifths"></div>`},
		{"Three fifths", SizeThreeFifths, `<div class="column is-three-fifths"></div>`},
		{"Four fifths", SizeFourFifths, `<div class="column is-four-fifths"></div>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Column(ColumnProps{Size: tt.size}).Render(context.Background(), &buf)
			if err != nil {
				t.Fatalf("render failed: %v", err)
			}
			got := strings.TrimSpace(buf.String())
			if got != tt.expect {
				t.Errorf("expected:\n%s\ngot:\n%s", tt.expect, got)
			}
		})
	}
}

func TestColumnOffsets(t *testing.T) {
	tests := []struct {
		name   string
		offset Size
		expect string
	}{
		{"Offset 0", Size0, `<div class="column is-offset-0"></div>`},
		{"Offset 1", Size1, `<div class="column is-offset-1"></div>`},
		{"Offset 6", Size6, `<div class="column is-offset-6"></div>`},
		{"Offset 12", Size12, `<div class="column is-offset-12"></div>`},
		{"Offset half", SizeHalf, `<div class="column is-offset-half"></div>`},
		{"Offset one third", SizeOneThird, `<div class="column is-offset-one-third"></div>`},
		{"Offset two thirds", SizeTwoThirds, `<div class="column is-offset-two-thirds"></div>`},
		{"Offset one quarter", SizeOneQuarter, `<div class="column is-offset-one-quarter"></div>`},
		{"Offset three quarters", SizeThreeQuarters, `<div class="column is-offset-three-quarters"></div>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Column(ColumnProps{Offset: tt.offset}).Render(context.Background(), &buf)
			if err != nil {
				t.Fatalf("render failed: %v", err)
			}
			got := strings.TrimSpace(buf.String())
			if got != tt.expect {
				t.Errorf("expected:\n%s\ngot:\n%s", tt.expect, got)
			}
		})
	}
}

func TestColumnResponsive(t *testing.T) {
	tests := []struct {
		name   string
		props  ColumnProps
		expect string
	}{
		// Mobile sizes
		{"Mobile half", ColumnProps{SizeMobile: SizeHalf}, `<div class="column is-half-mobile"></div>`},
		{"Mobile full", ColumnProps{SizeMobile: SizeFull}, `<div class="column is-full-mobile"></div>`},
		{"Mobile 6", ColumnProps{SizeMobile: Size6}, `<div class="column is-6-mobile"></div>`},

		// Tablet sizes
		{"Tablet half", ColumnProps{SizeTablet: SizeHalf}, `<div class="column is-half-tablet"></div>`},
		{"Tablet one third", ColumnProps{SizeTablet: SizeOneThird}, `<div class="column is-one-third-tablet"></div>`},
		{"Tablet 8", ColumnProps{SizeTablet: Size8}, `<div class="column is-8-tablet"></div>`},

		// Touch sizes
		{"Touch two thirds", ColumnProps{SizeTouch: SizeTwoThirds}, `<div class="column is-two-thirds-touch"></div>`},
		{"Touch 9", ColumnProps{SizeTouch: Size9}, `<div class="column is-9-touch"></div>`},

		// Desktop sizes
		{"Desktop one quarter", ColumnProps{SizeDesktop: SizeOneQuarter}, `<div class="column is-one-quarter-desktop"></div>`},
		{"Desktop 3", ColumnProps{SizeDesktop: Size3}, `<div class="column is-3-desktop"></div>`},

		// Widescreen sizes
		{"Widescreen narrow", ColumnProps{SizeWidescreen: SizeNarrow}, `<div class="column is-narrow-widescreen"></div>`},
		{"Widescreen 4", ColumnProps{SizeWidescreen: Size4}, `<div class="column is-4-widescreen"></div>`},

		// FullHD sizes
		{"FullHD two fifths", ColumnProps{SizeFullHD: SizeTwoFifths}, `<div class="column is-two-fifths-fullhd"></div>`},
		{"FullHD 5", ColumnProps{SizeFullHD: Size5}, `<div class="column is-5-fullhd"></div>`},

		// Responsive offsets
		{"Mobile offset", ColumnProps{OffsetMobile: SizeOneQuarter}, `<div class="column is-offset-one-quarter-mobile"></div>`},
		{"Tablet offset", ColumnProps{OffsetTablet: Size2}, `<div class="column is-offset-2-tablet"></div>`},
		{"Desktop offset", ColumnProps{OffsetDesktop: SizeHalf}, `<div class="column is-offset-half-desktop"></div>`},
		{"Widescreen offset", ColumnProps{OffsetWidescreen: Size3}, `<div class="column is-offset-3-widescreen"></div>`},
		{"FullHD offset", ColumnProps{OffsetFullHD: SizeOneThird}, `<div class="column is-offset-one-third-fullhd"></div>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Column(tt.props).Render(context.Background(), &buf)
			if err != nil {
				t.Fatalf("render failed: %v", err)
			}
			got := strings.TrimSpace(buf.String())
			if got != tt.expect {
				t.Errorf("expected:\n%s\ngot:\n%s", tt.expect, got)
			}
		})
	}
}

func TestGapValidation(t *testing.T) {
	tests := []struct {
		name   string
		gap    Gap
		expect string
	}{
		{"Gap 0 valid", Gap0, `<div class="columns is-0"></div>`},
		{"Gap 1 valid", Gap1, `<div class="columns is-1"></div>`},
		{"Gap 4 valid", Gap4, `<div class="columns is-4"></div>`},
		{"Gap 8 valid", Gap8, `<div class="columns is-8"></div>`},
		{"Gap nil (unset)", nil, `<div class="columns"></div>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Columns(ColumnsProps{Gap: tt.gap}).Render(context.Background(), &buf)
			if err != nil {
				t.Fatalf("render failed: %v", err)
			}
			got := strings.TrimSpace(buf.String())
			if got != tt.expect {
				t.Errorf("expected:\n%s\ngot:\n%s", tt.expect, got)
			}
		})
	}
}

func TestFormatGapClass(t *testing.T) {
	tests := []struct {
		name        string
		gap         Gap
		suffix      string
		expectClass string
		expectValid bool
	}{
		{"Valid gap 0", Gap0, "", "is-0", true},
		{"Valid gap 3", Gap3, "", "is-3", true},
		{"Valid gap 8", Gap8, "", "is-8", true},
		{"Valid gap with mobile suffix", Gap2, "mobile", "is-2-mobile", true},
		{"Valid gap with tablet suffix", Gap5, "tablet", "is-5-tablet", true},
		{"Valid gap with desktop suffix", Gap7, "desktop", "is-7-desktop", true},
		{"Nil gap", nil, "", "", false},
		{"Nil gap with suffix", nil, "mobile", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotClass, gotValid := formatGapClass(tt.gap, tt.suffix)
			if gotClass != tt.expectClass {
				t.Errorf("expected class %q, got %q", tt.expectClass, gotClass)
			}
			if gotValid != tt.expectValid {
				t.Errorf("expected valid %v, got %v", tt.expectValid, gotValid)
			}
		})
	}
}

func TestFormatSizeClass(t *testing.T) {
	tests := []struct {
		name        string
		prefix      string
		size        Size
		suffix      string
		expectClass string
		expectValid bool
	}{
		{"Valid size class", "is", SizeHalf, "", "is-half", true},
		{"Valid offset class", "is-offset", SizeOneThird, "", "is-offset-one-third", true},
		{"Valid responsive size", "is", Size6, "mobile", "is-6-mobile", true},
		{"Valid responsive offset", "is-offset", SizeOneQuarter, "desktop", "is-offset-one-quarter-desktop", true},
		{"Empty size", "is", Size(""), "", "", false},
		{"Empty size with suffix", "is", Size(""), "mobile", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotClass, gotValid := formatSizeClass(tt.prefix, tt.size, tt.suffix)
			if gotClass != tt.expectClass {
				t.Errorf("expected class %q, got %q", tt.expectClass, gotClass)
			}
			if gotValid != tt.expectValid {
				t.Errorf("expected valid %v, got %v", tt.expectValid, gotValid)
			}
		})
	}
}

func TestColumnsResponsiveGaps(t *testing.T) {
	tests := []struct {
		name   string
		props  ColumnsProps
		expect string
	}{
		{
			name: "All responsive gaps",
			props: ColumnsProps{
				Gap:               Gap4,
				GapMobile:         Gap1,
				GapTablet:         Gap2,
				GapTabletOnly:     Gap3,
				GapTouch:          Gap2,
				GapDesktop:        Gap5,
				GapDesktopOnly:    Gap6,
				GapWidescreen:     Gap7,
				GapWidescreenOnly: Gap4,
				GapFullHD:         Gap8,
			},
			expect: `<div class="columns is-4 is-1-mobile is-2-tablet is-3-tablet-only is-2-touch is-5-desktop is-6-desktop-only is-7-widescreen is-4-widescreen-only is-8-fullhd"></div>`,
		},
		{
			name: "Mixed layout with gaps",
			props: ColumnsProps{
				IsMobile:    true,
				IsMultiline: true,
				IsCentered:  true,
				Gap:         Gap3,
				GapMobile:   Gap1,
				GapDesktop:  Gap5,
			},
			expect: `<div class="columns is-mobile is-multiline is-centered is-3 is-1-mobile is-5-desktop"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Columns(tt.props).Render(context.Background(), &buf)
			if err != nil {
				t.Fatalf("render failed: %v", err)
			}
			got := strings.TrimSpace(buf.String())
			if got != tt.expect {
				t.Errorf("expected:\n%s\ngot:\n%s", tt.expect, got)
			}
		})
	}
}
