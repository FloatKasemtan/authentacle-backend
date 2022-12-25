package application

import "errors"

func (a applicationService) GetApp(id string, userId string) (*ApplicationResponse, error) {
	app, err := a.applicationRepository.GetAppById(id)
	if err != nil {
		return nil, err
	}
	if app.UserId.Hex() != userId {
		return nil, errors.New("Permission denied")
	}

	return &ApplicationResponse{Name: app.Name, Logo: app.Logo, Key: app.Key}, nil
}
