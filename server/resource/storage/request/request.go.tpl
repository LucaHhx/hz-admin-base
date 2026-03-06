{{- if .IsAdd}}
// 在结构体中新增如下字段
{{- range .Fields}}
    {{- if ne .FieldSearchType ""}}
        {{- if eq .FieldSearchType "BETWEEN" "NOT BETWEEN"}}
Start{{.FieldName}}  *{{.FieldType}}  `json:"start{{.FieldName}}" form:"start{{.FieldName}}"`
End{{.FieldName}}  *{{.FieldType}}  `json:"end{{.FieldName}}" form:"end{{.FieldName}}"`
        {{- else if eq .FieldSearchType "IN" "NOT IN"}}
            {{- if eq .FieldType "enum" }}
{{.FieldName}}  []int16 `json:"{{.FieldJson}}" form:"{{.FieldJson}}" `
            {{- else if eq .FieldType "string" }}
{{.FieldName}}  []string `json:"{{.FieldJson}}" form:"{{.FieldJson}}" `
            {{- else if or (eq .FieldType "int") (eq .FieldType "int32") (eq .FieldType "int64") }}
{{.FieldName}}  []{{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" `
            {{- else if eq .FieldType "float64" }}
{{.FieldName}}  []float64 `json:"{{.FieldJson}}" form:"{{.FieldJson}}" `
            {{- else if eq .FieldType "bool" }}
{{.FieldName}}  []bool `json:"{{.FieldJson}}" form:"{{.FieldJson}}" `
            {{- else }}
{{.FieldName}}  []{{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" `
            {{- end }}
        {{- else }}
            {{- if eq .FieldType "enum" }}
{{.FieldName}}  int16 `json:"{{.FieldJson}}" form:"{{.FieldJson}}" `
            {{- else if or (eq .FieldType "json") (eq .FieldType "array") }}
{{.FieldName}}  string `json:"{{.FieldJson}}" form:"{{.FieldJson}}" `
            {{- else }}
{{.FieldName}}  *{{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" `
            {{- end }}
        {{- end }}
    {{- end}}
{{- end }}
{{- if .NeedSort}}
Sort  string `json:"sort" form:"sort"`
Order string `json:"order" form:"order"`
{{- end}}
{{- else }}
package request

import (
{{- if not .OnlyTemplate }}
	"{{.Module}}/model/common/request"
	{{ if or .HasSearchTimer .GvaModel}}"time"{{ end }}
{{- end }}
)

type {{.StructName}}Search struct{
{{- if not .OnlyTemplate}}
{{- if .GvaModel }}
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
{{- end }}
{{- range .Fields}}
    {{- if ne .FieldSearchType ""}}
        {{- if eq .FieldSearchType "BETWEEN" "NOT BETWEEN"}}
            {{- if eq .FieldType "date" "datetime" }}
    Start{{.FieldName}}  *time.Time  `json:"start{{.FieldName}}" form:"start{{.FieldName}}"`
    End{{.FieldName}}    *time.Time  `json:"end{{.FieldName}}" form:"end{{.FieldName}}"`
            {{- else }}
    Start{{.FieldName}}  *{{.FieldType}}  `json:"start{{.FieldName}}" form:"start{{.FieldName}}"`
    End{{.FieldName}}    *{{.FieldType}}  `json:"end{{.FieldName}}" form:"end{{.FieldName}}"`
            {{- end}}
        {{- else if eq .FieldSearchType "IN" "NOT IN"}}
            {{- if eq .FieldType "enum" }}
    {{.FieldName}}  []int16 `json:"{{.FieldJson}}" form:"{{.FieldJson}}" `
            {{- else if eq .FieldType "string" }}
    {{.FieldName}}  []string `json:"{{.FieldJson}}" form:"{{.FieldJson}}" `
            {{- else if or (eq .FieldType "int") (eq .FieldType "int32") (eq .FieldType "int64") }}
    {{.FieldName}}  []{{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" `
            {{- else if eq .FieldType "float64" }}
    {{.FieldName}}  []float64 `json:"{{.FieldJson}}" form:"{{.FieldJson}}" `
            {{- else if eq .FieldType "bool" }}
    {{.FieldName}}  []bool `json:"{{.FieldJson}}" form:"{{.FieldJson}}" `
            {{- else }}
    {{.FieldName}}  []{{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" `
            {{- end }}
        {{- else }}
            {{- if eq .FieldType "enum" }}
    {{.FieldName}}  int16 `json:"{{.FieldJson}}" form:"{{.FieldJson}}" `
            {{- else if or (eq .FieldType "json") (eq .FieldType "array") }}
    {{.FieldName}}  string `json:"{{.FieldJson}}" form:"{{.FieldJson}}" `
            {{- else }}
    {{.FieldName}}  *string `json:"{{.FieldJson}}" form:"{{.FieldJson}}" `
            {{- end }}
        {{- end }}
    {{- end}}
{{- end }}
    request.PageInfo
    {{- if .NeedSort}}
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
    {{- end}}
{{- end}}
}
{{- end }}
