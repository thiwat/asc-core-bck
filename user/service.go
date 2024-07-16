package user

import (
	"asc-core/db"
	"asc-core/line"
	"asc-core/types"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

var SESSION_TTL = 30

func Login(input LoginInput) (LoginOutput, error) {

	lineProfile, err := line.GetLineProfile(input.LineToken)

	if err != nil {
		return LoginOutput{}, err
	}

	user, err := FindByUserId(lineProfile.Sub)

	if err != nil {
		user, _ = Create(User{
			UserId:       uuid.NewString(),
			Name:         lineProfile.Name,
			LineId:       lineProfile.Sub,
			ProfileImage: lineProfile.Picture,
			RegisteredAt: time.Now(),
		})
	}

	session := types.Session{
		Name:   user.Name,
		UserId: user.UserId,
		LineId: user.LineId,
	}

	token := uuid.NewString()
	out, _ := json.Marshal(session)

	db.SetKey("TOKEN_"+token, string(out), SESSION_TTL)

	res := LoginOutput{
		Profile: user,
		Token:   token,
	}

	return res, nil
}

func GetProfile(session types.Session) (User, error) {
	profile, err := FindByUserId(session.UserId)
	return profile, err
}
