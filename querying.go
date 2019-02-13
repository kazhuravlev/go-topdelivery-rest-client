package topdeliveryRestClient

import (
	"encoding/json"
	"errors"
)

type (
	Rule struct {
		Field string      `json:"field"`
		Op    string      `json:"op"`
		Data  interface{} `json:"data"`
	}

	Group struct {
		GroupOp string  `json:"groupOp"`
		Rules   []Rule  `json:"rules"`
		Groups  []Group `json:"groups"`
	}

	Sort struct {
		Sidx string `json:"sidx"`
		Sord string `json:"sord"`
	}

	Limit struct {
		Rows int `json:"rows"`
		Page int `json:"page"`
	}

	query struct {
		Search Group
		Sort   Sort
		Limit  Limit
	}
)

func NewQuery() *query {
	return &query{
		Search: Group{
			GroupOp: "AND",
			Rules:   make([]Rule, 0),
			Groups:  make([]Group, 0),
		},
	}
}

func (q *query) ToString() (queryString string, err error) {
	queryString = "?"

	if len(q.Search.Groups) > 0 || len(q.Search.Rules) > 0 {
		queryBytes, err := json.Marshal(q.Search)
		if err != nil {
			return "", errors.New("error marshaling query Group" + err.Error())
		}
		queryString += "filters=" + string(queryBytes) + "&"
	}

	if q.Sort.Sidx != "" && q.Sort.Sord != "" {
		queryBytes, err := json.Marshal(q.Search)
		if err != nil {
			return "", errors.New("error marshaling query Sort" + err.Error())
		}
		queryString += string(queryBytes) + "&"
	}

	if q.Limit.Rows > 0 || q.Limit.Page > 0 {
		queryBytes, err := json.Marshal(q.Limit)
		if err != nil {
			return "", errors.New("error marshaling query Limit" + err.Error())
		}
		queryString += string(queryBytes) + "&"
	}

	return queryString, nil
}
