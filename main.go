package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {

	// 13.03.2018 14:00:15, 12.03.2018 14:00:15
	//13.03.2018 14:00:15,12.03.2018 14:00:15

	//13.03.2018 14:00:15,12.03.2018 14:00:15
	//12.03.2018 14:00:15,13.03.2018 14:00:15
	fmt.Println("Введите дату в формате дд.мм.гггг чч:мм:сс, дд.мм.гггг чч:мм:сс")
	fmt.Println("сделала Олеся")
	fmt.Println("сделала Олеся")
	fmt.Println("сделала Олеся")
	fmt.Println("сделала Олеся")
	fmt.Println("сделала Олеся")
	inputedTime, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}

	splitedTime := strings.Split(inputedTime, ",")

	time1, err := time.Parse("02.01.2006 15:04:05", strings.TrimRight(splitedTime[0], "\n"))
	time2, err := time.Parse("02.01.2006 15:04:05", strings.TrimRight(splitedTime[1], "\n"))

	fmt.Println(time.Since(time1))
	fmt.Println(time.Since(time2))
	var res time.Duration
	if time1.After(time2) {
		res = time1.Sub(time2)
	} else {
		res = time2.Sub(time1)
	}

	fmt.Println(res.Round(time.Second))
	//
	//now := time.Now()
	//past := now.AddDate(0, 0, -1)
	//future := now.AddDate(0, 0, 1)
	//
	//// func Since(t Time) Duration
	//// вычисляет период между текущим моментом и заданным временем в прошлом
	//fmt.Println(now.Format("2006-01-02 15:04:05"), past.Format("2006-01-02 15:04:05"), future.Format("2006-01-02 15:04:05"))
	//
	//fmt.Println(time.Since(past).Round(time.Second))   // 24h0m0s
	//fmt.Println(time.Until(future).Round(time.Second)) // 24h0m0s

	//dur, err := time.ParseDuration("1h12m3s")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(dur.Round(time.Hour).Hours()) // 1
	//
	////2020-05-15 14:00:00
	////2020-05-15 08:00:00
	//
	//var timeString string
	//
	////fmt.Scan(&timeString)
	//rd := bufio.NewReader(os.Stdin)
	//
	//timeStr, err := rd.ReadString('\n')
	//if err != nil {
	//	panic(err)
	//}
	//
	//timeString = strings.TrimSuffix(timeStr, "\n")
	//const shortForm = "2006-01-02 15:04:05"
	//dt, err := time.Parse(shortForm, timeString)
	////	dt, err := time.Parse("2006-01-02 15:04:05", timeString)
	//if err != nil {
	//	panic(err)
	//}
	//hour := dt.Hour()
	//
	//if hour >= 13 {
	//	dt = dt.Add(time.Hour * 24)
	//}
	//fmt.Println(dt.Format("2006-01-02 15:04:05"))
}

