package response

type JrebelResponse struct {
	ServerVersion         string   `json:"serverVersion"`         // \"serverVersion\": \"3.2.4\",\n" +
	ServerProtocolVersion string   `json:"serverProtocolVersion"` // "    \"serverProtocolVersion\": \"1.1\",\n" +
	ServerGuid            string   `json:"serverGuid"`            // "    \"serverGuid\": \"a1b4aea8-b031-4302-b602-670a990272cb\",\n" +
	GroupType             string   `json:"groupType"`             // "    \"groupType\": \"managed\",\n" +
	Id                    int      `json:"id"`                    // "    \"id\": 1,\n" +
	LicenseType           int      `json:"licenseType"`           // "    \"licenseType\": 1,\n" +
	EvaluationLicense     bool     `json:"evaluationLicense"`     // "    \"evaluationLicense\": false,\n" +
	Signature             string   `json:"signature"`             // "    \"signature\": \"OJE9wGg2xncSb+VgnYT+9HGCFaLOk28tneMFhCbpVMKoC/Iq4LuaDKPirBjG4o394/UjCDGgTBpIrzcXNPdVxVr8PnQzpy7ZSToGO8wv/KIWZT9/ba7bDbA8/RZ4B37YkCeXhjaixpmoyz/CIZMnei4q7oWR7DYUOlOcEWDQhiY=\",\n" +
	ServerRandomness      string   `json:"serverRandomness"`      // "    \"serverRandomness\": \"H2ulzLlh7E0=\",\n" +
	SeatPoolType          string   `json:"seatPoolType"`          // "    \"seatPoolType\": \"standalone\",\n" +
	StatusCode            string   `json:"statusCode"`            // "    \"statusCode\": \"SUCCESS\",\n" +
	Offline               string   `json:"offline"`               // "    \"offline\": " + String.valueOf(offline) + ",\n" +
	ValidFrom             string   `json:"validFrom"`             // "    \"validFrom\": " + validFrom + ",\n" +
	ValidUntil            string   `json:"validUntil"`            // "    \"validUntil\": " + validUntil + ",\n" +
	Company               string   `json:"company"`               // "    \"company\": \"Administrator\",\n" +
	OrderId               string   `json:"orderId"`               // "    \"orderId\": \"\",\n" +
	ZeroIds               []string `json:"zeroIds"`               // "    \"zeroIds\": [\n" +  ],\n" +
	LicenseValidFrom      string   `json:"licenseValidFrom"`      // "    \"licenseValidFrom\": 1490544001000,\n" +
	LicenseValidUntil     string   `json:"licenseValidUntil"`     // "    \"licenseValidUntil\": 1691839999000\n" +
}

type JrebelResponse2 struct {
	ServerVersion         string `json:"serverVersion"`
	ServerProtocolVersion string `json:"serverProtocolVersion"`
	ServerGuid            string `json:"serverGuid"`
	GroupType             string `json:"groupType"`
	StatusCode            string `json:"statusCode"`
	Msg                   string `json:"msg"`
	StatusMessage         string `json:"statusMessage"`
	Username              string `json:"username"`
}

