
# Guardians of Privacy: Advanced Data Masking Challenge

## Problem Statement

You are provided with a string of text that contains users' personally identifiable information (PII) which includes email addresses, phone numbers, and credit card numbers. To maintain privacy, you need to mask part of each PII entry in the text. Specifically:
- For email addresses, replace the characters between the first character and the "@" symbol with asterisks (*).
- For phone numbers, replace the middle digits with asterisks, keeping the first two and the last two digits visible.
- For credit card numbers, replace all digits except the last four with asterisks.

## Input & Output

### Input Format

A single string of text containing email addresses, phone numbers, and credit card numbers.

### Input Constraints

- Email addresses follow the format: localpart@domain.
- Phone numbers are 10 digits long.
- Credit card numbers are 16 digits long.

### Output Format

A single string of text with masked PII.

## Examples

### Example 1

#### Input
```
alice@example.com 1234567890 1234567812345678
```

#### Output
```
a****e@example.com 12****90 ************5678
```

### Example 2

#### Input
```
john.doe@somecompany.com 9876543210 4321432143214321 jane_smith@anotherdomain.org 1231231234 1111222233334444 charles@domain.com 1111111111 5555666677778888
```

#### Output
```
j*******e@somecompany.com 98****10 ************4321 j********h@anotherdomain.org 12****34 ************4444 c*****s@domain.com 11****11 ************8888
```

## User Task

Your task is to complete the function `mask_pii()` and return the expected answer. You don't need to read the input, just write code within the function block directly.
