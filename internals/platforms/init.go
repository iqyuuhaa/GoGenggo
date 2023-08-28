package platforms

import "gogenggo/internals/platforms/db"

var platformObjects *PlatformModule

func Init() (*PlatformModule, error) {
	if platformObjects != nil {
		return platformObjects, nil
	}

	allPlatforms := new(PlatformModule)
	allPlatforms.DB = db.Init()

	platformObjects = allPlatforms

	return allPlatforms, nil
}
