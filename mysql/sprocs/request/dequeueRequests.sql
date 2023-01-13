use harvest;

DROP PROCEDURE IF EXISTS dequeueRequests;

DELIMITER $$

CREATE PROCEDURE dequeueRequests(
    IN crawlIdIn INT,
    IN scrapeIdIn INT,
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
    WHERE crawlId = crawlIdIn AND scrapeId IS NULL
    LIMIT numToDequeue
    FOR UPDATE SKIP LOCKED;

    UPDATE RequestQueue SET
        scrapeId = scrapeIdIn
    WHERE FIND_IN_SET(requestQueueId, dequeueIds);

    SET @numDequeued := ROW_COUNT();

    CALL updateCrawlStatus(crawlIdIn, scrapeIdIn, -1 * @numDequeued, 0, 0, 0);

    SELECT request, createdByRequestId FROM RequestQueue
    WHERE FIND_IN_SET(requestQueueId, dequeueIds);

    IF createTransaction THEN
        COMMIT;
    END IF;
END $$

DELIMITER ;