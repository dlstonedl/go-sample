package parser

import (
	"github.com/dlstonedl/go-sample/crawler/engine"
	"github.com/dlstonedl/go-sample/crawler/model"
	"regexp"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([^<]+)岁</div>`)
var marriageRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>(未婚|离异)</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([^<]+cm)</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([^<]+kg)</div>`)
var genderRe = regexp.MustCompile(`"genderString":"(男|女)士"`)
var educationRe = regexp.MustCompile(`"educationString":"([^"]+)"`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>月收入:([^<]+)</div>`)
var nationRe = regexp.MustCompile(`<div class="m-btn pink"[^>]*>([^>]+族)</div>`)
var carRe = regexp.MustCompile(`<div class="m-btn pink"[^>]*>(已买车)</div>`)
var houseRe = regexp.MustCompile(`<div class="m-btn pink"[^>]*>(已购房)</div>`)

func ParseProfile(contents []byte, name string, city string) engine.ParseResult {
	result := engine.ParseResult{}

	profile := model.Profile{}
	profile.Age = extractString(ageRe, contents)
	profile.Marriage = extractString(marriageRe, contents)
	profile.Height = extractString(heightRe, contents)
	profile.Weight = extractString(weightRe, contents)
	profile.Gender = extractString(genderRe, contents)
	profile.Education = extractString(educationRe, contents)
	profile.Income = extractString(incomeRe, contents)
	profile.Car = extractString(carRe, contents)
	profile.House = extractString(houseRe, contents)
	profile.Nation = extractString(nationRe, contents)
	profile.Name = name
	profile.City = city

	result.Items = append(result.Items, profile)
	return result
}

func extractString(re *regexp.Regexp, contents []byte) string {
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		return string(m[1])
	}
	return ""
}
