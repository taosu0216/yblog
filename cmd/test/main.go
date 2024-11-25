package main

import (
	"blug/internal/pkg"
	"blug/internal/pkg/markdown"
	"fmt"
)

func main() {
	markdown.InitRenderer()
	content, _ := pkg.GetArticleContent("_posts/initDemo.md")
	fmt.Println(markdown.GetRenderer().Render(content))
}
