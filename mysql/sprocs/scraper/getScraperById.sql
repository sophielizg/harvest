use harvest;

DROP PROCEDURE IF EXISTS getScraperById;

DELIMITER $$

CREATE PROCEDURE getScraperById(
    IN scraperIdIn INT
) BEGIN
    SELECT * FROM Scraper WHERE scraperId = scraperIdIn;
END $$

DELIMITER ;