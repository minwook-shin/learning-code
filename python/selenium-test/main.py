from selenium import webdriver
from selenium.webdriver.common.keys import Keys
import time
from bs4 import BeautifulSoup

browser = webdriver.Safari()
# browser = webdriver.Firefox()
# browser = webdriver.Chrome('/chromedriver')

browser.get('http://www.google.com')

browser.find_element_by_name('q').send_keys("weather", Keys.RETURN)

time.sleep(2)

browser.save_screenshot("test.png")

soup = BeautifulSoup(browser.page_source, 'html.parser')
locate = soup.select('#wob_loc')
date = soup.select('#wob_dts')
rain_percent = soup.select('#wob_pp')

for t in locate:
    print(t.text)

for t in date:
    print(t.text)

for t in rain_percent:
    print(t.text)

browser.quit()