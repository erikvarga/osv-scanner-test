package lockfile

// KnownEcosystems returns a list of ecosystems that `lockfile` supports
// automatically inferring an extractor for based on a file path.
func KnownEcosystems() []Ecosystem {
	return []Ecosystem{
		NpmEcosystem,
		NuGetEcosystem,
		BundlerEcosystem,
		ComposerEcosystem,
		GoEcosystem,
		MixEcosystem,
		MavenEcosystem,
		PipEcosystem,
		PubEcosystem,
		ConanEcosystem,
		CRANEcosystem,
		// Disabled temporarily,
		// see https://github.com/google/osv-scanner-test/pull/128 discussion for additional context
		// AlpineEcosystem,
	}
}
