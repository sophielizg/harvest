use harvest;

CALL addForeignKeyIfNotExists(
    'Result', 
    'FK_Result_Run_runId',
    'runId',
    'Run(runId)',
    'ON DELETE RESTRICT ON UPDATE CASCADE');

CALL addForeignKeyIfNotExists(
    'Result', 
    'FK_Result_Request_requestId',
    'requestId',
    'Request(requestId)',
    'ON DELETE RESTRICT ON UPDATE CASCADE');

CALL addForeignKeyIfNotExists(
    'Result', 
    'FK_Result_Parser_parserId',
    'parserId',
    'Parser(parserId)',
    'ON DELETE RESTRICT ON UPDATE CASCADE');
