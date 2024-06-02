package processor

import (
	"bytes"
	"fmt"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
	"log"
)

func ConvertToHtml(content []byte) string {
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM, extension.Typographer),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
	var buf bytes.Buffer
	if err := md.Convert(content, &buf); err != nil {
		log.Fatalf("Failed to convert markdown: %v", err)
	}
	htmlContent := buf.String()
	htmlContent += `
    <script type="module">
    import mermaid from 'https://cdn.jsdelivr.net/npm/mermaid@10/dist/mermaid.esm.min.mjs';
    mermaid.initialize({ startOnLoad: true });
    </script>
    `
	fmt.Println("html-content success:", htmlContent)
	return htmlContent
}
