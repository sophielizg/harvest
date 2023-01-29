use harvest;

DROP PROCEDURE IF EXISTS dequeueRunner;

DELIMITER $$

CREATE PROCEDURE dequeueRunner(
    IN createTransaction BOOL
)
BEGIN
    DECLARE dequeueRunnerId INT;
    DECLARE dequeueRunId INT;

    DECLARE CONTINUE HANDLER FOR SQLEXCEPTION
    BEGIN
        IF createTransaction THEN
            ROLLBACK;
        ELSE
            ROLLBACK TO dequeueRunner;
        END IF;
        RESIGNAL;
    END;

    IF createTransaction THEN
        START TRANSACTION;
    ELSE
        SAVEPOINT dequeueRunner;
    END IF;

    SELECT 
        runnerId, runId  
        INTO
        dequeueRunnerId, dequeueRunId
    FROM RunnerQueue
    WHERE startTimestamp IS NULL
    LIMIT 1
    FOR UPDATE SKIP LOCKED;

    UPDATE RunnerQueue SET
        startTimestamp = NOW()
    WHERE runnerId = dequeueRunnerId;

    SELECT 
        dequeueRunnerId AS runnerId,
        dequeueRunId AS runId,
        scraperId
    FROM Run
    WHERE runId = dequeueRunId;

    IF createTransaction THEN
        COMMIT;
    END IF;
END $$

DELIMITER ;