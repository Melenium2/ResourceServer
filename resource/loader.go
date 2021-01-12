package resource

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"time"
)

// Load function upload image by the 'rawurl' to 'topath' folder
// and generate unique filename for jpg picture then return this
// filename to caller
func Load(topath, rawurl string) (string, error) {
	urlstruct, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}
	filename := filename(urlstruct.Path)
	hashedname := fmt.Sprintf("%s.jpg", hash(filename, time.Now().UTC().String()))

	resp, err := request(rawurl)
	if err != nil {
		return "", err
	}
	if resp.StatusCode > 200 {
		return "", fmt.Errorf("url %s response with %d code", rawurl, resp.StatusCode)
	}


	defer resp.Body.Close()

	if err = Save(path.Join(topath, hashedname), resp.Body); err != nil {
		return "", err
	}

	return hashedname, nil
}

// Save copy io.ReadCloser body to file
func Save(filepath string, body io.ReadCloser) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, body)
	if err != nil {
		return err
	}

	return nil
}

// Generate md5 hash with 'str' and 'salt' then return string to caller
func hash(str, salt string) string {
	strWithTime := fmt.Sprintf("%s_%s", str, salt)
	hasher := md5.New()
	hasher.Write([]byte(strWithTime))

	return hex.EncodeToString(hasher.Sum(nil))
}

// filename function return last piece of given url
func filename(url string) string {
	urlpieces := strings.Split(url, "/")

	return urlpieces[len(urlpieces)-1]
}

func request(rawurl string) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", rawurl, nil)
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36")

	return client.Do(req)
}
