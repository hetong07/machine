package provision

import (
	"testing"

	"github.com/hetong07/machine/drivers/fakedriver"
	"github.com/hetong07/machine/libmachine/auth"
	"github.com/hetong07/machine/libmachine/engine"
	"github.com/hetong07/machine/libmachine/provision/provisiontest"
	"github.com/hetong07/machine/libmachine/swarm"
)

func TestRedHatDefaultStorageDriver(t *testing.T) {
	p := NewRedHatProvisioner("", &fakedriver.Driver{})
	p.SSHCommander = provisiontest.NewFakeSSHCommander(provisiontest.FakeSSHCommanderOptions{})
	p.Provision(swarm.Options{}, auth.Options{}, engine.Options{})
	if p.EngineOptions.StorageDriver != "overlay2" {
		t.Fatal("Default storage driver should be overlay2")
	}
}
