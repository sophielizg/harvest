use harvest;

DROP PROCEDURE IF EXISTS addScraper;

DELIMITER $$

CREATE PROCEDURE addScraper(
    IN nameIn VARCHAR(255),
    IN configIn JSON
) BEGIN
    INSERT INTO Scraper
        (name, createdTimestamp, config)
    VALUES
        (nameIn, NOW(), configIn);
    SELECT LAST_INSERT_ID() AS scraperId;
END $$

DELIMITER ;