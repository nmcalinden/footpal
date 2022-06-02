package services

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/nmcalinden/footpal/api/models"
	"github.com/nmcalinden/footpal/api/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type BookingRepositoryMock struct {
	mock.Mock
}

func (r *BookingRepositoryMock) Save(booking *models.Booking, matches *[]models.Match, pitchSlots *[]models.PitchSlot) (*int, error) {
	//TODO implement me
	panic("implement me")
}

func (r *BookingRepositoryMock) FindAvailableVenues(venueId *int, matchDate string, city *string, players *int) (*[]models.Venue, error) {
	//TODO implement me
	panic("implement me")
}

func (r *BookingRepositoryMock) IsExistingMatchPresent(matchDate *string, pitchTimeslotId *int) (*bool, error) {
	//TODO implement me
	panic("implement me")
}

func (r *BookingRepositoryMock) FindAll() (*[]models.Booking, error) {
	var bookingRecords []models.Booking

	for i := 1; i < 4; i++ {
		bookingId := i
		b := models.Booking{
			BookingId:       &bookingId,
			BookingStatusId: 4,
			CreatedBy:       i,
			Created:         time.Now(),
			LastUpdated:     time.Now(),
		}

		bookingRecords = append(bookingRecords, b)
	}
	return &bookingRecords, nil
}

func (r *BookingRepositoryMock) FindById(id *int) (*models.Booking, error) {
	b := models.Booking{
		BookingId:       id,
		BookingStatusId: 4,
		CreatedBy:       1,
		Created:         time.Time{},
		LastUpdated:     time.Time{},
	}
	return &b, nil
}

func (r *BookingRepositoryMock) FindMatchesByBookingId(id *int) (*[]models.Match, error) {
	var matches []models.Match
	return &matches, nil
}

func (r *BookingRepositoryMock) Update(booking *models.Booking) (*models.Booking, error) {
	return booking, nil
}

func TestBookingService_GetBookings(t *testing.T) {
	repo := BookingRepositoryMock{}
	repo.On("FindAll").Return([]models.Booking{}, nil)

	service := BookingService{&repo}
	got, err := service.GetBookings()
	if err != nil {
		assert.Fail(t, "Error", err)
	}

	for i, b := range *got {
		expectedId := i + 1
		assert.Equal(t, expectedId, *b.BookingId, fmt.Sprintf("Booking ID must equal: %d", expectedId))
		assert.True(t, b.BookingStatusId != 0, "Booking Status ID should not be zero")
	}
}

func TestBookingService_GetBookingById(t *testing.T) {
	id := 1
	repo := BookingRepositoryMock{}
	repo.On("FindById").Return(models.Booking{}, nil)

	service := BookingService{&repo}
	got, err := service.GetBookingById(&id)
	if err != nil {
		assert.Fail(t, "Error")
	}

	assert.Equal(t, id, *got.BookingId, fmt.Sprintf("Booking ID must equal: %d", id))
	assert.True(t, got.BookingStatusId != 0, "Booking Status ID should not be zero")
}

func TestBookingService_GetBookingByIdMock(t *testing.T) {
	id := 1
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockBookingRepo := mocks.NewMockBookingRepositoryI(mockCtrl)
	mockBookingRepo.EXPECT().FindById(&id).Return(&models.Booking{
		BookingId:       &id,
		BookingStatusId: 4,
		CreatedBy:       1,
		Created:         time.Now(),
		LastUpdated:     time.Now(),
	}, nil)

	bService := NewBookingService(mockBookingRepo)
	result, err := bService.GetBookingById(&id)
	if err != nil {
		t.Errorf("Error getting booking: %v", err)
	}

	assert.Equal(t, id, *result.BookingId, fmt.Sprintf("Booking ID must equal: %d", id))
	assert.True(t, result.BookingStatusId != 0, "Booking Status ID should not be zero")
}

func TestBookingService_GetBookingById_HandleError(t *testing.T) {
	id := 1
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	expectedError := errors.New("no results")
	mockBookingRepo := mocks.NewMockBookingRepositoryI(mockCtrl)
	mockBookingRepo.EXPECT().FindById(&id).Return(nil, expectedError)

	bService := NewBookingService(mockBookingRepo)
	result, err := bService.GetBookingById(&id)

	assert.True(t, result == nil, "No bookings returned")
	assert.True(t, err != nil, "Error expected")
	assert.True(t, errors.Is(err, expectedError))
}
