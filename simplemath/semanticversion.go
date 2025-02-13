package simplemath

import "fmt"

type SemanticVersion struct {
	major, minor, patch int
}

func NewSemanticVersion(major, minor, patch int) SemanticVersion {
	return SemanticVersion{
		major: major,
		minor: minor,
		patch: patch,
	}
}

func (version *SemanticVersion) String() string {
	return fmt.Sprintf("%d.%d.%d", version.minor, version.minor, version.patch)
}

func (version *SemanticVersion) IncrementMajor() {
	version.major += 1
}

func (version *SemanticVersion) IncrementMinor() {
	version.minor += 1
}

func (version *SemanticVersion) IncrementPatch() {
	version.patch += 1
}
