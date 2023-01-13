use harvest;

DROP PROCEDURE IF EXISTS getCrawlStatus;

DELIMITER $$

CREATE PROCEDURE getCrawlStatus(
    IN crawlIdIn INT,
    IN crawlRunIdIn INT
) BEGIN
    DECLARE crawlRunIdForStatus INT;

    IF crawlRunIdIn IS NULL THEN
        SELECT crawlRunId INTO crawlRunIdForStatus FROM CrawlRun
        WHERE crawlId = crawlIdIn 
        ORDER BY startTimestamp DESC
        LIMIT 1;
    ELSE
        SET crawlRunIdForStatus := crawlRunIdIn;
    END IF;

    SELECT
        SUM(cs.queued) AS queued,
        SUM(cs.successes) AS successes, 
        SUM(cs.errors) AS errors, 
        SUM(cs.missing) AS missing,
        MAX(cs.lastUpdatedTimestamp) AS lastUpdatedTimestamp
    FROM CrawlStatus cs
    LEFT JOIN Scrape s ON s.scrapeId = cs.scrapeId
    WHERE s.crawlRunId = crawlRunIdForStatus
    OR (crawlRunIdIn IS NULL AND cs.scrapeId IS NULL AND cs.crawlId = crawlIdIn);
END $$

DELIMITER ;