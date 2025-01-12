# GoSteamProto

[Steam Protocol Buffer](https://github.com/SteamDatabase/Protobufs.git) definitions for Go.
Contains compiled protobufs for Steam and its games (CSGO, Dota 2, etc).

## Installation

```bash
go get github.com/neigrok/gosteamproto
```

## Usage Example

Poll auth session status after login attempt:

```go

package main

import (
	"bytes"
	"encoding/base64"
	"github.com/neigrok/gosteamproto/compiled/steam"
	"google.golang.org/protobuf/proto"
	"io"
	"mime/multipart"
	"net/http"
)

func PollAuthSessionStatus(client *http.Client, clientId *uint64, requestId []byte) (*steam.CAuthentication_PollAuthSessionStatus_Response, error) {
	binData, err := proto.Marshal(&steam.CAuthentication_PollAuthSessionStatus_Request{
		ClientId:  clientId,
		RequestId: requestId,
	})
	if err != nil {
		return nil, err
	}

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	writer.WriteField("input_protobuf_encoded", base64.StdEncoding.EncodeToString(binData))
	writer.Close()

	u := "https://api.steampowered.com/IAuthenticationService/PollAuthSessionStatus/v1"
	req, _ := http.NewRequest(http.MethodPost, u, &body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var pollResp steam.CAuthentication_PollAuthSessionStatus_Response
	if err := proto.Unmarshal(respBody, &pollResp); err != nil {
		return nil, err
	}

	return &pollResp, nil
}
```
