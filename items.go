package topdeliveryRestClient

import (
	"errors"
	"net/http"
)

type Item struct {
	Article         string
	ClientPrice     int
	Count           int
	DeliveryCount   int
	Id              int
	Name            string
	OrderId         int
	Price           int
	ProblemCount    int
	ProblemStatusId int
	StatusId        int
	Type            string
	Vat             int
	Weight          int
}

func (c *Client) GetItems(searchString string) ([]Item, error) {

	req, err := http.NewRequest("GET", c.config.BaseUrl+"/Order/Item"+searchString, nil)

	if err != nil {
		return nil, errors.New("error creating GetItems request: " + err.Error())
	}

	var items []Item

	if err := c.callApi(req, &items); err != nil {
		return nil, err
	}

	return items, nil
}
