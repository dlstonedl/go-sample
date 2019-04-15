package parser

import (
	"github.com/dlstonedl/go-sample/crawler/engine"
	"github.com/dlstonedl/go-sample/crawler/model"
	"regexp"
)

var (
	ageRe       = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([^<]+)岁</div>`)
	marriageRe  = regexp.MustCompile(`<div class="m-btn purple"[^>]*>(未婚|离异)</div>`)
	heightRe    = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([^<]+cm)</div>`)
	weightRe    = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([^<]+kg)</div>`)
	genderRe    = regexp.MustCompile(`"genderString":"(男|女)士"`)
	educationRe = regexp.MustCompile(`"educationString":"([^"]+)"`)
	incomeRe    = regexp.MustCompile(`<div class="m-btn purple"[^>]*>月收入:([^<]+)</div>`)
	nationRe    = regexp.MustCompile(`<div class="m-btn pink"[^>]*>([^>]+族)</div>`)
	carRe       = regexp.MustCompile(`<div class="m-btn pink"[^>]*>(已买车)</div>`)
	houseRe     = regexp.MustCompile(`<div class="m-btn pink"[^>]*>(已购房)</div>`)
	useCityRe   = regexp.MustCompile(`<div class="des f-cl"[^>]*>([^ ]+) [^<]*</div>`)
	idRe        = regexp.MustCompile(`http://album.zhenai.com/u/([0-9]+)`)
)

func ParseProfile(contents []byte, name string, url string) engine.ParseResult {
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
	profile.City = extractString(useCityRe, contents)
	profile.Name = name

	item := engine.Item{}
	item.Url = url
	item.Id = extractString(idRe, []byte(url))
	item.Type = "zhenai"
	item.Data = profile

	result := engine.ParseResult{}
	result.Items = append(result.Items, item)
	return result
}

func extractString(re *regexp.Regexp, contents []byte) string {
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		return string(m[1])
	}
	return ""
}
