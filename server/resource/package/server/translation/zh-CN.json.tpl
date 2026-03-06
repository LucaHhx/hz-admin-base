{
  "columns": {
    {{- range $index, $field := .Fields }}
      "{{$field.FieldJson}}": "{{$field.FieldDesc}}",
    {{- end }}
      "ID": "ID",
      "CreatedAt": "创建时间",
      "UpdatedAt": "更新时间",
      "DeletedAt": "删除时间",
      "CreatedBy": "创建人",
      "UpdatedBy": "更新人",
      "DeletedBy": "删除人"
  },
  "enums": {
    {{- $enumFields := false }}
    {{- range $index, $field := .Fields }}
        {{- if eq $field.FieldType "enum" }}
            {{- if $enumFields }},{{- end }}
            {{- $enumFields = true }}
      "{{$field.FieldJson}}": {
            "0": "None",
            "1": "{{$field.FieldDesc}}-1",
            "2": "{{$field.FieldDesc}}-2"
        }
        {{- end }}
    {{- end }}
  },
  "messages": {
  }
}
