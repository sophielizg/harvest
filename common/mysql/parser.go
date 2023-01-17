package mysql

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strings"

	"github.com/sophielizg/harvest/common/harvest"
)

type ParserService struct {
	Db *sql.DB
}

type ParserAutoIncrementRules harvest.ParserAutoIncrementRules

func (autoIncrementRules *ParserAutoIncrementRules) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return errors.New("incompatible type for ParserAutoIncrementRules")
	}
	return json.Unmarshal(b, &autoIncrementRules)
}

func (autoIncrementRules *ParserAutoIncrementRules) Value() (driver.Value, error) {
	if autoIncrementRules == nil {
		return nil, nil
	}

	b, err := json.Marshal(autoIncrementRules)
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

func (p *ParserService) getParserTypes() (map[string]int, error) {
	rows, err := p.Db.Query("CALL getParserTypes();")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	parserTypes := make(map[string]int)
	for rows.Next() {
		var parserTypeId int
		var typeName string
		err = rows.Scan(&parserTypeId, &typeName)
		if err != nil {
			return parserTypes, err
		}
		parserTypes[typeName] = parserTypeId
	}
	return parserTypes, nil
}

func (p *ParserService) ParserTypes() ([]string, error) {
	parserTypes, err := p.getParserTypes()
	if err != nil {
		return nil, err
	}

	parserTypeNames := make([]string, 0, len(parserTypes))
	for name := range parserTypes {
		parserTypeNames = append(parserTypeNames, name)
	}

	return parserTypeNames, nil
}

func (p *ParserService) Parsers(crawlId int) ([]harvest.Parser, error) {
	rows, err := p.Db.Query("CALL getParsersForCrawl(?);", crawlId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var parsers []harvest.Parser
	for rows.Next() {
		var dummy int
		var jsonPath *string
		var tagStr *string
		var autoIncrementRules *ParserAutoIncrementRules
		var parser harvest.Parser

		err = rows.Scan(&parser.ParserId, &dummy, &parser.CreatedTimestamp, &dummy,
			&parser.Selector, &parser.Attr, &parser.Xpath, &jsonPath,
			&parser.EnqueueCrawlId, &autoIncrementRules, &parser.PageType, &tagStr)
		if err != nil {
			return nil, err
		}

		if jsonPath != nil {
			parser.JsonPath = strings.Split(*jsonPath, ",")
		}

		if tagStr != nil {
			parser.Tags = strings.Split(*tagStr, ",")
		}

		if autoIncrementRules != nil {
			convertedRules := harvest.ParserAutoIncrementRules(*autoIncrementRules)
			parser.AutoIncrementRules = &convertedRules
		}

		parsers = append(parsers, parser)
	}
	return parsers, nil
}

func (p *ParserService) AddParser(crawlId int, parser harvest.ParserFields) (int, error) {
	if parser.PageType == nil {
		return 0, errors.New("No parser type provided")
	}

	parserTypes, err := p.getParserTypes()
	if err != nil {
		return 0, err
	}

	parserTypeId, ok := parserTypes[*parser.PageType]
	if !ok {
		return 0, errors.New("Invalid parser type provided")
	}

	var autoIncrementRules *ParserAutoIncrementRules
	if parser.AutoIncrementRules != nil {
		convertedRules := ParserAutoIncrementRules(*parser.AutoIncrementRules)
		autoIncrementRules = &convertedRules
	}

	var jsonPathStr string
	if parser.JsonPath != nil {
		jsonPathStr = strings.Join(parser.JsonPath, ",")
	}

	rows, err := p.Db.Query("CALL addParser(?, ?, ?, ?, ?, ?, ?, ?);",
		crawlId, parserTypeId, parser.Selector, parser.Attr, parser.Xpath, jsonPathStr,
		parser.EnqueueCrawlId, autoIncrementRules)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var parserId int
		err = rows.Scan(&parserId)
		if err != nil {
			return 0, err
		}
		return parserId, nil
	}
	return 0, errors.New("Record created but no parserId returned")
}

func (p *ParserService) DeleteParser(parserId int) error {
	stmt, err := p.Db.Prepare("CALL deleteParser(?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(parserId)
	return err
}

func (p *ParserService) AddParserTag(parserId int, tag string) error {
	stmt, err := p.Db.Prepare("CALL addParserTag(?, ?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(parserId, tag)
	return err
}

func (p *ParserService) DeleteParserTag(parserId int, tag string) error {
	stmt, err := p.Db.Prepare("CALL deleteParserTag(?, ?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(parserId, tag)
	return err
}
