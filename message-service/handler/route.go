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
	var err error
	switch payload.ResolverName {
	case "saveMessage":
		var m model.Message
		err = json.Unmarshal([]byte(payload.Body), &m)
		if err != nil {
			fmt.Printf("body: %v", payload.Body)
			resp.StatusCode = 400
			resp.Body = err.Error()
		} else {
			_ = resolvers.SaveMessage(m)
		}
	case "UpdateMesageStatus":
		var p struct {
			OrganizationID string
			CreatedAt      string
			Status model.MessageStatus
		}
		if err = json.Unmarshal([]byte(payload.Body), &p); err != nil {
			resp.StatusCode = 400
			resp.Body = err.Error()
		}else {
			_ = resolvers.UpdateMesageStatus(p.OrganizationID, p.CreatedAt, p.Status)
		}

	default:
		resp.StatusCode = 400
		errText := fmt.Sprintf("%s is not found", payload.ResolverName)
		err = errors.New(errText)
	}
	return resp, err
}

