package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Reference struct {
	Datacenters []Datacenter `json:"datacenters"`
	Region      string       `json:"region"`
	Hardware    string       `json:"hardware"`
}

type Datacenter struct {
	Datacenter   string `json:"datacenter"`
	Availability string `json:"availability"`
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		log.Println("Nothing to check...")
		os.Exit(0)
	}

	for {
		if err := CheckAvailability(args); err != nil {
			log.Println(err)
		}
		time.Sleep(30 * time.Second)
	}
}

func CheckAvailability(references []string) error {
	resp, err := http.Get("https://www.ovh.com/engine/api/dedicated/server/availabilities?country=we")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	availabilities := []Reference{}
	err = json.Unmarshal(body, &availabilities)
	if err != nil {
		return err
	}

	for _, availability := range availabilities {
		if IsIndex(references, availability.Hardware) < len(references) && availability.Region == "europe" {
			for _, zone := range availability.Datacenters {
				if zone.Availability != "unavailable" {
					PrintAvailability(availability)
				}
			}
		}
	}

	return nil
}

func PrintAvailability(r Reference) {
	fmt.Printf("Reference %s (%s) is available on these datacenters :\n", r.Hardware, r.Region)
	for _, zone := range r.Datacenters {
		if zone.Availability != "unavailable" && zone.Datacenter != "default" {
			fmt.Printf("- %s (%s)\n", zone.Datacenter, zone.Availability)
		}
	}
}

// Returns the position of the lf element in the array. If the element is not found then return the length of the array.
func IsIndex[T comparable](array []T, lf T) int {
	for i, value := range array {
		if value == lf {
			return i
		}
	}

	return len(array)
}
