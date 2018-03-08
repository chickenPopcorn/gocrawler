package parser

import (
	"../../model"
	"../parser"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile(
		"profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := parser.ParseProfile(contents)
	if len(result.Items) != 1 {
		t.Errorf("Result should contain 1 element; but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)

	expected := model.Profile{
		Name:       "",
		Gender:     "男",
		Age:        53,
		Height:     170,
		Weight:     78,
		Income:     "3001-5000元",
		Marriage:   "未婚",
		Education:  "大专",
		Occupation: "其他职业",
		Hukou:      "北京东城区",
		Xinzuo:     "牡羊座",
		House:      "已购房",
		Car:        "未购车",
	}

	if profile != expected {
		t.Errorf("epxected %v; but was %v", expected, profile)
	}
}
