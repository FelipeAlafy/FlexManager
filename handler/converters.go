package handler

import (
	"fmt"
	"strconv"
	"strings"
)

func ConvertStringIntoFloat(str string) float64 {
	v, err := strconv.ParseFloat(strings.ReplaceAll(str, ",", "."), 64)
	if err != nil {return 0.0}
	return v
}

func ConvertFloatIntoString(v float64) string {
	s := fmt.Sprintf("%.2f", v)
	return strings.ReplaceAll(s, ".", ",")
}

func GetPreFormatedWhatsappUrl(str string) string {
	nstr := "https://wa.me/"
	if str[0:2] != "55" {
		nstr = fmt.Sprintf("%s55%s", nstr, str)
	} else {
		nstr = fmt.Sprintf("%s%s", nstr, str)
	}
	nstr = strings.ReplaceAll(nstr, " ", "")
	nstr = strings.ReplaceAll(nstr, "-", "")
	
	
	return nstr
}