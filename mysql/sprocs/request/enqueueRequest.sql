use harvest;

DROP PROCEDURE IF EXISTS enqueueRequest;

DELIMITER $$

CREATE PROCEDURE enqueueRequest(
    IN crawlIdIn INT,
    IN scrapeIdIn INT,
    IN requestIn BLOB,
    IN createdByRequestIdIn INT,
    IN isInitialRequestIn BOOL,
    IN createTransaction BOOL
) BEGIN
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
        (crawlId, createdTimestamp, request, createdByRequestId, isInitialRequest)
    VALUES
        (crawlIdIn, NOW(), requestIn, createdByRequestIdIn, isInitialRequestIn);

    CALL updateCrawlStatus(crawlIdIn, scrapeIdIn, 1, 0, 0, 0);

    IF createTransaction THEN
        COMMIT;
    END IF;
END $$

DELIMITER ;