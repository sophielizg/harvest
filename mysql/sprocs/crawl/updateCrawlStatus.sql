use harvest;

DROP PROCEDURE IF EXISTS updateCrawlStatus;

DELIMITER $$

CREATE PROCEDURE updateCrawlStatus(
    IN crawlIdIn INT,
    IN scrapeIdIn INT,
    IN addQueued INT,
    IN addSuccesses INT,
    IN addErrors INT,
    IN addMissing INT
) BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM CrawlStatus
        WHERE crawlId = crawlIdIn AND scrapeId = scrapeIdIn)
    THEN
        INSERT INTO CrawlStatus
            (crawlId, scrapeId, lastUpdatedTimestamp, 
            queued, successes, errors, missing)
        VALUES
            (crawlIdIn, scrapeIdIn, NOW(),
            addQueued, addSuccesses, addErrors, addMissing);
    ELSE
        UPDATE CrawlStatus SET
            lastUpdatedTimestamp = NOW(),
            queued = queued + addQueued,
            successes = successes + addSuccesses,
            errors = errors + addErrors,
            missing = missing + addMissing
        WHERE crawlId = crawlIdIn AND scrapeId = scrapeIdIn;
    END IF;
END $$

DELIMITER ;