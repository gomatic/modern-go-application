# Example Outputs

This document shows example outputs from the noop application.

## Command Structure

```
modern-go-application
├── resource (parent)
│   ├── create (demonstrates: strings, booleans, string slices)
│   └── list (demonstrates: filtering, pagination, string slices)
└── service (parent)
    ├── start (demonstrates: nested configs, database, server)
    └── stop (demonstrates: integer slices, signals)
```

## Example 1: Resource Create with Tags

Command:
```bash
./modern-go-application resource create \
  --name my-test-resource \
  --description "A test resource" \
  --tags prod,critical,app \
  --enabled
```

Output:
```json
{
  "success": true,
  "resource_id": "res-my-test-resource",
  "name": "my-test-resource",
  "description": "A test resource",
  "tags": [
    "prod",
    "critical",
    "app"
  ],
  "enabled": true,
  "dry_run": false,
  "force": false,
  "message": "Resource created successfully (noop)"
}
```

## Example 2: Resource List with Status Filter

Command:
```bash
./modern-go-application resource list \
  --status active \
  --limit 5
```

Output:
```json
{
  "resources": [
    {
      "id": "res-001",
      "name": "example-resource-1",
      "status": "active",
      "tags": [
        "prod",
        "critical"
      ],
      "created_at": "2025-01-01T00:00:00Z"
    },
    {
      "id": "res-003",
      "name": "example-resource-3",
      "status": "active",
      "tags": [
        "staging"
      ],
      "created_at": "2025-01-03T00:00:00Z"
    }
  ],
  "total": 2,
  "limit": 5,
  "offset": 0,
  "filter_statuses": [
    "active"
  ],
  "sort_by": "name",
  "ascending": true
}
```

## Example 3: Service Start with Nested Config

Command:
```bash
./modern-go-application service start \
  --service-name web-api \
  --environment prod \
  --db-host db.example.com \
  --db-port 5432 \
  --server-port 8080 \
  --workers 4 \
  --enable-cache \
  --debug
```

Output:
```json
{
  "success": true,
  "service_name": "web-api",
  "environment": "prod",
  "pid": 12345,
  "database": {
    "Host": "db.example.com",
    "Port": 5432,
    "Name": "appdb",
    "User": "app",
    "Password": ""
  },
  "server": {
    "Host": "localhost",
    "Port": 8080,
    "ReadTimeout": 30,
    "WriteTimeout": 30
  },
  "workers": 4,
  "enable_cache": true,
  "debug": true,
  "message": "Service started successfully (noop)"
}
```

## Example 4: Service Stop with Multiple PIDs

Command:
```bash
./modern-go-application service stop \
  --service-name web-api \
  --pid 1234 \
  --pid 5678 \
  --pid 9012 \
  --timeout 10 \
  --force
```

Output:
```json
{
  "success": true,
  "service_name": "web-api",
  "force": true,
  "timeout": 10,
  "pids": [
    1234,
    5678,
    9012
  ],
  "signal": "SIGTERM",
  "message": "Service force stopped successfully (noop)"
}
```

## Example 5: Using Environment Variables

Command:
```bash
export MODERN_GO_APP_RESOURCE_CREATE_NAME="env-resource"
export MODERN_GO_APP_RESOURCE_CREATE_ENABLED=true
export MODERN_GO_APP_RESOURCE_CREATE_TAGS="env-tag1,env-tag2"

./modern-go-application resource create --output /tmp/result.json
```

Output (written to `/tmp/result.json`):
```json
{
  "success": true,
  "resource_id": "res-env-resource",
  "name": "env-resource",
  "description": "",
  "tags": [
    "env-tag1",
    "env-tag2"
  ],
  "enabled": true,
  "dry_run": false,
  "force": false,
  "message": "Resource created successfully (noop)"
}
```

## Configuration Types Demonstrated

| Type | Example Flag | Environment Variable | Location |
|------|-------------|---------------------|----------|
| String | `--name my-resource` | `MODERN_GO_APP_RESOURCE_CREATE_NAME` | resource create |
| Boolean | `--enabled` | `MODERN_GO_APP_RESOURCE_CREATE_ENABLED` | resource create |
| String Slice | `--tags prod,critical` | `MODERN_GO_APP_RESOURCE_CREATE_TAGS` | resource create |
| Integer | `--limit 10` | `MODERN_GO_APP_RESOURCE_LIST_LIMIT` | resource list |
| Integer Slice | `--pid 1234 --pid 5678` | `MODERN_GO_APP_SERVICE_STOP_PIDS` | service stop |
| Nested Config | `--db-host localhost` | `MODERN_GO_APP_SERVICE_START_DB_HOST` | service start |

## Help Output

### Global Help
```bash
./modern-go-application --help
```

Shows:
- Global flags (log-level, log-format)
- Top-level commands (resource, service)
- Environment variable names for each flag

### Command Help
```bash
./modern-go-application resource create --help
```

Shows:
- Command description with examples
- All available flags with defaults
- Environment variable names
- Aliases (short flags)
- Required vs optional flags

