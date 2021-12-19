package PrometheusStruct

type GetsInfoRequest struct {

	PrometheusId string `form:"prometheusId" json:"prometheusId"`
	UserId string `form:"userId" json:"userId"`
}
