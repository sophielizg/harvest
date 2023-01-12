use harvest;

CALL addForeignKeyIfNotExists(
    'RequestQueue', 
    'FK_RequestQueue_Scrape_scrapeId',
    'scrapeId',
    'Scrape(scrapeId)',
    'ON DELETE RESTRICT ON UPDATE CASCADE');

CALL addForeignKeyIfNotExists(
    'RequestQueue', 
    'FK_RequestQueue_Crawl_crawlId',
    'crawlId',
    'Crawl(crawlId)',
    'ON DELETE RESTRICT ON UPDATE CASCADE');
