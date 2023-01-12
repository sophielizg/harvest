use harvest;

DROP PROCEDURE IF EXISTS deleteParser;

DELIMITER $$

CREATE PROCEDURE deleteParser(
    IN parserIdIn INT
) BEGIN
    DELETE FROM Parser
    WHERE parserId = parserIdIn;
END $$

DELIMITER ;