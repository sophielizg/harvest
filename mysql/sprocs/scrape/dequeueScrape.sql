use harvest;

DROP PROCEDURE IF EXISTS dequeueScrape;

DELIMITER $$

CREATE PROCEDURE dequeueScrape(
    IN createTransaction BOOL
)
BEGIN
    DECLARE dequeueScrapeId INT;
    DECLARE dequeueCrawlRunId INT;

    DECLARE CONTINUE HANDLER FOR SQLEXCEPTION
    BEGIN
        IF createTransaction THEN
            ROLLBACK;
        ELSE
            ROLLBACK TO dequeueScrape;
        END IF;
        RESIGNAL;
    END;

    IF createTransaction THEN
        START TRANSACTION;
    ELSE
        SAVEPOINT dequeueScrape;
    END IF;

    SELECT 
        scrapeId, crawlRunId  
        INTO
        dequeueScrapeId, dequeueCrawlRunId
    FROM Scrape
    WHERE startTimestamp IS NULL
    LIMIT 1
    FOR UPDATE SKIP LOCKED;

    UPDATE Scrape SET
        startTimestamp = NOW()
    WHERE scrapeId = dequeueScrapeId;

    SELECT 
        dequeueScrapeId AS scrapeId,
        dequeueCrawlRunId AS crawlRunId,
        crawlId
    FROM CrawlRun
    WHERE crawlRunId = dequeueCrawlRunId;

    IF createTransaction THEN
        COMMIT;
    END IF;
END $$

DELIMITER ;