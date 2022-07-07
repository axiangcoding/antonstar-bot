import re


def cleanup(string: str):
    if string is None:
        return 'N/A'
    return string.replace(" ", "").replace("\\t", "").replace("\r", "").replace("\n", "")

CLEAN = re.compile('<.*?>')

def cleanhtml(raw_html):
  cleantext = re.sub(CLEAN, '', raw_html)
  return cleantext
