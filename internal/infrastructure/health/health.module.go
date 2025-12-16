package health

import "github.com/dangduoc08/ginject/core"

var HealthModule = func() *core.Module {
	healthController := HealthController{}

	var module = core.ModuleBuilder().
		Imports().
		Controllers(healthController).
		Build()

	module.
		Prefix("health")

	return module
}
