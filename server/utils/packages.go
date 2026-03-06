package utils

import (
	"fmt"
	"go/types"
	"hab/global"
	"reflect"
	"strings"
	"testing"
	"unicode"

	"golang.org/x/tools/go/packages"
)

type TableInfo struct {
	Name      string
	TableName string
	Fields    []FieldInfo
}

type FieldInfo struct {
	Name      string
	Type      string
	JSONTag   string
	ColumnTag string
	Comment   string
}

var predefinedStructs = map[string][]FieldInfo{
	"HAB_MODEL": {
		{Name: "ID", Type: "uint", JSONTag: "ID", ColumnTag: "id"},
		{Name: "CreatedAt", Type: "time.Time", JSONTag: "CreatedAt", ColumnTag: "created_at", Comment: "创建日期"},
		{Name: "UpdatedAt", Type: "time.Time", JSONTag: "UpdatedAt", ColumnTag: "updated_at", Comment: "修改日期"},
		{Name: "DeletedAt", Type: "time.Time", JSONTag: "DeletedAt", ColumnTag: "deleted_at", Comment: "删除日期"},
	},
	"HAB_MODEL_NOD": {
		{Name: "ID", Type: "uint", JSONTag: "ID", ColumnTag: "id"},
		{Name: "CreatedAt", Type: "time.Time", JSONTag: "CreatedAt", ColumnTag: "created_at", Comment: "创建日期"},
		{Name: "UpdatedAt", Type: "time.Time", JSONTag: "UpdatedAt", ColumnTag: "updated_at", Comment: "修改日期"},
	},
	"HAB_MODEL_NOID": {
		{Name: "CreatedAt", Type: "time.Time", JSONTag: "CreatedAt", ColumnTag: "created_at", Comment: "创建日期"},
		{Name: "UpdatedAt", Type: "time.Time", JSONTag: "UpdatedAt", ColumnTag: "updated_at", Comment: "修改日期"},
		{Name: "DeletedAt", Type: "time.Time", JSONTag: "DeletedAt", ColumnTag: "deleted_at", Comment: "删除日期"},
	},
}

var typedPredefinedStructs = map[string]string{
	"gtype.Data": "object",
	"enum.":      "string",
	"gtype.Map":  "object",
	"gtype.Arr":  "array",
}

func GetTableInfo(name string) TableInfo {
	var tables map[string]TableInfo
	switch name[0] {
	case 'b':
		tables = CollectTables(fmt.Sprintf("%s/%s/model/business", global.HAB_CONFIG.AutoCode.Root, global.HAB_CONFIG.AutoCode.Server))
	case 'v':
		tables = CollectTables(fmt.Sprintf("%s/%s/model/view", global.HAB_CONFIG.AutoCode.Root, global.HAB_CONFIG.AutoCode.Server))

	}
	return tables[strings.ToLower(name)]
}

func TestCollectTables(t *testing.T) {
	tables := CollectTables("./model/business")
	for s, table := range tables {
		fmt.Println("Path:", s)
		fmt.Println("Name:", table.Name)
		fmt.Println("TableName:", table.TableName)
		fmt.Println("Fields:")
		for _, field := range table.Fields {
			fmt.Printf("  Name: %s, Type: %s, JSONTag: %s, ColumnTag: %s, Comment: %s\n",
				field.Name, field.Type, field.JSONTag, field.ColumnTag, field.Comment)
		}
	}
}

func CollectTables(dir string) map[string]TableInfo {
	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedName | packages.NeedFiles | packages.NeedDeps,
		Dir:  dir,
	}

	pkgs, err := packages.Load(cfg, "./...")
	if err != nil {
		panic(err)
	}

	parsed := make(map[string]TableInfo)
	visited := make(map[string]bool)

	for _, pkg := range pkgs {
		scope := pkg.Types.Scope()
		for _, name := range scope.Names() {
			obj := scope.Lookup(name)
			typeName, ok := obj.(*types.TypeName)
			if !ok {
				continue
			}
			named, ok := typeName.Type().(*types.Named)
			if !ok || named.Underlying() == nil {
				continue
			}
			if _, ok := named.Underlying().(*types.Struct); !ok {
				continue
			}
			if !implementsTableName(named, pkg) {
				continue
			}
			key := strings.ToLower(typeName.Name())
			if _, exists := parsed[key]; exists {
				continue
			}
			fields := []FieldInfo{}
			walkStructFromTypes(named, visited, &fields)
			parsed[key] = TableInfo{
				Name:      LowerFirst(typeName.Name()),
				TableName: typeName.Name(),
				Fields:    fields,
			}
		}
	}
	return parsed
}

func walkStructFromTypes(named *types.Named, visited map[string]bool, out *[]FieldInfo) {
	key := named.String()
	wasVisited := visited[key]
	visited[key] = true

	st, ok := named.Underlying().(*types.Struct)
	if !ok {
		return
	}

	for i := 0; i < st.NumFields(); i++ {
		field := st.Field(i)
		if !field.Exported() {
			continue
		}
		if field.Embedded() {
			if embeddedNamed, ok := deref(field.Type()).(*types.Named); ok {
				typeName := embeddedNamed.Obj().Name()
				if predefined, ok := predefinedStructs[typeName]; ok {
					*out = append(*out, predefined...)
				} else {
					walkStructFromTypes(embeddedNamed, visited, out)
				}
				continue
			}
		}
		tag := reflect.StructTag(st.Tag(i))
		gormMap := parseGormTag(tag.Get("gorm"))

		fieldType := field.Type().String()
		if mapped, ok := matchTypedAlias(fieldType); ok {
			fieldType = mapped
		}

		jsonName := tag.Get("json")
		if jsonName == "" {
			jsonName = field.Name()
		}
		*out = append(*out, FieldInfo{
			Name:      field.Name(),
			Type:      fieldType,
			JSONTag:   jsonName,
			ColumnTag: gormMap["column"],
			Comment:   gormMap["comment"],
		})
	}

	if wasVisited {
		visited[key] = false
	}
}

func implementsTableName(named *types.Named, pkg *packages.Package) bool {
	check := func(typ types.Type) bool {
		if method, _, _ := types.LookupFieldOrMethod(typ, false, pkg.Types, "TableName"); method != nil {
			if sig, ok := method.Type().(*types.Signature); ok &&
				sig.Params().Len() == 0 &&
				sig.Results().Len() == 1 &&
				types.Identical(sig.Results().At(0).Type(), types.Typ[types.String]) {
				return true
			}
		}
		return false
	}
	return check(named) || check(types.NewPointer(named))
}

func parseGormTag(tag string) map[string]string {
	result := make(map[string]string)
	parts := strings.Split(tag, ";")
	for _, part := range parts {
		if kv := strings.SplitN(part, ":", 2); len(kv) == 2 {
			result[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
		}
	}
	return result
}

func deref(typ types.Type) types.Type {
	if ptr, ok := typ.(*types.Pointer); ok {
		return ptr.Elem()
	}
	return typ
}

func matchTypedAlias(typ string) (string, bool) {
	for pattern, mapped := range typedPredefinedStructs {
		if strings.Contains(typ, pattern) {
			return mapped, true
		}
	}
	return "", false
}

func LowerFirst(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}
