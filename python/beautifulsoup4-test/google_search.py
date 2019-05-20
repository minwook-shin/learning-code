from urllib.request import Request, urlopen
from bs4 import BeautifulSoup

header = {'User-Agent': 'Mozilla/5.0'}

word = input()
html = Request('https://www.google.com/search?q='+str(word), headers=header)
source = urlopen(html)
soup = BeautifulSoup(source, "html.parser")

for search in soup.find_all("div", id="search"):
    for article in search.find_all("h3"):
        print(article.find('a').text)
