package parser

import (
	"../../engine"
	"../../model"
	"regexp"

	"strconv"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([0-9]+)岁</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span><span field="">([0-9]+)CM</span></td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([0-9]+)KG</span></td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var xinzuoRe = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
var hukouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var carRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)

func ParseProfile(contents []byte) engine.ParseResult {
	profile := model.Profile{}
	age, err := strconv.Atoi(extractString(contents, ageRe))

	if err == nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}

	profile.Income = extractString(contents, incomeRe)
	profile.Gender = extractString(contents, genderRe)
	profile.Car = extractString(contents, carRe)
	profile.Education = extractString(contents, educationRe)
	profile.Hukou = extractString(contents, hukouRe)
	profile.House = extractString(contents, houseRe)
	profile.Marriage = extractString(contents, marriageRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.Xinzuo = extractString(contents, xinzuoRe)
	//profile.Name  = extractString(contents, xinzuoRe)*/

	return engine.ParseResult{
		Items: []interface{}{profile},
	}

}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
