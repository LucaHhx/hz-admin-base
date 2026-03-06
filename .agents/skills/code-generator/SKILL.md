---
name: code-generator
description: GVA code generator automation. Use this skill whenever the user wants to generate CRUD code from database tables, mentions "code generator", "auto code", "generate code from table", "scaffold from database", or "create CRUD from table". Automatically generates front-end (Vue3) and back-end (Go) code from MySQL tables by directly calling the GVA service layer (no HTTP API needed).
---

# GVA Code Generator

Automates the gin-vue-admin code generation process. Generates complete CRUD functionality from database tables without manual intervention.

## Quick Start

```bash
# From project root
cd server/cmd/codegen
go run . <table_name> [package] [description]

# Example: Generate agent table (default package: business)
go run . b_agent business "代理商管理"

# Or with default package
go run . b_agent
```

## Workflow

### Step 1: Understand Requirements

Ask the user for:
- **Table name**: Which database table to generate code from
- **Package**: Target package name (default: `business`)
- **Description**: Chinese description of the entity (e.g., "用户管理")
- **Options**: Web/Server generation (default: both)

### Step 2: Read Table Structure

The generator automatically fetches table structure from MySQL using the configured database connection.

### Step 3: Auto-Generate Code

The script:
1. Reads table columns from database
2. Maps database types to Go/field types
3. Generates naming conventions (StructName, Abbreviation, etc.)
4. Creates all necessary files (model, service, API, router, frontend)
5. Injects code into enter.go files

### Step 4: Restart Backend

**IMPORTANT**: The backend service MUST be restarted for new code to take effect.

## Type Mapping

| Database Type | Go Type | Field Type | Search Type |
|---------------|---------|------------|-------------|
| varchar(n) | string | string | like |
| text/longtext | string | richtext | like |
| int | int64 | int | = |
| tinyint | int64 | bool | = |
| bigint | int64 | int64 | = |
| decimal/double | float64 | float64 | = |
| datetime | time.Time | datetime | between |
| date | time.Time | date | between |
| json | string | json | like |

## Auto-Generated Naming

| Input (table) | StructName | Abbreviation | PackageName | HumpPackageName |
|---------------|------------|--------------|-------------|-----------------|
| sys_user | User | user | user | user |
| b_agent | Agent | agent | agent | agent |
| tb_order | Order | order | order | order |
| t_product | Product | product | product | product |

## Generated Files

### Backend Files
```
server/model/{package}/{struct}.go           # Data model
server/model/{package}/request/{struct}.go    # Request DTOs
server/service/{package}/{struct}_service.go  # Business logic
server/api/v1/{package}/{struct}.go           # API handlers
server/router/{package}/{struct}_router.go    # Routes
server/translation/zh-CN/business/{pkg}.json  # i18n Chinese
server/translation/en-US/business/{pkg}.json  # i18n English
```

### Frontend Files
```
web/src/api/{package}/{struct}.js            # API calls
web/src/view/{package}/{struct}/{struct}.vue    # List page
web/src/view/{package}/{struct}/{struct}Form.vue # Form page
```

## Field Visibility Rules

### Form Visibility (form field)
- **Hidden**: `id`, `*_at`, `*_time`, `created_by`, `updated_by`, `deleted_by`
- **Shown**: All other fields

### Table Visibility (table column)
- **Shown**: `id`, `created_at`, `updated_at`, and all non-system fields
- **Hidden**: `deleted_at`, `deleted_by`, `*_by` fields

### Required Fields
Fields matching these patterns are auto-marked as required:
- `*_id`, `name`, `code`, `uuid`

### Special Field Handling

| Pattern | Treatment |
|---------|-----------|
| `*_id` | Foreign key, shown in form/table |
| `*_at`, `*_time` | DateTime, form hidden, table shown |
| `*_pic`, `*_img` | Picture type |
| `*_file` | File upload |
| `status`, `state` | Bool/Enum with dict support |
| `note`, `remark`, `content` | Rich text type |

## Error Handling

| Error | Solution |
|-------|----------|
| Table not found | Check database connection and table name in `config.yaml` |
| Package not exists | Use existing package (`system`, `example`) or create one first |
| Duplicate struct | Use different name or delete existing from `sys_auto_code_histories` |
| Template missing | Check `server/resource/` directory exists |

## Example Session

```
User: "Generate CRUD for b_agent table"
→ go run . b_agent business "代理商管理"

User: "Generate code from orders table"
→ go run . orders example "订单管理"

User: "Create scaffold for customer table"
→ go run . customer system "客户管理"
```

## Implementation Notes

### Why Direct Go Script Instead of HTTP API?

1. **No Authentication Needed**: HTTP API requires login token
2. **Direct Service Access**: Uses `AutoCodeTemplate.Create()` directly
3. **Faster**: No network overhead
4. **Reliable**: Works regardless of server status

### Key Implementation Details

1. **Initialization**: Uses `core.InitServer()` to properly initialize database connection
2. **SQL Query**: Uses table aliases to avoid column ambiguity
3. **PrimaryField**: Explicitly set for GvaModel tables to prevent template errors
4. **Field Inference**: Smart mapping from database types to UI field types

## Package Options

| Package | Description | Files Generated In |
|---------|-------------|-------------------|
| `business` | Business module (default) | `server/service/business/`, `server/api/v1/business/`, etc. |
| `system` | System module | `server/service/system/`, `server/api/v1/system/`, etc. |
| `example` | Example module | `server/service/example/`, etc. |

**Note**: Package must already exist in the project structure.
