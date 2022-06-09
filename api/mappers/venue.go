package mappers

import (
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/nmcalinden/footpal/api/models"
	"github.com/nmcalinden/footpal/api/utils"
	"github.com/nmcalinden/footpal/api/views"
	"gopkg.in/jeevatkm/go-model.v1"
	"time"
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

func MapToPitchSlotsByVenue(dest *views.PitchBookingDetails, b []models.VenueTimeSlot, pts []models.PitchTimeSlot, fd string, d string) error {
	var l []views.PitchTimeSlotBooking

	ct := time.Now()
	x, _ := time.Parse("2006-01-02", fd)
	isToday := ct.Truncate(24 * time.Hour).Equal(x.Truncate(24 * time.Hour))
	for _, pt := range pts {
		k := getPitchTimeSlotBooking(pt, b, ct, isToday)
		if k != nil {
			l = append(l, *k)
		}
	}

	dest.MatchDate = fd
	dest.DayOfWeek = d
	dest.TimeSlots = l
	return nil
}

func MapToPitchTimeslotView(dest *views.PitchTimeSlot, p models.Pitch, pts models.PitchTimeSlot) error {

	pi := views.Pitch{}
	mErrs := model.Copy(&pi, p)
	if mErrs != nil {
		var rErr error
		return multierror.Append(rErr, mErrs...)
	}

	st := utils.GetFormattedTime(pts.StartTime)
	et := utils.GetFormattedTime(pts.EndTime)
	pv := views.TimeSlot{
		PitchTimeSlotId: pts.PitchTimeSlotId,
		DayOfWeek:       pts.DayOfWeek,
		StartTime:       st,
		EndTime:         et,
	}

	dest.Pitch = pi
	dest.TimeSlot = pv
	return nil
}

func getPitchTimeSlotBooking(pts models.PitchTimeSlot, b []models.VenueTimeSlot, ct time.Time, isToday bool) *views.PitchTimeSlotBooking {
	ctH, _, _ := ct.Clock()
	ptsH, _, _ := pts.StartTime.Clock()
	if ptsH <= ctH && isToday {
		return nil
	}

	st := utils.GetFormattedTime(pts.StartTime)
	et := utils.GetFormattedTime(pts.EndTime)

	k := views.PitchTimeSlotBooking{
		PitchTimeSlotId: pts.PitchTimeSlotId,
		StartTime:       st,
		EndTime:         et,
		IsBooked:        isSlotBooked(pts, b),
	}

	return &k
}
func isSlotBooked(pts models.PitchTimeSlot, b []models.VenueTimeSlot) bool {
	var o = false
	for _, bs := range b {
		if pts.PitchTimeSlotId == bs.PitchTimeSlotId && pts.DayOfWeek == bs.DayOfWeek {
			o = true
			break
		}
	}

	return o
}

func buildPitchSummary(v models.Venue, ps models.Pitch) views.VenuePitchSummary {
	ref := fmt.Sprintf("/venues/%d/pitches/%d", *v.VenueId, *ps.PitchId)
	r := views.VenuePitchSummary{PitchId: ps.PitchId, Name: ps.Name, Href: ref}
	return r
}
