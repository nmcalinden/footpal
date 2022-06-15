package mappers

import (
	"fmt"

	"github.com/hashicorp/go-multierror"
	"github.com/nmcalinden/footpal/models"
	"github.com/nmcalinden/footpal/views"
	"gopkg.in/jeevatkm/go-model.v1"
)

func MapToPlayerView(dest *views.Player, mp models.Player, u models.User) error {
	mErr := model.Copy(dest, mp)
	if mErr != nil {
		var rErr error
		return multierror.Append(rErr, mErr...)
	}
	dest.Name = fmt.Sprintf("%s %s", u.Forename, u.Surname)
	return nil
}
