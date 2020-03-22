package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func Test_uploadFileHandler(t *testing.T) {
	url := "localhost:9092/v1.0/owners/7/students/72/classrooms/45/files"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open("/Users/josue/Documents/Cotización.pdf")
	defer file.Close()
	part1,
		errFile1 := writer.CreateFormFile("file", filepath.Base("/Users/josue/Documents/Cotización.pdf"))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {

		t.Error(errFile1)
	}
	err := writer.Close()
	if err != nil {
		t.Error(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		t.Error(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("SessionID", "fadfasfasf")
	req.Header.Add("Authorization", "Basic dGVzdDp0ZXN0")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(body))
}
