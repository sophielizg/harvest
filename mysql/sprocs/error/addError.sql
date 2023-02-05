use harvest;

DROP PROCEDURE IF EXISTS addError;

DELIMITER $$

CREATE PROCEDURE addError(
    IN runIdIn INT,
    IN runnerIdIn INT,
    IN requestIdIn INT,
    IN parserIdIn INT,
    IN statusCodeIn INT,
    IN responseIn TEXT,
    IN isMissngParseResultIn BOOL,
    IN errorMessageIn VARCHAR(4096),
    IN createTransaction BOOL
) BEGIN
    DECLARE newErrorId INT;

    DECLARE CONTINUE HANDLER FOR SQLEXCEPTION
    BEGIN
        IF createTransaction THEN
            ROLLBACK;
        ELSE
            ROLLBACK TO addError;
        END IF;
        RESIGNAL;
    END;

    IF createTransaction THEN
        START TRANSACTION;
    ELSE
        SAVEPOINT addError;
    END IF;

    INSERT INTO Error
        (runId, requestId, parserId, scrapedTimestamp,
         statusCode, response, isMissngParseResult, errorMessage)
    VALUES
        (runIdIn, requestIdIn, parserIdIn, NOW(),
         statusCodeIn, responseIn, isMissngParseResultIn, errorMessageIn);

    SELECT LAST_INSERT_ID() INTO newErrorId;

    CALL updateStatus(
        runIdIn, runnerIdIn, 0, 0, 
        IF(isMissngParseResultIn, 0, 1), 
        IF(isMissngParseResultIn, 1, 0));

    SELECT newErrorId AS errorId;
    
    IF createTransaction THEN
        COMMIT;
    END IF;
END $$

DELIMITER ;