//package main
//
//import (
//	"encoding/json"
//	"fmt"
//)
//
////func readTask() (value1, value2, operation interface{}) {
////
////	//return 5.0, 5.6, "/" //тут играемся с параметрами
////	return 5.0, 5, "%" //тут играемся с параметрами
////	//return 5.0, 5.6, "%" //тут играемся с параметрами
////}
//
////func errMassage(value interface{}) bool {
////
////	if _, ok := value.(float64); !ok {
////
////		fmt.Printf("value=%v: %T ", value, value)
////		os.Exit(0)
////		return false
////	} else {
////		return true
////	}
////
////}
////func errMess(err error) {
////	if err != nil {
////		panic(err)
////	}
//
//func main() {
//	//}
//
//	type Student struct {
//		LastName   string
//		FirstName  string
//		MiddleName string
//		Birthday   string
//		Address    string
//		Phone      string
//		Rating     []uint
//	}
//
//	type nGroup struct {
//		ID       int
//		Number   string
//		Year     int
//		StudentS []Student
//	}
//
//	var (
//		arrRatng []uint
//		arrStd   []Student
//	)
//
//	arrRatng = append(arrRatng, 1)
//	arrRatng = append(arrRatng, 3)
//	arrRatng = append(arrRatng, 2)
//
//	NewSt := Student{
//		LastName:   "Вещий",
//		FirstName:  "Лифон",
//		MiddleName: "Вениаминович",
//		Birthday:   "4апреля1970года",
//		Address:    "632432,г.Тобольск,ул.Киевская,дом6,квартира23",
//		Phone:      "+7(948)709-47-24",
//		Rating:     arrRatng,
//	}
//
//	arrRatng = append(arrRatng, 4)
//
//	NewSt2 := Student{
//
//		LastName:   "Ien",
//		FirstName:  "ccc",
//		MiddleName: "Вениаминович",
//		Birthday:   "4апреля1970года",
//		Address:    "632432,г.Тобольск,ул.Киевская,дом6,квартира23",
//		Phone:      "+7(948)709-47-24",
//		Rating:     arrRatng,
//	}
//
//	arrStd = append(arrStd, NewSt)
//	arrStd = append(arrStd, NewSt2)
//
//	newGroup := nGroup{
//		ID:       134,
//		Number:   "ИЛМ-1274",
//		Year:     2,
//		StudentS: arrStd}
//
//	data, err := json.Marshal(newGroup)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	//data, err := ioutil.ReadAll(os.Stdin)
//	//if err != nil {
//	//	fmt.Println(err)
//	//	return
//	//}
//
//	var (
//		count, countStud int
//		gr               nGroup
//	)
//
//	type Answer struct {
//		Average float32
//	}
//
//	json.Unmarshal(data, &gr)
//
//	for _, student := range gr.StudentS {
//		r := student.Rating
//		countStud++
//		count += len(r)
//
//	}
//
//	myAnswer := Answer{
//		Average: float32(count) / float32(countStud),
//	}
//	dataRez, err := json.MarshalIndent(myAnswer, "", "\t")
//	fmt.Printf("%s", dataRez)
//
//}
//
////var s int
////scanner := bufio.NewScanner(os.Stdin)
//////scanner.
////for scanner.Scan() {
////
////	aInt, err := strconv.Atoi(scanner.Text())
////
////	if err != nil {
////		break
////	}
////	s += aInt
////
////}
////io.WriteString(os.Stdout, strconv.Itoa(s))
//
////	a, err := bufio.NewReader(os.Stdin).ReadLine()
////if err := scanner.Err(); err != nil {
////
////}
//
///*for {
//
//	a, err := bufio.NewReader(os.Stdin).ReadString('\n')
//	a = strings.TrimSuffix(a, "\n")
//	aInt, err := strconv.Atoi(a)
//	if err != nil {
//		break
//	}
//	s += aInt
//}*/
//
////io.WriteString(os.Stdout, strconv.Itoa(s))
//
////func readFile(fileName string) {
////
////	databyte, err := ioutil.ReadFile(fileName)
////
////	errMess(err)
////
////	fmt.Println("Text data inside the file is \n", string(databyte))
////
////}
////func writeFile(fileName string, data []byte) {
////	ioutil.WriteFile(fileName, data, 0600)
////	ioutil.WriteFile(fileName, data, 0600)
////}
//
////fmt.Print()
////specialPrint(10, 4, 5, 8, 7)
////fmt.Println("Welcome to files in Go")
////content := "This is my file"
////
////file, err := os.Create("./myGofile.txt")
////checkNilErr(err)
////
////length, err := io.WriteString(file, content)
////
////checkNilErr(err)
////fmt.Println("length  is ", length)
////
////file.Close()
////readFile("./myGofile.txt")
//
////func specialPrint(arg ...int) error {
////	if _, err := fmt.Print(arg); err != nil {
////		return err
////	}
////	return nil
////}
////
////func readFile(filename string) {
////	dataByte, err := ioutil.ReadFile(filename)
////	checkNilErr(err)
////
////	fmt.Println("Text data inside a file i \n", string(dataByte))
////}
////
////func checkNilErr(err error) {
////	if err != nil {
////		panic(err)
////	}
////}
//
////func ExampleClosure() {
////	fn := func() func(int) int {
////		count := 0
////		return func(i int) int {
////			count++
////			return count * i
////		}
////	}()
////	i := 5
////	fmt.Printf("%T", fn(i))
////
////	//for i := 1; i <= 5; i++ {
////	//	fmt.Println(fn(i))
////	//}
////
////	// Output:
////	// 1
////	// 4
////	// 9
////	// 16
////	// 25
////}
//
////var a float64 = 101 / 2.0
////fmt.Println(uint(a))
////
////fmt.Println(a + float64(uint(a)))
//
////	groupCity := map[int][]string{
////		10:   []string{"Anapa", "Kerch", "Azov"}, // города с населением 10-99 тыс. человек
////		100:  []string{"Krasnodar", "Rostov"},    // города с населением 100-999 тыс. человек
////		1000: []string{"Msk", "Spb"},             // города с населением 1000 тыс. человек и более
////	}
////
////	cityPopulation := map[string]int{
////		"Anapa":     11,
////		"Kerch":     12,
////		"Krasnodar": 500,
////		"Rostov":    800,
////		"Msk":       5000,
////	}
////
////	var isInside bool
////
////	for indP, _ := range cityPopulation {
////		isInside = false
////		for _, valGC := range groupCity[100] {
////
////			if valGC == indP {
////				isInside = true
////			}
////		}
////		if !isInside {
////			delete(cityPopulation, indP)
////		}
////
////	}
////
////	fmt.Print(cityPopulation)
////}
//
////
////var a []int
////var b int
////
////for i := 0; i < 10; i++ {
////	fmt.Scan(&b)
////	a = append(a, b)
////}
////m := make(map[int]int)
////
////for i := 0; i < 10; i++ {
////
////	val, isOk := m[a[i]]
////
////	if isOk {
////		fmt.Println(val)
////	} else {
////		m[a[i]] = work(a[i])
////		fmt.Println(m[a[i]])
////	}
////}
