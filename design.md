
Design ideas
version: draft
date: 2017.12.09

we know what we're building here.
what components do we need ?

1. ticker data gathering tool, this should run continuously somewhere, concurrency here is important. i think this should be config driven (as opposed to flags driven). this way we can have different configs for different things, mix match crypto etfs with real etfs, mix match different types of data sources, etc.

2. a reader / processor to handle the data we gather, concurrency is also important here. should we do caching ? process blocks of time (1 month), the window we slide over is 6 months, each month of data that is process can be cached and then stored, only data from the new month needs to be processed and evaluated, this could ease up the rebalancing calcs (but it may not matter if we just rebalance infrequently)

3. i'd personally like some kind of visualizer to easily see whats going on month to month, get stats

