{{- $db := "" }}
{{- if eq .BusinessDB "" }}
 {{- $db = "global.GVA_DB" }}
{{- else}}
 {{- $db =  printf "global.MustGetGlobalDBByDBName(\"%s\")" .BusinessDB   }}
{{- end}}

{{- if .IsAdd}}

// Get{{.StructName}}InfoList 新增搜索语句
        {{- range .Fields}}
            {{- if .FieldSearchType}}
                {{- if or (eq .FieldType "enum") (eq .FieldType "pictures") (eq .FieldType "picture") (eq .FieldType "video") (eq .FieldType "json") }}
if info.{{.FieldName}} != "" {
        {{- if or (eq .FieldType "enum") }}
    db = db.Where("{{.ColumnName}} {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+ {{ end }}*info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
        {{- else}}
// Data type is complex, please implement complex type query business according to business requirements
        {{- end}}
}
    {{- else if eq .FieldSearchType "BETWEEN" "NOT BETWEEN"}}
if info.Start{{.FieldName}} != nil && info.End{{.FieldName}} != nil {
    db = db.Where("{{.ColumnName}} {{.FieldSearchType}} ? AND ? ",info.Start{{.FieldName}},info.End{{.FieldName}})
}
    {{- else}}
if info.{{.FieldName}} != nil{{- if eq .FieldType "string" }} && *info.{{.FieldName}} != ""{{- end }} {
    db = db.Where("{{.ColumnName}} {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}*info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
}
            {{- end }}
        {{- end }}
    {{- end }}


// Get{{.StructName}}InfoList 新增排序语句 请自行在搜索语句中添加orderMap内容
       {{- range .Fields}}
            {{- if .Sort}}
orderMap["{{.ColumnName}}"] = true
         	{{- end}}
       {{- end}}


{{- if .HasDataSource }}
//  Get{{.StructName}}DataSource()方法新增关联语句
	{{range $key, $value := .DataSourceMap}}
{{$key}} := make([]map[string]any, 0)
{{ $dataDB := "" }}
{{- if eq $value.DBName "" }}
{{ $dataDB = $db }}
{{- else}}
{{ $dataDB = printf "global.MustGetGlobalDBByDBName(\"%s\")" $value.DBName }}
{{- end}}
{{$dataDB}}.Table("{{$value.Table}}"){{- if $value.HasDeletedAt}}.Where("deleted_at IS NULL"){{ end }}.Select("{{$value.Label}} as label,{{$value.Value}} as value").Scan(&{{$key}})
res["{{$key}}"] = {{$key}}
	{{- end }}
{{- end }}
{{- else}}
package {{.Package}}

import (
    "forge-basic/global"
    "forge-basic/enum"
    "github.com/pkg/errors"
{{- if not .OnlyTemplate }}
    "github.com/gin-gonic/gin"
    "forge-basic/model/common/request"
    "forge-basic/utils"
	"{{.Module}}/model/{{.Package}}"
	{{- if not .IsTree}}
    //{{.Package}}Req "{{.Module}}/model/{{.Package}}/request"
    {{- else }}
    "{{.Module}}/utils"
    "errors"
    {{- end }}
    {{- if .AutoCreateResource }}
    "gorm.io/gorm"
    {{- end}}
{{- end }}
)

type {{.StructName}}Service struct {}

{{- if not .OnlyTemplate }}
// Create{{.StructName}} 创建{{.Description}}记录
// Author [yourname](https://github.com/yourname)
func ({{.Abbreviation}}Service *{{.StructName}}Service) Create{{.StructName}}({{.Abbreviation}} *{{.Package}}.{{.StructName}},c *gin.Context) (err error) {
	err = {{$db}}.Create({{.Abbreviation}}).Error
	return err
}

// Import{{.StructName}} 导入{{.Description}}记录
// Author [yourname](https://github.com/yourname)
func ({{.Abbreviation}}Service *{{.StructName}}Service) Import{{.StructName}}(tep enum.ImportType, list []{{.Package}}.{{.StructName}},c *gin.Context) (err error) {
		switch tep {
    	case enum.ImportType_Full:
            return {{$db}}.Transaction(func(tx *gorm.DB) error {
                // 截断表（清空表并重置自增ID）
                if err := tx.Exec("TRUNCATE TABLE {{.TableName}}").Error; err != nil {
                    return err
                }
                // 批量写入
                if err := tx.Create(&list).Error; err != nil {
                    return err
                }
                return nil
            })
    	case enum.ImportType_Append:
    		return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
    			for i, _ := range list {
    				nowTime := global.NowMySQLTime()
    				list[i].CreatedAt = nowTime
    				list[i].UpdatedAt = nowTime
    				list[i].ID = 0
    			}
    			// 批量写入
    			if err := tx.Create(&list).Error; err != nil {
    				return err
    			}
    			return nil
    		})
    	default:
    		return errors.New("Invalid import type")
    	}
}

// Delete{{.StructName}} 删除{{.Description}}记录
// Author [yourname](https://github.com/yourname)
func ({{.Abbreviation}}Service *{{.StructName}}Service)Delete{{.StructName}}({{.PrimaryField.FieldJson}} string{{- if .AutoCreateResource -}},userID uint{{- end -}},c *gin.Context) (err error) {
	{{- if .IsTree }}
       var count int64
	   err = {{$db}}.Find(&{{.Package}}.{{.StructName}}{},"parent_id = ?",{{.PrimaryField.FieldJson}}).Count(&count).Error
	   if count > 0 {
           return errors.New("This node has child nodes and cannot be deleted")
       }
       if err != nil {
           return err
       }
	{{- end }}

	{{- if .AutoCreateResource }}
	err = {{$db}}.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&{{.Package}}.{{.StructName}}{}).Where("{{.PrimaryField.ColumnName}} = ?", {{.PrimaryField.FieldJson}}).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&{{.Package}}.{{.StructName}}{},"{{.PrimaryField.ColumnName}} = ?",{{.PrimaryField.FieldJson}}).Error; err != nil {
              return err
        }
        return nil
	})
    {{- else }}
	err = {{$db}}.Delete(&{{.Package}}.{{.StructName}}{},"{{.PrimaryField.ColumnName}} = ?",{{.PrimaryField.FieldJson}}).Error
	{{- end }}
	return err
}

