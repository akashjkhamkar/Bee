import requests
import sys

def entry(request):
    print("hello", request, file=sys.stderr)
    return "Its working boiii"