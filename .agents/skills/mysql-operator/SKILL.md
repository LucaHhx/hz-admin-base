---
name: mysql-operator
description: MySQL database operations skill for GVA projects. Use this skill whenever the user wants to explore, query, or manipulate the MySQL database defined in server/config.yaml. Trigger on phrases like "query database", "check mysql", "show tables", "explore database", "sql query", "database operations", "check data", "look at database", or any request involving MySQL/MariaDB data access in the project.
---

# MySQL Operator

A skill for connecting to and operating the MySQL database defined in `server/config.yaml`.

## Prerequisites

This skill requires one of the following:

1. **MySQL client** (recommended - faster output formatting)
   - macOS: `brew install mysql-client`
   - Ubuntu/Debian: `sudo apt install mysql-client`
   - CentOS/RHEL: `sudo yum install mysql`

2. **Python with pymysql** (alternative)
   - `pip install pymysql PyYAML`

If neither is available, the skill will report the connection details so you can connect manually.

## Database Connection

The database connection is configured in `server/config.yaml` under the `mysql` section (lines 130-143):

```yaml
mysql:
  path: 16.163.7.125        # Database host
  port: "3306"               # Database port
  username: root             # Database user
  password: ***              # Database password
  db-name: gva-base          # Database name
```

## Workflow

### Step 1: Read Configuration

Read `server/config.yaml` and extract the MySQL connection parameters.

### Step 2: Test Connection

Before performing any operation, test the database connection:

```bash
./scripts/mysql_query.py "SELECT 1"
```

**If connection fails:**
- Exit immediately with a clear error message
- Report the connection failure details (host, port, user)
- **Do not** attempt further database operations

**If connection succeeds:**
- Proceed to execute the requested operation
- Report results to the user

### Step 3: Execute Operations

Use the bundled script at `scripts/mysql_query.py` for all database operations.

## Available Operations

### Explore Database Structure

**Show all tables:**
```bash
./scripts/mysql_query.py "SHOW TABLES"
```

**Describe a table structure:**
```bash
./scripts/mysql_query.py "DESCRIBE sys_users"
# or
./scripts/mysql_query.py "SHOW COLUMNS FROM sys_users"
```

**Show table details:**
```bash
./scripts/mysql_query.py "SHOW CREATE TABLE sys_users"
```

### Query Data

**List records with limit:**
```bash
./scripts/mysql_query.py "SELECT * FROM sys_users LIMIT 10"
```

**Count rows:**
```bash
./scripts/mysql_query.py "SELECT COUNT(*) as total FROM sys_users"
```

**Query with conditions:**
```bash
./scripts/mysql_query.py "SELECT * FROM sys_users WHERE username = 'admin'"
```

**Join queries:**
```bash
./scripts/mysql_query.py "SELECT u.username, r.name FROM sys_users u JOIN sys_authority r ON u.authority_id = r.authority_id"
```

### Data Manipulation

**⚠️ Destructive operations require user confirmation:**

Before executing UPDATE, DELETE, DROP, or TRUNCATE, ask the user:
```
"You are about to: <operation>
This will affect <X> rows. Continue? (yes/no)"
```

**Update data:**
```bash
./scripts/mysql_query.py "UPDATE sys_users SET nick_name = 'Admin' WHERE username = 'admin'"
```

**Delete data:**
```bash
./scripts/mysql_query.py "DELETE FROM sys_users WHERE id = 123"
```

**Insert data:**
```bash
./scripts/mysql_query.py "INSERT INTO sys_users (username, nick_name) VALUES ('testuser', 'Test User')"
```

### Common GVA Tables

The gin-vue-admin project typically has these tables:

| Table | Description |
|-------|-------------|
| `sys_users` | User accounts |
| `sys_authorities` | Roles/Authorities |
| `sys_apis` | API endpoints |
| `sys_menus` | Menu items |
| `sys_bases` | Base parameters |
| `sys_dictionaries` | Dictionary data |
| `sys_operation_records` | Operation logs |
| `jwt_blacklists` | JWT blacklist |

## Output Format

Present results clearly to the user:

### For Query Results:
```
## Query Results

| id | username | nick_name | authority_id |
|----|----------|-----------|--------------|
| 1  | admin    | Admin     | 888          |
| 2  | test     | Test User | 999          |

**Total rows:** 2
```

### For Empty Results:
```
## Query Results

No data found. The table may be empty or the condition matched no records.
```

### For Connection Errors:
```
## Connection Error

Could not connect to MySQL server.

**Connection Details:**
- Host: 16.163.7.125:3306
- Database: gva-base
- User: root

**Possible causes:**
- Database server is down
- Network connectivity issue
- Incorrect credentials
- Firewall blocking connection
```

## Best Practices

1. **Always use LIMIT** for SELECT queries on large tables to avoid overwhelming output
2. **Use COUNT(*)** first to understand table sizes before full queries
3. **Quote strings properly** in SQL queries (use single quotes)
4. **Check table structure** with DESCRIBE before writing complex queries
5. **For destructive operations**, always confirm with the user first

## Example Usage Patterns

```
User: "Show me all users"
You: Execute: SELECT * FROM sys_users LIMIT 100

User: "How many users are there?"
You: Execute: SELECT COUNT(*) FROM sys_users

User: "What's the structure of the users table?"
You: Execute: DESCRIBE sys_users

User: "Find user with ID 5"
You: Execute: SELECT * FROM sys_users WHERE id = 5

User: "List all admin users"
You: Execute: SELECT * FROM sys_users WHERE authority_id = 888
```

## Error Handling

If you encounter any of these errors, communicate clearly to the user:

| Error | Meaning | Solution |
|-------|---------|----------|
| `Can't connect to MySQL server` | Network/connection issue | Check server status, network, firewall |
| `Access denied for user` | Wrong credentials | Verify username/password in config |
| `Unknown database` | Database doesn't exist | Check db-name in config |
| `Table doesn't exist` | Table not found | Use SHOW TABLES to list available tables |
| `Column not found` | Wrong column name | Use DESCRIBE to check table structure |
