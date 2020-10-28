package controller

import (
	grpcfrom0 "Proto/github.com/monkrus/grpc-from0"
	"github.com/labstack/echo"
)

func Options(c echo.Context) error {



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

	//fmt.Printf("Workload Metrics : %d", Bench)

	var batch3 []*grpcfrom0.Batch

	ID := 1
	for i := 0; i <= totalT; {

		z := total-i
		b := total-i-size
		test := result[b: z]
		var sam []*grpcfrom0.Sample
		for k := size-1 ; k >= 0; k-- {
			if Bench == 0 {
				samp := &grpcfrom0.Sample{
					CPUUtilization: test[k].CpuUtilizationAverage,
				}
				sam = append(sam, samp)
			}else if Bench == 1 {
				samp := &grpcfrom0.Sample{
					NetworkIN: test[k].NetworkInAverage,
				}
				sam = append(sam, samp)
			}else if Bench == 2 {
				samp := &grpcfrom0.Sample{
					NetworkOUT: test[k].NetworkOutAverage,
				}
				sam = append(sam, samp)
			}else if Bench == 3 {
				samp := &grpcfrom0.Sample{
					MemoryUtilization: test[k].MemoryUtilizationAverage,
				}
				sam = append(sam, samp)
			}else {
				samp := &grpcfrom0.Sample{
					MemoryUtilization: test[k].MemoryUtilizationAverage,
				}
				sam = append(sam, samp)
			}


		}
		batch := &grpcfrom0.Batch{
			Batch_ID: int32(ID),
			Samples:  sam,
		}
		//log.Println(test)
		batch3 = append(batch3, batch)
		i += size
		ID++
	}




	if workload.BinarySerialization == "binary" {
		return Proto(c, workload.RFWID, int32(workload.BatchID), batch3)
	}

	return EncodeJson(c, workload.RFWID, int32(workload.BatchID), batch3)

}

