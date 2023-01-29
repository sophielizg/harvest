use harvest;

DROP PROCEDURE IF EXISTS stopRun;

DELIMITER $$

CREATE PROCEDURE stopRun(
    IN scraperIdIn INT,
    IN runIdIn INT,
    IN createTransaction BOOL
) BEGIN
    DECLARE currentRunId INT;

    DECLARE CONTINUE HANDLER FOR SQLEXCEPTION
    BEGIN
        IF createTransaction THEN
            ROLLBACK;
        ELSE
            ROLLBACK TO stopRun;
        END IF;
        RESIGNAL;
    END;

    IF createTransaction THEN
        START TRANSACTION;
    ELSE
        SAVEPOINT stopRun;
    END IF;

    IF (runIdIn IS NOT NULL) THEN
        SET currentRunId := runIdIn;
    ELSE 
        SELECT runId INTO currentRunId FROM Run
        WHERE scraperId = scraperIdIn AND isRunning = 1
        ORDER BY startTimestamp DESC
        LIMIT 1;
    END IF;

    UPDATE Run SET
        endTimestamp = NOW()
    WHERE runId = currentRunId;

    DELETE FROM RequestQueue
    WHERE runId = currentRunId AND runnerId IS NULL;

    SET @numDequeued := ROW_COUNT();
    CALL updateStatus(currentRunId, NULL, -1 * @numDequeued, 0, 0, 0);

    UPDATE Run SET
        running = 0
    WHERE runId = currentRunId;

    IF createTransaction THEN
        COMMIT;
    END IF;
END $$

DELIMITER ;