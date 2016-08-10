package controllers

import (
	"archive/zip"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type FileAttr struct {
	Name     string
	Path     string
	FullName string
	Ext      string
	Size     int64
	ModTime  time.Time
	IsDir    bool
}

//压缩文件
//files 文件数组，可以是不同dir下的文件或者文件夹
//dest 压缩文件存放地址
func Compress(files []*os.File, dest string) error {
	d, _ := os.Create(dest)
	defer d.Close()
	w := zip.NewWriter(d)
	defer w.Close()
	for _, file := range files {
		err := compress(file, "", w)
		if err != nil {
			return err
		}
	}
	return nil
}

const zone int64 = +8

//压缩
func compress(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, zw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

//解压
func DeCompress(zipFile, dest string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()
	//解压文件
	for _, file := range reader.File {
		filename := dest + file.Name
		if file.FileInfo().IsDir() {
			filename = dest + "/" + file.Name
			err = os.MkdirAll(filename, 0755)
			if err != nil {
				return err
			}
			continue
		}
		w, err := os.Create(filename)
		defer w.Close()

		for err != nil {
			end := strings.LastIndex(filename, "/")
			var p string
			if end > -1 {
				p = subString(filename, 0, end)
			}
			if Exist(p) {
				return err
			}
			//目录不存在
			err = os.MkdirAll(p, 0755)
			if err != nil {
				return err
			}
			w, err = os.Create(filename)
		}
		rc, err := file.Open()
		defer rc.Close()
		if err != nil {
			return err
		}
		_, err = io.Copy(w, rc)
		if err != nil {
			return err
		}
		w.Close()
		rc.Close()
	}
	return nil
}

func getDir(path string) string {
	return subString(path, 0, strings.LastIndex(path, "/"))
}

func subString(str string, start, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < start || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}

//文件或文件夹是否存在
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

//获取随机字符串
func RandString() (str string) {
	rand.Seed(time.Now().UnixNano())
	rnd := rand.Intn(999999)
	str = strconv.FormatInt(time.Now().Unix()+int64(rnd), 16)
	return
}

//拷贝文件
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
func ReadFileBytes(path string) []byte {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return fd
}
func ReadFileString(path string) string {
	return string(ReadFileBytes(path))
}
func GzipBytes(bs []byte) []byte {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	defer w.Close()
	w.Write(bs)
	w.Flush()
	return b.Bytes()
}

//获取指定目录下的所有文件和目录
func DirList(dirPth string) (list []string, err error) {
	//fmt.Println(dirPth)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	//PthSep := string(os.PathSeparator)
	PthSep := "/"
	// suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		list = append(list, dirPth+PthSep+fi.Name())
	}
	return list, nil
}
func FileAttrList(dirPth string) (list []FileAttr, err error) {
	//fmt.Println(dirPth)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	//PthSep := string(os.PathSeparator)
	PthSep := "/"
	// suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		fa := FileAttr{Name: fi.Name(), Path: dirPth, FullName: dirPth + PthSep + fi.Name(), ModTime: fi.ModTime(), Size: fi.Size(), IsDir: fi.IsDir()}
		if !fa.IsDir {
			fa.Ext = path.Ext(fa.FullName)
		}
		list = append(list, fa)
	}
	return list, nil
}
func ToJsonFile(v interface{}, jfile string) error {
	cfg, _ := json.Marshal(v)
	return ioutil.WriteFile(jfile, cfg, 0666) //写入文件(字节数组)
}
func StringToFile(content, file string) error {
	var bytes = []byte(content)
	return ioutil.WriteFile(file, bytes, 0666)
}
func BytesToFile(bytes []byte, file string) error {
	return ioutil.WriteFile(file, bytes, 0666)
}

// package ziptest

// import (
// 	"os"
// 	"testing"
// )

// func TestCompress(t *testing.T) {
// 	f1, err := os.Open("/home/zzw/test_data/ziptest/gophercolor16x16.png")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer f1.Close()
// 	f2, err := os.Open("/home/zzw/test_data/ziptest/readme.notzip")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer f2.Close()
// 	f3, err := os.Open("/home/zzw/test_data")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer f3.Close()
// 	var files = []*os.File{f1, f2, f3}
// 	dest := "/home/zzw/test_data/test.zip"
// 	err = Compress(files, dest)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }
// func TestDeCompress(t *testing.T) {
// 	err := DeCompress("/home/zzw/test_data/test.zip", "/home/zzw/test_data/de")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }
