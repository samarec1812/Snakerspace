package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"time"

	"github.com/samarec1812/Snakerspace/internal/adapters/noterepo"
	"github.com/samarec1812/Snakerspace/internal/app"
	_http "github.com/samarec1812/Snakerspace/internal/ports/http"
)

type noteData struct {
	ID        int64     `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}

type noteResponse struct {
	Data noteData `json:"data"`
}

type notesResponse struct {
	Data []noteData `json:"data"`
}

var (
	ErrBadRequest = fmt.Errorf("bad request")
	ErrForbidden  = fmt.Errorf("forbidden")
)

type testClient struct {
	client  *http.Client
	baseURL string
}

func getTestClient() *testClient {
	server := _http.NewHTTPServer(":18080", app.NewApp(noterepo.New()))
	testServer := httptest.NewServer(server.Handler)

	return &testClient{
		client:  testServer.Client(),
		baseURL: testServer.URL,
	}
}

func (tc *testClient) getResponse(req *http.Request, out any) error {
	resp, err := tc.client.Do(req)
	if err != nil {
		return fmt.Errorf("unexpected error: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusBadRequest {
			return ErrBadRequest
		}
		if resp.StatusCode == http.StatusForbidden {
			return ErrForbidden
		}
		return fmt.Errorf("unexpected status code: %s", resp.Status)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable to read response: %w", err)
	}

	err = json.Unmarshal(respBody, out)
	if err != nil {
		return fmt.Errorf("unable to unmarshal: %w", err)
	}

	return nil
}

func (tc *testClient) createNote(text string) (noteResponse, error) {
	body := map[string]any{
		"text": text,
	}

	data, err := json.Marshal(body)
	if err != nil {
		return noteResponse{}, fmt.Errorf("unable to marshal: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, tc.baseURL+"/note", bytes.NewReader(data))
	if err != nil {
		return noteResponse{}, fmt.Errorf("unable to create request: %w", err)
	}

	req.Header.Add("Content-Type", "application/json")

	var response noteResponse
	err = tc.getResponse(req, &response)
	if err != nil {
		return noteResponse{}, err
	}

	return response, nil
}

func (tc *testClient) updateNote(noteID int64, text string) (noteResponse, error) {
	body := map[string]any{
		"text": text,
	}
	data, err := json.Marshal(body)
	if err != nil {
		return noteResponse{}, fmt.Errorf("unable to marshal: %w", err)
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf(tc.baseURL+"/note/%d", noteID), bytes.NewReader(data))
	if err != nil {
		return noteResponse{}, fmt.Errorf("unable to create request: %w", err)
	}

	req.Header.Add("Content-Type", "application/json")

	var response noteResponse
	err = tc.getResponse(req, &response)
	if err != nil {
		return noteResponse{}, err
	}

	return response, nil
}

func (tc *testClient) listNote(param string) (notesResponse, error) {
	params := url.Values{}
	if param != "" {
		params.Add("order_by", param)
	}

	req, err := http.NewRequest(http.MethodGet, tc.baseURL+"/note"+params.Encode(), nil)
	if err != nil {
		return notesResponse{}, fmt.Errorf("unable to create request: %w", err)
	}

	var response notesResponse
	err = tc.getResponse(req, &response)
	if err != nil {
		return notesResponse{}, err
	}

	return response, nil
}

func (tc *testClient) noteByID(noteID int64) (noteResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(tc.baseURL+"/note/%d", noteID), nil)
	if err != nil {
		return noteResponse{}, fmt.Errorf("unable to create request: %w", err)
	}

	var response noteResponse
	err = tc.getResponse(req, &response)
	if err != nil {
		return noteResponse{}, err
	}

	return response, nil
}

func (tc *testClient) deleteNote(noteID int64) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf(tc.baseURL+"/note/%d", noteID), nil)
	if err != nil {
		return fmt.Errorf("unable to create request: %w", err)
	}

	var response noteResponse
	err = tc.getResponse(req, &response)
	if err != nil {
		return err
	}

	return nil
}
