use harvest;

DROP PROCEDURE IF EXISTS deleteParserTag;

DELIMITER $$

CREATE PROCEDURE deleteParserTag(
    IN parserIdIn INT,
    IN nameIn VARCHAR(255)
) BEGIN
    DELETE FROM ParserTag
    WHERE parserId = parserIdIn AND name = nameIn;
END $$

DELIMITER ;