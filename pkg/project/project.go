package project

var (
	bundleVersion = "0.0.1"
	description   = "The loadtest-operator does something."
	gitSHA        = "n/a"
	name          = "loadtest-operator"
	source        = "https://github.com/giantswarm/loadtest-operator"
	version       = "n/a"
)

func BundleVersion() string {
	return bundleVersion
}

func Description() string {
	return description
}

func GitSHA() string {
	return gitSHA
}

func Name() string {
	return name
}

func Source() string {
	return source
}

func Version() string {
	return version
}
