package controller

import (
	"Proto/library"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)



func J(c echo.Context) error {

	data, workload := DataSet(c)
	//size := strconv.Itoa(int(workload.Batch_Size))

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

	batchSize := (workload.BatchSize / workload.BatchUnit) - 1
	var batch3 []*library.Batch

	for i := batchSize; i > (batchSize)-workload.BatchID; i-- {
		//fmt.Println(i)
		f := batchSize - i

		var sam []*library.Sample

		c := i * workload.BatchUnit

		for j := c; j < c+workload.BatchUnit; j++ {

			switch Bench {
			case 0:
				samp := &library.Sample{
					CpuUtilization: data[j].CpuUtilizationAverage,
				}
				sam = append(sam, samp)
			case 1 :
				samp := &library.Sample{
					NetworkIN: data[j].NetworkInAverage,
				}
				sam = append(sam, samp)
			case 2:
				samp := &library.Sample{
					NetworkOUT: data[j].NetworkOutAverage,
				}
				sam = append(sam, samp)
			case 3:
				samp := &library.Sample{
					MemoryUtilization: data[j].MemoryUtilizationAverage,
				}
				sam = append(sam, samp)
			default:
				samp := &library.Sample{
					FinalTarget: data[j].FinalTarget,
				}
				sam = append(sam, samp)
			}

		}

		batch := &library.Batch{
			Batch_ID: f + 1,
			Samples:  sam,
		}
		batch3 = append(batch3, batch)
	}
	return  EncodeJson(c, workload.RFWID, workload.BatchID, batch3)

}

func EncodeJson(c echo.Context, RfwId string, LastBatchID int, samplesWorK []*library.Batch) error {
	RFW := &library.RFD{
		RFWID:       RfwId,
		LastBatchId: LastBatchID,
		Batches:     samplesWorK,
	}

	return c.JSONPretty(http.StatusOK, RFW, "")
}

