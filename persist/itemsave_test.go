package persist

import (
	"context"
	"encoding/json"
	"testing"

	"../model"
	"gopkg.in/olivere/elastic.v5"
)

func TestSave(t *testing.T) {
	expected := model.Profile{
		Name:       "西洋参",
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
	id, err := save(expected)
	if err != nil {
		panic(err)
	}

	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	resp, err := client.Get().
		Index("dating_profile").
		Type("zhenai").
		Id(id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	// TODO : Try to start up elastic search
	// here using docker go client
	t.Logf("%s", resp.Source)
	var actual model.Profile
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil {
		panic(err)
	}
	if actual != expected {
		t.Errorf("got %v expect %v", actual, expected)
	}

}
