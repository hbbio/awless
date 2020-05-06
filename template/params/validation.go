package params

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"os"
	"strings"
)

// Validate checks if `paramValues` match `all` Validators.
func Validate(all Validators, paramValues map[string]interface{}) error {
	msg := bytes.NewBufferString("param validation:")
	var hasErr bool
	for key, vFn := range all {
		if val, ok := paramValues[key]; ok {
			if err := vFn(val, paramValues); err != nil {
				hasErr = true
				msg.WriteString(fmt.Sprintf("\n\t\t- param '%s': %s", key, err))
			}
		}
	}
	if hasErr {
		return errors.New(msg.String())
	}
	return nil
}

// validatorFunc defines a validator function.
type validatorFunc func(val interface{}, others map[string]interface{}) error

// Validators are a map of validator functions.
type Validators map[string]validatorFunc

// IsInEnumIgnoreCase checks if value is in `items` (case insensitive).
func IsInEnumIgnoreCase(items ...string) validatorFunc {
	included := func(arr []string, s string) bool {
		for _, a := range arr {
			if strings.ToLower(s) == strings.ToLower(a) {
				return true
			}
		}
		return false
	}
	return func(i interface{}, others map[string]interface{}) error {
		s, err := toString(i)
		if err != nil {
			return err
		}
		if !included(items, s) {
			return fmt.Errorf("expected any of %s but got '%s'", items, s)
		}
		return nil
	}
}

// MaxLengthOf checks if value length is less than `l`.
func MaxLengthOf(l int) validatorFunc {
	return func(i interface{}, others map[string]interface{}) error {
		s, err := toString(i)
		if err != nil {
			return err
		}
		if actual := len(s); actual > l {
			return fmt.Errorf("expected max length of %d but got %d", l, actual)
		}
		return nil
	}
}

// MinLengthOf checks if value length is at least `l`.
func MinLengthOf(l int) validatorFunc {
	return func(i interface{}, others map[string]interface{}) error {
		s, err := toString(i)
		if err != nil {
			return err
		}
		if actual := len(s); actual < l {
			return fmt.Errorf("expected min length of %d but got %d", l, actual)
		}
		return nil
	}
}

// IsFilepath checks if value is a file path.
func IsFilepath(i interface{}, others map[string]interface{}) error {
	filepath, err := toString(i)
	if err != nil {
		return err
	}
	stat, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return fmt.Errorf("cannot find file '%s'", filepath)
	}
	if err != nil {
		return err
	}
	if stat.IsDir() {
		return fmt.Errorf("'%s' is a directory", filepath)
	}
	return nil
}

func IsCIDR(i interface{}, others map[string]interface{}) error {
	s, err := toString(i)
	if err != nil {
		return err
	}
	_, _, err = net.ParseCIDR(s)
	return err

}

func IsIP(i interface{}, others map[string]interface{}) (err error) {
	s, err := toString(i)
	if err != nil {
		return
	}
	if ip := net.ParseIP(s); ip == nil {
		err = fmt.Errorf("expected valid IP address but got '%s'", s)
	}
	return
}

func toString(i interface{}) (string, error) {
	s, ok := i.(string)
	if !ok {
		return s, fmt.Errorf("expected a string but got %T", i)
	}
	return s, nil
}
