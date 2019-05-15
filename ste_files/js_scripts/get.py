import requests
from bs4 import BeautifulSoup as bs
import time

def download_file(url):
    if not url:
        return False
    filename = url[url.rfind("/")+1:]
    if "?" in filename:
        filename = filename[:filename.rfind("?")]
    with requests.get(url, stream=True) as r:
        with open(filename, 'wb') as f:
            for chunk in r.iter_content(chunk_size=8192): 
                if chunk: # filter out keep-alive new chunks
                    f.write(chunk)
    return filename

def write_to_file(txt):
    fname = "script_" + str(int(time.time())) + ".js"
    with open(fname, 'w+') as f:
        f.write(txt)
    
def fix_url(url):
    if url.startswith('//'):
        return "https:" + url
    return url

url = "https://www.novelupdates.com/novelslisting/?st=1"

source = requests.get(url)

soup = bs(source.text, 'html.parser')

script_tags = soup.find_all('script')

for tag in script_tags:
    if tag.has_attr("src"):
        if not tag["src"]:
            print("none", tag["src"])
            continue
        if not download_file(fix_url(tag["src"])):
            print(tag['src'], "failed")
    else:
        write_to_file(tag.text)
print(len(script_tags))