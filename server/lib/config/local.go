package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	pkgerr "github.com/pkg/errors"
)

// LoadLocalSecrets given an interface, populates fields from keybase config file where the field implements the `localsecret` struct tag
func LoadLocalSecrets(o interface{}) error {
	// load config file into map, if files do not exist it will be an empty map
	contentMap, err := loadContentMap()
	if err != nil {
		return pkgerr.Wrap(err, "local secrets could not be loaded")
	}
	// no environments available
	if len(contentMap) < 1 {
		return nil
	}

	// kick off recursive replacement
	return setTaggedFields(contentMap, o)
}

func setTaggedFields(contentMap map[string]interface{}, o interface{}) error {
	if o == nil {
		return errors.New("config must not be nil")
	}
	if reflect.ValueOf(o).Kind() != reflect.Ptr {
		return errors.New("config must be a pointer")
	}

	val := reflect.ValueOf(o).Elem()
	typ := val.Type()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)        // struct field
		fieldVal := val.Field(i)     // reflect.Value of the value at this field in the config object
		fieldKind := fieldVal.Kind() // kind of the struct field

		// if value of this field is a struct we recurse
		if fieldKind == reflect.Struct {
			err := setTaggedFields(contentMap, fieldVal.Addr().Interface())
			if err != nil {
				return err
			}
			continue
		}

		key, ok := field.Tag.Lookup("localsecret")
		if !ok {
			continue
		}
		if !fieldVal.CanSet() {
			return fmt.Errorf("config value at field %s cannot be set", field.Name)
		}
		if fieldKind != reflect.String {
			return fmt.Errorf("config value for key %s at field %s must be string", key, field.Name)
		}

		envVal, err := getKey(contentMap, key)
		if err != nil {
			return err
		}

		envValKind := reflect.ValueOf(envVal).Kind()

		if envValKind != reflect.String {
			return fmt.Errorf("key %s at field %s expected string, got %s", key, field.Name, envValKind)
		}

		fieldVal.SetString(envVal.(string))
	}

	return nil
}

// getKey attempts to find a value in the content map using a JSON-style dot-notation key
func getKey(contentMap map[string]interface{}, key string) (interface{}, error) {
	path := strings.Split(key, ".")

	lastSegment := path[len(path)-1]

	segmentVal := contentMap

	for _, segment := range path {
		if segment == lastSegment {
			return segmentVal[segment], nil

		} else if newPath, ok := segmentVal[segment].(map[string]interface{}); !ok {
			return nil, fmt.Errorf("key %s not found at %s", key, segment)
		} else {
			segmentVal = newPath
		}
	}

	return nil, fmt.Errorf("key %s not found", key)
}

// loadContentMap fetches the configuration maps and merges them
func loadContentMap() (map[string]interface{}, error) {
	defaultConf, err := loadFileAsMap("NONPROD")
	if err != nil {
		return nil, err
	}

	prod, err := loadFileAsMap("PROD")
	if err != nil {
		return nil, err
	}
	for env, conf := range prod {
		defaultConf[env] = conf
	}

	techops, err := loadFileAsMap("TECHOPS")
	if err != nil {
		return nil, err
	}
	for env, conf := range techops {
		defaultConf[env] = conf
	}

	return defaultConf, nil
}

// loadFileAsMap attempts to load and parse the config file. If the file does not exist it returns an empty map.
func loadFileAsMap(system string) (contentMap map[string]interface{}, err error) {
	directory, ok := os.LookupEnv(fmt.Sprintf("CUVVA_%s_REPO", system))
	if !ok {
		return map[string]interface{}{}, nil
	}

	file, err := os.Open(fmt.Sprintf("%s/secrets.json", directory))
	if err != nil {
		if os.IsNotExist(err) || strings.Contains(err.Error(), "input/output error") || os.IsPermission(err) {
			return map[string]interface{}{}, nil
		}
		return
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	err = json.Unmarshal(content, &contentMap)

	return
}
