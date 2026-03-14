# ligoip

A small terminal CLI for checking your public IP address or looking up details about any IP address. It uses `ip-api.com` and can print either a readable summary or raw JSON.

## What you get

- **Check your current public IP** with no arguments
- **Look up any IP address** from the command line
- **Readable terminal output** for quick inspection
- **JSON output** for scripts and automation
- **No API key required**

---

## Install

### Option 1: Download from Releases

Open [Releases](https://github.com/theoldzoom/goip/releases) and download the file for your system:


| I use...              | Download this file  |
| --------------------- | ------------------- |
| Linux (x86_64)        | `goip-linux-amd64`  |
| Linux (ARM64)         | `goip-linux-arm64`  |
| macOS (Intel)         | `goip-darwin-amd64` |
| macOS (Apple Silicon) | `goip-darwin-arm64` |


Then make it executable and run it:

```bash
chmod +x goip-linux-amd64
./goip-linux-amd64
```

### Option 2: Install to `~/.local/bin`

This installs the latest matching release as `~/.local/bin/goip`:

```bash
curl -fsSL https://raw.githubusercontent.com/theoldzoom/goip/master/install.sh | sh
```

If `~/.local/bin` is already in your `PATH`, you can then run:

```bash
goip
```

### Option 3: Build from source

You need [Go](https://go.dev/) installed.

```bash
git clone https://github.com/theoldzoom/goip.git
cd goip
make build
./build/goip
```

## Usage

### Show your own public IP info

```bash
goip
```

Example formatted output:

```text
IP:        8.8.8.8
Country:   United States (US)
Region:    Virginia (VA)
City:      Ashburn
ZIP:       20149
Timezone:  America/New_York
ISP:       Google LLC
Org:       Google Public DNS
AS:        AS15169 Google LLC
Coords:    39.0300, -77.5000
```

### Look up a specific IP

```bash
goip 8.8.8.8
```

### Output JSON

```bash
goip --json
goip 8.8.8.8 --json
```

Example JSON:

```json
{"status":"success","country":"United States","countryCode":"US","region":"VA","regionName":"Virginia","city":"Ashburn","zip":"20149","lat":39.03,"lon":-77.5,"timezone":"America/New_York","isp":"Google LLC","org":"Google Public DNS","as":"AS15169 Google LLC","query":"8.8.8.8","message":""}
```

### Help

```bash
goip --help
```

Available flags:

- `--json` - output JSON instead of formatted text
- `--config <path>` - load a config file from a specific location

---

## Notes

- Lookups are powered by `ip-api.com`
- The tool does not require an account or API key
- A config file path can be provided, but the CLI currently has no user-facing settings beyond loading that file

---

## Project structure

```text
.
├── cmd/              # Cobra CLI entrypoint
├── internal/
│   ├── ipinfo/       # IP lookup client and response types
│   └── output/       # Human-readable output formatting
├── makefile          # Build and release targets
├── install.sh        # Install latest release to ~/.local/bin
├── main.go
├── go.mod
└── LICENSE
```

## Tech stack

- [Go](https://go.dev/)
- [Cobra](https://github.com/spf13/cobra) - CLI framework

---

## Contributing

Contributions are welcome.

```bash
git clone https://github.com/theoldzoom/goip.git
cd goip
make build
./build/goip
```

Before opening a PR:

- Keep changes focused
- Run `go test ./...`
- Run `make build`
- Follow the existing project style


## License

MIT. See `[LICENSE](LICENSE)`.