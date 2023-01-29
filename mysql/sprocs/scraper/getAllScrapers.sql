use harvest;

DROP PROCEDURE IF EXISTS getAllScrapers;

DELIMITER $$

CREATE PROCEDURE getAllScrapers()
BEGIN
    SELECT * FROM Scraper;
END $$

DELIMITER ;