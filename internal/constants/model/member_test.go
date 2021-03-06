package model

import (
	"testing"
	"time"

	"github.com/BrunoDM2943/church-members-api/internal/infra/i18n"
	"github.com/stretchr/testify/assert"
)

func TestFormattedContact(t *testing.T) {
	c := Contact{
		CellPhone:     953200587,
		CellPhoneArea: 11,
		Phone:         29435002,
		PhoneArea:     11,
	}
	if "(11) 953200587" != c.GetFormattedCellPhone() {
		t.Fail()
	}

	if "(11) 29435002" != c.GetFormattedPhone() {
		t.Fail()
	}
}

func TestClassification(t *testing.T) {
	now := time.Now()
	t.Run("Children", func(t *testing.T) {
		assert.Equal(t, "Children", Member{
			Person: Person{
				BirthDate: &now,
			},
		}.Classification())
	})
	t.Run("Teen", func(t *testing.T) {
		birthDate := time.Now().AddDate(-17, 0, 0)
		assert.Equal(t, "Teen", Member{
			Person: Person{
				BirthDate: &birthDate,
			},
		}.Classification())
	})
	t.Run("Young", func(t *testing.T) {
		birthDate := time.Now().AddDate(-29, 0, 0)
		assert.Equal(t, "Young", Member{
			Person: Person{
				BirthDate: &birthDate,
			},
		}.Classification())
	})
	t.Run("Adult Single", func(t *testing.T) {
		birthDate := time.Now().AddDate(-33, 0, 0)
		assert.Equal(t, "Adult", Member{
			Person: Person{
				BirthDate: &birthDate,
			},
		}.Classification())
	})
	t.Run("Adult Married", func(t *testing.T) {
		birthDate := time.Now().AddDate(-25, 0, 0)
		assert.Equal(t, "Adult", Member{
			Person: Person{
				BirthDate:    &birthDate,
				MarriageDate: &now,
			},
		}.Classification())
	})
	t.Run("Localized", func(t *testing.T) {
		birthDate := time.Now().AddDate(-25, 0, 0)
		assert.Equal(t, "Adult", Member{
			Person: Person{
				BirthDate:    &birthDate,
				MarriageDate: &now,
			},
		}.ClassificationLocalized(i18n.Localizer))
	})
}

func TestFormattedAddress(t *testing.T) {
	address := Address{
		Address:  "Rua xicas",
		District: "Parque feliz",
		Number:   2,
	}
	assert.Equal(t, "Rua xicas, 2 - Parque feliz", address.GetFormatted())
}

func TestGetFullName(t *testing.T) {
	assert.Equal(t, Person{
		FirstName: "John",
		LastName:  "Doe",
	}.GetFullName(), "John Doe")
}
