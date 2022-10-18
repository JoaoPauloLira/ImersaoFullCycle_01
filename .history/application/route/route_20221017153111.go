package route

type Route struct {
	ID        string `json:"routeId"`
	ClientID  string `json:"clientId"`
	Positions []Position `json:"position"`
}

type Position struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}