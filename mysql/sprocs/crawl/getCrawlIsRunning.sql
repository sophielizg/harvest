use harvest;

DROP PROCEDURE IF EXISTS getCrawlIsRunning;

DELIMITER $$

CREATE PROCEDURE getCrawlIsRunning(
    IN crawlIdIn INT,
    IN crawlRunIdIn INT
) BEGIN
    DECLARE crawlRunIdForIsRunning INT;

    IF crawlRunIdIn IS NULL THEN
        SELECT crawlRunId INTO crawlRunIdForIsRunning FROM CrawlRun
        WHERE crawlId = crawlIdIn 
        ORDER BY startTimestamp DESC
        LIMIT 1;
    ELSE
        SET crawlRunIdForIsRunning := crawlRunIdIn;
    END IF;

    SELECT running FROM CrawlRun WHERE crawlRunId = crawlRunIdForIsRunning;
END $$

DELIMITER ;