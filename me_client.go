package vimeo

// Me client
type MeClient struct{}

// View one user
func (c *MeClient) ShowInfomation() (*Me, error) {
	user := Me{}
	err := query("GET", "/me", nil, &user)
	if err == nil  {
		return &user, err
	}
	return &user, err
}
