use harvest;

DROP PROCEDURE IF EXISTS getCrawlStatus;

DELIMITER $$

CREATE PROCEDURE getCrawlStatus(
    IN crawlIdIn INT
) BEGIN
    SELECT @crawlRunId := crawlRunId FROM CrawlRun
    WHERE crawlId = crawlIdIn 
    ORDER BY startTimestamp DESC
    LIMIT 1;

    SELECT
        SUM(cs.queued) AS queued,
        SUM(cs.successes) AS successes, 
        SUM(cs.errors) AS errors, 
        SUM(cs.missing) AS missing,
        MAX(cs.lastUpdatedTimestamp) AS lastUpdatedTimestamp
    FROM CrawlStatus cs
    INNER JOIN Crawl c ON c.crawlId = cs.crawlId
    LEFT JOIN Scrape s ON s.scrapeId = cs.scrapeId
    WHERE cs.crawlId = crawlIdIn 
    AND (s.crawlRunId = @crawlRunId OR cs.scrapeId IS NULL);
END $$

DELIMITER ;