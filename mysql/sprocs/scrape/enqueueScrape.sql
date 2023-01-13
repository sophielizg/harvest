use harvest;

DROP PROCEDURE IF EXISTS enqueueScrape;

DELIMITER $$

CREATE PROCEDURE enqueueScrape(
    IN crawlIdIn INT
) BEGIN
    DECLARE crawlRunId INT;

    IF EXISTS (
        SELECT 1 FROM Crawl c
        WHERE crawlId = crawlIdIn AND running = 1
    ) THEN
        SELECT crawlRunId INTO crawlRunId FROM CrawlRun
        WHERE crawlId = crawlIdIn AND endTimestamp IS NULL
        ORDER BY startTimestamp DESC
        LIMIT 1;

        INSERT INTO Scrape (crawlRunId)
        VALUES (crawlRunId);

        SELECT LAST_INSERT_ID() AS scrapeId;
    END IF;
END $$

DELIMITER ;