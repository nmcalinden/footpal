package mappers

import (
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/nmcalinden/footpal/api/models"
	"github.com/nmcalinden/footpal/api/views"
	"gopkg.in/jeevatkm/go-model.v1"
)

func MapToVenueSummaries(v []models.Venue) (*[]views.VenueSummary, error) {
	var errs error
	var res []views.VenueSummary
	for _, vs := range v {
		ven := views.VenueSummary{}
		mErrs := model.Copy(&ven, vs)
		if mErrs != nil {
			errs = multierror.Append(errs, mErrs...)
		}
		res = append(res, ven)
	}

	return &res, errs
}

func MapToVenueView(v models.Venue, p []models.Pitch) (*views.Venue, error) {
	var vP []views.VenuePitchSummary
	for _, ps := range p {
		vP = append(vP, buildPitchSummary(v, ps))
	}

	var errs error
	venue := views.Venue{}
	mErrs := model.Copy(&venue, v)
	if mErrs != nil {
		errs = multierror.Append(errs, mErrs...)
		return nil, errs
	}

	venue.Pitches = vP
	return &venue, nil
}

func buildPitchSummary(v models.Venue, ps models.Pitch) views.VenuePitchSummary {
	ref := fmt.Sprintf("/venues/%d/pitches/%d", *v.VenueId, *ps.PitchId)
	r := views.VenuePitchSummary{PitchId: ps.PitchId, Name: ps.Name, Href: ref}
	return r
}
