package topdeliveryRestClient

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type JurFace struct {
	Title                string
	JurAddress           string
	FactAddress          string
	DirectorFio          string
	DirectorFio2         string
	BookkeeperFio        string
	BookkeeperFio2       string
	SignerNameGenitive   string
	SignerPosition       string
	Authority            string
	OrganizationType     string
	Inn                  string
	Kpp                  string
	Ogrn                 string
	Bank                 string
	Account              string
	CorrespondentAccount string
	Bik                  string
	Nds                  string
	Okpo                 string
	Id                   int
}



func (c *Client) PostJurFace(jurFace JurFace) (*JurFace, error) {

	jurFaceB,err := json.Marshal(jurFace)

	if err != nil{
		return nil,err
	}

	req, err := http.NewRequest("POST", c.config.BaseUrl+ "/JurFace", bytes.NewBuffer(jurFaceB))
	if err != nil {
		return nil, errors.New("error creating PostJurFace request: " + err.Error())
	}

	var items JurFace

	if err := c.callApi(req,&items); err != nil {
		return nil, err
	}

	return &items, nil
}
