package response

type MonitStatResponse struct {
	HostName  string  `json:"hostName"`
	IP        string  `json:"ip"`
	CPU       string  `json:"CPU"`
	RamTotal  uint64  `json:"ramTotal"`
	RamUse    uint64  `json:"ramUse"`
	Ram       float64 `json:"ram"`
	DiskTotal uint64  `json:"DiskTotal"`
	DiskUse   uint64  `json:"DiskUse"`
	Disk      float64 `json:"disk"`
	Flow      uint64  `json:"flow"`
}
