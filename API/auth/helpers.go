package auth

import (
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

// RespToJson extracts response body compressed using gzip compression scheme
// It helps in instances where trying to decode response body leads to garbled string
// It gives the user the control to save the response to a json file or not
func RespToJSON(resp *http.Response, saveToFile bool, directory ...string) (string, error) {

	if resp.StatusCode != http.StatusOK {
		s := "Got " + strconv.Itoa(resp.StatusCode) + " status code"
		e := errors.New(s)
		return "", errors.Wrap(e, "Status code not 200:Ok!")
	}
	defer resp.Body.Close()
	var u User
	gReader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "Could not extract response body")
	}
	defer gReader.Close()
	gBytes, err := ioutil.ReadAll(gReader)
	if err != nil {
		return "", errors.Wrap(err, "Unable to read uncompressed content")
	}
	err = json.Unmarshal(gBytes, &u)
	if err != nil {
		return "", errors.Wrap(err, "Error unmarshalling bytes from the reponse body")
	}
	respJson, err := json.MarshalIndent(u, "", "\t")
	if err != nil {
		return "", errors.Wrap(err, "Could not Marshal Json")
	}
	if saveToFile {
		if len(directory) != 0 {
			directoryExists, _ := ConfirmDirectoryExistence(directory[0])
			if directoryExists {
				err = ioutil.WriteFile("output.json", respJson, 0644)
				if err != nil {
					return "", errors.Wrap(err, "Unable to write json file")
				}
			} else {
				err = CreateDirectory(directory[0])
				if err != nil {
					return "", errors.Wrap(err, "Write directory does not exist and could not be created")
				}
				err = ioutil.WriteFile("output.json", respJson, 0644)
				if err != nil {
					return "", errors.Wrap(err, "Unable to write json file")
				}
			}
		} else {
			err = ioutil.WriteFile("output.json", respJson, 0644)
			if err != nil {
				return "", errors.Wrap(err, "Unable to write json file")
			}
		}
	}
	return string(respJson), nil
}
