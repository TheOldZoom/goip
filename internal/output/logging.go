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
package output

import (
	"fmt"
	"goip/internal/ipinfo"
	"strings"
)

func FormatIPInfo(info ipinfo.LookupResponse) string {
	var builder strings.Builder

	writeField := func(label, value string) {
		if value == "" {
			return
		}
		fmt.Fprintf(&builder, "%-10s %s\n", label+":", value)
	}

	writeField("IP", info.Query)
	writeField("Country", formatCountry(info.Country, info.CountryCode))
	writeField("Region", formatRegion(info.RegionName, info.Region))
	writeField("City", info.City)
	writeField("ZIP", info.Zip)
	writeField("Timezone", info.Timezone)
	writeField("ISP", info.ISP)
	writeField("Org", info.Org)
	writeField("AS", info.AS)

	if info.Lat != 0 || info.Lon != 0 {
		fmt.Fprintf(&builder, "%-10s %.4f, %.4f\n", "Coords:", info.Lat, info.Lon)
	}

	return strings.TrimRight(builder.String(), "\n")
}

func formatCountry(country, code string) string {
	if country == "" {
		return ""
	}
	if code == "" {
		return country
	}
	return fmt.Sprintf("%s (%s)", country, code)
}

func formatRegion(regionName, regionCode string) string {
	if regionName == "" {
		return ""
	}
	if regionCode == "" {
		return regionName
	}
	return fmt.Sprintf("%s (%s)", regionName, regionCode)
}
