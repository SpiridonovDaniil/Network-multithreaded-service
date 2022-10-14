package main

import (
	"diploma/internal/repository"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleConnection(w http.ResponseWriter, r *http.Request) {

	resultt, err := json.Marshal(repository.Resultt())
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(resultt)
	if err != nil {
		log.Println(err)
	}

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HandleConnection)

	err := http.ListenAndServe(":8282", r)
	if err != nil {
		log.Println(err)
	}

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
	//z := incident.ParseData("http://localhost:8383/accendent")
	//ap, err := json.Marshal(z)
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println(string(ap))
}
