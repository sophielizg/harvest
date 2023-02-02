use harvest;

DROP PROCEDURE IF EXISTS addRequest;

DELIMITER $$

CREATE PROCEDURE addRequest(
    IN runIdIn INT,
    IN requestBlobIn BLOB,
    IN parentRequestIdIn INT,
    IN originatorRequestIdIn INT
) BEGIN
    INSERT INTO Request
        (runId, visitedTimestamp, requestBlob, parentRequestId, originatorRequestId)
    VALUES
        (runIdIn, NOW(), requestBlobIn, parentRequestIdIn, originatorRequestIdIn);
    SELECT LAST_INSERT_ID() AS requestId;
END $$

DELIMITER ;