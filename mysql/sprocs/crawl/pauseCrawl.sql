use harvest;

DROP PROCEDURE IF EXISTS pauseCrawl;

DELIMITER $$

CREATE PROCEDURE pauseCrawl(
    IN crawlIdIn INT
) BEGIN
    DECLARE currentCrawlRunId INT;

    SELECT crawlRunId INTO currentCrawlRunId FROM CrawlRun
    WHERE crawlId = crawlIdIn 
    ORDER BY startTimestamp DESC
    LIMIT 1;

    UPDATE CrawlRun SET
        running = 0
    WHERE crawlRunId = currentCrawlRunId;
END $$

DELIMITER ;