package string_util

import (
	"strconv"
	"strings"
)

// SPLIT THE STRING FROM THE FILE
func SplitString(target string) string{
	parts := strings.Split(target, " ")
	numStr := parts[0];

	if strings.HasSuffix(numStr, "k"){
		parts = strings.Split(numStr, "k");
		numStr = parts[0];
		
		amount, err := strconv.Atoi(numStr)

		if err != nil{
			return "Error parsing line"
		}

		totalAmount := amount * 1000
		return strconv.Itoa(totalAmount);
	}
	return numStr;
}
