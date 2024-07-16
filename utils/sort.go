package utils

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func BuildSort(sort string) bson.M {
	var res = bson.M{}

	if sort == "" {
		return res
	}

	sorts := strings.Split(sort, ",")

	for _, word := range sorts {
		if word[0:1] == "-" {
			res[word[1:]] = -1
		} else {
			res[word] = 1
		}
	}

	return res
}
