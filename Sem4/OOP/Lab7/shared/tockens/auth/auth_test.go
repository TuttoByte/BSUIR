package auth

import (
	"Lab7/shared/tockens/models"
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestServiceGenerateNewKey(t *testing.T) {
	key := []byte("000f3e5799296cc4ce32c444cfde4962")
	pasettoToken, err := NewPasseto(key)
	require.NoError(t, err)

	token, err := pasettoToken.NewTocken(
		models.TockenData{
			Subject:  "test",
			Duration: 5 * time.Second,
			AdditionalClaims: models.AdditionalClaims{
				Name: "Nikolas",
				Role: "Admin",
			},
			Footer: models.Footer{
				"footer",
			},
		})

	require.NoError(t, err)
	require.NotEmpty(t, token)

	sc, err := pasettoToken.VerifyTocken(token)
	require.NoError(t, err)
	require.Equal(t, "footer", sc.Footer.MetaData)
	require.Equal(t, "Nikolas", sc.AdditionalClaims.Name)
	require.Equal(t, "Admin", sc.AdditionalClaims.Role)
}

func TestInvalidToken(t *testing.T) {
	badKey := []byte("00")
	pasetoToken, err := NewPasseto(badKey)
	require.ErrorIs(t, err, errors.ErrUnsupported)

	key := []byte("000f3e5799296cc4ce32c444cfde4962")
	pasetoToken, err = NewPasseto(key)
	require.NoError(t, err)

	token, err := pasetoToken.NewTocken(models.TockenData{
		Duration: -5 * time.Second,
	})

	require.NoError(t, err)

	sc, err := pasetoToken.VerifyTocken(token)
	require.Error(t, err)
	require.Error(t, sc.Valid())
}
