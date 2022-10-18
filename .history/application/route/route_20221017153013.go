package route

type Route struct {
	ID        string `json:"routeId"`
	ClientID  string `json:"clientId"`
	Positions []Position `json:"position"`
}