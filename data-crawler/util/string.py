def cleanup(string: str):
    return string.replace(" ", "").replace("\\t", "").replace("\r", "").replace("\n", "")


