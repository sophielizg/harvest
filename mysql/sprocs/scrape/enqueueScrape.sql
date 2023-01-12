use harvest;

DROP PROCEDURE IF EXISTS enqueueScrape;

DELIMITER $$

CREATE PROCEDURE enqueueScrape(
    IN crawlIdIn INT
) BEGIN
    INSERT INTO Scrape (crawlId)
    VALUES (crawlIdIn);
    SELECT LAST_INSERT_ID() AS scrapeId;
END $$

DELIMITER ;