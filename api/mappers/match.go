package mappers

import (
	"github.com/hashicorp/go-multierror"
	"github.com/nmcalinden/footpal/models"
	"github.com/nmcalinden/footpal/utils"
	"github.com/nmcalinden/footpal/views"
	"gopkg.in/jeevatkm/go-model.v1"
)

func MapToMatchSummaryView(dest *views.MatchSummary, mId int, mb models.MatchBookingDetail, sq models.Squad) error {
	var errs error
	mErrs := model.Copy(dest, mb)
	if mErrs != nil {
		errs = multierror.Append(errs, mErrs...)
		return errs
	}

	dest.Time = utils.GetFormattedTime(mb.StartTime)
	if sq.SquadId != nil {
		dest.SquadId = sq.SquadId
		dest.SquadName = &sq.Name
	}

	dest.MatchId = &mId
	return nil
}
