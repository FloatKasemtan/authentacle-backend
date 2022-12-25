package user

func (s userService) GetUser(userId string) (*Response, error) {
	user, err := s.userRepository.GetById(userId)
	if err != nil {
		return nil, err
	}

	return &Response{
		Username: user.Username,
		Email:    user.Email,
		IsVerify: user.IsVerify,
	}, nil
}
