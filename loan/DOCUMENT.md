# Document

1. Tech Stack:

    - Docker Compose: Simplifies containerization and environment setup
    - PostgreSQL: Reliable and robust relational database
    - Golang: Efficient, high-performance language suitable for backend services
    - Gorm: Simplifies database operations with Go

2. Endpoint

    1. Guest

        - POST `/api/auth/sign-up`
        - POST `/api/auth/login`
        - GET `/api/loan`: to show list approved and invested loan

    2. User

        - DELETE `/api/user/auth`: for logout
        - POST `/api/user/loan/propose`: user can propose for loan
        - POST `/api/user/loan/:loanId/invest`: user can invest a loan

    3. Admin

        - GET `/api/admin/loan`: to show list loan
        - POST `/api/admin/loan/:loanId/approve`: to approve a loan
        - POST `/api/admin/loan/:loanId/disburse`: to disburse a loan

3. Test Cases

    Authentication Sign Up

    - Missing required field (Bad Request)
    - Duplicate sign up email (Conflict)
    - Valid sign up (No Content)

    Authentication Login

    - Invalid login (Unauthorized)
    - Invalid credential (Unauthorized)
    - Valid login (Ok)

    User Authentication Logout

    - Invalid credential (Unauthorized)
    - Valid logout (No Content)

    User Loan Propose

    - Invalid credential (Unauthorized)
    - Missing required field (Bad Request)
    - Propose (No Content)

    Admin Loan Approve

    - Invalid credential (Unauthorized)
    - Missing required field (Bad Request)
    - Invalid loan (Not Found)
    - Invalid loan status (Conflict)
    - Approve (No Content)

    User Loan Invest

    - Invalid credential (Unauthorized)
    - Missing required field (Bad Request)
    - Invalid loan (Not Found)
    - Invalid loan status (Conflict)
    - Loan already fullfilled (Not Acceptable)
    - Invest amount invalid (Not Acceptable)
    - Invest (No Content)

    Admin Loan Disbursement

    - Invalid credential (Unauthorized)
    - Missing required field (Bad Request)
    - Invalid loan (Not Found)
    - Invalid loan status (Conflict)
    - Disbursement (No Content)
