package tools

// The MIT License (MIT)
// Author: Alexandr (Alex M.A.K.) Mikhailenko
// E-mail: alex-m.a.k@yandex.kz
// Phone: +7-(747)-137-71-54
// Copyright (c) 2021 AL_HILAL_CORE alex-m.a.k@yandex.kz

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/gofiber/fiber/v2"
	"github.com/hashicorp/go-uuid"
	jsoniter "github.com/json-iterator/go"
	"golang.org/x/crypto/bcrypt"
)

var camel = regexp.MustCompile("(^[^A-Z]*|[A-Z]*)([A-Z][^A-Z]+|$)")

// SendExportedFile - sends exported file
func SendExportedFile(c *fiber.Ctx, body io.Reader, filename, format string) error {
	file := fmt.Sprintf("%s_%s.%s", filename, time.Now().Format("2006.01.02"), format)
	c.Set("Content-Description", "File Transfer")
	c.Set("Content-Transfer-Encoding", "binary")
	c.Set("Content-Disposition", "attachment; filename="+file)

	var mimeType string
	switch format {
	case "pdf":
		mimeType = "application/pdf"
	case "csv":
		mimeType = "text/csv"
	case "json":
		mimeType = "application/json"
	case "xlsx":
		mimeType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"

	}
	c.Set(fiber.HeaderContentType, mimeType)

	return c.Status(fiber.StatusOK).SendStream(body)
}

// RouteMatch - check if the current route matches to the list of routes
func RouteMatch(route string, routes []string) bool {
	for _, excludedRoutes := range routes {
		if route == excludedRoutes {
			return true
		}
	}
	return false
}

func JSONDecode(c *fiber.Ctx, v interface{}) error {
	// custom json decoder
	// ex: err := tools.JSONDecode(c, &data)
	return jsoniter.ConfigFastest.Unmarshal(c.Body(), v)
}

func GenerateSubtitle(text string) string {
	if l := utf8.RuneCountInString(text); l > 100 {
		return string([]rune(text)[:100])
	}

	return text
}

// Pricify -
// ex. Pricify(20.000, 2)
// ex. Pricify(20.321, 2)
func Pricify(v float64, decimals int) float64 {
	var pow float64 = 1
	for i := 0; i < decimals; i++ {
		pow *= 10
	}
	if v < 0 {
		return float64(int((v*pow)-0.5)) / pow
	}
	return float64(int((v*pow)+0.5)) / pow
}

func WordHelper(word string, num int32) (resp string) {
	if strings.EqualFold(word, "персон") {
		switch num {
		case 0:
			resp = "персон"
		case 1:
			resp = "персону"
		default:
			resp = "персоны"
		}
	}

	if strings.EqualFold(word, "место") {
		switch num {
		case 1:
			resp = "место"
		default:
			resp = "мест"
		}
	}

	return
}

func HttpGet(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching url %s: %v", url, err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error fetching url %s: %v", url, err)
	}

	return body
}

// Find key in interface (recursively) and return value as interface
func Find(obj interface{}, key string) (interface{}, bool) {

	//if the argument is not a map, ignore it
	mobj, ok := obj.(map[string]interface{})
	if !ok {
		return nil, false
	}

	for k, v := range mobj {
		// key match, return value
		if k == key {
			return v, true
		}

		// if the value is a map, search recursively
		if m, ok := v.(map[string]interface{}); ok {
			if res, ok := Find(m, key); ok {
				return res, true
			}
		}
		// if the value is an array, search recursively
		// from each element
		if va, ok := v.([]interface{}); ok {
			for _, a := range va {
				if res, ok := Find(a, key); ok {
					return res, true
				}
			}
		}
	}

	// element not found
	return nil, false
}

// HashPassword hashes the password using bcrypt
func HashPassword(password string) string {
	hashedPW, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPW)
}

// HashMatchesPassword matches hash with password. Returns true if hash and password match.
func HashMatchesPassword(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func GetFileContentType(out *os.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

func IsContainCyrillic(text string) bool {
	for _, r := range []rune(text) {
		if r > '\u0410' && r < '\u044F' {
			return true
		}
	}
	return false
}

func ReformatTimeOrNull(a time.Time, f string) interface{} {
	return map[bool]interface{}{true: nil, false: a.Format(f)}[a.IsZero()]
}

func BeginningOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 0, -date.Day()+1)
}

func EndOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 1, -date.Day())
}

