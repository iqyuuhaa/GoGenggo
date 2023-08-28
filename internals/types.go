package internals

import (
	"gogenggo/internals/pkg"
	"gogenggo/internals/platforms"
	"gogenggo/internals/usecases"
)

type InternalModules struct {
	Usecase  *usecases.UsecaseWrapper
	Pkg      *pkg.PkgWrapper
	Platform *platforms.PlatformModule
}
