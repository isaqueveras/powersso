// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package geoip

import (
	"net"

	"github.com/isaqueveras/powersso/config"
	"github.com/oschwald/geoip2-golang"
)

var db *geoip2.Reader

func Load() (err error) {
	db, err = geoip2.Open(config.Get().Server.GeoIPDatabase)
	return
}

func Close() { _ = db.Close() }

func Get(ip net.IP) (l string) {
	if db == nil {
		return "Location unavailable"
	}

	city, err := db.City(ip)
	if err != nil {
		return "Undefined location"
	}

	language := config.Get().Server.GeoIPLocation
	l = city.City.Names[language]

	if len(city.Subdivisions) > 0 {
		for i := range city.Subdivisions {
			if len(city.Subdivisions[i].Names) > 0 && (city.Subdivisions[i].Names[language] != "") {
				l += " - " + city.Subdivisions[i].Names[language]
			}
		}
	}

	l += ", " + city.Country.Names[language]
	return
}
