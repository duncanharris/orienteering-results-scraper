package stats

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

type (
	EventStats struct {
		Title      string
		Courses    []CourseStats
		NonBOClubs []string
		Warnings   []string
	}

	CourseStats struct {
		Name string
		Members
	}

	Members struct {
		Yes, No AgeSplit
	}

	AgeSplit struct {
		Senior, Junior int
	}
)

func (r *EventStats) TotalMembers() Members {
	var total Members
	for _, c := range r.Courses {
		total.Add(c.Members)
	}
	return total
}

func (m *Members) Total() int { return m.Yes.Total() + m.No.Total() }
func (m *Members) Add(o Members) {
	m.Yes.Add(o.Yes)
	m.No.Add(o.No)
}
func (a *AgeSplit) Total() int { return a.Senior + a.Junior }
func (a *AgeSplit) Add(o AgeSplit) {
	a.Senior += o.Senior
	a.Junior += o.Junior
}
func (a *AgeSplit) AddIsJunior(junior bool) {
	if junior {
		a.Junior++
	} else {
		a.Senior++
	}
}

func ComputeEvent(title string, coursesData []CourseData) (EventStats, error) {
	res := EventStats{
		Title:   title,
		Courses: make([]CourseStats, len(coursesData)),
	}

	nonBOClubs := map[string]struct{}{}
	entrantCourses := map[string][]string{}

	for iCourse, courseData := range coursesData {
		res.Courses[iCourse].Name = courseData.Name
		for _, entrant := range courseData.Entrants {
			junior, ageErr := isJuniorAgeCat(entrant.AgeCat)
			if ageErr != nil {
				courseName := res.Courses[iCourse].Name
				junior = strings.EqualFold(courseName, "White") ||
					strings.EqualFold(courseName, "Yellow")
				res.Warnings = append(res.Warnings,
					fmt.Sprintf("%s, %s on %s has %v: counted as %s",
						entrant.Name, entrant.Club, res.Courses[iCourse].Name, ageErr,
						ifElse(junior, "junior", "senior")))
			}

			entrantCourses[entrant.Name+", "+entrant.Club] = append(entrantCourses[entrant.Name], courseData.Name)

			if members := &res.Courses[iCourse].Members; isBOClub(entrant.Club) {
				members.Yes.AddIsJunior(junior)
			} else {
				members.No.AddIsJunior(junior)
				if entrant.Club != "" {
					nonBOClubs[entrant.Club] = struct{}{}
				}
			}
		}
	}

	// check for entrants who ran multiple courses - we shouldn't pay levy for them
	for nameAndClub, courses := range entrantCourses {
		if len(courses) > 1 {
			res.Warnings = append(res.Warnings,
				fmt.Sprintf("%s ran multiple courses: %s", nameAndClub, strings.Join(courses, ", ")))
		}
	}

	res.NonBOClubs = slices.Sorted(maps.Keys(nonBOClubs))

	return res, nil
}

func ifElse[T any](cond bool, t, f T) T {
	if cond {
		return t
	}
	return f
}
