package main

import (
	"blog/fnService"
	"blog/message-service/resolvers"
	"blog/message-service/shared/model"
	"encoding/json"
	"errors"
	"fmt"
)

func Route(payload fnService.FnRequestPayload) (fnService.FnResponse, error) {
	resp := fnService.FnResponse{
		StatusCode:      200,
		Body:            "",
	}
	fmt.Printf("payload body: %v\r\n", payload.Body)
	var err error
	switch payload.ResolverName {
	case "SaveMessage":
		var m model.Message
		err = json.Unmarshal([]byte(payload.Body), &m)
		if err != nil {
			resp.StatusCode = 400
			resp.Body = err.Error()
		} else {
			if err = resolvers.SaveMessage(m); err != nil {
				resp.StatusCode = 400
				resp.Body = err.Error()
			}
		}
	case "UpdateMessageStatus":
		var p struct {
			OrganizationID string
			CreatedAt      string
			Status model.MessageStatus
		}
		if err = json.Unmarshal([]byte(payload.Body), &p); err != nil {
			resp.StatusCode = 400
			resp.Body = err.Error()
		}else {
			if err = resolvers.UpdateMesageStatus(p.OrganizationID, p.CreatedAt, p.Status); err != nil {
				resp.StatusCode = 400
				resp.Body = err.Error()
			}
		}

	default:
		resp.StatusCode = 400
		errText := fmt.Sprintf("%s is not found", payload.ResolverName)
		err = errors.New(errText)
	}
	return resp, err
}

