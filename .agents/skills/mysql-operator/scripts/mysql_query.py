#!/usr/bin/env python3
"""
MySQL Query Script
Reads database configuration from server/config.yaml and executes SQL queries

This script uses pymysql (if available) or falls back to mysql command-line client.
"""

import sys
import os
import re
import subprocess
import argparse

# Try to import pymysql for direct Python connection
try:
    import pymysql
    HAS_PYMYSQL = True
except ImportError:
    HAS_PYMYSQL = False


def parse_yaml_mysql_config(config_path):
    """Simple YAML parser to extract MySQL config from config.yaml"""

    if not os.path.exists(config_path):
        return None

    config = {
        'host': 'localhost',
        'port': 3306,
        'user': 'root',
        'password': '',
        'database': '',
    }

    with open(config_path, 'r') as f:
        content = f.read()

    # Extract mysql section
    lines = content.split('\n')

    in_mysql = False
    for line in lines:
        if re.match(r'^mysql:\s*$', line):
            in_mysql = True
            continue
        if in_mysql:
            if re.match(r'^[a-z-]+:\s*$', line) and not line.startswith(' '):
                break
            match = re.match(r'\s+(\S+):\s*(.+?)\s*$', line)
            if match:
                key, value = match.groups()
                if key == 'path':
                    config['host'] = value
                elif key == 'port':
                    config['port'] = int(value.strip('"'))
                elif key == 'username':
                    config['user'] = value
                elif key == 'password':
                    config['password'] = value
                elif key == 'db-name':
                    config['database'] = value

    return config


def format_table(results):
    """Format query results as a markdown table"""
    if not results:
        print("Empty result set")
        return

    # Get column names
    if isinstance(results[0], dict):
        columns = list(results[0].keys())
        rows = results
    else:
        # Assume list of tuples
        columns = results[0]._fields if hasattr(results[0], '_fields') else []
        rows = [dict(zip(columns, row)) for row in results]

    if not columns:
        print(results)
        return

    # Calculate column widths
    col_widths = {col: len(str(col)) for col in columns}
    for row in rows:
        for col in columns:
            col_widths[col] = max(col_widths[col], len(str(row.get(col, ''))))

    # Print header
    header = "| " + " | ".join(str(col).ljust(col_widths[col]) for col in columns) + " |"
    separator = "|-" + "-|-".join("-" * col_widths[col] for col in columns) + "-|"
    print(header)
    print(separator)

    # Print rows
    for row in rows:
        row_str = "| " + " | ".join(str(row.get(col, '')).ljust(col_widths[col]) for col in columns) + " |"
        print(row_str)

    print(f"\nTotal rows: {len(rows)}")


def execute_with_pymysql(config, query):
    """Execute query using PyMySQL"""
    try:
        connection = pymysql.connect(
            host=config['host'],
            port=config['port'],
            user=config['user'],
            password=config['password'],
            database=config['database'],
            charset='utf8mb4',
            cursorclass=pymysql.cursors.DictCursor
        )

        with connection.cursor() as cursor:
            cursor.execute(query)
            result = cursor.fetchall()

        connection.close()
        return 0, result, None

    except pymysql.err.OperationalError as e:
        error_msg = f"Connection failed: {e.args[1]} (Code: {e.args[0]})"
        return 1, None, error_msg

    except pymysql.err.MySQLError as e:
        error_msg = f"Query failed: {e}"
        return 1, None, error_msg

    except Exception as e:
        error_msg = f"Error: {e}"
        return 1, None, error_msg


def execute_with_mysql_client(config, query):
    """Execute query using mysql command-line client"""

    cmd = [
        'mysql',
        f'-h{config["host"]}',
        f'-P{config["port"]}',
        f'-u{config["user"]}',
    ]

    if config.get('password'):
        cmd.append(f'-p{config["password"]}')

    cmd.append(config['database'])
    cmd.extend(['-e', query])

    result = subprocess.run(
        cmd,
        capture_output=True,
        text=True
    )

    if result.returncode == 0:
        return 0, result.stdout, None
    else:
        return result.returncode, None, result.stderr


