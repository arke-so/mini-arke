# Coding challenge

This is the coding challenge for the Backend Engineer role.

You'll work with `mini-arke`, a simplified version of our manufacturing management system built on our actual stack. Your task is to implement a CSV export feature for orders.

**The goal**: Get a sense of your coding skills, problem-solving approach, and how you work with existing codebases. We want to see how you'd tackle a real feature in our async/remote environment.

**Important**: Treat this as production code that needs to ship. Write clean, maintainable code that follows the patterns already in the project. Don't overthink it, just do what you'd normally do on the job!

**Time expectation**: Our appetite for this feature is 2 hours. We value your time and designed this to be completable within that timeframe. Please don't spend significantly more time on it, we need to see some code, but we don't want to take up too much of your time. When submitting, let us know roughly how long you spent (no penalty if you go over, we'd just like to know!) 🙌

### What you need to implement

Create a `GET /orders/export` endpoint that:

1. Returns orders in CSV format
2. Accepts optional query parameters:
    - `startDate` (format: `YYYY-MM-DD`) - Filter orders from this date
    - `endDate` (format: `YYYY-MM-DD`) - Filter orders until this date
    - `status` - Filter orders by status (e.g., `pending`, `completed`, `cancelled`)

The CSV should include the following columns:

```csv
Order ID,Customer Name,Customer Email,Order Date,Status,Total Amount
550e8400-e29b-41d4-a716-446655440000,John Doe,john@example.com,2024-11-15,completed,299.99
660e8400-e29b-41d4-a716-446655440001,Jane Smith,jane@example.com,2024-11-16,pending,150.50
```

## Deliverable

1. Fork this repository into a **private repository** named `mini-arke`
2. Implement the CSV export endpoint
3. Ensure `make test` passes
4. Add `@ilpes` as collaborator to your private repo
5. Send us the link to your private repository

## What comes next?
The next step will be a 30-minute technical session where we'll work together on top of the code you submitted. This could be expanding the feature, adding a related endpoint, writing tests, or something similar. Make sure to keep the project running on your local environment so we can dive right in.

After the technical portion, we'll spend some time getting to know each other better, learning about your background, what you're looking for in your next role, and answering any questions you have about the team and how we work.
We'll follow up with you for next steps!

## Questions?

If you have any questions about the assignment, feel free to reach out!

---

**We can't wait to see your submission! Good luck! 🚀**