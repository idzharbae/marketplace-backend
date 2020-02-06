package app

type Repos struct {
}

func NewRepos() *Repos {
	return &Repos{}
}

func (r *Repos) Close() []error {
	var errs []error

	return errs
}
