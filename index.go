package main

import "goes/array"

func main()  {
	m := make(map[string]interface{})
	m["test"]="xxx"
	m["test1"]="xxx222"
	m["xxx"]=123
	array.ArrayChangeKeyCase(m,array.CASE_LOWER)
}