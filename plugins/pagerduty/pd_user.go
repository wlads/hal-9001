package pagerduty

/*
 * Copyright 2016 Albert P. Tobey <atobey@netflix.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// https://v2.developer.pagerduty.com/v2/page/api-reference#!/On-Calls/get_oncalls

func GetUsersOncall(token string) ([]Oncall, error) {
	out := make([]Oncall, 0)
	offset := 0
	limit := 100

	for {
		url := pagedUrl("/oncalls", offset, limit, nil)

		resp, err := authenticatedGet(url, token)
		if err != nil {
			log.Printf("GET %s failed: %s", url, err)
			return out, err
		}

		data, err := ioutil.ReadAll(resp.Body)

		oresp := OncallsResponse{}
		err = json.Unmarshal(data, &oresp)
		if err != nil {
			log.Printf("json.Unmarshal failed: %s", err)
			return out, err
		}

		out = append(out, oresp.Oncalls...)

		if oresp.More {
			offset = offset + limit
		} else {
			break
		}
	}

	return out, nil
}

func GetUsers(token string, params map[string][]string) ([]User, error) {
	out := make([]User, 0)
	offset := 0
	limit := 100

	for {
		url := pagedUrl("/users", offset, limit, params)

		resp, err := authenticatedGet(url, token)
		if err != nil {
			log.Printf("GET %s failed: %s", url, err)
			return out, err
		}

		data, err := ioutil.ReadAll(resp.Body)

		uresp := UsersResponse{}
		err = json.Unmarshal(data, &uresp)
		if err != nil {
			log.Printf("json.Unmarshal failed: %s", err)
			return out, err
		}

		out = append(out, uresp.Users...)

		if uresp.More {
			offset = offset + limit
		} else {
			break
		}
	}

	return out, nil
}
