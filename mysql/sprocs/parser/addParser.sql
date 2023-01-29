use harvest;

DROP PROCEDURE IF EXISTS addParser;

DELIMITER $$

CREATE PROCEDURE addParser(
    IN scraperIdIn INT,
    IN parserTypeIdIn INT,
    IN selectorIn VARCHAR(255),
    IN attrIn VARCHAR(255),
    IN xpathIn VARCHAR(255),
    IN jsonPathIn VARCHAR(255),
    IN enqueueScraperIdIn INT,
    IN autoIncrementRulesIn JSON
) BEGIN
    INSERT INTO Parser
        (scraperId, createdTimestamp, parserTypeId, 
        selector, attr, xpath, jsonPath, 
        enqueueScraperId, autoIncrementRules)
    VALUES
        (scraperIdIn, NOW(), parserTypeIdIn, 
        selectorIn, attrIn, xpathIn, jsonPathIn, 
        enqueueScraperIdIn, autoIncrementRulesIn);
    SELECT LAST_INSERT_ID() AS parserId;
END $$

DELIMITER ;