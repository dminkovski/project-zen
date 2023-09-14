from django.http import HttpResponse
from django.views.decorators.csrf import csrf_exempt
import os
import json
import openai
from dotenv import load_dotenv

load_dotenv()

openai.api_key = os.getenv("AZURE_OPENAI_KEY")
openai.api_base = os.getenv("AZURE_OPENAI_ENDPOINT") # your endpoint should look like the following https://YOUR_RESOURCE_NAME.openai.azure.com/
openai.api_type = 'azure'
openai.api_version = '2023-05-15' # this may change in the future

deployment_name='my-gpt-model' #This will correspond to the custom name you chose for your deployment when you deployed a model. 

@csrf_exempt
def index(request):
    if request.method == 'GET':
        return HttpResponse("Hello, world. You're at the polls index.")
    elif request.method == 'POST':
        received_json_data=json.loads(request.body)
        conversation=[{"role": "system", "content": "You are a newsletter editor who needs to summarize a newsletter for a busy executive. The executive wants to know the key information, important links, and any discount codes mentioned."}]
        conversation.append({"role": "user", "content": "Please summarize this latest newsletter I received in my email. I'm particularly interested in the key information, important links, and any discount codes mentioned. Make sure to highlight these elements for easy reference in bullet points. Maintain the summary in points format without any intro and outro and salutation. The summary should be shorter than the input text. Thank you!"+ received_json_data['message']})
        response = openai.ChatCompletion.create(
            engine=deployment_name,
            messages=conversation
        )
        conversation.append({"role": "assistant", "content": response["choices"][0]["message"]["content"]})
        print("\n" + response['choices'][0]['message']['content'] + "\n")
        return HttpResponse(response['choices'][0]['message']['content'])
        #return HttpResponse(received_json_data['title'])
        #return HttpResponse(request.body)
