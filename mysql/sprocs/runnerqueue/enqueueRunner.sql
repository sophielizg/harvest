use harvest;

DROP PROCEDURE IF EXISTS enqueueRunner;

DELIMITER $$

CREATE PROCEDURE enqueueRunner(
    IN scraperIdIn INT,
    IN runIdIn INT
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
    
    IF EXISTS (
        SELECT 1 FROM Run
        WHERE scraperId = scraperIdIn AND isRunning = 1
    ) THEN
        INSERT INTO RunnerQueue (runId)
        VALUES (currentRunId);

        SELECT LAST_INSERT_ID() AS runnerId;
    END IF;
END $$

DELIMITER ;