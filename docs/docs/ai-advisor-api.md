# AI Advisor API

Source: https://kb.kentik.com/docs/ai-advisor-api

---

The AI Advisor REST API provides programmatic access to Kentik's AI-powered network analysis capabilities. This asynchronous API enables you to create chat sessions by submitting natural language questions about your network and retrieving AI-generated insights and analysis.

* **API Version:** /ai\_advisor/v202511
* **Protocol:** HTTPS
* **Format:** JSON

## Kentik REST API Basics

### Authentication

All API requests require authentication using HTTP headers:

```
X-CH-Auth-Email: your-email@company.com
X-CH-Auth-API-Token: your-api-token-here
```

### Getting Your API Token

1. Log in to the Kentik Portal.
2. Navigate to **User Profile** (top-right menu).
3. Go to **Authentication** tab.
4. Copy your **API Token.**

### Kentik API URL

* If your company is in US cluster: `https://grpc.api.kentik.com`
* If your company is in EU cluster: `https://grpc.api.kentik.eu`

## API Endpoints

### 1. Create Chat Session

This endpoint creates a new AI Advisor conversation and submits an initial question.

**Endpoint:** `POST /ai_advisor/v202511/chat`

**Request Body**

```
{
  "prompt": "How are my devices doing?"
}
```

**Response**

```
{
  "id": "019ab80d-f7a1-739b-99e0-803f501fdea9",
  "status": "SESSION_STATUS_PENDING"
}
```

**cURL Example**

```
curl -X POST https://grpc.api.kentik.com/ai_advisor/v202511/chat \
  -H "X-CH-Auth-Email: your-email@company.com" \
  -H "X-CH-Auth-API-Token: your-api-token" \
  -H "Content-Type: application/json" \
  -d '{"prompt": "Show me top talkers in the last hour"}'
```

### 2. Get Chat Session

This endpoint retrieves the current status and results of a chat session. Use this endpoint to poll for completion.

**Endpoint:** `GET /ai_advisor/v202511/chat/{id}`

**Path Parameters**

* `id` (required): Session UUID returned from Create or Update

**Response example with status PROCESSING**

```
{
  "id": "019ab80d-f7a1-739b-99e0-803f501fdea9",
  "status": "SESSION_STATUS_PROCESSING",
  "messages": [
    {
      "id": "2e1e0146-fdde-472c-9ca5-fd4fefe5ff92",
      "status": "SESSION_STATUS_PROCESSING",
      "prompt": "How are my devices doing?",
      "finalAnswer": "",
      "reasoning": "I'll check the health status of your network devices...",
      "errorMessage": "",
      "createdAt": "2025-11-24T22:48:34.701Z",
      "updatedAt": "2025-11-24T22:48:51.422Z"
    }
  ]
}
```

**Response example with status COMPLETED**

```
{
  "id": "019ab80d-f7a1-739b-99e0-803f501fdea9",
  "status": "SESSION_STATUS_COMPLETED",
  "messages": [
    {
      "id": "2e1e0146-fdde-472c-9ca5-fd4fefe5ff92",
      "status": "SESSION_STATUS_COMPLETED",
      "prompt": "How are my devices doing?",
      "finalAnswer": "## Device Health Overview

Your network currently has **6 active devices**...",
      "reasoning": "I'll check the health status of your network devices...",
      "errorMessage": "",
      "createdAt": "2025-11-24T22:48:34.701Z",
      "updatedAt": "2025-11-24T22:48:51.422Z"
    }
  ]
}
```

**cURL Example**

```
curl -X GET https://grpc.api.kentik.com/ai_advisor/v202511/chat/019ab80d-f7a1-739b-99e0-803f501fdea9 \
  -H "X-CH-Auth-Email: your-email@company.com" \
  -H "X-CH-Auth-API-Token: your-api-token"
```

### 3. Update Chat Session

This endpoint adds a follow-up question to an existing conversation, maintaining context from previous messages.

**Endpoint:** `PUT /ai_advisor/v202511/chat`

**Request Body**

```
{
  "id": "019ab80d-f7a1-739b-99e0-803f501fdea9",
  "prompt": "What about interface utilization?"
}
```

**Response**

```
{
  "id": "019ab80d-f7a1-739b-99e0-803f501fdea9",
  "status": "SESSION_STATUS_PENDING"
}
```

**cURL Example**

```
curl -X PUT https://grpc.api.kentik.com/ai_advisor/v202511/chat \
  -H "X-CH-Auth-Email: your-email@company.com" \
  -H "X-CH-Auth-API-Token: your-api-token" \
  -H "Content-Type: application/json" \
  -d '{"id": "019ab80d-f7a1-739b-99e0-803f501fdea9", "prompt": "Show me alerts from the last 24 hours"}'
```

### Session Status Values

