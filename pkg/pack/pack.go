package pack

// Scanner describes types that support the behaviour of scanning the host for
// an installation of a given package.
type Scanner interface {
	Scan(Package) (Found, error)
}

// VersionCompareFunc is the type of a function that can compare one version
// string against another to determine if the versions match some criteria
// indicating that a version of a package was found.
type VersionCompareFunc func(string, string) Found

// Package describes a software package that may be installed on a host.
type Package struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// Found is a descriptive alias for a boolean produced by a Scanner that has
// performed a scan for a particular package.
type Found bool

const (
	WasFound Found = true
	NotFound Found = false
)

// Equals determines whether two packages describe the same thing.
func (pkg Package) Equals(other Package) bool {
	return pkg.Name == other.Name && pkg.Version == other.Version
}
