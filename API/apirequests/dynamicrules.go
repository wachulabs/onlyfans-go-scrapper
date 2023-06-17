package apirequests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// DynamicRules stores dynamic rules
type DynamicRulesStruct struct {
	StaticParam       string   `json:"static_param"`
	Format            string   `json:"format"`
	CheckSumIndexes   []int    `json:"checksum_indexes"`
	CheckSumConstants []int    `json:"checksum_constants"`
	CheckSumConstant  int      `json:"checksum_constant"`
	AppToken          string   `json:"app_token"`
	RemoveHeaders     []string `json:"remove_headers"`
	ErrorCode         int      `json:"error_code"`
	Message           string   `json:"message"`
}

// getDynamicRule gets request from https://raw.githubusercontent.com/DIGITALCRIMINALS/dynamic-rules/main/onlyfans.json. It returns a DynamicRulesStruct
func getDymanicRules() DynamicRulesStruct {
	var dr DynamicRulesStruct
	url := "https://raw.githubusercontent.com/DIGITALCRIMINALS/dynamic-rules/main/onlyfans.json"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := client.Do(req)
	gBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(gBytes, &dr)
	return dr
}

// GetDynamicRule converts DynamicRulesStruct to a map
func GetDynamicRules() map[string]any {
	DynamicRules := getDymanicRules()
	dynamicRules := make(map[string]any)
	dynamicRules["static_param"] = DynamicRules.StaticParam
	dynamicRules["format"] = DynamicRules.Format
	dynamicRules["checksum_indexes"] = DynamicRules.CheckSumIndexes
	dynamicRules["checksum_constants"] = DynamicRules.CheckSumConstants
	dynamicRules["checksum_constant"] = DynamicRules.CheckSumConstant
	dynamicRules["app_token"] = DynamicRules.AppToken
	dynamicRules["remove_headers"] = DynamicRules.RemoveHeaders
	dynamicRules["error_code"] = DynamicRules.ErrorCode
	dynamicRules["message"] = DynamicRules.Message
	return dynamicRules
}
