package user

import (
	"asc-core/line"
	"time"

	"github.com/google/uuid"
)

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

	res := LoginOutput{
		Profile: user,
		Token:   "TOKEN",
	}

	return res, nil
}
