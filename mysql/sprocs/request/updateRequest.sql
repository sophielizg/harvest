use harvest;

DROP PROCEDURE IF EXISTS updateRequest;

DELIMITER $$

CREATE PROCEDURE updateRequest(
    IN requestIdIn INT,
    IN requestBlobIn BLOB,
    IN parentRequestIdIn INT,
    IN originatorRequestIdIn INT
) BEGIN
    UPDATE Request SET
        requestBlob = requestBlobIn,
        parentRequestId = parentRequestIdIn,
        originatorRequestId = originatorRequestIdIn
    WHERE requestId = requestIdIn;
END $$

DELIMITER ;