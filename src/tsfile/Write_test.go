// tsfile project main.go
package main

import (
	"os"
	"testing"
	"time"
	"tsfile/common/constant"
	"tsfile/timeseries/write/sensorDescriptor"
	"tsfile/timeseries/write/tsFileWriter"
)

func TestWrite(t *testing.T) {
	fileName := "D:/test.ts"

	if _, err := os.Stat(fileName); !os.IsNotExist(err) {
		os.Remove(fileName)
	}

	// init tsFileWriter
	tfWriter, _ := tsFileWriter.NewTsFileWriter(fileName)

	// init sensorDescriptor
	sd1, _ := sensorDescriptor.New("sensor_1", constant.INT32, constant.RLE)
	sd2, _ := sensorDescriptor.New("sensor_2", constant.INT64, constant.RLE)
	sd3, _ := sensorDescriptor.New("sensor_3", constant.FLOAT, constant.RLE)
	sd4, _ := sensorDescriptor.New("sensor_4", constant.DOUBLE, constant.RLE)
	sd5, _ := sensorDescriptor.New("sensor_5", constant.INT32, constant.TS_2DIFF)
	sd6, _ := sensorDescriptor.New("sensor_6", constant.INT64, constant.TS_2DIFF)
	sd7, _ := sensorDescriptor.New("sensor_7", constant.FLOAT, constant.TS_2DIFF)
	sd8, _ := sensorDescriptor.New("sensor_8", constant.DOUBLE, constant.TS_2DIFF)
	sd9, _ := sensorDescriptor.New("sensor_9", constant.FLOAT, constant.GORILLA)
	sd10, _ := sensorDescriptor.New("sensor_10", constant.DOUBLE, constant.GORILLA)
	sd11, _ := sensorDescriptor.New("sensor_11", constant.INT32, constant.PLAIN)
	sd12, _ := sensorDescriptor.New("sensor_12", constant.INT64, constant.PLAIN)
	sd13, _ := sensorDescriptor.New("sensor_13", constant.FLOAT, constant.PLAIN)
	sd14, _ := sensorDescriptor.New("sensor_14", constant.DOUBLE, constant.PLAIN)
	sd15, _ := sensorDescriptor.New("sensor_15", constant.TEXT, constant.PLAIN)

	// add sensorDescriptor to tfFileWriter
	tfWriter.AddSensor(sd1)
	tfWriter.AddSensor(sd2)
	tfWriter.AddSensor(sd3)
	tfWriter.AddSensor(sd4)
	tfWriter.AddSensor(sd5)
	tfWriter.AddSensor(sd6)
	tfWriter.AddSensor(sd7)
	tfWriter.AddSensor(sd8)
	tfWriter.AddSensor(sd9)
	tfWriter.AddSensor(sd10)
	tfWriter.AddSensor(sd11)
	tfWriter.AddSensor(sd12)
	tfWriter.AddSensor(sd13)
	tfWriter.AddSensor(sd14)
	tfWriter.AddSensor(sd15)

	// create a tsRecord
	tr, _ := tsFileWriter.NewTsRecord(time.Now(), "device_1")

	// create data point
	fdp1, _ := tsFileWriter.NewInt("sensor_1", constant.INT32, 11)
	fdp2, _ := tsFileWriter.NewLong("sensor_2", constant.INT64, 1111111)
	fdp3, _ := tsFileWriter.NewFloat("sensor_3", constant.FLOAT, 11.1)
	fdp4, _ := tsFileWriter.NewDouble("sensor_4", constant.DOUBLE, 11.11111)
	fdp5, _ := tsFileWriter.NewInt("sensor_5", constant.INT32, 22)
	fdp6, _ := tsFileWriter.NewLong("sensor_6", constant.INT64, 2222222)
	fdp7, _ := tsFileWriter.NewFloat("sensor_7", constant.FLOAT, 22.2)
	fdp8, _ := tsFileWriter.NewDouble("sensor_8", constant.DOUBLE, 22.22222)
	fdp9, _ := tsFileWriter.NewFloat("sensor_9", constant.FLOAT, 33.3)
	fdp10, _ := tsFileWriter.NewDouble("sensor_10", constant.DOUBLE, 33.33333)
	fdp11, _ := tsFileWriter.NewInt("sensor_11", constant.INT32, 44)
	fdp12, _ := tsFileWriter.NewLong("sensor_12", constant.INT64, 4444444)
	fdp13, _ := tsFileWriter.NewFloat("sensor_13", constant.FLOAT, 44.4)
	fdp14, _ := tsFileWriter.NewDouble("sensor_14", constant.DOUBLE, 44.44444)
	fdp15, _ := tsFileWriter.NewString("sensor_15", constant.TEXT, "44.4abc")

	// add data points to ts record
	tr.AddTuple(fdp1)
	tr.AddTuple(fdp2)
	tr.AddTuple(fdp3)
	tr.AddTuple(fdp4)
	tr.AddTuple(fdp5)
	tr.AddTuple(fdp6)
	tr.AddTuple(fdp7)
	tr.AddTuple(fdp8)
	tr.AddTuple(fdp9)
	tr.AddTuple(fdp10)
	tr.AddTuple(fdp11)
	tr.AddTuple(fdp12)
	tr.AddTuple(fdp13)
	tr.AddTuple(fdp14)
	tr.AddTuple(fdp15)

	// write tsRecord to file
	tfWriter.Write(*tr)

	// close file descriptor
	tfWriter.Close()
}
