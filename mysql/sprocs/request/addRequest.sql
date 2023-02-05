use harvest;

DROP PROCEDURE IF EXISTS addRequest;

DELIMITER $$

CREATE PROCEDURE addRequest(
    IN runIdIn INT,
    IN urlIn VARCHAR(1024),
    IN methodIn VARCHAR(8),
    IN requestBlobIn BLOB,
    IN parentRequestIdIn INT,
    IN originatorRequestIdIn INT
) BEGIN
    INSERT INTO Request
        (runId, visitedTimestamp, method, url, requestBlob, 
         parentRequestId, originatorRequestId)
    VALUES
        (runIdIn, NOW(), methodIn, urlIn, requestBlobIn, 
         parentRequestIdIn, originatorRequestIdIn);
    SELECT LAST_INSERT_ID() AS requestId;
END $$

DELIMITER ;