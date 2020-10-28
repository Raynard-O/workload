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

/*Data mining and Algorithm Manipulation of Dataset request
Gets -  the Dataset and Client request
Return - slice of data as requested by the client
*/

func Client2(c echo.Context) error {
	// gets dataset and request params from the context
	data, workload := DataSet(c)
	log.Println(data)
	if data == nil {
		return BadRequestResponse(c, "Invalid Params")
	}

	//determine the workload metrics to be operated on
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


	// derive the last batch ID of the dataset with response to the batch unit size
	last_batch_id := workload.BSize/ workload.BatchUnit
	last_batch_id = last_batch_id + 1
	fmt.Println(Bench, last_batch_id)

	// new instance for arrays of  batch request
	var batch3 []*grpc_from0.Batch


	// iterate over the dataset array to create a new array for batch IDS
	ID := workload.BatchID-1
	BatchSize := ID+workload.BatchSize -1

	for i := ID; i <= BatchSize; {

		b := i*workload.BatchUnit
		z := b+workload.BatchUnit
		// creating a new batch per ID and appending the new batch to the array of batches
		newBatch := data[b+1:z+1]
		// new instance for arrays of  sample request
		var sam []*grpc_from0.Sample
		for k := 0 ; k <= workload.BatchUnit-1; k++ {
			// creating a new sample for each batch unit sample and appending the new sample to the array of samples

			if Bench == 0 {
				samp := &grpc_from0.Sample{
					CPUUtilization: newBatch[k].CpuUtilizationAverage,
				}
				sam = append(sam, samp)
			}else if Bench == 1 {
				samp := &grpc_from0.Sample{
					NetworkIN: newBatch[k].NetworkInAverage,
				}
				sam = append(sam, samp)
			}else if Bench == 2 {
				samp := &grpc_from0.Sample{
					NetworkOUT: newBatch[k].NetworkOutAverage,
				}
				sam = append(sam, samp)
			}else if Bench == 3 {
				samp := &grpc_from0.Sample{
					MemoryUtilization: newBatch[k].MemoryUtilizationAverage,
				}
				sam = append(sam, samp)
			}else {
				samp := &grpc_from0.Sample{
					MemoryUtilization: newBatch[k].MemoryUtilizationAverage,
				}
				sam = append(sam, samp)
			}


		}
		batch := &grpc_from0.Batch{
			Batch_ID: int32(i+1),
			Samples:  sam,
		}
		log.Println(newBatch)
		batch3 = append(batch3, batch)
		i++
	}
	// determining the type of files to return to the client depending on the Binary serialization request.
	if workload.BinarySerialization == "binary" {
		//returns protobuf binary
		return Proto(c, workload.RFWID, int32(last_batch_id), batch3)
	}
	//return json encoded data
	return EncodeJson(c, workload.RFWID, int32(last_batch_id), batch3)
}
