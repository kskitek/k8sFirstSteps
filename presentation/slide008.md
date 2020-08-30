
  Super simple "complex" architecture

┌───────────┐  POST /generate ┌───────────┐
│  Service A│---------------->│  Service B│
│      :8080│<----------------│      :8080│
└───────────┘  POST /data     └───────────┘
