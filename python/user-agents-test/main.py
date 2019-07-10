from user_agents import parse

ua_str = 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.1.1 ' \
         'Safari/605.1.15'
user_agent = parse(ua_str)

print(user_agent.browser, user_agent.browser.family, user_agent.browser.version, user_agent.browser.version_string)

print(user_agent.os, user_agent.os.family, user_agent.os.version, user_agent.os.version_string)

print(user_agent.device, user_agent.device.family, user_agent.device.brand, user_agent.device.model)

print(str(user_agent))

from user_agents import parse

ua_str = 'BlackBerry9700/5.0.0.862 Profile/MIDP-2.1 Configuration/CLDC-1.1 VendorID/331 UNTRUSTED/1.0 3gpp-gba'
user_agent = parse(ua_str)
print(user_agent.is_mobile, user_agent.is_tablet, user_agent.is_pc)
print(user_agent.is_touch_capable)
print(user_agent.is_bot)
print(str(user_agent))

ua_str = 'Mozilla/5.0 (Linux; U; Android 4.0.4; en-gb; GT-I9300 Build/IMM76D) AppleWebKit/534.30 (KHTML, ' \
         'like Gecko) Version/4.0 Mobile Safari/534.30 '
user_agent = parse(ua_str)
print(user_agent.is_mobile, user_agent.is_tablet, user_agent.is_pc)
print(user_agent.is_touch_capable)
print(user_agent.is_bot)
print(str(user_agent))

ua_str = 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.1.1 ' \
         'Safari/605.1.15 '
user_agent = parse(ua_str)
print(user_agent.is_mobile, user_agent.is_tablet, user_agent.is_pc)
print(user_agent.is_touch_capable)
print(user_agent.is_bot)
print(str(user_agent))

ua_str = 'Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0; Touch)'
user_agent = parse(ua_str)
print(user_agent.is_mobile, user_agent.is_tablet, user_agent.is_pc)
print(user_agent.is_touch_capable)
print(user_agent.is_bot)
print(str(user_agent))
