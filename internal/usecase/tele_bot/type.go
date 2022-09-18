package tele_bot

type EarthquakeInfoResp struct {
	InfoGempa EarthquakeInfo `json:"Infogempa"`
}

type EarthquakeInfo struct {
	Gempa Earthquake `json:"gempa"`
}

type EarthquakeInfoListResp struct {
	InfoGempa EarthquakeInfoList `json:"Infogempa"`
}

type EarthquakeInfoList struct {
	Gempa []Earthquake `json:"gempa"`
}

type Earthquake struct {
	Tanggal   string `json:"Tanggal"`
	Jam       string `json:"Jam"`
	Kordinat  string `json:"Coordinates"`
	Lintang   string `json:"Lintang"`
	Bujur     string `json:"Bujur"`
	Magnitude string `json:"Magnitude"`
	Kedalaman string `json:"Kedalaman"`
	Wilayah   string `json:"Wilayah"`
	Potensi   string `json:"Potensi"`
	Dirasakan string `json:"Dirasakan,omitempty"`
}
