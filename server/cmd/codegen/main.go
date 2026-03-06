package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"gorm.io/gorm/schema"

	"hab/core"
	"hab/global"
	modelsystem "hab/model/system"
	"hab/model/system/request"
	servicesystem "hab/service/system"
)

// ColumnInfo represents database column metadata
type ColumnInfo struct {
	ColumnName    string `json:"column_name"`
	DataType      string `json:"data_type"`
	DataTypeLong  string `json:"data_type_long"`
	ColumnComment string `json:"column_comment"`
	PrimaryKey    int    `json:"primary_key"`
	OrdinalPos    int    `json:"ordinal_position"`
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	tableName := os.Args[1]
	pkg := "business"
	description := ""

	if len(os.Args) >= 3 {
		pkg = os.Args[2]
	}
	if len(os.Args) >= 4 {
		description = os.Args[3]
	}

	fmt.Printf("=== HAB Code Generator ===\n")
	fmt.Printf("Table: %s | Package: %s | Description: %s\n\n", tableName, pkg, description)

	// Change to server directory for config file access
	if err := os.Chdir("../.."); err != nil {
		fmt.Printf("Error changing to server directory: %v\n", err)
		os.Exit(1)
	}

	// Initialize HAB using core.InitServer
	fmt.Println("Initializing HAB...")
	core.InitServer()

	// Get table columns
	fmt.Println("Reading table structure...")
	columns, err := getTableColumns(tableName)
	if err != nil {
		fmt.Printf("Error reading columns: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Found %d columns\n", len(columns))

	// Build AutoCode request
	autoCode := buildAutoCodeRequest(tableName, pkg, description, columns)

	// Call Pretreatment to set Module and other required fields
	if err := autoCode.Pretreatment(); err != nil {
		fmt.Printf("Error preprocessing request: %v\n", err)
		os.Exit(1)
	}

	// Check for duplicates and auto-delete them
	if servicesystem.AutocodeHistory.Repeat(autoCode.BusinessDB, autoCode.StructName, autoCode.Abbreviation, autoCode.Package) {
		fmt.Printf("Found duplicate struct/abbreviation: %s/%s - removing old records...\n", autoCode.StructName, autoCode.Abbreviation)
		// Delete all duplicate records using GORM (hard delete)
		if err := global.HAB_DB.Unscoped().Where("struct_name = ? OR abbreviation = ?", autoCode.StructName, autoCode.Abbreviation).Delete(&modelsystem.SysAutoCodeHistory{}).Error; err != nil {
			fmt.Printf("Warning: Could not delete duplicate records: %v\n", err)
			fmt.Println("Continuing anyway...")
		} else {
			fmt.Println("Old duplicate records removed successfully")
		}
	}

	// Generate code
	fmt.Println("Generating code...")
	ctx := context.Background()
	if err := servicesystem.AutoCodeTemplate.Create(ctx, *autoCode); err != nil {
		fmt.Printf("Error generating code: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n=== Code Generation Complete ===")
	printGeneratedFiles(autoCode)
	fmt.Println("\nIMPORTANT: Restart backend service to load new routes and APIs!")
}

func printUsage() {
	fmt.Println("Usage: go run . <table_name> [package] [description]")
	fmt.Println("")
	fmt.Println("Arguments:")
	fmt.Println("  table_name    Database table name (required)")
	fmt.Println("  package       Target package (default: system)")
	fmt.Println("  description   Chinese description (default: auto-generated)")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("  go run . b_agent system \"代理商管理\"")
	fmt.Println("  go run . sys_user system \"用户管理\"")
	fmt.Println("  go run . orders example \"订单管理\"")
}

func getTableColumns(tableName string) ([]ColumnInfo, error) {
	dbName := global.HAB_CONFIG.Mysql.Dbname

	query := `
		SELECT
			c.COLUMN_NAME as column_name,
			c.DATA_TYPE as data_type,
			CASE c.DATA_TYPE
				WHEN 'longtext' THEN c.CHARACTER_MAXIMUM_LENGTH
				WHEN 'varchar' THEN c.CHARACTER_MAXIMUM_LENGTH
				WHEN 'double' THEN CONCAT(c.NUMERIC_PRECISION, ',', c.NUMERIC_SCALE)
				WHEN 'decimal' THEN CONCAT(c.NUMERIC_PRECISION, ',', c.NUMERIC_SCALE)
				WHEN 'int' THEN c.NUMERIC_PRECISION
				WHEN 'bigint' THEN c.NUMERIC_PRECISION
				ELSE ''
			END AS data_type_long,
			c.COLUMN_COMMENT as column_comment,
			CASE WHEN kcu.COLUMN_NAME IS NOT NULL THEN 1 ELSE 0 END AS primary_key,
			c.ORDINAL_POSITION
		FROM INFORMATION_SCHEMA.COLUMNS c
		LEFT JOIN INFORMATION_SCHEMA.KEY_COLUMN_USAGE kcu
			ON c.TABLE_SCHEMA = kcu.TABLE_SCHEMA
			AND c.TABLE_NAME = kcu.TABLE_NAME
			AND c.COLUMN_NAME = kcu.COLUMN_NAME
			AND kcu.CONSTRAINT_NAME = 'PRIMARY'
		WHERE c.TABLE_NAME = ? AND c.TABLE_SCHEMA = ?
		ORDER BY c.ORDINAL_POSITION`

	var columns []ColumnInfo
	if err := global.HAB_DB.Raw(query, tableName, dbName).Scan(&columns).Error; err != nil {
		return nil, err
	}

	return columns, nil
}

func buildAutoCodeRequest(tableName, pkg, description string, columns []ColumnInfo) *request.AutoCode {
	// Generate naming conventions
	baseName := tableName
	for _, prefix := range []string{"sys_", "tb_", "t_", "b_"} {
		if strings.HasPrefix(baseName, prefix) {
			baseName = strings.TrimPrefix(baseName, prefix)
			break
		}
	}

	structName := toPascalCase(baseName)
	abbreviation := toCamelCase(baseName)
	packageName := toSnakeCase(baseName)
	humpPackageName := toCamelCase(baseName)

	if description == "" {
		description = inferDescription(baseName)
	}

	// Check if table has HAB_MODEL fields (before filtering)
	gvaModel := hasGvaModelColumns(columns)

	// Build fields from database columns
	// Skip HAB_MODEL fields to avoid duplication
	// Note: created_by, updated_by, deleted_by are auto-added by template when AutoCreateResource=true
	gvaModelFields := map[string]string{
		"id":         "ID",
		"created_at": "CreatedAt",
		"updated_at": "UpdatedAt",
		"deleted_at": "DeletedAt",
		"created_by": "CreatedBy",
		"updated_by": "UpdatedBy",
		"deleted_by": "DeletedBy",
	}

	fields := make([]*request.AutoCodeField, 0, len(columns))
	var primaryField *request.AutoCodeField

	for _, col := range columns {
		// Skip HAB_MODEL fields
		if _, exists := gvaModelFields[col.ColumnName]; exists {
			if col.PrimaryKey == 1 {
				// Still set primaryField for HAB_MODEL tables
				primaryField = &request.AutoCodeField{
					FieldName:    "ID",
					FieldDesc:    "主键ID",
					FieldType:    "uint",
					FieldJson:    "ID",
					DataTypeLong: "20",
					Comment:      "主键ID",
					ColumnName:   "id",
				}
			}
			continue
		}

		field := &request.AutoCodeField{
			ColumnName:   col.ColumnName,
			FieldJson:    toCamelCase(col.ColumnName),
			FieldName:    toPascalCase(col.ColumnName),
			FieldDesc:    inferFieldDesc(col.ColumnName, col.ColumnComment),
			DataType:     col.DataType,
			DataTypeLong: col.DataTypeLong,
			Comment:      col.ColumnComment,
			PrimaryKey:   col.PrimaryKey == 1,
			Require:      col.PrimaryKey == 1 || isRequiredField(col.ColumnName),
		}

		field.FieldType = mapFieldType(col.DataType, col.ColumnName)
		field.FieldSearchType = mapSearchType(field.FieldType)
		setFieldVisibility(field, col.ColumnName)

		if col.PrimaryKey == 1 {
			primaryField = field
		}

		fields = append(fields, field)
	}

	// Set PrimaryField for GvaModel if not found
	if gvaModel && primaryField == nil {
		primaryField = &request.AutoCodeField{
			FieldName:    "ID",
			FieldDesc:    "主键ID",
			FieldType:    "uint",
			FieldJson:    "ID",
			DataTypeLong: "20",
			Comment:      "主键ID",
			ColumnName:   "id",
		}
	}

	return &request.AutoCode{
		TableName:           tableName,
		Package:             pkg,
		StructName:          structName,
		PackageName:         packageName,
		Description:         description,
		Abbreviation:        abbreviation,
		HumpPackageName:     humpPackageName,
		BusinessDB:          "",
		GvaModel:            gvaModel,
		AutoMigrate:         true,
		AutoCreateResource:  true,
		AutoCreateApiToSql:  true,
		AutoCreateMenuToSql: true,
		AutoCreateBtnAuth:   true,
		OnlyTemplate:        false,
		IsTree:              false,
		IsAdd:               false,
		Fields:              fields,
		GenerateWeb:         true,
		GenerateServer:      true,
		PrimaryField:        primaryField,
	}
}

// Naming conversion functions
func toPascalCase(s string) string {
	namer := schema.NamingStrategy{}
	return namer.SchemaName(s)
}

func toCamelCase(s string) string {
	pascal := toPascalCase(s)
	if len(pascal) == 0 {
		return ""
	}
	return strings.ToLower(pascal[:1]) + pascal[1:]
}

func toSnakeCase(s string) string {
	var result []rune
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, r)
	}
	return strings.ToLower(string(result))
}

