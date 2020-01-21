package fequal

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

const ByteBufferSize = 4000

func AreEqual(fName1, fName2 string) (bool, error) {
	f1, err := os.Open(fName1)
	if err != nil {
		return false, err
	}
	defer f1.Close()

	f2, err := os.Open(fName2)
	if err != nil {
		return false, err
	}
	defer f2.Close()

	fs1, err := f1.Stat()
	if err != nil {
		return false, err
	}

	fs2, err := f2.Stat()
	if err != nil {
		return false, err
	}

	if os.SameFile(fs1, fs2) {
		return true, nil
	}

	if fs1.Size() != fs2.Size() {
		return false, nil
	}

	for {
		b1 := make([]byte, ByteBufferSize)
		_, err1 := f1.Read(b1)

		b2 := make([]byte, ByteBufferSize)
		_, err2 := f2.Read(b2)

		if err1 != nil || err2 != nil {
			if err1 == io.EOF && err2 == io.EOF {
				return true, nil
			} else if err1 == io.EOF || err2 == io.EOF {
				return false, nil
			} else {
				return false, errors.New(fmt.Sprintf("file contents were read in with the following errors: file1: %s, file2: %s\n", err1, err2))
			}
		}

		if !bytes.Equal(b1, b2) {
			return false, nil
		}
	}
}

func AreEqualTimed(fName1, fName2 string) (bool, error) {
	f1, err := os.Open(fName1)
	if err != nil {
		return false, err
	}
	defer f1.Close()

	f2, err := os.Open(fName2)
	if err != nil {
		return false, err
	}
	defer f2.Close()

	start := time.Now()
	err = f1.SetReadDeadline(start.Add(time.Second * 30))
	if err != nil {
		fmt.Printf("The first file %s one can not have a timeout set\n", fName1)
	}

	err = f2.SetReadDeadline(start.Add(time.Second * 30))
	if err != nil {
		fmt.Printf("The second file %s one can not have a timeout set\n", fName2)
	}

	fs1, err := f1.Stat()
	if err != nil {
		return false, err
	}

	fs2, err := f2.Stat()
	if err != nil {
		return false, err
	}

	if os.SameFile(fs1, fs2) {
		return true, nil
	}

	if fs1.Size() != fs2.Size() {
		return false, nil
	}

	for {
		b1 := make([]byte, ByteBufferSize)
		_, err1 := f1.Read(b1)

		b2 := make([]byte, ByteBufferSize)
		_, err2 := f2.Read(b2)

		if err1 != nil || err2 != nil {
			if err1 == io.EOF && err2 == io.EOF {
				return true, nil
			} else if err1 == io.EOF || err2 == io.EOF {
				return false, nil
			} else if os.IsTimeout(err1) || os.IsTimeout(err2) {
				return false, errors.New(fmt.Sprintf("One of the files timed out while being read.\nFirst Filed timed out? %t\nSecond file timed out? %t\n", os.IsTimeout(err1), os.IsTimeout(err2)))
			} else {
				return false, errors.New(fmt.Sprintf("file contents were read in with the following errors: file1: %s, file2: %s\n", err1, err2))
			}
		}

		if !bytes.Equal(b1, b2) {
			return false, nil
		}
	}
}
