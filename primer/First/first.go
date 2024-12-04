package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const EnglishHelloPrefix = "Hello, "
const FrenchHelloPrefix = "Hello, "
const SpanishHelloPrefix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World!"
	}

	return getPrefix(language) + name
}

func getPrefix(language string) string {
	prefix := EnglishHelloPrefix
	switch language {
	case "french":
		prefix = FrenchHelloPrefix
	case "spanish":
		prefix = SpanishHelloPrefix
	}
	return prefix
}
func appendTest() {
	s := make([]string, 20, 21)
	s[0] = "Hello"
	s[1] = "World"
	s[2] = "Hola"
	s = append(s, "zzh")
	fmt.Println(s)
}
func stringTest() {
	var AString string
	AString = "fucking, Go!"
	var BString = "啊？？？"
	CString := "哈哈哈哈" + BString
	IntNumber, _ := strconv.Atoi(CString)
	println(BString, AString, CString, IntNumber)

}
func mapTest() {
	m := make(map[string]int)
	m["其心智也"] = 1
	m["心声"] = 2
	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(m["心声"])
	value, ok := m["心声"]
	fmt.Println(value, ok)
	delete(m, "其心智也")
	fmt.Println(m)
	v, o := m["真心"]
	fmt.Println(v, o)
}
func rangeTest() {
	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for i, num := range nums {
		sum += num
		if num%2 == 0 {
			fmt.Println("index:", i, "num:", num)
		}
	}
	fmt.Println("sum:", sum)

	m := map[string]string{"a": "A", "b": "B"}
	for k, v := range m {
		fmt.Println("k:", k, "v:", v)
	}
	for v := range m {
		fmt.Println("v:", v)
	}
}

func funcTest() (v string, yes bool, no bool) {
	return "stirng", true, false
}

type user struct {
	Name     string
	password string
}

func (u *user) resetPassword(password string) {
	u.password = password
}
func (u user) checkPassword(password string) bool {
	return u.password == password
}
func StructPointTest() {
	u := user{"zzh", "1024"}
	u.resetPassword("1111")
	fmt.Println(u.checkPassword("1024"))
}
func findUser(users []user, name string) (v *user, err error) {
	for _, u := range users {
		if u.Name == name {
			return &u, nil
		}
	}
	return nil, errors.New("user not found")
}
func findUserTest() {
	u, err := findUser([]user{{"zzh", "1024"}}, "yzr")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u)
}
func stringFuncTest() {
	a := "i love harry forever"
	fmt.Println(strings.Contains(a, "rr"))
	fmt.Println(strings.Index(a, "rr"))
	fmt.Println(strings.HasPrefix(a, "i "))
	fmt.Println(strings.HasSuffix(a, "ver"))
	fmt.Println(strings.Join([]string{"a", "b", "c"}, "-"))
	fmt.Println(strings.Replace(a, "r", "R", 2))
	fmt.Println(strings.Split("a-b-c-", "-"))
}
func formatString() {
	u := user{"zzh", "1024"}
	fmt.Printf("user = %v\n", u)
	fmt.Printf("user = %+v\n", u)
	fmt.Printf("user = %#v\n", u)
}
func 序列化() {
	users := []user{{"zzh", "1024"}, {"zzh", "1024"}}
	buf, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(buf))
}

type userInfo struct {
	Name  string   `json:"name"`
	Age   int      `json:"age"`
	Hobby []string `json:"hobby"`
}

func UnMarshalTest() {
	user := userInfo{Name: "zzh", Age: 20, Hobby: []string{"Golang", "Java"}}
	buf, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
	fmt.Println(buf)
	fmt.Println(string(buf))

	buf, err = json.MarshalIndent(user, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))

	var b userInfo
	err = json.Unmarshal(buf, &b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", b)
}
func strconvTest() {
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	n, _ := strconv.ParseInt("123", 10, 16)
	fmt.Println(n)

	n2, _ := strconv.Atoi("12334323")
	fmt.Println(n2)
	n3, _ := strconv.Atoi("AAA")
	fmt.Println(n3)
}
func ProcessInfoTest() {
	fmt.Println(os.Args)
	fmt.Println(os.Getenv("PATH"))
	buf, err := exec.Command("netstat").CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))
}

func HttpServerTest() {
	http.Handle("/", http.FileServer(http.Dir("."))) //路由处理
	http.ListenAndServe(":8080", nil)                //服务器监听与启动
}
func main() {
	//UnMarshalTest()
	HttpServerTest()
}
