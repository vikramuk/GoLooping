package main

import (
	"fmt"
)

func main() {
	baseurl := "https://white.lightmetrics.co/v2/tsp/devices?limit=10&offset="
	Offset := []string{"10", "20", "30"}
	for _, Turl := range Offset {
		url := baseurl + Turl
		fmt.Println(string(url))
	}

}


package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

func main() {	  
	baseurl := "https://AABBCC.com/v2/devices?limit=10&offset="
	method := "GET"
	Offset := []string{"10", "20", "30"}
	for _, Turl := range Offset {
	  url := baseurl + Turl	  
	  client := &http.Client {
	  }
	  req, err := http.NewRequest(method, url, nil)
	  if err != nil {
		fmt.Println(err)
		return
	  }
	  req.Header.Add("authority", "white.lightmetrics.co")
	  req.Header.Add("acc_name", "local")
	  req.Header.Add("accept", "application/json, text/plain, */*")
	  req.Header.Add("x-access-token", "")
	  req.Header.Add("Cookie", "AWSALB=")
	  res, err := client.Do(req)
	  if err != nil {
		fmt.Println(err)
		return
	  }
	  defer res.Body.Close()
	  body, err := ioutil.ReadAll(res.Body)
	  if err != nil {
		fmt.Println(err)
		return
	  }
	  fmt.Println(string(body))
	}
}
