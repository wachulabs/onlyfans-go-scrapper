//This is the auth package.

package auth

import (
	"crypto/sha1"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// CreateHeader creates a http.Header property of the API request
func CreateHeader(dynamicRules map[string]any, authID, userAgent, cookie, xBC string) http.Header {
	header := http.Header{}
	header.Set("user-agent", userAgent)
	header.Set("user-id", authID)
	header.Set("accept", "application/json, text/plain, */*")
	header.Set("accept-encoding", "gzip, deflate")
	header.Set("app-token", dynamicRules["app_token"].(string))
	header.Set("cookie", cookie)
	header.Set("x-bc", xBC)
	header.Set("referrer", "https://onlyfans.com/")
	return header
}

// HeaderSign signs the API header. API header must be signed
func SignHeader(endpointlink string, dynamicRules map[string]any, header http.Header) http.Header {
	unixtime := time.Now().Unix()
	msg := dynamicRules["static_param"].(string) + "\n" + strconv.Itoa(int(unixtime)) + "\n" + endpointlink + "\n" + header.Get("user-id")
	h := sha1.New()
	h.Write([]byte(msg))
	Sha1 := hex.EncodeToString(h.Sum(nil))
	var cs []byte
	ci := dynamicRules["checksum_indexes"].([]int)
	for _, val := range ci {
		cs = append(cs, Sha1[val])
	}
	checksum := ArraySum(cs) + dynamicRules["checksum_constant"].(int)
	checksum1 := fmt.Sprintf("%x", checksum)
	Format := dynamicRules["format"].(string)
	header.Set("sign", StringFormat(Format, "{}", Sha1, "{:x}", checksum1))
	header.Set("time", strconv.Itoa(int(unixtime)))
	return header

}

func StringFormat(format string, args ...string) string {
	r := strings.NewReplacer(args...)
	return (r.Replace(format))
}

func Float64FromBytes(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

// Config aid to load config files from a
type Config struct {
	UserAgent string `json:"user-agent"`
	UserID    string `json:"user-id"`
	AppToken  string `json:"app-token"`
	Session   string `json:"cookie"`
	XBC       string `json:"x-bc"`
}

// Create Header from file creates an API header from a configuraton file
// The configuration file might be a json file, a text file or a .config file

func CreateHeaderFromFile(pathToFile, endpointlink string, dynamicRules map[string]any) (http.Header, error) {
	if _, err := os.Stat(pathToFile); errors.Is(err, os.ErrNotExist) {
		return nil, errors.New("File does not exist")
	}
	//Filtering too big files
	fileSize, _ := os.Stat(pathToFile)
	if (fileSize.Size() / int64(1000000)) > int64(20) {
		return nil, errors.New("The file is too big")
	}

	//Checking if a file is a .json file
	if filepath.Ext(pathToFile) == ".json" {
		jsonFile, err := os.Open(pathToFile)
		if err != nil {
			return nil, err
		}
		var config Config
		var header http.Header
		defer jsonFile.Close()
		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &config)

		header = CreateHeader(dynamicRules, config.UserID, config.UserAgent, config.Session, config.XBC)
		header = SignHeader(endpointlink, dynamicRules, header)
		return header, nil
	} else {
		return nil, errors.New("Handling for: " + filepath.Ext(pathToFile) + "extension not handled")
	}
}

func ArraySum(arr []byte) int {
	var sum int
	for _, val := range arr {
		sum += int(val)
	}
	return sum
}
