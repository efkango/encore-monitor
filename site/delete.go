package site

import "context"

//encore:api public method=DELETE path=/site/:siteID
func (s *Service) Delete(ctx context.Context, siteID int) error {
	return s.db.Delete(&Site{ID: siteID}).Error
}
