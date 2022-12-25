package application

func (a applicationService) GetAllApps(id string) ([]*ApplicationsResponse, error) {

	applications, err := a.applicationRepository.GetAllAppsByUserId(id)
	if err != nil {
		return nil, err
	}

	var applicationResponses []*ApplicationsResponse
	for _, app := range applications {
		applicationResponses = append(applicationResponses, &ApplicationsResponse{Name: app.Name, Logo: app.Logo})
	}

	return applicationResponses, nil
}
