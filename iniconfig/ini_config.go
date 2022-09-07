package iniconfig

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

func MarshalFile(data interface{}, filneme string) (err error) {
	result, err := Marshal(data)
	if err != nil {
		return
	}
	ioutil.WriteFile(filneme, result, 0755)
	return
}

func UnMarshalFile(filename string, result interface{}) (err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	return UnMarshal(data, result)
}

func Marshal(data interface{}) (result []byte, err error) {
	typeinfo := reflect.TypeOf(data)
	if typeinfo.Kind() != reflect.Struct {
		err = errors.New("please pass stuct")
		return
	}
	var conf []string

	valueInfo := reflect.ValueOf(data)
	for i := 0; i < typeinfo.NumField(); i++ {
		sectionField := typeinfo.Field(i)
		sectionVal := valueInfo.Field(i)

		fieldType := sectionField.Type
		if fieldType.Kind() != reflect.Struct {
			continue
		}
		tagVal := sectionField.Tag.Get("ini")
		if len(tagVal) == 0 {
			tagVal = sectionField.Name
		}
		section := fmt.Sprintf("\n[%s]\n", tagVal)
		conf = append(conf, section)

		for j := 0; j < fieldType.NumField(); j++ {
			keyField := fieldType.Field(j)
			fieldTagVal := keyField.Tag.Get("ini")
			if len(fieldTagVal) == 0 {
				fieldTagVal = keyField.Name
			}
			valField := sectionVal.Field(j)
			item := fmt.Sprintf("%s=%v\n", fieldTagVal, valField.Interface())
			// fmt.Println(item)
			conf = append(conf, item)
		}
	}

	for _, val := range conf {
		byteVal := []byte(val)
		result = append(result, byteVal...)
		// fmt.Println(val)
	}

	return
}

func UnMarshal(data []byte, result interface{}) (err error) {
	lineArr := strings.Split(string(data), "\n")
	typeInfo := reflect.TypeOf(result)
	if typeInfo.Kind() != reflect.Ptr {
		err = errors.New("please pass address")
		return
	}

	typeStruct := typeInfo.Elem()
	if typeStruct.Kind() != reflect.Struct {
		err = errors.New("please pass stuct")
		return
	}
	var lastFieldName string
	for index, line := range lineArr {

		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		// 如果是注释，直接忽略
		if line[0] == ';' || line[0] == '#' {
			continue
		}

		if line[0] == '[' {
			lastFieldName, err = parseSection(line, typeStruct)
			if err != nil {
				err = fmt.Errorf("%v lineno:%d", err, index+1)
				return
			}
			// fmt.Printf("selectionName:%s\n", lastFieldName)
			continue
		}
		err = parserItem(lastFieldName, line, result)
		if err != nil {
			err = fmt.Errorf("%v lineno:%d", err, index+1)
			return
		}
	}
	return
}

func parserItem(lastFieldName string, line string, result interface{}) (err error) {
	index := strings.Index(line, "=")
	if index == -1 {
		err = fmt.Errorf("sytax error, line:%s", line)
		return
	}
	key := strings.TrimSpace(line[0:index])
	val := strings.TrimSpace(line[index+1:])
	if len(key) == 0 {
		err = fmt.Errorf("syntax error line:%s", line)
		return
	}
	resultValue := reflect.ValueOf(result)
	selectionValue := resultValue.Elem().FieldByName(lastFieldName)

	selectionType := selectionValue.Type()
	if selectionType.Kind() != reflect.Struct {
		err = fmt.Errorf("field:%s must be struct", lastFieldName)
		return
	}

	keyFiledName := ""
	for i := 0; i < selectionType.NumField(); i++ {
		field := selectionType.Field(i)
		tagVal := field.Tag.Get("ini")
		if tagVal == key {
			keyFiledName = field.Name
			break
		}
	}
	if len(keyFiledName) == 0 {
		return
	}
	fieldValue := selectionValue.FieldByName(keyFiledName)
	if fieldValue == reflect.ValueOf(nil) {
		return
	}
	switch fieldValue.Type().Kind() {
	case reflect.String:
		fieldValue.SetString(val)
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		intval, errRet := strconv.ParseInt(val, 10, 64)
		if errRet != nil {
			err = errRet
			return
		}
		fieldValue.SetInt(intval)
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		unitval, errRet := strconv.ParseUint(val, 10, 64)
		if errRet != nil {
			err = errRet
			return
		}
		fieldValue.SetUint(unitval)
	case reflect.Float32, reflect.Float64:
		floatval, errRet := strconv.ParseFloat(val, 64)
		if errRet != nil {
			err = errRet
			return
		}
		fieldValue.SetFloat(floatval)
	default:
		err = fmt.Errorf("unsupport type: %d", fieldValue.Type().Kind())
	}

	return
}

func parseSection(line string, typeInfo reflect.Type) (fieldName string, err error) {
	if line[0] == '[' && len(line) <= 2 {
		err = fmt.Errorf("syntax error, invalid section:%s", line)
		return
	}
	if line[0] == '[' && line[len(line)-1] != ']' {
		err = fmt.Errorf("syntax error, invalid section:%s", line)
		return
	}
	if line[0] == '[' && line[len(line)-1] == ']' {
		sectionName := strings.TrimSpace(line[1 : len(line)-1])
		if len(sectionName) == 0 {
			err = fmt.Errorf("syntax error, invalid section:%s", line)
			return
		}

		for i := 0; i < typeInfo.NumField(); i++ {
			field := typeInfo.Field(i)
			tagValue := field.Tag.Get("ini")
			if tagValue == sectionName {
				fieldName = field.Name
				break
			}
		}
	}
	return
}
