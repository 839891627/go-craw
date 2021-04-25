package parser

import (
	"crawler/mode"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "洒脱遇到")

	if len(result.Items) != 1 {
		t.Errorf("Result should contain 1 element; but was %v", result.Items)
	}

	profile := result.Items[0].(mode.Profile)

	expected := mode.Profile{
		Name:       "洒脱遇到",
		Gender:     "女",
		Age:        71,
		Height:     22,
		Weight:     292,
		Income:     "2001-3000",
		Marriage:   "未婚",
		Education:  "高中",
		Occupation: "",
		Hokou:      "南京市",
		Xinzuo:     "白羊座",
		House:      "租房",
		Car:        "有车",
	}

	if profile != expected {
		t.Errorf("expected %v; but was %v", expected, profile)
	}
}
