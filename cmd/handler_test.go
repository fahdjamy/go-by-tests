package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func FuzzCalculateHighestHandler(f *testing.F) {

	srv := httptest.NewServer(http.HandlerFunc(CalculateHighestHandler))

	defer srv.Close()

	testCases := []ValuesRequest{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 455, 90}},
		{[]int{-1, 90902, 3, 94, -5, 6, 7, 8, 32455, -9990}},
	}

	for _, tc := range testCases {
		data, _ := json.Marshal(tc)
		f.Add(data)
	}

	f.Fuzz(func(t *testing.T, generatedData []byte) {
		// validate fuzzy generated data
		if !json.Valid(generatedData) {
			t.Errorf("Generated data is not valid JSON")
			t.Skip()
		}

		valReq := ValuesRequest{}
		err := json.Unmarshal(generatedData, &valReq)
		if err != nil {
			t.Errorf("Cannot unmarshal generated data: %v", err)
			t.Skip()
		}

		resp, err := http.DefaultClient.Post(srv.URL, "application/json", bytes.NewBuffer(generatedData))
		if err != nil {
			t.Errorf("error connecting to the server, err: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("want status code %d, got %d", 200, resp.StatusCode)
		}
		var response int
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			t.Errorf("error decoding response body: %v", err)
		}
	})

	// add seed co
	// f.Add("test", "902993")
}
