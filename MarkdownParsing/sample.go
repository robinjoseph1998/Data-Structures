// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"regexp"
// 	"strings"

// 	"github.com/google/uuid"
// )

// // Generate unique IDs for blocks
// func generateID() string {
// 	return uuid.New().String()
// }

// // Add default props to each block
// func defaultProps() map[string]interface{} {
// 	return map[string]interface{}{
// 		"textColor":       "default",
// 		"backgroundColor": "default",
// 		"textAlignment":   "left",
// 	}
// }

// func parseMarkDownBlock(markdown string) ([]map[string]interface{}, error) {
// 	var blocks []map[string]interface{}
// 	lines := strings.Split(markdown, "\n")

// 	var currentParent map[string]interface{}
// 	for _, line := range lines {
// 		trimmedLine := strings.TrimSpace(line)

// 		var block map[string]interface{}
// 		switch {
// 		case strings.HasPrefix(trimmedLine, "#"):
// 			block = parseHeading(trimmedLine)
// 		case strings.HasPrefix(trimmedLine, "- [ ]") || strings.HasPrefix(trimmedLine, "- [x]"):
// 			block = parseCheckListItem(trimmedLine)
// 		case strings.HasPrefix(trimmedLine, "-"):
// 			block = parseBulletListItem(trimmedLine)
// 		case isNumberedList(trimmedLine):
// 			block = parseNumberedListItem(trimmedLine)
// 		case isTableRow(trimmedLine):
// 			block = parseTable(trimmedLine)
// 		case isMediaBlock(trimmedLine):
// 			block = parseMedia()
// 		case isAlertBlock(trimmedLine):
// 			block = parseAlert(trimmedLine)
// 		default:
// 			block = parseParagraph(trimmedLine)
// 		}

// 		if block != nil {
// 			if currentParent != nil && block["type"] == "numberedListItem" || block["type"] == "bulletListItem" {
// 				children := currentParent["children"].([]map[string]interface{})
// 				children = append(children, block)
// 				currentParent["children"] = children
// 			} else {
// 				blocks = append(blocks, block)
// 				if block["type"] == "numberedListItem" || block["type"] == "bulletListItem" {
// 					block["children"] = []map[string]interface{}{}
// 					currentParent = block
// 				} else {
// 					currentParent = nil
// 				}
// 			}
// 		}
// 	}

// 	return blocks, nil
// }

// // Functions for parsing specific block types
// func parseHeading(line string) map[string]interface{} {
// 	level := strings.Count(line, "#")
// 	text := strings.TrimSpace(strings.TrimLeft(line, "#"))
// 	return createBlock("heading", map[string]interface{}{"level": level}, []map[string]interface{}{
// 		{"type": "text", "text": text, "styles": map[string]interface{}{}},
// 	})
// }

// func parseCheckListItem(line string) map[string]interface{} {
// 	checked := strings.HasPrefix(line, "- [x]")
// 	text := strings.TrimSpace(line[5:])
// 	return createBlock("checkListItem", map[string]interface{}{"checked": checked}, []map[string]interface{}{
// 		{"type": "text", "text": text, "styles": map[string]interface{}{}},
// 	})
// }

// func parseBulletListItem(line string) map[string]interface{} {
// 	text := strings.TrimSpace(strings.TrimPrefix(line, "-"))
// 	return createBlock("bulletListItem", nil, []map[string]interface{}{
// 		{"type": "text", "text": text, "styles": map[string]interface{}{}},
// 	})
// }

// func parseNumberedListItem(line string) map[string]interface{} {
// 	parts := strings.SplitN(line, ".", 2)
// 	text := strings.TrimSpace(parts[1])
// 	return createBlock("numberedListItem", nil, []map[string]interface{}{
// 		{"type": "text", "text": text, "styles": map[string]interface{}{}},
// 	})
// }

// func parseTable(line string) map[string]interface{} {
// 	columns := strings.Split(line, "|")
// 	var cells [][]map[string]interface{}
// 	for _, column := range columns {
// 		text := strings.TrimSpace(column)
// 		if text != "" {
// 			cells = append(cells, []map[string]interface{}{
// 				{"type": "text", "text": text, "styles": map[string]interface{}{}},
// 			})
// 		}
// 	}
// 	return createBlock("table", nil, map[string]interface{}{"rows": []map[string]interface{}{{"cells": cells}}})
// }

// func parseMedia() map[string]interface{} {
// 	return createBlock("media", map[string]interface{}{
// 		"url": "example.com",
// 	}, nil)
// }

// func parseAlert(line string) map[string]interface{} {
// 	alertType := strings.TrimSpace(strings.TrimPrefix(line, "!"))
// 	return createBlock("alert", map[string]interface{}{"type": alertType}, []map[string]interface{}{
// 		{"type": "text", "text": alertType, "styles": map[string]interface{}{"bold": true}},
// 	})
// }

// func parseParagraph(line string) map[string]interface{} {
// 	if line == "" {
// 		return nil
// 	}
// 	return createBlock("paragraph", nil, []map[string]interface{}{
// 		{"type": "text", "text": line, "styles": map[string]interface{}{}},
// 	})
// }

// // Utility to create a block with common structure
// func createBlock(blockType string, props map[string]interface{}, content interface{}) map[string]interface{} {
// 	if props == nil {
// 		props = defaultProps()
// 	}
// 	return map[string]interface{}{
// 		"id":       generateID(),
// 		"type":     blockType,
// 		"props":    props,
// 		"content":  content,
// 		"children": []map[string]interface{}{},
// 	}
// }

// func isNumberedList(line string) bool {
// 	re := regexp.MustCompile(`^\d+\.\s`)
// 	return re.MatchString(line)
// }

// func isTableRow(line string) bool {
// 	return strings.HasPrefix(line, "|") || strings.Contains(line, "|")
// }

// func isMediaBlock(line string) bool {
// 	return strings.HasPrefix(line, "http") &&
// 		(strings.HasSuffix(line, ".png") || strings.HasSuffix(line, ".jpg") || strings.HasSuffix(line, ".gif") ||
// 			strings.HasSuffix(line, ".svg") || strings.HasSuffix(line, ".webp") || strings.HasSuffix(line, ".mp4"))
// }

// func isAlertBlock(line string) bool {
// 	alertPrefixes := []string{"!warning", "!info", "!success", "!danger"}
// 	for _, prefix := range alertPrefixes {
// 		if strings.HasPrefix(line, prefix) {
// 			return true
// 		}
// 	}
// 	return false
// }

// // Main function
// func main() {
// 	content, err := ioutil.ReadFile("Markdowns/Example.md") // Replace with your file path
// 	if err != nil {
// 		log.Fatalf("Failed to read file: %v", err)
// 	}

// 	blocks, err := parseMarkDownBlock(string(content))
// 	if err != nil {
// 		log.Fatalf("Failed to parse markdown: %v", err)
// 	}

// 	jsonString, err := json.MarshalIndent(blocks, "", "  ")
// 	if err != nil {
// 		log.Fatalf("Failed to generate JSON: %v", err)
// 	}

// 	fmt.Println(string(jsonString))
// }
