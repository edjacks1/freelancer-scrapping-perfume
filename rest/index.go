package rest

type Rest struct {
	baseUrl string
}

func InitRest(url string) Rest {
	return Rest{baseUrl: url}
}
