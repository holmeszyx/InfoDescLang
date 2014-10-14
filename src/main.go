package main

import (
	"os"
	"fmt"
	"./idl"
)

func main(){
	testFilePath := "/home/holmes/IdeaProjects/InfoDescLang/files/api_desc.txt"
	file, err := os.Open(testFilePath)
	if err != nil{
		fmt.Println("can't open file")
		os.Exit(1)
	}
	defer file.Close()

	p := idl.NewSimpleIdlParser(file)
	infos := p.ParseInfomations()

	if len(infos) == 0{
		fmt.Println("nothing parsed")
		os.Exit(0)
	}


	for _, info := range(infos){
		printInformation(info)
	}

}

func printInformation(info *idl.Information){
	if info != nil{
		fmt.Println("info: ", info.Name)
		for _, attr := range(info.Attrs){
			fmt.Println("  - attr : ", attr.GetValue())
		}
	}
}
