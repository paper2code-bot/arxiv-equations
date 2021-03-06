package arxiv

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func DownloadTarball(url string, path string) error {
	// Create the file
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func parseXML(xmlStr string) Feed {
	var result Feed
	if err := xml.Unmarshal([]byte(xmlStr), &result); err != nil {
		log.Fatal(err)
	}

	return result
}

func SearchPapers(params map[string]string) (Feed, error) {
	// define api url
	u, err := url.Parse("http://export.arxiv.org/api/query")
	if err != nil {
		return Feed{}, err
	}

	// construct query string
	q := u.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()

	// send the request
	resp, err := http.Get(u.String())
	if err != nil {
		return Feed{}, err
	}

	// parse result xml
	xmlBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Feed{}, err
	}
	xmlObj := parseXML(string(xmlBytes))

	return xmlObj, nil
}
