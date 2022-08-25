package tools

import (
	"bufio"
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"runtime"
	"strings"
	"time"
	"unicode"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	uuid "github.com/satori/go.uuid"

	"math/rand"
	"os"

	"github.com/theckman/go-flock"
)

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func ByteToPoint(a byte) *byte {
	return &a
}

// LockOrDie - ...
func LockOrDie(ctx context.Context, dir string) *flock.Flock {
	f := flock.New(dir)
	success, err := f.TryLock()
	if err != nil {
		fmt.Printf("Locking AL_HILAL_CORE %s\n", err.Error())
	}

	if !success {
		fmt.Printf("Locking AL_HILAL_CORE %v\n", success)
	}

	return f
}

func RemoveSpaces(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}

// ShuffleSlice - ...
func ShuffleSlice(slice []string) {
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// UUID - ...
func UUID() string {
	return uuid.Must(uuid.NewV4(), nil).String()
}

// UserHomeDir - ...
func UserHomeDir() string {
	env := "HOME"
	if runtime.GOOS == "windows" {
		env = "USERPROFILE"
	} else if runtime.GOOS == "plan9" {
		env = "home"
	}
	return os.Getenv(env)
}

// MakeDirectory makes directory if is not exists
func MakeDirectory(dir string) error {
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			return os.Mkdir(dir, 0775)
		}
		return err
	}
	return nil
}

// StringInSlice - ...
func StringInSlice(slice []string, v string) bool {
	for _, item := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// SHA256 - ...
func SHA256(str string) (result string) {
	h := sha256.New()
	h.Write([]byte(str))
	result = fmt.Sprintf("%x", h.Sum(nil))
	return
}

// ReadFile - ...
func ReadFile(fileName string) (dat []byte, err error) {
	dat, err = ioutil.ReadFile(fileName)
	return
}

// GetCurrentPath - ...
func GetCurrentPath() (currentPath string) {
	if cwd, err := os.Getwd(); err == nil {
		currentPath = cwd
	}
	return
}

func ReadFileLine(fileName string) (lines []string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	file.Close()
	return
}

func StrUniq(s []string) []string {
	m := make(map[string]bool)
	for _, item := range s {
		if _, ok := m[item]; !ok {
			m[item] = true
		}
	}
	var result []string
	for item, _ := range m {
		result = append(result, item)
	}
	return result
}

func IntUniq(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// CopyFile - copy src (path) file into dst (path). Rewrites or creates dst.
func CopyFile(src, dst string) (err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("could not open source file %s: %v", src, err)
	}
	defer func() {
		// Close file properly
		closeErr := srcFile.Close()
		if err == nil && closeErr != nil {
			err = fmt.Errorf("could not close source file %s: %v", src, closeErr)
		}
	}()

	// Try to open destination file
	dstFile, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	//dstFile, err := os.Open(dst)
	if os.IsNotExist(err) {
		// Try to create file
		dstFile, err = os.Create(dst)
		if err != nil {
			return fmt.Errorf("could not create destination file %s: %v", dst, err)
		}
	}
	defer func() {
		// Close file properly
		closeErr := dstFile.Close()
		if err == nil && closeErr != nil {
			err = fmt.Errorf("could not close destination file %s: %v", dst, closeErr)
		}
	}()

	_, err = io.Copy(dstFile, srcFile)

	return
}

func RemoveInt64(elem int64, slice []int64) (result []int64) {
	for i := range slice {
		if slice[i] != elem {
			result = append(result, slice[i])
		}
	}
	return result
}

func RemoveString(elem string, slice []string) (result []string) {
	for i := range slice {
		if slice[i] != elem {
			result = append(result, slice[i])
		}
	}
	return result
}

func SetIfNotZero(field int64) interface{} {
	if field != 0 {
		return field
	} else {
		return nil
	}
}

func SetIfNotEmpty(field string) interface{} {
	if field != "" {
		return field
	} else {
		return nil
	}
}

func MaskCardNumber(number string) (result string) {
	if number == "" {
		return
	}
	if len(number) < 9 {
		return number
	}

	result = number[:6]
	for i := 0; i < len(number)-10; i++ {
		result += "X"
	}

	result += number[len(number)-4:]

	return
}

func ValidateTaxCode(taxCode string) bool {
	return CheckWithRegExp(taxCode, "^[0-9]{12}$")
}

func RemoveDuplicatesTime(slice []time.Time) []time.Time {
	set := make(map[time.Time]struct{})
	var result []time.Time

	for _, t := range slice {
		if _, ok := set[t]; !ok {
			result = append(result, t)
			set[t] = struct{}{}
		}
	}

	return result
}

func RemoveDuplicatesInt64(slice []int64) []int64 {
	set := make(map[int64]struct{})
	var result []int64

	for _, number := range slice {
		if _, ok := set[number]; !ok {
			result = append(result, number)
			set[number] = struct{}{}
		}
	}

	return result
}

func WrapError(desc string, err error) error {
	if err != nil {
		return errors.Wrap(err, desc)
	}
	return nil
}

// TrackTime - used to track execution time of function
// usage:
// defer TrackTime(time.Now(), "get all customer from database", logger.WorkLogger)
func TrackTime(start time.Time, msg string, logger *zap.Logger) {
	logger.Info(msg, zap.String("took", fmt.Sprintf("%s", time.Now().Sub(start))))
}
