# Configuration System

This project uses [UP (Unified Properties)](https://github.com/uplang/spec) for configuration, with support for templating and composition.

## Files

- **`template.up`** - Base configuration template with variables
- **`deployment.up`** - Development/deployment environment overlay
- **`production.up`** - Production environment overlay
- **`config.up`** - Generated development configuration (not committed, in project root)

## Usage

### Development

Generate development configuration (happens automatically on `make up`):

```bash
make config.up
# or just
make up
```

This processes `deployment.up` which:
1. Loads `template.up` as the base
2. Overlays development-specific variable values
3. Generates the final `config.up`

### Production

Generate production configuration:

```bash
make config-production
```

This creates `config.production.out.up` with production settings.

## How It Works

UP templating uses declarative composition following the [UP Templating spec](https://github.com/uplang/spec/blob/main/TEMPLATING.md):

### Base Template (`template.up`)

Defines variables and uses them throughout:

```up
vars {
  environment production
  debug!bool false
  http_addr :8000
}

environment $vars.environment
debug!bool $vars.debug

app {
  httpAddr $vars.http_addr
}
```

### Environment Overlays

Override variables for specific environments:

```up
# deployment.up
config!base template.up

vars!overlay {
  environment development
  debug!bool true
  http_addr 127.0.0.1:8000
}
```

### Processing

The `go tool up template process` command:
1. Loads the base configuration
2. Applies overlays (merging blocks)
3. Resolves all variable references iteratively
4. Outputs the final configuration

## Benefits

- **Type-safe** - Variables have explicit types (`!bool`, `!int`, etc.)
- **No string substitution** - Pure data composition
- **Declarative** - Describe what you want, not how to get it
- **Composable** - Layer configurations easily
- **Clean syntax** - Same UP syntax everywhere, no templating language

## Adding New Environments

Create a new overlay file:

```up
# config/staging.up
config!base template.up

vars!overlay {
  environment staging
  http_addr staging.example.com:8000
}
```

Add a Makefile target:

```makefile
.PHONY: config-staging
config-staging: config/template.up config/staging.up
	go tool up template process -i config/staging.up -o config.staging.out.up
```

## Documentation

- [UP Specification](https://github.com/uplang/spec)
- [UP Templating Guide](https://github.com/uplang/spec/blob/main/TEMPLATING.md)
- [UP Syntax Reference](https://github.com/uplang/spec/blob/main/SYNTAX-REFERENCE.md)

