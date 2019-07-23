package qasircore

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"net/http"
	"net/url"

	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type replacement struct {
	re *regexp.Regexp
	ch string
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

var (
	rExps = []replacement{
		{re: regexp.MustCompile(`[\xC0-\xC6]`), ch: "A"},
		{re: regexp.MustCompile(`[\xE0-\xE6]`), ch: "a"},
		{re: regexp.MustCompile(`[\xC8-\xCB]`), ch: "E"},
		{re: regexp.MustCompile(`[\xE8-\xEB]`), ch: "e"},
		{re: regexp.MustCompile(`[\xCC-\xCF]`), ch: "I"},
		{re: regexp.MustCompile(`[\xEC-\xEF]`), ch: "i"},
		{re: regexp.MustCompile(`[\xD2-\xD6]`), ch: "O"},
		{re: regexp.MustCompile(`[\xF2-\xF6]`), ch: "o"},
		{re: regexp.MustCompile(`[\xD9-\xDC]`), ch: "U"},
		{re: regexp.MustCompile(`[\xF9-\xFC]`), ch: "u"},
		{re: regexp.MustCompile(`[\xC7-\xE7]`), ch: "c"},
		{re: regexp.MustCompile(`[\xD1]`), ch: "N"},
		{re: regexp.MustCompile(`[\xF1]`), ch: "n"},
	}
	spacereg       = regexp.MustCompile(`\s+`)
	noncharreg     = regexp.MustCompile(`[^A-Za-z0-9-]`)
	minusrepeatreg = regexp.MustCompile(`\-{2,}`)
	subdomain      string
)

/**
 * For changing data map[string]interface{} into urlencode
 * @param data map[string]interface{}
 * @return string
 */
func UrlEncodedString(data map[string]interface{}) string {
	var Url *url.URL
	Url, err := url.Parse("")
	if err != nil {
		return ""
	}
	parameters := url.Values{}
	for key, val := range data {
		parameters.Add(key, fmt.Sprint(val))
	}
	Url.RawQuery = parameters.Encode()
	return Url.String()
}

func MathRound(f float64) int {
	if math.Abs(f) < 0.5 {
		return 0
	}
	return int(f + math.Copysign(0.5, f))
}

func HandleEmptyFunction(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST,PUT,GET")
	c.Writer.Header().Set("Accept", "application/json")
	c.JSON(200, map[string]interface{}{})
	return
}

func DateTimeNow() string {
	t := time.Now()
	return t.Format("2006-01-02 03:04:05")
}

func BetweenTwoDates(date string) int {
	t := time.Now()

	then, _ := time.Parse(time.RFC3339, date)
	diff := then.Sub(t)

	return int(diff.Hours() / 24)
}

func FormatDate(date string) string {
	then, _ := time.Parse(time.RFC3339, date)
	return then.Format("02-January-2006")
}

func FormatHours(date string) string {
	then, _ := time.Parse(time.RFC3339, date)
	return then.Format("15:4")
}

func SetSubdomain(c *gin.Context) {
	host := strings.Split(c.Request.Host, ".")
	if len(host) > 1 {
		subdomain = host[0]
	}
}

func GetSubdomain() string {
	return subdomain
}

func Base64ToByte(base64EncodeString string) []byte {
	e64 := base64.StdEncoding
	enc := []byte(base64EncodeString)
	maxDecLen := e64.DecodedLen(len(enc))
	decBuf := make([]byte, maxDecLen)

	n, err := e64.Decode(decBuf, enc)
	_ = err

	return decBuf[0:n]
}

func ToSlug(s string, lower ...bool) string {
	for _, r := range rExps {
		s = r.re.ReplaceAllString(s, r.ch)
	}

	s = strings.ToLower(s)
	s = spacereg.ReplaceAllString(s, "-")
	s = noncharreg.ReplaceAllString(s, "")
	s = minusrepeatreg.ReplaceAllString(s, "-")

	return s
}

func StringToInteger(value string) int {
	val, _ := strconv.Atoi(value)

	return val
}

func ToArrayInteger(listTextID string) []int {
	varID := strings.Split(listTextID, ",")
	arrayInteger := []int{}
	for _, index := range varID {
		if index != "" {
			j, err := strconv.Atoi(index)
			if err == nil {
				arrayInteger = append(arrayInteger, j)
			}
		}
	}
	return arrayInteger
}

/**
 * Change type data of io.ReadCloser to string
 * @param data io.ReadCloser
 * @return string
 */
func ReadCloserToString(data io.ReadCloser) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(data)
	return buf.String()
}

