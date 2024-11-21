package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type HTTPRequest struct {
	Method string
	Path string
	Version string
	Headers map[string]string
	Body string
}

func parseHTTPRequest(reader *bufio.Reader) (*HTTPRequest, error) {
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	parts := strings.Fields(strings.TrimSpace(requestLine))
	if len(parts) != 3 {
		return nil, fmt.Errorf("malformed request line: %q", requestLine)
	}

	method, path, version := parts[0], parts[1], parts[2]

	headers := make(map[string]string)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, fmt.Errorf("error reading header: %v", err)
		}
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("malformed header: %q", line)
			}
		headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
	}

	var body string
	if contentLength, ok := headers["Content-Length"]; ok {
		length, err := strconv.Atoi(contentLength)
		if err != nil {
			return nil, fmt.Errorf("invalid Content-Length: %v", err)
		}

		bodyBytes := make([]byte, length)
		if _, err = io.ReadFull(reader, bodyBytes); err != nil {
			return nil, fmt.Errorf("error reading body: %v", err)
		}
		body = string(bodyBytes)
	}

	return &HTTPRequest {
		Method: method,
		Path: path,
		Version: version,
		Headers: headers,
		Body: body,
	}, nil
}