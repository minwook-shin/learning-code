import sys

from logbook import Logger, StreamHandler


log = Logger('Stream handler logger')

StreamHandler(sys.stdout).push_application()

log.warn('warning')
log.error("error")


from logbook import StderrHandler


handler = StderrHandler()

handler.format_string = '{record.channel}: {record.message}'
handler.push_application()

log.warn('warning')
log.error("error")