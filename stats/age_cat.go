package stats

import (
	"fmt"
	"regexp"
	"strconv"
)

var (
	reAgeCat = regexp.MustCompile(`^([MmWwBbGg])(\d{1,2})$`)
)

func isJuniorAgeCat(ageCat string) (bool, error) {
	parts := reAgeCat.FindStringSubmatch(ageCat)
	if parts == nil {
		// count invalid age categories as seniors
		return false, fmt.Errorf("bad age category: %q", ageCat)
	}
	age, err := strconv.Atoi(parts[2])
	if err != nil {
		panic(err) // valid number in regex
	}
	switch parts[1] {
	case "B", "b", "G", "g":
		return true, nil
	case "M", "m", "W", "w":
		return age < 21, nil
	default:
		panic("unreachable")
	}
}
