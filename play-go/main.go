/*
go standard package
	builtin V
	fmt V
	errors V
	strings V
	strconv V
	math
	time V
*/
package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// b, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	// if err != nil {
	// 	log.Fatal("error:", err)
	// }

	// fmt.Println("pass:", string(b))
	if err := bcrypt.CompareHashAndPassword([]byte("$2a$10$6k2X4o6DVR7GUc8jaTiAj.74ysJk/7nhk0i9iZ9nKtogbuyCyvFfWSELECT"), []byte("123456")); err != nil {
		log.Fatal("compare failed:", err)
	}

	// timeUnmarshalTest()
	// reflectPlayground()
}

func updateDBTest() {
	f, err := os.Open("./update.sql")
	if err != nil {
		log.Fatal(err)
	}

	info, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, info.Size())
	_, err = f.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	b := new(strings.Builder)
	b.Write(data)
	fmt.Println(b.String())
}

func checkSumTest(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x", h.Sum(nil))
}

func timeUnmarshalTest() {
	t := PgTime{
		Time: time.Now(),
	}

	b, err := t.MarshalJSON()
	if err != nil {
		log.Fatal("time marshal json error:", err)
	}

	str := new(strings.Builder)
	str.Write(b)
	println(str.String())

	if err = t.UnmarshalJSON(b); err != nil {
		println("time unmarshal json error:", err)
	}

	println("finally:", t.String())
}
