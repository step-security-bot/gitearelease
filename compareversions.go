package gitearelease

import (
	"fmt"
	"strconv"
	"strings"
)

// CompareVersions compares two version strings and returns 1 if v1 is greater than v2, -1 if v1 is less than v2, and 0 if v1 is equal to v2
func CompareVersions(v1, v2 string) int {
	// Remove the version prefixs from the version strings
	v1 = TrimVersionPrefix(v1)
	v2 = TrimVersionPrefix(v2)

	// Split the version strings into individual version numbers
	v1Numbers := strings.Split(v1, ".")
	v2Numbers := strings.Split(v2, ".")

	// Convert the version numbers to integers and compare them
	for i := 0; i < len(v1Numbers) && i < len(v2Numbers); i++ {
		v1Num, err := strconv.Atoi(v1Numbers[i])
		if err != nil {
			fmt.Println("Invalid version number:", v1Numbers[i])
			return 0
		}
		v2Num, err := strconv.Atoi(v2Numbers[i])
		if err != nil {
			fmt.Println("Invalid version number:", v2Numbers[i])
			return 0
		}

		if v1Num > v2Num {
			return 1
		} else if v1Num < v2Num {
			return -1
		}
	}

	// If we got here, the version numbers are the same up to this point
	// Check if one of the versions has more version numbers than the other
	if len(v1Numbers) > len(v2Numbers) {
		return 1
	} else if len(v1Numbers) < len(v2Numbers) {
		return -1
	}

	// The version numbers are the same
	return 0
}
