package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/samarec1812/Snakerspace/internal/app"
	"github.com/samarec1812/Snakerspace/internal/notes"

	"github.com/gorilla/mux"
)

func createNote(a app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody createNoteRequest
		dec := json.NewDecoder(r.Body)
		if err := dec.Decode(&reqBody); err != nil {
			err = json.NewEncoder(w).Encode(NoteErrorResponse(err))
			if err != nil {
				log.Println(err)
				return
			}
			return
		}

		note, err := a.CreateNote(r.Context(), reqBody.Text)
		if err != nil {
			err = json.NewEncoder(w).Encode(NoteErrorResponse(err))
			if err != nil {
				log.Println(err)
				return
			}
			return
		}

		err = json.NewEncoder(w).Encode(NoteSuccessResponse(note))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func updateNote(a app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		noteID, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err = json.NewEncoder(w).Encode(NoteErrorResponse(err))
			if err != nil {
				log.Println(err)
				return
			}
			return
		}

		var reqBody updateNoteRequest

		dec := json.NewDecoder(r.Body)
		if err := dec.Decode(&reqBody); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err = json.NewEncoder(w).Encode(NoteErrorResponse(err))
			if err != nil {
				log.Println(err)
				return
			}
			return
		}

		note, err := a.UpdateNote(r.Context(), noteID, reqBody.Text)
		if err != nil {
			if errors.Is(err, notes.ErrInvalidParam) {
				w.WriteHeader(http.StatusBadRequest)
				err = json.NewEncoder(w).Encode(NoteErrorResponse(err))
				if err != nil {
					log.Println(err)
					return
				}
				return
			}

			if errors.Is(err, notes.ErrNoteNotFound) {
				w.WriteHeader(http.StatusNotFound)
				err = json.NewEncoder(w).Encode(NoteErrorResponse(err))
				if err != nil {
					log.Println(err)
					return
				}
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(NoteErrorResponse(err))
			if err != nil {
				log.Println(err)
				return
			}
			return
		}

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(NoteSuccessResponse(note))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func getNote(a app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		noteID, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err = json.NewEncoder(w).Encode(NoteErrorResponse(err))
			if err != nil {
				log.Println(err)
				return
			}
			return
		}

		note, err := a.GetNote(r.Context(), noteID)
		if err != nil {
			if errors.Is(err, notes.ErrNoteNotFound) {
				w.WriteHeader(http.StatusNotFound)
				err = json.NewEncoder(w).Encode(NoteErrorResponse(err))
				if err != nil {
					log.Println(err)
					return
				}
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(NoteErrorResponse(err))
			if err != nil {
				log.Println(err)
				return
			}
			return
		}

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(NoteSuccessResponse(note))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func getAllNote(a app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		param := r.URL.Query().Get("order_by")
		list, err := a.GetAllNote(r.Context(), param)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(NoteErrorResponse(err))
			if err != nil {
				log.Println(err)
				return
			}
			return
		}

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(ListNoteSuccessResponse(list))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func deleteNote(a app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		noteID, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err = json.NewEncoder(w).Encode(NoteErrorResponse(err))
			if err != nil {
				log.Println(err)
				return
			}
			return
		}

		err = a.DeleteNote(r.Context(), noteID)
		if err != nil {
			if errors.Is(err, notes.ErrNoteNotFound) {
				w.WriteHeader(http.StatusOK)
				err = json.NewEncoder(w).Encode(NoteDeleteResponse())
				if err != nil {
					log.Println(err)
					return
				}
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(NoteErrorResponse(err))
			if err != nil {
				log.Println(err)
				return
			}
			return
		}

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(NoteDeleteResponse())
		if err != nil {
			log.Println(err)
			return
		}
	}
}
