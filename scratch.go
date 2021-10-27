package scratch

import (
	"encoding/json"

	"github.com/Nv7-Github/scratch/assets"
	"github.com/Nv7-Github/scratch/blocks"
	"github.com/Nv7-Github/scratch/sprites"
	"github.com/Nv7-Github/scratch/types"
)

var Stage *sprites.Stage

// No target support
var Extensions []Extension
var Metadata types.ScratchMetadata

func init() {
	Clear()
}

func Build() types.ScratchProject {
	exts := make([]string, len(Extensions))
	for i, ext := range Extensions {
		exts[i] = extensionNames[ext]
	}
	monitors := make([]types.ScratchMonitor, len(Monitors))
	for i, m := range Monitors {
		monitors[i] = m.Build()
	}
	return types.ScratchProject{
		Targets:    []types.ScratchTarget{Stage.Build()},
		Monitors:   monitors,
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

func Clear() {
	Stage = sprites.NewStage()
	Extensions = make([]Extension, 0)
	Monitors = make([]Monitor, 0)
	Metadata = types.ScratchMetadata{
		SemVer: "3.0.0",
		VM:     "0.2.0-prerelease.20211015091140",                                                                                          // VM at time of making this
		Agent:  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.81 Safari/537.36", // My laptop
	}
	blocks.Clear()
	assets.Clear()
	types.Clear()
}
