package patterns

import (
	"../sources"
)

type SourceFactory struct{}

func (factory *SourceFactory) CreateSource(args ...string) sources.Source {
    if len(args) == 0 {
        return nil
    }
    
	if args[0] == "Guardian" && len(args) == 2 {
		return &sources.Guardian{ApiKey: args[1]}
	}

	if args[0] == "NYTimes" && len(args) == 2 {
		return &sources.NYTimes{CsvInput: args[1]}
	}

	return nil
}
