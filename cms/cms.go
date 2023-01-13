package cms

type CmsApp struct{}

func (app *CmsApp) Run() error {
	return nil
}

func New() CmsApp {
	return CmsApp{}
}
