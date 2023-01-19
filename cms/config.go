package cms

import "errors"

const DefaultAdminRootUrl = "admin"

type CmsConfig struct {
	AdminRootUrl string
}

func (c *CmsConfig) validate() error {
	if c == nil {
		return errors.New("cms: config cannot be nil")
	}

	if c.AdminRootUrl == "" {
		c.AdminRootUrl = DefaultAdminRootUrl
	}

	return nil
}
