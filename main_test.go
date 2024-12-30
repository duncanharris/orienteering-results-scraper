package main

import (
	"os"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"orienteering-results-scraper/stats"
)

func Test_extractStats(t *testing.T) {
	type Case struct {
		htmlFile    string
		onlyCourses string
		expResults  stats.EventStats
		expTotal    stats.Members
	}
	test := func(c Case) {
		b, err := os.ReadFile(c.htmlFile)
		require.NoError(t, err)
		var reOnlyCourses *regexp.Regexp
		if c.onlyCourses != "" {
			reOnlyCourses, err = regexp.Compile(c.onlyCourses)
			require.NoError(t, err)
		}
		eventStats, err := extractStats(string(b), reOnlyCourses)
		require.NoError(t, err)
		assert.Equal(t, c.expResults, eventStats)
		assert.Equal(t, c.expTotal, eventStats.TotalMembers())
	}

	test(Case{
		htmlFile: "testdata/Stadt-Moers-SL-09-Mar-2024.html",
		expResults: stats.EventStats{
			Title: "Results - Stadt Moers SL - 09 Mar 2024",
			Courses: []stats.CourseStats{{
				Name: "Green",
				Members: stats.Members{
					Yes: stats.AgeSplit{Senior: 24, Junior: 4},
					No:  stats.AgeSplit{Senior: 2, Junior: 1},
				},
			}, {
				Name: "Light Green",
				Members: stats.Members{
					Yes: stats.AgeSplit{Senior: 10, Junior: 1},
					No:  stats.AgeSplit{Senior: 0, Junior: 12},
				},
			}, {
				Name: "Orange",
				Members: stats.Members{
					Yes: stats.AgeSplit{Senior: 0, Junior: 0},
					No:  stats.AgeSplit{Senior: 0, Junior: 13},
				},
			}, {
				Name: "Yellow",
				Members: stats.Members{
					Yes: stats.AgeSplit{Senior: 0, Junior: 1},
					No:  stats.AgeSplit{Senior: 5, Junior: 16},
				},
			}},
			NonBOClubs: []string{"BIRK", "BRID", "BRIEP", "CHS", "FALL", "HILBRE", "IND", "MO", "UPT"},
		},
		expTotal: stats.Members{
			Yes: stats.AgeSplit{Senior: 34, Junior: 6},
			No:  stats.AgeSplit{Senior: 7, Junior: 42},
		},
	})
	test(Case{
		htmlFile:    "testdata/Stadt-Moers-SL-09-Mar-2024.html",
		onlyCourses: `Green`, // includes both "Green" and "Light Green"
		expResults: stats.EventStats{
			Title: "Results - Stadt Moers SL - 09 Mar 2024",
			Courses: []stats.CourseStats{{
				Name: "Green",
				Members: stats.Members{
					Yes: stats.AgeSplit{Senior: 24, Junior: 4},
					No:  stats.AgeSplit{Senior: 2, Junior: 1},
				},
			}, {
				Name: "Light Green",
				Members: stats.Members{
					Yes: stats.AgeSplit{Senior: 10, Junior: 1},
					No:  stats.AgeSplit{Senior: 0, Junior: 12},
				},
			}},
			NonBOClubs: []string{"BIRK", "CHS", "FALL", "IND", "MO", "UPT"},
		},
		expTotal: stats.Members{
			Yes: stats.AgeSplit{Senior: 34, Junior: 5},
			No:  stats.AgeSplit{Senior: 2, Junior: 13},
		},
	})
	test(Case{
		htmlFile: "testdata/Verdin-Park-DS-27-Jun-2024.html",
		expResults: stats.EventStats{
			Title: "Results - Verdin Park DS - 27 Jun 2024",
			Courses: []stats.CourseStats{{
				Name: "Yellow",
				Members: stats.Members{
					Yes: stats.AgeSplit{Senior: 0, Junior: 0},
					No:  stats.AgeSplit{Senior: 0, Junior: 8},
				},
			}, {
				Name: "Sprint",
				Members: stats.Members{
					Yes: stats.AgeSplit{Senior: 31, Junior: 3},
					No:  stats.AgeSplit{Senior: 2, Junior: 35},
				},
			}},
			NonBOClubs: []string{"IND"},
			Warnings:   []string{`Eleanor Liney, IND on Yellow has bad age category: "": counted as junior`},
		},
		expTotal: stats.Members{
			Yes: stats.AgeSplit{Senior: 31, Junior: 3},
			No:  stats.AgeSplit{Senior: 2, Junior: 43},
		},
	})
}
