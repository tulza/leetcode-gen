package gencmd

import (
	"encoding/json"
	"io"
	"lcgen/internal/config"
	"lcgen/model"
	"net/http"
)

func fetchProblem(slugName string) (model.Problem, error) {

	// fetch leetcode data from api
	API_URL := config.API_BASE_URL + `/select?titleSlug=` + slugName
	res, err := http.Get(API_URL)

	if err != nil {
		return model.Problem{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return model.Problem{}, err
	}

	// parse data
	var problem model.Problem
	err = json.Unmarshal(body, &problem)

	if err != nil {
		return model.Problem{}, err
	}

	return problem, nil
}

func fetchDaily() (model.Problem, error) {

	// fetch leetcode data from api
	API_URL := config.API_BASE_URL + `/daily`
	res, err := http.Get(API_URL)

	if err != nil {
		return model.Problem{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return model.Problem{}, err
	}

	// parse data
	var problem model.Problem
	err = json.Unmarshal(body, &problem)

	if err != nil {
		return model.Problem{}, err
	}

	return problem, nil
}