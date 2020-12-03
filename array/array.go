package array

import (
	"fmt"
	"strings"
)
const (
	CASE_UPPER  = 1
	CASE_LOWER  = 0
)





/**
array_change_key_case(array,case);
把数组中所有键更改为小写或大写
 */
func ArrayChangeKeyCase(array map[string]interface{}, case_val int) map[string]interface{} {
	array_len := len(array)
	if array_len < 1 {
		return array
	}
	return_arr := make(map[string]interface{})
	switch case_val {
	case CASE_UPPER:
		for key, value := range array {
			return_arr[strings.ToUpper(key)] = value
		}
	default:
		for key, value := range array {
			return_arr[strings.ToLower(key)] = value
		}
	}
	//fmt.Println(return_arr)
	return return_arr
}

/**
array_column(array,column_key,index_key)
返回输入数组中某个单一列的值。
 */
func ArrayColumn(array map[int]map[string]interface{},column_key string,index_key string) {
	if column_key == "" && index_key == "" {
		//return array
	}
	if column_key != ""{
		//return_arr := make([]string,10)
		for key, value := range array {
			fmt.Printf("addr:%p \t\tlen:%v ",key,value);

			/*if value[column_key] != "" {
				append(return_arr,value)
			}*/
		}
		//return return_arr
	}
	//return array
}
