# Website Liveliness Checker

A concurrent Go program that checks the liveliness of web servers listed in a file and reports how many are up or down.

## Features

- Checks multiple websites concurrently for efficiency.
- Robust error handling with detailed logging.
- Simple command-line interface.
- Configurable HTTP timeout (default: 10 seconds).

## Requirements

- Go 1.16 or higher.

## Installation

Clone the repository and build the program using Go:

```bash
git clone https://github.com/yourusername/website-liveliness-checker.git
cd website-liveliness-checker
go build -o website-liveliness-checker
```

## Usage
Create a file named websites.txt with one URL per line. For example:
```
http://google.com
http://example.com
http://nonexistent-site.xyz
```
## Run the program:
`./website-liveliness-checker`

The program will check each website and print a summary:
```
2025/03/13 16:51:12 Error checking http://nonexistent-site.xyz: Get "http://nonexistent-site.xyz": dial tcp: lookup nonexistent-site.xyz on 127.0.0.53:53: no such host
Total websites checked: 3
Websites Up: 2
Websites Down: 1

```

Errors are logged to the console with timestamps, for example:
`2025/03/13 16:45:17 Error checking http://nonexistent-site.xyz: Get "http://nonexistent-site.xyz": dial tcp: lookup nonexistent-site.xyz on 127.0.0.53:53: no such host`


### Error Handling
The program logs errors for each website check, including:
Network errors (e.g., timeouts, DNS lookup failures).
Non-success HTTP status codes (e.g., 404, 500).


## Contributing
Contributions are welcome! Please submit issues or pull requests on GitHub.

License   
`Chill Bro`