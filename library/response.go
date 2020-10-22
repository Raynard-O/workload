package library

type DataParamsRequest struct {
	RFWID          string `json:"rfwid"`
	BenchmarkType  string `json:"benchmark_type"`
	WorkloadMetric string `json:"workload_metric"`
	BatchUnit      int    `json:"batch_unit"`
	BatchID        int    `json:"batch_id"`
	BatchSize      int    `json:"batch_size"`
}

type sample struct {
	CPU_utilization    int `json:"cpu_utilization"`
	NetworkIN          int `json:"network_in"`
	NetworkOUT         int `json:"network_out"`
	Memory_utilization int `json:"memory_utilization"`
}
type batch struct {
	Batch_ID int      `json:"batch_id"`
	samples  []sample `json:"samples"`
}
type RFD struct {
	RFWID         string  `json:"rfwid"`
	last_batch_id int     `json:"last_batch_id"`
	batches       []batch `json:"batches"`
}

//
//func GenId(pref string, n int) string {
//	RHex, _ := RandomHex(n)
//	return fmt.Sprintf("%s_%s", pref, RHex)
//}
//
//// Generate random hash
//func RandomHex(n int) (string, error) {
//	bytes := make([]byte, n)
//	if _, err := rand.Read(bytes); err != nil {
//		return "", err
//	}
//	return hex.EncodeToString(bytes), nil
//}
//
