package fpp

import (
	"fmt"
	"strings"
)

type Person struct {
	PersonId   string `json:"person_id"`
	PersonName string `json:"person_name"`
	Faces      []Face `json:"face"`
}

func personParams(personId string, useNameInsteadOfId bool) map[string]string {
	paramKey := "person_id"
	if useNameInsteadOfId {
		paramKey = "person_name"
	}
	return map[string]string{paramKey: personId}
}

func GetPersonList() ([]Person, error) {
	var people struct {
		personList []Person `json:"person"`
	}
	err := GetRequest("info/get_person_list", map[string]string{}, &people)
	return people.personList, err
}

func CreatePerson(pName string, faceIds ...string) (string, error) {
	var jsonResp struct {
		personId string `json:"person_id"`
	}
	err := GetRequest("person/create", map[string]string{
		"person_name": pName,
		"face_id":     strings.Join(faceIds, ","),
	}, &jsonResp)
	return jsonResp.personId, err
}

func GetPersonInfo(personId string, useNameInsteadOfId bool) (Person, error) {
	var person Person
	err := GetRequest("person/get_info", personParams(personId, useNameInsteadOfId), &person)
	return person, err
}

func AddFacesToPerson(personId string, useNameInsteadOfId bool, faceIds ...string) error {
	var jsonResp struct {
		success bool `json:"success"`
	}
	params := personParams(personId, useNameInsteadOfId)
	params["face_id"] = strings.Join(faceIds, ",")
	err := GetRequest("person/add_face", params, &jsonResp)
	if err != nil {
		return err
	}
	if !jsonResp.success {
		return fmt.Errorf("Unable to add faces to person %s", personId)
	}
	return nil
}

func RemoveFacesFromPerson(personId string, useNameInsteadOfId bool, faceIds ...string) error {
	var jsonResp struct {
		success bool `json:"success"`
	}
	params := personParams(personId, useNameInsteadOfId)
	params["face_id"] = strings.Join(faceIds, ",")
	err := GetRequest("person/remove_face", params, &jsonResp)
	if err != nil {
		return err
	}
	if !jsonResp.success {
		return fmt.Errorf("Unable to remove faces from person %s", personId)
	}
	return nil
}

func DeletePerson(personId string, useNameInsteadOfId bool) error {
	var jsonResp struct {
		success bool `json:"success"`
	}
	err := GetRequest("person/delete", personParams(personId, useNameInsteadOfId), &jsonResp)
	if err != nil {
		return err
	}
	if !jsonResp.success {
		return fmt.Errorf("Unable to delete person %s", personId)
	}
	return nil
}
