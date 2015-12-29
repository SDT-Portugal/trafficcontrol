// Copyright 2015 Comcast Cable Communications Management, LLC

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was initially generated by gen_goto2.go (add link), as a start
// of the Traffic Ops golang data model

package todb

import (
	"encoding/json"
	"fmt"
	"gopkg.in/guregu/null.v3"
	"time"
)

type Log struct {
	Id          int64       `db:"id" json:"id"`
	Level       null.String `db:"level" json:"level"`
	Message     string      `db:"message" json:"message"`
	TmUser      int64       `db:"tm_user" json:"tmUser"`
	Ticketnum   null.String `db:"ticketnum" json:"ticketnum"`
	LastUpdated time.Time   `db:"last_updated" json:"lastUpdated"`
}

func handleLog(method string, id int, payload []byte) (interface{}, error) {
	if method == "GET" {
		ret := []Log{}
		if id >= 0 {
			err := globalDB.Select(&ret, "select * from log where id=$1", id)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
		} else {
			queryStr := "select * from log"
			err := globalDB.Select(&ret, queryStr)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
		}
		return ret, nil
	} else if method == "POST" {
		var v Asn
		err := json.Unmarshal(payload, &v)
		if err != nil {
			fmt.Println(err)
		}
		insertString := "INSERT INTO log("
		insertString += "level"
		insertString += ",message"
		insertString += ",tm_user"
		insertString += ",ticketnum"
		insertString += ") VALUES ("
		insertString += ":level"
		insertString += ",:message"
		insertString += ",:tm_user"
		insertString += ",:ticketnum"
		insertString += ")"
		result, err := globalDB.NamedExec(insertString, v)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		return result.LastInsertId()
	}
	return nil, nil
}
