package compressor

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var files []string

func init() {
	filepath.Walk("./testdata",
		func(path string, info os.FileInfo, err error) error {
			rel, err := filepath.Rel("./testdata", path)
			if err != nil {
				return err
			}

			files = append(files, fmt.Sprintf("./testdata/%s", rel))
			return nil
		})
}

func compress() error {
	file, _ := os.Create("sample.zip")
	defer file.Close()

	zw := zip.NewWriter(file)
	defer zw.Close()

	for _, f := range files {
		err := func() error {
			fw, err := os.Open(f)
			defer fw.Close()

			info, err := fw.Stat()
			hdr, err := setHileHeader(info)
			if err != nil {
				panic(err)
			}

			wf, err := zw.CreateHeader(hdr)
			if err != nil {
				return err
			}

			contents, err := ioutil.ReadAll(fw)
			_, err = wf.Write(contents)
			if err != nil {
				return err
			}

			return nil
		}()

		if err != nil {
			return err
		}
	}

	return nil
}

func compressNew() error {
	file, _ := os.Create("sample.zip")
	zw := zip.NewWriter(file)

	com := &compression{
		p:  make(chan int, 14),
		wg: new(sync.WaitGroup),
		zw: zw,
	}

	for _, f := range files {
		com.wg.Add(1)
		go com.hoge(f)
	}
	com.wg.Wait()
	file.Close()

	return nil
}

type compression struct {
	p  chan int
	wg *sync.WaitGroup
	zw *zip.Writer
}

func (c *compression) hoge(f string) {
	defer c.wg.Done()
	c.p <- 1

	fw, err := os.Open(f)
	if err != nil {
		panic(err)
	}

	info, err := fw.Stat()
	if err != nil {
		fw.Close()
		panic(err)
	}

	fw.Close()

	hdr, err := setHileHeader(info)
	if err != nil {
		panic(err)
	}

	wf, err := c.zw.CreateHeader(hdr)
	if err != nil {
		panic(err)
	}
	contents, _ := ioutil.ReadAll(fw)
	wf.Write(contents)
	<-c.p
}

func setHileHeader(info os.FileInfo) (*zip.FileHeader, error) {
	hdr, err := zip.FileInfoHeader(info)
	if err != nil {
		return nil, err
	}

	local := time.Now().Local()

	//現時刻のオフセットを取得
	_, offset := local.Zone()

	//差分を追加
	hdr.SetModTime(hdr.ModTime().Add(time.Duration(offset) * time.Second))

	return hdr, nil
}
