package filewriter

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"
)

type FileWriter struct {
	fp    	*os.File
	ch		chan *bytes.Buffer
	buf		*bytes.Buffer
	wg		sync.WaitGroup
	closed	int32
}

func New(fpath string) (f *FileWriter, err error) {
	fname := filepath.Base(fpath)
	if fname == "" {
		return nil, fmt.Errorf("log文件名为空")
	}
	dir := filepath.Dir(fpath)
	fi, err := os.Stat(dir)
	if err == nil && !fi.IsDir() {
		return nil, fmt.Errorf("%s log目录已经存在", dir)
	}
	if os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0755); err != nil {
			return nil, fmt.Errorf("%s 创建log目录失败: %s", dir, err.Error())
		}
	}
	fp, err := os.OpenFile(fpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	f = &FileWriter{
		fp: fp,
		ch: make(chan *bytes.Buffer, 1024),
		buf: new(bytes.Buffer),
	}
	f.wg.Add(1)
	go f.worker()
	return
}

/**
 * 实现io.write接口
 * 内容存入channel等待worker消费
 */
func (f *FileWriter) Write(p []byte) (int, error) {
	if atomic.LoadInt32(&f.closed) == 1 {
		return 0, fmt.Errorf("日志已关闭")
	}
	buf := f.buf
	buf.Write(p)
	select {
	case f.ch <- buf:
		return len(p), nil
	default:
		return 0, errors.New("队列已到上限，日志将被丢弃")
	}
}

/**
 * 平滑退出
 * 先关闭channel，等待worker退出后再退出
 */
func (f *FileWriter) Close() {
	atomic.StoreInt32(&f.closed, 1)
	close(f.ch)
	f.wg.Wait()
}

/**
 * gorouting
 * 先将内容暂存到buffer，再定时批量写文件。closed用来保证退出前buffer和channel内容已全部写入。
 */
func (f *FileWriter) worker() {
	tbuf := &bytes.Buffer{}
	aggstk := time.NewTicker(100 * time.Millisecond)
	var err error
	for {
		select {
		case buf, ok := <-f.ch:
			if ok {
				tbuf.Write(buf.Bytes())
			}
		case <-aggstk.C:
			if tbuf.Len() > 0 {
				if _, err = f.fp.Write(tbuf.Bytes()); err != nil {
					fmt.Printf("记录日志出错: %s\n", err)
				}
				tbuf.Reset()
			}
		}
		if atomic.LoadInt32(&f.closed) != 1 {
			continue
		}
		if _, err = f.fp.Write(tbuf.Bytes()); err != nil {
			fmt.Printf("记录日志出错: %s\n", err)
		}
		for buf := range f.ch {
			if _, err = f.fp.Write(buf.Bytes()); err != nil {
				fmt.Printf("记录日志出错: %s\n", err)
			}
		}
		break
	}
	f.wg.Done()
}