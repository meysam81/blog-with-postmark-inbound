package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"sort"
	"strings"
)

type ConfigField struct {
	Name         string
	KoanfTag     string
	Type         string
	DefaultValue interface{}
	Comment      string
}

func main() {
	configPath := "cmd/config/config.go"
	outputPath := ".env.example"

	fields, defaults, err := parseConfigFile(configPath)
	if err != nil {
		fmt.Printf("Error parsing config file: %v\n", err)
		os.Exit(1)
	}

	err = generateEnvExample(outputPath, fields, defaults)
	if err != nil {
		fmt.Printf("Error updating .env.example: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully updated %s\n", outputPath)
}

func parseConfigFile(filename string) ([]ConfigField, map[string]interface{}, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, nil, err
	}

	var fields []ConfigField
	defaults := make(map[string]interface{})

	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.TypeSpec:
			if x.Name.Name == "cfg" {
				if structType, ok := x.Type.(*ast.StructType); ok {
					for _, field := range structType.Fields.List {
						for _, name := range field.Names {
							configField := parseStructField(name.Name, field)
							if configField.KoanfTag != "" {
								fields = append(fields, configField)
							}
						}
					}
				}
			}
		case *ast.CallExpr:
			if sel, ok := x.Fun.(*ast.SelectorExpr); ok {
				if sel.Sel.Name == "Load" {
					parseDefaultValues(x, defaults)
				}
			}
		}
		return true
	})

	return fields, defaults, nil
}

func parseStructField(fieldName string, field *ast.Field) ConfigField {
	configField := ConfigField{
		Name: fieldName,
		Type: getTypeString(field.Type),
	}

	if field.Tag != nil {
		tag := strings.Trim(field.Tag.Value, "`")
		if strings.Contains(tag, "koanf:") {
			parts := strings.Split(tag, "\"")
			for i, part := range parts {
				if strings.Contains(part, "koanf:") && i+1 < len(parts) {
					configField.KoanfTag = parts[i+1]
					break
				}
			}
		}
	}

	configField.Comment = generateComment(configField.KoanfTag)

	return configField
}

func getTypeString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", getTypeString(t.X), t.Sel.Name)
	default:
		return "unknown"
	}
}

func parseDefaultValues(callExpr *ast.CallExpr, defaults map[string]interface{}) {
	for _, arg := range callExpr.Args {
		if call, ok := arg.(*ast.CallExpr); ok {
			if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
				if sel.Sel.Name == "Provider" && len(call.Args) > 0 {
					if comp, ok := call.Args[0].(*ast.CompositeLit); ok {
						parseMapLiteral(comp, defaults)
					}
				}
			}
		}
	}
}

func parseMapLiteral(comp *ast.CompositeLit, defaults map[string]interface{}) {
	for _, elt := range comp.Elts {
		if kv, ok := elt.(*ast.KeyValueExpr); ok {
			if key, ok := kv.Key.(*ast.BasicLit); ok {
				keyStr := strings.Trim(key.Value, "\"")
				value := parseValue(kv.Value)
				defaults[keyStr] = value
			}
		}
	}
}

func parseValue(expr ast.Expr) interface{} {
	switch v := expr.(type) {
	case *ast.BasicLit:
		switch v.Kind {
		case token.STRING:
			return strings.Trim(v.Value, "\"")
		case token.INT:
			return v.Value
		}
	case *ast.Ident:
		switch v.Name {
		case "true":
			return true
		case "false":
			return false
		}
	}
	return nil
}

func generateComment(koanfTag string) string {
	comments := map[string]string{
		"port":                           "Port for the web server to listen on",
		"base-url":                       "Base URL where Tarzan is accessible (used for generating links)",
		"auth.username":                  "Username for webhook authentication",
		"auth.password":                  "Password for webhook authentication",
		"redis.host":                     "Redis server hostname",
		"redis.port":                     "Redis server port",
		"redis.db":                       "Redis database number",
		"redis.password":                 "Redis authentication password",
		"redis.ssl":                      "Enable SSL/TLS for Redis connection",
		"datastore":                      "Data storage backend (sqlite or redis)",
		"filestore":                      "File storage backend (filesystem or redis)",
		"dir.db":                         "Path to the SQLite database file",
		"dir.storage":                    "Path to the storage directory for uploaded files",
		"endpoints.rss":                  "RSS endpoint path",
		"metrics.enabled":                "Enable metrics collection",
		"metrics.auth.enabled":           "Enable authentication for metrics endpoint",
		"metrics.auth.username":          "Username for metrics endpoint authentication",
		"metrics.auth.password":          "Password for metrics endpoint authentication",
		"dangerously-accept-all-senders": "Accept webhooks from any sender (dangerous in production)",
	}

	if comment, exists := comments[koanfTag]; exists {
		return comment
	}
	return fmt.Sprintf("Configuration for %s", koanfTag)
}

