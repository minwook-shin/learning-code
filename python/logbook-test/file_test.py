import logbook

only_debug = logbook.FileHandler('./debug.log', filter=lambda r, h: r.extra['debug'])
everything = logbook.FileHandler('./all.log', bubble=True)

with only_debug, everything:
    logbook.info('this is debug code', extra={'debug': True})
    logbook.info('this is not debug code')