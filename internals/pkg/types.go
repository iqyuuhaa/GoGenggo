package pkg

import "gogenggo/internals/platforms"

type PkgModules struct {
	Platform *platforms.PlatformModule
}

type PkgWrapper struct {
	Api        apiWrapperInterface
	Dialogflow dialogflowWrapperInterface
	Telegram   telegramWrapperInterface
}
