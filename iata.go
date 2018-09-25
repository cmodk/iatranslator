package iatranslator

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

func get_first_string(ss []string) string {
	for _, s := range ss {
		if len(s) > 0 {
			return s
		}
	}
	return ""
}

func SplitCode(code string) (string, int, error) {
	reg_num := regexp.MustCompile("[0-9]")
	reg_char := regexp.MustCompile("[a-zA-Z]")

	ps := reg_num.Split(code, -1)
	if len(ps) == 0 {
		return "", 0, errors.New(fmt.Sprintf("No airline in code: %s\n", code))
	}

	airline := ps[0]

	//Insert the number also
	ps = reg_char.Split(code, -1)
	if len(ps) == 0 {
		return "", 0, errors.New(fmt.Sprintf("No flight number in code: %s\n", code))
	}

	flight_number, err := strconv.Atoi(get_first_string(ps))
	if err != nil {
		return "", 0, err
	}

	return airline, flight_number, nil
}

func GetFlightIATA(icao string) (string, int, error) {

	airline_icao, flight_number, err := SplitCode(icao)
	if err != nil {
		return "", -1, err
	}

	//Convert to strings

	codes := map[string]string{
		"SAS": "SK",
		"TAP": "TP",
		"AFR": "AF",
		"BEL": "SK",
		"FDX": "FX",
		"KLM": "KL",
		"FIN": "AY",
		"RYR": "FR",
		"NAX": "DY",
		"IBK": "D8",
		"WZZ": "W6",
		"DLH": "LH",
		"BAW": "BA",
	}

	iata_id, ok := codes[airline_icao]
	if !ok {
		return "", -1, errors.New(fmt.Sprintf("Unknown icao code: %s", airline_icao))
	}

	return iata_id, flight_number, nil
}
