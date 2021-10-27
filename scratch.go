package scratch

import (
	"encoding/json"

	"github.com/Nv7-Github/scratch/assets"
	"github.com/Nv7-Github/scratch/sprites"
	"github.com/Nv7-Github/scratch/types"
)

var Stage = sprites.NewStage()

// No monitors support
// No target support
var Extensions = make([]Extension, 0)
var Metadata = types.ScratchMetadata{
	SemVer: "3.0.0",
	VM:     "0.2.0-prerelease.20211015091140",                                                                                          // VM at time of making this
	Agent:  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.81 Safari/537.36", // My laptop
}

func Build() types.ScratchProject {
	exts := make([]string, len(Extensions))
	for i, ext := range Extensions {
		exts[i] = extensionNames[ext]
	}
	return types.ScratchProject{
		Targets:    []types.ScratchTarget{Stage.Build()},
		Monitors:   []types.ScratchMonitor{},
		Metadata:   Metadata,
		Extensions: exts,
	}
}

func Save(fs types.FS) error {
	err := assets.Save(fs)
	if err != nil {
		return err
	}

	proj, err := fs.Create("project.json")
	if err != nil {
		return err
	}

	enc := json.NewEncoder(proj)
	data := Build()
	return enc.Encode(data)
}
