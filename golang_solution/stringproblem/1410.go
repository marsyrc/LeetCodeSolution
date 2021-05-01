package stringproblem

import "strings"

//strings.NewReplacer
func entityParser(s string) string {
	return strings.NewReplacer(`&quot;`, `"`, `&apos;`, `'`, `&gt;`, `>`, `&lt;`, `<`, `&frasl;`, `/`, `&amp;`, `&`).Replace(s)
}