func DefaultConfigPath() string {
	workDirectory := filepath.Join(func() string {
		if runtime.GOOS == "windows" {
			home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
			if home == "" {
				home = os.Getenv("USERPROFILE")
			}
			return home
		} else if runtime.GOOS == "linux" {
			home := os.Getenv("XDG_CONFIG_HOME")
			if home != "" {
				return home
			}
		}
		return os.Getenv("HOME")
	}(), ".AL_HILAL_CORE")

	_ = MakeDirectory(workDirectory)
	_ = MakeDirectory(filepath.Join(workDirectory, "logs"))
	_ = MakeDirectory(filepath.Join(workDirectory, "config"))
	_ = MakeDirectory(filepath.Join(workDirectory, "config", "site"))
	_ = MakeDirectory(filepath.Join(workDirectory, "plugins"))
	_ = MakeDirectory(filepath.Join(workDirectory, "workspace"))
	_ = MakeDirectory(filepath.Join(workDirectory, "workspace", "user-data"))
	_ = MakeDirectory(filepath.Join(workDirectory, "workspace", "system-data"))
	_ = MakeDirectory(filepath.Join(workDirectory, "workspace", "framework-temp"))

	return workDirectory
}

func NormalizeCamelCase(s string) string {
	var a []string
	for _, sub := range camel.FindAllStringSubmatch(s, -1) {
		if sub[1] != "" {
			a = append(a, sub[1])
		}
		if sub[2] != "" {
			a = append(a, sub[2])
		}
	}
	return strings.Join(a, " ")
}

func IsWeekend(currentDay time.Time) bool {
	return currentDay.Weekday() == time.Sunday || currentDay.Weekday() == time.Saturday
}

// EndOfDay returns time at the end of the day t
func EndOfDay(t time.Time) time.Time {
	y, m, d := t.Year(), t.Month(), t.Day()
	return time.Date(y, m, d, 23, 59, 59, 999999999, t.Location())
}

// StartOfDay returns time at the start of the day t
func StartOfDay(t time.Time) time.Time {
	y, m, d := t.Year(), t.Month(), t.Day()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

func AreDatesEqual(t1, t2 time.Time) bool {
	return t1.Year() == t2.Year() && t1.Month() == t2.Month() && t1.Day() == t2.Day()
}

// StartOfYear - returns time at start of year
// e.g. 22.07.2020 15:02:00 -> 01.01.2020 00:00:00
func StartOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}

// Contains - returns true if "slice" is a slice of type of "elem" and "slice" contains "elem".
// Now supports time.Time and other types
func Contains(elem interface{}, sliceRaw interface{}) bool {
	slice := reflect.ValueOf(sliceRaw)
	if slice.Kind() == reflect.Slice {
		for i := 0; i < slice.Len(); i++ {
			if slice.Index(i).Interface() == elem {
				return true
			}
		}
	}
	return false
}

// PathExists - ...
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// StringsJoin string array join
func StringsJoin(strs ...string) string {
	var str string
	var b bytes.Buffer
	strsLen := len(strs)
	if strsLen == 0 {
		return str
	}
	for i := 0; i < strsLen; i++ {
		b.WriteString(strs[i])
	}
	str = b.String()
	return str

}

// Capitalize Character initials capitalized
func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str) // The following are introduced
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 { // The following are introduced
				vv[i] -= 32 // string's Code Table is 32 bits apart
				upperStr += string(vv[i])
			} else {
				fmt.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

// GetAllFile - ...
func GetAllFile(pathname string, suffix string) (fileSlice []string) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		return
	}

	for _, fi := range rd {
		if fi.IsDir() {
			continue
			//GetAllFile(path.Join(pathname, fi.Name()))
		} else {
			if suffix != "" {
				if strings.HasSuffix(fi.Name(), suffix) {
					fileSlice = append(fileSlice, fi.Name())
				}
			} else {
				fileSlice = append(fileSlice, fi.Name())

			}
		}
	}
	return
}

func IsString(e interface{}, v string) bool {
	return reflect.TypeOf(e).Kind() == reflect.String && e.(string) == v
}

func IsInt(e interface{}, v int) bool {
	return reflect.TypeOf(e).Kind() == reflect.Int && e.(int) == v
}

func Substract(a, b float64) float64 {
	return map[bool]float64{true: b - a, false: a - b}[a < b]
}

func Int64Null(s int64) interface{} {
	return map[bool]interface{}{true: s, false: nil}[s > 0]
}

func Int32Null(s int32) interface{} {
	return map[bool]interface{}{true: s, false: nil}[s > 0]
}

func Float64Null(s float64) interface{} {
	return map[bool]interface{}{true: s, false: nil}[s > 0]
}

func Float32Null(s float32) interface{} {
	return map[bool]interface{}{true: s, false: nil}[s > 0]
}

func StringNull(s string) interface{} {
	return map[bool]interface{}{true: s, false: nil}[len(s) != 0 && s != ""]
}

