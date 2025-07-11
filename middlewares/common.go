package middlewares

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// ฟังก์ชัน parse วันที่ + เวลา + timezone ครบ ๆ
func ParseFlexibleDatetime(datetimeStr string) (time.Time, error) {
	var layouts = []string{
		"2006-01-02T15:04:05-07:00", // มี T + timezone เช่น 2025-04-30T08:59:55+07:00
		"2006-01-02 15:04:05-07:00", // ไม่มี T แต่มี timezone เช่น 2025-04-30 08:59:55+07:00
		"2006-01-02T15:04:05",       // มี T แต่ไม่มี timezone เช่น 2025-04-30T08:59:55
		"2006-01-02 15:04:05",       // ไม่มี T ไม่มี timezone เช่น 2025-04-30 08:59:55
		"2006-01-02",                // มีแค่วัน เช่น 2025-04-30
	}

	for _, layout := range layouts {
		t, err := time.Parse(layout, datetimeStr)
		if err == nil {
			return t, nil
		}
	}

	// ถ้า parse ไม่ได้เลย
	return time.Time{}, fmt.Errorf("unable to parse datetime: %s", datetimeStr)
}

func DistinceStringArray(array []string) []string {
	var unique []string
	for _, v := range array {
		skip := false
		for _, u := range unique {
			if v == u {
				skip = true
				break
			}
		}
		if !skip {
			unique = append(unique, v)
		}
	}
	return unique
}

func DistinceIntegerArray(array []int) []int {
	var unique []int
	for _, v := range array {
		skip := false
		for _, u := range unique {
			if v == u {
				skip = true
				break
			}
		}
		if !skip {
			unique = append(unique, v)
		}
	}
	return unique
}

func DifferenceArrayInt(slice1 []int, slice2 []int) []int {
	var diff []int
	// Loop two times, first to find slice1 ints not in slice2,
	// second loop to find slice2 ints not in slice1
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			// int not found. We add it to return slice
			if !found {
				diff = append(diff, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}

	return diff
}

func DifferenceArrayString(slice1 []string, slice2 []string) []string {
	var diff []string

	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				diff = append(diff, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}

	return diff
}

func GenerateDateTimeCode() string {
	datetime := time.Now()
	hour, min, sec := datetime.Clock()
	dateStr := fmt.Sprintf("%d%02d%02d%02d%02d%02d", datetime.Year(), datetime.Month(), datetime.Day(), hour, min, sec)
	return dateStr
}

// function compare datetime
func CompareDateTime(strDate1, strDate2 string, comType int, checkBefore bool) (bool, error) {
	datetime1, err1 := ParseFlexibleDatetime(strDate1)
	datetime2, err2 := ParseFlexibleDatetime(strDate2)

	if err1 != nil || err2 != nil {
		return false, fmt.Errorf("unable to parse datetime")
	}

	var start_datetime, end_datetime time.Time

	switch comType {
	case 1:
		// เทียบวันที่ + เวลา
		start_datetime, _ = time.Parse("2006-01-02 15:04:05", datetime1.Format("2006-01-02 15:04:05"))
		end_datetime, _ = time.Parse("2006-01-02 15:04:05", datetime2.Format("2006-01-02 15:04:05"))
	case 2:
		// เทียบแค่วันที่
		start_datetime, _ = time.Parse("2006-01-02", datetime1.Format("2006-01-02"))
		end_datetime, _ = time.Parse("2006-01-02", datetime2.Format("2006-01-02"))
	case 3:
		// เทียบแค่เดือน
		start_datetime, _ = time.Parse("2006-01", datetime1.Format("2006-01"))
		end_datetime, _ = time.Parse("2006-01", datetime2.Format("2006-01"))
	case 4:
		// เทียบแค่ปี
		start_datetime, _ = time.Parse("2006", datetime1.Format("2006"))
		end_datetime, _ = time.Parse("2006", datetime2.Format("2006"))
	default:
		return false, fmt.Errorf("invalid compare type")
	}

	if checkBefore {
		return start_datetime.Before(end_datetime), nil
	}
	return start_datetime.After(end_datetime), nil
}

func StringToFloat64(s string) (float64, error) {
	if s == "" {
		return 0, nil // Return 0 if the string is empty
	}
	res, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid float value: %s", s)
	}
	return res, nil
}

