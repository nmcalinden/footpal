package services

import (
	"github.com/jmoiron/sqlx"
	"github.com/nmcalinden/footpal/models"
	"github.com/nmcalinden/footpal/repository"
)

type VenueService struct {
	venueRepo *repository.VenueRepository
}

func NewVenueService(database *sqlx.DB) *VenueService {
	return &VenueService{venueRepo: repository.NewVenueRepository(database)}
}

func (service *VenueService) GetVenues() (*[]models.Venue, error) {
	return service.venueRepo.FindAll()
}

func (service *VenueService) GetVenueById(venueId *int) (*models.Venue, error) {
	return service.venueRepo.FindById(venueId)
}

func (service *VenueService) GetVenueAdmins(venueId *int) (*[]models.VenueAdmin, error) {
	return service.venueRepo.FindAdminsByVenue(venueId)
}

func (service *VenueService) GetVenuePitches(venueId *int) (*[]models.Pitch, error) {
	return service.venueRepo.FindPitchesByVenue(venueId)
}

func (service *VenueService) GetVenuePitch(venueId *int, pitchId *int) (*models.Pitch, error) {
	return service.venueRepo.FindPitchByVenueIdAndPitchId(venueId, pitchId)
}

func (service *VenueService) GetVenueTimeslots(venueId *int) (*[]models.VenueTimeSlot, error) {
	return service.venueRepo.FindTimeslotsByVenueId(venueId)
}

func (service *VenueService) GetVenuePitchTimeslots(venueId *int, pitchId *int) (*[]models.PitchTimeSlot, error) {
	return service.venueRepo.FindPitchTimeslots(pitchId)
}

func (service *VenueService) CreateNewVenue(venueRequest *models.VenueRequest) (*int, error) {
	newVenue := models.Venue{
		VenueId:  0,
		Name:     venueRequest.Name,
		Address:  venueRequest.Address,
		Postcode: venueRequest.Postcode,
		City:     venueRequest.City,
		PhoneNo:  venueRequest.PhoneNo,
		Email:    venueRequest.Email,
	}
	return service.venueRepo.Save(&newVenue)
}

func (service *VenueService) EditVenue(venueId *int, venueRequest *models.VenueRequest) (*models.Venue, error) {
	v, err := service.venueRepo.FindById(venueId)
	if err != nil {
		return nil, err
	}

	v.Address = venueRequest.Address
	v.Name = venueRequest.Name
	v.Postcode = venueRequest.Postcode
	v.City = venueRequest.City
	v.PhoneNo = venueRequest.PhoneNo
	v.Email = venueRequest.Email
	return service.venueRepo.Update(v)
}

func (service *VenueService) RemoveVenue(venueId *int) error {
	v, err := service.venueRepo.FindById(venueId)
	if err != nil {
		return err
	}

	return service.venueRepo.Delete(&v.VenueId)
}

func (service *VenueService) CreateNewVenueAdmin(venueRequest *models.VenueAdminRequest) (*int, error) {
	newAdmin := models.VenueAdmin{
		VenueAdminId: 0,
		UserId:       venueRequest.UserId,
		VenueId:      venueRequest.VenueId,
	}
	return service.venueRepo.SaveAdmin(&newAdmin)
}

func (service *VenueService) RemoveVenueAdmin(venueId *int, adminId *int) error {
	v, err := service.venueRepo.FindAdminByVenue(venueId, adminId)
	if err != nil {
		return err
	}

	return service.venueRepo.DeleteAdmin(&v.VenueAdminId)
}

func (service *VenueService) CreateNewVenuePitch(venueId *int, pitchRequest *models.PitchRequest) (*int, error) {
	newPitch := models.Pitch{
		PitchId:    0,
		VenueId:    *venueId,
		Name:       pitchRequest.Name,
		MaxPlayers: pitchRequest.MaxPlayers,
		Cost:       pitchRequest.Cost,
	}
	return service.venueRepo.SavePitch(&newPitch)
}

func (service *VenueService) EditVenuePitch(venueId *int, pitchId *int, pitchRequest *models.PitchRequest) (*models.Pitch, error) {
	p, err := service.venueRepo.FindPitchByVenueIdAndPitchId(venueId, pitchId)
	if err != nil {
		return nil, err
	}

	p.Name = pitchRequest.Name
	p.MaxPlayers = pitchRequest.MaxPlayers
	p.Cost = pitchRequest.Cost
	return service.venueRepo.UpdatePitch(p)
}

func (service *VenueService) RemoveVenuePitch(venueId *int, pitchId *int) error {
	p, err := service.venueRepo.FindPitchByVenueIdAndPitchId(venueId, pitchId)
	if err != nil {
		return err
	}

	return service.venueRepo.DeletePitch(&p.PitchId)
}
