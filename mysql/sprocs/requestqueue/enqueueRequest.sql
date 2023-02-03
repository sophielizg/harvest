use harvest;

DROP PROCEDURE IF EXISTS enqueueRequest;

DELIMITER $$

CREATE PROCEDURE enqueueRequest(
    IN scraperIdIn INT,
    IN runIdIn INT,
    IN runnerIdIn INT,
    IN requestBlobIn BLOB,
    IN isInitialRequestIn BOOL,
    IN createTransaction BOOL
) BEGIN
    DECLARE requestQueueId INT;

    DECLARE CONTINUE HANDLER FOR SQLEXCEPTION
    BEGIN
        IF createTransaction THEN
            ROLLBACK;
        ELSE
            ROLLBACK TO enqueueRequest;
        END IF;
        RESIGNAL;
    END;

    IF createTransaction THEN
        START TRANSACTION;
    ELSE
        SAVEPOINT enqueueRequest;
    END IF;

    INSERT INTO RequestQueue
        (scraperId, runId, createdTimestamp, requestBlob, isInitialRequest)
    VALUES
        (scraperIdIn, runIdIn, NOW(), requestBlobIn, isInitialRequestIn);

    SELECT LAST_INSERT_ID() INTO requestQueueId;

    IF (runIdIn IS NOT NULL) THEN
        CALL updateStatus(runIdIn, runnerIdIn, 1, 0, 0, 0);
    END IF;

    SELECT requestQueueId;

    IF createTransaction THEN
        COMMIT;
    END IF;
END $$

DELIMITER ;