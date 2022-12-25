package application

func (a appRepositoryDB) AddApp(application *Application) error {
	if err := a.coll.Create(application); err != nil {
		return err
	}

	return nil
}
