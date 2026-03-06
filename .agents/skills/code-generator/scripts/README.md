# Code Generator Scripts

## Important Note

The main code generator script is located at **`server/cmd/codegen/main.go`**.

## Usage

```bash
cd server/cmd/codegen
go run . <table_name> [package] [description]
```

## Why server/generate.go?

1. **Direct Service Access**: Uses `AutoCodeTemplate.Create()` directly without HTTP API
2. **No Authentication Needed**: Bypasses login requirement
3. **Reliable**: Works regardless of server running status
4. **Proper Initialization**: Uses `core.InitServer()` for correct DB setup

## Examples

```bash
# Generate agent table
go run generate.go b_agent business "代理商管理"

# Generate user table
go run generate.go sys_user system "用户管理"

# Generate with default description
go run generate.go orders example
```

## Alternative (Requires Auth)

The Python script (`generate_code.py`) can be used but requires:
1. Backend server running
2. Valid authentication token
3. `requests` library installed

**Recommendation**: Use `go run generate.go` for simplicity.
