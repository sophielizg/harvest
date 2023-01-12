use harvest;

DROP PROCEDURE IF EXISTS dequeueScrape;

DELIMITER $$

CREATE PROCEDURE dequeueScrape()
BEGIN
    SELECT 
        @dequeueScrapeId := scrapeId, 
        @dequeueCrawlId := crawlId 
    FROM Scrape
    WHERE startTimestamp IS NULL
    LIMIT 1
    FOR UPDATE SKIP LOCKED;

    UPDATE Scrape SET
        startTimestamp = NOW()
    WHERE scrapeId = @dequeueScrapeId;

    SELECT @dequeueScrapeId AS scrapeId, @dequeueCrawlId AS crawlId;
END $$

DELIMITER ;