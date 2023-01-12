use harvest;

DROP PROCEDURE IF EXISTS addParserTag;

DELIMITER $$

CREATE PROCEDURE addParserTag(
    IN parserIdIn INT,
    IN nameIn VARCHAR(255)
) BEGIN
    INSERT INTO ParserTag
        (parserId, createdTimestamp, name)
    VALUES
        (parserIdIn, NOW(), nameIn);
    SELECT LAST_INSERT_ID() AS parserTagId;
END $$

DELIMITER ;