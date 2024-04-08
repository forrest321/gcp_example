# Goread RSS Reader

## Overview
Goread is a command-line RSS Reader application that fetches and processes RSS feeds. It is part of a larger project structured to run on Google Cloud Platform, with the backend and frontend managed separately.

## Project Structure
- `backend/rssreader`: Contains the logic for the RSS Reader.
- `cmd/main.go`: Entry point for the command-line application.
- `Makefile`: Simplifies the build process for the project.

## Getting Started

### Prerequisites
- Ensure you have Go installed on your system.
- Optional: Docker, if you plan to containerize the application.

### Installation
1. Clone the repository:
```shell
git clone https://github.com/forrest321/gcp_example
```
2. Navigate to the project directory:
```shell
cd gcp_example
```

### Building the Application
- Run the following command to build the backend (RSS Reader):
```shell
make backend
```
This will compile the backend code and create a binary named `goread`.

### Running the Application
- The compiled `goread` binary can be run with various flags to read RSS feeds:
- To specify individual feed URLs:
  ```
  ./goread --feed "http://example.com/rss"
  ```
- To use a configuration file:
  ```
  ./goread --config "path/to/config.yaml"
  ```
The configuration file should follow the format defined in the `backend/rssreader` package.

## Contributing
Contributions to the project are welcome. Please ensure to follow the guidelines and code of conduct.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
