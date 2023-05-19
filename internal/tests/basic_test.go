package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNote(t *testing.T) {
	client := getTestClient()

	resp, err := client.createNote("alibaba")
	assert.NoError(t, err)
	assert.Zero(t, resp.Data.ID)
	assert.Equal(t, resp.Data.Text, "alibaba")
}

func TestUpdateNote(t *testing.T) {
	client := getTestClient()

	response, err := client.createNote("hello world")
	assert.NoError(t, err)

	response, err = client.updateNote(0, "привет мир")
	assert.NoError(t, err)
	assert.Equal(t, response.Data.Text, "привет мир")
	assert.Equal(t, response.Data.ID, int64(0))
}

func TestGetNoteByID(t *testing.T) {
	client := getTestClient()

	response, err := client.createNote("hello world")
	assert.NoError(t, err)

	response, err = client.noteByID(0)
	assert.NoError(t, err)
	assert.Equal(t, response.Data.Text, "hello world")
	assert.Equal(t, response.Data.ID, int64(0))

}

func TestDeleteNote(t *testing.T) {
	client := getTestClient()

	_, err := client.createNote("hello world")
	assert.NoError(t, err)

	_, err = client.createNote("привет")
	assert.NoError(t, err)

	_, err = client.createNote("желтый цвет")
	assert.NoError(t, err)

	err = client.deleteNote(2)
	assert.NoError(t, err)

}

func TestListAds(t *testing.T) {
	client := getTestClient()

	_, err := client.createNote("hello world")
	assert.NoError(t, err)

	_, err = client.createNote("привет")
	assert.NoError(t, err)

	_, err = client.createNote("желтый цвет")
	assert.NoError(t, err)

	list, err := client.listNote("")
	assert.NoError(t, err)
	assert.Len(t, list.Data, 3)
}