func inferDescription(name string) string {
	descMap := map[string]string{
		"user":     "用户",
		"agent":    "代理商",
		"customer": "客户",
		"role":     "角色",
		"menu":     "菜单",
		"api":      "接口",
		"dict":     "字典",
		"order":    "订单",
		"product":  "产品",
		"goods":    "商品",
		"category": "分类",
	}

	if desc, ok := descMap[name]; ok {
		return desc + "管理"
	}
	return name + "管理"
}

func inferFieldDesc(columnName, comment string) string {
	if comment != "" {
		return comment
	}

	descMap := map[string]string{
		"id":         "ID",
		"name":       "名称",
		"uuid":       "编号",
		"amount":     "金额",
		"status":     "状态",
		"state":      "状态",
		"note":       "备注",
		"remark":     "备注",
		"created_at": "创建时间",
		"updated_at": "更新时间",
		"deleted_at": "删除时间",
		"created_by": "创建者",
		"updated_by": "更新者",
		"deleted_by": "删除者",
	}

	if desc, ok := descMap[columnName]; ok {
		return desc
	}

	// Handle amount fields
	if strings.Contains(columnName, "amount") {
		if strings.Contains(columnName, "settle") {
			return "待结算金额"
		}
		if strings.Contains(columnName, "freeze") {
			return "冻结金额"
		}
		if strings.Contains(columnName, "sum") || strings.Contains(columnName, "total") {
			return "总金额"
		}
		if strings.Contains(columnName, "avail") {
			return "可用金额"
		}
		return "金额"
	}

	return columnName
}

