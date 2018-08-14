package main

import (
	"os"
	"sync"
	"github.com/kataras/iris/core/errors"
	"io"
)

type DataFile interface {
	Read() (rsn int64, d Data, err error)
	Write(d Data) (wsn int64, err error)
	Rsn() int64
	Wsn() int64
	DataLen() uint32
}

type Data []byte

type myDataFile struct {
	f        *os.File
	fmutex   sync.RWMutex
	woffset  int64
	roffset  int64
	wmutex   sync.Mutex
	rmutex   sync.Mutex
	dataLen  uint32
}



func main()  {
	NewDataFile("./", 200)
}


func NewDataFile(path string, dataLen uint32) (DataFile, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	if dataLen == 0 {
		return nil, errors.New("没有定义的长度")
	}

	df := &myDataFile{f: f, dataLen: dataLen}
	return df, nil
}

func (df *myDataFile) Read() (rsn int64, d Data, err error) {
	var offset int64
	df.rmutex.Lock()
	offset = df.roffset
	df.roffset += int64(df.dataLen)
	df.rmutex.Unlock()

	rsn = offset / int64(df.dataLen)
	bytes := make([]byte, df.dataLen)
	for {
		df.fmutex.RLock()
		_, err = df.f.ReadAt(bytes, offset)
		if err != nil {
			if err == io.EOF {
				df.fmutex.RUnlock()
				continue
			}
			df.fmutex.RUnlock()
			return
		}
		d = bytes
		df.fmutex.RUnlock()
	}
}

func (df *myDataFile) Write(d Data) (wsn int64, err error)  {
	var offset int64
	df.wmutex.Lock()
	offset = df.woffset
	df.woffset += int64(df.dataLen)
	df.wmutex.Unlock()

	wsn = offset / int64(df.dataLen)
	var bytes []byte
	if len(d) > int(df.dataLen) {
		bytes = d[0:df.dataLen]
	} else {
		bytes = d
	}
	df.fmutex.Lock()
	df.fmutex.Unlock()
	_, err = df.f.Write(bytes)
	return
}

func (df *myDataFile) Rsn() int64 {
	df.rmutex.Lock()
	defer df.rmutex.Unlock()
	return df.roffset / int64(df.dataLen)
}

func (df *myDataFile) Wsn() int64 {
	df.wmutex.Lock()
	defer df.wmutex.Unlock()
	return df.woffset / int64(df.dataLen)
}

func (df *myDataFile) DataLen() uint32 {
	lengthWord := new(myDataFile)
	return lengthWord.dataLen
}
