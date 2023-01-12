use harvest;

DROP PROCEDURE IF EXISTS addParser;

DELIMITER $$

CREATE PROCEDURE addParser(
    IN crawlIdIn INT,
    IN typeIdIn INT,
    IN selectorIn VARCHAR(255),
    IN attrIn VARCHAR(255),
    IN xpathIn VARCHAR(255),
    IN jsonPathIn VARCHAR(255),
    IN enqueueCrawlIdIn INT,
    IN autoIncrementRulesIn JSON
) BEGIN
    INSERT INTO Parser
        (crawlId, createdTimestamp, typeId, 
        selector, attr, xpath, jsonPath, 
        enqueueCrawlId, autoIncrementRules)
    VALUES
        (crawlIdIn, NOW(), typeIdIn, 
        selectorIn, attrIn, xpathIn, jsonPathIn, 
        enqueueCrawlIdIn, autoIncrementRulesIn);
    SELECT LAST_INSERT_ID() AS parserId;
END $$

DELIMITER ;