use harvest;

CALL addForeignKeyIfNotExists(
    'Error', 
    'FK_Error_Run_runId',
    'runId',
    'Run(runId)',
    'ON DELETE RESTRICT ON UPDATE CASCADE');

CALL addForeignKeyIfNotExists(
    'Error', 
    'FK_Error_Request_requestId',
    'requestId',
    'Request(requestId)',
    'ON DELETE RESTRICT ON UPDATE CASCADE');

CALL addForeignKeyIfNotExists(
    'Error', 
    'FK_Error_Parser_parserId',
    'parserId',
    'Parser(parserId)',
    'ON DELETE RESTRICT ON UPDATE CASCADE');
