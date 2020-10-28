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



/*Data set structure*/
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


/*
Retrieve Datasets base on request params from the front end

returns the full data to be processed, dataset size Bsize

*/
func DataSet(ctx echo.Context) ([]WORKLOAD, *library.DataParamsRequest) {
// create new data structure
	workload := new(library.DataParamsRequest)
	// bind/marshal query request to the data structure
	if err := ctx.Bind(workload); err != nil {
		log.Fatal(err)
	}
	fmt.Print(workload)
	//check for valid query params
	if workload.BatchUnit == 0 || workload.BatchID == 0 || workload.RFWID == "" || workload.BatchSize == 0 {
		return nil, nil
	}
	//count data files
	workload.BSize = 0

	var fileimage string
	//determine dataset to be worked on depending on query params
	switch workload.BenchmarkType {
	case "DVD":
		fileimage = "Workload_Data/DVD-training.csv"

	default:
		fileimage = "Workload_Data/NDBench-training.csv"

	}
	//get data from csv file
	file, err := os.Open(fileimage)
	if err != nil {
		log.Fatal(err.Error())
	}
	r := csv.NewReader(file)
	// Iterate through the records
	datas := new(DATAs)

	for {
		// Read each data from csv
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
		// writing the files to an array for better manipulation
		dvd.CpuUtilizationAverage = int32(i)
		dvd.NetworkInAverage = int32(j)
		dvd.NetworkOutAverage = int32(k)
		dvd.MemoryUtilizationAverage = float32(s)
		dvd.FinalTarget = float32(m)
		workload.BSize++
		datas.AddItems(*dvd)
	}

	return datas.WORKLOADs, workload
}

/*
Retrieve Datasets
Hard-coded for test of Dataset functoin at backend
returns the full data to be processed, dataset size Bsize

*/
func DataSet2(ctx echo.Context) ([]WORKLOAD, *library.DataParamsRequest) {

	workload := &library.DataParamsRequest{

		BenchmarkType:  "NDBENCH",
		WorkloadMetric: "CPU",
		BatchUnit:      3,
		BatchID:         3,
		BSize: 0,
		BatchSize:      4,
	}
	fmt.Print(workload)


	workload.BSize = 0

	var fileimage string

	switch workload.BenchmarkType {
	case "DVD":
		fileimage = "Workload_Data/DVD-training.csv"

	default:
		fileimage = "Workload_Data/NDBench-training.csv"

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
		dvd.NetworkInAverage = int32(j)
		dvd.NetworkOutAverage = int32(k)
		dvd.MemoryUtilizationAverage = float32(s)
		dvd.FinalTarget = float32(m)
		workload.BSize++
		datas.AddItems(*dvd)
	}

	return datas.WORKLOADs, workload
}
