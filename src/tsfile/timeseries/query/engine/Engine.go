package engine

import (
	"tsfile/file/metadata"
	"tsfile/timeseries/read"
	"tsfile/timeseries/query"
	"tsfile/timeseries/query/dataset"
	"log"
	"fmt"
	"tsfile/common/constant"
	"tsfile/timeseries/read/reader"
	impl2 "tsfile/timeseries/query/dataset/impl"
	"tsfile/timeseries/read/reader/impl/basic"
)

type Engine struct {
	reader *read.TsFileSequenceReader
	fileMeta *metadata.FileMetaData
}

func (e *Engine) Open(reader *read.TsFileSequenceReader) {
	e.reader = reader
	e.fileMeta = reader.ReadFileMetadata()
}

func (e *Engine) Close() {
	e.reader.Close()
	e.reader = nil
	e.fileMeta = nil
}

func (e *Engine) Query(exp *query.QueryExpression) dataset.IQueryDataSet{
	readerMap := make(map[string]reader.ISeriesReader)
	for _, path := range exp.SelectPaths(){
		readerMap[path] = e.constructReader(path)
	}
	for _, path := range exp.ConditionPaths(){
		if _, ok := readerMap[path]; !ok {
			readerMap[path] = e.constructReader(path)
		}
	}
	dataSet := impl2.NewQueryDataSet(exp.SelectPaths(), exp.ConditionPaths(), readerMap, exp.Filter())
	return dataSet
}

func (e *Engine) constructReader(path string) reader.ISeriesReader {
	dataType := e.getDataType(path)
	if dataType == constant.INVALID {
		log.Println(fmt.Sprintf("No such timeseries in this file : %s", path))
		return  nil
	}

	deviceMeta, ok := e.fileMeta.DeviceMap()[path]
	if !ok {
		log.Println(fmt.Sprintf("No such timeseries in this file : %s", path))
		return  nil
	}

	var offsets []int64
	var sizes []int
	// var headers []*header.PageHeader
	// find the offsets, sizes and headers(optional) of all pages of this path
	for ele := deviceMeta.RowGroupMetadataList().Front(); ele != nil; ele = ele.Next() {
		if rowGroupMeta, ok := ele.Value.(metadata.RowGroupMetaData); ok {
			for c := rowGroupMeta.TimeSeriesChunkMetaDataList().Front(); c != nil; c = c.Next() {
				if chunkMeta, ok := c.Value.(metadata.ChunkMetaData); ok {
					chunkHeader := e.reader.ReadChunkHeaderAt(chunkMeta.FileOffsetOfCorrespondingData())
					for i := 0; i < chunkHeader.GetNumberOfPages(); i ++ {
						pageHeader := e.reader.ReadPageHeader(dataType)
						offsets = append(offsets, e.reader.Pos())
						sizes = append(sizes, pageHeader.GetCompressedSize())
						// headers = append(headers, pageHeader)
					}
				}
			}
		}
	}
	return basic.NewSeriesReader(offsets, sizes, e.reader)
}

func (e* Engine) getDataType(path string) constant.TSDataType {
	if tsMeta, ok := e.fileMeta.TimeSeriesMetadataMap()[path]; ok {
		return tsMeta.DataType()
	}
	return constant.INVALID
}