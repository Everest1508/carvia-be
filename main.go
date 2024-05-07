package main

import "fmt"

func main(){
	n:=10
	for i:=0;i<n;i++{
		fmt.Println(space(n-1-i)+star(i+1))
	}
}
func space(j int) string {
	space_str := ""
	for i:=0;i<j;i++{
		space_str+=" "
	}
	return space_str
}

func star(j int) string {
	star_str := ""
	for i:=0;i<j;i++{
		star_str+="* "
	}
	return star_str
}

