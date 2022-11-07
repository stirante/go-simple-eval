package functions

import (
	"github.com/stirante/go-simple-query-language/query/utils"
)

func RegisterSemverFunctions() {
	RegisterFunction(JsonFunction{
		Name: "semver",
		Body: semverString,
	})
	RegisterFunction(JsonFunction{
		Name: "semver",
		Body: semverArray,
	})
	RegisterFunction(JsonFunction{
		Name: "semver",
		Body: semver,
	})
}

func semverString(str string) (utils.Semver, error) {
	return utils.ParseSemverString(str)
}

func semverArray(arr []interface{}) (utils.Semver, error) {
	return utils.ParseSemverArray(arr)
}

func semver(major, minor, patch utils.JsonNumber) (utils.Semver, error) {
	return utils.Semver{Major: int(major.IntValue()), Minor: int(minor.IntValue()), Patch: int(patch.IntValue())}, nil
}
