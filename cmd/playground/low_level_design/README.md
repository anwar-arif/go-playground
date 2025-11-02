🧩 Stripe-Style Coding Problems

Here’s a list of realistic Stripe-style questions — some are inspired by actual interviews and others by similar companies (like Coinbase, Square, DoorDash):

🧠 1. Transaction Stream Validator

Problem:
You receive a stream of transactions, each represented as:

{ "user_id": "u123", "amount": 50.0, "timestamp": 1739001234 }


Implement a function:

func DetectFraudulentUsers(transactions []Transaction, threshold float64, window time.Duration) []string


Return user IDs that spent more than threshold within any rolling window (e.g., $1000 in 5 minutes).

✅ Tests data structures, window logic, and careful state tracking.

🧾 2. Invoice Generator

Given:

A list of purchases (amount, user ID, timestamp)

Discounts or tax rates

Implement:

func GenerateInvoices(purchases []Purchase, taxRate float64, discountRate float64) map[string]Invoice


Return invoices per user, computing totals correctly and rounding to 2 decimals.

✅ Tests data modeling, math precision, and grouping logic.

🧮 3. String Compression Tracker

Given a stream of strings, you must “localize” them into a compact form.

Example:

Input: "internationalization" → Output: "i18n"


Implement a function that compresses and decompresses such words:

type Localizer struct { ... }
func (l *Localizer) Compress(word string) string
func (l *Localizer) Decompress(short string) string


✅ Tests clean API design, correctness, and edge cases.

💳 4. Payment Processor Queue

Design a minimal in-memory priority queue that processes payments based on priority and timestamp.

type Payment struct {
ID        string
Priority  int
Timestamp time.Time
}

type PaymentQueue struct { ... }

func (pq *PaymentQueue) Push(p Payment)
func (pq *PaymentQueue) Pop() Payment


✅ Tests concurrency safety and use of container/heap.

🔄 5. Currency Conversion Engine

Given a set of currency exchange rates:

USD → EUR = 0.9
EUR → GBP = 0.8
GBP → INR = 100


Implement:

func Convert(from, to string, amount float64, rates []Rate) (float64, error)


Support multi-hop conversions (USD→EUR→GBP→INR).

✅ Tests graph traversal (DFS/BFS) + error handling.

🧰 6. Log Aggregator

Implement a simple log aggregator that stores logs from multiple services, supports filtering, and provides statistics.

type Log struct {
Service string
Level   string
Message string
Time    time.Time
}

type Aggregator struct { ... }

func (a *Aggregator) AddLog(log Log)
func (a *Aggregator) Filter(service, level string) []Log
func (a *Aggregator) CountByService() map[string]int


✅ Tests clean struct design and modularity.

🧾 7. Receipt Reconciliation

Given a CSV of receipts and another CSV of bank statements, detect mismatches.

func FindMismatches(receipts, statements []Transaction) []string


✅ Tests attention to detail, string/number parsing, and clean error reporting.

⚙️ 8. Rate Limiter

Design a rate limiter:

func AllowRequest(user string) bool


Allow each user N requests per minute, sliding window style.

✅ Tests system design + time-based logic.

💡 Tips for Stripe-style questions

Structure your code cleanly.
Separate input parsing, core logic, and output.

Write helper types (like Transaction, Invoice, etc.) with proper names.

Handle errors explicitly.
Return error, don’t panic.

Add small tests.
Even in interviews, Stripe loves if you write a few test cases quickly.

Think in systems.
The problem may seem small, but they want to see production-grade thinking.