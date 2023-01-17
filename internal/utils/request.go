/*
Copyright 2022 Nethermind

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package utils

import (
	"net/http"
	"time"

	"github.com/cenkalti/backoff/v4"
	log "github.com/sirupsen/logrus"
)

/*
GetRequest :
Make a GET request to the given URL. Uses exponential retries with backoff.

params :-
a. url string
URL to make the request to
b. retryDuration time.Duration
Duration to wait between retries

returns :-
a. http.Response
Response from the request
b. error
Error if any
*/
func GetRequest(url string, retryDuration time.Duration) (*http.Response, error) {
	logFields := log.Fields{"Method": "GetRequest"}
	var response *http.Response

	// Adding exponential retry
	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = retryDuration

	err := backoff.Retry(func() (err error) {
		// To make a request with custom headers, use NewRequest and Client.Do.
		response, err = http.Get(url)
		if err != nil {
			log.WithFields(logFields).Errorf("request failed. Error: %v", err)
			log.WithFields(logFields).Info("Retrying request")
			return err
		} else if response.StatusCode != 200 {
			log.WithFields(logFields).Errorf("bad response, got: %d", response.StatusCode)
		}
		return nil
	}, b)
	if err != nil {
		return nil, err
	}

	return response, nil
}
