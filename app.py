from fastapi import FastAPI, Request
from fastapi.middleware.cors import CORSMiddleware
import requests

app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

API_KEY = "sk-roG3OusRr0TLCHAADks6lw"
BASE_URL = "https://openai-hub.neuraldeep.tech/v1/chat/completions"

# Simulated user data (this would normally come from the bank API or database)
USER_PROFILE = {
    "name": "Ali",
    "age": 22,
    "occupation": "student",
    "monthly_income": 180000,  # KZT
    "monthly_expenses": {
        "food": 50000,
        "transport": 15000,
        "entertainment": 20000,
        "savings": 15000
    },
    "goals": [
        {"name": "Buy a laptop", "target": 600000, "progress": 150000},
        {"name": "Trip to Almaty", "target": 200000, "progress": 50000}
    ],
    "financial_habits": [
        "Uses debit card for all payments",
        "Prefers halal/ethical banking options",
        "Struggles with impulse spending on delivery apps"
    ],
    "bank_products": [
        "Halal Deposit Plan",
        "Education Savings Account",
        "Sharia-compliant Credit Card"
    ]
}


@app.post("/chat")
async def chat(request: Request):
    data = await request.json()
    user_message = data.get("message", "")

    # Dynamic system prompt — gives the AI background about the user
    system_prompt = f"""
You are Tyler, an empathetic AI assistant for Zaman Bank that helps users achieve financial goals
using Islamic finance principles. You know this about the user:

Name: {USER_PROFILE['name']}
Age: {USER_PROFILE['age']}
Occupation: {USER_PROFILE['occupation']}
Monthly income: {USER_PROFILE['monthly_income']} KZT
Monthly expenses: {USER_PROFILE['monthly_expenses']}
Current goals: {USER_PROFILE['goals']}

You are also aware of the following:
- Financial habits and preferences the user has shared (e.g., "prefers halal banking or not religious at all", "struggles with impulse spending").
- The user's behavior and transactions within the app (such as how they interact with different financial products or make payments).
- Their banking history and usage (debit card usage, financial product interests).

Your mission:
- Help the user manage money wisely.
- Recommend suitable Zaman Bank products.
- Analyze spending patterns.
- Suggest better financial habits and stress-coping methods.
- Use natural, human-like tone (motivating and friendly).
- Always align with Islamic finance ethics (no interest/usury).

When giving advice, always:
- Mention their specific goals or habits if relevant.
- Keep responses short, friendly, and motivating.
- Provide actionable steps to improve their financial health.

When giving advice, always mention their specific goals or habits if relevant, 
keep everything short(like maximum of 30 words, exception when user ask for it) and simple so it wont unattract user.
"""

    payload = {
        "model": "gpt-4o-mini",
        "messages": [
            {"role": "system", "content": system_prompt},
            {"role": "user", "content": user_message}
        ]
    }

    headers = {"Authorization": f"Bearer {API_KEY}"}
    response = requests.post(BASE_URL, headers=headers, json=payload)
    result = response.json()

    answer = result["choices"][0]["message"]["content"]
    return {"reply": answer}
