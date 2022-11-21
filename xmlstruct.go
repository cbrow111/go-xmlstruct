// Package xmlstruct generates Go structs from multiple XML documents.
package xmlstruct

import (
	"encoding/xml"
	"unicode"

	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

const (
	DefaultFormatSource                 = true
	DefaultHeader                       = "// This file is automatically generated. DO NOT EDIT."
	DefaultTopLevelAttributes           = false
	DefaultIntType                      = "int"
	DefaultNamedTypes                   = false
	DefaultPackageName                  = "main"
	DefaultTimeLayout                   = "2006-01-02T15:04:05Z"
	DefaultUsePointersForOptionalFields = true
)

var (
	DefaultExportNameFunc = TitleFirstRune
	DefaultNameFunc       = IgnoreNamespaceNameFunc
)

// An ExportNameFunc returns the exported Go identifier for the given xml.Name.
type ExportNameFunc func(xml.Name) string

// A NameFunc modifies xml.Names observed in the XML documents.
type NameFunc func(xml.Name) xml.Name

// observeOptions contains options for observing XML documents.
type observeOptions struct {
	nameFunc           NameFunc
	timeLayout         string
	topLevelAttributes bool
	topLevelElements   map[xml.Name]*element
}

// generateOptions contains options for generating Go source.
type generateOptions struct {
	exportNameFunc               ExportNameFunc
	header                       string
	importPackageNames           map[string]struct{}
	intType                      string
	namedTypes                   map[xml.Name]*element
	simpleTypes                  map[xml.Name]struct{}
	usePointersForOptionalFields bool
}

// TitleFirstRune returns name.Local with the initial rune capitalized.
func TitleFirstRune(name xml.Name) string {
	runes := []rune(name.Local)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// The IdentityNameFunc returns name unchanged. The same local name in different
// namespaces will be treated as distinct names.
func IdentityNameFunc(name xml.Name) xml.Name {
	return name
}

// IgnoreNamespaceNameFunc returns name with name.Space cleared. The same local
// name in different namespaces will be treated as identical names.
func IgnoreNamespaceNameFunc(name xml.Name) xml.Name {
	return xml.Name{
		Local: name.Local,
	}
}

// sortedKeys returns the keys of m in order.
func sortedKeys[M ~map[K]V, K constraints.Ordered, V any](m M) []K {
	keys := maps.Keys(m)
	slices.Sort(keys)
	return keys
}
