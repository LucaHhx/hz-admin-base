{
  "columns": {
    {{- range $index, $field := .Fields }}
      "{{$field.FieldJson}}": "{{$field.FieldName}}",
    {{- end }}
      "ID": "ID",
      "CreatedAt": "Creation Time",
      "UpdatedAt": "Update Time",
      "DeletedAt": "Delete time",
      "CreatedBy": "Create By",
      "UpdatedBy": "Updated By",
      "DeletedBy": "Deleted By"
  },
  "enums": {
    {{- $enumFields := false }}
    {{- range $index, $field := .Fields }}
        {{- if eq $field.FieldType "enum" }}
            {{- if $enumFields }},{{- end }}
            {{- $enumFields = true }}
        "{{$field.FieldJson}}": {
            "0": "None",
            "1": "{{$field.FieldName}}-1",
            "2": "{{$field.FieldName}}-2"
        }
        {{- end }}
    {{- end }}
  },
  "messages": {
  }
}
