package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type Myevent struct {
	Name string `json:"What is your name?"`
	Age int `json:"How old are you?"`
}

type MyResponse struct{
	Message string `json:"Answer:"`


}

func HandleLambdaEvent(event Myevent) (MyResponse, error) {
	return MyResponse{Message: fmt.Sprintf("Hello %s, you are %d years old", event.Name, event.Age)}, nil
}

func main(){
	lambda.Start(HandleLambdaEvent)
}