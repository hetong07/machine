package provision

import (
	"github.com/hetong07/machine/libmachine/auth"
	"github.com/hetong07/machine/libmachine/engine"
)

type EngineConfigContext struct {
	DockerPort       int
	AuthOptions      auth.Options
	EngineOptions    engine.Options
	DockerOptionsDir string
}
