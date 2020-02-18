package types

type (

	// Config is the preset variables for your extractor
	Config struct {
		BaseURI     string
		WetPaths    string
		DataFolder  string
		MatchFolder string
		Start       int
		Stop        int
	}
)
