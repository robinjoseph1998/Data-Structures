package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

// Block structure for JSON data
type Block struct {
	ID       string                 `json:"id"`
	Type     string                 `json:"type"`
	Props    map[string]interface{} `json:"props"`
	Content  interface{}            `json:"content"`
	Children []Block                `json:"children"`
}

// Helper function to create a unique ID
func generateID() string {
	return fmt.Sprintf("%d", len(dummyData)+1)
}

// // Global dummy data storage for simplicity (replace with your actual handling)
var dummyData []Block

// Parse Markdown file to JSON block format
func ParseMarkdownToBlocks(mdContent string) []Block {
	var blocks []Block

	// Initialize goldmark parser with extensions
	mdParser := goldmark.New(
		goldmark.WithExtensions(
			extension.Strikethrough,  // Enables strikethrough parsing
			extension.Table,          // Enables table parsing
			extension.DefinitionList, // Enables definition list parsing
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
	).Parser()

	reader := text.NewReader([]byte(mdContent))
	doc := mdParser.Parse(reader)

	// Walk the Markdown AST
	ast.Walk(doc, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		switch n := node.(type) {
		case *ast.Heading:
			level := n.Level
			textContent := extractTextContent(n, reader.Source())
			blocks = append(blocks, Block{
				ID:   generateID(),
				Type: "heading",
				Props: map[string]interface{}{
					"textColor":       "default",
					"backgroundColor": "default",
					"textAlignment":   "left",
					"level":           level,
				},
				Content:  textContent,
				Children: nil,
			})
		case *ast.Paragraph:
			textContent := extractTextContent(n, reader.Source())
			blocks = append(blocks, Block{
				ID:   generateID(),
				Type: "paragraph",
				Props: map[string]interface{}{
					"textColor":       "default",
					"backgroundColor": "default",
					"textAlignment":   "left",
				},
				Content:  textContent,
				Children: nil,
			})
		case *ast.List:
			blocks = append(blocks, processListNode(n, reader.Source()))
		case *ast.Blockquote:
			textContent := extractTextContent(n, reader.Source())
			blocks = append(blocks, Block{
				ID:   generateID(),
				Type: "alert",
				Props: map[string]interface{}{
					"textColor":     "default",
					"textAlignment": "left",
					"type":          "warning", // You can determine type dynamically if needed
				},
				Content:  textContent,
				Children: nil,
			})
		default:
			// Handle other Markdown nodes if needed
		}

		return ast.WalkContinue, nil
	})

	return blocks
}

// Extract text content from an AST node
func extractTextContent(node ast.Node, source []byte) []map[string]interface{} {
	var content []map[string]interface{}

	for child := node.FirstChild(); child != nil; child = child.NextSibling() {
		switch child := child.(type) {
		case *ast.Text:
			// Plain text
			content = append(content, map[string]interface{}{
				"text":   string(child.Text(source)),
				"type":   "text",
				"styles": map[string]interface{}{}, // No styles for plain text
			})
		case *ast.Emphasis:
			// Emphasis can be italic, bold, or both
			childContent := extractTextContent(child, source)
			styles := map[string]interface{}{}

			if child.Level == 1 {
				styles["italic"] = true // Single `*` or `_`
			}
			if child.Level == 2 {
				styles["bold"] = true // Double `**` or `__`
			}

			for _, c := range childContent {
				for key, value := range styles {
					c["styles"].(map[string]interface{})[key] = value
				}
				content = append(content, c)
			}
		default:
			// Handle other inline elements if needed
		}
	}

	return content
}

// Process list nodes (bullet or numbered)
func processListNode(listNode *ast.List, source []byte) Block {
	children := []Block{}
	listType := "bulletListItem"
	if listNode.IsOrdered() {
		listType = "numberedListItem"
	}

	for item := listNode.FirstChild(); item != nil; item = item.NextSibling() {
		if listItem, ok := item.(*ast.ListItem); ok {
			listChildren := []Block{}
			for subItem := listItem.FirstChild(); subItem != nil; subItem = subItem.NextSibling() {
				if subList, ok := subItem.(*ast.List); ok {
					listChildren = append(listChildren, processListNode(subList, source))
				} else {
					listChildren = append(listChildren, Block{
						ID:       generateID(),
						Type:     listType,
						Props:    map[string]interface{}{},
						Content:  extractTextContent(subItem, source),
						Children: nil,
					})
				}
			}
			children = append(children, Block{
				ID:       generateID(),
				Type:     listType,
				Props:    map[string]interface{}{},
				Content:  extractTextContent(listItem, source),
				Children: listChildren,
			})
		}
	}

	return Block{
		ID:       generateID(),
		Type:     listType,
		Props:    map[string]interface{}{},
		Content:  nil,
		Children: children,
	}
}

// Convert blocks to JSON string
func ConvertBlocksToJSONString(blocks []Block) (string, error) {
	jsonData, err := json.MarshalIndent(blocks, "", "  ")
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

// Main function
func main() {
	FilePath := "Markdowns/Example.md"
	mdContent, err := ioutil.ReadFile(FilePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	blocks := ParseMarkdownToBlocks(string(mdContent))
	jsonString, err := ConvertBlocksToJSONString(blocks)
	if err != nil {
		log.Fatalf("Error converting blocks to JSON string: %v", err)
	}

	fmt.Println("JSON String Representation:")
	fmt.Println(jsonString)
}
