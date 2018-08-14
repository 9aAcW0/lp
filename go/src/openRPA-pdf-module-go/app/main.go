package main

import (
	"os/exec"
	"fmt"
	"bytes"
	"io/ioutil"
	"os"
	"encoding/base64"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	//r := gin.Default()
	//
	//routes.InitRoutes(r)
	//
	//r.Run(":9000") // listen and serve on 0.0.0.0:9000

	pdfNames := "test.pdf"
	pdfSecoundsNames := "A" + pdfNames
	pdfPasswords := "abcde"

	//pdfRequests := make([]modules.MailRequest, 1)

	//for i := range pdfRequests {
	//	pdfRequest := pdfRequests[i]
	//	pdfRequest = modules.MailRequest{
	//		FileName:pdfNames,
	//		FileContent:pdfContents[i],
	//		FilePassword:pdfPasswords[i],
	//
	//	}
	cd := "../"
	if err := exec.Command("cd %s",cd).Run();err != err{
		fmt.Print(err)
	}

	fmt.Print(pdfNames  + "\n")
	fmt.Print(pdfSecoundsNames  + "\n")

	fmt.Print(pdfPasswords + "\n")
	fmt.Print("pdftk",pdfNames,"output",pdfNames,"user_pw",pdfPasswords + "\n")



	cmd := exec.Command("pdftk", pdfNames, "output", pdfSecoundsNames, "user_pw", pdfPasswords)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Println("Result: " + out.String())
	Encode(pdfSecoundsNames)
	}


func Encode(fileName string)  {
	fmt.Print(fileName)
	f, err := os.Open("./" + fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b, _ := ioutil.ReadAll(f)

	out, _ := os.Create("out")
	base64.NewEncoder(base64.StdEncoding, out).Write(b)
	spew.Dump(string(b))

}


