package utils

import (
	"encoding/json"
	"fmt"
	"hz-admin-base/utils/gurl"
)

type IpInfo struct {
	Status      string `json:"status"`
	CountryCode string `json:"countryCode"`
	Region      string `json:"region"`
	City        string `json:"city"`
	Proxy       bool   `json:"proxy"`
}

func GetIpInfo(ip string) (*IpInfo, error) {
	ipInfo := &IpInfo{}
	do, err := gurl.NewGurl("get", fmt.Sprintf("https://pro.ip-api.com/json/%s?key=MCFIdiwEo8ahaL7&fields=status,message,countryCode,region,city,proxy", ip)).Do()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(do, &ipInfo)
	if err != nil {
		return nil, err
	}
	if ipInfo.Status != "success" {
		return nil, fmt.Errorf("ip info error")
	}
	return ipInfo, nil
}
