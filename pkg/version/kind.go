package version

import "regexp"

// Kind indicates the kind of a version string (development, RC, or release).
// Note that this is different from the `build flavor` defined in `pkg/buildinfo`, which
// describes the set of source files/build options used for a build -- both RC and release
// builds use the same options, hence have the same `release` build flavor.
type Kind int

const (
	// InvalidKind is the version kind for unrecognized version strings.
	InvalidKind Kind = iota
	// DevelopmentKind is the version kind for development version strings.
	DevelopmentKind
	// RCKind is the version kind for RC version strings.
	RCKind
	// ReleaseKind is the version kind for release version strings.
	ReleaseKind
	// NightlyKind is the version kind for nightly release version strings.
	NightlyKind
)

//go:generate stringer -type=Kind

const (
	releaseRegexStr  = `\d+(?:\.\d+)*`
	rcSuffixRegexStr = `-rc\.\d+`
	hashTagRegexStr  = `-g[0-9a-f]{10}`
)

var (
	releaseRegex = regexp.MustCompile(`^` + releaseRegexStr + `$`)
	rcRegex      = regexp.MustCompile(`^` + releaseRegexStr + rcSuffixRegexStr + `$`)
	nightlyRegex = regexp.MustCompile(`^.*-nightly-.*$`)
	devRegex     = regexp.MustCompile(`^` + releaseRegexStr + `(?:` + rcSuffixRegexStr + `|\.x)?-\d+` + hashTagRegexStr + `(?:-dirty)?$`)
)

// GetVersionKind returns the version kind (release, RC, development) of the given version string.
func GetVersionKind(versionStr string) Kind {
	switch {
	case releaseRegex.MatchString(versionStr):
		return ReleaseKind
	case rcRegex.MatchString(versionStr):
		return RCKind
	case nightlyRegex.MatchString(versionStr):
		return NightlyKind
	case devRegex.MatchString(versionStr):
		return DevelopmentKind
	default:
		return InvalidKind
	}
}
