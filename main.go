package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Credentials struct {
	SecretAccessKey string
	SessionToken    string
	AccessKeyId     string
}

type GetSessionTokenResponse struct {
	Credentials Credentials
}

func main() {
	rawData, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		panic(err)
	}

	var response GetSessionTokenResponse
	err = json.Unmarshal(rawData, &response)

	if err != nil {
		panic(err)
	}

	fmt.Printf("aws_access_key_id = %s\n", response.Credentials.AccessKeyId)
	fmt.Printf("aws_secret_access_key = %s\n", response.Credentials.SecretAccessKey)
	fmt.Printf("aws_session_token = %s\n", response.Credentials.SessionToken)
}
