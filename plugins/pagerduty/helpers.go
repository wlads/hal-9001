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
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// AuthenticatedGet authenticates with the provided token and GETs the url.
func authenticatedGet(geturl, token string) (*http.Response, error) {
	tokenHdr := fmt.Sprintf("Token token=%s", token)

	req, err := http.NewRequest("GET", geturl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/vnd.pagerduty+json;version=2")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", tokenHdr)

	client := &http.Client{}
	r, err := client.Do(req)

	log.Printf("pagerduty.authenticatedGet(%s, token) = %d", geturl, r.StatusCode)

	return r, err
}

// AuthenticatedPost authenticates with the provided token and posts the
// provided body.
func authenticatedPost(token string, body []byte) (*http.Response, error) {
	tokenHdr := fmt.Sprintf("Token token=%s", token)
	buf := bytes.NewBuffer(body)

	// TODO: make Endpoint a url parameter
	req, err := http.NewRequest("POST", Endpoint, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/vnd.pagerduty+json;version=2")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", tokenHdr)

	client := &http.Client{}
	return client.Do(req)
}

func pagedUrl(resource string, offset, limit int, params map[string][]string) string {
	out := fmt.Sprintf("https://api.pagerduty.com%s", resource)

	query := make([]string, 0)

	if limit != 0 {
		query = append(query, fmt.Sprintf("limit=%d", limit))
	}

	if offset != 0 {
		query = append(query, fmt.Sprintf("offset=%d", offset))
	}

	if params != nil {
		for k, vlist := range params {
			for _, vv := range vlist {
				query = append(query, fmt.Sprintf("%s=%s", k, url.QueryEscape(vv)))
			}
		}
	}

	if len(query) > 0 {
		return fmt.Sprintf("%s?%s", out, strings.Join(query, "&"))
	}

	return out
}
