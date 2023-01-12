use harvest;

DROP PROCEDURE IF EXISTS getCrawlStatus;

DELIMITER $$

CREATE PROCEDURE getCrawlStatus(
    IN crawlIdIn INT
) BEGIN
    SET @queued := (SELECT SUM(cs.queued) FROM CrawlStatus 
                    WHERE crawlId = crawlIdIn);

    SELECT
        @queued AS queued,
        SUM(cs.successes) AS successes, 
        SUM(cs.errors) AS errors, 
        SUM(cs.missing) AS missing,
        MAX(cs.lastUpdatedTimestamp) AS lastUpdatedTimestamp
    FROM CrawlStatus cs
    INNER JOIN Crawl c ON c.crawlId = cs.crawlId
    WHERE cs.crawlId = crawlIdIn 
    AND cs.lastUpdatedTimestamp >= c.currentRunStartTimestamp;
END $$

DELIMITER ;