func TimeNull(a time.Time, formats ...string) interface{} {
	var format string
	if len(formats) == 0 {
		format = "2006-01-02"
	} else {
		format = formats[0]
	}
	return map[bool]interface{}{true: nil, false: a.Format(format)}[a.IsZero()]
}

func TimeFormat(a *time.Time, formats ...string) interface{} {
	var format string
	if len(formats) == 0 {
		format = "2006-01-02"
	} else {
		format = formats[0]
	}
	return map[bool]interface{}{true: nil, false: a.Format(format)}[a.IsZero()]
}

func GetDatePeriod(selectedPeriod string, dateFrom, dateTo time.Time) (from, to time.Time) {
	loc, _ := time.LoadLocation("America/Los_Angeles")
	from, to = dateFrom.In(loc), dateTo.In(loc)
	if from.IsZero() && to.IsZero() {
		from, to = GetSelectedDatePeriod(selectedPeriod, time.Now())
	}
	return StartOfDay(from), EndOfDay(to)
}

func GetSelectedDatePeriod(selectedPeriod string, dateTo time.Time) (from, to time.Time) {
	to = dateTo
	switch strings.ToLower(selectedPeriod) {
	case "all":
		from = time.Time{}
	case "month":
		from = to.AddDate(0, -1, 0)
		break
	case "week":
		from = to.AddDate(0, 0, +7)
		break
	case "today", "day":
		from = to.AddDate(0, 0, 0)
	default:
		from = to.AddDate(+1, 0, 0)
		break
	}
	return
}

func GetRangeDate(selectedPeriod string) (from, to time.Time) {
	to = time.Now()
	switch strings.ToLower(selectedPeriod) {
	case "month":
		from = to.AddDate(0, -1, 0)
		break
	case "week":
		from = to.AddDate(0, 0, -7)
		break
	case "today":
	default:
		from = to.AddDate(0, 0, 1)
		break
	}
	return
}

// Exists until the start page is changed to 1 at the frontend
func NormalizePageOrder(page, size int) int {
	//if page <= 1 {
	//	return 0
	//} else {
	//	return (page - 1) * size
	//}
	return page * size
}
func RemoveHtmlTag(in string) string {
	// regex to match html tag
	const pattern = `(<\/?[a-zA-A]+?[^>]*\/?>)*`
	r := regexp.MustCompile(pattern)
	groups := r.FindAllString(in, -1)
	// should replace long string first
	sort.Slice(groups, func(i, j int) bool {
		return len(groups[i]) > len(groups[j])
	})
	for _, group := range groups {
		if strings.TrimSpace(group) != "" {
			in = strings.ReplaceAll(in, group, "")
		}
	}
	return in
}
func CamelToSnake(s string) string {
	var buff bytes.Buffer
	if len(s) > 0 {
		for i, letter := range s {
			if i == 0 {
				buff.WriteString(strings.ToLower(string(letter)))
				continue
			}
			if unicode.IsUpper(letter) {
				buff.WriteString("_")
				buff.WriteString(strings.ToLower(string(letter)))
			} else {
				buff.WriteRune(letter)
			}
		}
	}
	return buff.String()
}

// Used when in Java we have BigDecimal divide in HALF_UP mode
func RoundHalfUp(x float64, optionalPrec ...int) float64 {
	var prec int
	if len(optionalPrec) > 0 {
		prec = optionalPrec[0]
	} else {
		prec = 2
	}

	xs := strconv.FormatFloat(x, 'f', prec+1, 64)
	xf, _, _ := big.ParseFloat(xs, 10, 256, big.ToNearestEven)
	xf100, _ := new(big.Float).Mul(xf, big.NewFloat(math.Pow10(prec))).Float64()
	return math.Round(xf100) / math.Pow10(prec)
}

func FormatMoney(f float64, prec int) string {
	rate := RoundHalfUp(f, prec)
	return strconv.FormatFloat(rate, 'f', prec, 64)
}

// StructToMap - converts struct into map[string]interface{}.
// Uses ONLY exported (started with capital letter) fields!
// If tag != "" uses tag values as keys, skipping fields with no provided tag.
// If tag == "" uses fields names as keys.
// Returns nil if no interested fields found.
func StructToMap(obj interface{}, tag string) (result map[string]interface{}) {
	if reflect.TypeOf(obj).Kind() != reflect.Struct {
		return
	}
	objVal := reflect.ValueOf(obj)
	result = make(map[string]interface{}, objVal.NumField())

	for i := 0; i < objVal.NumField(); i++ {
		name := objVal.Type().Field(i).Name
		if unicode.IsLower([]rune(name)[0]) {
			continue
		}
		if tag == "" {
			val := objVal.Field(i).Interface()
			result[name] = val
		} else {
			if tagVal := objVal.Type().Field(i).Tag.Get(tag); tagVal != "" {
				val := objVal.Field(i).Interface()
				result[tagVal] = val
			}
		}
	}

	if len(result) == 0 {
		result = nil
	}

	return
}

