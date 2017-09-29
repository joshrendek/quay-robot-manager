package robots

import (
	"encoding/json"
	"errors"

	"fmt"

	"github.com/joshrendek/quay-robot-manager/requests"
	"github.com/rs/zerolog/log"
)

var BearerToken string

const (
	listRobotsURL = "https://quay.io/api/v1/organization/%s/robots"
	crudRobotURL  = "https://quay.io/api/v1/organization/%s/robots/%s"
)

type Robot struct {
	Token string `json:"token"`
	Name  string `json:"name"`
}

type ListRobots struct {
	Robots []Robot `json:"robots"`
}

func All(org string) ([]Robot, error) {
	url := fmt.Sprintf(listRobotsURL, org)
	resp, statusCode := requests.Request("GET", url, BearerToken, nil)
	if statusCode != 200 {
		log.Debug().Msgf("Response: %s", string(resp))
		return nil, errors.New(string(resp))
	}
	robots := ListRobots{}
	err := json.Unmarshal(resp, &robots)
	if err != nil {
		log.Fatal().Msgf("Error parsing json: %s", string(resp))
	}
	return robots.Robots, nil
}

func Get(name, org string) (*Robot, error) {
	url := fmt.Sprintf(crudRobotURL, org, name)
	resp, statusCode := requests.Request("GET", url, BearerToken, nil)
	if statusCode != 200 {
		log.Debug().Msgf("Response: %s", string(resp))
		return nil, errors.New(string(resp))
	}
	robot := Robot{}
	json.Unmarshal(resp, &robot)
	return &robot, nil
}

func Create(name, org string) (*Robot, error) {
	url := fmt.Sprintf(crudRobotURL, org, name)
	resp, statusCode := requests.Request("PUT", url, BearerToken, nil)
	if statusCode != 201 {
		log.Debug().Msgf("Response: %s", string(resp))
		return nil, errors.New(string(resp))
	}
	robot := Robot{}
	json.Unmarshal(resp, &robot)
	return &robot, nil
}

func Delete(name, org string) error {
	url := fmt.Sprintf(crudRobotURL, org, name)
	resp, statusCode := requests.Request("DELETE", url, BearerToken, nil)
	if statusCode != 204 {
		log.Debug().Msgf("Response: %s", string(resp))
		return errors.New(string(resp))
	}
	return nil
}
