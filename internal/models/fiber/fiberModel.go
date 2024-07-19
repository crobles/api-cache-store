package fiberModel

type Config struct {
	Prefork	bool  `json:"prefork"`
	CaseSensitive	bool  `json:"case_sensitive"`
	StrictRouting	bool  `json:"strict_routing"`
	ServerHeader	string  `json:"server_header"`
	AppName	string  `json:"app_name"`
}