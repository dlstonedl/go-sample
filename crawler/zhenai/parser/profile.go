package parser

import (
	"github.com/dlstonedl/go-sample/crawler/engine"
	"github.com/dlstonedl/go-sample/crawler/model"
	"log"
	"regexp"
	"strconv"
)

var (
	ageRe       = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([^<]+)岁</div>`)
	marriageRe  = regexp.MustCompile(`<div class="m-btn purple"[^>]*>(未婚|离异)</div>`)
	heightRe    = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([^<]+)cm</div>`)
	weightRe    = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([^<]+)kg</div>`)
	genderRe    = regexp.MustCompile(`"genderString":"(男|女)士"`)
	educationRe = regexp.MustCompile(`"educationString":"([^"]+)"`)
	incomeRe    = regexp.MustCompile(`<div class="m-btn purple"[^>]*>月收入:([^<]+)</div>`)
	nationRe    = regexp.MustCompile(`<div class="m-btn pink"[^>]*>([^>]+族)</div>`)
	carRe       = regexp.MustCompile(`<div class="m-btn pink"[^>]*>(已买车)</div>`)
	houseRe     = regexp.MustCompile(`<div class="m-btn pink"[^>]*>(已购房)</div>`)
	useCityRe   = regexp.MustCompile(`<div class="des f-cl"[^>]*>([^ ]+) [^<]*</div>`)
	idRe        = regexp.MustCompile(`http://album.zhenai.com/u/([0-9]+)`)
)

var profileCount = 0

func ParseProfile(contents []byte, url string, name string) engine.ParseResult {
	log.Printf("profileCount is #%d\n", profileCount)
	profileCount++

	profile := model.Profile{}
	profile.Age = convertStringToInt(extractString(ageRe, contents))
	profile.Marriage = extractString(marriageRe, contents)
	profile.Height = convertStringToInt(extractString(heightRe, contents))
	profile.Weight = convertStringToInt(extractString(weightRe, contents))
	profile.Gender = extractString(genderRe, contents)
	profile.Education = extractString(educationRe, contents)
	profile.Income = extractString(incomeRe, contents)
	profile.Car = extractString(carRe, contents)
	profile.House = extractString(houseRe, contents)
	profile.Nation = extractString(nationRe, contents)
	profile.City = extractString(useCityRe, contents)
	profile.Name = name

	return engine.ParseResult{
		Items: []engine.Item{
			{
				Url:  url,
				Id:   extractString(idRe, []byte(url)),
				Type: "zhenai",
				Data: profile,
			},
		},
	}
}

func convertStringToInt(num string) int {
	if num == "" {
		return 0
	}

	i, err := strconv.Atoi(num)
	if err != nil {
		log.Panicf("num is %s, %v\n",
			num, err)
		return 0
	}

	return i
}

func extractString(re *regexp.Regexp, contents []byte) string {
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		return string(m[1])
	}
	return ""
}

func ProfileParser(name string) engine.ParseFunc {
	return func(contents []byte, url string) engine.ParseResult {
		return ParseProfile(contents, url, name)
	}
}