func StringToInt64(s string) (int64, error) {
	if s == "" {
		return 0, nil // Return 0 if the string is empty
	}
	res, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid int value: %s", s)
	}
	return res, nil
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func objectPropertyToArray(obj interface{}, propertyName string) ([]interface{}, error) {
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("input must be a struct or a pointer to a struct")
	}

	field := val.FieldByName(propertyName)
	if !field.IsValid() {
		return nil, fmt.Errorf("field '%s' not found", propertyName)
	}

	if field.Kind() != reflect.Slice {
		return nil, fmt.Errorf("field '%s' is not a slice", propertyName)
	}

	arr := make([]interface{}, field.Len())
	for i := 0; i < field.Len(); i++ {
		arr[i] = field.Index(i).Interface()
	}

	return arr, nil
}

func StructToMap(obj any) (map[string]interface{}, error) {
	var result map[string]interface{}
	// แปลง struct → JSON
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	// แปลง JSON → map
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func ConvertToString(val interface{}) string {
	if val == nil {
		return ""
	}

	// handle pointer
	v := reflect.ValueOf(val)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return ""
		}
		return ConvertToString(v.Elem().Interface())
	}

	switch t := val.(type) {
	case string:
		return t
	case int, int8, int16, int32, int64:
		return fmt.Sprintf("%d", t)
	case uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", t)
	case float32, float64:
		return fmt.Sprintf("%.2f", t) // แสดงทศนิยม 2 ตำแหน่ง
	case bool:
		return fmt.Sprintf("%t", t)
	case time.Time:
		return t.Format("2006-01-02 15:04:05")
	default:
		// fallback to fmt
		return fmt.Sprintf("%v", val)
	}
}

// DiffDate calculates the difference between two dates and returns year, month, day
func DiffDate(from, to time.Time) (years, months, days int) {
	// Ensure from is before to
	if from.After(to) {
		from, to = to, from
	}

	// Calculate year and month difference
	years = to.Year() - from.Year()
	months = int(to.Month()) - int(from.Month())
	days = to.Day() - from.Day()

	if days < 0 {
		// borrow days from previous month
		prevMonth := to.AddDate(0, -1, 0)
		days += daysInMonth(prevMonth.Year(), prevMonth.Month())
		months--
	}

	if months < 0 {
		months += 12
		years--
	}

	return
}

// daysInMonth returns number of days in a given month
func daysInMonth(year int, month time.Month) int {
	return time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

func ContainsString(list []string, target string) bool {
	for _, item := range list {
		if item == target {
			return true
		}
	}
	return false
}

func ContainsInt(list []int, target int) bool {
	for _, item := range list {
		if item == target {
			return true
		}
	}
	return false
}

func DecodeJWT(token string) (map[string]interface{}, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid token format")
	}

	// Decode the payload (2nd part)
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("failed to decode payload: %v", err)
	}

	var claims map[string]interface{}
	if err := json.Unmarshal(payload, &claims); err != nil {
		return nil, fmt.Errorf("failed to unmarshal payload: %v", err)
	}

	return claims, nil
}

func GetSubstring(input string, count int, fromStart bool) string {
	if input != "" {
		if len(input) < count {
			return input
		}
		if fromStart {
			return input[:count]
		} else {
			return input[len(input)-count:]
		}
	} else {
		return ""
	}
}

// RandomNumberString generates a random number with `digitCount` digits and returns it as a string
func RandomNumberString(digitCount int) string {
	if digitCount <= 0 {
		return ""
	}

	rand.Seed(time.Now().UnixNano())

	// Ensure first digit is not 0
	firstDigit := rand.Intn(9) + 1
	result := strconv.Itoa(firstDigit)

	for i := 1; i < digitCount; i++ {
		result += strconv.Itoa(rand.Intn(10))
	}

	return result
}
