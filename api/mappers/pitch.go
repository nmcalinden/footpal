package mappers

import (
	"github.com/nmcalinden/footpal/api/models"
	"github.com/nmcalinden/footpal/api/utils"
	"github.com/nmcalinden/footpal/api/views"
	"time"
)

func MapPitchSlotsToOpeningHours(pts []models.PitchTimeSlot) (*[]views.VenueOpeningHour, error) {
	daysOfWeek := [7]time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday, time.Sunday}
	var res []views.VenueOpeningHour

	for _, w := range daysOfWeek {
		t := getPitchSlotsByDay(pts, w)
		vo := buildDayOpeningHours(t, w)
		res = append(res, vo)
	}
	return &res, nil
}

func getPitchSlotsByDay(pts []models.PitchTimeSlot, d time.Weekday) []models.PitchTimeSlot {
	var t []models.PitchTimeSlot
	for _, pt := range pts {
		if pt.DayOfWeek == d.String() {
			t = append(t, pt)
		}
	}

	return t
}

func buildDayOpeningHours(t []models.PitchTimeSlot, d time.Weekday) views.VenueOpeningHour {
	o := ""
	c := ""
	if len(t) > 0 {
		o = utils.GetFormattedTime(t[0].StartTime)
		c = utils.GetFormattedTime(t[len(t)-1].EndTime)
	}

	return views.VenueOpeningHour{
		DayOfWeek: d.String(),
		Open:      o,
		Close:     c,
	}
}
