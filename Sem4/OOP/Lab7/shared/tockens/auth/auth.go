package auth

import (
	"Lab7/shared/tockens/models"
	"errors"
	"github.com/vk-rv/pvx"
	"time"
)

type PasetoAuth struct {
	pasetoKey    *pvx.SymKey
	symmetricKey []byte
}

const keySize = 32

func NewPasseto(key []byte) (*PasetoAuth, error) {
	if len(key) != keySize {
		return nil, errors.ErrUnsupported
	}

	pasetoKey := pvx.NewSymmetricKey(key, pvx.Version4)

	return &PasetoAuth{
		pasetoKey:    pasetoKey,
		symmetricKey: key,
	}, nil
}

func (p *PasetoAuth) NewTocken(data models.TockenData) (string, error) {
	serviceClaims := &models.ServiceClaims{}
	iss := time.Now()
	exp := iss.Add(data.Duration)

	serviceClaims.IssuedAt = &iss
	serviceClaims.Expiration = &exp
	serviceClaims.Subject = data.Subject

	serviceClaims.AdditionalClaims = data.AdditionalClaims
	serviceClaims.Footer = data.Footer

	pv4 := pvx.NewPV4Local()

	authTocken, err := pv4.Encrypt(p.pasetoKey, serviceClaims, pvx.WithFooter(serviceClaims.Footer))
	if err != nil {
		return "", err
	}

	return authTocken, nil
}

func (p *PasetoAuth) VerifyTocken(tocken string) (*models.ServiceClaims, error) {
	pv4 := pvx.NewPV4Local()
	tk := pv4.Decrypt(tocken, p.pasetoKey)

	f := models.Footer{}
	cs := models.ServiceClaims{
		Footer: f,
	}

	err := tk.Scan(&cs, &f)
	if err != nil {
		return &cs, err
	}
	return &cs, nil
}
