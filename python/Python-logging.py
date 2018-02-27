import logging

log = logging.getLogger("main")

stream = logging.StreamHandler()

log.addHandler(stream)

log.setLevel(logging.DEBUG)
log.debug("디버그")
log.info("안내")
log.warning("주의")
log.error("오류")
log.critical("치명적")

print("-"*10)

log.setLevel(logging.INFO)
log.debug("디버그")
log.info("안내")
log.warning("주의")
log.error("오류")
log.critical("치명적")

print("-"*10)

log.setLevel(logging.WARNING)
log.debug("디버그")
log.info("안내")
log.warning("주의")
log.error("오류")
log.critical("치명적")

print("-"*10)

log.setLevel(logging.ERROR)
log.debug("디버그")
log.info("안내")
log.warning("주의")
log.error("오류")
log.critical("치명적")

print("-"*10)

log.setLevel(logging.CRITICAL)
log.debug("디버그")
log.info("안내")
log.warning("주의")
log.error("오류")
log.critical("치명적")

print("-"*10)