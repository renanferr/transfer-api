package jwt

import (
	"errors"
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/pedroyremolo/transfer-api/pkg/authenticating"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var InvalidLoginErr = errors.New("it seems your login credentials are invalid, verify them and try again")

type Gatekeeper struct {
	hs  *jwt.HMACSHA
	iss string
}

func NewGatekeeper(tokenSecret string, issuer string) *Gatekeeper {
	return &Gatekeeper{
		hs:  jwt.NewHS256([]byte(tokenSecret)),
		iss: issuer,
	}
}

func (g *Gatekeeper) Sign(login authenticating.Login, secretDigest string, clientID string) (authenticating.Token, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(secretDigest), []byte(login.Secret)); err != nil {
		return authenticating.Token{}, InvalidLoginErr
	}

	currentTime := time.Now().UTC()
	id := primitive.NewObjectID()
	token, err := jwt.Sign(Token{
		Payload: jwt.Payload{
			Issuer:         g.iss,
			ExpirationTime: jwt.NumericDate(currentTime.Add(time.Minute * 30)),
			IssuedAt:       jwt.NumericDate(currentTime),
			JWTID:          id.Hex(),
		},
		ClientID: clientID,
	}, g.hs)
	if err != nil {
		return authenticating.Token{}, err
	}

	return authenticating.Token{ID: id, ClientID: clientID, Digest: string(token)}, nil
}

func (g *Gatekeeper) Verify(_ string) authenticating.Token {
	panic("implement me")
}
