package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"

	"orienteering-results-scraper/http"
	"orienteering-results-scraper/sitiming"
	"orienteering-results-scraper/stats"
)

var (
	onlyCourses = flag.String("only-courses", "", "only process courses matching the supplied regular expression")
	help        = flag.Bool("help", false, "show this help message")
)

func main() {
	flag.Usage = func() {
		fmt.Println("\nUsage: orienteering-results-scraper [options] URL")
		flag.PrintDefaults()
	}
	flag.Parse()
	if *help {
		flag.Usage()
		os.Exit(0)
	}
	if len(flag.Args()) < 1 {
		_, _ = fmt.Fprintln(flag.CommandLine.Output(), "Please provide a URL on the command line")
		flag.Usage()
		os.Exit(1)
	}
	var reOnlyCourses *regexp.Regexp
	if *onlyCourses != "" {
		var err error
		reOnlyCourses, err = regexp.Compile(*onlyCourses)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "Invalid regular expression for --only-courses:", err)
			os.Exit(1)
		}
	}
	if err := run(flag.Arg(0), reOnlyCourses); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(url string, onlyCourses *regexp.Regexp) error {
	fmt.Printf("Fetching web page: %s\n", url)
	htmlContent, err := http.GetHTMLContent(url)
	if err != nil {
		return err
	}
	eventStats, err := extractStats(htmlContent, onlyCourses)
	if err != nil {
		return err
	}

	onlyCoursesDescription := ":"
	if onlyCourses != nil {
		onlyCoursesDescription = fmt.Sprintf(" matching: %s", onlyCourses.String())
	}

	fmt.Printf("Title: %s\n", eventStats.Title)
	fmt.Printf("Found %d courses%s\n", len(eventStats.Courses), onlyCoursesDescription)

	longestCourseName := 0
	for _, c := range eventStats.Courses {
		longestCourseName = max(longestCourseName, len(c.Name))
	}
	for _, c := range eventStats.Courses {
		fmt.Printf("  %s %s: %3d ;  MS %2d, MJ %2d ;  NS %2d, NJ %2d\n",
			c.Name, strings.Repeat(" ", longestCourseName-len(c.Name)), c.Members.Total(),
			c.Members.Yes.Senior, c.Members.Yes.Junior,
			c.Members.No.Senior, c.Members.No.Junior)
	}
	totalMembers := eventStats.TotalMembers()
	fmt.Printf("Total entrants: %d\n", totalMembers.Total())
	fmt.Printf("Members:     Senior %3d , Junior %3d\n", totalMembers.Yes.Senior, totalMembers.Yes.Junior)
	fmt.Printf("Non-members: Senior %3d , Junior %3d\n", totalMembers.No.Senior, totalMembers.No.Junior)
	if len(eventStats.Warnings) > 0 {
		fmt.Println()
		for _, warning := range eventStats.Warnings {
			fmt.Println(warning)
		}
	}
	fmt.Printf("\nNon-BO clubs: %s\n", strings.Join(eventStats.NonBOClubs, ", "))
	return nil
}

func extractStats(htmlContent string, onlyCourses *regexp.Regexp) (stats.EventStats, error) {
	title, coursesData, err := sitiming.ExtractCoursesData(htmlContent)
	if err != nil {
		return stats.EventStats{}, err
	}
	if onlyCourses != nil {
		coursesData = slices.DeleteFunc(coursesData, func(c stats.CourseData) bool {
			return !onlyCourses.MatchString(c.Name)
		})
	}
	return stats.ComputeEvent(title, coursesData)
}
