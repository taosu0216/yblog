package markdown

import "gitlab.com/golang-commonmark/markdown"

type Renderer struct {
	markdown *markdown.Markdown
}

var Render *Renderer

func InitRenderer() {
	md := markdown.New(markdown.XHTMLOutput(true))
	Render = &Renderer{markdown: md}
}

func GetRenderer() *Renderer {
	return Render
}

func (r *Renderer) Render(content string) string {
	return r.markdown.RenderToString([]byte(content))
}
