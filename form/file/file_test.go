package file

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestFile(t *testing.T) {
	tests := []struct {
		name   string
		props  FileProps
		expect string
	}{
		{
			name:   "Default",
			props:  FileProps{},
			expect: `<div class="file"></div>`,
		},
		{
			name:   "With ID and classes",
			props:  FileProps{ID: "test-file", Class: []string{"custom", "upload"}},
			expect: `<div id="test-file" class="file custom upload"></div>`,
		},
		{
			name:   "Has name",
			props:  FileProps{HasName: true},
			expect: `<div class="file has-name"></div>`,
		},
		{
			name:   "Is boxed",
			props:  FileProps{IsBoxed: true},
			expect: `<div class="file is-boxed"></div>`,
		},
		{
			name:   "Is fullwidth",
			props:  FileProps{IsFullwidth: true},
			expect: `<div class="file is-fullwidth"></div>`,
		},
		{
			name:   "Is centered",
			props:  FileProps{IsCentered: true},
			expect: `<div class="file is-centered"></div>`,
		},
		{
			name:   "Is right",
			props:  FileProps{IsRight: true},
			expect: `<div class="file is-right"></div>`,
		},
		{
			name:   "Small size",
			props:  FileProps{Size: IsSmall},
			expect: `<div class="file is-small"></div>`,
		},
		{
			name:   "Large size",
			props:  FileProps{Size: IsLarge},
			expect: `<div class="file is-large"></div>`,
		},
		{
			name:   "Primary color",
			props:  FileProps{Color: IsPrimary},
			expect: `<div class="file is-primary"></div>`,
		},
		{
			name:   "Danger color",
			props:  FileProps{Color: IsDanger},
			expect: `<div class="file is-danger"></div>`,
		},
		{
			name: "Complete configuration",
			props: FileProps{
				ID:          "upload-file",
				Class:       []string{"form-upload", "document"},
				HasName:     true,
				IsBoxed:     true,
				IsFullwidth: true,
				Size:        IsMedium,
				Color:       IsSuccess,
			},
			expect: `<div id="upload-file" class="file has-name is-boxed is-fullwidth is-medium is-success form-upload document"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := File(tt.props).Render(context.Background(), &buf)
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

func TestFileLabel(t *testing.T) {
	tests := []struct {
		name   string
		props  FileLabelProps
		expect string
	}{
		{
			name:   "Default",
			props:  FileLabelProps{},
			expect: `<label class="file-label"></label>`,
		},
		{
			name:   "With ID and classes",
			props:  FileLabelProps{ID: "test-label", Class: []string{"custom"}},
			expect: `<label id="test-label" class="file-label custom"></label>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := FileLabel(tt.props).Render(context.Background(), &buf)
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

func TestFileInput(t *testing.T) {
	tests := []struct {
		name   string
		props  FileInputProps
		expect string
	}{
		{
			name:   "Default",
			props:  FileInputProps{},
			expect: `<input type="file" class="file-input">`,
		},
		{
			name:   "With ID and name",
			props:  FileInputProps{ID: "test-input", Name: "upload"},
			expect: `<input id="test-input" type="file" name="upload" class="file-input">`,
		},
		{
			name:   "With accept",
			props:  FileInputProps{Accept: "image/*"},
			expect: `<input type="file" accept="image/*" class="file-input">`,
		},
		{
			name:   "Multiple",
			props:  FileInputProps{Multiple: true},
			expect: `<input type="file" multiple class="file-input">`,
		},
		{
			name:   "Disabled",
			props:  FileInputProps{Disabled: true},
			expect: `<input type="file" disabled class="file-input">`,
		},
		{
			name:   "Required",
			props:  FileInputProps{Required: true},
			expect: `<input type="file" required class="file-input">`,
		},
		{
			name: "Complete configuration",
			props: FileInputProps{
				ID:       "document-upload",
				Name:     "documents",
				Accept:   ".pdf,.doc,.docx",
				Multiple: true,
				Required: true,
				Class:    []string{"custom-input"},
			},
			expect: `<input id="document-upload" type="file" name="documents" accept=".pdf,.doc,.docx" multiple required class="file-input custom-input">`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := FileInput(tt.props).Render(context.Background(), &buf)
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

func TestFileCTA(t *testing.T) {
	tests := []struct {
		name   string
		props  FileCTAProps
		expect string
	}{
		{
			name:   "Default",
			props:  FileCTAProps{},
			expect: `<span class="file-cta"></span>`,
		},
		{
			name:   "With ID and classes",
			props:  FileCTAProps{ID: "test-cta", Class: []string{"custom"}},
			expect: `<span id="test-cta" class="file-cta custom"></span>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := FileCTA(tt.props).Render(context.Background(), &buf)
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

func TestFileIcon(t *testing.T) {
	tests := []struct {
		name   string
		props  FileIconProps
		expect string
	}{
		{
			name:   "Default",
			props:  FileIconProps{},
			expect: `<span class="file-icon"></span>`,
		},
		{
			name:   "With ID and classes",
			props:  FileIconProps{ID: "test-icon", Class: []string{"custom"}},
			expect: `<span id="test-icon" class="file-icon custom"></span>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := FileIcon(tt.props).Render(context.Background(), &buf)
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

func TestFileText(t *testing.T) {
	tests := []struct {
		name   string
		props  FileTextProps
		expect string
	}{
		{
			name:   "Default",
			props:  FileTextProps{},
			expect: `<span class="file-label"></span>`,
		},
		{
			name:   "With ID and classes",
			props:  FileTextProps{ID: "test-text", Class: []string{"custom"}},
			expect: `<span id="test-text" class="file-label custom"></span>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := FileText(tt.props).Render(context.Background(), &buf)
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

func TestFileName(t *testing.T) {
	tests := []struct {
		name   string
		props  FileNameProps
		expect string
	}{
		{
			name:   "Default",
			props:  FileNameProps{},
			expect: `<span class="file-name"></span>`,
		},
		{
			name:   "With ID and classes",
			props:  FileNameProps{ID: "test-name", Class: []string{"custom"}},
			expect: `<span id="test-name" class="file-name custom"></span>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := FileName(tt.props).Render(context.Background(), &buf)
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

func TestFileSizes(t *testing.T) {
	tests := []struct {
		name   string
		size   Size
		expect string
	}{
		{"Small", IsSmall, `<div class="file is-small"></div>`},
		{"Normal", IsNormal, `<div class="file is-normal"></div>`},
		{"Medium", IsMedium, `<div class="file is-medium"></div>`},
		{"Large", IsLarge, `<div class="file is-large"></div>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := File(FileProps{Size: tt.size}).Render(context.Background(), &buf)
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

func TestFileColors(t *testing.T) {
	tests := []struct {
		name   string
		color  Color
		expect string
	}{
		{"White", IsWhite, `<div class="file is-white"></div>`},
		{"Black", IsBlack, `<div class="file is-black"></div>`},
		{"Light", IsLight, `<div class="file is-light"></div>`},
		{"Dark", IsDark, `<div class="file is-dark"></div>`},
		{"Primary", IsPrimary, `<div class="file is-primary"></div>`},
		{"Link", IsLink, `<div class="file is-link"></div>`},
		{"Info", IsInfo, `<div class="file is-info"></div>`},
		{"Success", IsSuccess, `<div class="file is-success"></div>`},
		{"Warning", IsWarning, `<div class="file is-warning"></div>`},
		{"Danger", IsDanger, `<div class="file is-danger"></div>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := File(FileProps{Color: tt.color}).Render(context.Background(), &buf)
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

func TestFileAttributes(t *testing.T) {
	props := FileProps{
		Attributes: templ.Attributes{
			"data-testid": "file-upload",
			"aria-label":  "File upload",
		},
	}

	var buf strings.Builder
	err := File(props).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}
	got := strings.TrimSpace(buf.String())

	for key, value := range props.Attributes {
		expectedAttr := key + `="` + fmt.Sprintf("%v", value) + `"`
		if !strings.Contains(got, expectedAttr) {
			t.Errorf("expected attribute %s to be present in: %s", expectedAttr, got)
		}
	}
}
