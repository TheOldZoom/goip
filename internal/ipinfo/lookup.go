/*
Copyright © 2026 Zoom theoldzoom@proton.me

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package ipinfo

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type LookupResponse struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	ISP         string  `json:"isp"`
	Org         string  `json:"org"`
	AS          string  `json:"as"`
	Query       string  `json:"query"`
	Message     string  `json:"message"`
}

var client = &http.Client{Timeout: 5 * time.Second}

func GetMyIP() (LookupResponse, error) {
	return GetIPInfo("")
}

func GetIPInfo(ip string) (LookupResponse, error) {
	url := "http://ip-api.com/json/"
	if ip != "" {
		url += ip
	}

	res, err := client.Get(url)
	if err != nil {
		return LookupResponse{}, err
	}
	defer res.Body.Close()

	var info LookupResponse
	if err := json.NewDecoder(res.Body).Decode(&info); err != nil {
		return LookupResponse{}, err
	}

	if info.Status != "" && info.Status != "success" {
		if info.Message == "" {
			info.Message = "lookup failed"
		}
		return info, errors.New(info.Message)
	}

	return info, nil
}
