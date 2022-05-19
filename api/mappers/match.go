package mappers

import (
	"github.com/hashicorp/go-multierror"
	"github.com/nmcalinden/footpal/api/models"
	"github.com/nmcalinden/footpal/api/views"
	"gopkg.in/jeevatkm/go-model.v1"
)

const (
	layoutTime = "15:04"
)

func MapToMatchSummaryView(dest *views.MatchSummary, mId int, mb models.MatchBookingDetail, sq models.Squad) error {
	var errs error
	mErrs := model.Copy(dest, mb)
	if mErrs != nil {
		errs = multierror.Append(errs, mErrs...)
		return errs
	}

	dest.Time = mb.StartTime.Format(layoutTime)
	if sq.SquadId != nil {
		dest.SquadId = sq.SquadId
		dest.SquadName = &sq.Name
	}

	dest.MatchId = &mId
	return nil
}
