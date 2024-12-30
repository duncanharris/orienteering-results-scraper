package http

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
)

func GetHTMLContent(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("error creating HTTP request: %w", err)
	}
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml,text/plain")
	req.Header.Add("User-Agent", `Chrome`)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error performing HTTP request: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error: status code %d", resp.StatusCode)
	}
	const maxInspectSize = 1024
	inspectPrefix := bytes.NewBuffer(make([]byte, 0, maxInspectSize))
	_, errCopy := io.CopyN(inspectPrefix, resp.Body, maxInspectSize)
	if errCopy != nil && errCopy != io.EOF {
		return "", errCopy
	}

	contentType := resp.Header.Get("Content-Type")
	encoding, _, _ := charset.DetermineEncoding(inspectPrefix.Bytes(), contentType)
	var body io.Reader = inspectPrefix
	if errCopy == nil { // not EOF
		body = io.MultiReader(inspectPrefix, resp.Body)
	}

	utf8body, err := io.ReadAll(transform.NewReader(body, encoding.NewDecoder()))
	if err != nil {
		return "", err
	}

	return string(utf8body), nil
}
