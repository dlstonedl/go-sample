package parser

import (
	"fmt"
	"github.com/dlstonedl/go-sample/crawler/engine"
	"github.com/dlstonedl/go-sample/crawler/model"
	"regexp"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([^<]+)岁</div>`)
var marriageRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>(未婚|离异)</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([^<]+)cm</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([^<]+)kg</div>`)
var genderRe = regexp.MustCompile(`"genderString":"(男|女)士"`)
var educationRe = regexp.MustCompile(`"educationString":"([^"]+)"`)
var incomeRe = regexp.MustCompile(`<div class="des f-cl"[^>]*>([^<]+)</div>`)

func ParseProfile(contents []byte) engine.ParseResult {
	result := engine.ParseResult{}

	profile := model.Profile{}

	fmt.Printf("%s\n", extractString(ageRe, contents))
	fmt.Printf("%s\n", extractString(marriageRe, contents))
	fmt.Printf("%s\n", extractString(heightRe, contents))
	fmt.Printf("%s\n", extractString(weightRe, contents))
	fmt.Printf("%s\n", extractString(genderRe, contents))
	fmt.Printf("%s\n", extractString(educationRe, contents))
	fmt.Printf("%s\n", extractString(incomeRe, contents))

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
