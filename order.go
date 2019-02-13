package topdeliveryRestClient

import (
	"errors"
	"net/http"
	"time"
)

type (
	Order struct {
		Additional            int
		AdditionalOrderId     int
		BarCode               string
		CheckByCc             int
		CityCurrentId         int
		CityDeliveryId        int
		CityWebshopId         int
		ClientEmail           string
		ClientFio             string
		ClientFullCost        float64
		ClientPaid            int
		ClientPhone           string
		Comment               string
		CourierFio            string
		CourierPhone          string
		CurrentCost           int
		CurrentOwnerType      string
		CurrentPartnerId      int
		CurrentShipmentId     int
		DateCreate            time.Time
		DateEndOfStorage      time.Time
		DateFactDelivery      time.Time
		DateFinalStatus       time.Time
		DateLastStatus        time.Time
		DateProblem           time.Time
		DateProblemStatus     time.Time
		DateReceiveInTd       time.Time
		DateReceiveWebshop    time.Time
		DeliveryCost          float64
		DeliveryItemsCost     float64
		DeliveryStreet        string
		DeliveryType          string
		DeliveryZipcode       int
		DenyReasonId          int
		DenyType              string
		ExtPartnerInfo        string
		ForChoice             int
		Id                    int
		InReportStatus        int
		Locked                int
		NeedMarking           int
		NeedSmsNotify         int
		NotOpen               int
		OrderSubtype          string
		OrderUrl              string
		PackDone              int
		PackSize              int
		PackType              int
		Partner1TariffScaleId int
		Partner2TariffScaleId int
		PartnerExecutorId     int
		PartnerWebshopId      int
		PayedGroupId          int
		PaymentType           string
		PickupAddressId       int
		Problem               int
		ProblemStatusId       int
		ReceiverClientFio     string
		RegionCurrentId       int
		RegionDeliveryId      int
		RegionWebshopId       int
		Reimbursable          bool
		ReturnItemsCost       float64
		ServiceType           string
		ShiftCountCall        int
		ShiftCountPlace       int
		ShiftCountUncall      int
		StatusId              int
		VolumeWeightDelivery  int
		WebshopBarcode        string
		WebshopId             int
		WebshopNumber         string
		WebshopTariffScaleId  int
		WeightInfoDelivery    weight
		WeightInfoReturn      weight
		WorkStatusId          int
		//wsDesireDateDelivery: "2019-11-12 11:11:11"
		//wsDesireDeliveryInterval: "11:11:11/11:11:11"
	}

	weight struct {
		physical int
		volume   [3]int
	}
)

func (c *Client) GetOrders(searchString string) ([]Order, error) {

	req, err := http.NewRequest("GET", c.config.BaseUrl+ "/Order/Order" +searchString, nil)
	if err != nil {
		return nil, errors.New("error creating PostJurFace request: " + err.Error())
	}

	var orders []Order

	if err := c.callApi(req,&orders); err != nil {
		return nil, err
	}

	return orders, nil
}
