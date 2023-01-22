use harvest;

DROP PROCEDURE IF EXISTS unpauseCrawl;

DELIMITER $$

CREATE PROCEDURE unpauseCrawl(
    IN crawlIdIn INT
) BEGIN
    DECLARE currentCrawlRunId INT;

    SELECT crawlRunId INTO currentCrawlRunId FROM CrawlRun
    WHERE crawlId = crawlIdIn 
    ORDER BY startTimestamp DESC
    LIMIT 1;

    UPDATE CrawlRun SET
        running = 1
    WHERE crawlRunId = currentCrawlRunId;
END $$

DELIMITER ;