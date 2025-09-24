package grid

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestFixedGrid(t *testing.T) {
	tests := []struct {
		name   string
		props  FixedGridProps
		expect string
	}{
		{
			name:   "Default",
			props:  FixedGridProps{},
			expect: `<div class="fixed-grid"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  FixedGridProps{ID: "grid1", Class: []string{"custom-grid"}},
			expect: `<div id="grid1" class="fixed-grid custom-grid"></div>`,
		},
		{
			name:   "3 columns",
			props:  FixedGridProps{Cols3: true},
			expect: `<div class="fixed-grid has-3-cols"></div>`,
		},
		{
			name:   "12 columns",
			props:  FixedGridProps{Cols12: true},
			expect: `<div class="fixed-grid has-12-cols"></div>`,
		},
		{
			name:   "With mobile responsive",
			props:  FixedGridProps{Cols3: true, Cols1Mobile: true},
			expect: `<div class="fixed-grid has-3-cols has-1-cols-mobile"></div>`,
		},
		{
			name:   "With tablet responsive",
			props:  FixedGridProps{Cols4: true, Cols2Tablet: true},
			expect: `<div class="fixed-grid has-4-cols has-2-cols-tablet"></div>`,
		},
		{
			name:   "With desktop responsive",
			props:  FixedGridProps{Cols6: true, Cols4Desktop: true},
			expect: `<div class="fixed-grid has-6-cols has-4-cols-desktop"></div>`,
		},
		{
			name:   "With widescreen responsive",
			props:  FixedGridProps{Cols8: true, Cols6Widescreen: true},
			expect: `<div class="fixed-grid has-8-cols has-6-cols-widescreen"></div>`,
		},
		{
			name:   "With fullhd responsive",
			props:  FixedGridProps{Cols10: true, Cols8FullHD: true},
			expect: `<div class="fixed-grid has-10-cols has-8-cols-fullhd"></div>`,
		},
		{
			name:   "With auto count",
			props:  FixedGridProps{HasAutoCount: true},
			expect: `<div class="fixed-grid has-auto-count"></div>`,
		},
		{
			name: "All breakpoints",
			props: FixedGridProps{
				Cols4:           true,
				Cols1Mobile:     true,
				Cols2Tablet:     true,
				Cols3Desktop:    true,
				Cols5Widescreen: true,
				Cols6FullHD:     true,
			},
			expect: `<div class="fixed-grid has-4-cols has-1-cols-mobile has-2-cols-tablet has-3-cols-desktop has-5-cols-widescreen has-6-cols-fullhd"></div>`,
		},
		{
			name: "Complete configuration",
			props: FixedGridProps{
				ID:           "main-grid",
				Class:        []string{"foo", "bar"},
				Cols3:        true,
				Cols1Mobile:  true,
				HasAutoCount: false,
				Attributes:   templ.Attributes{"data-testid": "grid-container"},
			},
			expect: `<div id="main-grid" class="fixed-grid has-3-cols has-1-cols-mobile foo bar" data-testid="grid-container"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := FixedGrid(tt.props).Render(context.Background(), &buf)
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

func TestFixedGridColumns(t *testing.T) {
	tests := []struct {
		name   string
		props  FixedGridProps
		expect string
	}{
		{"1 column", FixedGridProps{Cols1: true}, `<div class="fixed-grid has-1-cols"></div>`},
		{"2 columns", FixedGridProps{Cols2: true}, `<div class="fixed-grid has-2-cols"></div>`},
		{"3 columns", FixedGridProps{Cols3: true}, `<div class="fixed-grid has-3-cols"></div>`},
		{"4 columns", FixedGridProps{Cols4: true}, `<div class="fixed-grid has-4-cols"></div>`},
		{"5 columns", FixedGridProps{Cols5: true}, `<div class="fixed-grid has-5-cols"></div>`},
		{"6 columns", FixedGridProps{Cols6: true}, `<div class="fixed-grid has-6-cols"></div>`},
		{"7 columns", FixedGridProps{Cols7: true}, `<div class="fixed-grid has-7-cols"></div>`},
		{"8 columns", FixedGridProps{Cols8: true}, `<div class="fixed-grid has-8-cols"></div>`},
		{"9 columns", FixedGridProps{Cols9: true}, `<div class="fixed-grid has-9-cols"></div>`},
		{"10 columns", FixedGridProps{Cols10: true}, `<div class="fixed-grid has-10-cols"></div>`},
		{"11 columns", FixedGridProps{Cols11: true}, `<div class="fixed-grid has-11-cols"></div>`},
		{"12 columns", FixedGridProps{Cols12: true}, `<div class="fixed-grid has-12-cols"></div>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := FixedGrid(tt.props).Render(context.Background(), &buf)
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

func TestGrid(t *testing.T) {
	tests := []struct {
		name   string
		props  GridProps
		expect string
	}{
		{
			name:   "Default",
			props:  GridProps{},
			expect: `<div class="grid"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  GridProps{ID: "smart-grid", Class: []string{"custom-smart"}},
			expect: `<div id="smart-grid" class="grid custom-smart"></div>`,
		},
		{
			name:   "With column min 1",
			props:  GridProps{ColMin: ColMin1},
			expect: `<div class="grid is-col-min-1"></div>`,
		},
		{
			name:   "With column min 6 (default)",
			props:  GridProps{ColMin: ColMin6},
			expect: `<div class="grid is-col-min-6"></div>`,
		},
		{
			name:   "With column min 12",
			props:  GridProps{ColMin: ColMin12},
			expect: `<div class="grid is-col-min-12"></div>`,
		},
		{
			name:   "With column min 32",
			props:  GridProps{ColMin: ColMin32},
			expect: `<div class="grid is-col-min-32"></div>`,
		},
		{
			name:   "With auto fill",
			props:  GridProps{IsAutoFill: true},
			expect: `<div class="grid is-auto-fill"></div>`,
		},
		{
			name:   "With column min and auto fill",
			props:  GridProps{ColMin: ColMin8, IsAutoFill: true},
			expect: `<div class="grid is-col-min-8 is-auto-fill"></div>`,
		},
		{
			name: "Complete configuration",
			props: GridProps{
				ID:         "smart-grid-main",
				Class:      []string{"responsive", "flexible"},
				ColMin:     ColMin4,
				IsAutoFill: true,
				Attributes: templ.Attributes{"data-layout": "smart"},
			},
			expect: `<div id="smart-grid-main" class="grid is-col-min-4 is-auto-fill responsive flexible" data-layout="smart"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Grid(tt.props).Render(context.Background(), &buf)
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

func TestGridColMinValues(t *testing.T) {
	tests := []struct {
		name   string
		colMin ColMin
		expect string
	}{
		{"ColMin 1", ColMin1, `<div class="grid is-col-min-1"></div>`},
		{"ColMin 5", ColMin5, `<div class="grid is-col-min-5"></div>`},
		{"ColMin 10", ColMin10, `<div class="grid is-col-min-10"></div>`},
		{"ColMin 15", ColMin15, `<div class="grid is-col-min-15"></div>`},
		{"ColMin 20", ColMin20, `<div class="grid is-col-min-20"></div>`},
		{"ColMin 25", ColMin25, `<div class="grid is-col-min-25"></div>`},
		{"ColMin 30", ColMin30, `<div class="grid is-col-min-30"></div>`},
		{"ColMin 32", ColMin32, `<div class="grid is-col-min-32"></div>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Grid(GridProps{ColMin: tt.colMin}).Render(context.Background(), &buf)
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

func TestCell(t *testing.T) {
	tests := []struct {
		name   string
		props  CellProps
		expect string
	}{
		{
			name:   "Default",
			props:  CellProps{},
			expect: `<div class="cell"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  CellProps{ID: "cell1", Class: []string{"custom-cell"}},
			expect: `<div id="cell1" class="cell custom-cell"></div>`,
		},
		{
			name:   "Column start",
			props:  CellProps{ColStart: 2},
			expect: `<div class="cell is-col-start-2"></div>`,
		},
		{
			name:   "Column end",
			props:  CellProps{ColEnd: 6},
			expect: `<div class="cell is-col-end-6"></div>`,
		},
		{
			name:   "Column span",
			props:  CellProps{ColSpan: 3},
			expect: `<div class="cell is-col-span-3"></div>`,
		},
		{
			name:   "Column from end",
			props:  CellProps{ColFromEnd: 2},
			expect: `<div class="cell is-col-from-end-2"></div>`,
		},
		{
			name:   "Row start",
			props:  CellProps{RowStart: 1},
			expect: `<div class="cell is-row-start-1"></div>`,
		},
		{
			name:   "Row span",
			props:  CellProps{RowSpan: 2},
			expect: `<div class="cell is-row-span-2"></div>`,
		},
		{
			name:   "Is col start end",
			props:  CellProps{IsColStartEnd: true},
			expect: `<div class="cell is-col-start-end"></div>`,
		},
		{
			name:   "Is row start end",
			props:  CellProps{IsRowStartEnd: true},
			expect: `<div class="cell is-row-start-end"></div>`,
		},
		{
			name:   "Mobile responsive column",
			props:  CellProps{ColSpanMobile: 6},
			expect: `<div class="cell is-col-span-6-mobile"></div>`,
		},
		{
			name:   "Tablet responsive column",
			props:  CellProps{ColStartTablet: 3},
			expect: `<div class="cell is-col-start-3-tablet"></div>`,
		},
		{
			name:   "Desktop responsive column",
			props:  CellProps{ColSpanDesktop: 4},
			expect: `<div class="cell is-col-span-4-desktop"></div>`,
		},
		{
			name:   "Widescreen responsive column",
			props:  CellProps{ColEndWidescreen: 8},
			expect: `<div class="cell is-col-end-8-widescreen"></div>`,
		},
		{
			name:   "FullHD responsive column",
			props:  CellProps{ColFromEndFullHD: 2},
			expect: `<div class="cell is-col-from-end-2-fullhd"></div>`,
		},
		{
			name:   "Mobile responsive row",
			props:  CellProps{RowSpanMobile: 3},
			expect: `<div class="cell is-row-span-3-mobile"></div>`,
		},
		{
			name:   "Desktop responsive row",
			props:  CellProps{RowStartDesktop: 2},
			expect: `<div class="cell is-row-start-2-desktop"></div>`,
		},
		{
			name:   "Multiple column properties",
			props:  CellProps{ColStart: 1, ColSpan: 4, ColEnd: 5},
			expect: `<div class="cell is-col-start-1 is-col-end-5 is-col-span-4"></div>`,
		},
		{
			name:   "Multiple row properties",
			props:  CellProps{RowStart: 1, RowSpan: 2, RowFromEnd: 3},
			expect: `<div class="cell is-row-start-1 is-row-from-end-3 is-row-span-2"></div>`,
		},
		{
			name: "Complex responsive layout",
			props: CellProps{
				ColSpan:        12,
				ColSpanMobile:  6,
				ColSpanTablet:  4,
				ColSpanDesktop: 3,
				RowSpan:        1,
				RowSpanMobile:  2,
			},
			expect: `<div class="cell is-col-span-12 is-row-span-1 is-col-span-6-mobile is-row-span-2-mobile is-col-span-4-tablet is-col-span-3-desktop"></div>`,
		},
		{
			name: "All properties combined",
			props: CellProps{
				ID:             "main-cell",
				Class:          []string{"featured", "highlight"},
				ColStart:       2,
				ColSpan:        4,
				RowStart:       1,
				RowSpan:        2,
				ColSpanMobile:  12,
				ColStartTablet: 1,
				ColSpanDesktop: 3,
				IsColStartEnd:  false,
				IsRowStartEnd:  false,
				Attributes:     templ.Attributes{"data-role": "main-content"},
			},
			expect: `<div id="main-cell" class="cell is-col-start-2 is-col-span-4 is-row-start-1 is-row-span-2 is-col-span-12-mobile is-col-start-1-tablet is-col-span-3-desktop featured highlight" data-role="main-content"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Cell(tt.props).Render(context.Background(), &buf)
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

func TestCellPositioning(t *testing.T) {
	tests := []struct {
		name   string
		props  CellProps
		expect string
	}{
		// Column positioning
		{"Col start 1", CellProps{ColStart: 1}, `<div class="cell is-col-start-1"></div>`},
		{"Col start 6", CellProps{ColStart: 6}, `<div class="cell is-col-start-6"></div>`},
		{"Col start 12", CellProps{ColStart: 12}, `<div class="cell is-col-start-12"></div>`},
		{"Col end 1", CellProps{ColEnd: 1}, `<div class="cell is-col-end-1"></div>`},
		{"Col end 6", CellProps{ColEnd: 6}, `<div class="cell is-col-end-6"></div>`},
		{"Col end 12", CellProps{ColEnd: 12}, `<div class="cell is-col-end-12"></div>`},
		{"Col span 1", CellProps{ColSpan: 1}, `<div class="cell is-col-span-1"></div>`},
		{"Col span 6", CellProps{ColSpan: 6}, `<div class="cell is-col-span-6"></div>`},
		{"Col span 12", CellProps{ColSpan: 12}, `<div class="cell is-col-span-12"></div>`},
		{"Col from end 1", CellProps{ColFromEnd: 1}, `<div class="cell is-col-from-end-1"></div>`},
		{"Col from end 6", CellProps{ColFromEnd: 6}, `<div class="cell is-col-from-end-6"></div>`},
		{"Col from end 12", CellProps{ColFromEnd: 12}, `<div class="cell is-col-from-end-12"></div>`},

		// Row positioning
		{"Row start 1", CellProps{RowStart: 1}, `<div class="cell is-row-start-1"></div>`},
		{"Row start 6", CellProps{RowStart: 6}, `<div class="cell is-row-start-6"></div>`},
		{"Row start 12", CellProps{RowStart: 12}, `<div class="cell is-row-start-12"></div>`},
		{"Row end 1", CellProps{RowEnd: 1}, `<div class="cell is-row-end-1"></div>`},
		{"Row end 6", CellProps{RowEnd: 6}, `<div class="cell is-row-end-6"></div>`},
		{"Row end 12", CellProps{RowEnd: 12}, `<div class="cell is-row-end-12"></div>`},
		{"Row span 1", CellProps{RowSpan: 1}, `<div class="cell is-row-span-1"></div>`},
		{"Row span 6", CellProps{RowSpan: 6}, `<div class="cell is-row-span-6"></div>`},
		{"Row span 12", CellProps{RowSpan: 12}, `<div class="cell is-row-span-12"></div>`},
		{"Row from end 1", CellProps{RowFromEnd: 1}, `<div class="cell is-row-from-end-1"></div>`},
		{"Row from end 6", CellProps{RowFromEnd: 6}, `<div class="cell is-row-from-end-6"></div>`},
		{"Row from end 12", CellProps{RowFromEnd: 12}, `<div class="cell is-row-from-end-12"></div>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Cell(tt.props).Render(context.Background(), &buf)
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

func TestCellResponsive(t *testing.T) {
	tests := []struct {
		name   string
		props  CellProps
		expect string
	}{
		// Mobile responsive
		{"Mobile col start", CellProps{ColStartMobile: 2}, `<div class="cell is-col-start-2-mobile"></div>`},
		{"Mobile col span", CellProps{ColSpanMobile: 6}, `<div class="cell is-col-span-6-mobile"></div>`},
		{"Mobile row start", CellProps{RowStartMobile: 1}, `<div class="cell is-row-start-1-mobile"></div>`},
		{"Mobile row span", CellProps{RowSpanMobile: 3}, `<div class="cell is-row-span-3-mobile"></div>`},

		// Tablet responsive
		{"Tablet col start", CellProps{ColStartTablet: 3}, `<div class="cell is-col-start-3-tablet"></div>`},
		{"Tablet col span", CellProps{ColSpanTablet: 4}, `<div class="cell is-col-span-4-tablet"></div>`},
		{"Tablet row start", CellProps{RowStartTablet: 2}, `<div class="cell is-row-start-2-tablet"></div>`},
		{"Tablet row span", CellProps{RowSpanTablet: 2}, `<div class="cell is-row-span-2-tablet"></div>`},

		// Desktop responsive
		{"Desktop col start", CellProps{ColStartDesktop: 1}, `<div class="cell is-col-start-1-desktop"></div>`},
		{"Desktop col span", CellProps{ColSpanDesktop: 8}, `<div class="cell is-col-span-8-desktop"></div>`},
		{"Desktop row start", CellProps{RowStartDesktop: 3}, `<div class="cell is-row-start-3-desktop"></div>`},
		{"Desktop row span", CellProps{RowSpanDesktop: 4}, `<div class="cell is-row-span-4-desktop"></div>`},

		// Widescreen responsive
		{"Widescreen col start", CellProps{ColStartWidescreen: 4}, `<div class="cell is-col-start-4-widescreen"></div>`},
		{"Widescreen col span", CellProps{ColSpanWidescreen: 6}, `<div class="cell is-col-span-6-widescreen"></div>`},
		{"Widescreen row start", CellProps{RowStartWidescreen: 1}, `<div class="cell is-row-start-1-widescreen"></div>`},
		{"Widescreen row span", CellProps{RowSpanWidescreen: 5}, `<div class="cell is-row-span-5-widescreen"></div>`},

		// FullHD responsive
		{"FullHD col start", CellProps{ColStartFullHD: 2}, `<div class="cell is-col-start-2-fullhd"></div>`},
		{"FullHD col span", CellProps{ColSpanFullHD: 10}, `<div class="cell is-col-span-10-fullhd"></div>`},
		{"FullHD row start", CellProps{RowStartFullHD: 2}, `<div class="cell is-row-start-2-fullhd"></div>`},
		{"FullHD row span", CellProps{RowSpanFullHD: 3}, `<div class="cell is-row-span-3-fullhd"></div>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Cell(tt.props).Render(context.Background(), &buf)
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

func TestCellInvalidValues(t *testing.T) {
	tests := []struct {
		name   string
		props  CellProps
		expect string
	}{
		{"Col start 0 (invalid)", CellProps{ColStart: 0}, `<div class="cell"></div>`},
		{"Col start 13 (invalid)", CellProps{ColStart: 13}, `<div class="cell"></div>`},
		{"Col span -1 (invalid)", CellProps{ColSpan: -1}, `<div class="cell"></div>`},
		{"Row start 15 (invalid)", CellProps{RowStart: 15}, `<div class="cell"></div>`},
		{"Mobile col span 0 (invalid)", CellProps{ColSpanMobile: 0}, `<div class="cell"></div>`},
		{"Desktop row span 20 (invalid)", CellProps{RowSpanDesktop: 20}, `<div class="cell"></div>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Cell(tt.props).Render(context.Background(), &buf)
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

func TestFormatCellClass(t *testing.T) {
	tests := []struct {
		name        string
		prefix      string
		value       int
		suffix      string
		expectClass string
		expectValid bool
	}{
		{"Valid base class", "is-col-start", 5, "", "is-col-start-5", true},
		{"Valid responsive class", "is-col-span", 3, "mobile", "is-col-span-3-mobile", true},
		{"Invalid value 0", "is-col-start", 0, "", "", false},
		{"Invalid value 13", "is-col-end", 13, "", "", false},
		{"Invalid negative value", "is-row-span", -5, "tablet", "", false},
		{"Valid edge case 1", "is-row-start", 1, "desktop", "is-row-start-1-desktop", true},
		{"Valid edge case 12", "is-col-from-end", 12, "fullhd", "is-col-from-end-12-fullhd", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotClass, gotValid := formatCellClass(tt.prefix, tt.value, tt.suffix)
			if gotClass != tt.expectClass {
				t.Errorf("expected class %q, got %q", tt.expectClass, gotClass)
			}
			if gotValid != tt.expectValid {
				t.Errorf("expected valid %v, got %v", tt.expectValid, gotValid)
			}
		})
	}
}

func TestFixedGridResponsiveBreakpoints(t *testing.T) {
	tests := []struct {
		name   string
		props  FixedGridProps
		expect string
	}{
		{
			name: "All mobile breakpoints",
			props: FixedGridProps{
				Cols1Mobile:  true,
				Cols6Mobile:  true,
				Cols12Mobile: true,
			},
			expect: `<div class="fixed-grid has-1-cols-mobile has-6-cols-mobile has-12-cols-mobile"></div>`,
		},
		{
			name: "All tablet breakpoints",
			props: FixedGridProps{
				Cols2Tablet: true,
				Cols4Tablet: true,
				Cols8Tablet: true,
			},
			expect: `<div class="fixed-grid has-2-cols-tablet has-4-cols-tablet has-8-cols-tablet"></div>`,
		},
		{
			name: "All desktop breakpoints",
			props: FixedGridProps{
				Cols3Desktop: true,
				Cols6Desktop: true,
				Cols9Desktop: true,
			},
			expect: `<div class="fixed-grid has-3-cols-desktop has-6-cols-desktop has-9-cols-desktop"></div>`,
		},
		{
			name: "All widescreen breakpoints",
			props: FixedGridProps{
				Cols4Widescreen:  true,
				Cols8Widescreen:  true,
				Cols12Widescreen: true,
			},
			expect: `<div class="fixed-grid has-4-cols-widescreen has-8-cols-widescreen has-12-cols-widescreen"></div>`,
		},
		{
			name: "All fullhd breakpoints",
			props: FixedGridProps{
				Cols1FullHD:  true,
				Cols5FullHD:  true,
				Cols10FullHD: true,
			},
			expect: `<div class="fixed-grid has-1-cols-fullhd has-5-cols-fullhd has-10-cols-fullhd"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := FixedGrid(tt.props).Render(context.Background(), &buf)
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
