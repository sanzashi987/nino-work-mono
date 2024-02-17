package utils

import (
	"fmt"
	"strconv"
	"strings"
)

const DefaultVersion = "0.1.0"

// Version represents a version string with major, minor, and patch components.
type Version struct {
	Major int
	Minor int
	Patch int
}

// NewVersion creates a new Version instance from a version string.
func VersionFromString(versionStr string) (*Version, error) {
	parts := strings.Split(versionStr, ".")

	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid version format: %s", versionStr)
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, fmt.Errorf("invalid major version: %s", parts[0])
	}

	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("invalid minor version: %s", parts[1])
	}

	patch, err := strconv.Atoi(parts[2])
	if err != nil {
		return nil, fmt.Errorf("invalid patch version: %s", parts[2])
	}

	return &Version{Major: major, Minor: minor, Patch: patch}, nil
}

// String returns the string representation of the Version.
func (v *Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

// Compare compares two versions and returns an integer indicating their relationship.
// -1 if v is less than other, 0 if they are equal, and 1 if v is greater than other.
func (v *Version) Compare(other *Version) int {
	switch {
	case v.Major < other.Major:
		return -1
	case v.Major > other.Major:
		return 1
	case v.Minor < other.Minor:
		return -1
	case v.Minor > other.Minor:
		return 1
	case v.Patch < other.Patch:
		return -1
	case v.Patch > other.Patch:
		return 1
	default:
		return 0
	}
}

// BumpMajor increments the Major version number and resets Minor and Patch to zero.
func (v *Version) BumpMajor() {
	v.Major++
	v.Minor = 0
	v.Patch = 0
}

// BumpMinor increments the Minor version number and resets Patch to zero.
func (v *Version) BumpMinor() {
	v.Minor++
	v.Patch = 0
}

// BumpPatch increments the Patch version number.
func (v *Version) BumpPatch() {
	v.Patch++
}
