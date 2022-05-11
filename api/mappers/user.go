package mappers

import (
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/nmcalinden/footpal/api/models"
	"github.com/nmcalinden/footpal/api/views"
	"gopkg.in/jeevatkm/go-model.v1"
)

func MapToUser(dest *views.PlayerUser, mp models.Player, u models.User) error {
	mErr := model.Copy(dest, mp)
	if mErr != nil {
		var rErr error
		return multierror.Append(rErr, mErr...)
	}
	dest.Name = fmt.Sprintf("%s %s", u.Forename, u.Surname)
	dest.Email = u.Email
	return nil
}
