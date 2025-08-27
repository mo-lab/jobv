<div dir=rtl>
# jobv API Documentation

این راهنمای مسیر های API هست

---

## Endpoints

### 1. Send OTP

- **URL:** `/otp/send`
- **Method:** `GET`
- **Description:** یک درخواست ازنوع GET به همراه پارامتر phone و مقدار از نوع string ارسال می شود و نتیجه در کنسول چاپ شده و به شما نمایش داده میشود
- **Request:** *http://localhost:8080/otp/send?phone=1234*
- **Response:** *OTP sent successfully*
- **Handler:** `SendOTPHandler`

---

### 2. Verify OTP

- **URL:** `/otp/verify`
- **Method:** `POST`
- **Description:** کد نمایش داده شده به همراه تلفن با فرمت json به این اند پوینت ارسال شده و مقدار توکن jwt در پاسخ با فرمت json برگشت داده میشود
- **Request Body:** `{"phone":"1234","code":"99999"}`
- **Response:** `{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTYyNTA0NjUsInJvbGUiOiJ1c2VyIiwidWlkIjoiIn0.HgBTZKf9IixHmSpnlyFSOE4M5FJ7Rv-JmLMWDTwIVGY"
}`
- **Handler:** `VerifyOTPHandler`

---

### 3. Get Users

- **URL:** `/users`
- **Method:** `GET`
- **Headers:** 
- **Description:** ق یک لیست از کاربران را با فرمت json دریافت میکنیم
- **Response:** [
    {
        "id": "",
        "phone": "1234",
        "created_at": "2025-08-26T18:49:20.14Z"
    }
]
- **Handler:** `GetUsersHandler`
- **Params:** `page` شماره صفحه و `limit` تعداد کاربران دریافتی را مشخص میکند

---

### 4. Search Users

- **URL:** `/users/search`
- **Method:** `GET`
- **Description:** Searches for users based on specified criteria.
- **Request Example:** *(http://localhost:8080/users/search?phone=123456&page=2&limit=10)*
- **Response:** *[
    {
        "id": "68aee7ca1328cf7d00b72cf3",
        "phone": "123456",
        "created_at": "0001-01-01T00:00:00Z"
    }
]*
- **Handler:** `SearchUsersHandler`
 **Params:** `page` شماره صفحه و `limit` تعداد کاربران دریافتی را مشخص میکند `phone`و `role` پارامتر های سرچ هستند
---

### 5. Protected Search Users

- **URL:** `/protected/users/search`
- **Method:** `GET`
- **Description:** میدلور jwt برای این اند پوینت فعال شده و با قرار دادن توکن دریافتی از اند پوینت `/verify`در هدر امکان دسترسی به آن امکان پذیر است
- **Request Header:** *"Authorization":  Bearer <YOUR_JWT_HERE>*
- **Response:** *(example needed)*
- **Handler:** `SearchUsersHandler`
 **Params:** `page` شماره صفحه و `limit` تعداد کاربران دریافتی را مشخص میکند `phone`و `role` پارامتر های سرچ هستند
---
</div>