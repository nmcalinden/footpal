package mappers

import (
	"fmt"

	"github.com/hashicorp/go-multierror"
	"github.com/nmcalinden/footpal/models"
	"github.com/nmcalinden/footpal/views"
	"gopkg.in/jeevatkm/go-model.v1"
)

func MapToUserProfile(dest *views.UserProfile, mp models.Player, u models.User) error {
	mErr := model.Copy(dest, mp)
	if mErr != nil {
		var rErr error
		return multierror.Append(rErr, mErr...)
	}
	dest.Id = u.UserId
	dest.Name = fmt.Sprintf("%s %s", u.Forename, u.Surname)
	dest.Email = u.Email
	return nil
}

func MapToPlayer(dest *views.PlayerUser, mp models.Player, u models.User) error {
	mErr := model.Copy(dest, mp)
	if mErr != nil {
		var rErr error
		return multierror.Append(rErr, mErr...)
	}
	dest.Name = fmt.Sprintf("%s %s", u.Forename, u.Surname)
	dest.Email = u.Email
	return nil
}
