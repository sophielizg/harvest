use harvest;

DROP PROCEDURE IF EXISTS createRun;

DELIMITER $$

CREATE PROCEDURE createRun(
    IN scraperIdIn INT,
    IN createTransaction BOOL
) BEGIN
    DECLARE newRunId INT;

    DECLARE CONTINUE HANDLER FOR SQLEXCEPTION
    BEGIN
        IF createTransaction THEN
            ROLLBACK;
        ELSE
            ROLLBACK TO createRun;
        END IF;
        RESIGNAL;
    END;

    IF createTransaction THEN
        START TRANSACTION;
    ELSE
        SAVEPOINT createRun;
    END IF;

    INSERT INTO Run
        (scraperId, isRunning, startTimestamp)
    VALUES
        (scraperIdIn, 1, NOW());

    SELECT LAST_INSERT_ID() INTO newRunId;

    INSERT INTO RequestQueue
        (scraperId, runId, runnerId, createdTimestamp, requestBlob, isInitialRequest)
    SELECT
        scraperId, newRunId, NULL, NOW(), requestBlob, 0
    FROM RequestQueue
    WHERE isInitialRequest = 1;

    SET @numQueued := ROW_COUNT();
    CALL updateStatus(newRunId, NULL, @numQueued, 0, 0, 0);

    SELECT newRunId AS runId;

    IF createTransaction THEN
        COMMIT;
    END IF;
END $$

DELIMITER ;