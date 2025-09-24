package table

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestContainer(t *testing.T) {
	tests := []struct {
		name   string
		props  ContainerProps
		expect string
	}{
		{
			name:   "Default",
			props:  ContainerProps{},
			expect: `<div class="table-container"></div>`,
		},
		{
			name:   "With ID and custom classes",
			props:  ContainerProps{ID: "container1", Class: []string{"custom-container"}},
			expect: `<div id="container1" class="table-container custom-container"></div>`,
		},
		{
			name: "With attributes",
			props: ContainerProps{
				ID:         "responsive-table",
				Attributes: templ.Attributes{"data-testid": "table-wrapper"},
			},
			expect: `<div id="responsive-table" class="table-container" data-testid="table-wrapper"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Container(tt.props).Render(context.Background(), &buf)
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

func TestTable(t *testing.T) {
	tests := []struct {
		name   string
		props  TableProps
		expect string
	}{
		{
			name:   "Default",
			props:  TableProps{},
			expect: `<table class="table"></table>`,
		},
		{
			name:   "With ID and custom classes",
			props:  TableProps{ID: "table1", Class: []string{"custom-table"}},
			expect: `<table id="table1" class="table custom-table"></table>`,
		},
		{
			name:   "With caption",
			props:  TableProps{Caption: "Sales Data"},
			expect: `<table class="table"><caption>Sales Data</caption></table>`,
		},
		{
			name:   "With bordered style",
			props:  TableProps{IsBordered: true},
			expect: `<table class="table is-bordered"></table>`,
		},
		{
			name:   "With fullwidth style",
			props:  TableProps{IsFullwidth: true},
			expect: `<table class="table is-fullwidth"></table>`,
		},
		{
			name:   "With narrow style",
			props:  TableProps{IsNarrow: true},
			expect: `<table class="table is-narrow"></table>`,
		},
		{
			name:   "With striped style",
			props:  TableProps{IsStriped: true},
			expect: `<table class="table is-striped"></table>`,
		},
		{
			name:   "With vcentered layout",
			props:  TableProps{IsVCentered: true},
			expect: `<table class="table is-vcentered"></table>`,
		},
		{
			name:   "With hoverable state",
			props:  TableProps{IsHoverable: true},
			expect: `<table class="table is-hoverable"></table>`,
		},
		{
			name:   "Multiple modifiers",
			props:  TableProps{IsStriped: true, IsHoverable: true, IsBordered: true},
			expect: `<table class="table is-bordered is-striped is-hoverable"></table>`,
		},
		{
			name: "All fields combined",
			props: TableProps{
				ID:          "main-table",
				Class:       []string{"foo", "bar"},
				Caption:     "Monthly Report",
				IsBordered:  true,
				IsFullwidth: true,
				IsStriped:   true,
				IsHoverable: true,
				Attributes:  templ.Attributes{"data-testid": "report-table"},
			},
			expect: `<table id="main-table" class="table is-bordered is-fullwidth is-striped is-hoverable foo bar" data-testid="report-table"><caption>Monthly Report</caption></table>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Table(tt.props).Render(context.Background(), &buf)
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

func TestHead(t *testing.T) {
	tests := []struct {
		name   string
		props  HeadProps
		expect string
	}{
		{
			name:   "Default",
			props:  HeadProps{},
			expect: `<thead></thead>`,
		},
		{
			name:   "With ID and custom classes",
			props:  HeadProps{ID: "head1", Class: []string{"custom-head"}},
			expect: `<thead id="head1" class="custom-head"></thead>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Head(tt.props).Render(context.Background(), &buf)
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

func TestBody(t *testing.T) {
	tests := []struct {
		name   string
		props  BodyProps
		expect string
	}{
		{
			name:   "Default",
			props:  BodyProps{},
			expect: `<tbody></tbody>`,
		},
		{
			name:   "With ID and custom classes",
			props:  BodyProps{ID: "body1", Class: []string{"custom-body"}},
			expect: `<tbody id="body1" class="custom-body"></tbody>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Body(tt.props).Render(context.Background(), &buf)
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

func TestFoot(t *testing.T) {
	tests := []struct {
		name   string
		props  FootProps
		expect string
	}{
		{
			name:   "Default",
			props:  FootProps{},
			expect: `<tfoot></tfoot>`,
		},
		{
			name:   "With ID and custom classes",
			props:  FootProps{ID: "foot1", Class: []string{"custom-foot"}},
			expect: `<tfoot id="foot1" class="custom-foot"></tfoot>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Foot(tt.props).Render(context.Background(), &buf)
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

func TestRow(t *testing.T) {
	tests := []struct {
		name   string
		props  RowProps
		expect string
	}{
		{
			name:   "Default",
			props:  RowProps{},
			expect: `<tr></tr>`,
		},
		{
			name:   "With ID and custom classes",
			props:  RowProps{ID: "row1", Class: []string{"custom-row"}},
			expect: `<tr id="row1" class="custom-row"></tr>`,
		},
		{
			name:   "With color",
			props:  RowProps{Color: IsPrimary},
			expect: `<tr class="is-primary"></tr>`,
		},
		{
			name:   "With selected state",
			props:  RowProps{IsSelected: true},
			expect: `<tr class="is-selected"></tr>`,
		},
		{
			name:   "With color and selected",
			props:  RowProps{Color: IsSuccess, IsSelected: true},
			expect: `<tr class="is-success is-selected"></tr>`,
		},
		{
			name: "All fields combined",
			props: RowProps{
				ID:         "main-row",
				Class:      []string{"foo", "bar"},
				Color:      IsDanger,
				IsSelected: true,
				Attributes: templ.Attributes{"data-id": "123"},
			},
			expect: `<tr id="main-row" class="is-danger is-selected foo bar" data-id="123"></tr>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Row(tt.props).Render(context.Background(), &buf)
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

func TestHeader(t *testing.T) {
	tests := []struct {
		name   string
		props  HeaderProps
		expect string
	}{
		{
			name:   "Default",
			props:  HeaderProps{},
			expect: `<th></th>`,
		},
		{
			name:   "With ID and custom classes",
			props:  HeaderProps{ID: "header1", Class: []string{"custom-header"}},
			expect: `<th id="header1" class="custom-header"></th>`,
		},
		{
			name:   "With color",
			props:  HeaderProps{Color: IsPrimary},
			expect: `<th class="is-primary"></th>`,
		},
		{
			name:   "With narrow layout",
			props:  HeaderProps{IsNarrow: true},
			expect: `<th class="is-narrow"></th>`,
		},
		{
			name:   "With vcentered layout",
			props:  HeaderProps{IsVCentered: true},
			expect: `<th class="is-vcentered"></th>`,
		},
		{
			name:   "With selected state",
			props:  HeaderProps{IsSelected: true},
			expect: `<th class="is-selected"></th>`,
		},
		{
			name:   "With scope attribute",
			props:  HeaderProps{Scope: "col"},
			expect: `<th scope="col"></th>`,
		},
		{
			name:   "With row scope",
			props:  HeaderProps{Scope: "row"},
			expect: `<th scope="row"></th>`,
		},
		{
			name: "All fields combined",
			props: HeaderProps{
				ID:          "main-header",
				Class:       []string{"foo", "bar"},
				Color:       IsInfo,
				IsNarrow:    true,
				IsVCentered: true,
				IsSelected:  true,
				Scope:       "colgroup",
				Attributes:  templ.Attributes{"data-sort": "asc"},
			},
			expect: `<th id="main-header" class="is-info is-narrow is-vcentered is-selected foo bar" scope="colgroup" data-sort="asc"></th>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := Header(tt.props).Render(context.Background(), &buf)
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
			expect: `<td></td>`,
		},
		{
			name:   "With ID and custom classes",
			props:  CellProps{ID: "cell1", Class: []string{"custom-cell"}},
			expect: `<td id="cell1" class="custom-cell"></td>`,
		},
		{
			name:   "With color",
			props:  CellProps{Color: IsPrimary},
			expect: `<td class="is-primary"></td>`,
		},
		{
			name:   "With narrow layout",
			props:  CellProps{IsNarrow: true},
			expect: `<td class="is-narrow"></td>`,
		},
		{
			name:   "With vcentered layout",
			props:  CellProps{IsVCentered: true},
			expect: `<td class="is-vcentered"></td>`,
		},
		{
			name:   "With selected state",
			props:  CellProps{IsSelected: true},
			expect: `<td class="is-selected"></td>`,
		},
		{
			name: "All fields combined",
			props: CellProps{
				ID:          "main-cell",
				Class:       []string{"foo", "bar"},
				Color:       IsWarning,
				IsNarrow:    true,
				IsVCentered: true,
				IsSelected:  true,
				Attributes:  templ.Attributes{"data-value": "123"},
			},
			expect: `<td id="main-cell" class="is-warning is-narrow is-vcentered is-selected foo bar" data-value="123"></td>`,
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

func TestTableColors(t *testing.T) {
	colorTests := []struct {
		name   string
		color  Color
		expect string
	}{
		{"Black", IsBlack, "is-black"},
		{"Danger", IsDanger, "is-danger"},
		{"Dark", IsDark, "is-dark"},
		{"Info", IsInfo, "is-info"},
		{"Light", IsLight, "is-light"},
		{"Link", IsLink, "is-link"},
		{"Primary", IsPrimary, "is-primary"},
		{"Success", IsSuccess, "is-success"},
		{"Text", IsText, "is-text"},
		{"Warning", IsWarning, "is-warning"},
		{"White", IsWhite, "is-white"},
	}

	t.Run("Row colors", func(t *testing.T) {
		for _, tt := range colorTests {
			t.Run(tt.name, func(t *testing.T) {
				var buf strings.Builder
				err := Row(RowProps{Color: tt.color}).Render(context.Background(), &buf)
				if err != nil {
					t.Fatalf("render failed: %v", err)
				}
				got := strings.TrimSpace(buf.String())
				expect := `<tr class="` + tt.expect + `"></tr>`
				if got != expect {
					t.Errorf("expected:\n%s\ngot:\n%s", expect, got)
				}
			})
		}
	})

	t.Run("Header colors", func(t *testing.T) {
		for _, tt := range colorTests {
			t.Run(tt.name, func(t *testing.T) {
				var buf strings.Builder
				err := Header(HeaderProps{Color: tt.color}).Render(context.Background(), &buf)
				if err != nil {
					t.Fatalf("render failed: %v", err)
				}
				got := strings.TrimSpace(buf.String())
				expect := `<th class="` + tt.expect + `"></th>`
				if got != expect {
					t.Errorf("expected:\n%s\ngot:\n%s", expect, got)
				}
			})
		}
	})

	t.Run("Cell colors", func(t *testing.T) {
		for _, tt := range colorTests {
			t.Run(tt.name, func(t *testing.T) {
				var buf strings.Builder
				err := Cell(CellProps{Color: tt.color}).Render(context.Background(), &buf)
				if err != nil {
					t.Fatalf("render failed: %v", err)
				}
				got := strings.TrimSpace(buf.String())
				expect := `<td class="` + tt.expect + `"></td>`
				if got != expect {
					t.Errorf("expected:\n%s\ngot:\n%s", expect, got)
				}
			})
		}
	})
}
