package main

import (
	"diploma/internal/repository/incident"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	http.Serve()

	r := mux.NewRouter()
	
	//for _, s := range sms.ParseData("simulator/sms.data") {
	//	fmt.Println(s)
	//}
	//ans := mms.ParseData("http://localhost:8383/mms")
	//a, err := json.Marshal(ans)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(a))
	//for _, s := range voicecall.ParseData("simulator/voice.data") {
	//	fmt.Println(s)
	//}
	//for _, s := range email.ParseData("simulator/email.data") {
	//	fmt.Println(s)
	//}
	//k := biling.ParseData("simulator/billing.data")
	//fmt.Println(k)
	//ans := support.ParseData("http://localhost:8383/support")
	//a, err := json.Marshal(ans)
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println(string(a))
	ans := incident.ParseData("http://localhost:8383/accendent")
	a, err := json.Marshal(ans)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(a))
}
