package internals

import (
	"log"

	"gogenggo/internals/pkg"
	"gogenggo/internals/platforms"
	"gogenggo/internals/usecases"
)

var internalObjects *InternalModules

func InitAllInternals() (*InternalModules, error) {
	if internalObjects != nil {
		return internalObjects, nil
	}

	allPlatforms, err := platforms.Init()
	if err != nil {
		log.Fatalln("[Internals - InitAllInternals] Error initializing platforms, err: ", err)
		return nil, err
	}

	pkgModule := new(pkg.PkgModules)
	pkgModule.Platform = allPlatforms
	allPkg, err := pkg.Init(pkgModule)
	if err != nil {
		log.Fatalln("[Internals - InitAllInternals] Error initializing pkg, err: ", err)
		return nil, err
	}

	usecaseModule := new(usecases.UsecaseModules)
	usecaseModule.Pkg = allPkg
	usecaseModule.Platform = allPlatforms
	allUsecase, err := usecases.Init(usecaseModule)
	if err != nil {
		log.Fatalln("[Internals - InitAllInternals] Error initializing usecases, err: ", err)
		return nil, err
	}

	allInternal := new(InternalModules)
	allInternal.Usecase = allUsecase
	allInternal.Pkg = allPkg
	allInternal.Platform = allPlatforms

	internalObjects = allInternal

	return internalObjects, nil
}
