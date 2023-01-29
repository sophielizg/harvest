use harvest;

DROP PROCEDURE IF EXISTS addOrUpdateRequest;

DELIMITER $$

CREATE PROCEDURE addOrUpdateRequest(
    IN runIdIn INT,
    IN requestIdIn INT UNSIGNED,
    IN requestBlobIn BLOB,
    IN parentRequestIdIn INT,
    IN originatorRequestIdIn INT
) BEGIN
    INSERT INTO Request
        (requestId, runId, visitedTimestamp, requestBlob, parentRequestId, originatorRequestId)
    VALUES
        (requestIdIn, runIdIn, NOW(), requestBlobIn, parentRequestIdIn, originatorRequestIdIn)
    ON DUPLICATE KEY UPDATE
        requestBlob = requestBlobIn,
        parentRequestId = parentRequestIdIn,
        originatorRequestId = originatorRequestIdIn;
END $$

DELIMITER ;