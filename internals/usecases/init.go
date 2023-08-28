package usecases

var usecaseObjects *UsecaseWrapper

func Init(usecaseWrapper *UsecaseModules) (*UsecaseWrapper, error) {
	if usecaseObjects != nil {
		return usecaseObjects, nil
	}

	allUsecase := new(UsecaseWrapper)
	allUsecase.Chat = usecaseWrapper

	usecaseObjects = allUsecase

	return usecaseObjects, nil
}
