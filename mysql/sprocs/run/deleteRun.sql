use harvest;

DROP PROCEDURE IF EXISTS deleteRun;

DELIMITER $$

CREATE PROCEDURE deleteRun(
    IN runIdIn INT,
    IN createTransaction BOOL
) BEGIN
    DECLARE relatedScraperId INT;
    DECLARE totalLeftInQueue INT;

    DECLARE CONTINUE HANDLER FOR SQLEXCEPTION
    BEGIN
        IF createTransaction THEN
            ROLLBACK;
        ELSE
            ROLLBACK TO deleteRun;
        END IF;
        RESIGNAL;
    END;

    IF createTransaction THEN
        START TRANSACTION;
    ELSE
        SAVEPOINT deleteRun;
    END IF;

    SELECT scraperId INTO relatedScraperId FROM Run
    WHERE runId = runIdIn;

    UPDATE Request
    SET 
        parentRequestId = NULL,
        originatorRequestId = NULL
    WHERE runId = runIdIn;

    DELETE FROM Result
    WHERE runId = runIdIn;

    DELETE FROM Error
    WHERE runId = runIdIn;

    DELETE FROM Status
    WHERE runId = runIdIn;

    DELETE FROM RequestQueue
    WHERE runId = runIdIn;

    DELETE FROM Run
    WHERE runId = runIdIn;

    IF createTransaction THEN
        COMMIT;
    END IF;
END $$

DELIMITER ;