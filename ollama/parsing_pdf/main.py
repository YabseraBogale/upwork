
from google import genai
from pypdf import PdfReader
# The client gets the API key from the environment variable `GEMINI_API_KEY`.

reader=PdfReader("./app.pdf")

page=""

for i in reader.pages:
    page+=i.extract_text()
try:

    client = genai.Client()

    response = client.models.generate_content(
        model="gemini-2.5-flash", contents=f"can you give the returns from the route as json in {page}"
    )
    print(response.text)
except Exception as e:
    print(e)