def execute_query(config, query):
    """Execute SQL query using available method"""

    # Try PyMySQL first (if available)
    if HAS_PYMYSQL:
        returncode, result, error = execute_with_pymysql(config, query)

        if returncode == 0:
            format_table(result)
            return 0
        else:
            print(f"Error: {error}", file=sys.stderr)
            print(f"\nConnection details:", file=sys.stderr)
            print(f"  Host: {config['host']}:{config['port']}", file=sys.stderr)
            print(f"  User: {config['user']}", file=sys.stderr)
            print(f"  Database: {config['database']}", file=sys.stderr)
            return 1

    # Fall back to mysql client
    try:
        returncode, result, error = execute_with_mysql_client(config, query)

        if returncode == 0:
            print(result, end='')
            return 0
        else:
            print(f"Error: {error}", file=sys.stderr)
            return 1

    except FileNotFoundError:
        print("Error: No MySQL client or PyMySQL available.", file=sys.stderr)
        print("\nTo connect to MySQL, install one of:", file=sys.stderr)
        print("  1. Python: pip install pymysql", file=sys.stderr)
        print("  2. MySQL Client: brew install mysql-client", file=sys.stderr)
        print("\nConnection details:", file=sys.stderr)
        print(f"  Host: {config['host']}:{config['port']}", file=sys.stderr)
        print(f"  User: {config['user']}", file=sys.stderr)
        print(f"  Database: {config['database']}", file=sys.stderr)
        return 1

    except Exception as e:
        print(f"Error: {e}", file=sys.stderr)
        return 1


def main():
    parser = argparse.ArgumentParser(description='Execute MySQL queries using config from server/config.yaml')
    parser.add_argument('query', nargs='?', help='SQL query to execute')
    parser.add_argument('-f', '--file', help='Read query from file')
    parser.add_argument('-c', '--config', default='server/config.yaml',
                        help='Path to config file (default: server/config.yaml)')
    parser.add_argument('-v', '--verbose', action='store_true', help='Verbose output')
    parser.add_argument('-j', '--json', action='store_true', help='Output as JSON (requires PyMySQL)')

    args = parser.parse_args()

    # Get query
    if args.file:
        with open(args.file, 'r') as f:
            query = f.read().strip()
    elif args.query:
        query = args.query
    elif not sys.stdin.isatty():
        query = sys.stdin.read().strip()
    else:
        parser.print_help()
        print("\nExamples:", file=sys.stderr)
        print("  mysql_query.py \"SHOW TABLES\"", file=sys.stderr)
        print("  mysql_query.py \"SELECT * FROM users LIMIT 10\"", file=sys.stderr)
        sys.exit(1)

    if not query:
        print("Error: Empty query", file=sys.stderr)
        sys.exit(1)

    # Find project root
    script_dir = os.path.dirname(os.path.abspath(__file__))
    project_root = os.path.abspath(os.path.join(script_dir, '../../../..'))

    # If config path is relative, make it relative to project root
    if not os.path.isabs(args.config):
        config_path = os.path.join(project_root, args.config)
    else:
        config_path = args.config

    if args.verbose:
        print(f"Using config: {config_path}", file=sys.stderr)
        print(f"PyMySQL available: {HAS_PYMYSQL}", file=sys.stderr)

    # Load config
    mysql_config = parse_yaml_mysql_config(config_path)

    if mysql_config is None:
        print(f"Error: Cannot read config from {config_path}", file=sys.stderr)
        sys.exit(1)

    if args.verbose:
        print(f"Connecting to {mysql_config['host']}:{mysql_config['port']}", file=sys.stderr)

    # Handle JSON output (requires PyMySQL)
    if args.json:
        if not HAS_PYMYSQL:
            print("Error: JSON output requires PyMySQL", file=sys.stderr)
            sys.exit(1)

        import json
        returncode, result, error = execute_with_pymysql(mysql_config, query)

        if returncode == 0:
            print(json.dumps(result, indent=2, default=str))
            sys.exit(0)
        else:
            print(f"Error: {error}", file=sys.stderr)
            sys.exit(1)

    # Execute query
    sys.exit(execute_query(mysql_config, query))


if __name__ == '__main__':
    main()
