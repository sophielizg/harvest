use harvest;

DROP PROCEDURE IF EXISTS getParserTypes;

DELIMITER $$

CREATE PROCEDURE getParserTypes()
BEGIN
    SELECT * FROM ParserType;
END $$

DELIMITER ;