| Status | Description |
| --- | --- |
| `SESSION_STATUS_PENDING` | Request received, processing not yet started |
| `SESSION_STATUS_PROCESSING` | AI is actively analyzing and generating response |
| `SESSION_STATUS_COMPLETED` | Analysis complete, final answer available |
| `SESSION_STATUS_FAILED` | Processing failed, check `errorMessage` field |

## Polling Pattern

Since AI Advisor processes requests asynchronously, you must poll the **Get Chat Session** endpoint to retrieve results.

### Recommended Polling Strategy

```
import time
import requests

def wait_for_completion(session_id, headers, max_wait=60, interval=2):
    """Poll until session completes or timeout"""
    elapsed = 0
    while elapsed < max_wait:
        response = requests.get(
            f"https://grpc.api.kentik.com/ai_advisor/v202511/chat/{session_id}",
            headers=headers
        )
        data = response.json()

        if data["status"] in ["SESSION_STATUS_COMPLETED", "SESSION_STATUS_FAILED"]:
            return data

        time.sleep(interval)
        elapsed += interval

    raise TimeoutError("Session did not complete in time")
```

### Polling Guidelines

* Start polling immediately after creating/updating a session
* Use 2 second intervals for responsive results
* Do not use less than 2 second interval
* Typical completion time: 5-30 seconds depending on query complexity
* Set reasonable timeout (60-120 seconds recommended)

## Response Fields

### ChatMessage Object

| Field | Type | Description |
| --- | --- | --- |
| `id` | string | Unique message identifier (UUID) |
| `status` | string | Message processing status |
| `prompt` | string | Original user question |
| `finalAnswer` | string | AI-generated response (Markdown formatted) |
| `reasoning` | string | AI's internal reasoning process |
| `errorMessage` | string | Error details if status is FAILED |
| `createdAt` | timestamp | Message creation time |
| `updatedAt` | timestamp | Last update time |

## Workflow Example

```
# Step 1: Create session
SESSION_ID=$(curl -s -X POST https://grpc.api.kentik.com/ai_advisor/v202511/chat \
  -H "X-CH-Auth-Email: user@company.com" \
  -H "X-CH-Auth-API-Token: abc123..." \
  -H "Content-Type: application/json" \
  -d '{"prompt": "Show top 10 interfaces by utilization"}' \
  | jq -r '.id')

echo "Session ID: $SESSION_ID"

# Step 2: Poll for completion
while true; do
  RESPONSE=$(curl -s -X GET \
    "https://grpc.api.kentik.com/ai_advisor/v202511/chat/$SESSION_ID" \
    -H "X-CH-Auth-Email: user@company.com" \
    -H "X-CH-Auth-API-Token: abc123...")

  STATUS=$(echo $RESPONSE | jq -r '.status')

  if [ "$STATUS" = "SESSION_STATUS_COMPLETED" ]; then
    echo $RESPONSE | jq '.messages[0].finalAnswer'
    break
  elif [ "$STATUS" = "SESSION_STATUS_FAILED" ]; then
    echo "Error:" $(echo $RESPONSE | jq '.messages[0].errorMessage')
    break
  fi

  sleep 2
done

# Step 3: Ask follow-up question
curl -X PUT https://grpc.api.kentik.com/ai_advisor/v202511/chat \
  -H "X-CH-Auth-Email: user@company.com" \
  -H "X-CH-Auth-API-Token: abc123..." \
  -H "Content-Type: application/json" \
  -d '{"id": "$SESSION_ID", "prompt": "What about errors on those interfaces?"}'
```

## Rate Limiting

API requests are subject to additional rate limiting policies. The current rate limits are as follows:

* **Create Chat Session** and **Update Chat Session:**

  + 4 requests per minute (current limit)
  + 60 requests per hour (will be enforced in the future)
  + 120 requests per day (will be enforced in the future)
* **Get Chat Session:**

  + 60 requests per minute (current limit)

## Resources

* **AI Advisor OpenAPI Spec (GitHub)**
* **AI Advisor KB Article**
* **Kentik APIs Overview**
* **All Kentik API Specs (GitHub)**

---

### FAQs

#### What is the AI Advisor REST API used for?

The AI Advisor REST API provides programmatic access to Kentik's AI-powered network analysis capabilities, allowing users to create chat sessions and retrieve AI-generated insights.

#### How do I authenticate API requests?

All API requests require authentication using HTTP headers: `X-CH-Auth-Email` and `X-CH-Auth-API-Token`.

#### What should I do if the chat session status is SESSION\_STATUS\_FAILED?

If the session status is `SESSION_STATUS_FAILED`, check the `errorMessage` field for details about the failure.

#### What is the rate limit for creating chat sessions?

The current rate limit for creating and updating chat sessions is 4 requests per minute.

#### How do I retrieve the current status of a chat session?

You can retrieve the current status of a chat session using the `GET /ai_advisor/v202511/chat/{id}` endpoint.
