package doTest


type ReqBody struct {
	Token 	string	`json:"token"`
	Action 	string	`json:"action"`
	Timestamp 	int	`json:"timestamp"`
	IPInfo 	IP_Info	`json:"ip_info"`
	PingData Ping_Data	`json:"ping_data"`
}

type IP_Info struct {
	City 		string	`json:"city"`
	Country 	string	`json:"country"`
	IP 			string		`json:"ip"`
	Location 	string	`json:"location"`
	Org 		string		`json:"org"`
	Region 		string		`json:"region"`
}

type Ping_Data struct {
	TTL   string `json:"TTL"`
	Delay int `json:"delay"`
	DstIP string `json:"dst_ip"`
	Loss  int    `json:"loss"`
	SrcIP string    `json:"src_ip"`
}


var TestBody = &ReqBody{
	Token: 	"",
	Action: 	"ping",
	Timestamp: 	1545733726,
	IPInfo: IP_Info{
		City: "shanghai_yangpu",
		Country: "china",
		IP: 	"192.168.152.12",
		Location: "string",
		Org: 	"string",
		Region: "shanghai",
	},
	PingData: Ping_Data{
		TTL: "3",
		Delay: 1,
		DstIP: "192.168.153.218",
		Loss: 0,
		SrcIP: "192.168.152.12",
	},
}

var pingData = &Ping_Data{
	TTL: "3",
	Delay: 1,
	DstIP: "192.168.153.218",
	Loss: 0,
	SrcIP: "192.168.152.12",
}

var ipInfo = &IP_Info{
	City: "shanghai_yangpu_longchanglu",
	Country: "china",
	IP: 	"192.168.152.12",
	Location: "string",
	Org: 	"string",
	Region: "shanghai",
}