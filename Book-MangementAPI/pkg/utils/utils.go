package utils

import(
	"encoding/json"
	"io/ioutil"
	"net/http"
)



func ParseBody(request *http.Request, x interface{}){

	if body,err:= ioutil.ReadAll(request.Body); err ==nil{
		// Then we will Decode the body 

		if err:= json.Unmarshal([]byte(body),x); err !=nil{
			return 
		}


	}


}










// Its Like The Json decode and encode in Node