func generateEnvExample(filename string, fields []ConfigField, defaults map[string]interface{}) error {
	var content strings.Builder

	content.WriteString("# =============================================================================\n")
	content.WriteString("# TARZAN CONFIGURATION\n")
	content.WriteString("# =============================================================================\n")
	content.WriteString("# Copy this file to .env and update the values according to your setup\n\n")

	sections := map[string][]ConfigField{
		"SERVER CONFIGURATION":   {},
		"WEBHOOK AUTHENTICATION": {},
		"REDIS CONFIGURATION":    {},
		"STORAGE CONFIGURATION":  {},
		"DIRECTORY PATHS":        {},
		"ENDPOINTS":              {},
		"METRICS":                {},
		"SECURITY":               {},
	}

	for _, field := range fields {
		switch {
		case strings.Contains(field.KoanfTag, "port") || strings.Contains(field.KoanfTag, "base-url"):
			sections["SERVER CONFIGURATION"] = append(sections["SERVER CONFIGURATION"], field)
		case strings.Contains(field.KoanfTag, "auth") && !strings.Contains(field.KoanfTag, "metrics"):
			sections["WEBHOOK AUTHENTICATION"] = append(sections["WEBHOOK AUTHENTICATION"], field)
		case strings.Contains(field.KoanfTag, "redis"):
			sections["REDIS CONFIGURATION"] = append(sections["REDIS CONFIGURATION"], field)
		case strings.Contains(field.KoanfTag, "store"):
			sections["STORAGE CONFIGURATION"] = append(sections["STORAGE CONFIGURATION"], field)
		case strings.Contains(field.KoanfTag, "dir"):
			sections["DIRECTORY PATHS"] = append(sections["DIRECTORY PATHS"], field)
		case strings.Contains(field.KoanfTag, "endpoints"):
			sections["ENDPOINTS"] = append(sections["ENDPOINTS"], field)
		case strings.Contains(field.KoanfTag, "metrics"):
			sections["METRICS"] = append(sections["METRICS"], field)
		case strings.Contains(field.KoanfTag, "dangerous"):
			sections["SECURITY"] = append(sections["SECURITY"], field)
		default:
			sections["SERVER CONFIGURATION"] = append(sections["SERVER CONFIGURATION"], field)
		}
	}

	sectionOrder := []string{
		"SERVER CONFIGURATION",
		"WEBHOOK AUTHENTICATION",
		"REDIS CONFIGURATION",
		"STORAGE CONFIGURATION",
		"DIRECTORY PATHS",
		"ENDPOINTS",
		"METRICS",
		"SECURITY",
	}

	for _, sectionName := range sectionOrder {
		sectionFields := sections[sectionName]
		if len(sectionFields) == 0 {
			continue
		}

		content.WriteString("# =============================================================================\n")
		content.WriteString(fmt.Sprintf("# %s\n", sectionName))
		content.WriteString("# =============================================================================\n")

		sort.Slice(sectionFields, func(i, j int) bool {
			return sectionFields[i].KoanfTag < sectionFields[j].KoanfTag
		})

		for _, field := range sectionFields {
			content.WriteString(fmt.Sprintf("# %s\n", field.Comment))

			envVar := "TARZAN_" + strings.ToUpper(strings.ReplaceAll(field.KoanfTag, ".", "_"))

			var defaultValue string
			if def, exists := defaults[field.KoanfTag]; exists {
				defaultValue = formatDefaultValue(def)
			} else {
				defaultValue = getPlaceholderValue(field.Type, field.KoanfTag)
			}

			content.WriteString(fmt.Sprintf("%s=%s\n\n", envVar, defaultValue))
		}
	}

	finalContent := strings.TrimSpace(content.String()) + "\n"

	return os.WriteFile(filename, []byte(finalContent), 0644)
}

func formatDefaultValue(value interface{}) string {
	switch v := value.(type) {
	case string:
		if v == "" {
			return ""
		}
		return v
	case bool:
		if v {
			return "true"
		}
		return "false"
	case int, int32, int64, uint, uint32, uint64:
		return fmt.Sprintf("%v", v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

func getPlaceholderValue(fieldType, koanfTag string) string {
	switch fieldType {
	case "string":
		if strings.Contains(koanfTag, "username") {
			return "your_username"
		} else if strings.Contains(koanfTag, "password") {
			return "your_password"
		} else if strings.Contains(koanfTag, "host") {
			return "localhost"
		} else if strings.Contains(koanfTag, "url") {
			return "http://example.com"
		} else if strings.Contains(koanfTag, "path") || strings.Contains(koanfTag, "dir") {
			return "path/to/directory"
		}
		return "your_value"
	case "int", "uint32":
		return "8000"
	case "uint8":
		return "0"
	case "bool":
		return "false"
	default:
		return "your_value"
	}
}
