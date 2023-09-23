package outputs

import (
	"bytes"
	"text/template"

	"github.com/evertras/code-todos/internal/todos"
)

var markdownTableTemplate = template.Must(template.New("markdownTable").Parse(`| Filename | Package | Line | Text |
| -------- | ------- | ---- | ---- |
{{ range . }}| [{{ .Filename }}](./{{ .Filename }}#{{ .Line }}) | {{ .PackageName }} | {{ .Line }} | {{ .Text }} |
{{ end -}}`))

func MarkdownTable(todos []todos.Todo) (string, error) {
	buf := new(bytes.Buffer)
	err := markdownTableTemplate.ExecuteTemplate(buf, "markdownTable", todos)

	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