func mapFieldType(dbType, columnName string) string {
	switch strings.ToLower(dbType) {
	case "varchar", "char":
		return "string"
	case "text", "longtext", "mediumtext":
		if strings.Contains(columnName, "content") || strings.Contains(columnName, "note") || strings.Contains(columnName, "remark") {
			return "richtext"
		}
		return "string"
	case "int":
		return "int"
	case "tinyint":
		if strings.Contains(columnName, "status") || strings.Contains(columnName, "state") {
			return "bool"
		}
		return "int64"
	case "bigint":
		return "int64"
	case "decimal", "double", "float":
		return "float64"
	case "datetime", "timestamp":
		return "datetime"
	case "date":
		return "date"
	case "json":
		return "json"
	default:
		return "string"
	}
}

func mapSearchType(fieldType string) string {
	switch fieldType {
	case "string", "richtext":
		return "like"
	case "bool", "int64", "float64":
		return "="
	case "datetime", "date":
		return "between"
	default:
		return "="
	}
}

func setFieldVisibility(field *request.AutoCodeField, columnName string) {
	field.Form = true
	field.Table = true
	field.Desc = true

	// Hide system fields in form
	systemFields := []string{"created_at", "updated_at", "deleted_at", "created_by", "updated_by", "deleted_by"}
	for _, sys := range systemFields {
		if columnName == sys || strings.HasSuffix(columnName, "_at") || strings.HasSuffix(columnName, "_time") {
			field.Form = false
			if columnName != "created_at" && columnName != "updated_at" {
				field.Table = false
			}
			return
		}
	}

	if columnName == "id" {
		field.Form = false
	}
}

func isRequiredField(columnName string) bool {
	requiredPatterns := []string{"_id", "name", "code", "uuid"}
	for _, p := range requiredPatterns {
		if strings.Contains(columnName, p) {
			return true
		}
	}
	return false
}

// hasGvaModelColumns checks if the original table columns contain HAB_MODEL fields
func hasGvaModelColumns(columns []ColumnInfo) bool {
	gvaFields := map[string]bool{
		"id": true, "created_at": true, "updated_at": true, "deleted_at": true,
	}

	for _, col := range columns {
		if gvaFields[col.ColumnName] {
			return true
		}
	}
	return false
}

func printGeneratedFiles(autoCode *request.AutoCode) {
	fmt.Println("\n=== Generated Files ===")

	if autoCode.GenerateServer {
		fmt.Println("Backend:")
		fmt.Printf("  server/model/%s/%s.go\n", autoCode.Package, autoCode.HumpPackageName)
		fmt.Printf("  server/model/%s/request/%s.go\n", autoCode.Package, autoCode.HumpPackageName)
		fmt.Printf("  server/service/%s/%s.go\n", autoCode.Package, autoCode.HumpPackageName)
		fmt.Printf("  server/api/v1/%s/%s.go\n", autoCode.Package, autoCode.HumpPackageName)
		fmt.Printf("  server/router/%s/%s.go\n", autoCode.Package, autoCode.HumpPackageName)
	}

	if autoCode.GenerateWeb {
		fmt.Println("Frontend:")
		fmt.Printf("  web/src/api/%s/%s.js\n", autoCode.Package, autoCode.PackageName)
		fmt.Printf("  web/src/view/%s/%s/%s.vue\n", autoCode.Package, autoCode.PackageName, autoCode.PackageName)
		fmt.Printf("  web/src/view/%s/%s/%sForm.vue\n", autoCode.Package, autoCode.PackageName, autoCode.PackageName)
	}
}
