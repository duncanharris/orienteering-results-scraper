package sitiming

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"regexp"
	"strings"

	"orienteering-results-scraper/stats"
)

func ExtractCoursesData(htmlContent string) (string, []stats.CourseData, error) {
	title := "<none>"
	if m := reTitle.FindStringSubmatch(htmlContent); m != nil {
		title = m[1]
	}

	coursesJSON := reCourse.FindAllStringSubmatch(htmlContent, -1)
	courseNamesData := reCourseNames.FindAllStringSubmatch(htmlContent, -1)
	if len(courseNamesData) != len(coursesJSON) {
		return "", nil, fmt.Errorf("found %d course names but %d courses", len(courseNamesData), len(coursesJSON))
	}
	coursesData := make([]stats.CourseData, 0, len(coursesJSON))
	for i, courseJSON := range coursesJSON {
		courseName := courseNamesData[i][1]
		var data [][]string
		if jsonErr := json.NewDecoder(bytes.NewBufferString(courseJSON[1])).Decode(&data); jsonErr != nil {
			return "", nil, fmt.Errorf("decoding JSON for course %d: %w", i+1, jsonErr)
		}
		courseData := stats.CourseData{
			Name:     courseName,
			Entrants: make([]stats.Entrant, len(data)),
		}
		for j, entrantData := range data {
			if len(entrantData) < 5 {
				return "", nil, fmt.Errorf("course %d, entrant %d: expected at least 5 fields, got %d",
					i+1, j+1, len(entrantData))
			}
			courseData.Entrants[j] = stats.Entrant{
				Name:   strings.TrimSpace(html.UnescapeString(entrantData[2])),
				Club:   strings.TrimSpace(html.UnescapeString(entrantData[3])),
				AgeCat: strings.TrimSpace(html.UnescapeString(entrantData[4])),
			}
		}
		coursesData = append(coursesData, courseData)
	}
	return title, coursesData, nil
}

const (
	reJSONString      = `"(?:[^"\\]|\\[\\"/bfnrt]|\\u[0-9a-fA-F]{4})*"`
	reJSONStringArray = `\[\s*` + reJSONString + `(?:\s*,\s*` + reJSONString + `)*\s*\]`
)

var (
	reTitle       = regexp.MustCompile(`(?i)<title>(.*?)</title>`)
	reCourse      = regexp.MustCompile(`if\s*\(\s*tableNumber\s*===?\s*\d+\)\s*return\s*(\[\s*(?:` + reJSONStringArray + `(?:\s*,\s*` + reJSONStringArray + `)*\s*)?\])`)
	reCourseNames = regexp.MustCompile(`<div\s+class="results-block-title"\s*>\s*<h3>\s*([^<>]*?)\s*</h3>\s*</div>`)
)
