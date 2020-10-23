package library

type DataParamsRequest struct {
	RFWID          string `json:"rfwid"`
	BenchmarkType  string `json:"benchmark_type"`
	WorkloadMetric string `json:"workload_metric"`
	BatchUnit      int    `json:"batch_unit"`
	BatchID        int    `json:"batch_id"`
	BatchSize      int    `json:"batch_size"`
	BinarySerialization	string `json:"binary_serialization"`
}

type Sample struct {
	CpuUtilization    int32 `json:"cpu_utilization"`
	NetworkIN         int32 `json:"network_in"`
	NetworkOUT        int32 `json:"network_out"`
	MemoryUtilization float32 `json:"memory_utilization"`
	FinalTarget		float32 `json:"final_target"`
}
type Batch struct {
	Batch_ID int      `json:"batch_id"`
	Samples  []*Sample `json:"samples"`
}
type RFD struct {
	RFWID       string  `json:"rfwid"`
	LastBatchId int     `json:"last_batch_id"`
	Batches     []*Batch `json:"batches"`
}

