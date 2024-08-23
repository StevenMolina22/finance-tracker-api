# 🤑 FINANCE TRACKER PROJECT API

This API powers a finance tracker application that allows users to record their income, expenses, and manage their financial well-being.

## 🤖 Functionalities

**Implemented**

- **User Management:** Create, retrieve, update, and delete user accounts.
- **Transaction Management:** Track income and expenses with categories, amounts, and dates.
- **Category Management:** Define categories to organize transactions.
- **Asset & Liability Management:** Track assets and liabilities to calculate net worth.
- **Goal Setting:** Define saving, debt reduction, or investment goals with progress tracking.
- **Reporting & Analytics:** Generate reports for income vs expenses, category spending, net worth trends, and goal progress.

## 🎓 Data Models

- **Users:** Stores user information (id, username, email, etc.)
- **Transactions:** Records income and expenses (id, user_id, type, category_id, amount, date, etc.)
- **Categories:** Defines categories for transactions (id, name, description, etc.)
- **Assets:** Tracks user's assets (id, user_id, name, value, etc.)
- **Liabilities:** Tracks user's liabilities (id, user_id, name, value, etc.)
- **Goals:** Defines user's financial goals (id, user_id, type, amount, start_date, end_date, status, notes)
- **Goal Progress:** Tracks progress towards achieving goals (id, goal_id, date, amount_saved/paid/invested)

## 🦾 API Endpoints (Implemented)

### **User Management:**

- `GET /users`: List all users.
- `GET /users/:id`: Retrieve a specific user.
- `POST /users`: Create a new user.
- `PUT /users/:id`: Update a user.
- `DELETE /users/:id`: Delete a user.

### **Transaction Management:**

- `GET /transactions`: List all transactions for a user.
- `GET /transactions/:id`: Retrieve a specific transaction.
- `POST /transactions`: Create a new transaction.
- `PUT /transactions/:id`: Update a transaction.
- `DELETE /transactions/:id`: Delete a transaction.

### **Category Management:**

- `GET /categories`: List all categories.
- `GET /categories/:id`: Retrieve a specific category.
- `POST /categories`: Create a new category.
- `PUT /categories/:id`: Update a category.
- `DELETE /categories/:id`: Delete a category.

### **Asset Management:**

- `GET /assets`: List all assets for a user.
- `GET /assets/:id`: Retrieve a specific asset.
- `POST /assets`: Create a new asset.
- `PUT /assets/:id`: Update an asset.
- `DELETE /assets/:id`: Delete an asset.

### **Liability Management:**

- `GET /liabilities`: List all liabilities for a user.
- `GET /liabilities/:id`: Retrieve a specific liability.
- `POST /liabilities`: Create a new liability.
- `PUT /liabilities/:id`: Update a liability.
- `DELETE /liabilities/:id`: Delete a liability.

### **Goal Management:**

- `GET /goals`: List all goals for a user. (Optional)
- `GET /goals/:id`: Retrieve a specific goal.
- `POST /goals`: Create a new goal.
- `PUT /goals/:id`: Update a goal.
- `DELETE /goals/:id`: Delete a goal.

- `POST /goals/:id/progress`: Create goal progress entry.
- `GET /goals/:id/progress`: Retrieve goal progress data.

### **Reporting & Analytics:**

- `GET /reports/income_vs_expenses`: Get income vs expense report for a date range.
- `GET /reports/category_spending`: Get category spending report for a date range.
- `GET /reports/net_worth`: Get net worth trend report for a date range.
- `GET /reports/goal_progress`: Get progress report for a specific goal.

## ✅ TODO

- Implement `GET /goals` endpoint to retrieve all goals for a user.
- Implement `GET /goals/:id` endpoint to retrieve a specific goal.
- Implement `POST /goals` endpoint to create a new goal.
- Implement `PUT /goals/:id` endpoint to update a goal.
- Implement `DELETE /goals/:id` endpoint to delete a goal.
- Implement `POST /goals/:id/progress` endpoint to create goal progress entry.
- Implement `GET /goals/:id/progress` endpoint to retrieve goal progress data.

### Roadmap

1. **Define Goal Model:** Create a data model for goals with necessary fields (id, user_id, type, amount, start_date, end_date, status, notes).
2. **Define Goal Progress Model:** Create a data model for goal progress with necessary fields (id, goal_id, date, amount_saved/paid/invested).
3. **Implement Goal CRUD Operations:** Create, read, update, and delete operations for goals.
4. **Implement Goal Progress Creation:** Allow users to add progress entries for goals.
5. **Implement Goal Progress Retrieval:** Retrieve progress data for a specific goal.
6. **Test and Refine:** Thoroughly test all endpoints and handle potential errors.

## Reporting & Analytics

### TODO

- Implement `GET /reports/income_vs_expenses` endpoint to generate income vs expense report.
- Implement `GET /reports/category_spending` endpoint to generate category spending report.
- Implement `GET /reports/net_worth` endpoint to generate net worth trend report.
- Implement `GET /reports/goal_progress` endpoint to generate goal progress report.

### Roadmap

1. **Data Aggregation:** Develop functions to calculate totals, averages, and other relevant metrics.
2. **Report Structure:** Define report formats (JSON, CSV) and include necessary data fields.
3. **Income vs Expenses Report:** Implement calculations for income and expenses, and generate the report.
4. **Category Spending Report:** Calculate spending per category and generate the report.
5. **Net Worth Report:** Calculate net worth based on assets and liabilities, and generate the report.
6. **Goal Progress Report:** Retrieve goal progress data and calculate progress percentage.
7. **Data Filtering:** Allow users to filter reports by date range or other criteria.
8. **Testing and Refinement:** Test report generation and accuracy, and optimize performance.

**Additional Notes:**

- All endpoints support filtering and pagination (to be implemented).
- Authentication and authorization mechanisms are not included in this documentation (to be implemented).

## Project Resources

- **Excalidraw:** Link to Excalidraw project for visualizing app flow (if applicable).
- **FinanceTracker (Optional):** Link to the deployed finance tracker application (if available).

## Development Status

This project is actively under development. Check the "TODO" section for
