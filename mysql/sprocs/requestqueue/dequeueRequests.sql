use harvest;

DROP PROCEDURE IF EXISTS dequeueRequests;

DELIMITER $$

CREATE PROCEDURE dequeueRequests(
    IN runIdIn INT,
    IN runnerIdIn INT,
    IN numToDequeue INT,
    IN createTransaction BOOL
)
BEGIN
    DECLARE dequeueIds VARCHAR(4096);

    DECLARE CONTINUE HANDLER FOR SQLEXCEPTION
    BEGIN
        IF createTransaction THEN
            ROLLBACK;
        ELSE
            ROLLBACK TO dequeueRequests;
        END IF;
        RESIGNAL;
    END;

    IF createTransaction THEN
        START TRANSACTION;
    ELSE
        SAVEPOINT dequeueRequests;
    END IF;

    SELECT 
        GROUP_CONCAT(requestQueueId) INTO dequeueIds
    FROM RequestQueue
    WHERE runId = runIdIn AND runnerId IS NULL
    LIMIT numToDequeue
    FOR UPDATE SKIP LOCKED;

    UPDATE RequestQueue SET
        runnerId = runnerIdIn
    WHERE FIND_IN_SET(requestQueueId, dequeueIds);

    SET @numDequeued := ROW_COUNT();

    CALL updateStatus(runIdIn, runnerIdIn, -1 * @numDequeued, 0, 0, 0);

    SELECT requestBlob FROM RequestQueue
    WHERE FIND_IN_SET(requestQueueId, dequeueIds);

    IF createTransaction THEN
        COMMIT;
    END IF;
END $$

DELIMITER ;