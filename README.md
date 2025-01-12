# GoSteamProto

[Steam Protocol Buffer](https://github.com/SteamDatabase/Protobufs.git) definitions for Go.
Contains compiled protobufs for Steam and its games (CSGO, Dota 2, etc).

## Installation

```bash
go get github.com/neigrok/gosteamproto
```

## Usage Example

```go
package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/neigrok/gosteamproto/compiled/steam"
	"google.golang.org/protobuf/proto"
	"io"
	"mime/multipart"
	"net/http"
)

func GetAuthSessionStatus(client *http.Client, clientId *uint64, requestId []byte) (*steam.CAuthentication_PollAuthSessionStatus_Response, error) {
	// Prepare request
	pollReq := &steam.CAuthentication_PollAuthSessionStatus_Request{
		ClientId:  clientId,
		RequestId: requestId,
	}

	binData, err := proto.Marshal(pollReq)
	if err != nil {
		return nil, fmt.Errorf("marshal error: %v", err)
	}

	// Build multipart form
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	writer.WriteField("input_protobuf_encoded", base64.StdEncoding.EncodeToString(binData))
	writer.Close()

	// Send request
	req, _ := http.NewRequest("POST", "https://api.steampowered.com/IAuthenticationService/PollAuthSessionStatus/v1", &body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading PollAuthSessionStatus response body: %v", err)
	}

	// Parse response
	var pollResp steam.CAuthentication_PollAuthSessionStatus_Response
	if err := proto.Unmarshal(respBody, &pollResp); err != nil {
		return nil, fmt.Errorf("unmarshal error: %v", err)
	}

	return &pollResp, nil
}
```
