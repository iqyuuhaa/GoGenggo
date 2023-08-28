package handlers

import "gogenggo/internals/usecases"

func Init(usecase *usecases.UsecaseWrapper) (*Handler, error) {
	handler := new(Handler)
	handler.usecases = usecase

	return handler, nil
}
