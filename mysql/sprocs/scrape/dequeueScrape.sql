use harvest;

DROP PROCEDURE IF EXISTS dequeueScrape;

DELIMITER $$

CREATE PROCEDURE dequeueScrape()
BEGIN
    DECLARE `_rollback` BOOL DEFAULT 0;
    DECLARE CONTINUE HANDLER FOR SQLEXCEPTION SET `_rollback` = 1;

    START TRANSACTION;

    SELECT 
        @dequeueScrapeId := scrapeId, 
        @dequeueCrawlId := crawlId,
        @dequeueCrawlRunId := crawlRunId
    FROM Scrape
    WHERE startTimestamp IS NULL
    LIMIT 1
    FOR UPDATE SKIP LOCKED;

    UPDATE Scrape SET
        startTimestamp = NOW()
    WHERE scrapeId = @dequeueScrapeId;

    SELECT 
        @dequeueScrapeId AS scrapeId, 
        @dequeueCrawlId AS crawlId,
        @dequeueCrawlRunId AS crawlRunId;

    IF `_rollback` THEN
        ROLLBACK;
    ELSE
        COMMIT;
    END IF;
END $$

DELIMITER ;