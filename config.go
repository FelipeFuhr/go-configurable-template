package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
)

// Change these values to add more configurations
type appConfig struct {
	Name             string
	Description      string
	ServerAddress    string
	OptionalArgument string
}

var optionalFields []string

func init() {
	// Add fields to specify they are optional
	optionalFields = append(optionalFields, "OptionalArgument",
		"OptionalArgument2",
		"OptionalArgument3")
}

// Reads the JSON Configuration File for the Application
func ReadConfig(cfName string) (appConfig, error) {
	config := appConfig{}

	cf, err := os.Open(cfName)
	if err != nil {
		return config, err
	}
	defer cf.Close()

	d := json.NewDecoder(cf)
	if err := d.Decode(&config); err != nil {
		return config, err
	}

	err = checkMandatoryFields(config)

	return config, err
}

// Function to check whether the mandatory JSON fields are present in the config file
func checkMandatoryFields(config appConfig) error {
	// Missing Required Fields
	var mrfs []string
	// Optional Fields
	optfs := optionalFields
	// Use reflect to get Name and Value of each Config Struct Element
	v := reflect.ValueOf(&config).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i).Name
		fval := v.Field(i).Interface()

		isfNil := reflect.ValueOf(fval).IsZero()
		var isfReq bool
		if isfReq, optfs = checkIsFieldRequired(field, optfs); isfNil && isfReq {
			mrfs = append(mrfs, field)
		}
	}
	if len(mrfs) > 0 {
		return &MissingFieldError{
			MissingFields: mrfs,
			Err:           errors.New("Required Field(s) Missing in JSON config file"),
		}
	} else {
		return nil
	}
}

func checkIsFieldRequired(f string, optfs []string) (bool, []string) {
	isfReq := true
	for i, v := range optfs {
		if strings.ToLower(f) == strings.ToLower(v) {
			isfReq = false
			defer func() {
				if len(optfs) > 1 {
					optfs[i] = optfs[len(optfs)-1]
					optfs[len(optfs)-1] = ""
					optfs = optfs[:len(optfs)-1]
				} else {
					optfs = []string{}
				}
			}()
		}
	}
	return isfReq, optfs
}

// Error used when mandatory JSON fields are not in the config file
type MissingFieldError struct {
	MissingFields []string
	Err           error
}

func (e *MissingFieldError) Error() string {
	return fmt.Sprintf("%v: %s", e.Err, e.MissingFields)
}
