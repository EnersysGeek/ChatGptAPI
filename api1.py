import openai

# Set your API Key
openai.api_key = "Your_Key"

# Send a request to the ChatGPT API
response = openai.ChatCompletion.create(
    model="gpt-4o",  # Choose the model
    messages=[
        {"role": "system", "content": "You are an AI your name is EnersysBot"},
        {"role": "user", "content": "เราพบว่าโลกไม่แบน เมื่อไหร่ โดยใคร"}
    ]
)

# Print the response
print(response.choices[0].message["content"])
