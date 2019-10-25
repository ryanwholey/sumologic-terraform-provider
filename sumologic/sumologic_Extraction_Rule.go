// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Sumo Logic and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------\
package sumologic

import (
	"encoding/json"
	"fmt"
)

func (s *Client) CreateFieldExtractionRule(fieldExtractionRule FieldExtractionRule) (string, error) {
	data, err := s.Post("v1/extractionRules", fieldExtractionRule)
	if err != nil {
		return "", err
	}

	var createdfieldExtractionRule FieldExtractionRule
	err = json.Unmarshal(data, &createdfieldExtractionRule)
	if err != nil {
		return "", err
	}

	return createdfieldExtractionRule.ID, nil
}

func (s *Client) DeleteFieldExtractionRule(id string) error {
	_, err := s.Delete(fmt.Sprintf("v1/extractionRules/%s", id))
	return err
}

func (s *Client) GetFieldExtractionRule(id string) (*FieldExtractionRule, error) {
	data, _, err := s.Get(fmt.Sprintf("v1/extractionRules/%s", id))
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}

	var fieldExtractionRule FieldExtractionRule
	err = json.Unmarshal(data, &fieldExtractionRule)
	if err != nil {
		return nil, err
	}
	return &fieldExtractionRule, nil
}

func (s *Client) UpdateFieldExtractionRule(fieldExtractionRule FieldExtractionRule) error {
	url := fmt.Sprintf("v1/extractionRules/%s", fieldExtractionRule.ID)

	fieldExtractionRule.ID = ""

	_, err := s.Put(url, fieldExtractionRule)
	return err
}

// models
type FieldExtractionRule struct {
	ID string `json:"id,omitempty"`
	// Name of the field extraction rule. Use a name that makes it easy to identify the rule.
	Name string `json:"name"`
	// Scope of the field extraction rule. This could be a sourceCategory, sourceHost, or any other metadata that describes the data you want to extract from. Think of the Scope as the first portion of an ad hoc search, before the first pipe ( | ). You'll use the Scope to run a search against the rule.
	Scope string `json:"scope"`
	// Describes the fields to be parsed.
	ParseExpression string `json:"parseExpression"`
	// Is the field extraction rule enabled.
	Enabled bool `json:"enabled"`
}