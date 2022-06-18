package helper

type WebResposne struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
