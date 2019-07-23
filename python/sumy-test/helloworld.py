from sumy.nlp.stemmers import Stemmer
from sumy.nlp.tokenizers import Tokenizer
from sumy.parsers.html import HtmlParser
from sumy.summarizers.lsa import LsaSummarizer
from sumy.utils import get_stop_words

url = "https://en.wikipedia.org/wiki/%22Hello,_World!%22_program"
parser = HtmlParser.from_url(url, Tokenizer("english"))

stemmer = Stemmer("english")

summarizer = LsaSummarizer(stemmer)
summarizer.stop_words = get_stop_words("english")

for sentence in summarizer(parser.document, 8):
    print(sentence)