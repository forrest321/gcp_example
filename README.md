# RSS Reader CLI Tool

## Overview
The RSS Reader is a command-line application designed to fetch and process RSS feeds from various sources. It's built in Go (Golang) and can be configured to read multiple RSS feeds, either through command-line arguments or a configuration file.

## Key Features
- Read RSS feeds from specified URLs.
- Optionally use a YAML configuration file to specify multiple RSS feed URLs.
- Process and display RSS feed data in a structured format.

## Technology Stack
- **Language**: Go (Golang)
- **RSS Parsing**: `github.com/mmcdole/gofeed` for parsing RSS and Atom feeds.
- **CLI Framework**: `github.com/spf13/cobra` for building the command-line interface.
- **Configuration Management**: YAML format for feed configuration, processed using Go's standard libraries.

## Setup and Installation
1. Clone the repository: `git clone [repository-url]`
2. Navigate to the project directory: `cd [project-directory]`
3. Build the CLI tool: `go build -o rssreader ./cmd`
4. Run the tool: `./rssreader --config path/to/config.yaml` or `./rssreader --feed http://example.com/rss`

## Usage
- **Command-Line Arguments**:
    - `--feed`: Specify individual RSS feed URLs directly.
    - `--config`: Path to a YAML configuration file containing a list of RSS feed URLs.
- Example: `./rssreader --feed https://example.com/rssfeed`

## Configuration File Format
The configuration file is in YAML format. Example structure:

```yaml
feeds:
  - url: "https://example.com/rssfeed1"
  - url: "https://example2.com/rssfeed2"
