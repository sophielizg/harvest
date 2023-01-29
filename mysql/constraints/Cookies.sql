use harvest;

CALL addForeignKeyIfNotExists(
    'Cookies', 
    'FK_Cookies_Run_runId',
    'runId',
    'Run(runId)',
    'ON DELETE CASCADE ON UPDATE CASCADE');
