package controller

import (
	"Proto/library"
	"encoding/csv"
	"fmt"
	"github.com/labstack/echo"
	"io"
	"log"
	"os"
	"strconv"
)

type WORKLOAD struct {
	CpuUtilizationAverage    int32
	NetworkInAverage         int32
	NetworkOutAverage        int32
	MemoryUtilizationAverage float32
	FinalTarget              float32
}
type DATAs struct {
	WORKLOADs []WORKLOAD
}

func (DATA *DATAs) AddItems(data WORKLOAD) {
	DATA.WORKLOADs = append(DATA.WORKLOADs, data)
}

func DataSet(ctx echo.Context) ([]WORKLOAD, *library.DataParamsRequest) {

	workload := new(library.DataParamsRequest)
	if err := ctx.Bind(workload); err != nil {
		log.Fatal(err)
	}
	fmt.Print(workload)
	if workload.BenchmarkType == "" || workload.RFWID == "" || workload.WorkloadMetric == "" {
		log.Fatal("Params can not be empty")
	}

		workload.BatchSize = 0
	//workload := library.DataParamsRequest{
	//
	//	Benchmark_Type:  "NDBENCH",
	//	Workload_Metric: "CPU",
	//	Batch_Unit:      2,
	//	BatchId:         5,
	//	Batch_Size:      0,
	//}
	var fileimage string

	switch workload.BenchmarkType {
	case "DVD":
		fileimage = "Workload_Data/DVD-training.csv"
		//if workload.Workload_Metric == "TRAIN" {
		//	fileimage = "Workload_Data/DVD-training.csv"
		//}
		//else {
		//	fileimage = "Workload_Data/DVD-testing.csv"
		//}
	default:
		fileimage = "Workload_Data/NDBench-training.csv"
		//if workload.Workload_Metric != "TRAIN" {
		//	fileimage = "Workload_Data/NDBench-training.csv"
		//}else {
		//	fileimage = "Workload_Data/NDBench-testing.csv"
		//}
	}

	file, err := os.Open(fileimage)
	if err != nil {
		log.Fatal(err.Error())
	}
	r := csv.NewReader(file)
	// Iterate through the records
	datas := new(DATAs)

	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		i, _ := strconv.Atoi(record[0])
		j, _ := strconv.Atoi(record[1])
		k, _ := strconv.Atoi(record[2])
		s, _ := strconv.ParseFloat(record[3], 64)
		m, _ := strconv.ParseFloat(record[4], 64)

		dvd := new(WORKLOAD)

		dvd.CpuUtilizationAverage = int32(i)
		dvd.NetworkInAverage = j
		dvd.NetworkOutAverage = int32(k)
		dvd.MemoryUtilizationAverage = float32(s)
		dvd.FinalTarget = m
		workload.BatchSize++
		datas.AddItems(*dvd)
	}

	return datas.WORKLOADs, workload
	//ctx.JSON(200, datas.DVD_TESTs[dataNumber-1])
}
