package services

import (
	"github.com/nmcalinden/footpal/api/mappers"
	"github.com/nmcalinden/footpal/api/models"
	"github.com/nmcalinden/footpal/api/payloads"
	"github.com/nmcalinden/footpal/api/repository"
	"github.com/nmcalinden/footpal/api/views"
	"log"
	"time"
)

type VenueService struct {
	venueRepo repository.VenueRepositoryI
	pitchRepo repository.PitchRepositoryI
}

func NewVenueService(venueRepo repository.VenueRepositoryI, pitchRepo repository.PitchRepositoryI) *VenueService {
	return &VenueService{venueRepo: venueRepo, pitchRepo: pitchRepo}
}

func (s *VenueService) GetVenues() (*[]views.Venue, error) {
	venues, err := s.venueRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var res []views.Venue
	for _, v := range *venues {
		m, err := s.getVenueResponse(v)
		if err != nil {
			return nil, err
		}
		res = append(res, *m)
	}

	return &res, nil
}

func (s *VenueService) GetVenueSummaries() (*[]views.VenueSummary, error) {
	venues, err := s.venueRepo.FindAll()
	if err != nil {
		return nil, err
	}

	res, err := mappers.MapToVenueSummaries(*venues)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *VenueService) GetVenueById(venueId *int) (*views.Venue, error) {
	v, err := s.venueRepo.FindById(venueId)
	if err != nil {
		return nil, err
	}
	return s.getVenueResponse(*v)
}

func (s *VenueService) GetVenueAdmins(venueId *int) (*[]models.VenueAdmin, error) {
	return s.venueRepo.FindAdminsByVenue(venueId)
}

func (s *VenueService) GetVenuePitches(venueId *int) (*[]models.Pitch, error) {
	return s.venueRepo.FindPitchesByVenue(venueId)
}

func (s *VenueService) GetVenuePitch(venueId *int, pitchId *int) (*models.Pitch, error) {
	return s.venueRepo.FindPitchByVenueIdAndPitchId(venueId, pitchId)
}

func (s *VenueService) GetVenueTimeslots(v *int, f string, t string) (*[]views.PitchBookingDetails, error) {
	fd, err := time.Parse("2006-01-02", f)

	if t == "" {
		t = f
	}
	td, err := time.Parse("2006-01-02", t)

	b, err := s.venueRepo.FindTimeslotsByVenueIdAndDateRange(*v, f, t)
	if err != nil {
		return nil, err
	}

	var bd []views.PitchBookingDetails
	for fd.Before(td) {
		dayOfWeek := fd.Weekday().String()

		p, err := s.pitchRepo.FindAllByVenueIdAndDay(*v, dayOfWeek)
		if err != nil {
			return nil, err
		}

		var res views.PitchBookingDetails
		err = mappers.MapToPitchSlotsByVenue(&res, *b, *p, fd.Format("2006-01-02"), dayOfWeek)

		if err != nil {
			return nil, err
		}

		bd = append(bd, res)
		fd = fd.AddDate(0, 0, 1)
	}

	return &bd, nil
}

func (s *VenueService) GetVenuePitchTimeslots(venueId *int, pitchId *int) (*[]models.PitchTimeSlot, error) {
	return s.venueRepo.FindPitchTimeslots(pitchId)
}

func (s *VenueService) CreateNewVenue(venueRequest *payloads.VenueRequest) (*int, error) {
	newVenue := models.Venue{
		Name:     venueRequest.Name,
		Address:  venueRequest.Address,
		Postcode: venueRequest.Postcode,
		City:     venueRequest.City,
		PhoneNo:  venueRequest.PhoneNo,
		Email:    venueRequest.Email,
	}
	return s.venueRepo.Save(&newVenue)
}

func (s *VenueService) EditVenue(venueId *int, venueRequest *payloads.VenueRequest) (*models.Venue, error) {
	v, err := s.venueRepo.FindById(venueId)
	if err != nil {
		return nil, err
	}

	v.Address = venueRequest.Address
	v.Name = venueRequest.Name
	v.Postcode = venueRequest.Postcode
	v.City = venueRequest.City
	v.PhoneNo = venueRequest.PhoneNo
	v.Email = venueRequest.Email
	return s.venueRepo.Update(v)
}

func (s *VenueService) RemoveVenue(venueId *int) error {
	v, err := s.venueRepo.FindById(venueId)
	if err != nil {
		return err
	}

	return s.venueRepo.Delete(v.VenueId)
}

func (s *VenueService) CreateNewVenueAdmin(venueRequest *payloads.VenueAdminRequest) (*int, error) {
	newAdmin := models.VenueAdmin{
		UserId:  venueRequest.UserId,
		VenueId: venueRequest.VenueId,
	}
	return s.venueRepo.SaveAdmin(&newAdmin)
}

func (s *VenueService) RemoveVenueAdmin(venueId *int, adminId *int) error {
	v, err := s.venueRepo.FindAdminByVenue(venueId, adminId)
	if err != nil {
		return err
	}

	return s.venueRepo.DeleteAdmin(v.VenueAdminId)
}

func (s *VenueService) CreateNewVenuePitch(venueId *int, pitchRequest *payloads.PitchRequest) (*int, error) {
	newPitch := models.Pitch{
		VenueId:    *venueId,
		Name:       pitchRequest.Name,
		MaxPlayers: pitchRequest.MaxPlayers,
		Cost:       pitchRequest.Cost,
	}
	return s.venueRepo.SavePitch(&newPitch)
}

func (s *VenueService) EditVenuePitch(venueId *int, pitchId *int, pitchRequest *payloads.PitchRequest) (*models.Pitch, error) {
	p, err := s.venueRepo.FindPitchByVenueIdAndPitchId(venueId, pitchId)
	if err != nil {
		return nil, err
	}

	p.Name = pitchRequest.Name
	p.MaxPlayers = pitchRequest.MaxPlayers
	p.Cost = pitchRequest.Cost
	return s.venueRepo.UpdatePitch(p)
}

func (s *VenueService) RemoveVenuePitch(venueId *int, pitchId *int) error {
	p, err := s.venueRepo.FindPitchByVenueIdAndPitchId(venueId, pitchId)
	if err != nil {
		return err
	}

	return s.venueRepo.DeletePitch(p.PitchId)
}

func (s *VenueService) getVenueResponse(v models.Venue) (*views.Venue, error) {
	p, err := s.venueRepo.FindPitchesByVenue(v.VenueId)
	if err != nil {
		return nil, err
	}
	m, errs := mappers.MapToVenueView(v, *p)

	if errs != nil {
		log.Println(errs.Error())
		return nil, errs
	}
	return m, nil
}

func (s *VenueService) GetVenueOpeningHours(v *int) (*[]views.VenueOpeningHour, error) {
	p, err := s.pitchRepo.FindAllByVenueId(*v)
	if err != nil {
		return nil, err
	}

	return mappers.MapPitchSlotsToOpeningHours(*p)
}
