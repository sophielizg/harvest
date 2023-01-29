use harvest;

CALL addForeignKeyIfNotExists(
    'RequestQueue', 
    'FK_RequestQueue_Scraper_scraperId',
    'scraperId',
    'Scraper(scraperId)',
    'ON DELETE RESTRICT ON UPDATE CASCADE');

CALL addForeignKeyIfNotExists(
    'RequestQueue', 
    'FK_RequestQueue_Run_runId',
    'runId',
    'Run(runId)',
    'ON DELETE RESTRICT ON UPDATE CASCADE');

CALL addForeignKeyIfNotExists(
    'RequestQueue', 
    'FK_RequestQueue_RunnerQueue_runnerId',
    'runnerId',
    'RunnerQueue(runnerId)',
    'ON DELETE RESTRICT ON UPDATE CASCADE');
