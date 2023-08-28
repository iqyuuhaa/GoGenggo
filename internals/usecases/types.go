package usecases

import (
	"gogenggo/internals/pkg"
	"gogenggo/internals/platforms"
)

type UsecaseModules struct {
	Pkg      *pkg.PkgWrapper
	Platform *platforms.PlatformModule
}

type UsecaseWrapper struct {
	Chat chatWrapperInterface
}