// Delete{{.StructName}}ByIds 批量删除{{.Description}}记录
// Author [yourname](https://github.com/yourname)
func ({{.Abbreviation}}Service *{{.StructName}}Service)Delete{{.StructName}}ByIds({{.PrimaryField.FieldJson}}s []string {{- if .AutoCreateResource }},deleted_by uint{{- end}},c *gin.Context) (err error) {
	{{- if .AutoCreateResource }}
	err = {{$db}}.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&{{.Package}}.{{.StructName}}{}).Where("{{.PrimaryField.ColumnName}} in ?", {{.PrimaryField.FieldJson}}s).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("{{.PrimaryField.ColumnName}} in ?", {{.PrimaryField.FieldJson}}s).Delete(&{{.Package}}.{{.StructName}}{}).Error; err != nil {
            return err
        }
        return nil
    })
    {{- else}}
	err = {{$db}}.Delete(&[]{{.Package}}.{{.StructName}}{},"{{.PrimaryField.ColumnName}} in ?",{{.PrimaryField.FieldJson}}s).Error
    {{- end}}
	return err
}

// Update{{.StructName}} 更新{{.Description}}记录
// Author [yourname](https://github.com/yourname)
func ({{.Abbreviation}}Service *{{.StructName}}Service)Update{{.StructName}}({{.Abbreviation}} {{.Package}}.{{.StructName}},c *gin.Context) (err error) {
	return {{$db}}.Save(&{{.Abbreviation}}).Error
}

// Get{{.StructName}} 根据{{.PrimaryField.FieldJson}}获取{{.Description}}记录
// Author [yourname](https://github.com/yourname)
func ({{.Abbreviation}}Service *{{.StructName}}Service)Get{{.StructName}}({{.PrimaryField.FieldJson}} string,c *gin.Context) ({{.Abbreviation}} {{.Package}}.{{.StructName}}, err error) {
	// err = {{$db}}.Where("{{.PrimaryField.ColumnName}} = ?", {{.PrimaryField.FieldJson}}).First(&{{.Abbreviation}}).Error
    db := {{$db}}.Model(&{{.Package}}.{{.StructName}}{})
    err = db.Where("{{.PrimaryField.ColumnName}} = ?", {{.PrimaryField.FieldJson}}).First(&{{.Abbreviation}}).Error
	return
}


{{- if .IsTree }}
// Get{{.StructName}}InfoList 分页获取{{.Description}}记录,Tree模式下不添加分页和搜索
// Author [yourname](https://github.com/yourname)
func ({{.Abbreviation}}Service *{{.StructName}}Service)Get{{.StructName}}InfoList(c *gin.Context) (list []*{{.Package}}.{{.StructName}},err error) {
    // 创建db
	db := {{$db}}.Model(&{{.Package}}.{{.StructName}}{})
    var {{.Abbreviation}}s []*{{.Package}}.{{.StructName}}

	err = db.Find(&{{.Abbreviation}}s).Error

	return utils.BuildTree({{.Abbreviation}}s), err
}
{{- else }}
// Get{{.StructName}}InfoList 分页获取{{.Description}}记录
// Author [yourname](https://github.com/yourname)
func ({{.Abbreviation}}Service *{{.StructName}}Service)Get{{.StructName}}InfoList(info request.QueryInfo,c *gin.Context) (list []{{.Package}}.{{.StructName}}, total int64, err error) {
    // Create db
	db := {{$db}}.Model(&{{.Package}}.{{.StructName}}{})
    var {{.Abbreviation}}s []{{.Package}}.{{.StructName}}
    total, err = utils.TableQuery(db, info,c)
	if err != nil {
		return
	}
	err = db.Select(utils.GetColumns(db, c)).Find(&{{.Abbreviation}}s).Error
	return  {{.Abbreviation}}s, total, err
}

{{- end }}

{{- if .HasDataSource }}
func ({{.Abbreviation}}Service *{{.StructName}}Service)Get{{.StructName}}DataSource() (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)
	{{range $key, $value := .DataSourceMap}}
	   {{$key}} := make([]map[string]any, 0)
	   {{ $dataDB := "" }}
	   {{- if eq $value.DBName "" }}
       {{ $dataDB = $db }}
       {{- else}}
       {{ $dataDB = printf "global.MustGetGlobalDBByDBName(\"%s\")" $value.DBName }}
       {{- end}}
       {{$dataDB}}.Table("{{$value.Table}}"){{- if $value.HasDeletedAt}}.Where("deleted_at IS NULL"){{ end }}.Select("{{$value.Label}} as label,{{$value.Value}} as value").Scan(&{{$key}})
	   res["{{$key}}"] = {{$key}}
	{{- end }}
	return
}
{{- end }}
{{- end }}
func ({{.Abbreviation}}Service *{{.StructName}}Service)Get{{.StructName}}Public() {
    // This method is for getting data source defined data
    // Please implement it yourself
}
{{- end }}
