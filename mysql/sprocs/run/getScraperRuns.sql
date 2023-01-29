use harvest;

DROP PROCEDURE IF EXISTS getScraperRuns;

DELIMITER $$

CREATE PROCEDURE getScraperRuns(
    IN scraperIdIn INT
) BEGIN
    SELECT * FROM Run
    WHERE scraperId = scraperIdIn;
END $$

DELIMITER ;