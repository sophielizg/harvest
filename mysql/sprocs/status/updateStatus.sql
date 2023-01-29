use harvest;

DROP PROCEDURE IF EXISTS updateStatus;

DELIMITER $$

CREATE PROCEDURE updateStatus(
    IN runIdIn INT,
    IN runnerIdIn INT,
    IN addQueued INT,
    IN addSuccesses INT,
    IN addErrors INT,
    IN addMissing INT
) BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM Status
        WHERE runId = runIdIn AND runnerId = runnerIdIn)
    THEN
        INSERT INTO Status
            (runId, runnerId, lastUpdatedTimestamp, 
            queued, successes, errors, missing)
        VALUES
            (runIdIn, runnerIdIn, NOW(),
            addQueued, addSuccesses, addErrors, addMissing);
    ELSE
        UPDATE Status SET
            lastUpdatedTimestamp = NOW(),
            queued = queued + addQueued,
            successes = successes + addSuccesses,
            errors = errors + addErrors,
            missing = missing + addMissing
        WHERE runId = runIdIn AND runnerId = runnerIdIn;
    END IF;
END $$

DELIMITER ;