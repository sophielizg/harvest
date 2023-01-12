use harvest;

DROP PROCEDURE IF EXISTS endScrape;

DELIMITER $$

CREATE PROCEDURE endScrape(
    IN scrapeIdIn INT
)
BEGIN
    UPDATE Scrape SET
        endTimestamp = NOW()
    WHERE scrapeId = scrapeIdIn;
END $$

DELIMITER ;