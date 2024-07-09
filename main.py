import os
import argparse
import openai

openai.api_key = "****************************"
openai.base_url = "*******************"
openai.default_headers = {"x-foo": "true"}

def get_ai_response(user_input):
    completion = openai.chat.completions.create(
        model="gpt-4o",
        messages=[
            {
                "role": "user",
                "content": user_input,
            },
        ],
    )
    return completion.choices[0].message.content

def main():
    parser = argparse.ArgumentParser(description="AI Chat Service")
    parser.add_argument('input', type=str, help="User input for the AI to process")
    args = parser.parse_args()

    ai_response = get_ai_response(args.input)

    print(ai_response)

if __name__ == '__main__':
    main()
