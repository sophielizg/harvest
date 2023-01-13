use harvest;

DROP PROCEDURE IF EXISTS getRequestIsVisited;

DELIMITER $$

CREATE PROCEDURE getRequestIsVisited(
    IN scrapeIdIn INT,
    IN requestHashIn VARCHAR(16)
) BEGIN
    DECLARE currentCrawlRunId INT;

    SELECT crawlRunId INTO currentCrawlRunId FROM CrawlRun
    WHERE crawlId = crawlIdIn 
    ORDER BY startTimestamp DESC
    LIMIT 1;

    SELECT 1 FROM Request r
    INNER JOIN Scrape s ON r.scrapeId = s.scrapeId
    WHERE r.requestHash = requestHashIn AND s.crawlRunId = currentCrawlRunId;
END $$

DELIMITER ;