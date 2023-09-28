package igdb

import (
	"db_gen/internal/igdb/query"
	"encoding/json"
	"errors"
	"github.com/biter777/countries"
)

const (
	involvedCompaniesURL = "https://api.igdb.com/v4/involved_companies"
	companiesURL         = "https://api.igdb.com/v4/companies"
)

type involvedCompaniesResponseItem struct {
	ID        int `json:"id"`
	CompanyID int `json:"company"`
}

type involvedCompaniesResponse []involvedCompaniesResponseItem

func (c *Client) involvedCompanies(ids []int) (map[int]int, error) {
	q := query.New().
		Fields("id,company").
		Where(query.Equal("id", query.IDsToString(ids))).Limit(len(ids))

	req, err := c.createRequest(involvedCompaniesURL, q)
	if err != nil {
		return nil, err
	}

	resp, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var respParsed involvedCompaniesResponse
	err = json.NewDecoder(resp.Body).Decode(&respParsed)
	if err != nil {
		return nil, err
	}

	if len(respParsed) == 0 {
		return nil, errors.New("companies not found")
	}

	involvedCompaniesMap := make(map[int]int, len(respParsed))
	for _, g := range respParsed {
		involvedCompaniesMap[g.ID] = g.CompanyID
	}

	return involvedCompaniesMap, nil
}

type companiesResponseItem struct {
	ID      int    `json:"id"`
	Country int    `json:"country"`
	Name    string `json:"name"`
}

type companiesResponse []companiesResponseItem

type companyInfo struct {
	Name    string
	Country string
}

func (c *Client) companies(ids []int) (map[int]companyInfo, error) {
	q := query.New().
		Fields("id,name,country").
		Where(query.Equal("id", query.IDsToString(ids))).Limit(len(ids))

	req, err := c.createRequest(companiesURL, q)
	if err != nil {
		return nil, err
	}

	resp, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var respParsed companiesResponse
	err = json.NewDecoder(resp.Body).Decode(&respParsed)
	if err != nil {
		return nil, err
	}

	if len(respParsed) == 0 {
		return nil, errors.New("company not found")
	}

	companiesMap := make(map[int]companyInfo, len(respParsed))
	for _, g := range respParsed {
		companiesMap[g.ID] = companyInfo{
			Name:    g.Name,
			Country: countries.CountryCode(g.Country).String(),
		}
	}

	return companiesMap, nil
}

func (c *Client) companyCountry(ids []int) (map[int]companyInfo, error) {
	involvedCompaniesMap, err := c.involvedCompanies(ids)
	if err != nil {
		return nil, err
	}

	companiesIDs := make([]int, 0, len(involvedCompaniesMap))
	for _, val := range involvedCompaniesMap {
		companiesIDs = append(companiesIDs, val)
	}

	companiesMap, err := c.companies(companiesIDs)
	if err != nil {
		return nil, err
	}

	companiesInfo := make(map[int]companyInfo, len(companiesMap))
	for key, val := range involvedCompaniesMap {
		companiesInfo[key] = companiesMap[val]
	}

	return companiesInfo, nil
}
