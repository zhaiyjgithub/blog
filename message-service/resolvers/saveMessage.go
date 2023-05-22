package resolvers

import (
	"blog/message-service/model"
	"encoding/json"
	"fmt"
)

func SaveMessage(m model.Message) error {
	jb, err := json.Marshal(&m)
	if err != nil {
		fmt.Println("parse body failed", err.Error())
		return err
	}
	fmt.Printf("save message: %s", string(jb))
	return nil
}