// NilZeroValues - recursively sets zero values to nil
// for numbers:  0 -> nil
// for strings: "" -> nil
func NilZeroValues(m map[string]interface{}) {
	for key := range m {
		if m[key] == nil {
			continue
		}
		switch reflect.TypeOf(m[key]).Kind() {
		case reflect.Bool:
			continue
		case reflect.Map:
			val, ok := m[key].(map[string]interface{})
			if !ok {
				res, _ := json.Marshal(m[key])
				_ = json.Unmarshal(res, &val)
			}
			NilZeroValues(val)
		default:
			if reflect.ValueOf(m[key]).IsZero() {
				m[key] = nil
			}
		}
	}
}

func NilZeroValuesSlice(s []map[string]interface{}) {
	for i := range s {
		NilZeroValues(s[i])
	}
}

// PreciseAtLeast - formats "number" at least to given "precision".
// If precision of "number" > "precision" then original precision is used.
func PreciseAtLeast(number float64, precision int) string {
	if originalPrecision := DecimalPlaces(number); originalPrecision > precision {
		precision = originalPrecision
	}
	return strconv.FormatFloat(number, 'f', precision, 64)

}

// DecimalPlaces - counts decimal place in "number".
func DecimalPlaces(number float64) (result int) {
	stringNumber := strconv.FormatFloat(number, 'f', -1, 64)
	index := bytes.IndexRune([]byte(stringNumber), '.')
	if index > -1 {
		result = len(stringNumber) - index - 1
	}
	return
}

// OrderByMap - Used to order after concurrency by map
// todo: refactor the parts of code using this func from O(2) to O(1) :)
func OrderByMap(objMap map[int]interface{}) (result []interface{}) {
	result = make([]interface{}, len(objMap))
	for k, v := range objMap {
		result[k] = v
	}
	return
}

func GetQueryArray(ctx *fiber.Ctx, key, sep string) (result []string) {
	defer func() {
		if len(result) == 1 && result[0] == "" {
			result = nil
		}
	}()
	//if strings.Contains(ctx.Query(key), sep) {
	return strings.Split(ctx.Query(key), sep)
	//}
	//return ctx.QueryArray(key)
}

func CharAt(s string, i int) (result string) {
	if s != "" {
		result = string(s[i])
	} else {
		result = ""
	}
	return
}

func CheckWithRegExp(s, pattern string) bool {
	matched, _ := regexp.MatchString(pattern, s)
	return matched
}

func AddBusinessDate(t time.Time, numBusinessDays int) time.Time {
	for i := 0; i < numBusinessDays; i++ {
		t = t.AddDate(0, 0, 1)
		for t.Weekday() == time.Sunday || t.Weekday() == time.Saturday {
			t = t.AddDate(0, 0, 1)
		}
	}
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}

func GetMin(a, b int) int {
	return map[bool]int{
		true:  a,
		false: b,
	}[a < b]
}

func GetMax(a, b int) int {
	return map[bool]int{
		true:  a,
		false: b,
	}[a > b]
}

// StringsToSet - returns unique elements from strings
func StringsToSet(strings []string) (set []string) {
	seen := make(map[string]struct{})
	for i := range strings {
		if _, ok := seen[strings[i]]; !ok {
			seen[strings[i]] = struct{}{}
			set = append(set, strings[i])
		}
	}
	return
}

func StringsToMap(strings []string) (_map map[string]struct{}) {
	_map = make(map[string]struct{})
	for _, str := range strings {
		_map[str] = struct{}{}
	}
	return _map
}

// for ISmsSenderService.SendMessage
func GenerateSourceID() (result string, err error) {
	resultRaw, err := uuid.GenerateUUID()
	if err != nil {
		return
	}
	result = strings.Replace(resultRaw, "-", "", -1)
	return
}

func MonthsBetweenTwoDates(first, second time.Time) int {
	months := 0
	month := first.Month()
	for first.Before(second) {
		first = first.Add(time.Hour * 24)
		nextMonth := first.Month()
		if nextMonth != month {
			months++
		}
		month = nextMonth
	}

	return months
}

func DivideFloat64Safe(a, b float64) float64 {
	defer func() {
		_ = recover()
	}()
	return a / b
}

func MakeRoleList(roleIds ...int64) string {
	roleIdsStringArr := make([]string, 0, len(roleIds))
	for _, roleId := range roleIds {
		roleIdsStringArr = append(roleIdsStringArr, strconv.FormatInt(roleId, 10))
	}
	return "[" + strings.Join(roleIdsStringArr, ",") + "]" // not ", "
}