func ConvertingListVariantID(listVariantID []string) string {
	var stringVariantID string

	for _, val := range listVariantID {
		stringVariantID = stringVariantID + "," + val
	}
	return stringVariantID
}

func In_Array(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

/**
 * @brief      Gets the current date.
 * @return     The current date.
 */
func getCurrentDate() string {
	current_time := time.Now().Local()
	return current_time.Format("2006-01-02")
}

func GetExtensionFile(filename string) string {
	return filepath.Ext(filename)
}

func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func StringRandom(length int) string {
	return StringWithCharset(length, charset)
}

func GenerateRandomInt() int {
	rand.Seed(time.Now().UnixNano())
	return Random(1111, 9999)
}

func Random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func TrimQuote(s string) string {
	if len(s) > 0 && s[0] == '"' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '"' {
		s = s[:len(s)-1]
	}
	return s
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetIP(req *http.Request) string {

	// Try Request Headers (X-Forwarder). Client could be behind a Proxy
	ip, err := getClientIPByHeaders(req)
	if err == nil {
		log.Printf("debug: Found IP using Request Headers sniffing. ip: %v", ip)
		return ip
	}

	// Try by Request
	ip, err = getClientIPByRequestRemoteAddr(req)
	if err == nil {
		log.Printf("debug: Found IP using Request sniffing. ip: %v", ip)
		return ip
	}
	return ""
}

// getClientIPByRequest tries to get directly from the Request.
// https://blog.golang.org/context/userip/userip.go
func getClientIPByRequestRemoteAddr(req *http.Request) (ip string, err error) {

	// Try via request
	ip, port, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		log.Printf("debug: Getting req.RemoteAddr %v", err)
		return "", err
	} else {
		log.Printf("debug: With req.RemoteAddr found IP:%v; Port: %v", ip, port)
	}

	userIP := net.ParseIP(ip)
	if userIP == nil {
		message := fmt.Sprintf("debug: Parsing IP from Request.RemoteAddr got nothing.")
		log.Printf(message)
		return "", fmt.Errorf(message)

	}
	log.Printf("debug: Found IP: %v", userIP)
	return userIP.String(), nil

}

// getClientIPByHeaders tries to get directly from the Request Headers.
// This is only way when the client is behind a Proxy.
func getClientIPByHeaders(req *http.Request) (ip string, err error) {

	// Client could be behid a Proxy, so Try Request Headers (X-Forwarder)
	ipSlice := []string{}

	ipSlice = append(ipSlice, req.Header.Get("X-Forwarded-For"))
	ipSlice = append(ipSlice, req.Header.Get("x-forwarded-for"))
	ipSlice = append(ipSlice, req.Header.Get("X-FORWARDED-FOR"))

	for _, v := range ipSlice {
		log.Printf("debug: client request header check gives ip: %v", v)
		if v != "" {
			return v, nil
		}
	}
	err = errors.New("error: Could not find clients IP address from the Request Headers")
	return "", err

}

// getMyInterfaceAddr gets this private network IP. Basically the Servers IP.
func getMyInterfaceAddr() (net.IP, error) {

	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	addresses := []net.IP{}
	for _, iface := range ifaces {

		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			addresses = append(addresses, ip)
		}
	}
	if len(addresses) == 0 {
		return nil, fmt.Errorf("no address Found, net.InterfaceAddrs: %v", addresses)
	}
	//only need first
	return addresses[0], nil
}

func ChangeMobilePhoneFormat(phone string) string {
	var mobile string
	mobilePhone := string(phone[0])

	// check mobile phone
	if mobilePhone == "0" {
		mobile = strings.Replace(phone, "0", "+62", 1)
	} else {
		mobile = phone
	}

	// return new mobile phone
	return mobile
}

func HashHmac(data string, secret string) string {
	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))

	// Write Data to it
	h.Write([]byte(data))

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))

	return sha
}
