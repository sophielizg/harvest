use harvest;

DROP PROCEDURE IF EXISTS updateScraper;

DELIMITER $$

CREATE PROCEDURE updateScraper(
    IN scraperIdIn INT,
    IN nameIn VARCHAR(255),
    IN configIn JSON
) BEGIN
    UPDATE Scraper SET
        name = COALESCE(nameIn, name),
        config = COALESCE(configIn, config)
    WHERE scraperId = scraperIdIn;
END $$

DELIMITER ;