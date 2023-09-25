package outputs

import (
	"bytes"
	"fmt"
	"path"
	"text/template"

	"github.com/evertras/code-todos/internal/todos"
)

var markdownTableTemplate = template.Must(template.New("markdownTable").Parse(`| Filename | Package | Line | Text |
| -------- | ------- | ---- | ---- |
{{ range . }}| [{{ .Filename }}]({{ .Link }}) | {{ .PackageName }} | {{ .Line }} | {{ .Text }} |
{{ end -}}`))

type MarkdownTableConfig struct {
	LinkPrefix string
}

func MarkdownTable(list []todos.Todo, config MarkdownTableConfig) (string, error) {
	data := make([]struct {
		todos.Todo
		Link string
	}, len(list))

	for i, t := range list {
		data[i].Todo = t
		data[i].Link = fmt.Sprintf("%s#L%d", path.Join(config.LinkPrefix, t.Filename), t.Line)
	}

	buf := new(bytes.Buffer)
	err := markdownTableTemplate.ExecuteTemplate(buf, "markdownTable", data)

	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
