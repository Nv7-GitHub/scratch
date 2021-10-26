package scratch

import (
	"github.com/Nv7-Github/scratch/sprites"
	"github.com/Nv7-Github/scratch/types"
)

type Project struct {
	Stage *sprites.Stage
	// No monitor support yet
	// No target support yet
	Extensions []Extension
	Metadata   types.ScratchMetadata
}

func NewProject() *Project {
	return &Project{
		Stage:      sprites.NewStage(),
		Extensions: make([]Extension, 0),
		Metadata: types.ScratchMetadata{
			SemVer: "3.0.0",
			VM:     "0.2.0-prerelease.20211015091140",                                                                                          // VM at time of making this
			Agent:  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.81 Safari/537.36", // My laptop
		},
	}
}

func (p *Project) Build() types.ScratchProject {
	exts := make([]string, len(p.Extensions))
	for i, ext := range p.Extensions {
		exts[i] = extensionNames[ext]
	}
	return types.ScratchProject{
		Targets:    []types.ScratchTarget{p.Stage.Build()},
		Metadata:   p.Metadata,
		Extensions: exts,
	}
}
