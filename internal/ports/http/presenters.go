package http

import "github.com/samarec1812/Snakerspace/internal/notes"

type createNoteRequest struct {
	Text string `json:"text"`
}

type updateNoteRequest struct {
	Text string `json:"text"`
}

func NoteSuccessResponse(note notes.Note) map[string]interface{} {
	resp := make(map[string]interface{})
	resp["data"] = note
	resp["error"] = nil

	return resp
}

func ListNoteSuccessResponse(notes []notes.Note) map[string]interface{} {
	resp := make(map[string]interface{})
	resp["data"] = notes
	resp["error"] = nil

	return resp
}

func NoteDeleteResponse() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["data"] = nil
	resp["error"] = nil

	return resp
}

func NoteErrorResponse(err error) map[string]interface{} {
	resp := make(map[string]interface{})
	resp["data"] = nil
	resp["error"] = err.Error()

	return resp
}
