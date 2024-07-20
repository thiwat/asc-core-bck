package user

import (
	"asc-core/db"
	"asc-core/line"
	"asc-core/types"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

var SESSION_TTL = 30

func Login(input LoginInput) (LoginOutput, error) {

	lineProfile, err := line.GetLineProfile(input.LineToken)

	if err != nil {
		return LoginOutput{}, err
	}

	user, err := findOne(bson.M{"line_id": lineProfile.Sub})

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
	profile, err := findOne(bson.M{"user_id": session.UserId})
	return profile, err
}
