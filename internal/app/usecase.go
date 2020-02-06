package app

type UseCases struct {
}

func NewUsecase() *UseCases {
	return &UseCases{}
}

func (ucs *UseCases) Close() []error {
	var errs []error

	return errs
}
