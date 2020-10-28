package controller

import (
	grpc_from0 "Proto/github.com/monkrus/grpc-from0"
	"fmt"
	"github.com/labstack/echo"
	"log"
)

func Summary(c echo.Context) error {

	data, workload := DataSet(c)

	var size int = workload.BatchUnit
	var unit int = workload.BatchID
	bSize := workload.BatchSize
	total := size * unit
	totalT := total - 1
	result := data[bSize-total : bSize]

	var Bench int8
	switch workload.WorkloadMetric {
	case "CPU":
		Bench = 0
		//serverRFW.Benchmark_Type = "CPUUtilization_Average"
	case "NETIN":
		Bench = 1
		//serverRFW.Benchmark_Type = "NetworkIn_Average"
	case "NETOUT":
		Bench = 2
		//serverRFW.Benchmark_Type = "NetworkOut_Average"
	case "MEMUTI":
		Bench = 3
		//serverRFW.Benchmark_Type = "MemoryUtilization_Average"
	default:
		Bench = 4
		//serverRFW.Benchmark_Type = "Final_Target"
	}

	fmt.Println(Bench)

	var batch3 []*grpc_from0.Batch

	ID := 1
	for i := 0; i <= totalT; {

		z := total-i
		b := total-i-size
		test := result[b: z]
		var sam []*grpc_from0.Sample
		for k := 0 ; k <= size-1; k++ {
			if Bench == 0 {
				samp := &grpc_from0.Sample{
					CPUUtilization: test[k].CpuUtilizationAverage,
				}
				sam = append(sam, samp)
			}else if Bench == 1 {
				samp := &grpc_from0.Sample{
					NetworkIN: test[k].NetworkInAverage,
				}
				sam = append(sam, samp)
			}else if Bench == 2 {
				samp := &grpc_from0.Sample{
					NetworkOUT: test[k].NetworkOutAverage,
				}
				sam = append(sam, samp)
			}else if Bench == 3 {
				samp := &grpc_from0.Sample{
					MemoryUtilization: test[k].MemoryUtilizationAverage,
				}
				sam = append(sam, samp)
			}else {
				samp := &grpc_from0.Sample{
					MemoryUtilization: test[k].MemoryUtilizationAverage,
				}
				sam = append(sam, samp)
			}


		}
		batch := &grpc_from0.Batch{
			Batch_ID: int32(ID),
			Samples:  sam,
		}
		log.Println(test)
		batch3 = append(batch3, batch)
		i += size
		ID++
	}

	return c.JSONPretty(200, batch3, "")
}



func Client2(c echo.Context) error {

	data, workload := DataSet(c)
	var Bench int8
	switch workload.WorkloadMetric {
	case "CPU":
		Bench = 0
		//serverRFW.Benchmark_Type = "CPUUtilization_Average"
	case "NETIN":
		Bench = 1
		//serverRFW.Benchmark_Type = "NetworkIn_Average"
	case "NETOUT":
		Bench = 2
		//serverRFW.Benchmark_Type = "NetworkOut_Average"
	case "MEMUTI":
		Bench = 3
		//serverRFW.Benchmark_Type = "MemoryUtilization_Average"
	default:
		Bench = 4
		//serverRFW.Benchmark_Type = "Final_Target"
	}

	last_batch_id := workload.BSize/ workload.BatchUnit
	last_batch_id = last_batch_id + 1
	fmt.Println(Bench, last_batch_id)
	ID := workload.BatchID-1
	BatchSize := ID+workload.BatchSize -1

	var batch3 []*grpc_from0.Batch

	for i := ID; i <= BatchSize; {

		b := i*workload.BatchUnit
		z := b+workload.BatchUnit
		test := data[b+1:z+1]
		var sam []*grpc_from0.Sample
		for k := 0 ; k <= workload.BatchUnit-1; k++ {
			if Bench == 0 {
				samp := &grpc_from0.Sample{
					CPUUtilization: test[k].CpuUtilizationAverage,
				}
				sam = append(sam, samp)
			}else if Bench == 1 {
				samp := &grpc_from0.Sample{
					NetworkIN: test[k].NetworkInAverage,
				}
				sam = append(sam, samp)
			}else if Bench == 2 {
				samp := &grpc_from0.Sample{
					NetworkOUT: test[k].NetworkOutAverage,
				}
				sam = append(sam, samp)
			}else if Bench == 3 {
				samp := &grpc_from0.Sample{
					MemoryUtilization: test[k].MemoryUtilizationAverage,
				}
				sam = append(sam, samp)
			}else {
				samp := &grpc_from0.Sample{
					MemoryUtilization: test[k].MemoryUtilizationAverage,
				}
				sam = append(sam, samp)
			}


		}
		batch := &grpc_from0.Batch{
			Batch_ID: int32(i+1),
			Samples:  sam,
		}
		log.Println(test)
		batch3 = append(batch3, batch)
		//i += size
		i++
	}

	if workload.BinarySerialization == "binary" {
		return Proto(c, workload.RFWID, int32(last_batch_id), batch3)
	}

	return EncodeJson(c, workload.RFWID, int32(last_batch_id), batch3)


}
