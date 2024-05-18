package services

import "genesis-ses4/models"

var subscriptions subscriptionsService
var Subscriptions = &subscriptions

type subscriptionsService struct{}

func (s subscriptionsService) IsSubscribed(email string) (bool, error) {
	var count int64

	res := models.DB.Model(&models.Subscriber{}).Where("email = ?", email).Count(&count)

	if res.Error != nil {
		return false, res.Error
	}

	return count != 0, nil
}

func (s subscriptionsService) AddSubscriber(email string) (models.Subscriber, error) {
	new_sub := models.Subscriber{Email: email}

	res := models.DB.Create(&new_sub)

	if res.Error != nil {
		return new_sub, res.Error
	}

	return new_sub, nil
}

func (s subscriptionsService) GetSubscribersEmails() ([]string, error) {
	var subs []models.Subscriber

	res := models.DB.Find(&subs)
	if res.Error != nil {
		return []string{}, res.Error
	}

	emails := make([]string, len(subs))

	for i, sub := range subs {
		emails[i] = sub.Email
	}

	return emails, nil
}
