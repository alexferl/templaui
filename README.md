# TemplaUI

**⚠️ This is a pre-v1 release - APIs may change as we work toward a stable v1.0.**

A Go library providing [Bulma](https://bulma.io/) components for the [templ](https://templ.guide) templating language.

## Requirements
- Go 1.24+

## Installation
```shell
go get github.com/alexferl/templaui
```

## Usage
`hello.templ`:
```templ
package main

import (
	"github.com/alexferl/templaui"
	"github.com/alexferl/templaui/elements/tag"
	"github.com/alexferl/templaui/elements/title"
	"github.com/alexferl/templaui/layout/container"
)

templ Hello() {
	@templaui.Document(templaui.DocumentProps{
		Title:       "My TemplaUI App",
		Description: "A Hello World example showcasing Bulma components",
	}) {
		@container.Container() {
			@title.Title(title.TitleProps{Size: title.Is1}) {
				Hello, World!
			}
			@title.Subtitle(title.SubtitleProps{Size: title.Is4}) {
				Built with TemplaUI & Bulma CSS
			}
			@tag.Tags() {
				@tag.Tag(tag.TagProps{Color: tag.IsPrimary}) {
					Go
				}
				@tag.Tag(tag.TagProps{Color: tag.IsInfo}) {
					Templ
				}
				@tag.Tag(tag.TagProps{Color: tag.IsSuccess}) {
					Bulma
				}
			}
		}
	}
}
```

`main.go`:
```go
package main

import (
    "fmt"
    "net/http"

    "github.com/a-h/templ"
)

func main() {
    http.Handle("/", templ.Handler(Hello()))
    fmt.Println("Listening on :3000")
    http.ListenAndServe(":3000", nil)
}
```
