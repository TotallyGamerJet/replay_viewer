package utils

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

//GetZippedFile takes the zippedFile and searchs for the target file inside. It will either return the found file as a bytes.Buffer or an error
func GetZippedFile(zippedFile, target string) (*bytes.Buffer, error) {
	// Open a zip archive for reading.
	r, err := zip.OpenReader(zippedFile)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	for _, f := range r.File {
		//fmt.Println("Found: ", f.Name)
		if f.Name == target {
			rc, err := f.Open()
			if err != nil {
				log.Fatal(err)
			}
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				log.Fatal(err)
			}
			rc.Close()
			return bytes.NewBuffer(data), nil
		}
	}
	return nil, fmt.Errorf("Failed to locate file %s in zipped file %s", target, zippedFile)
}

//Unzip unzips the givin file 'src' into 'dest'
/*func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}
*/
