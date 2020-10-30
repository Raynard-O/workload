package library

type DataParamsRequest struct {
	RFWID          string `json:"rfwid"`
	BenchmarkType  string `json:"benchmark_type"`
	WorkloadMetric string `json:"workload_metric"`
	BatchUnit      int    `json:"batch_unit"`
	BatchID        int    `json:"batch_id"`
	BatchSize      int    `json:"batch_size"`
	BSize 	int `json:"b_size"`
	BinarySerialization	string `json:"binary_serialization"`
}
