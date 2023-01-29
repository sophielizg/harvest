use harvest;

DROP PROCEDURE IF EXISTS toggleIsRunning;

DELIMITER $$

CREATE PROCEDURE toggleIsRunning(
    IN scraperIdIn INT,
    IN runIdIn INT,
    IN isRunningIn BOOL
) BEGIN
    DECLARE currentRunId INT;

    IF (runIdIn IS NOT NULL) THEN
        SET currentRunId := runIdIn;
    ELSE 
        SELECT runId INTO currentRunId FROM Run
        WHERE scraperId = scraperIdIn AND isRunning = 1
        ORDER BY startTimestamp DESC
        LIMIT 1;
    END IF;

    UPDATE Run SET
        isRunning = isRunningIn
    WHERE runId = currentRunId;
END $$

DELIMITER ;