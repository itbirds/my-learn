package main

import "net/http"
import "fmt"
import "io/ioutil"

func main (){
	resp,err:= http.Get("http://rmzb.peopleapp.com/api/LoginApi/login?deviceCode=b73240634f1160079273bec264b6874931d55b22&password=ODg4ODg4&userName=lvhaiqiangs&userType=1")
	if err!=nil {
		fmt.Println("can not request",err)
		return
	}
	result,err:=ioutil.ReadAll(resp.Body)
	fmt.Printf("result is %s",result)
	defer resp.Body.Close()
}
