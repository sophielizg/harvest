use harvest;

DROP PROCEDURE IF EXISTS addResult;

DELIMITER $$

CREATE PROCEDURE addResult(
    IN runIdIn INT,
    IN runnerIdIn INT,
    IN requestIdIn INT,
    IN parserIdIn INT,
    IN elementIndexIn INT,
    IN valueIn TEXT,
    IN createTransaction BOOL
) BEGIN
    DECLARE newResultId INT;

    DECLARE CONTINUE HANDLER FOR SQLEXCEPTION
    BEGIN
        IF createTransaction THEN
            ROLLBACK;
        ELSE
            ROLLBACK TO addResult;
        END IF;
        RESIGNAL;
    END;

    IF createTransaction THEN
        START TRANSACTION;
    ELSE
        SAVEPOINT addResult;
    END IF;

    INSERT INTO Result
        (runId, requestId, parserId, elementIndex, scrapedTimestamp, value)
    VALUES
        (runIdIn, requestIdIn, parserIdIn, elementIndexIn, NOW(), valueIn);

    SELECT LAST_INSERT_ID() INTO newResultId;

    CALL updateStatus(runIdIn, runnerIdIn, 0, 1, 0, 0);

    SELECT newResultId AS resultId;
    
    IF createTransaction THEN
        COMMIT;
    END IF;
END $$

DELIMITER ;